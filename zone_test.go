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
		{"us", []string{"en-US"}},
		{"中国", []string{"zh-CN", "zh-Hans-CN"}},
		{"中國", []string{"zh-CN", "zh-Hant-CN"}},
		{"台湾", []string{"hak-Hant-TW", "nan-Hant-TW", "zh-Hans-TW", "zh-Hant-TW"}},
		{"台灣", []string{"hak-Hant-TW", "nan-Hant-TW", "zh-Hant-TW"}},
		{"香港", []string{"en-HK", "zh-Hans-HK", "zh-Hant-HK"}},
	}

	for _, tt := range tests {
		t.Run(tt.domain, func(t *testing.T) {
			domain := ToASCII(tt.domain)
			z := ZoneMap[domain]
			got := z.languages
			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("Zone.Language(), got: %#v, want: %#v", got, tt.want)
			}
		})
	}
}

func TestRDAPURLs(t *testing.T) {
	tests := []struct {
		domain string
		want   []string
	}{
		{"nr", nil},
		{"uk", []string{"https://rdap.nominet.uk/uk/"}},
		{"co.uk", []string{"https://rdap.nominet.uk/uk/"}},
		{"ac.uk", []string{"https://rdap.nominet.uk/uk/"}},
		{"bbc", []string{"https://rdap.nominet.uk/bbc/"}},
		{"lol", []string{"https://rdap.centralnic.com/lol/"}},
	}

	for _, tt := range tests {
		t.Run(tt.domain, func(t *testing.T) {
			domain := ToASCII(tt.domain)
			z := ZoneMap[domain]
			got := z.RDAPURLs()
			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("Zone.RDAPURLs(), got: %#v, want: %#v", got, tt.want)
			}
		})
	}
}
