package build

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// ETagCache stores URL→ETag mappings for conditional HTTP requests.
type ETagCache struct {
	path    string
	entries map[string]string
}

// NewETagCache creates an ETagCache that persists to the given file path.
func NewETagCache(path string) *ETagCache {
	return &ETagCache{
		path:    path,
		entries: make(map[string]string),
	}
}

// Load reads cached ETags from the JSON file. Returns nil if the file
// does not exist (first run).
func (c *ETagCache) Load() error {
	data, err := os.ReadFile(c.path)
	if os.IsNotExist(err) {
		return nil
	}
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &c.entries)
}

// Save writes cached ETags to the JSON file, creating directories as needed.
func (c *ETagCache) Save() error {
	if err := os.MkdirAll(filepath.Dir(c.path), 0o755); err != nil {
		return err
	}
	data, err := json.MarshalIndent(c.entries, "", "\t")
	if err != nil {
		return err
	}
	return os.WriteFile(c.path, data, 0o644)
}

// Get returns the stored ETag for a URL, or "" if not cached.
func (c *ETagCache) Get(url string) string {
	return c.entries[url]
}

// Set stores an ETag for a URL.
func (c *ETagCache) Set(url, etag string) {
	c.entries[url] = etag
}
