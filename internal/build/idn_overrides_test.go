package build

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApplyIDNOverrides(t *testing.T) {
	// Use the full fixture to populate the source zone with real
	// IANA IDN table data, then verify the override copies it.
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

	origOverrides := IDNOverrides
	IDNOverrides = map[string]string{"target": "source"}
	defer func() { IDNOverrides = origOverrides }()

	zones := map[string]*Zone{
		"source": {Domain: "aaa"},  // .aaa exists in the fixture
		"target": {Domain: "target"},
	}

	// Populate source from fixture.
	// Note: source has Domain "aaa" so FetchIDNTablesFromIANA matches
	// IANA entries for .aaa. But the zones map key is "source" so the
	// fetch won't find it. We need to use the real domain.
	// Let's restructure: use real domain names.
	zones = map[string]*Zone{
		"aaa": {Domain: "aaa"},
		"cc":  {Domain: "cc"},
	}
	IDNOverrides = map[string]string{"cc": "aaa"}

	if err := FetchIDNTablesFromIANA(zones); err != nil {
		t.Fatalf("FetchIDNTablesFromIANA: %v", err)
	}

	sourcePolicies := idnPolicies(zones["aaa"])
	if len(sourcePolicies) == 0 {
		t.Fatal("source zone has no IDN policies after fetch")
	}

	// .cc is not in the IANA fixture, so it should have no policies yet.
	if len(zones["cc"].Policies) != 0 {
		t.Fatalf("target zone should have no policies before override, got %d", len(zones["cc"].Policies))
	}

	ApplyIDNOverrides(zones)

	targetPolicies := idnPolicies(zones["cc"])

	// Every IDN policy from the source should now exist on the target.
	for key, wantValue := range sourcePolicies {
		gotValue, ok := targetPolicies[key]
		if !ok {
			t.Errorf("expected policy %q on target, not found", key)
		} else if gotValue != wantValue {
			t.Errorf("policy %q value = %q, want %q", key, gotValue, wantValue)
		}
	}

	// Target should have the same languages as source.
	sourceLangs := make(map[string]bool)
	for _, l := range zones["aaa"].Languages {
		sourceLangs[l] = true
	}
	targetLangs := make(map[string]bool)
	for _, l := range zones["cc"].Languages {
		targetLangs[l] = true
	}
	for lang := range sourceLangs {
		if !targetLangs[lang] {
			t.Errorf("expected language %q on target, not found", lang)
		}
	}
}

func TestApplyIDNOverrides_PreservesExisting(t *testing.T) {
	// Target zone has a pre-existing non-IANA policy. The override
	// should add source policies without removing the existing one.
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

	origOverrides := IDNOverrides
	IDNOverrides = map[string]string{"cc": "aaa"}
	defer func() { IDNOverrides = origOverrides }()

	existingPolicy := Policy{Type: TypeIDNTable, Key: "mul-Latn", Value: "https://example.com/cc-latn.txt"}

	zones := map[string]*Zone{
		"aaa": {Domain: "aaa"},
		"cc":  {Domain: "cc", Policies: []Policy{existingPolicy}, Languages: []string{"mul-Latn"}},
	}

	if err := FetchIDNTablesFromIANA(zones); err != nil {
		t.Fatalf("FetchIDNTablesFromIANA: %v", err)
	}

	ApplyIDNOverrides(zones)

	targetPolicies := idnPolicies(zones["cc"])

	// The pre-existing policy should still be there.
	if got, ok := targetPolicies["mul-Latn"]; !ok {
		t.Error("pre-existing policy mul-Latn should be preserved")
	} else if got != existingPolicy.Value {
		t.Errorf("pre-existing policy value = %q, want %q", got, existingPolicy.Value)
	}

	// Source policies should also be present.
	sourcePolicies := idnPolicies(zones["aaa"])
	for key, wantValue := range sourcePolicies {
		if key == "mul-Latn" {
			continue // already checked, target's own value takes precedence
		}
		gotValue, ok := targetPolicies[key]
		if !ok {
			t.Errorf("expected source policy %q on target, not found", key)
		} else if gotValue != wantValue {
			t.Errorf("policy %q value = %q, want %q", key, gotValue, wantValue)
		}
	}
}
