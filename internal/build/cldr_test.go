package build

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestParseCLDRTerritoryInfo(t *testing.T) {
	f, err := os.Open("testdata/territoryInfo.json")
	if err != nil {
		t.Fatalf("opening test fixture: %v", err)
	}
	defer f.Close()

	info, err := ParseCLDRTerritoryInfo(f)
	if err != nil {
		t.Fatalf("ParseCLDRTerritoryInfo: %v", err)
	}

	if len(info.Supplemental.TerritoryInfo) < 200 {
		t.Fatalf("expected at least 200 territories, got %d", len(info.Supplemental.TerritoryInfo))
	}

	// Check that RU territory exists
	ru, ok := info.Supplemental.TerritoryInfo["RU"]
	if !ok {
		t.Fatal("territory RU not found")
	}
	if len(ru.LanguagePopulation) == 0 {
		t.Fatal("RU has no language population data")
	}

	// Check that "ru" is an official language in RU
	ruLang, ok := ru.LanguagePopulation["ru"]
	if !ok {
		t.Fatal("language 'ru' not found in territory RU")
	}
	if ruLang.OfficialStatus != "official" {
		t.Errorf("ru in RU: officialStatus = %q, want %q", ruLang.OfficialStatus, "official")
	}
}

func TestOfficialLanguages(t *testing.T) {
	f, err := os.Open("testdata/territoryInfo.json")
	if err != nil {
		t.Fatalf("opening test fixture: %v", err)
	}
	defer f.Close()

	info, err := ParseCLDRTerritoryInfo(f)
	if err != nil {
		t.Fatalf("ParseCLDRTerritoryInfo: %v", err)
	}

	tests := []struct {
		territory string
		wantLang  string // At least one expected language
	}{
		{"RU", "ru"},
		{"BR", "pt"},
		{"DE", "de"},
		{"JP", "ja"},
		{"CN", "zh"},
		{"FR", "fr"},
	}

	for _, tt := range tests {
		langs := info.OfficialLanguages(tt.territory)
		if len(langs) == 0 {
			t.Errorf("OfficialLanguages(%q) returned empty", tt.territory)
			continue
		}
		found := false
		for _, l := range langs {
			if l == tt.wantLang {
				found = true
			}
		}
		if !found {
			t.Errorf("OfficialLanguages(%q) = %v, expected to contain %q", tt.territory, langs, tt.wantLang)
		}
	}
}

func TestTerritoryFromCCTLD(t *testing.T) {
	tests := []struct {
		domain string
		want   string
	}{
		{"ru", "RU"},
		{"br", "BR"},
		{"uk", "GB"},
		{"de", "DE"},
		{"us", "US"},
	}

	for _, tt := range tests {
		got := TerritoryFromCCTLD(tt.domain)
		if got != tt.want {
			t.Errorf("TerritoryFromCCTLD(%q) = %q, want %q", tt.domain, got, tt.want)
		}
	}
}

func TestComposeBCP47Tags(t *testing.T) {
	tests := []struct {
		langs     []string
		territory string
		want      string // At least one expected tag
	}{
		{[]string{"ru"}, "RU", "ru-RU"},
		{[]string{"pt"}, "BR", "pt-BR"},
		{[]string{"de"}, "DE", "de-DE"},
		{[]string{"en"}, "GB", "en-GB"},
	}

	for _, tt := range tests {
		tags := ComposeBCP47Tags(tt.langs, tt.territory)
		if len(tags) == 0 {
			t.Errorf("ComposeBCP47Tags(%v, %q) returned empty", tt.langs, tt.territory)
			continue
		}
		found := false
		for _, tag := range tags {
			if tag == tt.want {
				found = true
			}
		}
		if !found {
			t.Errorf("ComposeBCP47Tags(%v, %q) = %v, expected to contain %q", tt.langs, tt.territory, tags, tt.want)
		}
	}
}

func TestCountryName(t *testing.T) {
	tests := []struct {
		territory string
		want      string
	}{
		{"RU", "Russia"},
		{"BR", "Brazil"},
		{"DE", "Germany"},
		{"GB", "United Kingdom"},
		{"JP", "Japan"},
	}

	for _, tt := range tests {
		got := CountryName(tt.territory)
		if got != tt.want {
			t.Errorf("CountryName(%q) = %q, want %q", tt.territory, got, tt.want)
		}
	}
}

