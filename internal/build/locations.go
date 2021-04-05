package build

import (
	"os"

	"github.com/wsxiaoys/terminal/color"
)

func AddLocations(zones map[string]*Zone, locations []string) {
	var added, modified int
	for _, z := range zones {
		s := NewSet(z.Locations...)
		n := len(s)
		s.Add(locations...)
		m := len(s) - n
		if m != 0 {
			z.Locations = s.Values()
			added += m
			modified++
		}
	}
	color.Fprintf(os.Stderr, "@{.}Added %d locations(s) to %d zone(s)\n", added, modified)
}

func RemoveLocations(zones map[string]*Zone, locations []string) {
	var removed, modified int
	for _, z := range zones {
		s := NewSet(z.Locations...)
		n := len(s)
		for _, l := range locations {
			delete(s, l)
		}
		m := n - len(s)
		if m != 0 {
			z.Locations = s.Values()
			removed += m
			modified++
		}
	}
	color.Fprintf(os.Stderr, "@{.}Removed %d locations(s) from %d zone(s)\n", removed, modified)
}
