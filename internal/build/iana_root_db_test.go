package build

import (
	"os"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func loadRootDBFixture(t *testing.T) *goquery.Document {
	t.Helper()
	f, err := os.Open("testdata/iana-root-db.html")
	if err != nil {
		t.Fatalf("opening test fixture: %v", err)
	}
	defer f.Close()

	doc, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		t.Fatalf("parsing HTML: %v", err)
	}
	return doc
}

// zonesFromEntries builds a zones map keyed by normalized (Unicode) domain,
// matching the convention used by ReadZones/ReadMetadata in production.
func zonesFromEntries(entries []RootDBEntry) map[string]*Zone {
	zones := make(map[string]*Zone, len(entries))
	for _, e := range entries {
		domain := Normalize(e.Punycode)
		zones[domain] = &Zone{Domain: domain}
	}
	return zones
}

// applyRootDBFixture parses the fixture, builds a zones map, applies
// applyRootDBIndex, and returns both the entries and zones for subtests.
func applyRootDBFixture(t *testing.T) ([]RootDBEntry, map[string]*Zone) {
	t.Helper()
	doc := loadRootDBFixture(t)
	entries, err := ParseRootDBIndex(doc)
	if err != nil {
		t.Fatalf("ParseRootDBIndex: %v", err)
	}
	zones := zonesFromEntries(entries)

	if err := applyRootDBIndex(zones, doc); err != nil {
		t.Fatalf("applyRootDBIndex: %v", err)
	}
	return entries, zones
}

func TestParseRootDBIndex(t *testing.T) {
	doc := loadRootDBFixture(t)
	entries, err := ParseRootDBIndex(doc)
	if err != nil {
		t.Fatalf("ParseRootDBIndex: %v", err)
	}

	if len(entries) < 1000 {
		t.Fatalf("expected at least 1000 entries, got %d", len(entries))
	}

	// Check for known entries
	var foundRU, foundXnP1ai, foundCom bool
	for _, e := range entries {
		switch e.Punycode {
		case "ru":
			foundRU = true
			if e.Type != "country-code" {
				t.Errorf(".ru: type = %q, want %q", e.Type, "country-code")
			}
			if e.Organisation != "Coordination Center for TLD RU" {
				t.Errorf(".ru: organisation = %q, want %q", e.Organisation, "Coordination Center for TLD RU")
			}
		case "xn--p1ai":
			foundXnP1ai = true
			if e.Domain != "рф" {
				t.Errorf(".рф: domain = %q, want %q", e.Domain, "рф")
			}
			if e.Type != "country-code" {
				t.Errorf(".рф: type = %q, want %q", e.Type, "country-code")
			}
			if e.Organisation != "Coordination Center for TLD RU" {
				t.Errorf(".рф: organisation = %q, want %q", e.Organisation, "Coordination Center for TLD RU")
			}
		case "com":
			foundCom = true
			if e.Type != "generic" {
				t.Errorf(".com: type = %q, want %q", e.Type, "generic")
			}
		}
	}

	if !foundRU {
		t.Error("entry for .ru not found")
	}
	if !foundXnP1ai {
		t.Error("entry for .рф (xn--p1ai) not found")
	}
	if !foundCom {
		t.Error("entry for .com not found")
	}

	// HTML entity decoding: .am registry operator contains literal quotation
	// marks from &quot;Internet Society&quot; in the HTML source.
	t.Run("HTMLEntityDecoding", func(t *testing.T) {
		for _, e := range entries {
			if e.Punycode == "am" {
				if !strings.Contains(e.Organisation, `"Internet Society"`) {
					t.Errorf(".am organisation = %q, want to contain literal quotes around Internet Society", e.Organisation)
				}
				return
			}
		}
		t.Error("entry for .am not found")
	})
}

