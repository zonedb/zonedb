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

const (
	goSrc = `// Automatically generated

package zonedb

func init() {
	initZones()
}

// Separate function report allocs in initialization.
func initZones() {
	_z = _y
}

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
var Zones = _z[:]

// TLDs is a slice of all top-level domain Zones.
var TLDs = _z[:{{len .TLDs}}]

// _z is a static array of Zones.
// Other global variables have pointers into this array.
var _z [{{len .Zones}}]Zone

// _y and _z are separated to fix circular references.
var _y = [{{len .Zones}}]Zone{
	{{range $d := .Domains}} \
		{{$z := (index $.Zones $d)}} \
		{ \
			"{{ascii $d}}", \
			{{if $z.IsIDN}}/* {{$d}} */{{end }} \
			{{if $z.ParentDomain}} &_z[{{$z.POffset}}] {{else}} nil {{end}}, \
			{{if $z.SEnd}} _z[{{$z.SOffset}}:{{$z.SEnd}}] {{else}} nil {{end}}, \
			{{if $z.CPEnd}} _c[{{$z.CPOffset}}:{{$z.CPEnd}}] {{else}} nil {{end}}, \
			{{if $z.IDNCPs}} \
				IDNT{ \
				{{range $idnLang := $z.Langs}} \
					{{$idnCPIndexes := (index $z.IDNCPs $idnLang)}} \
					"{{ascii $idnLang}}": _c[{{index $idnCPIndexes 0}}:{{index $idnCPIndexes 1}}], \
				{{end}} \
				} \
			{{else}} \
				nil \
			{{end}}, \
			{{if $z.NameServers}} NS{ {{range $z.NameServers}}"{{ascii .}}",{{end}}} {{else}} nil {{end}}, \
			{{if $z.Locations}} L{ {{range $z.Locations}}"{{ascii .}}",{{end}}} {{else}} nil {{end}}, \
			"{{ascii $z.WhoisServer}}", \
			"{{ascii $z.WhoisURL}}", \
			"{{ascii $z.InfoURL}}", \
			{{printf "0x%x" $z.TagBits}}, \
		},
	{{end}} \
}

// ZoneMap maps Unicode domain names to Zones.
var ZoneMap = map[string]*Zone{
	{{range $d := .Domains}}  \
		{{$z := (index $.Zones $d)}} \
		{{$o := index $.Offsets $d}} \
		"{{ascii $d}}": &_z[{{$o}}], \
		{{if $z.IsIDN}}// {{$d}}{{end }}
	{{end}} \
}

// _c stores Unicode code point ranges allowed in each Zone by the registry.
// Rune values alternate low, high.
var _c = []rune("{{range .CodePoints}}{{printf "%c" .}}{{end}}")
`
)

// Experiment: split on each code point range from lo-hi
// var _c = []rune("{{range .CodePoints}}{{if rewound .}}" +
// 	"{{ end }}{{printf "%c" .}}{{end}}")

func cont(s string) string {
	return strings.Replace(s, "\\\n", "", -1)
}

func odd(i int) bool {
	return (i & 0x1) != 0
}

func mod0(i, radix int) bool {
	return (i % radix) == 0
}

var lastC rune

func rewound(c rune) (out bool) {
	if c < lastC {
		out = true
	}
	lastC = c
	return
}

var (
	funcMap = template.FuncMap{
		"odd":     odd,
		"mod0":    mod0,
		"rewound": rewound,
		"title":   strings.Title,
		"ascii":   ToASCII,
	}
	goTemplate = template.Must(template.New("").Funcs(funcMap).Parse(cont(goSrc)))
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

	// Sort domains by number of code points
	cpDomains := make([]string, len(domains))
	copy(cpDomains, domains)
	// Disabled for now (it doesn’t meaningfully change size of generated code)
	// sort.SliceStable(cpDomains, func(i, j int) bool {
	// 	return len(zones[cpDomains[j]].CodePoints) < len(zones[cpDomains[i]].CodePoints)
	// })

	var codePoints []rune
	for _, d := range cpDomains {
		z := zones[d]
		z.Normalize() // just in case
		z.POffset = offsets[z.ParentDomain()]
		if len(z.Subdomains) > 0 {
			z.SOffset = offsets[z.Subdomains[0]]
			z.SEnd = z.SOffset + len(z.Subdomains)
		}
		if z.ProhibitIDN {
			z.CPOffset, z.CPEnd = IndexOrAppendRunes(&codePoints, ASCII.Runes())
		} else {
			z.CPOffset, z.CPEnd = IndexOrAppendRunes(&codePoints, z.CodePoints.Runes())
		}
		z.IDNCPs = make(map[string]IDNCPIndexes)
		z.Langs = make([]string, 0, len(z.IDNTables))
		for lang := range z.IDNTables {
			z.Langs = append(z.Langs, lang)
		}
		sort.Strings(z.Langs)
		for _, lang := range z.Langs {
			idnOffset, idnEnd := IndexOrAppendRunes(&codePoints, z.IDNTables[lang].Runes())
			z.IDNCPs[lang] = IDNCPIndexes{idnOffset, idnEnd}
		}
		z.TagBits = tagBits(tagValues, z.Tags)
	}

	var singles int
	var gap int
	var max CodeRange
	for i := 0; i < len(codePoints); i += 2 {
		if codePoints[i] == codePoints[i+1] {
			singles++
		} else {
			n := int(codePoints[i+1] - codePoints[i])
			gap += n
			if n > int(max.Hi-max.Lo) {
				max.Lo = codePoints[i]
				max.Hi = codePoints[i+1]
			}
		}
	}
	ranges := len(codePoints) / 2
	pct := float64(singles) / float64(ranges) * 100
	avg := float64(gap) / float64(ranges)
	color.Fprintf(os.Stderr, "@{.}%d / %d single code point ranges (%.3f%%)\n", singles, ranges, pct)
	color.Fprintf(os.Stderr, "@{.}Range average: %.3f Max: %d (%s%s)\n", avg, (max.Hi - max.Lo), string(max.Lo), string(max.Hi))

	data := struct {
		Zones      map[string]*Zone
		TLDs       map[string]*Zone
		Domains    []string
		Offsets    map[string]int
		CodePoints []rune
		TagType    string
		Tags       []string
		TagValues  map[string]uint64
	}{
		zones,
		tlds,
		domains,
		offsets,
		codePoints,
		tagType,
		tags,
		tagValues,
	}

	buf := new(bytes.Buffer)
	err := goTemplate.Execute(buf, &data)
	if err != nil {
		return err
	}
	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		return err
	}

	fn := filepath.Join(BaseDir, "zones.go")
	color.Fprintf(os.Stderr, "@{.}Generating Go source code: %s\n", fn)
	f, err := os.Create(fn)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(formatted)
	if err != nil {
		return err
	}

	return nil
}
