package build

import (
	"context"
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

	limiter := make(chan struct{}, Concurrency)
	var wg sync.WaitGroup
	for token := range dns.ParseZone(res.Body, "", "") {
		if token.Error != nil {
			continue
		}
		h := token.RR.Header()
		if h.Rrtype != dns.TypeNS {
			continue
		}
		d := Normalize(h.Name)
		if d == "" {
			continue
		}

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
		if ns, ok := token.RR.(*dns.NS); ok {
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
	color.Fprintf(os.Stderr, "@{.}Fetching name servers for %d zones\n", len(zones))
	var found, added, removed, skipped, failed int32
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

		// Clear out existing name servers
		z.oldNameServers = z.NameServers
		z.NameServers = nil

		// Source from Google, Cloudflare, and parent
		nameServers := []string{"8.8.8.8", "1.1.1.1"}
		nameServers = append(nameServers, parent.NameServers...)

		// Iterate over name servers
		counts := make(map[string]int, 8)
		for _, host := range nameServers {
			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()
			rmsg, err := exchange(ctx, host, z.ASCII(), dns.TypeNS)
			if err != nil {
				if to, ok := err.(timeouter); ok {
					if to.Timeout() {
						continue
					}
				}
				color.Fprintf(os.Stderr, "@{r}Error querying %s for @{r!}%s@{r}: %s\n", host, z.Domain, err.Error())
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
		// if max > 0 && max < len(parent.NameServers) {
		// 	color.Fprintf(os.Stderr, "@{y}Warning: inconsistent DNS responses (%d < %d) for @{y!}%s\n", max, len(parent.NameServers), z.Domain)
		// }

		// Store new name servers
		for ns, count := range counts {
			// Ignore non-consensus values (FIXME: this criteria is subjective)
			if count == 1 && max > 3 {
				color.Fprintf(os.Stderr, "@{y}Warning: non-consensus name server for %s: %s (%d < %d)\n", z.Domain, ns, count, max)
				atomic.AddInt32(&skipped, 1)
				continue
			}
			if verifyNS(ns) != nil {
				atomic.AddInt32(&failed, 1)
				continue
			}
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
	color.Fprintf(os.Stderr, "@{.}\nFound %d name servers, %d added, %d removed, %d non-consensus, %d failed\n", found, added, removed, skipped, failed)

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
	Trace("@{.}.")
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
				if tag == "retired" || tag == "withdrawn" {
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
	// 	color.Printf("@{.}%s ", ns)
	// }
	// color.Printf("\n")
	color.Fprintf(os.Stderr, "@{.}Counted %d unique name servers in %d of %d zone(s)\n", len(all), found, len(zones))
}

// FindWildcards finds wildcard DNS records for a zone.
func FindWildcards(zones map[string]*Zone) error {
	color.Fprintf(os.Stderr, "@{.}Finding wildcard IPs for %d zones...\n", len(zones))
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
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()
			name := randLabel(32) + "." + z.ASCII()
			rrs, _ := resolver.LookupIPAddr(ctx, name)
			for _, rr := range rrs {
				addr := rr.IP.String()
				// Ignore ICANN name collisions
				// https://www.icann.org/resources/pages/name-collision-2013-12-06-en#127.0.53.53
				if addr == "127.0.53.53" {
					continue
				}
				addrs.Add(addr)
				resolved++
			}
			if i > 1 && len(addrs) == 0 {
				break
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
	qmsg.MsgHdr.RecursionDesired = false
	qmsg.SetQuestion(dns.Fqdn(qname), qtype)
	client := &dns.Client{}
	rmsg, _, err := client.ExchangeContext(ctx, qmsg, host+":53")
	if err == nil {
		return rmsg, err
	}
	switch err := err.(type) {
	case *net.DNSError:
		return nil, err
		// case *dns.Error:
		// 	if err == dns.ErrTruncated {
		// 		return rmsg, nil // Use truncated msg
		// 	}
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
