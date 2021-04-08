package build

import (
	"net"
	"net/url"
	"sort"
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

// ToASCII normalizes a domain name to ASCII/punycode.
func ToASCII(s string) string {
	s, _ = idna.ToASCII(s)
	return s
}

const hostnameProxy = "_____HOSTNAME_____"

// NormalizeURL normalizes the domain name part of a URL, with a Unicode (non-IDNA encoded) hostname.
// Returns a blank string if it is unable to parse s as a valid URL.
func NormalizeURL(s string) string {
	if s == "" {
		return s
	}
	u, _ := url.Parse(s)
	if u == nil {
		return ""
	}
	h, p, err := net.SplitHostPort(u.Host)
	if err != nil {
		h = u.Hostname()
	}
	if (u.Scheme == "http" && p == "80") || (u.Scheme == "https" && p == "443") {
		p = ""
	}
	h = Normalize(h)
	if p != "" {
		h = net.JoinHostPort(h, p)
	}
	u.Host = hostnameProxy
	if (u.Scheme == "http" || u.Scheme == "https") && u.Path == "" {
		u.Path = "/"
	}
	return strings.Replace(u.String(), hostnameProxy, h, 1)
}

// ToASCIIURL normalizes the domain name part of a URL to ASCII/punycode.
func ToASCIIURL(s string) string {
	if s == "" {
		return s
	}
	u, _ := url.Parse(s)
	if u == nil {
		return ""
	}
	u.Host = ToASCII(u.Host)
	return u.String()
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

// NormalizeDomains normalizes a slice of domain names.
// Returns a sorted, de-duplicated, normalized list of domain names.
func NormalizeDomains(in []string) []string {
	out := in[:]
	for i := range out {
		out[i] = Normalize(out[i])
	}
	out = NewSet(out...).Values()
	sort.Strings(out)
	return out
}

// IndexOfString finds s as a substring in haystack, returning the index and offset of s in haystack.
// Returns -1, -1 if s is not present in haystack.
func IndexOfString(haystack []string, s string) (int, int) {
	for i := range haystack {
		j := strings.Index(haystack[i], s)
		if j < 0 {
			continue
		}
		return i, j
	}
	return -1, -1
}

// IndexOrAppendString finds or adds s to haystack, returning the index and offset of s in haystack.
// Returns 0, 0, 0 for an empty string.
func IndexOrAppendString(haystack *[]string, s string) (int, int, int) {
	if len(s) == 0 {
		return 0, 0, 0
	}
	i, j := IndexOfString(*haystack, s)
	if i < 0 {
		i = len(*haystack)
		*haystack = append(*haystack, s)
	}
	return i, j, j + len(s)
}

// IndexOfStrings finds a rank-sorted slice of domain names (needle)
// within a larger slice (haystack).
func IndexOfStrings(haystack []string, needle []string) int {
outer:
	for i := range haystack {
		for j := range needle {
			if i+j >= len(haystack) || needle[j] != haystack[i+j] {
				continue outer
			}
		}
		return i
	}
	return -1
}

// IndexOrAppendStrings finds or appends a slice of rank-sorted domain names (needle)
// Returns 0,0 for a zero-length needle.
func IndexOrAppendStrings(haystack *[]string, needle []string) (int, int) {
	if len(needle) == 0 {
		return 0, 0
	}
	idx := IndexOfStrings(*haystack, needle)
	if idx < 0 {
		idx = len(*haystack)
		*haystack = append(*haystack, needle...)
	}
	return idx, idx + len(needle)
}
