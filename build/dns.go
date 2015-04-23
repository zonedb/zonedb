package build

import (
	"os"
	"sync"
	"sync/atomic"

	"github.com/domainr/dnsr"
	"github.com/miekg/dns"
	"github.com/wsxiaoys/terminal/color"
)

const rootZoneURL = "http://www.internic.net/domain/root.zone"

func FetchRootZone(zones map[string]*Zone, addNew bool) error {
	res, err := Fetch(rootZoneURL)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	seen := NewSet()
	color.Fprintf(os.Stderr, "@{.}Parsing %s\n", rootZoneURL)
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

		// Nuke any preexisting name servers
		if !seen[d] {
			seen.Add(d)
			z.NameServers = nil
		}

		// Extract name server
		if ns, ok := token.RR.(*dns.NS); ok {
			z.NameServers = append(z.NameServers, Normalize(ns.Ns))
		}
	}

	return nil
}

var resolver = dnsr.New(10000)

func FetchNameServers(zones map[string]*Zone) error {
	color.Fprintf(os.Stderr, "@{.}Fetching name servers for %d zones...\n", len(zones))
	var found int32
	limiter := make(chan struct{}, Concurrency)
	var wg sync.WaitGroup
	for _, z := range zones {
		limiter <- struct{}{}
		wg.Add(1)
		go func(z *Zone) {
			defer func() {
				<-limiter
				wg.Done()
			}()
			name := z.ACE()
			rrs := resolver.Resolve(name, "NS")
			for _, rr := range rrs {
				if rr.Type != "NS" || Normalize(rr.Name) != z.Domain {
					continue
				}
				z.NameServers = append(z.NameServers, Normalize(rr.Value))
				atomic.AddInt32(&found, 1)
			}
		}(z)
	}
	wg.Wait()
	color.Fprintf(os.Stderr, "@{.}Found %d name servers\n", found)
	return nil
}
