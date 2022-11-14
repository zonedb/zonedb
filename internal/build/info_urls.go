package build

import (
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func UpdateInfoURLs(zones map[string]*Zone) {
	Trace("@{.}Updating info URLs for %d zones...\n", len(zones))

	transport := http.DefaultTransport.(*http.Transport).Clone()
	transport.TLSHandshakeTimeout = 5 * time.Second
	transport.MaxIdleConnsPerHost = 10
	client := &http.Client{
		Transport: transport,
		Timeout:   10 * time.Second,
	}

	mapZones(zones, func(z *Zone) {
		var urls []string

		if strings.HasPrefix(z.InfoURL, "http:") {
			urls = []string{
				strings.Replace(z.InfoURL, "http:", "https:", 1),
				z.InfoURL,
			}
		} else if !strings.HasPrefix(z.InfoURL, "https://newgtlds.icann.org") {
			urls = []string{
				z.InfoURL,
			}
		}

		if z.IsTLD() {
			urls = append(urls,
				// Try NIC websites
				"https://nic."+z.Domain,
				"https://www.nic."+z.Domain,
				"http://nic."+z.Domain,
				"http://www.nic."+z.Domain,

				// Try ICANN first
				"https://www.icann.org/en/registry-agreements/details/"+z.ASCII(),

				// Then fall back to IANA
				"https://www.iana.org/domains/root/db/"+z.ASCII()+".html",
			)
		}

		var infoURL string
		for _, u := range urls {
			if u == "" {
				continue
			}
			res, err := client.Get(u)
			if err != nil {
				if u == z.InfoURL {
					Trace("@{y!}Warning:@{y} error fetching info URL for @{y!}%s@{y}: (%s): %v\n", z.Domain, u, err)
				}
				continue
			} else if res.StatusCode != http.StatusOK {
				if u == z.InfoURL {
					Trace("@{y!}Warning:@{y} non-200 status for info URL for @{y!}%s@{y}: (%s): %s\n", z.Domain, u, res.Status)
				}
				continue
			}
			CloseN(res.Body, 10_000_000)
			infoURL = NormalizeURL(res.Request.URL.String())
			break
		}

		// Do not rewrite URLs that just add tracking or query string info
		if infoURL != z.InfoURL && !strings.HasPrefix(infoURL, z.InfoURL) {
			if infoURL == "" {
				Trace("@{.!}Removed@{.} info URL for @{c}%s@{c}: @{y}%s@{c}\n", z.Domain, z.InfoURL)

			} else {
				Trace("@{.}Updated info URL for @{c}%s@{c}: @{y}%s@{c} → @{g}%s\n", z.Domain, z.InfoURL, infoURL)
			}
			z.InfoURL = infoURL
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
