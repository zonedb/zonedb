package build

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"net"
	"os"
	"slices"
	"sort"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/miekg/dns"
	"github.com/wsxiaoys/terminal/color"
	"golang.org/x/net/idna"
	"golang.org/x/sync/errgroup"
)

const rootZoneURL = "https://www.internic.net/domain/root.zone"

// FetchRootZone fetches the IANA root zone file and adds name servers to zones.
func FetchRootZone(ctx context.Context, zones map[string]*Zone, addNew bool, cache *ETagCache) error {
	res, err := FetchWithETag(ctx, rootZoneURL, cache)
	if err != nil {
		return err
	}
	if res == nil {
		return nil // 304 Not Modified
	}
	defer res.Body.Close()

	color.Fprintf(os.Stderr, "@{.}Parsing %s\n", rootZoneURL)

	// Clear out old name servers for TLDs
	for _, z := range zones {
		if !z.IsTLD() {
			continue
		}
		z.oldNameServers = z.NameServers
		z.NameServers = nil
	}

	zp := dns.NewZoneParser(res.Body, "", "")
	var rrs []dns.RR

	for rr, ok := zp.Next(); ok; rr, ok = zp.Next() {
		h := rr.Header()
		if h.Rrtype != dns.TypeNS {
			continue
		}
		d := Normalize(h.Name)
		if d != "" {
			rrs = append(rrs, rr)
		}
	}

	if err := zp.Err(); err != nil {
		return fmt.Errorf("parsing root zone: %w", err)
	}

	color.Fprintf(os.Stderr, "@{.}Verifying %d RRs in root zone...\n", len(rrs))

	g, gctx := errgroup.WithContext(ctx)
	g.SetLimit(Concurrency)

	for _, rr := range rrs {
		d := Normalize(rr.Header().Name)

		// Identify the zone
		z := zones[d]
		if z == nil {
			if !addNew {
				continue
			}
			color.Fprintf(os.Stderr, "@{g}New domain: %s\n", d)
			z = &Zone{Domain: d}
			zones[d] = z
		}

		// Extract name server
		if ns, ok := rr.(*dns.NS); ok {
			host := Normalize(ns.Ns)
			g.Go(func() error {
				if slices.Contains(z.oldNameServers, host) || verifyNS(gctx, host) == nil {
					z.mut.Lock()
					z.NameServers = append(z.NameServers, host)
					z.mut.Unlock()
				} else if gctx.Err() != nil {
					// Surface group cancellation; per-host verifyNS failures still return nil.
					return gctx.Err()
				}
				return nil
			})
		}
	}

	if err := g.Wait(); err != nil {
		return err
	}

	// Detect retired or withdrawn TLDs
	for _, z := range zones {
		if !z.IsTLD() {
			continue
		}
		if len(z.NameServers) == 0 && len(z.oldNameServers) > 0 {
			color.Fprintf(os.Stderr, "@{y}TLD no longer present in root.zone: @{y!}%s@{y}\n", z.Domain)
			// z.Retire()
		}
	}

	return nil
}

