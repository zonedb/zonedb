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

// Tags are stored in a single integer as a bit field.
const (
	TagGeneric = 1 << iota
	TagGeo
	TagCity
	TagRegion
	TagCountry
	TagCommunity
	TagSponsored
	TagBrand
	TagAdult
	TagClosed
	TagPrivate
	TagInfrastructure
	TagRetired
	TagWithdrawn
	numTags = iota
)

// TagStrings maps integer tag values to strings.
var TagStrings = map[uint32]string{
	TagGeneric:        "generic",
	TagGeo:            "geo",
	TagCity:           "city",
	TagRegion:         "region",
	TagCountry:        "country",
	TagCommunity:      "community",
	TagSponsored:      "sponsored",
	TagBrand:          "brand",
	TagAdult:          "adult",
	TagClosed:         "closed",
	TagPrivate:        "private",
	TagInfrastructure: "infrastructure",
	TagRetired:        "retired",
	TagWithdrawn:      "withdrawn",
}

// TagValues maps tag names to integer values.
var TagValues map[string]uint32

func init() {
	TagValues = make(map[string]uint32, len(TagStrings))
	for t, s := range TagStrings {
		TagValues[s] = t
	}
}
