package build

import (
	"encoding/json"
	"sort"
	"strings"
	"sync"
	"unicode/utf8"

	"golang.org/x/net/idna"
)

// Zone represents a build-time DNS zone (public suffix).
type Zone struct {
	Domain       string            `json:"domain,omitempty"`
	InfoURL      string            `json:"infoURL,omitempty"`
	Tags         []string          `json:"tags,omitempty"`
	Locations    []string          `json:"locations,omitempty"`
	WhoisServer  string            `json:"whoisServer,omitempty"`
	WhoisURL     string            `json:"whoisURL,omitempty"`
	NameServers  []string          `json:"nameServers,omitempty"`
	Policies     []Policy          `json:"policies,omitempty"`
	IDNTableURLs map[string]string `json:"idnTableURLs,omitempty"`
	ProhibitIDN  bool              `json:"prohibitIDN,omitempty"`
	IDNTableURLs map[string]string `json:"idnTableURLs,omitempty"`

	// Internal
	subdomains []string
	m          sync.Mutex

	// Exported for use in text/template
	ParentOffset, SubdomainsOffset, SubdomainsEnd int    `json:"-"`
	TagBits                                       uint64 `json:"-"`
}

// Normalize formats a build.Zone into normal form suitable for serialization.
func (z *Zone) Normalize() {
	z.Domain = Normalize(z.Domain)
	var tags []string
	for _, t := range z.Tags {
		tags = append(tags, t)
	}
	z.Tags = NewSet(tags...).Values()
	z.NameServers = NewSet(z.NameServers...).Values()
	for lang, u := range z.IDNTableURLs {
		norm, err := normalizeLang(lang)
		if err == nil && norm != lang {
			delete(z.IDNTableURLs, lang)
			if _, ok := z.IDNTableURLs[norm]; !ok {
				z.IDNTableURLs[norm] = u
			}
		}
	}
}

// IsIDN returns true if the Zone label(s) use non-ASCII characters.
func (z *Zone) IsIDN() bool {
	for i := 0; i < len(z.Domain); i++ {
		if z.Domain[i] >= utf8.RuneSelf {
			return true
		}
	}
	return false
}

// HasMetadata returns true if the Zone has any metadata necessary for emitting a metadata JSON file.
func (z *Zone) HasMetadata() bool {
	z2 := *z
	z2.Domain = ""
	if j, _ := json.Marshal(&z2); string(j) == "{}" {
		return false
	}
	return true
}

// ASCII returns the ACE encoded form of the Zone’s label(s).
func (z *Zone) ASCII() string {
	s, _ := idna.ToASCII(z.Domain)
	return s
}

// ParentDomain returns the parent domain name (if any) for the Zone.
func (z *Zone) ParentDomain() string {
	labels := strings.Split(z.Domain, ".")
	if len(labels) == 1 {
		return ""
	}
	return strings.Join(labels[1:], ".")
}

type IDNCPIndexes [2]int

// TLDs filters a zone set for top-level domains.
func TLDs(zones map[string]*Zone) map[string]*Zone {
	tlds := make(map[string]*Zone)
	for d, z := range zones {
		if !strings.Contains(d, ".") {
			tlds[d] = z
		}
	}
	return tlds
}

// SortedDomains returns a list of domain names sorted by rank.
func SortedDomains(zones map[string]*Zone) []string {
	domains := make([]string, 0, len(zones))
	for d := range zones {
		domains = append(domains, d)
	}
	Sort(domains)
	return domains
}

// mapZones concurrently applies a function to a set of Zones.
func mapZones(zones map[string]*Zone, fn func(*Zone)) {
	limiter := make(chan struct{}, Concurrency)
	var wg sync.WaitGroup
	for _, z := range zones {
		limiter <- struct{}{}
		wg.Add(1)
		go func(z *Zone) {
			defer func() {
				<-limiter
				wg.Done()
			}()
			fn(z)
		}(z)
	}
	wg.Wait()
}

// Sort sorts a slice of domain names by rank. Rank sort defined as:
// 1. Label count (TLDs followed by second- and third-level domains), then
// 2. lexically sorted reversed label order (TLD first, then second-level label, etc.)
// Example: com, net, org, uk, uk.com, ac.uk, co.uk, ...
func Sort(domains []string) {
	sort.Sort(sortDomains(domains))
}

// sort.Interface implementation
type sortDomains []string

func (ds sortDomains) Len() int      { return len(ds) }
func (ds sortDomains) Swap(i, j int) { ds[i], ds[j] = ds[j], ds[i] }
func (ds sortDomains) Less(i, j int) bool {
	a, b := strings.Split(ds[i], "."), strings.Split(ds[j], ".")
	alen, blen := len(a), len(b)
	// Sort TLDs before second- and third-level domains
	if alen != blen {
		return alen < blen
	}
	// Sort
	for k := alen - 1; k >= 0; k-- {
		if a[k] != b[k] {
			return a[k] < b[k]
		}
	}
	return false
}
