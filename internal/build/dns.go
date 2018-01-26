package build

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/domainr/dnsr"
	"github.com/miekg/dns"
	"github.com/wsxiaoys/terminal/color"
	"golang.org/x/net/idna"
)

const rootZoneURL = "https://www.internic.net/domain/root.zone"

func FetchRootZone(zones map[string]*Zone, addNew bool) error {
	res, err := Fetch(rootZoneURL)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	color.Fprintf(os.Stderr, "@{.}Parsing %s\n", rootZoneURL)

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

	return nil
}

var resolver = dnsr.New(10000)

// FetchNameServers fetches NS records for zones.
func FetchNameServers(zones map[string]*Zone) error {
	color.Fprintf(os.Stderr, "@{.}Fetching name servers for %d zones...\n", len(zones))
	var found int32
	mapZones(zones, func(z *Zone) {
		name := z.ASCII()
		rrs := resolver.Resolve(name, "NS")
		for _, rr := range rrs {
			if rr.Type != "NS" || Normalize(rr.Name) != z.Domain {
				continue
			}
			ns := Normalize(rr.Value)
			if verifyNS(ns) == nil {
				z.NameServers = append(z.NameServers, ns)
				atomic.AddInt32(&found, 1)
			}
		}
	})
	color.Fprintf(os.Stderr, "@{.}Found %d name servers\n", found)
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
	color.Fprintf(os.Stderr, "@{.}Found %d unique DNS servers for %d zone(s)\n", len(all), found)
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
			name := randLabel(32) + "." + z.ASCII()
			rrs := resolver.Resolve(name, "A")
			for _, rr := range rrs {
				if rr.Type != "A" || Normalize(rr.Name) != name {
					continue
				}
				// Ignore ICANN name collisions
				// https://www.icann.org/resources/pages/name-collision-2013-12-06-en#127.0.53.53
				if rr.Value == dnsr.NameCollision {
					continue
				}
				addrs.Add(rr.Value)
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
