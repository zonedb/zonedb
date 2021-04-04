package build

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/wsxiaoys/terminal/color"
)

func UpdateInfoURLs(zones map[string]*Zone) {
	color.Fprintf(os.Stderr, "@{.}Updating info URLs for %d zones...\n", len(zones))

	transport := http.DefaultTransport.(*http.Transport).Clone()
	transport.TLSHandshakeTimeout = 2 * time.Second
	transport.MaxIdleConnsPerHost = 10
	client := &http.Client{
		Transport: transport,
		Timeout:   4 * time.Second,
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
			CloseN(res.Body, 10_000_000)
			break
		}
	})
}

// CloseN drains rc up to a maximum of n bytes and closes rc.
// It returns the number of bytes drained and the first error encountered.
// Regardless of any errors, rc.Close() is guaranteed to be called.
func CloseN(rc io.ReadCloser, n int64) (int64, error) {
	n, cerr := io.Copy(ioutil.Discard, io.LimitReader(rc, n))
	err := rc.Close()
	if cerr != nil {
		err = cerr
	}
	return n, err
}
