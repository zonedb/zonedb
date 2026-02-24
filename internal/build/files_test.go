package build

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"golang.org/x/text/language"
)

// readAllZones is a test helper that reads zones.txt and metadata from the
// repository root.
func readAllZones(t *testing.T) map[string]*Zone {
	t.Helper()
	oldBaseDir := BaseDir
	BaseDir = "../.."
	defer func() { BaseDir = oldBaseDir }()

	zones, errs := ReadZones()
	if len(errs) > 0 {
		t.Fatalf("ReadZones: %d error(s)", len(errs))
	}
	return zones
}

func TestZonesFileValid(t *testing.T) {
	// This test validates that zones.txt contains only valid domain names.
	// It will fail if any invalid zones are present (e.g., " (deprecated)").

	// Set BaseDir to repository root (two directories up from internal/build)
	oldBaseDir := BaseDir
	BaseDir = "../.."
	defer func() { BaseDir = oldBaseDir }()

	zones, errs := ReadZonesFile()
	if len(errs) > 0 {
		for _, err := range errs {
			t.Error(err)
		}
	}

	// Additional validation: check for invalid characters that Normalize() doesn't catch
	for domain := range zones {
		if strings.ContainsAny(domain, " ()") {
			t.Errorf("invalid domain name contains spaces or parentheses: %q", domain)
		}
		if strings.HasPrefix(domain, " ") {
			t.Errorf("invalid domain name has leading space: %q", domain)
		}
	}
}

// TestWriteMetadataFileNoHTMLEscape verifies that writeMetadataFile does not
// HTML-escape special characters like & in JSON output. Go's json.Encoder
// escapes & as \u0026 by default unless SetEscapeHTML(false) is set.
func TestWriteMetadataFileNoHTMLEscape(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "test.json")

	z := &Zone{
		Domain:           "ag",
		RegistryOperator: "UHSA School of Medicine",
		CountryName:      "Antigua & Barbuda",
	}

	if err := writeMetadataFile(z, path); err != nil {
		t.Fatalf("writeMetadataFile: %v", err)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("reading output: %v", err)
	}
	content := string(data)

	if strings.Contains(content, `\u0026`) {
		t.Errorf("JSON contains escaped \\u0026, want literal &:\n%s", content)
	}
	if !strings.Contains(content, "Antigua & Barbuda") {
		t.Errorf("JSON missing literal 'Antigua & Barbuda':\n%s", content)
	}
}

// TestMetadataLanguageTagsValid reads all metadata files and validates that
// every language tag is a recognized BCP 47 tag with known subtags (not just
// syntactically well-formed). language.Parse returns an error for tags like
// "xx-YY" where the subtag is well-formed but not in the IANA subtag registry.
func TestMetadataLanguageTagsValid(t *testing.T) {
	zones := readAllZones(t)

	var checked int
	for _, z := range zones {
		for _, lang := range z.Languages {
			tag, err := language.Parse(lang)
			if err != nil {
				t.Errorf("zone %s: invalid language tag %q: %v", z.Domain, lang, err)
				continue
			}
			if tag == language.Und {
				t.Errorf("zone %s: language tag %q parsed to undefined", z.Domain, lang)
			}
			checked++
		}
	}
	t.Logf("validated %d language tags across %d zones", checked, len(zones))
}

// asciiCCTLD returns the ASCII ccTLD domain for a zone, used for exception
// lookups. For IDN ccTLDs this is domainAscii; for ASCII ccTLDs it's the
// domain itself.
func asciiCCTLD(z *Zone) string {
	if z.DomainAscii != "" {
		return z.DomainAscii
	}
	return z.Domain
}

// TestMetadataCCTLDsHaveCountryName validates that every country-code TLD has
// a countryName populated.
func TestMetadataCCTLDsHaveCountryName(t *testing.T) {
	zones := readAllZones(t)

	var checked int
	for _, z := range zones {
		if !z.IsTLD() || !NewSet(z.Tags...)[TagCountry] {
			continue
		}
		if z.CountryName == "" {
			t.Errorf("zone %s: country-code TLD missing countryName", z)
		}
		checked++
	}
	t.Logf("validated countryName for %d country-code TLDs", checked)
}

// TestMetadataCCTLDsHaveLanguages validates that country-code TLDs with
// languages have valid, non-empty language tags, and that a reasonable
// number of ccTLDs have languages assigned.
// Not all ccTLDs receive languages: zones with idn-disallowed policies,
// multi-language territories without a single dominant language (≥50%),
// and zones in CCTLDsWithoutOfficialLanguages are intentionally skipped.
func TestMetadataCCTLDsHaveLanguages(t *testing.T) {
	zones := readAllZones(t)

	var withLangs, withoutLangs, skipped int
	for _, z := range zones {
		if !z.IsTLD() || !NewSet(z.Tags...)[TagCountry] {
			continue
		}
		if CCTLDsWithoutOfficialLanguages[asciiCCTLD(z)] {
			skipped++
			continue
		}
		if len(z.Languages) == 0 {
			withoutLangs++
			continue
		}
		withLangs++
	}
	t.Logf("ccTLDs with languages: %d, without: %d, skipped: %d", withLangs, withoutLangs, skipped)
	// Sanity check: at least 40 ccTLDs should have languages (single-dominant-language
	// territories like .ru, .br, .de, .fr, .jp, etc. plus IDN table policy zones).
	if withLangs < 40 {
		t.Errorf("only %d ccTLDs have languages, expected at least 40", withLangs)
	}
}

// TestMetadataIDNccTLDsComplete validates that every IDN country-code TLD has
// domainPunycode and domainAscii, and that the ASCII counterpart's domainIdn
// contains the reverse mapping.
func TestMetadataIDNccTLDsComplete(t *testing.T) {
	zones := readAllZones(t)

	var checked int
	for _, z := range zones {
		if !z.IsTLD() || !z.IsIDN() || !NewSet(z.Tags...)[TagCountry] {
			continue
		}
		if z.DomainPunycode == "" {
			t.Errorf("zone %s: IDN ccTLD missing domainPunycode", z)
		}
		if z.DomainAscii == "" {
			t.Errorf("zone %s: IDN ccTLD missing domainAscii", z)
			continue
		}
		ascii := zones[z.DomainAscii]
		if ascii == nil {
			t.Errorf("zone %s: domainAscii %q not found in zones", z, z.DomainAscii)
			continue
		}
		if !NewSet(ascii.DomainIdn...)[z.Domain] {
			t.Errorf("zone %s: ASCII counterpart %q domainIdn=%v missing %q",
				z, z.DomainAscii, ascii.DomainIdn, z.Domain)
		}
		checked++
	}
	t.Logf("validated %d IDN ccTLD mappings", checked)
}
