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

var (
	_z [{{len .Zones}}]Zone
	Zones = _z[:]
	TLDs = _z[:{{len .TLDs}}]
)

func init() {
	initZones()
}

// Separate function to report allocs
func initZones() {
	_z = [{{len .Zones}}]Zone{
		{{range $d := .Domains}} \
			{{$z := (index $.Zones $d)}} \
			{ \
				"{{$d}}", \
				{{if $z.ParentDomain}} &_z[{{$z.POffset}}] {{else}} nil {{end}}, \
				{{if $z.SEnd}} _z[{{$z.SOffset}}:{{$z.SEnd}}] {{else}} nil {{end}}, \
				{{if $z.NSEnd}} _n[{{$z.NSOffset}}:{{$z.NSEnd}}] {{else}} nil {{end}}, \
				"{{$z.WhoisServer}}", \
				"{{$z.WhoisURL}}", \
				"{{$z.InfoURL}}", \
			},
		{{end}} \
	}
}

var (
	_n = [{{len .NameServers}}]string{
		{{range .NameServers}} \
			"{{.}}",
		{{end}} \
	}
)
`
)

var goTemplate = template.Must(template.New("go").Parse(cont(goSrc)))

func cont(s string) string {
	return strings.Replace(s, "\\\n", "", -1)
}

func GenerateGo(zones map[string]*Zone) error {
	tlds := TLDs(zones)
	domains := SortedDomains(zones)
	offsets := make(map[string]int, len(domains))
	for i, d := range domains {
		offsets[d] = i
	}
	var nameServers []string
	for _, d := range domains {
		z := zones[d]
		z.Normalize() // just in case
		z.POffset = offsets[z.ParentDomain()]
		if len(z.Subdomains) > 0 {
			z.SOffset = offsets[z.Subdomains[0]]
			z.SEnd = z.SOffset + len(z.Subdomains)
		}
		z.NSOffset, z.NSEnd = IndexOrAppend(&nameServers, z.NameServers)
	}

	data := struct {
		Zones       map[string]*Zone
		TLDs        map[string]*Zone
		Domains     []string
		NameServers []string
	}{
		zones,
		tlds,
		domains,
		nameServers,
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
