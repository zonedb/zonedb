package build

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"net"
	"os"
	"sort"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/miekg/dns"
	"github.com/wsxiaoys/terminal/color"
	"golang.org/x/net/idna"
)

const rootZoneURL = "https://www.internic.net/domain/root.zone"

// FetchRootZone fetches the IANA root zone file and adds name servers to zones.
func FetchRootZone(zones map[string]*Zone, addNew bool) error {
	res, err := Fetch(rootZoneURL)
	if err != nil {
		return err
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
		return err
	}

	color.Fprintf(os.Stderr, "@{.}Verifying %d RRs in root zone...\n", len(rrs))

	limiter := make(chan struct{}, Concurrency)
	var wg sync.WaitGroup

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
			limiter <- struct{}{}
			wg.Add(1)
			go func(z *Zone, host string) {
				if verifyNS(host) == nil {
					z.m.Lock()
					z.NameServers = append(z.NameServers, Normalize(host))
					z.m.Unlock()
				}
				<-limiter
				wg.Done()
			}(z, ns.Ns)
		}
	}

	wg.Wait()

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
func FetchNameServers(zones, allZones map[string]*Zone) error {
	color.Fprintf(os.Stderr, "@{.}Fetching name servers for %d zone(s)...\n", len(zones))
	var nx sync.Map
	var found, added, removed, skipped, failed, unresolved int32
	mapZones(zones, func(z *Zone) {
		// Skip TLDs
		if z.IsTLD() {
			return
		}

		// Get parent
		parent := allZones[z.ParentDomain()]
		if parent == nil {
			color.Fprintf(os.Stderr, "@{r}Error: parent zone %s does not exist for %s\n", z.ParentDomain(), z.Domain)
			return
		}

		// Copy parent name servers
		parent.m.Lock()
		parentNS := make([]string, len(parent.NameServers))
		copy(parentNS, parent.NameServers)
		parent.m.Unlock()

		// Clear out existing name servers
		z.oldNameServers = z.NameServers
		z.NameServers = nil

		// Iterate over name servers
		counts := make(map[string]int, 8)
		for _, host := range parentNS {
			if _, ok := nx.Load(host); ok {
				continue
			}
			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()
			rmsg, err := exchange(ctx, host, z.ASCII(), dns.TypeNS)
			if err != nil {
				var to timeouter
				if errors.As(err, &to) && to.Timeout() {
					Trace("@{.}Timeout querying %s for @{!}%s@{.}: %s\n", host, z.Domain, err.Error())
					continue
				}
				color.Fprintf(os.Stderr, "@{r}Error querying %s for @{r!}%s@{r}: %s\n", host, z.Domain, err.Error())

				// Cache host not found
				if strings.Contains(err.Error(), "no such host") {
					if _, ok := nx.LoadOrStore(host, struct{}{}); !ok {
						atomic.AddInt32(&unresolved, 1)
					}
				}
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
				// color.Fprintf(os.Stderr, "@{y}Warning: %s returned NXDOMAIN for %s (NS)\n", host, z.Domain)
				continue
			}
			for _, rr := range append(rmsg.Answer, rmsg.Ns...) {
				if ns, ok := rr.(*dns.NS); ok {
					v := Normalize(ns.Ns)
					counts[v] = counts[v] + 1
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
		// if max > 0 && max < len(parentNS) {
		// 	color.Fprintf(os.Stderr, "@{y}Warning: inconsistent DNS responses (%d < %d) for @{y!}%s\n", max, len(parentNS), z.Domain)
		// }

		// Store new name servers
		for ns, count := range counts {
			// Ignore non-consensus values where > 1 name server does not return ns
			// FIXME: this criteria is subjective)
			if count == 1 && max > 2 {
				color.Fprintf(os.Stderr, "@{y}Warning: non-consensus name server for %s: %s (%d < %d)\n", z.Domain, ns, count, max)
				atomic.AddInt32(&skipped, 1)
				continue
			}
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
			color.Fprintf(os.Stderr, "@{y}Zone lost all name servers: @{y!}%s@{y}\n", z.Domain)
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
	})
	color.Fprintf(os.Stderr, "@{.}\nFound %d name servers, %d added, %d removed, %d non-consensus, %d failed, %d NXDOMAIN\n", found, added, removed, skipped, failed, unresolved)

	return nil
}

// VerifyNameServers verifies the name servers for zones.
func VerifyNameServers(zones map[string]*Zone) {
	color.Fprintf(os.Stderr, "@{.}Verifying name servers for %d zones...\n", len(zones))
	mapZones(zones, func(z *Zone) {
		var nameServers []string
		for _, ns := range z.NameServers {
			if err := verifyNS(ns); err != nil {
				LogWarning(fmt.Errorf("can’t verify name server %s: %s", ns, err))
			} else {
				nameServers = append(nameServers, ns)
			}
		}
		z.NameServers = nameServers
	})
}

func verifyNS(host string) error {
	host, err := idna.ToASCII(Normalize(host))
	if err != nil {
		return err
	}
	// Do long colored lines break GitHub Actions logging?
	// Trace("@{.}.")
	err = CanDial("udp", host+":53")
	return err
}

// CountNameServers counts unique name servers for zones.
func CountNameServers(zones map[string]*Zone) {
	var found int32
	var mu sync.Mutex
	all := NewSet()
	mapZones(zones, func(z *Zone) {
		if len(z.NameServers) > 0 {
			atomic.AddInt32(&found, 1)
			for _, tag := range z.Tags {
				if tag == TagRetired || tag == TagWithdrawn {
					color.Fprintf(os.Stderr, "@{y}Retired/withdrawn zone with active name servers: %s\n", z.Domain)
				}
			}
		}
		for _, ns := range z.NameServers {
			mu.Lock()
			all.Add(ns)
			mu.Unlock()
		}
	})
	// for ns := range all {
	// 	color.Fprintf(os.Stderr, "@{.}%s ", ns)
	// }
	// ccolor.Fprintf(os.Stderr, "\n")
	color.Fprintf(os.Stderr, "@{.}Counted %d unique name servers in %d of %d zone(s)\n", len(all), found, len(zones))
}

// FindWildcards finds wildcard DNS records for a zone.
func FindWildcards(zones map[string]*Zone) error {
	color.Fprintf(os.Stderr, "@{.}Finding wildcards for %d zones...\n", len(zones))
	var found int32
	var mu sync.Mutex
	all := NewSet()
	mapZones(zones, func(z *Zone) {
		// Zero out wildcards
		z.Wildcards = nil

		// Try n random domains unlikely to be registered
		const n = 8
		var resolved int
		addrs := NewSet()
		for i := 0; i < n; i++ {
			if i > 2 && len(addrs) == 0 {
				break
			}
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()
			rand := randLabel(32)
			name := rand + "." + z.ASCII()
			cname, err := resolver.LookupCNAME(ctx, name)
			if err == nil && cname != "" {
				// LookupCNAME returns name and nil error if no CNAME records are found.
				if strings.Contains(cname, rand) {
					// color.Fprintf(os.Stderr, "@{y}Wildcard matching random name: @{w}%s@{.} ← @{c}%s@{.}\n", cname, z.Domain)
				} else {
					addrs.Add(Normalize(cname))
				}
			}
			ipaddrs, err := resolver.LookupIPAddr(ctx, name)
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
			return
		}

		// Record unique addrs
		atomic.AddInt32(&found, 1)
		mu.Lock()
		all.Add(addrs.Values()...)
		mu.Unlock()
		z.Wildcards = addrs.Values()
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
	return nil
}

func randLabel(n int) string {
	const ascii = "0123456789abcdefghijklmnopqrstuvwxyz-"
	b := make([]byte, n)
	for i := range b {
		b[i] = ascii[rand.Int63()%int64(len(ascii))]
	}
	return string(b)
}

func exchange(ctx context.Context, host, qname string, qtype uint16) (*dns.Msg, error) {
	qmsg := &dns.Msg{}
	qmsg.MsgHdr.RecursionDesired = true
	qmsg.SetQuestion(dns.Fqdn(qname), qtype)
	client := &dns.Client{}
	if deadline, ok := ctx.Deadline(); ok {
		client.DialTimeout = time.Until(deadline)
	}
	// rmsg, _, err := client.ExchangeContext(ctx, qmsg, host+":53")
	// if err == nil {
	// 	return rmsg, err
	// }
	// var derr *net.DNSError
	// if errors.As(err, &derr) {
	// 	return nil, err
	// }
	client.Net = "tcp" // Retry with TCP
	rmsg, _, err := client.ExchangeContext(ctx, qmsg, host+":53")
	return rmsg, err
}

type timeouter interface {
	Timeout() bool
}

var resolver = &net.Resolver{
	PreferGo:     true,
	StrictErrors: true,
}
