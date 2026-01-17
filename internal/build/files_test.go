package build

import (
	"strings"
	"testing"
)

func TestZonesFileValid(t *testing.T) {
	// This test validates that zones.txt contains only valid domain names.
	// It will fail if any invalid zones are present (e.g., " (deprecated)").

	// Set BaseDir to repository root (two directories up from internal/build)
	oldBaseDir := BaseDir
	BaseDir = "../.."
	defer func() { BaseDir = oldBaseDir }()

	zones, errs := ReadZonesFile()
	if len(errs) > 0 {
		for _, err := range errs {
			t.Error(err)
		}
	}

	// Additional validation: check for invalid characters that Normalize() doesn't catch
	for domain := range zones {
		if strings.ContainsAny(domain, " ()") {
			t.Errorf("invalid domain name contains spaces or parentheses: %q", domain)
		}
		if strings.HasPrefix(domain, " ") {
			t.Errorf("invalid domain name has leading space: %q", domain)
		}
	}
}
