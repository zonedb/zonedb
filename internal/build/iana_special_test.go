package build

import (
	"encoding/csv"
	"os"
	"testing"
)

func TestStripDeprecatedSuffix(t *testing.T) {
	tests := []struct {
		input        string
		wantDomain   string
		wantRetired  bool
	}{
		// Deprecated markers should be stripped
		{"eap-noob.arpa. (DEPRECATED)", "eap-noob.arpa", true},
		{"eap-noob.arpa. (deprecated)", "eap-noob.arpa", true},
		{"example.com (DEPRECATED)", "example.com", true},
		{"example.com (Deprecated)", "example.com", true},
		{"test.arpa.  (DEPRECATED)", "test.arpa", true},

		// Normal domains should pass through unchanged
		{"example.com", "example.com", false},
		{"arpa.", "arpa.", false},
		{"home.arpa.", "home.arpa.", false},
		{"localhost.", "localhost.", false},

		// Edge cases
		{"", "", false},
		{"(deprecated)", "", true}, // Just marker, no domain - stripped to empty string
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			gotDomain, gotRetired := stripDeprecatedSuffix(tt.input)
			if gotDomain != tt.wantDomain {
				t.Errorf("stripDeprecatedSuffix(%q) domain = %q, want %q", tt.input, gotDomain, tt.wantDomain)
			}
			if gotRetired != tt.wantRetired {
				t.Errorf("stripDeprecatedSuffix(%q) retired = %v, want %v", tt.input, gotRetired, tt.wantRetired)
			}
		})
	}
}

func TestDeprecatedSuffixPattern(t *testing.T) {
	// Ensure the pattern matches various deprecated formats
	shouldMatch := []string{
		"domain. (DEPRECATED)",
		"domain. (deprecated)",
		"domain. (Deprecated)",
		"domain.(DEPRECATED)",
		"domain.  (DEPRECATED)", // extra space
	}

	shouldNotMatch := []string{
		"domain.com",
		"deprecated.com",
		"domain.deprecated",
		"(deprecated).com",
		"domain (other)",
		"domain (deprecatedish)",
		"domain. (OBSOLETE)", // only deprecated is supported
	}

	for _, s := range shouldMatch {
		if !deprecatedSuffixPattern.MatchString(s) {
			t.Errorf("expected pattern to match %q", s)
		}
	}

	for _, s := range shouldNotMatch {
		if deprecatedSuffixPattern.MatchString(s) {
			t.Errorf("expected pattern NOT to match %q", s)
		}
	}
}

func TestParseIANASpecialUseDomainsCSV(t *testing.T) {
	// This test uses the actual IANA CSV file that contains deprecated entries.
	// It verifies that:
	// 1. Deprecated markers are stripped from domain names
	// 2. No invalid zones (like " (deprecated)") are created
	// 3. The cleaned domain is tagged as retired

	f, err := os.Open("testdata/2026-01-16-special-use-domains.csv")
	if err != nil {
		t.Fatalf("failed to open test fixture: %v", err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	cols, err := r.Read()
	if err != nil {
		t.Fatalf("failed to read CSV header: %v", err)
	}

	var foundDeprecated bool
	records, err := r.ReadAll()
	if err != nil {
		t.Fatalf("failed to read CSV records: %v", err)
	}

	for _, row := range records {
		rec := make(map[string]string)
		for i := range cols {
			rec[cols[i]] = row[i]
		}

		rawName := rec["Name"]
		cleanName, isDeprecated := stripDeprecatedSuffix(rawName)

		if isDeprecated {
			foundDeprecated = true

			// Verify the deprecated entry is the one we expect
			if rawName != "eap-noob.arpa. (DEPRECATED)" {
				t.Errorf("unexpected deprecated entry: %q", rawName)
			}

			// Verify the domain is correctly cleaned
			if cleanName != "eap-noob.arpa" {
				t.Errorf("stripDeprecatedSuffix(%q) = %q, want %q", rawName, cleanName, "eap-noob.arpa")
			}

			// Verify the normalized domain is valid
			normalized := Normalize(cleanName)
			if normalized != "eap-noob.arpa" {
				t.Errorf("Normalize(%q) = %q, want %q", cleanName, normalized, "eap-noob.arpa")
			}
		}

		// Verify no domain name contains deprecated markers after stripping
		if deprecatedSuffixPattern.MatchString(cleanName) {
			t.Errorf("cleaned domain still contains deprecated marker: %q", cleanName)
		}
	}

	if !foundDeprecated {
		t.Error("test fixture should contain at least one deprecated entry")
	}
}

func TestNoInvalidZonesFromDeprecatedDomains(t *testing.T) {
	// This test simulates what FetchSpecialUseDomainsFromIANA does and verifies
	// that no invalid zones (like " (deprecated)") are created when processing
	// deprecated domain names.

	zones := make(map[string]*Zone)

	// Simulate processing "eap-noob.arpa. (DEPRECATED)" from the IANA CSV
	rawName := "eap-noob.arpa. (DEPRECATED)"
	cleanName, isDeprecated := stripDeprecatedSuffix(rawName)

	if !isDeprecated {
		t.Fatal("expected entry to be marked as deprecated")
	}

	domain := Normalize(cleanName)
	d := domain

	// Walk up parent domains like FetchSpecialUseDomainsFromIANA does
	for d != "" {
		z := zones[d]
		if z == nil {
			z = &Zone{Domain: d}
			zones[d] = z
		}
		d = z.ParentDomain()
	}

	// Verify the zones that were created
	expectedZones := []string{"eap-noob.arpa", "arpa"}
	for _, expected := range expectedZones {
		if _, ok := zones[expected]; !ok {
			t.Errorf("expected zone %q to be created", expected)
		}
	}

	// Verify NO invalid zones were created
	invalidZones := []string{" (deprecated)", "(deprecated)", "arpa. (deprecated)", " (DEPRECATED)"}
	for _, invalid := range invalidZones {
		if _, ok := zones[invalid]; ok {
			t.Errorf("invalid zone %q should NOT have been created", invalid)
		}
	}

	// Verify total zone count
	if len(zones) != 2 {
		t.Errorf("expected 2 zones, got %d: %v", len(zones), zones)
	}
}
