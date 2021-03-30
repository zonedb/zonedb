package build

import (
	"bytes"
	"go/format"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/template"

	"github.com/wsxiaoys/terminal/color"
)

// GenerateGo generates a Go source representation of ZoneDB.
func GenerateGo(zones map[string]*Zone) error {
	tlds := TLDs(zones)
	domains := SortedDomains(zones)
	offsets := make(map[string]int, len(domains))

	tagSet := NewSet()
	for i, d := range domains {
		offsets[d] = i
		z := zones[d]
		tagSet.Add(z.Tags...)
	}
	tags := tagSet.Values()
	sort.Strings(tags)
	tagValues := make(map[string]uint64)
	for i, t := range tags {
		tagValues[t] = 1 << uint64(i)
	}
	tagType := "uint32"
	if len(tags) > 32 {
		tagType = "uint64"
	}

	for _, d := range domains {
		z := zones[d]
		z.Normalize() // just in case
		z.ParentOffset = offsets[z.ParentDomain()]
		if len(z.subdomains) > 0 {
			z.SubdomainsOffset = offsets[z.subdomains[0]]
			z.SubdomainsEnd = z.SubdomainsOffset + len(z.subdomains)
		}
		z.TagBits = tagBits(tagValues, z.Tags)
	}

	data := struct {
		Zones     map[string]*Zone
		TLDs      map[string]*Zone
		Domains   []string
		Offsets   map[string]int
		TagType   string
		Tags      []string
		TagValues map[string]uint64
	}{
		zones,
		tlds,
		domains,
		offsets,
		tagType,
		tags,
		tagValues,
	}

	err := generate("zones.go", zonesGoSrc, &data)
	if err != nil {
		return err
	}

	return nil
}

// Helper funcs

func cont(s string) string {
	return strings.ReplaceAll(s, "\\\n", "")
}

func quoted(s string) string {
	if s == "" {
		return "e" // const e = ""
	}
	return `"` + s + `"`
}

func quotedDomain(s string) string {
	return quoted(ToASCII(s))
}

func quotedURL(s string) string {
	return quoted(ToASCIIURL(s))
}

var (
	funcMap = template.FuncMap{
		"title":        strings.Title,
		"quoted":       quoted,
		"quotedDomain": quotedDomain,
		"quotedURL":    quotedURL,
	}
)

func generate(fn, src string, data interface{}) error {
	t := template.Must(template.New("").Funcs(funcMap).Parse(cont(src)))
	buf := new(bytes.Buffer)
	err := t.Execute(buf, data)
	if err != nil {
		return err
	}
	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		return err
	}
	fn = filepath.Join(BaseDir, fn)
	color.Fprintf(os.Stderr, "@{.}Generating Go source code: %s\n", fn)
	f, err := os.Create(fn)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(formatted)
	return err
}

const (
	zonesGoSrc = `// Automatically generated

package zonedb

func init() {
	initZones() // Separate function to report allocs in initialization
}

func initZones() {
	z = y
}

// Type s is an alias for []string to generate smaller source code
type s []string

// Constants to generate smaller source code
const (
	t = true
	f = false
	e = ""
)

// Nil variables to generate smaller source code
var (
	n []string // == nil
	x []Zone // nil (no subdomains)
	r *Zone // nil (root zone for TLDs)
)

// Tags are stored in a single integer as a bit field.
type Tags {{.TagType}}

// Tag values corresponding to bit shifts (1 << iota)
const (
	{{range $i, $t := .Tags}} \
		Tag{{title $t}} = {{index $.TagValues $t}}
	{{end}} \
	numTags = {{len .Tags}}
)

// TagStrings maps integer tag values to strings.
var TagStrings = map[Tags]string{
	{{range $t := .Tags}} \
		Tag{{title $t}}: "{{$t}}",
	{{end }}
}

// TagValues maps tag names to integer values.
var TagValues  = map[string]Tags{
	{{range $t := .Tags}} \
		"{{$t}}": Tag{{title $t}},
	{{end }}
}

// Zones is a slice of all Zones in the database.
var Zones = z[:]

// TLDs is a slice of all top-level domain Zones.
var TLDs = z[:{{len .TLDs}}]

// z is a static array of Zones.
// Other global variables have pointers into this array.
var z [{{len .Zones}}]Zone

// y and z are separated to fix circular references.
var y = [{{len .Zones}}]Zone{
	{{range $d := .Domains}} \
		{{$z := (index $.Zones $d)}} \
		{ \
			{{quotedDomain $d}}, \
			{{if $z.IsIDN}}/* {{$d}} */{{end }} \
			{{if $z.ParentDomain}} &z[{{$z.ParentOffset}}] {{else}} r {{end}}, \
			{{if $z.SubdomainsEnd}} z[{{$z.SubdomainsOffset}}:{{$z.SubdomainsEnd}}] {{else}} x {{end}}, \
			{{if $z.NameServers}} s{ {{range $z.NameServers}}{{quotedDomain .}},{{end}}} {{else}} n {{end}}, \
			{{if $z.Wildcards}} s{ {{range $z.Wildcards}}{{quoted .}},{{end}}} {{else}} n {{end}}, \
			{{if $z.Locations}} s{ {{range $z.Locations}}{{quoted .}},{{end}}} {{else}} n {{end}}, \
			{{if $z.Languages}} s{ {{range $z.Languages}}{{quoted .}},{{end}}} {{else}} n {{end}}, \
			{{quotedDomain $z.WhoisServer}}, \
			{{quotedURL $z.WhoisURL}}, \
			{{quotedURL $z.InfoURL}}, \
			{{printf "0x%x" $z.TagBits}}, \
			{{if $z.IDNDisallowed}} f {{else}} t {{end}}, \
		},
	{{end}} \
}

// ZoneMap maps domain names to Zones.
var ZoneMap = map[string]*Zone{
	{{range $d := .Domains}}  \
		{{$z := (index $.Zones $d)}} \
		{{$o := index $.Offsets $d}} \
		{{quotedDomain $d}}: &z[{{$o}}], \
		{{if $z.IsIDN}}// {{$d}}{{end }}
	{{end}} \
}
`
)
