package build

import (
	"context"
	"slices"
	"sort"
	"testing"

	"github.com/miekg/dns"
)

// scriptedExchange builds a drop-in replacement for the package-level exchange
// var, used by tests to return canned DNS responses per parent hostname.
func scriptedExchange(t *testing.T, scenarios map[string]scriptedResponse) func(ctx context.Context, host, qname string, qtype uint16) (*dns.Msg, error) {
	t.Helper()
	return func(ctx context.Context, host, qname string, qtype uint16) (*dns.Msg, error) {
		s, ok := scenarios[host]
		if !ok {
			t.Fatalf("scriptedExchange: unexpected host %q", host)
		}
		if s.timeout {
			return nil, context.DeadlineExceeded
		}
		msg := &dns.Msg{}
		msg.SetQuestion(dns.Fqdn(qname), qtype)
		msg.Response = true
		msg.Rcode = s.rcode
		for _, ns := range s.nsRecords {
			msg.Ns = append(msg.Ns, &dns.NS{
				Hdr: dns.RR_Header{Name: dns.Fqdn(qname), Rrtype: dns.TypeNS, Class: dns.ClassINET, Ttl: 3600},
				Ns:  dns.Fqdn(ns),
			})
		}
		return msg, nil
	}
}

type scriptedResponse struct {
	timeout   bool
	rcode     int
	nsRecords []string
}

func sorted(s []string) []string {
	out := slices.Clone(s)
	sort.Strings(out)
	return out
}

// TestFetchNameServers_AllParentsTimeOut verifies that when every parent NS
// query fails at the network layer, the zone's prior NS list is preserved.
// This is the regression guard for the nightly-metadata flap.
func TestFetchNameServers_AllParentsTimeOut(t *testing.T) {
	prior := []string{"ns-a.example", "ns-b.example"}
	parent := &Zone{Domain: "tld", NameServers: []string{"p1", "p2", "p3", "p4", "p5"}}
	child := &Zone{Domain: "example.tld", NameServers: slices.Clone(prior)}
	allZones := map[string]*Zone{"tld": parent, "example.tld": child}

	origExchange := exchange
	t.Cleanup(func() { exchange = origExchange })
	exchange = scriptedExchange(t, map[string]scriptedResponse{
		"p1": {timeout: true},
		"p2": {timeout: true},
		"p3": {timeout: true},
		"p4": {timeout: true},
		"p5": {timeout: true},
	})

	if err := FetchNameServers(t.Context(), map[string]*Zone{"example.tld": child}, allZones); err != nil {
		t.Fatalf("FetchNameServers returned error: %v", err)
	}

	got := sorted(child.NameServers)
	want := sorted(prior)
	if !slices.Equal(got, want) {
		t.Errorf("NS list should be preserved when all parents fail; got %v, want %v", got, want)
	}
}

// TestFetchNameServers_AllParentsReturnNXDOMAIN verifies that when every
// parent authoritatively reports the zone as nonexistent, the NS list is
// cleared (not preserved). This guards against immortalizing retired zones.
func TestFetchNameServers_AllParentsReturnNXDOMAIN(t *testing.T) {
	prior := []string{"ns-a.example", "ns-b.example"}
	parent := &Zone{Domain: "tld", NameServers: []string{"p1", "p2", "p3", "p4", "p5"}}
	child := &Zone{Domain: "example.tld", NameServers: slices.Clone(prior)}
	allZones := map[string]*Zone{"tld": parent, "example.tld": child}

	origExchange := exchange
	t.Cleanup(func() { exchange = origExchange })
	exchange = scriptedExchange(t, map[string]scriptedResponse{
		"p1": {rcode: dns.RcodeNameError},
		"p2": {rcode: dns.RcodeNameError},
		"p3": {rcode: dns.RcodeNameError},
		"p4": {rcode: dns.RcodeNameError},
		"p5": {rcode: dns.RcodeNameError},
	})

	if err := FetchNameServers(t.Context(), map[string]*Zone{"example.tld": child}, allZones); err != nil {
		t.Fatalf("FetchNameServers returned error: %v", err)
	}

	if len(child.NameServers) != 0 {
		t.Errorf("NS list should be cleared on authoritative NXDOMAIN; got %v", child.NameServers)
	}
}

// 2 of 10 parents return extras: count==2, max==10, filter keeps them.
func TestFetchNameServers_PartialConsensus_TwoParents_AddsExtras(t *testing.T) {
	core := []string{"ns1.core.example", "ns2.core.example", "ns3.core.example", "ns4.core.example"}
	extras := []string{"extra-a.example", "extra-b.example"}

	parent := &Zone{Domain: "tld", NameServers: []string{"p1", "p2", "p3", "p4", "p5", "p6", "p7", "p8", "p9", "p10"}}
	child := &Zone{Domain: "child.tld", NameServers: slices.Clone(core)}
	allZones := map[string]*Zone{"tld": parent, "child.tld": child}

	origExchange := exchange
	t.Cleanup(func() { exchange = origExchange })
	// 2 parents return core + extras; 8 return core only.
	withExtras := append(slices.Clone(core), extras...)
	exchange = scriptedExchange(t, map[string]scriptedResponse{
		"p1":  {nsRecords: withExtras},
		"p2":  {nsRecords: withExtras},
		"p3":  {nsRecords: core},
		"p4":  {nsRecords: core},
		"p5":  {nsRecords: core},
		"p6":  {nsRecords: core},
		"p7":  {nsRecords: core},
		"p8":  {nsRecords: core},
		"p9":  {nsRecords: core},
		"p10": {nsRecords: core},
	})

	if err := FetchNameServers(t.Context(), map[string]*Zone{"child.tld": child}, allZones); err != nil {
		t.Fatalf("FetchNameServers returned error: %v", err)
	}

	got := sorted(child.NameServers)
	want := sorted(append(slices.Clone(core), extras...))
	if !slices.Equal(got, want) {
		t.Errorf("extras returned by 2/10 parents (count==2) should be kept; got %v, want %v", got, want)
	}
}

