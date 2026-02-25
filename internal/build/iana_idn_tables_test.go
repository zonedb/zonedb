package build

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// idnPolicies returns the set of idn-table policy keys for a zone.
func idnPolicies(z *Zone) map[string]string {
	m := make(map[string]string)
	for _, p := range z.Policies {
		if p.Type == TypeIDNTable {
			m[p.Key] = p.Value
		}
	}
	return m
}

func TestFetchIDNTablesFromIANA_StaleRemoval(t *testing.T) {
	// Single server for the whole test so URLs are stable across fetches.
	srv := httptest.NewServer(http.FileServer(http.Dir("testdata")))
	defer srv.Close()

	origTablesURL := ianaTablesURL
	origBaseURL := ianaBaseURL
	ianaBaseURL = srv.URL
	defer func() {
		ianaTablesURL = origTablesURL
		ianaBaseURL = origBaseURL
	}()

	// Step 1: Populate .aaa from the full fixture.
	ianaTablesURL = srv.URL + "/idn-tables-full.html"
	zones := map[string]*Zone{
		"aaa": {Domain: "aaa"},
	}
	if err := FetchIDNTablesFromIANA(zones, nil); err != nil {
		t.Fatalf("fetch from full fixture: %v", err)
	}

	fullPolicies := idnPolicies(zones["aaa"])
	if len(fullPolicies) == 0 {
		t.Fatal("full fixture produced no idn-table policies for .aaa")
	}

	// Step 2: Add a non-IANA policy that must survive the next fetch.
	nonIANAKey := "mul-Latn"
	nonIANAValue := "https://example.com/aaa-latn.txt"
	zones["aaa"].AddPolicy(TypeIDNTable, nonIANAKey, nonIANAValue, "", "")
	zones["aaa"].Languages = append(zones["aaa"].Languages, nonIANAKey)

	// Step 3: Re-fetch from the reduced fixture (some entries removed).
	ianaTablesURL = srv.URL + "/idn-tables-reduced.html"
	if err := FetchIDNTablesFromIANA(zones, nil); err != nil {
		t.Fatalf("fetch from reduced fixture: %v", err)
	}

	afterPolicies := idnPolicies(zones["aaa"])

	// Step 4: Get expected IANA entries by parsing the reduced fixture
	// against a fresh zone.
	expectedZones := map[string]*Zone{
		"aaa": {Domain: "aaa"},
	}
	if err := FetchIDNTablesFromIANA(expectedZones, nil); err != nil {
		t.Fatalf("fetch expected from reduced fixture: %v", err)
	}

	expectedPolicies := idnPolicies(expectedZones["aaa"])

	// Verify: entries in the full fixture but not the reduced fixture
	// should have been removed.
	for key := range fullPolicies {
		if _, inReduced := expectedPolicies[key]; inReduced {
			continue
		}
		if _, still := afterPolicies[key]; still {
			t.Errorf("stale IANA policy %q should have been removed", key)
		}
	}

	// Verify: all entries from the reduced fixture are present.
	for key, wantValue := range expectedPolicies {
		gotValue, ok := afterPolicies[key]
		if !ok {
			t.Errorf("expected IANA policy %q not found", key)
		} else if gotValue != wantValue {
			t.Errorf("policy %q value = %q, want %q", key, gotValue, wantValue)
		}
	}

	// Verify: non-IANA policy is preserved with its original value.
	if got, ok := afterPolicies[nonIANAKey]; !ok {
		t.Error("non-IANA policy should be preserved")
	} else if got != nonIANAValue {
		t.Errorf("non-IANA policy value = %q, want %q", got, nonIANAValue)
	}

	// Verify: languages include non-IANA entry.
	langSet := make(map[string]bool)
	for _, l := range zones["aaa"].Languages {
		langSet[l] = true
	}
	if !langSet[nonIANAKey] {
		t.Error("non-IANA language should be preserved")
	}
}

// TestFetchIDNTablesFromIANA_PreservesIndependentLanguages is a regression test
// for the bug where z.Languages was rebuilt solely from remaining policy keys,
// dropping languages that were independently set (e.g., from CLDR or manual
// curation). This is the .ru/.рф scenario: a zone has a language like "ru-RU"
// that was set independently of any IDN table policy.
func TestFetchIDNTablesFromIANA_PreservesIndependentLanguages(t *testing.T) {
	srv := httptest.NewServer(http.FileServer(http.Dir("testdata")))
	defer srv.Close()

	origTablesURL := ianaTablesURL
	origBaseURL := ianaBaseURL
	ianaBaseURL = srv.URL
	defer func() {
		ianaTablesURL = origTablesURL
		ianaBaseURL = origBaseURL
	}()

	// Step 1: Populate .aaa from the full fixture (gives it IANA policies + languages).
	ianaTablesURL = srv.URL + "/idn-tables-full.html"
	zones := map[string]*Zone{
		"aaa": {Domain: "aaa"},
	}
	if err := FetchIDNTablesFromIANA(zones, nil); err != nil {
		t.Fatalf("fetch from full fixture: %v", err)
	}

	// Step 2: Add an independent language not backed by any policy.
	// This simulates what CLDR automation does: adding "ru-RU" to a zone
	// based on territory data, not IDN table data.
	independentLang := "ru-RU"
	zones["aaa"].Languages = append(zones["aaa"].Languages, independentLang)

	// Step 3: Re-fetch from the reduced fixture. Some IANA policies are removed,
	// but the independent language must survive.
	ianaTablesURL = srv.URL + "/idn-tables-reduced.html"
	if err := FetchIDNTablesFromIANA(zones, nil); err != nil {
		t.Fatalf("fetch from reduced fixture: %v", err)
	}

	langSet := make(map[string]bool)
	for _, l := range zones["aaa"].Languages {
		langSet[l] = true
	}
	if !langSet[independentLang] {
		t.Errorf("independent language %q was removed — languages = %v", independentLang, zones["aaa"].Languages)
	}
}

