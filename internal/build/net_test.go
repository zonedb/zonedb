package build

import (
	"context"
	"net"
	"testing"
)

// TestCanDialCacheSkipsCtxErrors verifies that CanDial does not cache errors
// caused by context cancellation or deadline. A cancelled dial must not
// poison the cache for subsequent callers passing a fresh ctx — otherwise
// the first Ctrl-C during a build would cause every future build to
// instant-fail for the cached hosts.
//
// NOTE: this test directly touches the package-global dialCache / mtx.
// When the follow-up to scope the cache per-run lands, update this test
// accordingly (see net.go's TODO above CanDial).
func TestCanDialCacheSkipsCtxErrors(t *testing.T) {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = ln.Close() })
	addr := ln.Addr().String()

	// Ensure the cache starts clean for this address — other tests or
	// prior runs in the same process could have seeded it.
	key := "tcp " + addr
	mtx.Lock()
	delete(dialCache, key)
	mtx.Unlock()
	t.Cleanup(func() {
		mtx.Lock()
		delete(dialCache, key)
		mtx.Unlock()
	})

	// First call with an already-cancelled ctx should error but NOT cache.
	ctx, cancel := context.WithCancel(t.Context())
	cancel()
	if err := CanDial(ctx, "tcp", addr); err == nil {
		t.Fatal("expected error from cancelled CanDial, got nil")
	}

	mtx.RLock()
	_, cached := dialCache[key]
	mtx.RUnlock()
	if cached {
		t.Fatal("context.Canceled was cached; cache would poison future callers")
	}

	// A fresh ctx should successfully dial, proving the cache wasn't poisoned.
	if err := CanDial(t.Context(), "tcp", addr); err != nil {
		t.Fatalf("fresh ctx failed to dial: %v", err)
	}
}
