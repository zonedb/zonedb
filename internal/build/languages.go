package build

import (
	"os"
	"strings"

	"github.com/wsxiaoys/terminal/color"
)

func AddLanguages(zones map[string]*Zone, languages []string) {
	var added, modified int
	for _, z := range zones {
		s := NewSet(z.Languages...)
		n := len(s)
		s.Add(languages...)
		m := len(s) - n
		if m != 0 {
			z.Languages = s.Values()
			z.normalizeLanguages()
			m = len(z.Languages) - n
			added += m
			if m != 0 {
				modified++
			}
		}
	}
	color.Fprintf(os.Stderr, "@{.}Added %d language(s) to %d zone(s)\n", added, modified)
}

func GuessLanguages(zones map[string]*Zone) {
	guessed := make(map[string]int)
	var modified int
	for _, z := range zones {
		if len(z.Languages) != 0 {
			continue
		}
		for _, p := range z.Policies {
			if p.Type != TypeIDNTable || p.Key == "" {
				continue
			}
			z.Languages = append(z.Languages, p.Key)
			guessed[p.Key] = guessed[p.Key] + 1
		}
		modified++
		color.Fprintf(os.Stderr, "@{.}Guessed language(s) for @{!}%s@{.}: @{c}%s\n", z.Domain, strings.Join(z.Languages, " "))
	}
	color.Fprintf(os.Stderr, "@{.}Guessed %d language(s) on %d zone(s)\n", len(guessed), modified)
}
