package build

import (
	"context"
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
	"xn--fzc2c9e2c":    "lk", // .ලංකා (Sinhala)
	"xn--xkc2al3hye2a": "lk", // .இலங்கை (Tamil)

	// Ukraine: "Ukrainian Network Information Centre (UANIC), Inc." vs "Hostmaster Ltd."
	"xn--j1amh": "ua", // .укр

	// Mongolia: "Datacom Co.,Ltd" vs "Datacom Co., Ltd." (punctuation)
	"xn--l1acc": "mn", // .мон

	// TRA disambiguations: Telecommunications Regulatory Authority manages both .bh and .om
	"xn--mgb9awbf":     "om", // .عمان (Oman)
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

// normalizeRootDBType canonicalizes the root DB type column to the slug form
// the tagging switch expects. IANA now renders display text ("Country Code")
// where it previously emitted slugs ("country-code"). Downstream matching is
// exact, so an unnormalized value silently drops country tags.
//
// "Special Use" is deliberately not mapped to "infrastructure": it covers
// nine TLDs (.alt, .example, .internal, .invalid, .local, .localhost, .onion,
// .test and .arpa), and only .arpa is infrastructure. Leaving it unmapped
// means those zones get no type tag from this source, which is correct —
// their tags come from FetchSpecialUseDomainsFromIANA.
func normalizeRootDBType(s string) string {
	s = strings.ToLower(strings.Join(strings.Fields(s), " "))
	return strings.ReplaceAll(s, " ", "-")
}

// isRootDBOrganisation reports whether the TLD Manager column holds an
// organisation name rather than a prose purpose. IANA reuses that column for
// both: special-use TLDs carry descriptions ("Used for clearly invalid domain
// names.") that must not be written into registryOperator.
func isRootDBOrganisation(tldType string) bool {
	return tldType != "special-use"
}

// rootDBColumns locates the "Type" and "TLD Manager" columns by header text.
// IANA added a "Status" column ahead of them in 2026, shifting the fixed
// offsets this code previously assumed. Falls back to the historic 1 and 2
// when no usable header is present.
func rootDBColumns(table *goquery.Selection) (typeCol, orgCol int) {
	typeCol, orgCol = 1, 2
	headers := table.Find("thead tr").First().Find("th")
	if headers.Length() == 0 {
		return typeCol, orgCol
	}
	foundType, foundOrg := false, false
	headers.Each(func(i int, th *goquery.Selection) {
		switch h := strings.ToLower(strings.TrimSpace(th.Text())); {
		case !foundType && h == "type":
			typeCol, foundType = i, true
		// The manager column has been titled "TLD Manager", "Sponsoring
		// Organisation", and "TLD Manager / Purpose" over time.
		case !foundOrg && (strings.Contains(h, "manager") || strings.Contains(h, "organisation")):
			orgCol, foundOrg = i, true
		}
	})
	return typeCol, orgCol
}

// ParseRootDBIndex parses the IANA root zone database index HTML page and
// returns a slice of entries, one per TLD row.
func ParseRootDBIndex(doc *goquery.Document) ([]RootDBEntry, error) {
	var entries []RootDBEntry
	table := doc.Find("table#ddtable, table#tld-table").First()
	typeCol, orgCol := rootDBColumns(table)
	// IANA renamed the table from #tld-table to #ddtable; match either so a
	// cached or reverted copy still parses.
	table.Find("tbody tr").Each(func(_ int, row *goquery.Selection) {
		cells := row.Find("td")
		if cells.Length() < 3 {
			return
		}

		// Column 1: domain link
		link := cells.Eq(0).Find("a")
		domainText := strings.TrimPrefix(strings.TrimSpace(link.Text()), ".")
		href, _ := link.Attr("href")

		// Extract punycode from href. The suffix is optional: the current
		// layout uses "/domains/root/db/xn--p1ai", the older one appended
		// ".html".
		punycode := strings.TrimSuffix(path.Base(href), ".html")

		// The current layout inserts a "Status" column before "Type", so
		// locate the columns by header rather than fixed offset.
		tldType := normalizeRootDBType(cells.Eq(typeCol).Text())
		var org string
		if isRootDBOrganisation(tldType) {
			org = strings.TrimSpace(cells.Eq(orgCol).Text())
		}

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
func FetchRootDBIndex(ctx context.Context, zones map[string]*Zone, cache *ETagCache) error {
	res, err := FetchWithETag(ctx, ianaRootDBURL, cache)
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
	normalized := make([]string, len(entries)) // parallel to entries
	orgToAscii := make(map[string][]string)    // org → []ASCII domain
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
