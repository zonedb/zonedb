package build

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/wsxiaoys/terminal/color"
)

const httpUserAgent = "zonedb/1.0 (https://github.com/zonedb/zonedb)"

var (
	// Concurrency specifies the maximimum number of concurrent build operations to execute.
	Concurrency = 32

	// Timeout specifies the maximum duration to wait for network operations to complete.
	Timeout = 10 * time.Second
)

// Fetch fetches a URL.
func Fetch(u string) (*http.Response, error) {
	color.Fprintf(os.Stderr, "@{.}Fetching %s\n", u)
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", httpUserAgent)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP error for %s: %s", u, res.Status)
	}
	return res, nil
}

// FetchWithETag performs a conditional GET using a stored ETag.
// Returns (response, nil) on 200 OK (updates cache with new ETag).
// Returns (nil, nil) on 304 Not Modified (content unchanged).
// Returns (nil, error) on failure.
// If cache is nil, behaves like Fetch.
func FetchWithETag(u string, cache *ETagCache) (*http.Response, error) {
	color.Fprintf(os.Stderr, "@{.}Fetching %s\n", u)
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", httpUserAgent)
	if cache != nil {
		if etag := cache.Get(u); etag != "" {
			req.Header.Set("If-None-Match", etag)
		}
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode == http.StatusNotModified {
		res.Body.Close()
		return nil, nil
	}
	if res.StatusCode != http.StatusOK {
		res.Body.Close()
		return nil, fmt.Errorf("HTTP error for %s: %s", u, res.Status)
	}
	if cache != nil {
		if etag := res.Header.Get("ETag"); etag != "" {
			cache.Set(u, etag)
		}
	}
	return res, nil
}

// CanDial verifies if possible to connect to a given network and address.
// Returns nil for successful dial or an error.
func CanDial(network, addr string) error {
	k := network + " " + addr
	mtx.RLock()
	err, ok := dialCache[k]
	mtx.RUnlock()
	if ok {
		return err
	}
	c, err := net.DialTimeout(network, addr, Timeout)
	mtx.Lock()
	dialCache[k] = err
	mtx.Unlock()
	if err != nil {
		return err
	}
	c.Close()
	return nil
}

var (
	dialCache = make(map[string]error)
	mtx       sync.RWMutex
)
