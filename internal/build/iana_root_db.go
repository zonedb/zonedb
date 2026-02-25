package build

import (
	"fmt"
	"path"
	"sort"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const ianaRootDBURL = "https://www.iana.org/domains/root/db"

// idnToAsciiOverrides maps IDN ccTLD punycode domains to their ASCII ccTLD
// counterpart when the IANA root DB index page organisation names don't match
// exactly. This handles cases where:
// - The IDN and ASCII ccTLDs are managed by different organisations
// - Organisation names differ in capitalisation or punctuation
// - Multiple ASCII ccTLDs share the same organisation (TRA: .bh/.om)
//
// When adding new entries, include a comment with the IDN org and ASCII org
// for future reference.
var idnToAsciiOverrides = map[string]string{
	// India: "National Internet eXchange of India" vs "National Internet Exchange of India"
	"xn--2scrj9c":    "in", // .ಭಾರತ (Kannada)
	"xn--3hcrj9c":    "in", // .ଭାରତ (Odia)
	"xn--45br5cyl":   "in", // .ভাৰত (Assamese)
	"xn--h2breg3eve": "in", // .भारतम् (Sanskrit)
	"xn--h2brj9c8c":  "in", // .भारोत (Maithili)
	"xn--mgbbh1a":    "in", // .بارت (Kashmiri)
	"xn--mgbgu82a":   "in", // .ڀارت (Sindhi)
	"xn--rvc1e0am3e": "in", // .ഭാരതം (Malayalam)

	// South Korea: "KISA (Korea Internet & Security Agency)" vs "Korea Internet & Security Agency (KISA)"
	"xn--3e0b707e": "kr", // .한국

	// Bulgaria: "Imena.BG AD" vs "Register.BG"
	"xn--90ae": "bg", // .бг

	// Sri Lanka: "LK Domain Registry" vs "Council for Information Technology LK Domain Registrar"
	"xn--fzc2c9e2c":   "lk", // .ලංකා (Sinhala)
	"xn--xkc2al3hye2a": "lk", // .இலங்கை (Tamil)

	// Ukraine: "Ukrainian Network Information Centre (UANIC), Inc." vs "Hostmaster Ltd."
	"xn--j1amh": "ua", // .укр

	// Mongolia: "Datacom Co.,Ltd" vs "Datacom Co., Ltd." (punctuation)
	"xn--l1acc": "mn", // .мон

	// TRA disambiguations: Telecommunications Regulatory Authority manages both .bh and .om
	"xn--mgb9awbf":   "om", // .عمان (Oman)
	"xn--mgbcpq6gpa1a": "bh", // .البحرين (Bahrain)

	// Iran: "Institute for Research in Fundamental Sciences (IPM)" vs without "(IPM)"
	"xn--mgba3a4f16a": "ir", // .ایران

	// Pakistan: "National Telecommunication Corporation" vs "PKNIC"
	"xn--mgbai9azgqp6j": "pk", // .پاکستان

	// Georgia: "Information Technologies Development Center (ITDC)" vs "Caucasus Online LLC"
	"xn--node": "ge", // .გე

	// Egypt: "National Telecommunication Regulatory Authority - NTRA" vs "Egyptian Universities Network (EUN)"
	"xn--wgbh1c": "eg", // .مصر

	// Palestine: "Ministry of Telecom & Information Technology (MTIT)" vs different name
	"xn--ygbi2ammx": "ps", // .فلسطين
}

// RootDBEntry represents a single TLD row from the IANA root zone database index page.
type RootDBEntry struct {
	Domain       string // Unicode domain (e.g., "рф")
	Punycode     string // ASCII/punycode form (e.g., "xn--p1ai")
	Type         string // "country-code", "generic", "sponsored", "infrastructure", "test"
	Organisation string // Registry operator name
}

// ParseRootDBIndex parses the IANA root zone database index HTML page and
// returns a slice of entries, one per TLD row.
func ParseRootDBIndex(doc *goquery.Document) ([]RootDBEntry, error) {
	var entries []RootDBEntry
	doc.Find("table#tld-table tbody tr").Each(func(_ int, row *goquery.Selection) {
		cells := row.Find("td")
		if cells.Length() < 3 {
			return
		}

		// Column 1: domain link
		link := cells.Eq(0).Find("a")
		domainText := strings.TrimPrefix(strings.TrimSpace(link.Text()), ".")
		href, _ := link.Attr("href")

		// Extract punycode from href (e.g., "/domains/root/db/xn--p1ai.html" → "xn--p1ai")
		punycode := strings.TrimSuffix(path.Base(href), ".html")

		// Column 2: type
		tldType := strings.TrimSpace(cells.Eq(1).Text())

		// Column 3: organisation
		org := strings.TrimSpace(cells.Eq(2).Text())

		if domainText == "" || punycode == "" {
			return
		}

		entries = append(entries, RootDBEntry{
			Domain:       domainText,
			Punycode:     punycode,
			Type:         tldType,
			Organisation: org,
		})
	})

	if len(entries) == 0 {
		return nil, fmt.Errorf("no TLD entries found in IANA root DB index page")
	}
	return entries, nil
}

// FetchRootDBIndex fetches the IANA root zone database index page, parses it,
// and populates zone metadata: registryOperator, domainPunycode, domainAscii,
// domainIdn, and country tags.
func FetchRootDBIndex(zones map[string]*Zone, cache *ETagCache) error {
	res, err := FetchWithETag(ianaRootDBURL, cache)
	if err != nil {
		return err
	}
	if res == nil {
		return nil // 304 Not Modified
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return err
	}

	return applyRootDBIndex(zones, doc)
}

// applyRootDBIndex applies parsed IANA root DB data to zones. Separated from
// FetchRootDBIndex for testability.
func applyRootDBIndex(zones map[string]*Zone, doc *goquery.Document) error {
	entries, err := ParseRootDBIndex(doc)
	if err != nil {
		return err
	}

	// Pre-compute normalized (Unicode) domain for each entry, and build
	// lookups needed for IDN→ASCII mapping.
	normalized := make([]string, len(entries))         // parallel to entries
	orgToAscii := make(map[string][]string)            // org → []ASCII domain
	for i, e := range entries {
		normalized[i] = Normalize(e.Punycode)
		if e.Type == "country-code" && !strings.HasPrefix(e.Punycode, "xn--") {
			orgToAscii[e.Organisation] = append(orgToAscii[e.Organisation], e.Domain)
		}
	}

	// First pass: populate registryOperator, domainPunycode, country tags
	var operatorsSet, punycodeSet, tagsSet int
	for i, e := range entries {
		z := zones[normalized[i]]
		if z == nil {
			continue
		}

		// Registry operator
		if e.Organisation != "" && e.Organisation != "Not assigned" {
			if z.RegistryOperator == "" {
				operatorsSet++
			}
			z.RegistryOperator = e.Organisation
		}

		// Punycode for IDN TLDs
		if strings.HasPrefix(e.Punycode, "xn--") {
			z.DomainPunycode = e.Punycode
			punycodeSet++
		}

		// Apply type-based tags from the IANA root DB type column
		var typeTags []string
		switch e.Type {
		case "country-code":
			typeTags = []string{TagCountry, TagGeo}
		case "generic", "generic-restricted":
			typeTags = []string{TagGeneric}
		case "sponsored":
			typeTags = []string{TagSponsored}
		case "infrastructure":
			typeTags = []string{TagInfrastructure}
		case "test":
			typeTags = []string{TagTest}
		}
		if d := z.AddTags(typeTags...); d > 0 {
			tagsSet++
		}
	}

	// Second pass: derive domainAscii for IDN ccTLDs and build reverse mapping
	var asciiSet, idnSet, unmatched int
	asciiToIdn := make(map[string][]string) // ASCII domain → []Unicode IDN domain
	for i, e := range entries {
		if e.Type != "country-code" || !strings.HasPrefix(e.Punycode, "xn--") {
			continue
		}

		// Resolve the ASCII counterpart
		var asciiDomain string

		// Check manual override first
		if ascii, ok := idnToAsciiOverrides[e.Punycode]; ok {
			asciiDomain = ascii
		} else {
			// Try organisation matching
			candidates := orgToAscii[e.Organisation]
			switch len(candidates) {
			case 0:
				Trace("@{r}Error: IDN ccTLD %s (%s) has no ASCII counterpart — add to idnToAsciiOverrides\n", e.Domain, e.Punycode)
				unmatched++
				continue
			case 1:
				asciiDomain = candidates[0]
			default:
				// Multiple ASCII ccTLDs share the same organisation — needs override
				Trace("@{r}Error: IDN ccTLD %s (%s) has %d ASCII candidates: %s — add to idnToAsciiOverrides\n",
					e.Domain, e.Punycode, len(candidates), strings.Join(candidates, ", "))
				unmatched++
				continue
			}
		}

		// Apply domainAscii to IDN ccTLD
		if z := zones[normalized[i]]; z != nil {
			z.DomainAscii = asciiDomain
			asciiSet++
			asciiToIdn[asciiDomain] = append(asciiToIdn[asciiDomain], z.Domain)
		}
	}

	// Apply domainIdn to ASCII ccTLDs (reverse mapping)
	for asciiDomain, idnDomains := range asciiToIdn {
		z := zones[asciiDomain]
		if z == nil {
			continue
		}
		sort.Strings(idnDomains)
		z.DomainIdn = idnDomains
		idnSet++
	}

	Trace("@{.}IANA root DB: set %d registry operators, %d punycode, %d domainAscii, %d domainIdn, %d type tags\n",
		operatorsSet, punycodeSet, asciiSet, idnSet, tagsSet)

	if unmatched > 0 {
		return fmt.Errorf("%d IDN ccTLD(s) have no ASCII counterpart mapping — add to idnToAsciiOverrides in iana_root_db.go", unmatched)
	}

	return nil
}
