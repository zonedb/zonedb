package build

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestCompareVersions(t *testing.T) {
	tests := []struct {
		a, b string
		want int
	}{
		{"1.0", "1.1", -1},
		{"1.2", "1.1", 1},
		{"2.0", "1.9", 1},
		{"3.0", "3.0", 0},
		{"1.0", "2.0", -1},
		{"2.6", "1.2", 1},
		{"3.2", "3.0", 1},
	}
	for _, tt := range tests {
		got := compareVersions(tt.a, tt.b)
		if got != tt.want {
			t.Errorf("compareVersions(%q, %q) = %d, want %d", tt.a, tt.b, got, tt.want)
		}
	}
}

func TestParseCentralNicIndex(t *testing.T) {
	f, err := os.Open("testdata/centralnic-idn-tables-index.html")
	if err != nil {
		t.Fatalf("open fixture: %v", err)
	}
	defer f.Close()

	doc, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		t.Fatalf("parse fixture: %v", err)
	}

	baseURL, _ := url.Parse("https://centralnicregistry.com")
	entries := parseCentralNicIndex(doc, baseURL)

	if len(entries) == 0 {
		t.Fatal("parseCentralNicIndex returned no entries")
	}

	// Arabic should be version-deduped: 1.1 and 1.2 in fixture, 1.2 wins
	ar, ok := entries["ar"]
	if !ok {
		t.Fatal("missing entry for 'ar'")
	}
	if ar.Version != "1.2" {
		t.Errorf("ar version = %q, want 1.2 (should pick highest)", ar.Version)
	}
	if ar.Type != "Language" {
		t.Errorf("ar type = %q, want Language", ar.Type)
	}
	if !strings.Contains(ar.URL, "/services/idn-table/ar/651") {
		t.Errorf("ar URL = %q, expected to contain /services/idn-table/ar/651", ar.URL)
	}

	// zyyy (emoji) should be present
	zyyy, ok := entries["zyyy"]
	if !ok {
		t.Fatal("missing entry for 'zyyy' (emoji)")
	}
	if zyyy.Type != "Script" {
		t.Errorf("zyyy type = %q, want Script", zyyy.Type)
	}

	// latn should be version-deduped (multiple versions in index, highest wins)
	latn, ok := entries["latn"]
	if !ok {
		t.Fatal("missing entry for 'latn'")
	}
	if compareVersions(latn.Version, "1.0") <= 0 {
		t.Errorf("latn version = %q, expected > 1.0", latn.Version)
	}
}

func TestParseCentralNicDetailPage_Language(t *testing.T) {
	f, err := os.Open("testdata/centralnic-idn-table-detail-ar.html")
	if err != nil {
		t.Fatalf("open fixture: %v", err)
	}
	defer f.Close()

	doc, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		t.Fatalf("parse fixture: %v", err)
	}

	zones := parseCentralNicDetailPage(doc)
	if len(zones) == 0 {
		t.Fatal("parseCentralNicDetailPage returned no zones")
	}

	// The Arabic detail page should list these zones
	expected := map[string]bool{
		"best":    false,
		"qpon":    false,
		"tickets": false,
		"uno":     false,
	}
	for _, z := range zones {
		if _, ok := expected[z]; ok {
			expected[z] = true
		}
	}
	for zone, found := range expected {
		if !found {
			t.Errorf("expected zone %q not found in Arabic detail page", zone)
		}
	}
}

func TestParseCentralNicDetailPage_Script(t *testing.T) {
	f, err := os.Open("testdata/centralnic-idn-table-detail-latn.html")
	if err != nil {
		t.Fatalf("open fixture: %v", err)
	}
	defer f.Close()

	doc, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		t.Fatalf("parse fixture: %v", err)
	}

	zones := parseCentralNicDetailPage(doc)
	if len(zones) == 0 {
		t.Fatal("parseCentralNicDetailPage returned no zones for Latin script")
	}

	found := false
	for _, z := range zones {
		if z == "aquarelle" {
			found = true
		}
	}
	if !found {
		t.Errorf("expected zone 'aquarelle' in Latin detail page, got %v", zones)
	}
}

// centralNicTestHandler serves the full HTML fixtures at the paths
// that appear in the index page's hrefs.
func centralNicTestHandler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/idn-tables/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "testdata/centralnic-idn-tables-index.html")
	})
	mux.HandleFunc("/services/idn-table/ar/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "testdata/centralnic-idn-table-detail-ar.html")
	})
	mux.HandleFunc("/services/idn-table/latn/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "testdata/centralnic-idn-table-detail-latn.html")
	})
	// For any other detail page, serve an empty page (no zones)
	mux.HandleFunc("/services/idn-table/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(`<!doctype html><html><body></body></html>`))
	})

	return mux
}

