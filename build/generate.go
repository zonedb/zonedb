package build

import (
	"bytes"
	"go/format"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/wsxiaoys/terminal/color"
)

const (
	goVarName = "_z"
	goSrc     = `// Automatically generated from zonedb data
// DO NOT EDIT

package zonedb

func init() {
	initZones()
}

// Separate function to fix circular reference and report allocs.
func initZones() {
	_z = _y
}

// _z is a static array of Zones.
// Other global variables have pointers into this array.
var _z [{{len .Zones}}]Zone

// Zones is a slice of all Zones in the database.
var Zones = _z[:]

// TLDs is a slice of all top-level domain Zones.
var TLDs = _z[:{{len .TLDs}}]

// ZoneMap is an index of domain names to Zones.
var ZoneMap = map[string]*Zone{
	{{range $d := .Domains}}  \
		{{$o := index $.Offsets $d}} \
		"{{$d}}": &_z[{{$o}}],
	{{end}} \
}

var _y = [{{len .Zones}}]Zone{
	{{range $d := .Domains}} \
		{{$z := (index $.Zones $d)}} \
		{ \
			"{{$d}}", \
			{{if $z.ParentDomain}} &_z[{{$z.POffset}}] {{else}} nil {{end}}, \
			{{if $z.SEnd}} _z[{{$z.SOffset}}:{{$z.SEnd}}] {{else}} nil {{end}}, \
			{{if $z.CPEnd}} _c[{{$z.CPOffset}}:{{$z.CPEnd}}] {{else}} nil {{end}}, \
			{{if $z.NSEnd}} _n[{{$z.NSOffset}}:{{$z.NSEnd}}] {{else}} nil {{end}}, \
			"{{$z.WhoisServer}}", \
			"{{$z.WhoisURL}}", \
			"{{$z.InfoURL}}", \
		},
	{{end}} \
}

var _n = [{{len .NameServers}}]string{
	{{range .NameServers}} \
		"{{.}}",
	{{end}} \
}
	
var _c = [{{len .CodePoints}}]rune{
	{{range $i, $cp := .CodePoints}} \
		'{{printf "%c" .}}', \
		{{if odd $i}}
		{{end}} \
	{{end}} \
}
`
)

func cont(s string) string {
	return strings.Replace(s, "\\\n", "", -1)
}

func odd(i int) bool {
	return (i & 0x1) != 0
}

var (
	funcMap    = template.FuncMap{"odd": odd}
	goTemplate = template.Must(template.New("go").Funcs(funcMap).Parse(cont(goSrc)))
)

func GenerateGo(zones map[string]*Zone) error {
	tlds := TLDs(zones)
	domains := SortedDomains(zones)
	zs := SortedZones(zones)
	offsets := make(map[string]int, len(domains))
	for i, d := range domains {
		offsets[d] = i
	}
	var nameServers []string
	var codePoints []rune
	for _, z := range zs {
		z.Normalize() // just in case
		z.POffset = offsets[z.ParentDomain()]
		if len(z.Subdomains) > 0 {
			z.SOffset = offsets[z.Subdomains[0]]
			z.SEnd = z.SOffset + len(z.Subdomains)
		}
		z.NSOffset, z.NSEnd = IndexOrAppendStrings(&nameServers, z.NameServers)
		z.CPOffset, z.CPEnd = IndexOrAppendRunes(&codePoints, z.CodePoints.Runes())
	}

	data := struct {
		Zones       map[string]*Zone
		TLDs        map[string]*Zone
		Domains     []string
		Offsets     map[string]int
		NameServers []string
		CodePoints  []rune
	}{
		zones,
		tlds,
		domains,
		offsets,
		nameServers,
		codePoints,
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
	return err
}
