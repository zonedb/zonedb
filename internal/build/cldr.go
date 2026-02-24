package build

import (
	"encoding/json"
	"fmt"
	"io"
	"sort"
	"strings"

	"golang.org/x/text/language"
	"golang.org/x/text/language/display"
)

var cldrTerritoryInfoURL = "https://raw.githubusercontent.com/unicode-org/cldr-json/main/cldr-json/cldr-core/supplemental/territoryInfo.json"

// CLDRTerritoryInfo holds the parsed CLDR territoryInfo.json data.
type CLDRTerritoryInfo struct {
	Supplemental struct {
		TerritoryInfo map[string]CLDRTerritory `json:"territoryInfo"`
	} `json:"supplemental"`
}

// CLDRTerritory holds territory-level data from CLDR.
type CLDRTerritory struct {
	LanguagePopulation map[string]CLDRLanguagePopulation `json:"languagePopulation"`
}

// CLDRLanguagePopulation holds per-language data within a territory.
type CLDRLanguagePopulation struct {
	PopulationPercent string `json:"_populationPercent"`
	OfficialStatus    string `json:"_officialStatus"`
}

// ParseCLDRTerritoryInfo parses a CLDR territoryInfo.json file.
func ParseCLDRTerritoryInfo(r io.Reader) (*CLDRTerritoryInfo, error) {
	var info CLDRTerritoryInfo
	if err := json.NewDecoder(r).Decode(&info); err != nil {
		return nil, fmt.Errorf("parsing CLDR territoryInfo.json: %w", err)
	}
	return &info, nil
}

// OfficialLanguages returns the official and de facto official language codes
// for a given ISO 3166 territory code (e.g., "RU", "BR").
func (info *CLDRTerritoryInfo) OfficialLanguages(territory string) []string {
	t, ok := info.Supplemental.TerritoryInfo[territory]
	if !ok {
		return nil
	}
	var langs []string
	for lang, pop := range t.LanguagePopulation {
		switch pop.OfficialStatus {
		case "official", "de_facto_official":
			langs = append(langs, lang)
		}
	}
	sort.Strings(langs)
	return langs
}

// ccTLDLangTerritory maps ccTLD domains to ISO 3166 territory codes for
// CLDR language lookup. These overrides point to territories that have
// language data in CLDR, which may differ from the ccTLD's own territory.
// Country names are derived separately (see CountryNameForCCTLD).
var ccTLDLangTerritory = map[string]string{
	"uk": "GB", // .uk → United Kingdom (ISO 3166: GB)

	// Defunct/historical ccTLDs → successor states (for language data)
	"an": "CW", // .an → Netherlands Antilles → Curaçao
	"su": "RU", // .su → Soviet Union → Russia
	"tp": "TL", // .tp → East Timor → Timor-Leste
	"yu": "RS", // .yu → Yugoslavia → Serbia

	// Uninhabited/minimal territories → administering countries (for language data)
	"ac": "SH", // .ac → Ascension Island → Saint Helena
	"bv": "NO", // .bv → Bouvet Island → Norway
	"hm": "AU", // .hm → Heard & McDonald Islands → Australia
	"tf": "FR", // .tf → French Southern Territories → France
}

// ccTLDCountryName provides explicit country names for defunct ccTLDs whose
// ISO 3166 codes are not recognized by golang.org/x/text.
var ccTLDCountryName = map[string]string{
	"an": "Netherlands Antilles",
	"su": "USSR",
	"yu": "Yugoslavia",
}

// CCTLDsWithoutOfficialLanguages lists ASCII ccTLD domains whose territories
// have no official languages in CLDR. These zones get countryName but not
// languages. IDN ccTLDs inherit this via their domainAscii mapping.
var CCTLDsWithoutOfficialLanguages = NewSet(
	"aq", // Antarctica — no permanent population, no official language
	"eu", // European Union — supranational; member states have their own ccTLDs
)

// TerritoryFromCCTLD derives the ISO 3166 territory code for CLDR language
// lookup from an ASCII ccTLD domain. Most ccTLDs are simply the uppercased
// domain (e.g., "ru" → "RU"), with overrides for territories where CLDR
// has no language data under the direct code.
func TerritoryFromCCTLD(domain string) string {
	if t, ok := ccTLDLangTerritory[domain]; ok {
		return t
	}
	return strings.ToUpper(domain)
}

