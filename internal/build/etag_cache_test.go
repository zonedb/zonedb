package build

import (
	"os"
	"path/filepath"
	"testing"
)

func TestETagCache_LoadSave(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "etags.json")

	c := NewETagCache(path)
	c.Set("https://example.com/a", `"abc123"`)
	c.Set("https://example.com/b", `"def456"`)

	if err := c.Save(); err != nil {
		t.Fatalf("Save: %v", err)
	}

	c2 := NewETagCache(path)
	if err := c2.Load(); err != nil {
		t.Fatalf("Load: %v", err)
	}

	if got := c2.Get("https://example.com/a"); got != `"abc123"` {
		t.Errorf("Get(a) = %q, want %q", got, `"abc123"`)
	}
	if got := c2.Get("https://example.com/b"); got != `"def456"` {
		t.Errorf("Get(b) = %q, want %q", got, `"def456"`)
	}
	if got := c2.Get("https://example.com/missing"); got != "" {
		t.Errorf("Get(missing) = %q, want empty", got)
	}
}

func TestETagCache_MissingFile(t *testing.T) {
	path := filepath.Join(t.TempDir(), "nonexistent", "etags.json")
	c := NewETagCache(path)
	if err := c.Load(); err != nil {
		t.Fatalf("Load from missing file should not error: %v", err)
	}
	if got := c.Get("https://example.com"); got != "" {
		t.Errorf("Get from empty cache = %q, want empty", got)
	}
}

func TestETagCache_SaveCreatesDir(t *testing.T) {
	dir := filepath.Join(t.TempDir(), "sub", "dir")
	path := filepath.Join(dir, "etags.json")

	c := NewETagCache(path)
	c.Set("https://example.com", `"etag"`)

	if err := c.Save(); err != nil {
		t.Fatalf("Save: %v", err)
	}
	if _, err := os.Stat(path); err != nil {
		t.Fatalf("cache file not created: %v", err)
	}
}
