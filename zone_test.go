package zonedb

import (
	"testing"
	"unsafe"
)

func TestZone(t *testing.T) {
	var z Zone
	t.Logf("sizeof Zone = %d", unsafe.Sizeof(z))
}

func TestZones(t *testing.T) {
	t.Logf("len(zones) = %d", len(Zones))
	t.Logf("%+v\n", Zones[1442])
	t.Logf("len(uk Subdomains)%d\n", len(Zones[1442].Subdomains))
	t.Logf("cap(uk Subdomains)%d\n", cap(Zones[1442].Subdomains))
}

func TestTLDS(t *testing.T) {
	t.Logf("%d top-level domains (%s to %s)", len(TLDs), TLDs[0].Domain, TLDs[len(TLDs)-1].Domain)
}

func BenchmarkInitAllocs(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		initZones()
	}
}