func TestApplyCLDRMetadata(t *testing.T) {
	f, err := os.Open("testdata/territoryInfo.json")
	if err != nil {
		t.Fatalf("opening test fixture: %v", err)
	}
	defer f.Close()

	info, err := ParseCLDRTerritoryInfo(f)
	if err != nil {
		t.Fatalf("ParseCLDRTerritoryInfo: %v", err)
	}

	// Zones map keyed by Unicode domain, matching ReadZones convention.
	zones := map[string]*Zone{
		"ru": {
			Domain: "ru",
			Tags:   []string{"country", "geo"},
		},
		"рф": {
			Domain:      "рф",
			DomainAscii: "ru",
			Tags:        []string{"country", "geo"},
			Languages:   []string{"mul-Cyrl"},
		},
	}

	ApplyCLDRMetadata(zones, info)

	// .ru should get countryName and languages
	ru := zones["ru"]
	if ru.CountryName != "Russia" {
		t.Errorf(".ru countryName = %q, want %q", ru.CountryName, "Russia")
	}
	hasRuRU := false
	for _, l := range ru.Languages {
		if l == "ru-RU" {
			hasRuRU = true
		}
	}
	if !hasRuRU {
		t.Errorf(".ru languages = %v, expected to contain ru-RU", ru.Languages)
	}

	// .рф should get countryName and preserve mul-Cyrl while adding ru-RU
	rf := zones["рф"]
	if rf.CountryName != "Russia" {
		t.Errorf(".рф countryName = %q, want %q", rf.CountryName, "Russia")
	}
	hasMulCyrl := false
	hasRuRU = false
	for _, l := range rf.Languages {
		if l == "mul-Cyrl" {
			hasMulCyrl = true
		}
		if l == "ru-RU" {
			hasRuRU = true
		}
	}
	if !hasMulCyrl {
		t.Errorf(".рф languages = %v, expected to contain mul-Cyrl", rf.Languages)
	}
	if !hasRuRU {
		t.Errorf(".рф languages = %v, expected to contain ru-RU", rf.Languages)
	}
}

// TestApplyCLDRMetadata_IDNRequiresDomainAscii verifies that IDN ccTLDs without
// domainAscii set (i.e., before FetchRootDBIndex runs) do NOT get countryName
// or languages — ensuring the pipeline order dependency is explicit.
func TestApplyCLDRMetadata_IDNRequiresDomainAscii(t *testing.T) {
	f, err := os.Open("testdata/territoryInfo.json")
	if err != nil {
		t.Fatalf("opening test fixture: %v", err)
	}
	defer f.Close()

	info, err := ParseCLDRTerritoryInfo(f)
	if err != nil {
		t.Fatalf("ParseCLDRTerritoryInfo: %v", err)
	}

	// IDN ccTLD WITHOUT domainAscii — simulates running CLDR before root DB
	zones := map[string]*Zone{
		"рф": {
			Domain:    "рф",
			Tags:      []string{"country", "geo"},
			Languages: []string{"mul-Cyrl"},
		},
	}

	ApplyCLDRMetadata(zones, info)

	rf := zones["рф"]
	if rf.CountryName != "" {
		t.Errorf(".рф countryName = %q, want empty (domainAscii not set)", rf.CountryName)
	}
	// Languages should be unchanged (only mul-Cyrl, no ru-RU added)
	if len(rf.Languages) != 1 || rf.Languages[0] != "mul-Cyrl" {
		t.Errorf(".рф languages = %v, want [mul-Cyrl] (domainAscii not set)", rf.Languages)
	}
}

