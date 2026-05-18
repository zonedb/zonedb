package build

import (
	"context"
	"net"
	"testing"
	"time"
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

	// Ensure the cache starts clean for this address. Other tests or
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

// TestCanDialDoesNotCacheTransientErrors: a single stub-resolver flake
// must not poison the cache for later callers of the same host.
func TestCanDialDoesNotCacheTransientErrors(t *testing.T) {
	// dnsTimeout reproduces the 2026-05-13 a-dns.pl failure; dnsTemporary
	// covers non-timeout transient resolver failures (e.g. SERVFAIL).
	dnsTimeout := &net.OpError{
		Op:  "dial",
		Net: "udp",
		Err: &net.DNSError{
			Err:       "i/o timeout",
			Name:      "transient-flake.example.invalid",
			Server:    "127.0.0.53:53",
			IsTimeout: true,
		},
	}
	dnsTemporary := &net.DNSError{
		Err:         "server misbehaving",
		Name:        "temp-flake.example.invalid",
		Server:      "127.0.0.53:53",
		IsTemporary: true,
	}

	cases := []struct {
		name string
		addr string
		err  error
	}{
		{"dns_timeout_wrapped_in_op_error", "transient-flake.example.invalid:53", dnsTimeout},
		{"dns_temporary", "temp-flake.example.invalid:53", dnsTemporary},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			// Isolate from prior runs and other tests that may share dialCache.
			key := "udp " + tc.addr
			mtx.Lock()
			delete(dialCache, key)
			mtx.Unlock()
			t.Cleanup(func() {
				mtx.Lock()
				delete(dialCache, key)
				mtx.Unlock()
			})

			// Substitute the dialer to force the chosen transient error and
			// count invocations. Restored via t.Cleanup.
			var calls int
			orig := dialContext
			t.Cleanup(func() { dialContext = orig })
			dialContext = func(_ context.Context, _, _ string) (net.Conn, error) {
				calls++
				return nil, tc.err
			}

			// First call: must return the transient error to the caller.
			if err := CanDial(t.Context(), "udp", tc.addr); err == nil {
				t.Fatal("expected transient error from first call, got nil")
			}

			// The contract under test: the transient error must not be cached.
			mtx.RLock()
			_, cached := dialCache[key]
			mtx.RUnlock()
			if cached {
				t.Fatal("transient error was cached; cache would poison future callers (the May 13 .pl flap)")
			}

			// Second call: must re-dial (calls == 2), not short-circuit on cache.
			if err := CanDial(t.Context(), "udp", tc.addr); err == nil {
				t.Fatal("expected transient error from second call, got nil")
			}
			if calls != 2 {
				t.Fatalf("dialer was called %d times; want 2 (cache must not short-circuit transient failures)", calls)
			}
		})
	}
}

// TestCanDialCachesDurableErrors: connection-refused stays cached, so
// known-bad endpoints aren't re-dialled within a build run.
func TestCanDialCachesDurableErrors(t *testing.T) {
	// Bind then immediately close to obtain an address that the OS
	// will refuse connections on — a deterministic durable error.
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	addr := ln.Addr().String()
	if err := ln.Close(); err != nil {
		t.Fatal(err)
	}

	// Isolate from prior runs and other tests that may share dialCache.
	key := "tcp " + addr
	mtx.Lock()
	delete(dialCache, key)
	mtx.Unlock()
	t.Cleanup(func() {
		mtx.Lock()
		delete(dialCache, key)
		mtx.Unlock()
	})

	// Act: a single dial attempt that should hit the closed port.
	if err := CanDial(t.Context(), "tcp", addr); err == nil {
		t.Fatal("expected connection-refused error, got nil")
	}

	// The contract under test: durable errors must be cached.
	mtx.RLock()
	_, cached := dialCache[key]
	mtx.RUnlock()
	if !cached {
		t.Fatal("durable error was not cached; we should not re-dial known-bad endpoints")
	}
}

// TestVerifyDialTimeout pins the policy that verify probes use a
// tighter dial budget than DialTimeout.
func TestVerifyDialTimeout(t *testing.T) {
	if VerifyDialTimeout >= DialTimeout {
		t.Fatalf("VerifyDialTimeout (%v) must be tighter than DialTimeout (%v)", VerifyDialTimeout, DialTimeout)
	}
	if VerifyDialTimeout <= 0 {
		t.Fatalf("VerifyDialTimeout must be positive, got %v", VerifyDialTimeout)
	}
	if VerifyDialTimeout > 5*time.Second {
		t.Fatalf("VerifyDialTimeout (%v) is generous for a UDP local-socket probe; expected ≤5s", VerifyDialTimeout)
	}
}
