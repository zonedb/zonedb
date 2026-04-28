package build

import (
	"context"
	"net"
	"testing"
)

// TestCanDialCacheSkipsCtxErrors: a cancelled dial in one build phase must
// not poison the shared dialCache for the next phase's fresh ctx. Touches
// package-global dialCache/mtx; revisit when the cache is scoped per-run.
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
