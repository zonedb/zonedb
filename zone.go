package zonedb

import "strings"

//go:generate go run build/cmd/zonedb/main.go -generate-go

// NS represents a slice of name servers.
type NS []string

// L represents a slice of locations.
type L []string

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

// Zone represents a single DNS zone.
type Zone struct {
	// Normalized UTF-8 domain name
	Domain string

	// Parent Zone (nil if Zone is a TLD)
	Parent *Zone

	// Slice of subdomain (child) Zones (nil if empty)
	Subdomains []Zone

	// Unicode codepoint ranges allowed by the registry.
	// Alternating low, high (inclusive)
	CodePoints []rune

	// DNS name servers for the Zone
	NameServers NS

	// Whois server responding on port 43
	whoisServer string

	// URL to look up whois info for a subdomain of this Zone
	whoisURL string

	// Informational URL for this Zone
	InfoURL string

	// Tags stored as an integer bit field.
	Tags Tags

	// Locations associated with the Zone
	Locations L
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
// (lowercase, Unicode). Malformed input, IP addresses, or other non-domain
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
