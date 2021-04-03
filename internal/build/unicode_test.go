package build

import "testing"

func TestNormalize(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{"com", "com"},
		{"net", "net"},
		{"org", "org"},
		{"COM", "com"},
		{"nEt", "net"},
		{".ORg", "org"},
		{"--HEl01-", "hel01"},
		{"--SMOOTH.OPERATOR--", "smooth.operator"},
		{".МОСКВА", "москва"},
		{"通用电气公司", "通用电气公司"},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			got := Normalize(tt.s)
			if got != tt.want {
				t.Errorf("Normalize(%q), want: %v, got: %v", tt.s, tt.want, got)
			}
		})
	}
}
func TestToASCII(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{"com", "com"},
		{"net", "net"},
		{"org", "org"},
		{"москва", "xn--80adxhks"},
		{"通用电气公司", "xn--55qx5d8y0buji4b870u"},
		{"abc.com:80", "abc.com:80"},
		{"abc.com:443", "abc.com:443"},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			got := ToASCII(tt.s)
			if got != tt.want {
				t.Errorf("ToASCII(%q), want: %v, got: %v", tt.s, tt.want, got)
			}
		})
	}
}

func TestNormalizeURL(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{"http://www.google.com/", "http://www.google.com/"},
		{"https://www.google.com/", "https://www.google.com/"},
		{"https://WWW.ALPHA.BET/", "https://www.alpha.bet/"},
		{"https://nic.dev", "https://nic.dev/"},
		{"https://nic.москва/", "https://nic.москва/"},
		{"https://nic.xn--80adxhks/", "https://nic.москва/"},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			got := NormalizeURL(tt.s)
			if got != tt.want {
				t.Errorf("NormalizeURL(%q), want: %v, got: %v", tt.s, tt.want, got)
			}
		})
	}
}

func TestToASCIIURL(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{"http://www.google.com/", "http://www.google.com/"},
		{"https://www.google.com/", "https://www.google.com/"},
		{"https://nic.москва/", "https://nic.xn--80adxhks/"},
		{"https://nic.xn--80adxhks/", "https://nic.xn--80adxhks/"},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			got := ToASCIIURL(tt.s)
			if got != tt.want {
				t.Errorf("ToASCIIURL(%q), want: %v, got: %v", tt.s, tt.want, got)
			}
		})
	}
}
