package build

import "github.com/wsxiaoys/terminal/color"

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
	color.Printf("@{.}Added %d tag(s) to %d zone(s)\n", added, modified)
}

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
	color.Printf("@{.}Removed %d tag(s) from %d zone(s)\n", removed, modified)
}

func tagBits(tagValues map[string]uint32, tags []string) uint32 {
	var bits uint32
	for _, t := range tags {
		bits |= tagValues[t]
	}
	return bits
}
