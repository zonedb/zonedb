package build

import (
	"os"

	"github.com/wsxiaoys/terminal/color"
)

// Well-known tags
const (
	TagAdult          = "adult"
	TagBrand          = "brand"
	TagCity           = "city"
	TagClosed         = "closed"
	TagCommunity      = "community"
	TagCountry        = "country"
	TagGeneric        = "generic"
	TagGeo            = "geo"
	TagInfrastructure = "infrastructure"
	TagPrivate        = "private"
	TagRegion         = "region"
	TagRetired        = "retired"
	TagSponsored      = "sponsored"
	TagTest           = "test"
	TagWithdrawn      = "withdrawn"
)

// AddTags adds one or more tags to a Zone.
func AddTags(zones map[string]*Zone, tags []string) {
	var added, modified int
	for _, z := range zones {
		d := z.AddTags(tags...)
		if d > 0 {
			modified++
			added += d
		}
	}
	color.Fprintf(os.Stderr, "@{.}Added %d tag(s) to %d zone(s)\n", added, modified)
}

// RemoveTags removes one or more tags from a Zone.
func RemoveTags(zones map[string]*Zone, tags []string) {
	var removed, modified int
	for _, z := range zones {
		d := z.RemoveTags(tags...)
		if d > 0 {
			modified++
			removed += d
		}
	}
	color.Fprintf(os.Stderr, "@{.}Removed %d tag(s) from %d zone(s)\n", removed, modified)
}

func tagBits(tagValues map[string]uint64, tags []string) (bits uint64) {
	for _, t := range tags {
		bits |= tagValues[t]
	}
	return
}
