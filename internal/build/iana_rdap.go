package build

import (
	"encoding/json"
)

const (
	ianaRdapURL = "https://data.iana.org/rdap/dns.json"
)

// FetchRDAPFromIANA retrieves the map of zones to RDAP service endpoints from IANA.
func FetchRDAPFromIANA(zones map[string]*Zone) error {
	res, err := Fetch(ianaRdapURL)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	x := struct {
		Services [][][]string // I'm screaming at this schema
	}{}
	d := json.NewDecoder(res.Body)
	err = d.Decode(&x)
	if err != nil {
		return err
	}
	for _, svc := range x.Services {
		tlds := svc[0]

		for _, tld := range tlds {
			z, ok := zones[tld]
			if !ok {
				continue
			}

			z.RdapURLs = []string{}
			for _, rdapURL := range svc[1] {
				normURL := NormalizeURL(rdapURL)
				if normURL == "" {
					Trace("@{r}invalid RDAP URL %q\n", rdapURL)
					continue
				}
				z.RdapURLs = append(z.RdapURLs, normURL)
			}
		}
	}

	return nil
}
