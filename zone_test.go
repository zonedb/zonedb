package zonedb

import (
	"testing"
	"unsafe"

	"golang.org/x/net/idna"
	"golang.org/x/text/language"
)

// ToASCII normalizes a domain name or URL to ASCII/punycode.
func ToASCII(s string) string {
	s, _ = idna.ToASCII(s)
	return s
}

func TestSizeofZone(t *testing.T) {
	var z Zone
	t.Logf("sizeof Zone = %d", unsafe.Sizeof(z))
}

func TestLanguage(t *testing.T) {
	tests := []struct {
		domain string
		want   string
	}{
		{"com", ""},
	}

	for _, tt := range tests {
		t.Run(tt.domain, func(t *testing.T) {
			domain := ToASCII(tt.domain)
			z := ZoneMap[domain]
			got := z.Language()
			want, _ := language.Parse(tt.want)
			if got != want {
				t.Errorf("Zone.Language(), got: %v, want: %v", got, want)
			}
		})
	}
}
