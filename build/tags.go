package build

import (
	"fmt"

	"github.com/wsxiaoys/terminal/color"
	"github.com/zonedb/zonedb"
)

func AddTags(zones map[string]*Zone, tags []string) {
	for _, t := range tags {
		verifyTag(t)
	}
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
	color.Printf("@{.}Added %d tag(s) to %d zone(s)\n", added, modified)
}

func RemoveTags(zones map[string]*Zone, tags []string) {
	for _, t := range tags {
		verifyTag(t)
	}
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
	color.Printf("@{.}Removed %d tag(s) from %d zone(s)\n", removed, modified)
}

var unknownTags = NewSet()

func verifyTag(t string) bool {
	if unknownTags[t] {
		return false
	}
	if _, ok := zonedb.TagValues[t]; ok {
		return true
	}
	unknownTags.Add(t)
	LogWarning(fmt.Errorf("Unknown tag: %s", t))
	return false
}

func tagBits(tags []string) uint32 {
	var bits uint32
	for _, t := range tags {
		if verifyTag(t) {
			bits |= zonedb.TagValues[t]
		}
	}
	return bits
}
