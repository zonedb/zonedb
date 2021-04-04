package build

import (
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/wsxiaoys/terminal/color"
)

func UpdateInfoURLs(zones map[string]*Zone) {
	color.Fprintf(os.Stderr, "@{.}Updating info URLs for %d zones...\n", len(zones))

	client := &http.Client{
		Timeout: 2 * time.Second,
	}

	mapZones(zones, func(z *Zone) {
		if z.InfoURL == "" {
			return
		}
		urls := []string{
			z.InfoURL,
		}
		if strings.HasPrefix(z.InfoURL, "http:") {
			urls = []string{
				strings.Replace(z.InfoURL, "http:", "https:", 1),
				z.InfoURL,
			}
		}
		for _, u := range urls {
			res, err := client.Get(u)
			if err != nil {
				// color.Fprintf(os.Stderr, "@{y!}Warning:@{y} error fetching info URL for @{y!}%s@{y}: (%s): %v\n", z.Domain, u, err)
				continue
			}
			ru := NormalizeURL(res.Request.URL.String())
			if ru != z.InfoURL {
				z.InfoURL = ru
				color.Fprintf(os.Stderr, "@{.}Updated info URLs for @{c}%s@{c}: @{y}%s\n", z.Domain, z.InfoURL)
			}
			break
		}
	})
}
