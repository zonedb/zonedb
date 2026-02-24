package build

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"golang.org/x/text/language"
)

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
	oldBaseDir := BaseDir
	BaseDir = "../.."
	defer func() { BaseDir = oldBaseDir }()

	zones, errs := ReadZones()
	if len(errs) > 0 {
		t.Fatalf("ReadZones: %d error(s)", len(errs))
	}

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
