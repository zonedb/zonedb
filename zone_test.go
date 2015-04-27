package zonedb

import (
	"fmt"
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

func TestTags_And(t *testing.T) {
	var tags Tags
	tags = TagGeneric | TagCountry
	if tags.And(TagGeneric) != true {
		t.Errorf("tags.And(TagGeneric) != true")
	}
	if tags.And(TagCountry) != true {
		t.Errorf("tags.And(TagCountry) != true")
	}
	if tags.And(TagAdult) != false {
		t.Errorf("tags.And(TagAdult) != false")
	}
	if tags.And(0) != false {
		t.Errorf("tags.And(0) != false")
	}
}

func ExampleTags_And() {
	var z *Zone
	z = ZoneMap["bananarepublic"]
	fmt.Println(z.Tags.And(TagBrand | TagGeo))
	// Output: true
}

func TestZone_WhoisServer(t *testing.T) {
	data := map[string]string{
		"com":    ZoneMap["net"].WhoisServer(),
		"er":     "",
		"uk.com": ZoneMap["us.com"].WhoisServer(),
		"com.er": "",
		"ac.uk":  ZoneMap["gov.uk"].WhoisServer(),
		"co.uk":  ZoneMap["uk"].WhoisServer(),
		"org.uk": ZoneMap["uk"].WhoisServer(),
	}
	for k, v := range data {
		g := ZoneMap[k].WhoisServer()
		if g != v {
			t.Errorf(`Expected Zones[%q].WhoisServer() == %q, got %q`, k, v, g)
		}
	}
}

func TestZone_WhoisURL(t *testing.T) {
	data := map[string]string{
		"com":   "",
		"net":   "",
		"org":   "",
		"co.az": ZoneMap["az"].WhoisURL(),
	}
	for k, v := range data {
		g := ZoneMap[k].WhoisURL()
		if g != v {
			t.Errorf(`Expected Zones[%q].WhoisURL() == %q, got %q`, k, v, g)
		}
	}
}

func TestZone_IsTLD(t *testing.T) {
	data := map[string]bool{
		"com":    true,
		"um":     true,
		"co.uk":  false,
		"org.br": false,
	}
	for k, v := range data {
		g := ZoneMap[k].IsTLD()
		if g != v {
			t.Errorf(`Expected Zones[%q].IsTLD() == %t, got %t`, k, v, g)
		}
	}
}

func TestZone_IsDelegated(t *testing.T) {
	data := map[string]bool{
		"com":    true,
		"um":     false,
		"yu":     false,
		"co.uk":  true,
		"org.za": true,
		"db.za":  false,
	}
	for k, v := range data {
		g := ZoneMap[k].IsDelegated()
		if g != v {
			t.Errorf(`Expected Zones[%q].IsDelegated() == %t, got %t`, k, v, g)
		}
	}
}

func TestZone_IsInRootZone(t *testing.T) {
	data := map[string]bool{
		"com":    true,
		"net":    true,
		"org":    true,
		"um":     false,
		"yu":     false,
		"co.uk":  false,
		"org.br": false,
	}
	for k, v := range data {
		g := ZoneMap[k].IsInRootZone()
		if g != v {
			t.Errorf(`Expected Zones[%q].IsInRootZone() == %t, got %t`, k, v, g)
		}
	}
}

func TestZone_AllowsRegistration(t *testing.T) {
	data := map[string]bool{
		"com":          true,
		"net":          true,
		"org":          true,
		"ck":           false,
		"yu":           false,
		"arpa":         false,
		"cadillac":     false,
		"amazon":       false,
		"co.uk":        true,
		"in-addr.arpa": false,
	}
	for k, v := range data {
		g := ZoneMap[k].AllowsRegistration()
		if g != v {
			t.Errorf(`Expected Zones[%q].AllowsRegistration() == %t, got %t`, k, v, g)
		}
	}
}

func BenchmarkInitAllocs(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		initZones()
	}
}
