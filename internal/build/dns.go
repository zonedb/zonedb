package build

import (
	"fmt"
	"os"
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
