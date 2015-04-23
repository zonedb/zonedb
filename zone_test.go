package zonedb

import (
	"testing"
	"unsafe"
)

func TestSizeofZone(t *testing.T) {
	var z Zone
	t.Logf("sizeof Zone = %d", unsafe.Sizeof(z))
}

func TestTLDs(t *testing.T) {
	t.Logf("%d top-level domains (%s to %s)", len(TLDs), TLDs[0].Domain, TLDs[len(TLDs)-1].Domain)
}

func TestZone_DomainACE(t *testing.T) {
	examples := map[string]string{
		"com":     "com",
		"net":     "net",
		"org":     "org",
		"アマゾン":    "xn--cckwcxetd",
		"العليان": "xn--mgba7c0bbn0a",
	}
	for domain, expected := range examples {
		z := ZoneMap[domain]
		got := z.DomainACE()
		if got != expected {
			t.Errorf("Expected %s, got %s for %s", expected, got, domain)
		}
	}
}

func BenchmarkInitAllocs(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		initZones()
	}
}
