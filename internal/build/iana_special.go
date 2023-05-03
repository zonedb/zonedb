package build

import (
	"encoding/csv"
	"io"
	"os"
	"strings"

	"github.com/wsxiaoys/terminal/color"
)

const (
	// https://www.iana.org/assignments/special-use-domain-names/special-use-domain-names.xhtml
	ianaSpecialUseURL      = "https://www.iana.org/assignments/special-use-domain-names/special-use-domain.csv"
	ietfDataTrackerBaseURL = "https://datatracker.ietf.org/doc/"
)

// FetchSpecialUseDomainsFromIANA fetches special use domains from the IANA website.
func FetchSpecialUseDomainsFromIANA(zones map[string]*Zone, addNew bool) error {
	res, err := Fetch(ianaSpecialUseURL)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	r := csv.NewReader(res.Body)

	cols, err := r.Read()
	if err != nil {
		return err
	}

	var zonesAdded, zonesModified int
	for {
		row, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		rec := make(map[string]string)
		for i := range cols {
			rec[cols[i]] = row[i]
		}

		// Mangle reference into IETF data tracker URL
		ref := strings.Trim(rec["Reference"], "[]")
		ref = strings.Replace(ref, "RFC-ietf-", "draft-", 1)
		infoURL := ietfDataTrackerBaseURL + strings.ToLower(ref)

		domain := Normalize(rec["Name"])
		d := domain
		for d != "" {
			var added, modified bool
			z := zones[d]
			if z == nil {
				z = &Zone{Domain: d}
				if addNew {
					zones[d] = z
					added = true
				}
			}

			// Fix up metadata
			if z.RegistryOperator != "IANA" {
				z.RegistryOperator = "IANA"
				modified = true
			}
			if z.InfoURL != infoURL {
				z.InfoURL = infoURL
			}
			if z.AddTags("infrastructure", "closed") != 0 {
				modified = true
			}

			if added {
				zonesAdded++
				color.Fprintf(os.Stderr, "@{g}     New zone @{g!}%s@{g}\t%s\n", d, infoURL)
			} else if modified {
				zonesModified++
				color.Fprintf(os.Stderr, "@{g}Modified zone @{g!}%s@{g}\t%s\n", d, infoURL)
			}

			d = z.ParentDomain()
		}
	}

	color.Fprintf(os.Stderr, "@{.}Added %d and modified %d special domain (infrastructure) zone(s)\n", zonesAdded, zonesModified)

	return nil
}
