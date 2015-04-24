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

func TestTags(t *testing.T) {
	if numTags != len(TagStrings) {
		t.Errorf("numTags (%d) != len(TagStrings) (%d)", numTags, len(TagStrings))
	}
	if len(TagStrings) != len(TagValues) {
		t.Errorf("len(TagStrings) (%d) != len(TagValues) (%d)", len(TagStrings), len(TagValues))
	}
}

func BenchmarkInitAllocs(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		initZones()
	}
}
