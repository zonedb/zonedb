package build

import (
	"fmt"
	"os"
	"slices"
	"sort"
	"strings"
	"sync"

	"github.com/wsxiaoys/terminal/color"
)

// NS debug logging.
//
// Enabled by setting ZONEDB_DEBUG_NS in the environment. Emits three kinds
// of lines to stderr:
//
//	NS_CHANGE step=fetch|verify  per-zone, only when the NS list changed.
//	NS_TRACE  step=fetch|verify  per-zone, always, but only for zones whose
//	                             domain matches one of the suffixes in
//	                             ZONEDB_DEBUG_NS_SUFFIX (comma-separated).
//	                             Used to capture "healthy" runs for
//	                             comparison against wipe runs.
//	PARENT_STATS                 one line per distinct parent host queried,
//	                             emitted at the end of FetchNameServers.
//
// Every emit includes a branch= tag identifying which code path fired, so the
// log can be grepped for specific preservation/removal behavior.
//
// Zero cost when ZONEDB_DEBUG_NS is unset. sync.Once caches the lookup, all
// record/emit helpers early-return.

type parentOutcome int

const (
	parentOutcomeNOERROR parentOutcome = iota
	parentOutcomeNXDOMAIN
	parentOutcomeTimeout
	parentOutcomeOtherErr
)

// fetchBranch tags which preservation/removal path fired in FetchNameServers.
type fetchBranch string

const (
	fetchBranchUnchanged          fetchBranch = "unchanged"            // old == new
	fetchBranchChanged            fetchBranch = "changed"              // normal update
	fetchBranchPreservedAllFailed fetchBranch = "preserved-all-failed" // PR #1124 escape hatch: successfulResponses==0
	fetchBranchLostAllNS          fetchBranch = "lost-all-ns"          // cleared despite having prior NS
)

// verifyBranch tags which path fired in VerifyNameServers.
type verifyBranch string

const (
	verifyBranchUnchanged verifyBranch = "unchanged"
	verifyBranchDropped   verifyBranch = "dropped"
)

type parentObservation struct {
	host       string
	outcome    string // "noerror" | "nxdomain" | "timeout" | "err"
	errMsg     string
	nsReturned int
}

type nsVerifyObservation struct {
	ns     string
	errMsg string // empty means verification succeeded
}

type parentStatRow struct {
	queries, noerror, nxdomain, timeout, otherErr int
}

var (
	nsDebugOnce     sync.Once
	nsDebugOn       bool
	nsDebugSuffixes []string

	parentStatsMu sync.Mutex
	parentStats   = map[string]*parentStatRow{}
)

// nsDebugEnabled reports whether debug logging is on. Cached after first call.
func nsDebugEnabled() bool {
	nsDebugOnce.Do(func() {
		nsDebugOn = os.Getenv("ZONEDB_DEBUG_NS") != ""
		if s := os.Getenv("ZONEDB_DEBUG_NS_SUFFIX"); s != "" {
			for _, part := range strings.Split(s, ",") {
				part = strings.TrimSpace(part)
				if part != "" {
					nsDebugSuffixes = append(nsDebugSuffixes, part)
				}
			}
		}
	})
	return nsDebugOn
}

// nsDebugMatchesTraceSuffix reports whether a zone domain is covered by the
// always-on NS_TRACE logging, based on ZONEDB_DEBUG_NS_SUFFIX. Returns false
// when no suffix is configured or when the domain matches none of them.
func nsDebugMatchesTraceSuffix(domain string) bool {
	if !nsDebugEnabled() || len(nsDebugSuffixes) == 0 {
		return false
	}
	for _, suf := range nsDebugSuffixes {
		if strings.HasSuffix(domain, suf) {
			return true
		}
	}
	return false
}

func resetParentStats() {
	parentStatsMu.Lock()
	defer parentStatsMu.Unlock()
	parentStats = map[string]*parentStatRow{}
}

func recordParentStat(host string, outcome parentOutcome) {
	if !nsDebugEnabled() {
		return
	}
	parentStatsMu.Lock()
	defer parentStatsMu.Unlock()
	row := parentStats[host]
	if row == nil {
		row = &parentStatRow{}
		parentStats[host] = row
	}
	row.queries++
	switch outcome {
	case parentOutcomeNOERROR:
		row.noerror++
	case parentOutcomeNXDOMAIN:
		row.nxdomain++
	case parentOutcomeTimeout:
		row.timeout++
	case parentOutcomeOtherErr:
		row.otherErr++
	}
}

