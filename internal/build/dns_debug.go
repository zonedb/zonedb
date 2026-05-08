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

// NS flap debug logging. Gated by ZONEDB_DEBUG_NS; ZONEDB_DEBUG_NS_SUFFIX
// (comma-separated) additionally emits NS_TRACE for matching zones unchanged.

type parentOutcome int

const (
	parentOutcomeNOERROR parentOutcome = iota
	parentOutcomeNXDOMAIN
	parentOutcomeTimeout
	parentOutcomeOtherErr
)

// parentOutcomeStrings maps parentOutcome values to log-line labels.
var parentOutcomeStrings = map[parentOutcome]string{
	parentOutcomeNOERROR:  "noerror",
	parentOutcomeNXDOMAIN: "nxdomain",
	parentOutcomeTimeout:  "timeout",
	parentOutcomeOtherErr: "err",
}

// String returns the short label used in NS_CHANGE/NS_TRACE log lines.
func (o parentOutcome) String() string {
	return parentOutcomeStrings[o]
}

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
	outcome    parentOutcome
	errMsg     string // populated only on timeout/err
	nsReturned int    // populated only on noerror
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

// nsDebugMatchesTraceSuffix reports whether a zone's domain matches any
// suffix in ZONEDB_DEBUG_NS_SUFFIX.
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

// nsDebugger is the per-invocation debug state. Goroutine-safe; nil-safe.
type nsDebugger struct {
	mut         sync.Mutex
	parentStats map[string]*parentStatRow
}

// newNSDebugger returns an initialized debugger when ZONEDB_DEBUG_NS is set,
// nil otherwise. A nil *nsDebugger is the normal "debug off" value.
func newNSDebugger() *nsDebugger {
	if !nsDebugEnabled() {
		return nil
	}
	return &nsDebugger{parentStats: map[string]*parentStatRow{}}
}

// recordParentStat tallies a single parent-query outcome. No-op on nil.
func (d *nsDebugger) recordParentStat(host string, outcome parentOutcome) {
	if d == nil {
		return
	}
	d.mut.Lock()
	defer d.mut.Unlock()
	row := d.parentStats[host]
	if row == nil {
		row = &parentStatRow{}
		d.parentStats[host] = row
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

// EmitParentStats writes one PARENT_STATS line per distinct parent host seen.
// No-op on nil.
func (d *nsDebugger) EmitParentStats() {
	if d == nil {
		return
	}
	// Snapshot under a single lock, then sort and emit without holding it.
	d.mut.Lock()
	snapshot := make(map[string]parentStatRow, len(d.parentStats))
	for h, r := range d.parentStats {
		snapshot[h] = *r
	}
	d.mut.Unlock()

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

// emitKind decides whether a zone's observation should produce an NS_CHANGE
// line, an NS_TRACE line, or nothing. Returns empty string to skip.
func emitKind(domain string, changed bool) string {
	if changed {
		return "NS_CHANGE"
	}
	if nsDebugMatchesTraceSuffix(domain) {
		return "NS_TRACE"
	}
	return ""
}

// fetchZoneObserver collects per-parent observations for a single zone.
type fetchZoneObserver struct {
	debugger  *nsDebugger
	zone      *Zone
	perParent []parentObservation
}

// ForFetchZone returns an observer for a single zone. No-op on nil: returns
// nil, which every *fetchZoneObserver method tolerates.
func (d *nsDebugger) ForFetchZone(z *Zone, parentCount int) *fetchZoneObserver {
	if d == nil {
		return nil
	}
	return &fetchZoneObserver{
		debugger:  d,
		zone:      z,
		perParent: make([]parentObservation, 0, parentCount),
	}
}

// ObserveParent records a parent-query observation both per-zone (for emit)
// and in the cross-zone parent-stats tally. No-op on nil.
func (o *fetchZoneObserver) ObserveParent(observation parentObservation) {
	if o == nil {
		return
	}
	o.perParent = append(o.perParent, observation)
	o.debugger.recordParentStat(observation.host, observation.outcome)
}

// EmitFetch writes NS_CHANGE for changed zones and NS_TRACE for zones matching
// ZONEDB_DEBUG_NS_SUFFIX. No-op on nil.
func (o *fetchZoneObserver) EmitFetch(parentNS []string, counts map[string]int, successful int, branch fetchBranch) {
	if o == nil {
		return
	}
	kind := emitKind(o.zone.Domain, branch != fetchBranchUnchanged)
	if kind == "" {
		return
	}
	var b strings.Builder
	fmt.Fprintf(&b, "%s step=fetch branch=%s zone=%s old=%s new=%s parents=%d successful=%d counts=%s\n",
		kind, branch, o.zone.Domain,
		formatNSList(o.zone.oldNameServers), formatNSList(o.zone.NameServers),
		len(parentNS), successful, formatCounts(counts))
	for _, p := range o.perParent {
		switch p.outcome {
		case parentOutcomeNOERROR:
			fmt.Fprintf(&b, "  parent=%s outcome=%s ns_returned=%d\n", p.host, p.outcome, p.nsReturned)
		case parentOutcomeNXDOMAIN:
			fmt.Fprintf(&b, "  parent=%s outcome=%s\n", p.host, p.outcome)
		case parentOutcomeTimeout, parentOutcomeOtherErr:
			fmt.Fprintf(&b, "  parent=%s outcome=%s err=%q\n", p.host, p.outcome, p.errMsg)
		}
	}
	color.Fprintf(os.Stderr, "%s", b.String())
}

// verifyZoneObserver collects per-NS verification results for a single zone.
type verifyZoneObserver struct {
	zone  *Zone
	prior []string
	perNS []nsVerifyObservation
}

// ForVerifyZone returns an observer for a single zone. No-op on nil.
func (d *nsDebugger) ForVerifyZone(z *Zone, prior []string) *verifyZoneObserver {
	if d == nil {
		return nil
	}
	return &verifyZoneObserver{
		zone:  z,
		prior: prior,
		perNS: make([]nsVerifyObservation, 0, len(prior)),
	}
}

// ObserveNS records a verifyNS outcome. err is nil on success. No-op on nil.
func (o *verifyZoneObserver) ObserveNS(ns string, err error) {
	if o == nil {
		return
	}
	observation := nsVerifyObservation{ns: ns}
	if err != nil {
		observation.errMsg = err.Error()
	}
	o.perNS = append(o.perNS, observation)
}

// EmitVerify writes NS_CHANGE for dropped NSs and NS_TRACE for zones matching
// ZONEDB_DEBUG_NS_SUFFIX. No-op on nil.
func (o *verifyZoneObserver) EmitVerify(branch verifyBranch) {
	if o == nil {
		return
	}
	kind := emitKind(o.zone.Domain, branch != verifyBranchUnchanged)
	if kind == "" {
		return
	}
	var b strings.Builder
	fmt.Fprintf(&b, "%s step=verify branch=%s zone=%s old=%s new=%s\n",
		kind, branch, o.zone.Domain,
		formatNSList(o.prior), formatNSList(o.zone.NameServers))
	for _, p := range o.perNS {
		if p.errMsg == "" {
			fmt.Fprintf(&b, "  ns=%s outcome=ok\n", p.ns)
		} else {
			fmt.Fprintf(&b, "  ns=%s outcome=fail err=%q\n", p.ns, p.errMsg)
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
