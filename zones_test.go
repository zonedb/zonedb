package zonedb

import (
	"fmt"
	"testing"
)

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

func TestTags_String(t *testing.T) {
	tests := map[Tags]string{
		(TagGeneric):                            "generic",
		(TagWithdrawn):                          "withdrawn",
		(TagCountry | TagGeo):                   "country geo",
		(TagCountry | TagGeo | TagSponsored):    "country geo sponsored",
		(TagAdult | TagGeo | TagInfrastructure): "adult geo infrastructure",
	}
	for tags, v := range tests {
		g := tags.String()
		if g != v {
			t.Errorf(`Expected tags.String() == %q, got %q`, v, g)
		}
	}
}

func ExampleTags_String() {
	var z *Zone
	z = ZoneMap["aero"]
	fmt.Println(z.Tags.String())
	// Output: generic sponsored
}

func TestZone_WhoisServer(t *testing.T) {
	tests := map[string]string{
		"com":    ZoneMap["net"].WhoisServer(),
		"er":     "",
		"uk.com": ZoneMap["us.com"].WhoisServer(),
		"com.er": "",
		"ac.uk":  ZoneMap["gov.uk"].WhoisServer(),
		"co.uk":  ZoneMap["uk"].WhoisServer(),
		"org.uk": ZoneMap["uk"].WhoisServer(),
	}
	for k, v := range tests {
		g := ZoneMap[k].WhoisServer()
		if g != v {
			t.Errorf(`Expected Zones[%q].WhoisServer() == %q, got %q`, k, v, g)
		}
	}
}

func TestZone_WhoisURL(t *testing.T) {
	tests := map[string]string{
		"com":   "",
		"net":   "",
		"org":   "",
		"co.az": ZoneMap["az"].WhoisURL(),
	}
	for k, v := range tests {
		g := ZoneMap[k].WhoisURL()
		if g != v {
			t.Errorf(`Expected Zones[%q].WhoisURL() == %q, got %q`, k, v, g)
		}
	}
}

func TestZonesHaveWhoisServerOrURLButNotBoth(t *testing.T) {
	for _, z := range Zones {
		if z.whoisServer != "" && z.whoisURL != "" {
			t.Errorf(`Zone %q has both whoisServer (%q) and whoisURL (%q) set`, z.Domain, z.whoisServer, z.whoisURL)
		}
	}
}

func TestZone_IsTLD(t *testing.T) {
	tests := map[string]bool{
		"com":    true,
		"um":     true,
		"co.uk":  false,
		"org.br": false,
	}
	for k, v := range tests {
		g := ZoneMap[k].IsTLD()
		if g != v {
			t.Errorf(`Expected Zones[%q].IsTLD() == %t, got %t`, k, v, g)
		}
	}
}

func TestZone_IsDelegated(t *testing.T) {
	tests := map[string]bool{
		"com":    true,
		"um":     false,
		"yu":     false,
		"co.uk":  true,
		"org.za": true,
		"db.za":  false,
	}
	for k, v := range tests {
		g := ZoneMap[k].IsDelegated()
		if g != v {
			t.Errorf(`Expected Zones[%q].IsDelegated() == %t, got %t`, k, v, g)
		}
	}
}

func TestZone_IsInRootZone(t *testing.T) {
	tests := map[string]bool{
		"com":    true,
		"net":    true,
		"org":    true,
		"um":     false,
		"yu":     false,
		"co.uk":  false,
		"org.br": false,
	}
	for k, v := range tests {
		g := ZoneMap[k].IsInRootZone()
		if g != v {
			t.Errorf(`Expected Zones[%q].IsInRootZone() == %t, got %t`, k, v, g)
		}
	}
}

func TestZone_AllowsIDN(t *testing.T) {
	data := map[string]bool{
		"com":   true,
		"net":   true,
		"org":   true,
		"aero":  false,
		"co.uk": false,
	}
	for k, v := range data {
		g := ZoneMap[k].AllowsIDN()
		if g != v {
			t.Errorf(`Expected Zones[%q].AllowsIDN() == %t, got %t`, k, v, g)
		}
	}
}

func TestZone_AllowsRegistration(t *testing.T) {
	tests := map[string]bool{
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
	for k, v := range tests {
		g := ZoneMap[k].AllowsRegistration()
		if g != v {
			t.Errorf(`Expected Zones[%q].AllowsRegistration() == %t, got %t`, k, v, g)
		}
	}
}

func TestIsZone(t *testing.T) {
	tests := map[string]bool{
		"com":                true,
		"um":                 true,
		"xn--node":           true,
		"co.uk":              true,
		"org.br":             true,
		"hashtag-not-a-zone": false,
	}
	for k, v := range tests {
		g := IsZone(k)
		if g != v {
			t.Errorf(`Expected IsZone(%q) == %t, got %t`, k, v, g)
		}
	}
}

func TestIsTLD(t *testing.T) {
	tests := map[string]bool{
		"com":                true,
		"um":                 true,
		"xn--node":           true,
		"co.uk":              false,
		"org.br":             false,
		"hashtag-not-a-zone": false,
	}
	for k, v := range tests {
		g := IsTLD(k)
		if g != v {
			t.Errorf(`Expected IsTLD(%q) == %t, got %t`, k, v, g)
		}
	}
}

func TestPublicZone(t *testing.T) {
	tests := map[string]*Zone{
		"com":           ZoneMap["com"],
		".com":          ZoneMap["com"],
		"foobar.com":    ZoneMap["com"],
		"acme.co.uk":    ZoneMap["co.uk"],
		"brazil.com.br": ZoneMap["com.br"],
		"foo.xn--node":  ZoneMap["xn--node"],
		"unknown.":      nil,
		"COM":           nil,
		"bar.გე":        nil,
	}
	for k, v := range tests {
		g := PublicZone(k)
		if g != v {
			t.Errorf(`Expected List.PublicSuffix(%q) == %v, got %v`, k, v, g)
		}
	}
}

func TestList_PublicSuffix(t *testing.T) {
	tests := map[string]string{
		"com":           "com",
		".com":          "com",
		"foobar.com":    "com",
		"foo.xn--node":  "xn--node",
		"acme.co.uk":    "co.uk",
		"brazil.com.br": "com.br",
		"unknown.":      "unknown.",
		"COM":           "COM",
		"bar.გე":        "bar.გე",
	}
	for k, v := range tests {
		g := List.PublicSuffix(k)
		if g != v {
			t.Errorf(`Expected List.PublicSuffix(%q) == %q, got %q`, k, v, g)
		}
	}
}

func BenchmarkInitAllocs(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		initZones()
	}
}
