package zonedb

//go:generate go run build/cmd/zonedb/main.go -generate-go

// NS represents a slice of name servers.
type NS []string

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
	WhoisServer string

	// URL to look up whois info for a subdomain of this Zone
	WhoisURL string

	// Informational URL for this Zone
	InfoURL string

	// Tags stored as an integer bit field.
	Tags uint32
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
// undelegated, withdrawn, retired, or infrastructure zones.
func (z *Zone) AllowsRegistration() bool {
	t := z.Tags & (TagClosed | TagWithdrawn | TagRetired | TagInfrastructure)
	return t == 0 && z.IsDelegated()
}
