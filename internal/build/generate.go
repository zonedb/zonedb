package build

import (
	"bytes"
	"fmt"
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

	data := templateData{
		Zones:     zones,
		TLDs:      tlds,
		Domains:   domains,
		Offsets:   offsets,
		TagType:   tagType,
		Tags:      tags,
		TagValues: tagValues,
		Strings:   []string{},
	}

	// Pre-index strings and string slices
	for _, d := range domains {
		data.domainString(d)
	}
	for _, d := range domains {
		data.domainStringSlice(zones[d].NameServers)
	}
	for _, d := range domains {
		data.domainStringSlice(zones[d].Wildcards)
	}
	for _, d := range domains {
		data.domainString(zones[d].WhoisServer)
	}
	for _, d := range domains {
		data.urlString(zones[d].WhoisURL)
	}
	for _, d := range domains {
		data.urlString(zones[d].InfoURL)
	}
	for _, d := range domains {
		data.indexedStringSlice(zones[d].Locations)
	}
	for _, d := range domains {
		data.indexedStringSlice(zones[d].Languages)
	}

	err := generate("zones.go", zonesGoSrc, &data)
	if err != nil {
		return err
	}

	return nil
}

type templateData struct {
	Zones     map[string]*Zone
	TLDs      map[string]*Zone
	Domains   []string
	Offsets   map[string]int
	TagType   string
	Tags      []string
	TagValues map[string]uint64
	Strings   []string
}

func (data *templateData) indexedString(s string) string {
	i, _ := IndexOrAppendStrings(&data.Strings, []string{s})
	return fmt.Sprintf("s[%d]", i)
}

func (data *templateData) indexedStringSlice(slice []string) string {
	i, j := IndexOrAppendStrings(&data.Strings, slice)
	return fmt.Sprintf("s[%d:%d]", i, j)
}

func (data *templateData) domainString(s string) string {
	return data.indexedString(ToASCII(s))
}

func (data *templateData) domainStringSlice(slice []string) string {
	needle := make([]string, len(slice))
	for i := range slice {
		needle[i] = ToASCII(slice[i])
	}
	return data.indexedStringSlice(needle)
}

func (data *templateData) urlString(s string) string {
	return data.indexedString(ToASCIIURL(s))
}

func (data *templateData) urlStringSlice(slice []string) string {
	needle := make([]string, len(slice))
	for i := range slice {
		needle[i] = ToASCIIURL(slice[i])
	}
	return data.indexedStringSlice(needle)
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

func generate(filename, src string, data *templateData) error {
	funcMap := template.FuncMap{
		"title":        strings.Title,
		"quoted":       quoted,
		"quotedDomain": quotedDomain,
		"quotedURL":    quotedURL,
		"string":       data.indexedString,
		"stringSlice":  data.indexedStringSlice,
		"domain":       data.domainString,
		"domainSlice":  data.domainStringSlice,
		"url":          data.urlString,
		"urlSlice":     data.urlStringSlice,
	}

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
	filename = filepath.Join(BaseDir, filename)
	color.Fprintf(os.Stderr, "@{.}Generating Go source code: %s\n", filename)
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(formatted)
	return err
}

func cont(s string) string {
	return strings.ReplaceAll(s, "\\\n", "")
}

const zonesGoSrc = `// Automatically generated

package zonedb

func init() {
	initZones() // Separate function to report allocs in initialization
}

func initZones() {
	z = y
	ZoneMap = make(map[string]*Zone, len(z))
	for i := range z {
		ZoneMap[z[i].Domain] = &z[i]
	}
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

// ZoneMap maps domain names to Zones.
var ZoneMap map[string]*Zone

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

// s is an array of unique string slices used in zone data.
var s = [{{len .Strings}}]string{
	{{range .Strings}}{{quoted .}},
	{{end}}
}
`
