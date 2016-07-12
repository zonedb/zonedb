package build

import "github.com/wsxiaoys/terminal/color"

func AddMetros(zones map[string]*Zone, metros []string) {
	var added, modified int
	for _, z := range zones {
		s := NewSet(z.Metros...)
		n := len(s)
		s.Add(metros...)
		m := len(s) - n
		if m != 0 {
			z.Metros = s.Values()
			added += m
			modified++
		}
	}
	color.Printf("@{.}Added %d metros(s) to %d zone(s)\n", added, modified)
}

func RemoveMetros(zones map[string]*Zone, metros []string) {
	var removed, modified int
	for _, z := range zones {
		s := NewSet(z.Metros...)
		n := len(s)
		for _, l := range metros {
			delete(s, l)
		}
		m := n - len(s)
		if m != 0 {
			z.Metros = s.Values()
			removed += m
			modified++
		}
	}
	color.Printf("@{.}Removed %d metros(s) from %d zone(s)\n", removed, modified)
}