// FetchNameServers fetches NS records for zones.
func FetchNameServers(ctx context.Context, zones, allZones map[string]*Zone) error {
	color.Fprintf(os.Stderr, "@{.}Fetching name servers for %d zone(s)...\n", len(zones))
	var nx sync.Map
	var found, added, removed, skipped, failed, unresolved int32
	err := mapZones(ctx, zones, func(gctx context.Context, z *Zone) error {
		// Skip TLDs
		if z.IsTLD() {
			return nil
		}

		// Get parent
		parent := allZones[z.ParentDomain()]
		if parent == nil {
			color.Fprintf(os.Stderr, "@{r}Error: parent zone %s does not exist for %s\n", z.ParentDomain(), z.Domain)
			return nil
		}

		// Copy parent name servers
		parent.mut.Lock()
		parentNS := make([]string, len(parent.NameServers))
		copy(parentNS, parent.NameServers)
		parent.mut.Unlock()

		// Clear out existing name servers
		z.oldNameServers = z.NameServers
		z.NameServers = nil

		// Iterate over name servers
		counts := make(map[string]int, 8)
		for _, host := range parentNS {
			// Short-circuit on cancellation. Without this, the callback would fall
			// off the end returning nil and errgroup wouldn't see the cancellation.
			if gctx.Err() != nil {
				return gctx.Err()
			}
			if _, ok := nx.Load(host); ok {
				continue
			}
			qctx, cancel := context.WithTimeout(gctx, 30*time.Second)
			rmsg, err := exchange(qctx, host, z.ASCII(), dns.TypeNS)
			cancel()
			if err != nil {
				var to timeouter
				if errors.As(err, &to) && to.Timeout() {
					Trace("@{.}Timeout querying %s for @{!}%s@{.}: %s\n", host, z, err.Error())
					continue
				}
				color.Fprintf(os.Stderr, "@{r}Error querying %s for @{r!}%s@{r}: %s\n", host, z, err.Error())

				// Cache host not found
				if strings.Contains(err.Error(), "no such host") {
					if _, ok := nx.LoadOrStore(host, struct{}{}); !ok {
						atomic.AddInt32(&unresolved, 1)
					}
				}
				// miekg/dns doesn't always wrap to *net.DNSError, so the
				// string match above is used instead. Revisit if that changes.
				//
				// var derr *net.DNSError
				// if errors.As(err, &derr) && derr.IsNotFound {
				// 	if _, ok := nx.LoadOrStore(host, struct{}{}); !ok {
				// 		atomic.AddInt32(&unresolved, 1)
				// 		Trace("@{y}! ")
				// 	}
				// }
				continue
			}
			if rmsg.Rcode == dns.RcodeNameError {
				// Expected. Many parent NS don't serve the child zone authoritatively. Uncomment to debug.
				//
				// color.Fprintf(os.Stderr, "@{y}Warning: %s returned NXDOMAIN for %s (NS)\n", host, z.Domain)
				continue
			}
			for _, rr := range append(rmsg.Answer, rmsg.Ns...) {
				if ns, ok := rr.(*dns.NS); ok {
					v := Normalize(ns.Ns)
					counts[v] = counts[v] + 1
					// Uncomment to trace which parent NS returned which child NS record.
					//
					// color.Fprintf(os.Stderr, "@{.}DNS record for %s: %s\n", z.Domain, v)
				}
			}
		}

		// Get max (consensus) count
		max := 0
		for _, count := range counts {
			if count > max {
				max = count
			}
		}
		// Partial consensus is normal during delegation updates; the per-NS
		// voting below handles it. Uncomment to investigate count mismatches.
		//
		// if max > 0 && max < len(parentNS) {
		// 	color.Fprintf(os.Stderr, "@{y}Warning: inconsistent DNS responses (%d < %d) for @{y!}%s\n", max, len(parentNS), z.Domain)
		// }

		// Store new name servers
		for ns, count := range counts {
			// Ignore non-consensus values where > 1 name server does not return ns.
			// FIXME: this criteria is subjective.
			if count == 1 && max > 2 {
				color.Fprintf(os.Stderr, "@{y}Warning: non-consensus name server for %s: %s (%d < %d)\n", z, ns, count, max)
				atomic.AddInt32(&skipped, 1)
				continue
			}
			// Disabled: many valid NS don't respond to unsolicited UDP (firewalls),
			// producing false negatives. See VerifyNameServers for the offline path.
			//
			// if verifyNS(ns) != nil {
			// 	color.Fprintf(os.Stderr, "@{y}Warning: unable to verify name server for %s: %s\n", z.Domain, ns)
			// 	atomic.AddInt32(&failed, 1)
			// 	continue
			// }
			z.NameServers = append(z.NameServers, ns)
			atomic.AddInt32(&found, 1)
		}

		// Sanity check
		if len(z.NameServers) == 0 && len(z.oldNameServers) > 0 {
			color.Fprintf(os.Stderr, "@{y}Zone lost all name servers: @{y!}%s@{y}\n", z)
		}

		// Record changes
		olds := NewSet(z.oldNameServers...)
		news := NewSet(z.NameServers...)
		for ns := range news {
			if !olds[ns] {
				atomic.AddInt32(&added, 1)
			}
		}
		for ns := range olds {
			if !news[ns] {
				atomic.AddInt32(&removed, 1)
			}
		}
		return nil
	})
	color.Fprintf(os.Stderr, "@{.}\nFound %d name servers, %d added, %d removed, %d non-consensus, %d failed, %d NXDOMAIN\n", found, added, removed, skipped, failed, unresolved)
	return err
}

// VerifyNameServers verifies the name servers for zones.
// Per-zone verification failures are logged and the failing entries dropped.
func VerifyNameServers(ctx context.Context, zones map[string]*Zone) error {
	color.Fprintf(os.Stderr, "@{.}Verifying name servers for %d zones...\n", len(zones))
	return mapZones(ctx, zones, func(gctx context.Context, z *Zone) error {
		var nameServers []string
		for _, ns := range z.NameServers {
			if err := verifyNS(gctx, ns); err != nil {
				if gctx.Err() != nil {
					return gctx.Err()
				}
				LogWarning(fmt.Errorf("can’t verify name server %s: %s", ns, err))
			} else {
				nameServers = append(nameServers, ns)
			}
		}
		z.NameServers = nameServers
		return nil
	})
}

func verifyNS(ctx context.Context, host string) error {
	host, err := idna.ToASCII(Normalize(host))
	if err != nil {
		return fmt.Errorf("normalizing host %q: %w", host, err)
	}
	return CanDial(ctx, "udp", host+":53")
}

