package zonedb

import (
	"reflect"
	"testing"
	"unsafe"

	"golang.org/x/net/idna"
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

func TestLanguages(t *testing.T) {
	tests := []struct {
		domain string
		want   []string
	}{
		{"com", nil},
		{"中国", []string{"zh-Hans-CN"}},
		{"中國", []string{"zh-Hant-CN"}},
		{"台灣", []string{"zh-Hant-TW"}},
		{"台灣", []string{"zh-Hant-TW"}},
		{"香港", []string{"zh-Hans-HK"}},
	}

	for _, tt := range tests {
		t.Run(tt.domain, func(t *testing.T) {
			domain := ToASCII(tt.domain)
			z := ZoneMap[domain]
			got := z.languages
			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("Zone.Language(), got: %v, want: %v", got, tt.want)
			}
		})
	}
}
