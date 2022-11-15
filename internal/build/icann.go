package build

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

const icannGTLDsURL = "https://www.icann.org/resources/registries/gtlds/v2/gtlds.json"

// FetchGTLDsFromICANN retrieves the list of gTLDs from ICANN.
func FetchGTLDsFromICANN(zones map[string]*Zone) error {
	res, err := Fetch(icannGTLDsURL)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	var x icannGTLDResponse
	d := json.NewDecoder(res.Body)
	err = d.Decode(&x)
	if err != nil {
		return err
	}

	if x.Version != 2 {
		return fmt.Errorf("unknown ICANN gTLDs list version: %d", x.Version)
	}

	var tagsAdded, tagsRemoved, zonesWithdrawn, zonesRetired, zonesModified int
	gTLDs := make(map[string]icannGTLD, len(x.GTLDs))

	// Iterate over ICANN-known gTLDs
	for _, g := range x.GTLDs {
		domain := Normalize(g.GTLD)
		gTLDs[domain] = g
		z, ok := zones[domain]
		if !ok {
			if len(zones) > 100 {
				Trace("@{y}Unknown domain in ICANN gTLD feed: %s\n", domain)
			}
			continue
		}

		// Compare ICANN uLabel to our normalized Unicode form
		if g.ULabel != "" && g.ULabel != domain {
			Trace("@{y}Warning: IDN U-label %s != @{y!}%s\n", g.ULabel, z.Domain)
		}

		var modified bool

		switch {
		// gTLD was withdrawn
		case g.DelegationDate.IsZero() && g.ContractTerminated:
			z.AddTags(TagWithdrawn)
			z.RemoveTags(TagRetired)
			zonesWithdrawn++
			modified = true

		// gTLD was retired
		case !g.RemovalDate.IsZero():
			z.RemoveTags(TagWithdrawn)
			z.AddTags(TagRetired)
			zonesRetired++
			modified = true

		// Pending delegation
		case g.DelegationDate.IsZero() && !g.ContractTerminated:
			n := len(z.NameServers)
			if n > 0 {
				Trace("@{y}Warning: undelegated gTLD with %d name servers: @{y!}%s\n", n, z.Domain)
			}
		}

		// Brand TLD?
		if g.Specification13 {
			d := z.AddTags(TagBrand)
			if d != 0 {
				tagsAdded += d
				modified = true
			}
		} else if !z.IsRetiredOrWithdrawn() && z.IsBrand() {
			// https://newgtlds.icann.org/en/applicants/agb/base-agreement-contracting/specification-13-applications
			Trace("@{y}Brand gTLD without ICANN Specification 13: @{y!}%s\n", z.Domain)
		}

		// Extract registry operator
		if g.RegistryOperator != "" {
			z.RegistryOperator = g.RegistryOperator
			modified = true
		}

		if modified {
			zonesModified++
		}
	}

	// Iterate over work zones
	for _, z := range zones {
		if !z.IsTLD() || z.IsRetiredOrWithdrawn() {
			continue
		}
		_, ok := gTLDs[z.Domain]
		if !ok && len(z.NameServers) == 0 {
			z.AddTags(TagWithdrawn)
		}
	}

	Trace("@{.}Added %d tag(s), removed %d tag(s) from %d zone(s)\n", tagsAdded, tagsRemoved, zonesModified)
	Trace("@{.}Withdrew %d and retired %d zone(s)\n", zonesWithdrawn, zonesRetired)

	return nil
}

type icannGTLDResponse struct {
	GTLDs []icannGTLD `json:"gTLDs"`
	// UpdatedOn time.Time `json:"updatedOn"` // Ignored because of nonstandard format
	Version int `json:"version"`
}

type icannGTLD struct {
	ApplicationID           string  `json:"applicationID"`
	ContractTerminated      bool    `json:"contractTerminated"`
	DateOfContractSignature ISODate `json:"dateOfContractSignature"`
	DelegationDate          ISODate `json:"delegationDate"`
	GTLD                    string  `json:"gTLD"`
	// RegistryClassDomainNameList   interface{} `json:"registryClassDomainNameList"` // (always null)
	RegistryOperator string `json:"registryOperator"`
	// RegistryOperatorCountryCode   *string     `json:"registryOperatorCountryCode"` // (always null)
	RemovalDate                   ISODate `json:"removalDate"`
	Specification13               bool    `json:"specification13"`               // Brand TLD
	ThirdOrLowerLevelRegistration bool    `json:"thirdOrLowerLevelRegistration"` // (always false or null)
	ULabel                        string  `json:"uLabel"`                        // Unicode IDN label
}

/*
{
	"gTLDs": [
		{
			"applicationId": "1-1386-27446",
			"contractTerminated": false,
			"dateOfContractSignature": "2015-02-26",
			"delegationDate": "2015-08-28",
			"gTLD": "aaa",
			"registryClassDomainNameList": null,
			"registryOperator": "American Automobile Association, Inc.",
			"registryOperatorCountryCode": null,
			"removalDate": null,
			"specification13": true,
			"thirdOrLowerLevelRegistration": false,
			"uLabel": null
		},
		{
			"applicationId": "1-1169-4534",
			"contractTerminated": false,
			"dateOfContractSignature": "2015-05-21",
			"delegationDate": "2015-11-03",
			"gTLD": "aarp",
			"registryClassDomainNameList": null,
			"registryOperator": "AARP",
			"registryOperatorCountryCode": null,
			"removalDate": null,
			"specification13": true,
			"thirdOrLowerLevelRegistration": false,
			"uLabel": null
		}
	],
	"updatedOn": "2021-07-16T15:54:03.596+0000",
	"version": 2
}
*/

type ISODate struct {
	time.Time
}

func (d *ISODate) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	if s == "null" {
		return nil
	}
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	*d = ISODate{t}
	return nil
}

// func (d ISODate) MarshalJSON() ([]byte, error) {
// 	return json.Marshal(d)
// }

// func (j ISODate) Format(s string) string {
// 	t := time.Time(j)
// 	return t.Format(s)
// }