func TestFetchIDNTablesFromCentralNic_IANAPrecedence(t *testing.T) {
	srv := httptest.NewServer(centralNicTestHandler())
	defer srv.Close()

	origIndexURL := centralNicIndexURL
	origBaseURL := centralNicBaseURL
	centralNicIndexURL = srv.URL + "/idn-tables/"
	centralNicBaseURL = srv.URL
	defer func() {
		centralNicIndexURL = origIndexURL
		centralNicBaseURL = origBaseURL
	}()

	// Pre-populate "best" with an IANA-sourced "ar" policy
	origIANABase := ianaBaseURL
	defer func() { ianaBaseURL = origIANABase }()
	ianaBaseURL = "https://www.iana.org"

	zones := map[string]*Zone{
		"best": {
			Domain: "best",
			Policies: []Policy{
				{Type: TypeIDNTable, Key: "ar", Value: "https://www.iana.org/domains/idn-tables/tables/best_ar_1.2.txt", Source: "https://www.iana.org/domains/idn-tables"},
			},
			Languages: []string{"ar"},
		},
		"qpon":    {Domain: "qpon"},
		"tickets": {Domain: "tickets"},
	}

	if err := FetchIDNTablesFromCentralNic(zones, nil); err != nil {
		t.Fatalf("FetchIDNTablesFromCentralNic: %v", err)
	}

	// "best" should still have the IANA-sourced "ar" policy, not CentralNic
	bestPolicies := idnPolicies(zones["best"])
	if v, ok := bestPolicies["ar"]; !ok {
		t.Error("best should have 'ar' policy")
	} else if !strings.HasPrefix(v, "https://www.iana.org/") {
		t.Errorf("best 'ar' policy should still be IANA-sourced, got %q", v)
	}

	// "qpon" should have CentralNic-sourced "ar" policy (no IANA precedence)
	qponPolicies := idnPolicies(zones["qpon"])
	if _, ok := qponPolicies["ar"]; !ok {
		t.Error("qpon should have 'ar' policy from CentralNic")
	}
}

func TestFetchIDNTablesFromCentralNic_StaleRemoval(t *testing.T) {
	srv := httptest.NewServer(centralNicTestHandler())
	defer srv.Close()

	origIndexURL := centralNicIndexURL
	origBaseURL := centralNicBaseURL
	centralNicIndexURL = srv.URL + "/idn-tables/"
	centralNicBaseURL = srv.URL
	defer func() {
		centralNicIndexURL = origIndexURL
		centralNicBaseURL = origBaseURL
	}()

	origIANABase := ianaBaseURL
	defer func() { ianaBaseURL = origIANABase }()
	ianaBaseURL = "https://www.iana.org"

	// Pre-populate with a stale CentralNic policy that should be removed
	staleCNSource := srv.URL + "/idn-tables/"
	zones := map[string]*Zone{
		"best": {
			Domain: "best",
			Policies: []Policy{
				{Type: TypeIDNTable, Key: "stale-lang", Value: srv.URL + "/stale", Source: staleCNSource},
			},
			Languages: []string{"stale-lang"},
		},
		"qpon": {Domain: "qpon"},
	}

	// Also add a non-CentralNic policy to "best" that should survive
	zones["best"].Policies = append(zones["best"].Policies,
		Policy{Type: TypeIDNTable, Key: "mul-Latn", Value: "https://example.com/custom.txt"},
	)
	zones["best"].Languages = append(zones["best"].Languages, "mul-Latn")

	if err := FetchIDNTablesFromCentralNic(zones, nil); err != nil {
		t.Fatalf("FetchIDNTablesFromCentralNic: %v", err)
	}

	bestPolicies := idnPolicies(zones["best"])

	// Stale CentralNic policy should be removed
	if _, ok := bestPolicies["stale-lang"]; ok {
		t.Error("stale CentralNic policy 'stale-lang' should have been removed")
	}

	// Non-CentralNic policy should be preserved
	if _, ok := bestPolicies["mul-Latn"]; !ok {
		t.Error("non-CentralNic policy 'mul-Latn' should be preserved")
	}

	// Fresh CentralNic data should be present (ar from the fixture)
	if _, ok := bestPolicies["ar"]; !ok {
		t.Error("fresh CentralNic policy 'ar' should have been added")
	}
}
