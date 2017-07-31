package build

import (
	"os"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/wsxiaoys/terminal/color"
)

var (
	Verbose     bool
	Concurrency = 32
	Timeout     = 10 * time.Second
)

func LogFatal(err error) {
	color.Fprintf(os.Stderr, "@{r!}Fatal:@{r} %s\n", err)
	os.Exit(1)
}

func LogError(err error) {
	color.Fprintf(os.Stderr, "@{r!}Error:@{r} %s\n", err)
}

func LogWarning(err error) {
	color.Fprintf(os.Stderr, "@{y!}Warning:@{y} %s\n", err)
}

func LogWarningFor(err error, label string) {
	color.Fprintf(os.Stderr, "@{y!}Warning:@{y} %s@{|}{@.} [%s]\n", err, label)
}

func LogWarningForAt(err error, label string, offset int) {
	color.Fprintf(os.Stderr, "@{y!}Warning:@{y} %s@{|}@{.} [%s]@@%d\n", err, label, offset)
}

func Trace(spec string, args ...interface{}) {
	color.Fprintf(os.Stderr, spec, args...)
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
	for d, _ := range zones {
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

// IndexOfStrings finds a rank-sorted slice of domain names (needle)
// within a larger slice (haystack).
func IndexOfStrings(haystack []string, needle []string) int {
outer:
	for i := range haystack {
		for j := range needle {
			if i+j >= len(haystack) || needle[j] != haystack[i+j] {
				continue outer
			}
		}
		return i
	}
	return -1
}

// IndexOrAppendStrings finds or appends a slice of rank-sorted domain names (needle)
// Returns 0,0 for a zero-length needle.
func IndexOrAppendStrings(haystack *[]string, needle []string) (int, int) {
	if len(needle) == 0 {
		return 0, 0
	}
	idx := IndexOfStrings(*haystack, needle)
	if idx < 0 {
		idx = len(*haystack)
		*haystack = append(*haystack, needle...)
	}
	return idx, idx + len(needle)
}
