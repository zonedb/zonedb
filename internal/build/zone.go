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
	Domain      string   `json:"domain,omitempty"`
	InfoURL     string   `json:"infoURL,omitempty"`
	Tags        []string `json:"tags,omitempty"`
	Locations   []string `json:"locations,omitempty"`
	WhoisServer string   `json:"whoisServer,omitempty"`
	WhoisURL    string   `json:"whoisURL,omitempty"`
	NameServers []string `json:"nameServers,omitempty"`
	Wildcards   []string `json:"wildcards,omitempty"`
	Policies    []Policy `json:"policies,omitempty"`

	// Transitional
	CodePoints   string            `json:"codePoints,omitempty"`
	ProhibitIDN  bool              `json:"prohibitIDN,omitempty"`
	IDNTableURLs map[string]string `json:"idnTableURLs,omitempty"`

	// Internal
	subdomains     []string
	oldNameServers []string
	m              sync.Mutex

	// Exported for use in text/template
	IDNDisallowed                                 bool   `json:"-"`
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
	sort.Strings(z.NameServers)
	sort.Strings(z.Wildcards)
	sort.Strings(z.subdomains)
	z.normalizePolicies()
}

func (z *Zone) normalizePolicies() {
	// De-dupe
	var set = make(map[Policy]struct{}, len(z.Policies))
	for _, p := range z.Policies {
		if p.Type == TypeIDNDisallowed {
			z.IDNDisallowed = true
		}
		set[p] = struct{}{}
	}
	z.Policies = make([]Policy, 0, len(set))
	for p := range set {
		z.Policies = append(z.Policies, p)
	}

	// Sort
	sortPolicies(z.Policies)
}

func (z *Zone) transition() {
	// Transition ProhibitIDN and ASCII CodePoints
	if z.ProhibitIDN || z.CodePoints == "--09az" {
		z.AddPolicy(TypeIDNDisallowed, "", "", "")
	}

	// Transition IDNTableURLs
	for lang, u := range z.IDNTableURLs {
		z.AddPolicy(TypeIDNTable, lang, u, "")
	}

	// Transition policies
	for i := range z.Policies {
		p := &z.Policies[i]

		// Transition TypeMinLength to TypeLength
		if p.Type == "min-length" {
			p.Type = TypeLength
			p.Key = KeyMin
		}

		// Transition Language to Key
		if p.Language != "" {
			p.Key = p.Language
			p.Language = ""
		}

		// Transition invalid BCP 47 language tags
		if p.Type == TypeIDNTable && p.Key != "" {
			lang, err := normalizeLang(p.Key)
			if err != nil {
				Trace("Zone %s has policy with bad language: %s %s\n", z.Domain, p.Type, p.Key)
			} else if lang != p.Key {
				p.Key = lang
			}
		}
	}

	// Zero out legacy data
	z.ProhibitIDN = false
	z.IDNTableURLs = nil
	z.CodePoints = ""
}

// AddPolicy adds a single policy to Zone z.
func (z *Zone) AddPolicy(ptype, key, value, comment string) {
	z.Policies = append(z.Policies, Policy{
		Type:    ptype,
		Key:     key,
		Value:   value,
		Comment: comment,
	})
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
	j, _ := json.Marshal(z)
	j2, _ := json.Marshal(&Zone{Domain: z.Domain})
	if string(j) == string(j2) {
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

// IsTLD returns true if z is a TLD.
func (z *Zone) IsTLD() bool {
	return !strings.Contains(z.Domain, ".")
}

// Retire marks a zone as retired or withdrawn.
func (z *Zone) Retire() {
	// Brand TLDs are withdrawn, not retired
	for _, tag := range z.Tags {
		if tag == "brand" {
			z.Tags = append(z.Tags, "withdrawn")
			return
		}
	}
	z.Tags = append(z.Tags, "retired")
}

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
