package zonedb

import "testing"

func TestZonesHaveWhoisServerOrURLButNotBoth(t *testing.T) {
	for _, z := range Zones {
		if z.whoisServer != "" && z.whoisURL != "" {
			t.Errorf(`Zone %q has both whoisServer (%q) and whoisURL (%q) set`, z.Domain, z.whoisServer, z.whoisURL)
		}
	}
}