func emitParentStats() {
	if !nsDebugEnabled() {
		return
	}
	// Snapshot under a single lock, then sort and emit without holding it.
	parentStatsMu.Lock()
	snapshot := make(map[string]parentStatRow, len(parentStats))
	for h, r := range parentStats {
		snapshot[h] = *r
	}
	parentStatsMu.Unlock()

	hosts := make([]string, 0, len(snapshot))
	for h := range snapshot {
		hosts = append(hosts, h)
	}
	sort.Strings(hosts)
	for _, h := range hosts {
		r := snapshot[h]
		color.Fprintf(os.Stderr, "PARENT_STATS host=%s queries=%d noerror=%d nxdomain=%d timeout=%d other_err=%d\n",
			h, r.queries, r.noerror, r.nxdomain, r.timeout, r.otherErr)
	}
}

// emitFetchObservation is called once per zone after FetchNameServers has
// finished processing it. Emits NS_CHANGE for changed zones, NS_TRACE for
// zones matching ZONEDB_DEBUG_NS_SUFFIX regardless of change.
func emitFetchObservation(
	z *Zone,
	parentNS []string,
	obs []parentObservation,
	counts map[string]int,
	successful int,
	branch fetchBranch,
) {
	if !nsDebugEnabled() {
		return
	}
	changed := branch != fetchBranchUnchanged
	trace := nsDebugMatchesTraceSuffix(z.Domain)
	if !changed && !trace {
		return
	}
	kind := "NS_CHANGE"
	if !changed {
		kind = "NS_TRACE"
	}
	var b strings.Builder
	fmt.Fprintf(&b, "%s step=fetch branch=%s zone=%s old=%s new=%s parents=%d successful=%d counts=%s\n",
		kind, branch, z.Domain,
		formatNSList(z.oldNameServers), formatNSList(z.NameServers),
		len(parentNS), successful, formatCounts(counts))
	for _, o := range obs {
		switch o.outcome {
		case "noerror":
			fmt.Fprintf(&b, "  parent=%s outcome=noerror ns_returned=%d\n", o.host, o.nsReturned)
		case "nxdomain":
			fmt.Fprintf(&b, "  parent=%s outcome=nxdomain\n", o.host)
		case "timeout":
			fmt.Fprintf(&b, "  parent=%s outcome=timeout err=%q\n", o.host, o.errMsg)
		case "err":
			fmt.Fprintf(&b, "  parent=%s outcome=err err=%q\n", o.host, o.errMsg)
		}
	}
	color.Fprintf(os.Stderr, "%s", b.String())
}

// emitVerifyObservation is the Verify-step analog of emitFetchObservation.
func emitVerifyObservation(
	z *Zone,
	prior []string,
	obs []nsVerifyObservation,
	branch verifyBranch,
) {
	if !nsDebugEnabled() {
		return
	}
	changed := branch != verifyBranchUnchanged
	trace := nsDebugMatchesTraceSuffix(z.Domain)
	if !changed && !trace {
		return
	}
	kind := "NS_CHANGE"
	if !changed {
		kind = "NS_TRACE"
	}
	var b strings.Builder
	fmt.Fprintf(&b, "%s step=verify branch=%s zone=%s old=%s new=%s\n",
		kind, branch, z.Domain, formatNSList(prior), formatNSList(z.NameServers))
	for _, o := range obs {
		if o.errMsg == "" {
			fmt.Fprintf(&b, "  ns=%s outcome=ok\n", o.ns)
		} else {
			fmt.Fprintf(&b, "  ns=%s outcome=fail err=%q\n", o.ns, o.errMsg)
		}
	}
	color.Fprintf(os.Stderr, "%s", b.String())
}

func formatNSList(ns []string) string {
	if len(ns) == 0 {
		return "[]"
	}
	return "[" + strings.Join(ns, ",") + "]"
}

func formatCounts(counts map[string]int) string {
	if len(counts) == 0 {
		return "{}"
	}
	keys := make([]string, 0, len(counts))
	for k := range counts {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	parts := make([]string, 0, len(keys))
	for _, k := range keys {
		parts = append(parts, fmt.Sprintf("%s:%d", k, counts[k]))
	}
	return "{" + strings.Join(parts, ",") + "}"
}

func slicesEqualUnordered(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	ac := slices.Clone(a)
	bc := slices.Clone(b)
	slices.Sort(ac)
	slices.Sort(bc)
	return slices.Equal(ac, bc)
}
