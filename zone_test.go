package zonedb

import (
	"errors"
	"fmt"
	"testing"
	"unsafe"

	"golang.org/x/net/idna"
)

// ToASCII normalizes a domain name or URL to ASCII/punycode.
func ToASCII(s string) string {
	s, _ = idna.ToASCII(s)
	return s
}

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

func TestTags_String(t *testing.T) {
	data := map[Tags]string{
		(TagGeneric):                            "generic",
		(TagWithdrawn):                          "withdrawn",
		(TagCountry | TagGeo):                   "country geo",
		(TagCountry | TagGeo | TagSponsored):    "country geo sponsored",
		(TagAdult | TagGeo | TagInfrastructure): "adult geo infrastructure",
	}
	for tags, v := range data {
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

type idnTest struct {
	Zone        string
	Domain      string
	Valid       bool
	ValidTables []string
	Error       error
}

func TestZone_IsValidDomain(t *testing.T) {
	data := []idnTest{
		{Zone: "jobs", Domain: "test.jobs", Valid: true},
		{Zone: "jobs", Domain: "tést.jobs", Valid: false},
		{Zone: "jobs", Domain: "test.com", Valid: false},
		{Zone: "com", Domain: "test.com", Valid: true},
		{Zone: "com", Domain: "tést.com", Valid: true},
		{Zone: ToASCII("рф"), Domain: ToASCII("russia.рф"), Valid: false},
		{Zone: ToASCII("рф"), Domain: ToASCII("россия.рф"), Valid: true},
	}
	for _, d := range data {
		res := ZoneMap[d.Zone].IsValidDomain(d.Domain)
		if res != d.Valid {
			t.Errorf(`Expected Zones[%q].IsValidDomain(%q) == %t, got %t`, d.Zone, d.Domain, d.Valid, res)
		}
	}
}

func TestZone_IDNTable(t *testing.T) {
	data := []idnTest{
		{Zone: "jobs", Domain: "test.jobs", ValidTables: []string{""}},
		{Zone: "jobs", Domain: "tést.jobs", Error: errors.New("domain is not a valid member of the zone")},
		{Zone: "jobs", Domain: "test.com", Error: errors.New("domain is not a member of the zone")},
		{Zone: "com", Domain: "test.com", ValidTables: []string{""}},
		{Zone: "bg", Domain: "здравей.bg", ValidTables: []string{"bg", "bg-bg", "ru", "ru-bg"}},
		{Zone: "bg", Domain: "здравей.всё.bg", ValidTables: []string{"ru", "ru-bg"}},
		{Zone: "bg", Domain: "tést.bg", Error: errors.New("domain is not a valid member of the zone")},
	}
	for _, d := range data {
		table, err := ZoneMap[d.Zone].IDNTable(d.Domain)
		if err != nil && d.Error == nil {
			t.Errorf(`Expected Zones[%q].IDNTable(%q) to not error, got %q`, d.Zone, d.Domain, err.Error())
		} else if err == nil && d.Error != nil {
			t.Errorf(`Expected Zones[%q].IDNTable(%q) to error with %q, got nothing`, d.Zone, d.Domain, d.Error.Error())
		} else if err != nil && d.Error != nil && err.Error() != d.Error.Error() {
			t.Errorf(`Expected Zones[%q].IDNTable(%q) to error with %q, got %q`, d.Zone, d.Domain, d.Error.Error(), err.Error())
		}
		if len(d.ValidTables) > 0 {
			found := false
			for _, validTable := range d.ValidTables {
				if table == validTable {
					found = true
					break
				}
			}
			if !found {
				t.Errorf(`Expected Zones[%q].IDNTable(%q) table in %q, got %q`, d.Zone, d.Domain, d.ValidTables, table)
			}
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

func TestIsZone(t *testing.T) {
	data := map[string]bool{
		"com":                true,
		"um":                 true,
		"xn--node":           true,
		"co.uk":              true,
		"org.br":             true,
		"hashtag-not-a-zone": false,
	}
	for k, v := range data {
		g := IsZone(k)
		if g != v {
			t.Errorf(`Expected IsZone(%q) == %t, got %t`, k, v, g)
		}
	}
}

func TestIsTLD(t *testing.T) {
	data := map[string]bool{
		"com":                true,
		"um":                 true,
		"xn--node":           true,
		"co.uk":              false,
		"org.br":             false,
		"hashtag-not-a-zone": false,
	}
	for k, v := range data {
		g := IsTLD(k)
		if g != v {
			t.Errorf(`Expected IsTLD(%q) == %t, got %t`, k, v, g)
		}
	}
}

func TestPublicZone(t *testing.T) {
	data := map[string]*Zone{
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
	for k, v := range data {
		g := PublicZone(k)
		if g != v {
			t.Errorf(`Expected List.PublicSuffix(%q) == %v, got %v`, k, v, g)
		}
	}
}

func TestList_PublicSuffix(t *testing.T) {
	data := map[string]string{
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
	for k, v := range data {
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
