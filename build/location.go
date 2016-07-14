package build

import "github.com/wsxiaoys/terminal/color"

type Location struct {
	Cities     []string `json:"cities,omitempty"`
	ISOCodes   []string `json:"iso_codes,omitempty"`
	MetroCodes []string `json:"metro_codes,omitempty"`
}

func (l Location) GetField(field string) []string {
	switch field {
	case "cities":
		return l.Cities
	case "iso":
		return l.ISOCodes
	case "metro":
		return l.MetroCodes
	}
	return nil
}

func (l *Location) SetField(field string, values []string) {
	switch field {
	case "cities":
		l.Cities = values
	case "iso":
		l.ISOCodes = values
	case "metro":
		l.MetroCodes = values
	}
}

func AddLocationField(zones map[string]*Zone, field string, values []string) {
	var added, modified int
	for _, z := range zones {
		if z.Location == nil {
			z.Location = &Location{}
		}
		s := NewSet(z.Location.GetField(field)...)
		n := len(s)
		s.Add(values...)
		m := len(s) - n
		if m != 0 {
			z.Location.SetField(field, s.Values())
			added += m
			modified++
		}
	}
	color.Printf("@{.}Added %d location values(s) to %d zone(s)\n", added, modified)
}

func RemoveLocationField(zones map[string]*Zone, field string, values []string) {
	var removed, modified int
	for _, z := range zones {
		if z.Location == nil {
			z.Location = &Location{}
		}
		s := NewSet(z.Location.GetField(field)...)
		n := len(s)
		for _, l := range values {
			delete(s, l)
		}
		m := n - len(s)
		if m != 0 {
			z.Location.SetField(field, s.Values())
			removed += m
			modified++
		}
	}
	color.Printf("@{.}Removed %d location values(s) from %d zone(s)\n", removed, modified)
}