// CountNameServers counts unique name servers for zones.
func CountNameServers(ctx context.Context, zones map[string]*Zone) error {
	var found int32
	var mu sync.Mutex
	all := NewSet()
	err := mapZones(ctx, zones, func(_ context.Context, z *Zone) error {
		if len(z.NameServers) > 0 {
			atomic.AddInt32(&found, 1)
			for _, tag := range z.Tags {
				if tag == TagRetired || tag == TagWithdrawn {
					color.Fprintf(os.Stderr, "@{y}Retired/withdrawn zone with active name servers: %s\n", z)
				}
			}
		}
		for _, ns := range z.NameServers {
			mu.Lock()
			all.Add(ns)
			mu.Unlock()
		}
		return nil
	})
	// Uncomment to enumerate every unique nameserver.
	//
	// for ns := range all {
	// 	color.Fprintf(os.Stderr, "@{.}%s ", ns)
	// }
	// color.Fprintf(os.Stderr, "\n")
	color.Fprintf(os.Stderr, "@{.}Counted %d unique name servers in %d of %d zone(s)\n", len(all), found, len(zones))
	return err
}

// FindWildcards finds wildcard DNS records for a zone.
func FindWildcards(ctx context.Context, zones map[string]*Zone) error {
	color.Fprintf(os.Stderr, "@{.}Finding wildcards for %d zones...\n", len(zones))
	var found int32
	var mu sync.Mutex
	all := NewSet()
	err := mapZones(ctx, zones, func(gctx context.Context, z *Zone) error {
		// Zero out wildcards
		z.Wildcards = nil

		// Try n random domains unlikely to be registered
		const n = 8
		var resolved int
		addrs := NewSet()
		for i := 0; i < n; i++ {
			// Short-circuit on cancellation. Without this, the loop would finish
			// and return nil; errgroup wouldn't see the cancellation.
			if gctx.Err() != nil {
				return gctx.Err()
			}
			if i > 2 && len(addrs) == 0 {
				break
			}
			qctx, cancel := context.WithTimeout(gctx, 10*time.Second)
			rand := randLabel(32)
			name := rand + "." + z.ASCII()
			cname, err := resolver.LookupCNAME(qctx, name)
			if err == nil && cname != "" {
				// Skip self-referential CNAMEs (cname contains the random label).
				// That's a normal wildcard shape. Uncomment to trace matches.
				//
				// if strings.Contains(cname, rand) {
				// 	color.Fprintf(os.Stderr, "@{y}Wildcard matching random name: @{w}%s@{.} ← @{c}%s@{.}\n", cname, z.Domain)
				// }
				if !strings.Contains(cname, rand) {
					addrs.Add(Normalize(cname))
				}
			}
			ipaddrs, err := resolver.LookupIPAddr(qctx, name)
			// Explicit cancel, not defer: deferred cancels would accumulate across n=8 iterations.
			cancel()
			if err != nil {
				continue
			}
			for _, ipaddr := range ipaddrs {
				addr := ipaddr.IP.String()
				// Ignore ICANN name collisions
				// https://www.icann.org/resources/pages/name-collision-2013-12-06-en#127.0.53.53
				if addr == "127.0.53.53" {
					continue
				}
				addrs.Add(addr)
				resolved++
			}
		}

		// Must have resolved at least n/2 domains
		if resolved < n/2 {
			return nil
		}

		// Record unique addrs
		atomic.AddInt32(&found, 1)
		mu.Lock()
		all.Add(addrs.Values()...)
		mu.Unlock()
		z.Wildcards = addrs.Values()
		return nil
	})

	// Make reverse mapping
	rev := make(map[string][]string)
	for _, z := range zones {
		for _, w := range z.Wildcards {
			rev[w] = append(rev[w], z.Domain)
		}
	}

	// Print mapping
	sorted := all.Values()
	sort.Strings(sorted)
	for _, addr := range sorted {
		color.Fprintf(os.Stderr, "@{.}Found wildcard: @{w}%s@{.} ← @{c}%s@{.}\n", addr, strings.Join(rev[addr], " "))
	}

	color.Fprintf(os.Stderr, "@{.}Found %d unique wildcard addresses for %d zone(s) (%d not wildcarded)\n",
		len(all), found, int32(len(zones))-found)
	return err
}

func randLabel(n int) string {
	const ascii = "0123456789abcdefghijklmnopqrstuvwxyz"
	b := make([]byte, n)
	for i := range b {
		b[i] = ascii[rand.Int63()%int64(len(ascii))]
	}
	return string(b)
}

func exchange(ctx context.Context, host, qname string, qtype uint16) (*dns.Msg, error) {
	qmsg := &dns.Msg{}
	qmsg.RecursionDesired = true
	qmsg.SetQuestion(dns.Fqdn(qname), qtype)
	client := &dns.Client{}
	// ExchangeContext already honors ctx; setting DialTimeout would be redundant.
	//
	// if deadline, ok := ctx.Deadline(); ok {
	// 	client.DialTimeout = time.Until(deadline)
	// }
	rmsg, _, err := client.ExchangeContext(ctx, qmsg, host+":53")
	if err == nil {
		return rmsg, err
	}
	var derr *net.DNSError
	if errors.As(err, &derr) {
		return nil, err
	}
	client.Net = "tcp" // Retry with TCP
	rmsg, _, err = client.ExchangeContext(ctx, qmsg, host+":53")
	return rmsg, err
}

type timeouter interface {
	Timeout() bool
}

var resolver = &net.Resolver{
	PreferGo:     true,
	StrictErrors: true,
}
