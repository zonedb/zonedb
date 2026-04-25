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

	// Timeout specifies the maximum duration to wait for network operations to complete.
	// It applies as a cap on top of any deadline already on the ctx passed to
	// Fetch, FetchWithETag, and CanDial: the earliest deadline wins.
	Timeout = 10 * time.Second
)

// httpClient is the shared HTTP client for all build-package fetches.
//
// Introduced alongside the context.Context refactor of Fetch / FetchWithETag:
// threading ctx into those functions is only meaningful if the underlying
// transport has default bounds, otherwise callers passing context.Background()
// would get the same unbounded behavior the ctx-aware API suggests it prevents.
// Timeout is applied to the TCP dial, TLS handshake, and the overall request-
// response cycle; per-call ctx deadlines compose on top via
// http.NewRequestWithContext (earliest deadline wins).
//
// KeepAlive matches the Domainr house convention in mustang
// (internal/maintenance/fetcher.go).
var httpClient = &http.Client{
	Timeout: Timeout,
	Transport: &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   Timeout,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		TLSHandshakeTimeout: 5 * time.Second,
	},
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

// CanDial verifies if possible to connect to a given network and address.
// Returns nil for successful dial or an error.
//
// Results are cached in a process-global map keyed by "network addr". Errors
// from context cancellation or deadline are NOT cached, so that a cancelled
// dial does not poison future attempts with a fresh ctx.
//
// TODO: the process-global cache should be per-run or per-test to avoid
// cross-run state leakage. Tracked as a follow-up to the ctx refactor.
func CanDial(ctx context.Context, network, addr string) error {
	k := network + " " + addr
	mtx.RLock()
	err, ok := dialCache[k]
	mtx.RUnlock()
	if ok {
		return err
	}
	d := &net.Dialer{Timeout: Timeout}
	c, err := d.DialContext(ctx, network, addr)
	// Do not cache ctx-induced errors: a future caller with a fresh ctx
	// should get a real dial attempt.
	if !errors.Is(err, context.Canceled) && !errors.Is(err, context.DeadlineExceeded) {
		mtx.Lock()
		dialCache[k] = err
		mtx.Unlock()
	}
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
