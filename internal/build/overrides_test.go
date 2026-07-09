package build

import (
	"os"
	"path/filepath"
	"slices"
	"testing"
)

// writeOverrides points BaseDir at a temp dir containing content/overrides.hujson
// with body, restoring BaseDir when the test finishes.
func writeOverrides(t *testing.T, body string) {
	t.Helper()
	dir := t.TempDir()
	if err := os.MkdirAll(filepath.Join(dir, "content"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, OverridesFile), []byte(body), 0o644); err != nil {
		t.Fatal(err)
	}
	prev := BaseDir
	BaseDir = dir
	t.Cleanup(func() { BaseDir = prev })
}

func TestReadOverrides(t *testing.T) {
	writeOverrides(t, `{
		// HuJSON comment
		"brand_exempt": {
			"nexus": "open",
			".MONSTER": "odd casing and a leading dot, plus a trailing comma",
		}
	}`)
	o, err := ReadOverrides()
	if err != nil {
		t.Fatalf("ReadOverrides: %v", err)
	}
	if got, want := o.BrandExempt["nexus"], "open"; got != want {
		t.Errorf("nexus reason = %q, want %q", got, want)
	}
	if _, ok := o.BrandExempt["monster"]; !ok {
		t.Errorf(".MONSTER not normalized to monster; keys = %v", o.BrandExempt)
	}
}

func TestReadOverridesUnknownField(t *testing.T) {
	writeOverrides(t, `{ "brand_exempts": { "nexus": "typo in category name" } }`)
	if _, err := ReadOverrides(); err == nil {
		t.Fatal("expected error for unknown top-level field, got nil")
	}
}

func TestReadOverridesMissingFile(t *testing.T) {
	prev := BaseDir
	BaseDir = t.TempDir()
	t.Cleanup(func() { BaseDir = prev })
	o, err := ReadOverrides()
	if err != nil {
		t.Fatalf("ReadOverrides on missing file: %v", err)
	}
	if len(o.BrandExempt) != 0 {
		t.Errorf("expected empty overrides, got %v", o.BrandExempt)
	}
}

func TestApplyBrandExempt(t *testing.T) {
	zones := map[string]*Zone{
		"nexus":  {Domain: "nexus", Tags: []string{"brand", "generic"}},
		"google": {Domain: "google", Tags: []string{"brand"}},
	}
	o := &Overrides{BrandExempt: map[string]string{"nexus": "open"}}
	if errs := o.ApplyBrandExempt(zones); len(errs) != 0 {
		t.Fatalf("unexpected errors: %v", errs)
	}
	if slices.Contains(zones["nexus"].Tags, TagBrand) {
		t.Errorf("nexus should be de-branded, tags = %v", zones["nexus"].Tags)
	}
	if !slices.Contains(zones["google"].Tags, TagBrand) {
		t.Errorf("google (not exempt) should stay branded, tags = %v", zones["google"].Tags)
	}
}

// TestOverridesRepoFile guards the committed overrides file: it must parse, and
// every entry must resolve to a real zone that ends up de-branded.
func TestOverridesRepoFile(t *testing.T) {
	prev := BaseDir
	BaseDir = filepath.Join("..", "..") // repo root, relative to internal/build
	t.Cleanup(func() { BaseDir = prev })

	o, err := ReadOverrides()
	if err != nil {
		t.Fatalf("ReadOverrides(%s): %v", OverridesFile, err)
	}
	for _, zone := range []string{"nexus", "monster", "sbs", "diy", "food"} {
		if _, ok := o.BrandExempt[zone]; !ok {
			t.Errorf("brand_exempt missing seeded zone %q", zone)
		}
	}

	// Every brand_exempt entry must resolve to a real zone; ApplyBrandExempt
	// returns an error per unknown zone, so zero errors == the file is clean.
	zones, errs := ReadZones()
	if len(errs) > 0 {
		t.Fatalf("ReadZones: %v", errs)
	}
	if aerrs := o.ApplyBrandExempt(zones); len(aerrs) > 0 {
		t.Errorf("brand_exempt has unknown zone(s): %v", aerrs)
	}
	for _, zone := range []string{"nexus", "monster", "sbs", "diy", "food"} {
		if z := zones[zone]; z != nil && z.IsBrand() {
			t.Errorf("%s still tagged brand after ApplyBrandExempt", zone)
		}
	}
}

func TestApplyBrandExemptUnknownZone(t *testing.T) {
	zones := map[string]*Zone{"nexus": {Domain: "nexus", Tags: []string{"brand"}}}
	o := &Overrides{BrandExempt: map[string]string{"nexux": "misspelled"}}
	if errs := o.ApplyBrandExempt(zones); len(errs) != 1 {
		t.Fatalf("expected 1 error for unknown zone, got %d: %v", len(errs), errs)
	}
}
