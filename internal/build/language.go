package build

import (
	"os"

	"github.com/wsxiaoys/terminal/color"
)

func GuessLanguage(zones map[string]*Zone) {
	guessed := make(map[string]int)
	var modified int
	for _, z := range zones {
		if z.Language != "" {
			continue
		}
		var lang string
		for _, p := range z.Policies {
			if p.Type != TypeIDNTable || p.Key == lang {
				continue
			}
			if lang != "" {
				lang = "" // Skip this zone if we find > 1 language table
				// color.Fprintf(os.Stderr, "@{.}Skipped language guessing for %s\n", z.Domain)
				break
			}
			lang = p.Key
		}
		if lang == "" {
			continue
		}
		z.Language = lang
		guessed[lang] = guessed[lang] + 1
		modified++
		color.Fprintf(os.Stderr, "@{.}Guessed language for @{!}%s@{.}: @{c}%s\n", z.Domain, z.Language)
	}
	color.Fprintf(os.Stderr, "@{.}Guessed %d language(s) on %d zone(s)\n", len(guessed), modified)
}
