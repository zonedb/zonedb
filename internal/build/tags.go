package build

import (
	"os"

	"github.com/wsxiaoys/terminal/color"
)

// AddTags adds one or more tags to a Zone.
func AddTags(zones map[string]*Zone, tags []string) {
	var added, modified int
	for _, z := range zones {
		s := NewSet(z.Tags...)
		n := len(s)
		s.Add(tags...)
		m := len(s) - n
		if m != 0 {
			z.Tags = s.Values()
			added += m
			modified++
		}
	}
	color.Fprintf(os.Stderr, "@{.}Added %d tag(s) to %d zone(s)\n", added, modified)
}

// RemoveTags removes one or more tags from a Zone.
func RemoveTags(zones map[string]*Zone, tags []string) {
	var removed, modified int
	for _, z := range zones {
		s := NewSet(z.Tags...)
		n := len(s)
		for _, t := range tags {
			delete(s, t)
		}
		m := n - len(s)
		if m != 0 {
			z.Tags = s.Values()
			removed += m
			modified++
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
