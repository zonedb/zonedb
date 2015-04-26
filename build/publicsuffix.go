package build

import (
	"os"
	"strings"

	"github.com/wsxiaoys/terminal/color"
	"golang.org/x/net/idna"
	"golang.org/x/net/publicsuffix"
)

const pfx = "zonedb-test."

func CheckPublicSuffix(zones map[string]*Zone) {
	color.Fprintf(os.Stderr, "@{.}Checking against the Public Suffix List for %d zones...\n", len(zones))
	mapZones(zones, func(z *Zone) {
		host, err := idna.ToASCII(pfx + z.Domain)
		if err != nil {
			LogWarning(err)
			return
		}
		s, _ := publicsuffix.PublicSuffix(host)
		s = Normalize(s)
		switch {
		// ZoneDB and PSL agree
		case s == z.Domain:
			return

		// PSL wildcard
		case strings.HasPrefix(s, pfx) && len(z.Subdomains) != 0:
			return

		// ZoneDB and PSL disagree
		default:
			color.Fprintf(os.Stderr, "@{y}Public Suffix List: @{y!}%s@{y} for @{y!}%s\n", s, z.Domain)
		}
	})
}
