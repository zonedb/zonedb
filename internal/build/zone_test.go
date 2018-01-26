package build

import "testing"

func TestZone_HasMetadata(t *testing.T) {
	z := &Zone{Domain: "example.com"}
	if z.HasMetadata() != false {
		t.Errorf("expected z.Metadata() to return false, got true")
	}
	z = &Zone{Domain: "google.com", InfoURL: "https://www.google.com/"}
	if z.HasMetadata() != true {
		t.Errorf("expected z.Metadata() to return true, got false")
	}
}
