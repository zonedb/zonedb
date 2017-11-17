package build

import (
	"strings"

	"golang.org/x/net/idna"
	"golang.org/x/text/language"
	"golang.org/x/text/unicode/norm"
)

// Normalize normalizes a domain name into valid, lowercase, Unicode form.
func Normalize(s string) string {
	s = norm.NFKC.String(s)
	s = strings.Trim(s, "-.")
	s = strings.ToLower(s)
	s, _ = idna.ToASCII(s)
	s, _ = idna.ToUnicode(s)
	return s
}

// ToASCII normalizes a domain name or URL to ASCII/punycode.
func ToASCII(s string) string {
	s, _ = idna.ToASCII(s)
	return s
}

// normalizeLang normalizes a 4-char ISO 15924 script code (e.g. “Latn”)
// or a BCP 47 language tag, e.g. (“en” or “en-us”).
func normalizeLang(s string) (string, error) {
	if s == "none" {
		s = "und"
	}
	tag, err := language.Parse(s)
	if err != nil {
		if len(s) != 4 {
			return "", err
		}
		tag, err = language.Parse("mul-" + s) // Try parsing as a script, e.g. “Latn”
	}
	return tag.String(), err
}
