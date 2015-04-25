// Automatically generated from zonedb data

package zonedb

// Tags are stored in a single integer as a bit field.
const (
	TagAdult          = 1
	TagBrand          = 2
	TagClosed         = 4
	TagCommunity      = 8
	TagCountry        = 16
	TagGeneric        = 32
	TagGeo            = 64
	TagInfrastructure = 128
	TagPrivate        = 256
	TagRetired        = 512
	TagSponsored      = 1024
	TagWithdrawn      = 2048
	numTags           = 12
)

// TagStrings maps integer tag values to strings.
var TagStrings = map[uint32]string{
	TagAdult:          "adult",
	TagBrand:          "brand",
	TagClosed:         "closed",
	TagCommunity:      "community",
	TagCountry:        "country",
	TagGeneric:        "generic",
	TagGeo:            "geo",
	TagInfrastructure: "infrastructure",
	TagPrivate:        "private",
	TagRetired:        "retired",
	TagSponsored:      "sponsored",
	TagWithdrawn:      "withdrawn",
}

// TagValues maps tag names to integer values.
var TagValues = map[string]uint32{
	"adult":          TagAdult,
	"brand":          TagBrand,
	"closed":         TagClosed,
	"community":      TagCommunity,
	"country":        TagCountry,
	"generic":        TagGeneric,
	"geo":            TagGeo,
	"infrastructure": TagInfrastructure,
	"private":        TagPrivate,
	"retired":        TagRetired,
	"sponsored":      TagSponsored,
	"withdrawn":      TagWithdrawn,
}
