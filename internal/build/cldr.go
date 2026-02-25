package build

import (
	"encoding/json"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"
	"unicode"

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

// DominantOfficialLanguages returns official/de facto official language codes
// that are spoken by at least the given percentage of the population.
// This filters out regional or minority official languages that, while
// recognized, do not represent the dominant language(s) of the territory.
func (info *CLDRTerritoryInfo) DominantOfficialLanguages(territory string, minPopulationPercent float64) []string {
	t, ok := info.Supplemental.TerritoryInfo[territory]
	if !ok {
		return nil
	}
	var langs []string
	for lang, pop := range t.LanguagePopulation {
		switch pop.OfficialStatus {
		case "official", "de_facto_official":
		default:
			continue
		}
		pct, err := strconv.ParseFloat(pop.PopulationPercent, 64)
		if err != nil {
			continue
		}
		if pct >= minPopulationPercent {
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

// CCTLDsWithoutOfficialLanguages lists ASCII ccTLD domains that should not
// receive language data from CLDR. These zones get countryName but not
// languages. IDN ccTLDs inherit this via their domainAscii mapping.
var CCTLDsWithoutOfficialLanguages = NewSet(
	"aq", // Antarctica — no permanent population, no official language
	"eu", // European Union — supranational; member states have their own ccTLDs

	// Defunct/historical ccTLDs — not actively associated with a living language community
	"an", // Netherlands Antilles (dissolved 2010)
	"su", // Soviet Union (dissolved 1991)
	"tp", // East Timor (replaced by .tl)
	"yu", // Yugoslavia (dissolved 2003)
)

// hanScriptVariants maps IDN ccTLD punycode domains to Han script variants
// (Hans = simplified, Hant = traditional). Unicode cannot distinguish
// simplified from traditional Chinese characters — both are "Han" script —
// so variant pairs for the same territory must be specified explicitly.
// Domains not in this table that use Han script derive their variant from
// CLDR territory data (e.g., HK → zh_Hant → "Hant").
var hanScriptVariants = map[string]string{
	"xn--fiqs8s":  "Hans", // .中国 — simplified Chinese
	"xn--fiqz9s":  "Hant", // .中國 — traditional Chinese
	"xn--kprw13d": "Hans", // .台湾 — simplified Chinese
	"xn--kpry57d": "Hant", // .台灣 — traditional Chinese
}

// scriptAmbiguousOverrides provides explicit language assignments for IDN
// ccTLDs where multiple domains share the same Unicode script for the same
// territory, making automated script-matching ambiguous. When the pre-scan
// detects such ambiguity, only domains in this table receive CLDR languages.
var scriptAmbiguousOverrides = map[string][]string{
	"xn--h2brj9c": {"hi-IN"}, // .भारत — Hindi (multiple Devanagari ccTLDs for IN)
}

// iso15924ToUnicode maps BCP 47 / ISO 15924 script codes to Go Unicode range
// tables. Used to match IDN domain scripts against CLDR language scripts.
var iso15924ToUnicode = map[string]*unicode.RangeTable{
	"Arab": unicode.Arabic,
	"Armn": unicode.Armenian,
	"Beng": unicode.Bengali,
	"Cyrl": unicode.Cyrillic,
	"Deva": unicode.Devanagari,
	"Geor": unicode.Georgian,
	"Grek": unicode.Greek,
	"Gujr": unicode.Gujarati,
	"Guru": unicode.Gurmukhi,
	"Hang": unicode.Hangul,
	"Hans": unicode.Han,
	"Hant": unicode.Han,
	"Khmr": unicode.Khmer,
	"Knda": unicode.Kannada,
	"Laoo": unicode.Lao,
	"Mlym": unicode.Malayalam,
	"Mymr": unicode.Myanmar,
	"Orya": unicode.Oriya,
	"Sinh": unicode.Sinhala,
	"Taml": unicode.Tamil,
	"Telu": unicode.Telugu,
	"Thai": unicode.Thai,
	"Tibt": unicode.Tibetan,
}

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

// domainUsesHan reports whether any rune in the domain belongs to the Han
// (CJK Unified Ideographs) Unicode block.
func domainUsesHan(domain string) bool {
	for _, r := range domain {
		if unicode.Is(unicode.Han, r) {
			return true
		}
	}
	return false
}

// domainUsesScript reports whether any rune in the domain belongs to the
// Unicode range table for the given BCP 47 / ISO 15924 script code.
func domainUsesScript(domain, scriptCode string) bool {
	rt, ok := iso15924ToUnicode[scriptCode]
	if !ok {
		return false
	}
	for _, r := range domain {
		if unicode.Is(rt, r) {
			return true
		}
	}
	return false
}

// dominantNonHanScript returns the BCP 47 script code of the first non-ASCII,
// non-Han rune in the domain, or "" if no recognized script is found.
// Han domains are handled separately via hanScriptVariants.
func dominantNonHanScript(domain string) string {
	for _, r := range domain {
		if r < 0x80 || unicode.Is(unicode.Han, r) {
			continue
		}
		for code, table := range iso15924ToUnicode {
			if code == "Hans" || code == "Hant" {
				continue
			}
			if unicode.Is(table, r) {
				return code
			}
		}
	}
	return ""
}

// idnScriptCountsByTerritory counts IDN ccTLDs per territory+script pair.
// Used to detect ambiguous script-matching (e.g., multiple Devanagari domains
// for India) where automated language assignment would be incorrect.
func idnScriptCountsByTerritory(zones map[string]*Zone) map[string]int {
	counts := map[string]int{}
	for _, z := range zones {
		if !z.IsTLD() || !z.IsIDN() {
			continue
		}
		tags := NewSet(z.Tags...)
		asciiDomain := z.ASCIICcTLD()
		if !tags[TagCountry] || asciiDomain == "" {
			continue
		}
		territory := TerritoryFromCCTLD(asciiDomain)
		script := dominantNonHanScript(z.Domain)
		if script != "" {
			counts[territory+":"+script]++
		}
	}
	return counts
}

// hanVariantForIDN determines the Han script variant (Hans or Hant) for an
// IDN ccTLD. It checks the hanScriptVariants table first (for variant pairs
// like .中国/.中國), then derives the variant from CLDR territory data
// (e.g., HK has zh_Hant → "Hant"), defaulting to "Hans".
func hanVariantForIDN(punycode string, cldr *CLDRTerritoryInfo, territory string) string {
	if v, ok := hanScriptVariants[punycode]; ok {
		return v
	}
	// Derive from CLDR: look for zh with explicit script in territory
	t, ok := cldr.Supplemental.TerritoryInfo[territory]
	if !ok {
		return "Hans"
	}
	for lang := range t.LanguagePopulation {
		if !strings.HasPrefix(lang, "zh") {
			continue
		}
		tag, err := language.Parse(strings.ReplaceAll(lang, "_", "-"))
		if err != nil {
			continue
		}
		if s, conf := tag.Script(); conf >= language.High {
			return s.String()
		}
	}
	return "Hans"
}

// languageScriptMatchesDomain checks whether a CLDR language code, when
// composed with a territory, produces a BCP 47 tag whose script matches
// the Unicode script used in the domain name.
func languageScriptMatchesDomain(lang, territory, domain string) bool {
	lang = strings.ReplaceAll(lang, "_", "-")
	base, err := language.Parse(lang)
	if err != nil {
		return false
	}
	reg, err := language.ParseRegion(territory)
	if err != nil {
		return false
	}
	tag, _ := language.Compose(base, reg)
	script, _ := tag.Script()
	return domainUsesScript(domain, script.String())
}

// composeSingleBCP47 composes a single BCP 47 tag from a CLDR language code
// and a parsed region.
func composeSingleBCP47(lang string, reg language.Region) string {
	lang = strings.ReplaceAll(lang, "_", "-")
	base, err := language.Parse(lang)
	if err != nil {
		return ""
	}
	tag, _ := language.Compose(base, reg)
	return tag.String()
}

// scriptMatchedLanguages returns BCP 47 language tags from CLDR for the
// given territory whose scripts match the Unicode script used in the domain.
// It tries official/de_facto_official languages first, falling back to
// include official_regional if no script matches are found.
func scriptMatchedLanguages(domain string, cldr *CLDRTerritoryInfo, territory string) []string {
	t, ok := cldr.Supplemental.TerritoryInfo[territory]
	if !ok {
		return nil
	}
	reg, err := language.ParseRegion(territory)
	if err != nil {
		return nil
	}

	// First pass: official and de_facto_official
	var matched []string
	for lang, pop := range t.LanguagePopulation {
		switch pop.OfficialStatus {
		case "official", "de_facto_official":
		default:
			continue
		}
		if languageScriptMatchesDomain(lang, territory, domain) {
			if tag := composeSingleBCP47(lang, reg); tag != "" {
				matched = append(matched, tag)
			}
		}
	}
	if len(matched) > 0 {
		sort.Strings(matched)
		return matched
	}

	// Second pass: include official_regional
	for lang, pop := range t.LanguagePopulation {
		switch pop.OfficialStatus {
		case "official", "de_facto_official", "official_regional":
		default:
			continue
		}
		if languageScriptMatchesDomain(lang, territory, domain) {
			if tag := composeSingleBCP47(lang, reg); tag != "" {
				matched = append(matched, tag)
			}
		}
	}
	sort.Strings(matched)
	return matched
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
//
// Language handling differs by zone type:
//   - IDN ccTLDs: languages are determined by matching the domain's Unicode
//     script against CLDR language data for the territory. Han script domains
//     use hanScriptVariants for simplified/traditional disambiguation. IDN
//     table policy languages are merged with CLDR-derived languages.
//   - ASCII ccTLDs with IDN table policies: languages are rebuilt from those
//     policies only, removing any stale CLDR-derived languages from prior builds.
//   - ASCII ccTLDs without IDN table policies: CLDR adds a language only when
//     the territory has exactly one dominant official language (≥50% population)
//     and the zone does not have idn-disallowed. Multi-language territories
//     (e.g., HK) are skipped so language search results come from IDN variants.
func ApplyCLDRMetadata(zones map[string]*Zone, cldr *CLDRTerritoryInfo) {
	// Pre-scan: count IDN ccTLDs per territory+script to detect ambiguity.
	// When multiple IDN ccTLDs share the same script for the same territory
	// (e.g., three Devanagari domains for India), automated script-matching
	// can't determine which domain represents which language.
	idnScriptCounts := idnScriptCountsByTerritory(zones)

	var langsSet, langsCleaned, namesSet int
	for _, z := range zones {
		if !z.IsTLD() {
			continue
		}

		tags := NewSet(z.Tags...)
		if !tags[TagCountry] {
			continue
		}

		// Determine the ASCII ccTLD domain for territory lookup
		asciiDomain := z.ASCIICcTLD()
		if asciiDomain == "" {
			continue
		}

		// Country name: use the ccTLD's own territorial identity
		if name := CountryNameForCCTLD(asciiDomain); name != "" {
			if z.CountryName == "" {
				namesSet++
			}
			z.CountryName = name
		}

		// --- Language handling for IDN ccTLDs ---
		if z.IsIDN() {
			// Collect IDN table policy languages
			var idnTableLangs []string
			for _, p := range z.Policies {
				if p.Type == TypeIDNTable && p.Key != "" {
					idnTableLangs = append(idnTableLangs, p.Key)
				}
			}

			// Skip excluded zones (keep IDN table languages only)
			if CCTLDsWithoutOfficialLanguages[asciiDomain] {
				if len(idnTableLangs) > 0 {
					sort.Strings(idnTableLangs)
					z.Languages = idnTableLangs
				} else if len(z.Languages) > 0 {
					z.Languages = nil
					langsCleaned++
				}
				continue
			}

			langTerritory := TerritoryFromCCTLD(asciiDomain)

			// Script-based language matching
			var cldrLangs []string
			if domainUsesHan(z.Domain) {
				// Han script: compose zh-{variant}-{region}
				variant := hanVariantForIDN(z.DomainPunycode, cldr, langTerritory)
				reg, err := language.ParseRegion(langTerritory)
				if err == nil {
					script, err2 := language.ParseScript(variant)
					if err2 == nil {
						zh := language.MustParse("zh")
						tag, _ := language.Compose(zh, script, reg)
						cldrLangs = []string{tag.String()}
					}
				}
			} else {
				// Non-Han: check for script ambiguity before matching
				script := dominantNonHanScript(z.Domain)
				key := langTerritory + ":" + script
				if script != "" && idnScriptCounts[key] > 1 {
					// Multiple domains share this script for this territory.
					// Use explicit override if available; otherwise skip.
					if override, ok := scriptAmbiguousOverrides[z.DomainPunycode]; ok {
						cldrLangs = override
					}
				} else {
					// Unambiguous: match CLDR languages by detected domain script
					cldrLangs = scriptMatchedLanguages(z.Domain, cldr, langTerritory)
				}
			}

			// Merge IDN table languages + CLDR-derived languages
			merged := make(map[string]bool)
			for _, l := range idnTableLangs {
				merged[l] = true
			}
			for _, l := range cldrLangs {
				merged[l] = true
			}

			if len(merged) > 0 {
				var langs []string
				for l := range merged {
					langs = append(langs, l)
				}
				sort.Strings(langs)
				z.Languages = langs
				langsSet++
			} else if len(z.Languages) > 0 {
				z.Languages = nil
				langsCleaned++
			}
			continue
		}

		// --- Language handling for ASCII ccTLDs ---

		// Rebuild languages from IDN table policies, removing any stale
		// CLDR-derived languages from previous builds.
		var idnTableLangs []string
		for _, p := range z.Policies {
			if p.Type == TypeIDNTable && p.Key != "" {
				idnTableLangs = append(idnTableLangs, p.Key)
			}
		}
		sort.Strings(idnTableLangs)

		if len(idnTableLangs) > 0 {
			// IDN table policies are the authoritative source of languages.
			if len(z.Languages) != len(idnTableLangs) {
				langsCleaned++
			}
			z.Languages = idnTableLangs
			continue
		}

		// No IDN table policies. Use CLDR as a fallback only when the
		// territory has a single dominant official language.
		// Skip zones with idn-disallowed (e.g., .us) or those in the
		// explicit exclusion set.
		if hasIDNDisallowedPolicy(z) || CCTLDsWithoutOfficialLanguages[asciiDomain] {
			if len(z.Languages) > 0 {
				z.Languages = nil
				langsCleaned++
			}
			continue
		}

		langTerritory := TerritoryFromCCTLD(asciiDomain)
		officialLangs := cldr.DominantOfficialLanguages(langTerritory, 50)

		// Only add when exactly one dominant language exists.
		// Multi-language territories (e.g., HK) are skipped so that
		// language search results come from IDN variants instead.
		if len(officialLangs) != 1 {
			if len(z.Languages) > 0 {
				z.Languages = nil
				langsCleaned++
			}
			continue
		}

		bcp47Tags := ComposeBCP47Tags(officialLangs, langTerritory)
		if len(bcp47Tags) == 0 {
			if len(z.Languages) > 0 {
				z.Languages = nil
				langsCleaned++
			}
			continue
		}

		z.Languages = bcp47Tags
		langsSet++
	}

	Trace("@{.}CLDR: set %d country names, added languages to %d zone(s), cleaned %d zone(s)\n", namesSet, langsSet, langsCleaned)
}

// hasIDNDisallowedPolicy reports whether the zone has an idn-disallowed policy.
func hasIDNDisallowedPolicy(z *Zone) bool {
	for _, p := range z.Policies {
		if p.Type == TypeIDNDisallowed {
			return true
		}
	}
	return false
}