// TestPipelineOrder_RootDBThenCLDR is an integration test verifying that when
// FetchRootDBIndex runs first (setting domainAscii), then ApplyCLDRMetadata
// runs second, IDN ccTLDs correctly get countryName and CLDR-derived languages.
// This is the exact pipeline order in cmd/zonedb/main.go.
func TestPipelineOrder_RootDBThenCLDR(t *testing.T) {
	// Step 1: Parse root DB to set domainAscii
	doc := loadRootDBFixture(t)
	entries, err := ParseRootDBIndex(doc)
	if err != nil {
		t.Fatalf("ParseRootDBIndex: %v", err)
	}

	// Build zones map keyed by Unicode domain (matching production)
	zones := zonesFromEntries(entries)

	if err := applyRootDBIndex(zones, doc); err != nil {
		t.Fatalf("applyRootDBIndex: %v", err)
	}

	// Verify domainAscii was set (prerequisite for CLDR)
	rf := zones["рф"]
	if rf == nil {
		t.Fatal("zone рф not found")
	}
	if rf.DomainAscii != "ru" {
		t.Fatalf(".рф domainAscii = %q, want %q", rf.DomainAscii, "ru")
	}

	// Step 2: Apply CLDR metadata
	f, err := os.Open("testdata/territoryInfo.json")
	if err != nil {
		t.Fatalf("opening CLDR fixture: %v", err)
	}
	defer f.Close()

	cldr, err := ParseCLDRTerritoryInfo(f)
	if err != nil {
		t.Fatalf("ParseCLDRTerritoryInfo: %v", err)
	}

	ApplyCLDRMetadata(zones, cldr)

	// .рф should now have countryName and ru-RU
	if rf.CountryName != "Russia" {
		t.Errorf(".рф countryName = %q, want %q", rf.CountryName, "Russia")
	}
	langSet := NewSet(rf.Languages...)
	if !langSet["ru-RU"] {
		t.Errorf(".рф languages = %v, expected to contain ru-RU", rf.Languages)
	}

	// .ru should also have countryName and ru-RU
	ru := zones["ru"]
	if ru.CountryName != "Russia" {
		t.Errorf(".ru countryName = %q, want %q", ru.CountryName, "Russia")
	}
	ruLangSet := NewSet(ru.Languages...)
	if !ruLangSet["ru-RU"] {
		t.Errorf(".ru languages = %v, expected to contain ru-RU", ru.Languages)
	}

	// .eu IDN counterparts (Greek + Cyrillic) should both get countryName
	// via their shared domainAscii → "eu" → territory "EU"
	for _, idnDomain := range []string{"ею", "ευ"} {
		z := zones[idnDomain]
		if z == nil {
			t.Errorf("zone %q not found", idnDomain)
			continue
		}
		if z.DomainAscii != "eu" {
			t.Errorf(".%s domainAscii = %q, want %q", idnDomain, z.DomainAscii, "eu")
		}
		if z.CountryName != "European Union" {
			t.Errorf(".%s countryName = %q, want %q", idnDomain, z.CountryName, "European Union")
		}
	}
}

// TestFetchAndApplyCLDRMetadata_304 verifies that when the CLDR fetch returns
// 304 Not Modified, FetchAndApplyCLDRMetadata returns nil without modifying
// any zone data. This tests the scenario where the ETag cache has a stale entry
// from a previous run where domainAscii wasn't populated.
func TestFetchAndApplyCLDRMetadata_304(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotModified)
	}))
	defer srv.Close()

	origURL := cldrTerritoryInfoURL
	cldrTerritoryInfoURL = srv.URL + "/territoryInfo.json"
	defer func() { cldrTerritoryInfoURL = origURL }()

	zones := map[string]*Zone{
		"ru": {
			Domain: "ru",
			Tags:   []string{"country", "geo"},
		},
	}

	cache := NewETagCache("")
	// Pre-populate the cache with an ETag so the request includes If-None-Match
	cache.Set(cldrTerritoryInfoURL, `"some-etag"`)

	err := FetchAndApplyCLDRMetadata(zones, cache)
	if err != nil {
		t.Fatalf("FetchAndApplyCLDRMetadata: %v", err)
	}

	// Zones should be untouched — no countryName, no languages added
	ru := zones["ru"]
	if ru.CountryName != "" {
		t.Errorf(".ru countryName = %q, want empty (304 should skip apply)", ru.CountryName)
	}
	if len(ru.Languages) != 0 {
		t.Errorf(".ru languages = %v, want empty (304 should skip apply)", ru.Languages)
	}
}
