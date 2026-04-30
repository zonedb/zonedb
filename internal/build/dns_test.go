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

	if err := FetchNameServers(map[string]*Zone{"example.tld": child}, allZones); err != nil {
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

	if err := FetchNameServers(map[string]*Zone{"example.tld": child}, allZones); err != nil {
		t.Fatalf("FetchNameServers returned error: %v", err)
	}

	if len(child.NameServers) != 0 {
		t.Errorf("NS list should be cleared on authoritative NXDOMAIN; got %v", child.NameServers)
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

	if err := FetchNameServers(map[string]*Zone{"example.tld": child}, allZones); err != nil {
		t.Fatalf("FetchNameServers returned error: %v", err)
	}

	got := sorted(child.NameServers)
	want := sorted([]string{"ns-a.example", "ns-b.example", "ns-extra.example"})
	if !slices.Equal(got, want) {
		t.Errorf("previously-known NS should not be dropped by consensus filter; got %v, want %v", got, want)
	}
}
