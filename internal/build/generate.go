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
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
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

	// Pre-index 2-letter pairs
	var pairs string
	const alphabet = "abcdefghijklmnopqrstuvwxyz"
	for _, i := range alphabet {
		for _, j := range alphabet {
			s := string(i) + string(j)
			if strings.Contains(pairs, s) {
				continue
			}
			if len(pairs) > 0 && pairs[len(pairs)-1] == s[0] {
				pairs = pairs + s[1:]
			} else {
				pairs = pairs + s
			}
		}
	}
	data.indexedString(pairs)

	// Pre-index strings and string slices
	set := NewSet()
	for _, z := range zones {
		set.Add(ToASCII(z.WhoisServer))
		set.Add(ToASCIIURL(z.WhoisURL))
		set.Add(ToASCIIURL(z.InfoURL))
	}

	// Sort by length desc, then alphabetically asc
	all := set.Values()
	sort.Slice(all, func(i, j int) bool {
		if len(all[i]) == len(all[j]) {
			return all[i] < all[j]
		}
		return len(all[i]) > len(all[j])
	})
	for _, s := range all {
		data.indexedString(s)
	}

	for _, d := range domains {
		data.domainStringSlice(zones[d].NameServers)
	}
	for _, d := range domains {
		data.domainStringSlice(zones[d].Wildcards)
	}
	for _, d := range domains {
		data.indexedStringSlice(zones[d].Locations)
	}
	for _, d := range domains {
		data.indexedStringSlice(zones[d].Languages)
	}

	// Add domains last, as these are most likely to be the shortest strings
	for i := range domains {
		data.domainString(domains[len(domains)-i-1])
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
	if s == "" {
		return "e"
	}
	i, j, k := IndexOrAppendString(&data.Strings, s)
	if j == 0 && k == len(data.Strings[i]) {
		return fmt.Sprintf("s[%d]", i)
	} else if j == 0 {
		return fmt.Sprintf("s[%d][:%d]", i, k)
	}
	return fmt.Sprintf("s[%d][%d:%d]", i, j, k)
}

func (data *templateData) indexedStringSlice(slice []string) string {
	if len(slice) == 0 {
		return "n"
	}
	// if len(slice) == 1 {
	// 	return "[]string{" + data.indexedString(slice[0]) + "}"
	// }
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

// identifier converts a kebab-case string to a Go identifier string.
func identifier(s string) string {
	s = tagCaser.String(s)
	return strings.ReplaceAll(s, "-", "")
}

var tagCaser = cases.Title(language.English)

func quoted(s string) string {
	if s == "" {
		return "e" // const e = ""
	}
	return fmt.Sprintf("%q", s)
}

func quotedDomain(s string) string {
	return quoted(ToASCII(s))
}

func quotedURL(s string) string {
	return quoted(ToASCIIURL(s))
}

func mod(i, j, k int) bool {
	return i%j == k
}

func generate(filename, src string, data *templateData) error {
	funcMap := template.FuncMap{
		"mod":          mod,
		"identifier":   identifier,
		"quoted":       quoted,
		"quotedDomain": quotedDomain,
		"quotedURL":    quotedURL,
		"string":       data.indexedString,
		"stringSlice":  data.indexedStringSlice,
		"domainString": data.domainString,
		"domainSlice":  data.domainStringSlice,
		"urlString":    data.urlString,
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

import "os"

// Type w is an alias for []string to generate smaller source code
type w []string

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
		Tag{{identifier $t}} = {{index $.TagValues $t}}
	{{end}} \
	numTags = {{len .Tags}}
)

// TagStrings maps integer tag values to strings.
var TagStrings = map[Tags]string{
	{{range $t := .Tags}} \
		Tag{{identifier $t}}: "{{$t}}",
	{{end }}
}

// TagValues maps tag names to integer values.
var TagValues  = map[string]Tags{
	{{range $t := .Tags}} \
		"{{$t}}": Tag{{identifier $t}},
	{{end }}
}

// z is a static slice of Zones.
// Other global variables have pointers into this slice.
var z = make([]Zone, {{len .Zones}})

// Zones is a slice of all Zones in the database.
var Zones = z[:]

// TLDs is a slice of all top-level domain Zones.
var TLDs = z[:{{len .TLDs}}]

// ZoneMap maps domain names to Zones.
var ZoneMap map[string]*Zone

//go:noinline
func i0() {
{{range $i, $d := .Domains}} \
	{{$z := (index $.Zones $d)}} \
		z[{{$i}}] = Zone{ \
			{{quotedDomain $d}}, \
			{{if $z.IsIDN}}/* {{$d}} */{{end }} \
			{{if $z.ParentDomain}} &z[{{$z.ParentOffset}}] {{else}} r {{end}}, \
			{{if $z.SubdomainsEnd}} z[{{$z.SubdomainsOffset}}:{{$z.SubdomainsEnd}}] {{else}} x {{end}}, \
			{{if $z.TagBits}} {{printf "0x%x" $z.TagBits}} {{else}} 0 {{end}}, \
			{{quoted $z.RegistryOperator}}, \
			{{quotedURL $z.InfoURL}}, \
			{{if $z.NameServers}} w{ {{range $z.NameServers}}{{quotedDomain .}},{{end}}} {{else}} n {{end}}, \
			{{if $z.Wildcards}} w{ {{range $z.Wildcards}}{{quotedDomain .}},{{end}}} {{else}} n {{end}}, \
			{{if $z.Locations}} w{ {{range $z.Locations}}{{quoted .}},{{end}}} {{else}} n {{end}}, \
			{{if $z.Languages}} w{ {{range $z.Languages}}{{quoted .}},{{end}}} {{else}} n {{end}}, \
			{{quotedDomain $z.WhoisServer}}, \
			{{quotedURL $z.WhoisURL}}, \
			{{if $z.RDAPURLs}} w{ {{range $z.RDAPURLs}}{{quotedDomain .}},{{end}}} {{else}} n {{end}}, \
			{{if $z.IDNDisallowed}} f {{else}} t {{end}}, \
		}
{{if mod $i 8 7}} \
	}

//go:noinline
func i{{$i}}() {
{{end}} \
{{end}}
}

func initZones() {
	i0()
{{range $i, $d := .Domains}} \
	{{if mod $i 8 7}} \
i{{$i}}()
	{{end}} \
{{end}}
	ZoneMap = make(map[string]*Zone, len(z))
	for i := range z {
		ZoneMap[z[i].Domain] = &z[i]
	}
}

func init() {
	if os.Getenv("ZONEDB_SKIP_INIT") != "" {
		return
	}
	initZones() // Separate function to report allocs in initialization
}

`
