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
	transport.TLSHandshakeTimeout = 5 * time.Second
	transport.MaxIdleConnsPerHost = 10
	client := &http.Client{
		Transport: transport,
		Timeout:   10 * time.Second,
	}

	mapZones(zones, func(z *Zone) {
		if z.InfoURL == "" {
			return
		}
		var urls []string
		if strings.HasPrefix(z.InfoURL, "https://newgtlds.icann.org") {
			if len(z.NameServers) == 0 {
				color.Fprintf(os.Stderr, "@{.}Removed info URL for @{c}%s@{c} (no DNS entries): @{y}%s@{c}\n", z.Domain, z.InfoURL)
				z.InfoURL = ""
				return
			}
			urls = []string{
				"https://nic." + z.Domain,
				"http://nic." + z.Domain,
				z.InfoURL,
			}
		} else if strings.HasPrefix(z.InfoURL, "http:") {
			urls = []string{
				strings.Replace(z.InfoURL, "http:", "https:", 1),
				z.InfoURL,
			}
		} else {
			urls = []string{
				z.InfoURL,
			}
		}
		for _, u := range urls {
			res, err := client.Get(u)
			if err != nil {
				color.Fprintf(os.Stderr, "@{y!}Warning:@{y} error fetching info URL for @{y!}%s@{y}: (%s): %v\n", z.Domain, u, err)
				continue
			}
			CloseN(res.Body, 10_000_000)
			ru := NormalizeURL(res.Request.URL.String())
			if ru != z.InfoURL && !strings.HasPrefix(ru, "https://newgtlds.icann.org") {
				color.Fprintf(os.Stderr, "@{.}Updated info URL for @{c}%s@{c}: @{y}%s@{c} → @{g}%s\n", z.Domain, z.InfoURL, ru)
				z.InfoURL = ru
			}
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

// trimURL trims query strings and /index.htm(l)? from a URL.
func trimURL(u string) string {
	u, _, _ = strings.Cut(u, "?")
	if strings.HasSuffix(u, "/index.html") {
		u = strings.TrimSuffix(u, "/index.html")
	} else {
		u = strings.TrimSuffix(u, "/index.htm")
	}
	return u
}