// TestFetchIDNTablesFromIANA_RemovesOrphanedLanguages verifies that when an
// IANA policy is removed and no other policy backs that language, the language
// IS removed from z.Languages (the stale-removal should still clean up policy-
// derived languages that become orphaned).
func TestFetchIDNTablesFromIANA_RemovesOrphanedLanguages(t *testing.T) {
	srv := httptest.NewServer(http.FileServer(http.Dir("testdata")))
	defer srv.Close()

	origTablesURL := ianaTablesURL
	origBaseURL := ianaBaseURL
	ianaBaseURL = srv.URL
	defer func() {
		ianaTablesURL = origTablesURL
		ianaBaseURL = origBaseURL
	}()

	// Step 1: Populate .aaa from the full fixture.
	ianaTablesURL = srv.URL + "/idn-tables-full.html"
	zones := map[string]*Zone{
		"aaa": {Domain: "aaa"},
	}
	if err := FetchIDNTablesFromIANA(zones, nil); err != nil {
		t.Fatalf("fetch from full fixture: %v", err)
	}

	fullLangs := make(map[string]bool)
	for _, l := range zones["aaa"].Languages {
		fullLangs[l] = true
	}
	if len(fullLangs) == 0 {
		t.Fatal("full fixture produced no languages")
	}

	// Step 2: Re-fetch from the reduced fixture.
	ianaTablesURL = srv.URL + "/idn-tables-reduced.html"
	if err := FetchIDNTablesFromIANA(zones, nil); err != nil {
		t.Fatalf("fetch from reduced fixture: %v", err)
	}

	// Step 3: Get the expected languages from a clean parse of the reduced fixture.
	expectedZones := map[string]*Zone{
		"aaa": {Domain: "aaa"},
	}
	if err := FetchIDNTablesFromIANA(expectedZones, nil); err != nil {
		t.Fatalf("fetch expected: %v", err)
	}
	expectedLangs := make(map[string]bool)
	for _, l := range expectedZones["aaa"].Languages {
		expectedLangs[l] = true
	}

	// Languages that were in the full fixture but not the reduced fixture
	// should be removed (they were only backed by IANA policies that got removed).
	afterLangs := make(map[string]bool)
	for _, l := range zones["aaa"].Languages {
		afterLangs[l] = true
	}
	for lang := range fullLangs {
		if expectedLangs[lang] {
			continue // Still expected
		}
		if afterLangs[lang] {
			t.Errorf("orphaned language %q should have been removed", lang)
		}
	}
}

func TestFetchIDNTablesFromIANA_NonIANAOnly(t *testing.T) {
	// A zone not listed on the IANA IDN tables page, with policies
	// from another source, must be left completely untouched.
	srv := httptest.NewServer(http.FileServer(http.Dir("testdata")))
	defer srv.Close()

	origTablesURL := ianaTablesURL
	origBaseURL := ianaBaseURL
	ianaTablesURL = srv.URL + "/idn-tables-full.html"
	ianaBaseURL = srv.URL
	defer func() {
		ianaTablesURL = origTablesURL
		ianaBaseURL = origBaseURL
	}()

	origPolicy := Policy{Type: TypeIDNTable, Key: "mul-Latn", Value: "https://example.com/idn-table.txt", Comment: "from registry"}

	zones := map[string]*Zone{
		"notiniana": {
			Domain:    "notiniana",
			Languages: []string{"mul-Latn"},
			Policies:  []Policy{origPolicy},
		},
	}

	if err := FetchIDNTablesFromIANA(zones, nil); err != nil {
		t.Fatalf("FetchIDNTablesFromIANA: %v", err)
	}

	z := zones["notiniana"]

	if len(z.Policies) != 1 {
		t.Fatalf("expected 1 policy, got %d", len(z.Policies))
	}
	if z.Policies[0] != origPolicy {
		t.Errorf("policy was modified: got %+v, want %+v", z.Policies[0], origPolicy)
	}
	if len(z.Languages) != 1 || z.Languages[0] != "mul-Latn" {
		t.Errorf("languages were modified: %v", z.Languages)
	}
}

func TestFetchIDNTablesFromIANA_Full(t *testing.T) {
	// Smoke test: parsing the full IANA fixture produces policies
	// and languages for .aaa.
	srv := httptest.NewServer(http.FileServer(http.Dir("testdata")))
	defer srv.Close()

	origTablesURL := ianaTablesURL
	origBaseURL := ianaBaseURL
	ianaTablesURL = srv.URL + "/idn-tables-full.html"
	ianaBaseURL = srv.URL
	defer func() {
		ianaTablesURL = origTablesURL
		ianaBaseURL = origBaseURL
	}()

	zones := map[string]*Zone{
		"aaa": {Domain: "aaa"},
	}

	if err := FetchIDNTablesFromIANA(zones, nil); err != nil {
		t.Fatalf("FetchIDNTablesFromIANA: %v", err)
	}

	z := zones["aaa"]
	policies := idnPolicies(z)
	if len(policies) == 0 {
		t.Fatal("expected idn-table policies for .aaa, got none")
	}
	if len(z.Languages) == 0 {
		t.Fatal("expected languages for .aaa, got none")
	}
}
