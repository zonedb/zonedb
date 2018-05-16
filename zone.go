package zonedb

import (
	"strings"
)

//go:generate go run cmd/zonedb/main.go -generate-go

// And performs a bitwise AND between tags and q,
// comparing the result to zero. Returns true if any
// tags match q.
func (tags Tags) And(q Tags) bool {
	if (tags & q) != 0 {
		return true
	}
	return false
}

// String returns a space-delimited list of values for tags.
func (tags Tags) String() string {
	var a [numTags]string
	s := a[0:0]
	for i := uint64(0); i < numTags; i++ {
		t := Tags(1 << i)
		if (tags & t) != 0 {
			s = append(s, TagStrings[t])
		}
	}
	return strings.Join(s, " ")
}

// Zone represents a single DNS zone (a public suffix), where subdomains may be registered or created.
type Zone struct {
	// Normalized (ASCII, punycode) fully-qualified domain name
	Domain string

	// Parent Zone (nil if Zone is a TLD)
	Parent *Zone

	// Slice of subdomain (child) Zones (nil if empty)
	Subdomains []Zone

	// CodePoints is deprecated and IDNTables has been removed. We searched for public code
	// on GitHub that imported the zonedb package and found no packages using these symbols.
	// CodePoints []rune

	// DNS name servers for the Zone
	NameServers []string

	// Wildcard addresses for unregistered subdomains
	Wildcards []string

	// Locations associated with the Zone
	Locations []string

	// Whois server responding on port 43
	whoisServer string

	// URL to look up whois info for a subdomain of this Zone
	whoisURL string

	// Informational URL for this Zone
	InfoURL string

	// Tags stored as an integer bit field
	Tags Tags

	// Transitional: does the zone operator allow registration of non-ASCII subdomains?
	allowsIDN bool
}

// WhoisServer returns the whois server that responds on port 43
// for the zone. It first searches the specific zone, then the parent,
// returning an empty string if none found.
func (z *Zone) WhoisServer() string {
	if z.whoisServer != "" {
		return z.whoisServer
	}
	if z.Parent != nil {
		return z.Parent.WhoisServer()
	}
	return ""
}

// WhoisURL returns a URL to retrieve whois data for a subdomain
// of the zone. It first searches the specific zone, then the parent,
// returning an empty string if none found.
func (z *Zone) WhoisURL() string {
	if z.whoisURL != "" {
		return z.whoisURL
	}
	if z.Parent != nil {
		return z.Parent.WhoisURL()
	}
	return ""
}

// IsTLD returns true if the Zone is a top-level domain.
func (z *Zone) IsTLD() bool {
	return z.Parent == nil
}

// IsDelegated returns true if the Zone has name servers.
func (z *Zone) IsDelegated() bool {
	return len(z.NameServers) != 0
}

// IsInRootZone returns true if the Zone is a top-level domain.
// present in the root DNS zone.
func (z *Zone) IsInRootZone() bool {
	return z.IsTLD() && z.IsDelegated()
}

// IsValidDomain and IDNTable have been removed. We searched for public code
// on GitHub that imported the zonedb package and found no open-source clients
// using these methods.
//
// Rationale: the generation and validation of (IDN) labels requires state
// beyond what can be represented in simple list if code points.
//
// This branch attempts to reconcile the objective of documenting and normalizing
// zone metadata, including policies that define allowed labels such as
// IDN tables, Label Generation Rulesets (LGRs), and implicit or undocumented
// policies such as applied by ccTLDs.

// AllowsIDN returns true if the zone operator (registry)
// permits registration of non-ASCII labels under this Zone.
func (z *Zone) AllowsIDN() bool {
	return z.allowsIDN
}

// AllowsRegistration returns true if the Zone’s authority (registry)
// permits registration of subdomains of this Zone. Examples:
// closed, withdrawn, retired, or infrastructure zones.
func (z *Zone) AllowsRegistration() bool {
	return !z.Tags.And(TagClosed | TagWithdrawn | TagRetired | TagInfrastructure)
}

// IsZone returns true if the input domain is a Zone.
func IsZone(domain string) bool {
	_, ok := ZoneMap[domain]
	return ok
}

// IsTLD returns true if the input domain is a top-level domain.
func IsTLD(domain string) bool {
	z, ok := ZoneMap[domain]
	if !ok {
		return false
	}
	return z.IsTLD()
}

// PublicZone returns the public zone for a given domain name
// or nil if none found.
// Input must be normalized by the client (lowercase, ASCII-encoded).
func PublicZone(domain string) *Zone {
	sfx := domain
	for {
		if z, ok := ZoneMap[sfx]; ok {
			return z
		}
		if i := strings.Index(sfx, "."); i >= 0 {
			sfx = sfx[i+1:]
		} else {
			break
		}
	}
	return nil
}

// List implements the cookiejar.PublicSuffixList interface.
var List list

type list struct{}

// PublicSuffix returns the public suffix (zone) for a given domain name
// by calling PublicZone. Input must be normalized by the client
// (lowercase, ASCII, punycode). Malformed input, IP addresses, or other non-domain
// name strings will be returned unmodified.
func (l list) PublicSuffix(domain string) string {
	z := PublicZone(domain)
	if z == nil {
		return domain
	}
	return z.Domain
}

// String returns a description of source of this public suffix list.
func (l list) String() string {
	return "ZoneDB"
}
