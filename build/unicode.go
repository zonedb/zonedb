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
	if ct == nil || *ct == nil {
		return
	}
	ct.Sort()
	j := 0
	for i := 1; i < len(*ct); i++ {
		left, right := &(*ct)[j], &(*ct)[i]
		if right.Lo-left.Hi <= 1 {
			left.Hi = right.Hi
		} else {
			j++
		}
	}
	*ct = (*ct)[0 : j+1]
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