// 1 of 10 parents returns extras: count==1, max==10, filter drops them.
func TestFetchNameServers_PartialConsensus_OneParent_DropsExtras(t *testing.T) {
	core := []string{"ns1.core.example", "ns2.core.example", "ns3.core.example", "ns4.core.example"}
	extras := []string{"extra-a.example", "extra-b.example"}

	parent := &Zone{Domain: "tld", NameServers: []string{"p1", "p2", "p3", "p4", "p5", "p6", "p7", "p8", "p9", "p10"}}
	child := &Zone{Domain: "child.tld", NameServers: slices.Clone(core)}
	allZones := map[string]*Zone{"tld": parent, "child.tld": child}

	origExchange := exchange
	t.Cleanup(func() { exchange = origExchange })
	// 1 parent returns core + extras; 9 return core only.
	withExtras := append(slices.Clone(core), extras...)
	exchange = scriptedExchange(t, map[string]scriptedResponse{
		"p1":  {nsRecords: withExtras},
		"p2":  {nsRecords: core},
		"p3":  {nsRecords: core},
		"p4":  {nsRecords: core},
		"p5":  {nsRecords: core},
		"p6":  {nsRecords: core},
		"p7":  {nsRecords: core},
		"p8":  {nsRecords: core},
		"p9":  {nsRecords: core},
		"p10": {nsRecords: core},
	})

	if err := FetchNameServers(t.Context(), map[string]*Zone{"child.tld": child}, allZones); err != nil {
		t.Fatalf("FetchNameServers returned error: %v", err)
	}

	got := sorted(child.NameServers)
	want := sorted(core)
	if !slices.Equal(got, want) {
		t.Errorf("extras returned by 1/10 parents (count==1, max==10) should be dropped by consensus filter; got %v, want %v", got, want)
	}
}

// Extras in oldNameServers survive count==1 via PR #1124's escape hatch.
func TestFetchNameServers_PartialConsensus_OneParent_KeepsExtrasIfKnown(t *testing.T) {
	core := []string{"ns1.core.example", "ns2.core.example", "ns3.core.example", "ns4.core.example"}
	extras := []string{"extra-a.example", "extra-b.example"}
	prior := append(slices.Clone(core), extras...)

	parent := &Zone{Domain: "tld", NameServers: []string{"p1", "p2", "p3", "p4", "p5", "p6", "p7", "p8", "p9", "p10"}}
	child := &Zone{Domain: "child.tld", NameServers: slices.Clone(prior)}
	allZones := map[string]*Zone{"tld": parent, "child.tld": child}

	origExchange := exchange
	t.Cleanup(func() { exchange = origExchange })
	withExtras := append(slices.Clone(core), extras...)
	exchange = scriptedExchange(t, map[string]scriptedResponse{
		"p1":  {nsRecords: withExtras},
		"p2":  {nsRecords: core},
		"p3":  {nsRecords: core},
		"p4":  {nsRecords: core},
		"p5":  {nsRecords: core},
		"p6":  {nsRecords: core},
		"p7":  {nsRecords: core},
		"p8":  {nsRecords: core},
		"p9":  {nsRecords: core},
		"p10": {nsRecords: core},
	})

	if err := FetchNameServers(t.Context(), map[string]*Zone{"child.tld": child}, allZones); err != nil {
		t.Fatalf("FetchNameServers returned error: %v", err)
	}

	got := sorted(child.NameServers)
	want := sorted(prior)
	if !slices.Equal(got, want) {
		t.Errorf("known extras (in oldNameServers) should survive count==1 filter; got %v, want %v", got, want)
	}
}

// Verifies that a previously-known NS still returned by at least one
// parent survives the consensus filter, even when count==1 and max>2.
// An NS absent from every parent response is dropped by a different
// codepath, not exercised here.
func TestFetchNameServers_ConsensusFilterKeepsKnownNS(t *testing.T) {
	prior := []string{"ns-a.example", "ns-b.example", "ns-extra.example"}
	parent := &Zone{Domain: "tld", NameServers: []string{"p1", "p2", "p3", "p4", "p5"}}
	child := &Zone{Domain: "example.tld", NameServers: slices.Clone(prior)}
	allZones := map[string]*Zone{"tld": parent, "example.tld": child}

	origExchange := exchange
	t.Cleanup(func() { exchange = origExchange })
	exchange = scriptedExchange(t, map[string]scriptedResponse{
		"p1": {nsRecords: []string{"ns-a.example", "ns-b.example"}},
		"p2": {nsRecords: []string{"ns-a.example", "ns-b.example"}},
		"p3": {nsRecords: []string{"ns-a.example", "ns-b.example"}},
		"p4": {nsRecords: []string{"ns-a.example", "ns-b.example"}},
		"p5": {nsRecords: []string{"ns-a.example", "ns-b.example", "ns-extra.example"}},
	})

	if err := FetchNameServers(t.Context(), map[string]*Zone{"example.tld": child}, allZones); err != nil {
		t.Fatalf("FetchNameServers returned error: %v", err)
	}

	got := sorted(child.NameServers)
	want := sorted([]string{"ns-a.example", "ns-b.example", "ns-extra.example"})
	if !slices.Equal(got, want) {
		t.Errorf("previously-known NS should not be dropped by consensus filter; got %v, want %v", got, want)
	}
}
