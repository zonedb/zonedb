package build

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/wsxiaoys/terminal/color"
)

const ianaRDAPURL = "https://data.iana.org/rdap/dns.json"

// FetchRDAPFromIANA retrieves the map of zones to RDAP service endpoints from IANA.
func FetchRDAPFromIANA(zones map[string]*Zone) error {
	res, err := Fetch(ianaRDAPURL)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// The IANA bootstrap RDAP schema is an array of pairs with the first element a list of
	// domains and the second element a list of RDAP service endpoints for those domains, e.g.
	// {
	//   ...
	//   "services": [
	//     [
	//	  ["foo", "bar", "baz"], ["https://example.com/rdap"]
	//     ],
	//     ...
	//   ],
	//   "version": "1.0"
	// }
	x := struct {
		Services [][][]string // I'm screaming at this schema
	}{}
	d := json.NewDecoder(res.Body)
	err = d.Decode(&x)
	if err != nil {
		return err
	}
	if len(x.Services) == 0 {
		// An empty mapping implies bad data or a version change, best stop
		return fmt.Errorf("empty RDAP services array")
	}

	for _, svc := range x.Services {
		domains := svc[0]

		for _, domain := range domains {
			z, ok := zones[domain]
			if !ok {
				color.Fprintf(os.Stderr, "@{y}domain %s not found in zones map\n", domain)
				continue
			}

			z.RDAPURLs = []string{}
			for _, rdapURL := range svc[1] {
				if rdapURL == "" {
					continue
				}

				normURL := NormalizeURL(rdapURL)
				if normURL == "" {
					Trace("@{r}invalid RDAP URL %q\n", rdapURL)
					continue
				}
				z.RDAPURLs = append(z.RDAPURLs, normURL)
			}
		}
	}

	return nil
}

func AddRDAPURLs(zones map[string]*Zone, urls []string) {
	var added, modified int

	for _, z := range zones {
		oldAdded := added
	URLloop:
		for _, rdapURL := range urls {
			normURL := NormalizeURL(rdapURL)
			for _, zru := range z.RDAPURLs {
				if zru == normURL {
					// If the URL already exists on the zone then skip
					continue URLloop
				}
			}
			z.RDAPURLs = append(z.RDAPURLs, normURL)
			added++
		}
		if oldAdded < added {
			modified++
		}
	}

	color.Fprintf(os.Stderr, "@{.}Added %d RDAP URL(s) to %d zone(s)\n", added, modified)
}
