package build

import (
	"bytes"
	"errors"
	"sort"
	"strings"

	"golang.org/x/net/idna"
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

// ASCII is a Unicode CodeTable for ASCII domain names (a-z, 0-9, and -)
var ASCII, _ = NewCodeTable(`--09az`)

// CodeTable is a slice of Unicode code point ranges.
type CodeTable []CodeRange

// NewCodeTable returns a valid, sorted, compressed CodeTable from the input string.
func NewCodeTable(s string) (CodeTable, error) {
	var ct CodeTable
	if err := ct.UnmarshalText([]byte(s)); err != nil {
		return nil, err
	}
	if err := ct.Validate(); err != nil {
		return ct, err
	}
	ct.Compress()
	return ct, nil
}

var (
	OddRuneCount = errors.New("odd number of runes")
	InvalidRange = errors.New("invalid Unicode range")
)

// Validate determines if a CodeTable is valid, returning an error if invalid.
func (ct CodeTable) Validate() error {
	for _, c := range ct {
		if c.Hi < c.Lo {
			return InvalidRange
		}
	}
	return nil
}

// Compress sorts and compresses a CodeTable, joining or eliminating overlapping CodeRanges.
func (ct *CodeTable) Compress() {
	if ct == nil || *ct == nil || len(*ct) < 2 {
		return
	}
	ct.Sort()
	newCt := &CodeTable{}
	j := 0
	l := len(*ct) - 1
	for i := 1; i <= l; i++ {
		left, right := &(*ct)[j], &(*ct)[i]
		if right.Lo-left.Hi <= 1 {
			if right.Hi > left.Hi {
				left.Hi = right.Hi
			}
			if i == l {
				*newCt = append(*newCt, *left)
			}
		} else {
			*newCt = append(*newCt, *left)
			if i == l {
				*newCt = append(*newCt, *right)
			}
			j = i
		}
	}
	*ct = *newCt
}

func (ct *CodeTable) UnmarshalText(b []byte) error {
	runes := bytes.Runes(b)
	if len(runes)%2 != 0 {
		return OddRuneCount
	}
	*ct = make(CodeTable, len(runes)/2)
	for i := 0; i < len(*ct); i++ {
		(*ct)[i].Lo = runes[i*2]
		(*ct)[i].Hi = runes[i*2+1]
	}
	return ct.Validate()
}

func (ct CodeTable) MarshalText() ([]byte, error) {
	var buf bytes.Buffer
	for _, c := range ct {
		_, err := buf.WriteRune(c.Lo)
		if err != nil {
			return nil, err
		}
		_, err = buf.WriteRune(c.Hi)
		if err != nil {
			return nil, err
		}
	}
	return buf.Bytes(), nil
}

func (ct CodeTable) Runes() []rune {
	var runes []rune
	for _, cr := range ct {
		runes = append(runes, cr.Lo)
		runes = append(runes, cr.Hi)
	}
	return runes
}

// Adds a rune, extending the last CodeRange, if possible
func (ct *CodeTable) AppendRune(char rune) {
	tableLen := len(*ct)
	if tableLen > 0 {
		if (*ct)[tableLen-1].Hi == char-1 {
			(*ct)[tableLen-1].Hi = char
			return
		}
	}
	*ct = append(*ct, CodeRange{char, char})
}

// Adds a range of runes
func (ct *CodeTable) AppendRange(start, end rune) {
	*ct = append(*ct, CodeRange{start, end})
}

// Appends another complete CodeTable
func (ct *CodeTable) AppendTable(other *CodeTable) {
	*ct = append(*ct, *other...)
}

func (ct CodeTable) Sort()         { sort.Sort(ct) }
func (ct CodeTable) Len() int      { return len(ct) }
func (ct CodeTable) Swap(i, j int) { ct[i], ct[j] = ct[j], ct[i] }
func (ct CodeTable) Less(i, j int) bool {
	a, b := ct[i], ct[j]
	switch {
	case a.Lo != b.Lo:
		return a.Lo < b.Lo
	case a.Hi != b.Hi:
		return a.Hi < b.Hi
	}
	return false
}

type CodeRange struct {
	Lo rune
	Hi rune
}

// IndexOfRunes finds a slice of runes (needle)
// within a larger slice (haystack).
func IndexOfRunes(haystack []rune, needle []rune) int {
outer:
	for i := range haystack {
		for j := range needle {
			if i+j >= len(haystack) {
				return -1
			}
			if needle[j] != haystack[i+j] {
				continue outer
			}
		}
		return i
	}
	return -1
}

// IndexOrAppendRunes finds or appends a slice of runes (needle)
// Returns 0,0 for a zero-length needle.
func IndexOrAppendRunes(haystack *[]rune, needle []rune) (int, int) {
	if len(needle) == 0 {
		return 0, 0
	}
	idx := IndexOfRunes(*haystack, needle)
	if idx < 0 {
		idx = len(*haystack)
		*haystack = append(*haystack, needle...)
	}
	return idx, idx + len(needle)
}

// Dup returns a copy of a CodeTable, hiding implementation details; it must
// recurse as much as needed such that the original and the duplicate can be
// independently modified.
func (ct *CodeTable) Dup() *CodeTable {
	d := make(CodeTable, len(*ct))
	copy(d, *ct)
	return &d
}
