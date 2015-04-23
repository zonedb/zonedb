package build

import (
	"encoding/json"
	"sort"
	"strings"

	"golang.org/x/net/idna"
)

type Zone struct {
	Domain      string    `json:"domain,omitempty"`
	InfoURL     string    `json:"infoURL,omitempty"`
	Tags        []string  `json:"tags,omitempty"`
	WhoisServer string    `json:"whoisServer,omitempty"`
	WhoisURL    string    `json:"whoisURL,omitempty"`
	NameServers []string  `json:"nameServers,omitempty"`
	CodePoints  CodeTable `json:"codePoints,omitempty"`
	Subdomains  []string  `json:"-"`

	// Exported for use in text/template
	POffset, SOffset, SEnd int `json:"-"`
	NSOffset, NSEnd        int `json:"-"`
	CPOffset, CPEnd        int `json:"-"`
}

func (z *Zone) Normalize() {
	z.Domain = Normalize(z.Domain)
	z.Tags = NewSet(z.Tags...).Values()
	z.NameServers = NewSet(z.NameServers...).Values()
	z.CodePoints.Compress()
}

func (z *Zone) HasMetadata() bool {
	z2 := *z
	z2.Domain = ""
	if j, _ := json.Marshal(&z2); string(j) == "{}" {
		return false
	}
	return true
}

func (z *Zone) ACE() string {
	s, _ := idna.ToASCII(z.Domain)
	return s
}

func (z *Zone) ParentDomain() string {
	labels := strings.Split(z.Domain, ".")
	if len(labels) == 1 {
		return ""
	}
	return strings.Join(labels[1:], ".")
}

// SortedZones returns a list of domain names sorted by complexity.
// More-complex Zones sort first to minimize diff size of generated code.
func SortedZones(zones map[string]*Zone) []*Zone {
	zs := make([]*Zone, 0, len(zones))
	for _, z := range zones {
		zs = append(zs, z)
	}
	sort.Sort(sortZones(zs))
	return zs
}

// sort.Interface implementation
type sortZones []*Zone

func (zs sortZones) Len() int      { return len(zs) }
func (zs sortZones) Swap(i, j int) { zs[i], zs[j] = zs[j], zs[i] }
func (zs sortZones) Less(i, j int) bool {
	a, b := zs[i], zs[j]
	switch {
	case len(a.NameServers) != len(b.NameServers):
		return len(a.NameServers) > len(b.NameServers)
	case len(a.CodePoints) != len(b.CodePoints):
		return len(a.CodePoints) > len(b.CodePoints)
	}
	return a.Domain < b.Domain
}
