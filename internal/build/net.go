package build

import (
	"context"
	"errors"
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

	// DialTimeout bounds the TCP/UDP dial performed by CanDial and queryWhois.
	// HTTP fetches use httpClient.Timeout (see below), not this value.
	DialTimeout = 10 * time.Second
)

// httpClient is the shared HTTP client for all build-package fetches.
// Timeout caps the whole request (DNS → body); per-call ctx deadlines compose on top.
var httpClient = &http.Client{
	Timeout: time.Minute,
}

// Fetch fetches a URL.
func Fetch(ctx context.Context, u string) (*http.Response, error) {
	color.Fprintf(os.Stderr, "@{.}Fetching %s\n", u)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, fmt.Errorf("building request for %s: %w", u, err)
	}
	req.Header.Set("User-Agent", httpUserAgent)
	res, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("fetching %s: %w", u, err)
	}
	if res.StatusCode != http.StatusOK {
		res.Body.Close()
		return nil, fmt.Errorf("HTTP error for %s: %s", u, res.Status)
	}
	return res, nil
}

// FetchWithETag performs a conditional GET using a stored ETag.
// Returns (response, nil) on 200 OK (updates cache with new ETag).
// Returns (nil, nil) on 304 Not Modified (content unchanged).
// Returns (nil, error) on failure.
// If cache is nil, behaves like Fetch.
func FetchWithETag(ctx context.Context, u string, cache *ETagCache) (*http.Response, error) {
	color.Fprintf(os.Stderr, "@{.}Fetching %s\n", u)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, fmt.Errorf("building request for %s: %w", u, err)
	}
	req.Header.Set("User-Agent", httpUserAgent)
	if cache != nil {
		if etag := cache.Get(u); etag != "" {
			req.Header.Set("If-None-Match", etag)
		}
	}
	res, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("fetching %s: %w", u, err)
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

// dialer is reused by CanDial and queryWhois to avoid allocating a net.Dialer
// per call. It has no per-call state that varies with ctx.
var dialer = &net.Dialer{Timeout: DialTimeout}

// CanDial reports whether network/addr accepts a connection, discarding it.
// Results are cached process-wide; ctx-induced errors are not cached so a
// cancelled dial doesn't poison future callers.
// FIXME: scope the cache per-run instead of package-global.
func CanDial(ctx context.Context, network, addr string) error {
	k := network + " " + addr
	mtx.RLock()
	err, ok := dialCache[k]
	mtx.RUnlock()
	if ok {
		return err
	}
	c, err := dialer.DialContext(ctx, network, addr)
	// Do not cache ctx-induced errors: a future caller with a fresh ctx
	// should get a real dial attempt.
	if err != nil {
		if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
			return err
		}
		mtx.Lock()
		dialCache[k] = err
		mtx.Unlock()
		return err
	}
	c.Close()
	mtx.Lock()
	dialCache[k] = nil
	mtx.Unlock()
	return nil
}

var (
	dialCache = make(map[string]error)
	mtx       sync.RWMutex
)
