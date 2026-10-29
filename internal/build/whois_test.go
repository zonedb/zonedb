package build

import "testing"

// TestZone_Normalize_WhoisExclusive covers the #812 invariant: a zone has a
// whoisServer or a whoisURL, never both. The nightly update broke on this
// when IANA began publishing a reachable port-43 server for .sr and .jo,
// which were carrying a curated whoisURL.
//
// Normalize enforces it because it runs after VerifyWhois has cleared any
// unreachable server, so a zone is never left with neither.
func TestZone_Normalize_WhoisExclusive(t *testing.T) {
	tests := []struct {
		name       string
		zone       *Zone
		wantServer string
		wantURL    string
	}{
		{
			name:       "verified server clears curated URL",
			zone:       &Zone{Domain: "sr", WhoisServer: "whois.sr", WhoisURL: "http://www.register.sr/"},
			wantServer: "whois.sr",
		},
		{
			name:       "server only is untouched",
			zone:       &Zone{Domain: "com", WhoisServer: "whois.verisign-grs.com"},
			wantServer: "whois.verisign-grs.com",
		},
		{
			// VerifyWhois cleared an unreachable server earlier in the
			// pipeline; the URL must survive as the only way to reach whois.
			name:    "URL survives when there is no server",
			zone:    &Zone{Domain: "bb", WhoisURL: "http://whois.telecoms.gov.bb/search_domain.php"},
			wantURL: "http://whois.telecoms.gov.bb/search_domain.php",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			z := tt.zone
			z.Normalize()
			if z.WhoisServer != tt.wantServer {
				t.Errorf("WhoisServer = %q, want %q", z.WhoisServer, tt.wantServer)
			}
			if z.WhoisURL != tt.wantURL {
				t.Errorf("WhoisURL = %q, want %q", z.WhoisURL, tt.wantURL)
			}
		})
	}
}