func TestApplyRootDBIndex(t *testing.T) {
	entries, zones := applyRootDBFixture(t)

	t.Run("RU", func(t *testing.T) {
		ru := zones["ru"]
		if ru.RegistryOperator != "Coordination Center for TLD RU" {
			t.Errorf(".ru registryOperator = %q, want %q", ru.RegistryOperator, "Coordination Center for TLD RU")
		}
		// .ru should get domainIdn linking to .рф
		idnSet := NewSet(ru.DomainIdn...)
		if !idnSet["рф"] {
			t.Errorf(".ru domainIdn = %v, expected to contain рф", ru.DomainIdn)
		}
	})

	t.Run("IDN_RU", func(t *testing.T) {
		rf := zones["рф"]
		if rf.RegistryOperator != "Coordination Center for TLD RU" {
			t.Errorf(".рф registryOperator = %q, want %q", rf.RegistryOperator, "Coordination Center for TLD RU")
		}
		if rf.DomainPunycode != "xn--p1ai" {
			t.Errorf(".рф domainPunycode = %q, want %q", rf.DomainPunycode, "xn--p1ai")
		}
		if rf.DomainAscii != "ru" {
			t.Errorf(".рф domainAscii = %q, want %q", rf.DomainAscii, "ru")
		}
	})

	t.Run("GenericTLD", func(t *testing.T) {
		com := zones["com"]
		if com.RegistryOperator == "" {
			t.Error(".com registryOperator is empty")
		}
		if com.DomainPunycode != "" {
			t.Errorf(".com domainPunycode = %q, want empty", com.DomainPunycode)
		}
	})

	t.Run("India_MultipleIDN", func(t *testing.T) {
		in := zones["in"]
		if len(in.DomainIdn) < 10 {
			t.Errorf(".in domainIdn has %d entries, expected at least 10 (India has 15 IDN ccTLDs)", len(in.DomainIdn))
		}
	})

	t.Run("TRA_Disambiguation", func(t *testing.T) {
		oman := zones[Normalize("xn--mgb9awbf")]
		if oman.DomainAscii != "om" {
			t.Errorf(".عمان domainAscii = %q, want %q", oman.DomainAscii, "om")
		}
		bahrain := zones[Normalize("xn--mgbcpq6gpa1a")]
		if bahrain.DomainAscii != "bh" {
			t.Errorf(".البحرين domainAscii = %q, want %q", bahrain.DomainAscii, "bh")
		}
	})

	// EU IDN variants (Greek .ευ and Cyrillic .ею) both map to domainAscii "eu",
	// and .eu lists both in domainIdn.
	t.Run("EU_MultipleIDN", func(t *testing.T) {
		for _, tc := range []struct {
			punycode string
			label    string
		}{
			{"xn--e1a4c", "Greek .ευ"},
			{"xn--qxa6a", "Cyrillic .ею"},
		} {
			domain := Normalize(tc.punycode)
			z := zones[domain]
			if z == nil {
				t.Errorf("%s: zone %q not found", tc.label, domain)
				continue
			}
			if z.DomainAscii != "eu" {
				t.Errorf("%s: domainAscii = %q, want %q", tc.label, z.DomainAscii, "eu")
			}
		}

		eu := zones["eu"]
		if eu == nil {
			t.Fatal("zone 'eu' not found")
		}
		idnSet := NewSet(eu.DomainIdn...)
		for _, want := range []string{"ευ", "ею"} {
			if !idnSet[want] {
				t.Errorf(".eu domainIdn = %v, expected to contain %q", eu.DomainIdn, want)
			}
		}
	})

	// Chinese IDN variants (simplified .中国 and traditional .中國) both map to
	// domainAscii "cn", and .cn lists both in domainIdn.
	t.Run("CN_MultipleIDN", func(t *testing.T) {
		for _, tc := range []struct {
			punycode string
			label    string
		}{
			{"xn--fiqs8s", "Simplified .中国"},
			{"xn--fiqz9s", "Traditional .中國"},
		} {
			domain := Normalize(tc.punycode)
			z := zones[domain]
			if z == nil {
				t.Errorf("%s: zone %q not found", tc.label, domain)
				continue
			}
			if z.DomainAscii != "cn" {
				t.Errorf("%s: domainAscii = %q, want %q", tc.label, z.DomainAscii, "cn")
			}
		}

		cn := zones["cn"]
		if cn == nil {
			t.Fatal("zone 'cn' not found")
		}
		idnSet := NewSet(cn.DomainIdn...)
		for _, want := range []string{"中国", "中國"} {
			if !idnSet[want] {
				t.Errorf(".cn domainIdn = %v, expected to contain %q", cn.DomainIdn, want)
			}
		}
	})

	// Every IDN ccTLD should get a domainAscii mapping. If this fails, a new
	// IDN ccTLD was added to the IANA root DB and needs an entry in
	// idnToAsciiOverrides.
	t.Run("AllIDNccTLDsHaveDomainAscii", func(t *testing.T) {
		for _, e := range entries {
			if e.Type != "country-code" || !strings.HasPrefix(e.Punycode, "xn--") {
				continue
			}
			z := zones[Normalize(e.Punycode)]
			if z.DomainAscii == "" {
				t.Errorf("IDN ccTLD %s (%s) has no domainAscii — add to idnToAsciiOverrides", e.Domain, e.Punycode)
			}
		}
	})

	// Every assigned TLD should get a registryOperator.
	t.Run("AllAssignedHaveRegistryOperator", func(t *testing.T) {
		for _, e := range entries {
			if e.Organisation == "Not assigned" {
				continue
			}
			z := zones[Normalize(e.Punycode)]
			if z.RegistryOperator == "" {
				t.Errorf("TLD %s (%s) has no registryOperator", e.Domain, e.Punycode)
			}
		}
	})

	// Every TLD should get a type-based tag matching its IANA classification.
	t.Run("TypeTags", func(t *testing.T) {
		typeToTag := map[string]string{
			"country-code":     TagCountry,
			"generic":          TagGeneric,
			"generic-restricted": TagGeneric,
			"sponsored":        TagSponsored,
			"infrastructure":   TagInfrastructure,
			"test":             TagTest,
		}
		for _, e := range entries {
			wantTag, ok := typeToTag[e.Type]
			if !ok {
				continue
			}
			z := zones[Normalize(e.Punycode)]
			tags := NewSet(z.Tags...)
			if !tags[wantTag] {
				t.Errorf("TLD %s (%s) type %q: missing tag %q, got %v", e.Domain, e.Punycode, e.Type, wantTag, z.Tags)
			}
			// country-code TLDs should also have "geo"
			if e.Type == "country-code" && !tags[TagGeo] {
				t.Errorf("country-code TLD %s (%s) missing %q tag", e.Domain, e.Punycode, TagGeo)
			}
		}
	})
}

// TestIDNccTLDOverridesAreValid verifies that every entry in the override map
// references an actual IDN ccTLD punycode from the fixture and maps to a valid
// ASCII ccTLD.
func TestIDNccTLDOverridesAreValid(t *testing.T) {
	doc := loadRootDBFixture(t)
	entries, err := ParseRootDBIndex(doc)
	if err != nil {
		t.Fatalf("ParseRootDBIndex: %v", err)
	}

	punycodes := make(map[string]bool, len(entries))
	asciiCCTLDs := make(map[string]bool)
	for _, e := range entries {
		punycodes[e.Punycode] = true
		if e.Type == "country-code" && !strings.HasPrefix(e.Punycode, "xn--") {
			asciiCCTLDs[e.Domain] = true
		}
	}

	for punycode, ascii := range idnToAsciiOverrides {
		if !punycodes[punycode] {
			t.Errorf("override %q → %q: punycode %q not found in IANA root DB", punycode, ascii, punycode)
		}
		if !asciiCCTLDs[ascii] {
			t.Errorf("override %q → %q: ASCII ccTLD %q not found in IANA root DB", punycode, ascii, ascii)
		}
	}
}
