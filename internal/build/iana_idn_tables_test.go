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