// CountryNameForCCTLD returns the country name for a ccTLD domain.
// It prefers the ccTLD's own territorial identity (e.g., .tf → "French
// Southern Territories") over the language-lookup override (which would
// give "France"). For defunct ccTLDs whose codes are not in golang.org/x/text,
// explicit names are provided via ccTLDCountryName.
func CountryNameForCCTLD(domain string) string {
	// Explicit override for defunct codes not in x/text
	if name, ok := ccTLDCountryName[domain]; ok {
		return name
	}
	// Try the ccTLD's own territory code first (preserves territorial identity)
	if name := CountryName(strings.ToUpper(domain)); name != "" {
		return name
	}
	// Fall back to the language-lookup territory
	return CountryName(TerritoryFromCCTLD(domain))
}

// ComposeBCP47Tags composes BCP 47 language-region tags from a list of CLDR
// language codes and a territory code.
// CLDR language codes may include script (e.g., "az_Cyrl") which is converted
// to BCP 47 form (e.g., "az-Cyrl-RU").
func ComposeBCP47Tags(langs []string, territory string) []string {
	reg, err := language.ParseRegion(territory)
	if err != nil {
		return nil
	}
	var tags []string
	for _, lang := range langs {
		// CLDR uses underscore, BCP 47 uses hyphen
		lang = strings.ReplaceAll(lang, "_", "-")
		base, err := language.Parse(lang)
		if err != nil {
			Trace("@{y}Warning: cannot parse CLDR language %q: %v\n", lang, err)
			continue
		}
		tag, _ := language.Compose(base, reg)
		tags = append(tags, tag.String())
	}
	sort.Strings(tags)
	return tags
}

// CountryName returns the English name for an ISO 3166 territory code.
func CountryName(territory string) string {
	reg, err := language.ParseRegion(territory)
	if err != nil {
		return ""
	}
	return display.English.Regions().Name(reg)
}

// FetchAndApplyCLDRMetadata fetches CLDR territoryInfo.json and applies
// language and country metadata to country-code TLDs.
func FetchAndApplyCLDRMetadata(zones map[string]*Zone, cache *ETagCache) error {
	res, err := FetchWithETag(cldrTerritoryInfoURL, cache)
	if err != nil {
		return err
	}
	if res == nil {
		Trace("@{.}CLDR: 304 Not Modified (cached)\n")
		return nil
	}
	defer res.Body.Close()

	cldr, err := ParseCLDRTerritoryInfo(res.Body)
	if err != nil {
		return err
	}

	ApplyCLDRMetadata(zones, cldr)
	return nil
}

// ApplyCLDRMetadata applies CLDR-derived language and country metadata to
// country-code TLDs. It uses the CLDR territory info, zone domainAscii
// mappings (set by FetchRootDBIndex), and golang.org/x/text for BCP 47
// composition and country names.
func ApplyCLDRMetadata(zones map[string]*Zone, cldr *CLDRTerritoryInfo) {
	var langsSet, namesSet int
	for _, z := range zones {
		if !z.IsTLD() {
			continue
		}

		// Determine the ASCII ccTLD domain for territory lookup
		var asciiDomain string
		tags := NewSet(z.Tags...)
		if !tags[TagCountry] {
			continue
		}

		if z.DomainAscii != "" {
			// IDN ccTLD — use the linked ASCII domain
			asciiDomain = z.DomainAscii
		} else if !z.IsIDN() {
			// ASCII ccTLD
			asciiDomain = z.Domain
		} else {
			continue
		}

		// Country name: use the ccTLD's own territorial identity
		if name := CountryNameForCCTLD(asciiDomain); name != "" {
			if z.CountryName == "" {
				namesSet++
			}
			z.CountryName = name
		}

		// Language territory may differ (e.g., .tf → FR for French language data)
		langTerritory := TerritoryFromCCTLD(asciiDomain)

		// Official languages from CLDR
		officialLangs := cldr.OfficialLanguages(langTerritory)
		if len(officialLangs) == 0 {
			continue
		}

		bcp47Tags := ComposeBCP47Tags(officialLangs, langTerritory)
		if len(bcp47Tags) == 0 {
			continue
		}

		// Merge CLDR-derived languages into existing languages (additive)
		existing := NewSet(z.Languages...)
		before := len(existing)
		existing.Add(bcp47Tags...)
		if len(existing) > before {
			langsSet++
		}
		z.Languages = existing.Values()
	}

	Trace("@{.}CLDR: set %d country names, added languages to %d zone(s)\n", namesSet, langsSet)
}
