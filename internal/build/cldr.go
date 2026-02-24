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

// ccTLDToTerritory maps ccTLD domains to ISO 3166 territory codes where
// the mapping is not a simple uppercase transformation.
var ccTLDToTerritory = map[string]string{
	"uk": "GB", // .uk → United Kingdom (ISO 3166: GB)
}

// TerritoryFromCCTLD derives the ISO 3166 territory code from an ASCII ccTLD domain.
// Most ccTLDs are simply the uppercased domain (e.g., "ru" → "RU"), with
// exceptions like "uk" → "GB".
func TerritoryFromCCTLD(domain string) string {
	if t, ok := ccTLDToTerritory[domain]; ok {
		return t
	}
	return strings.ToUpper(domain)
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

		territory := TerritoryFromCCTLD(asciiDomain)

		// Country name
		if name := CountryName(territory); name != "" {
			if z.CountryName == "" {
				namesSet++
			}
			z.CountryName = name
		}

		// Official languages from CLDR
		officialLangs := cldr.OfficialLanguages(territory)
		if len(officialLangs) == 0 {
			continue
		}

		bcp47Tags := ComposeBCP47Tags(officialLangs, territory)
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
