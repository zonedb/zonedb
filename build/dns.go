package build

import (
	"os"

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
