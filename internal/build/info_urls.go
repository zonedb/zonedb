package build

import (
	"context"
	"io"
	"net/http"
	"strings"
	"time"
)

// UpdateInfoURLs probes candidate info URLs for each zone and rewrites
// z.InfoURL to the first reachable one.
func UpdateInfoURLs(ctx context.Context, zones map[string]*Zone) error {
	Trace("@{.}Updating info URLs for %d zones...\n", len(zones))

	// Dedicated client: many short GETs per registry host; MaxIdleConnsPerHost
	// helps reuse connections across the candidate URLs tried per zone.
	transport := http.DefaultTransport.(*http.Transport).Clone()
	transport.TLSHandshakeTimeout = 5 * time.Second
	transport.MaxIdleConnsPerHost = 10
	client := &http.Client{
		Transport: transport,
		Timeout:   10 * time.Second,
	}

	return mapZones(ctx, zones, func(gctx context.Context, z *Zone) error {
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
			u = NormalizeURL(u)
			req, err := http.NewRequestWithContext(gctx, http.MethodGet, u, nil)
			if err != nil {
				Trace("@{y!}Warning:@{y} error fetching info URL for @{y!}%s@{y}: (%s): %v\n", z.Domain, u, err)
				continue
			}
			req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36")
			res, err := client.Do(req)
			if err != nil {
				// Abort only on group cancellation; per-request failures mean "try the next URL".
				if gctx.Err() != nil {
					return gctx.Err()
				}
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
			_, _ = CloseN(res.Body, 10_000_000)
			// Don’t use redirected URL, use the URL we crafted
			// infoURL = NormalizeURL(res.Request.URL.String())
			infoURL = u
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
		return nil
	})
}

// CloseN drains rc up to a maximum of n bytes and closes rc.
// It returns the number of bytes drained and the first error encountered.
// Regardless of any errors, rc.Close() is guaranteed to be called.
func CloseN(rc io.ReadCloser, n int64) (int64, error) {
	n, cerr := io.Copy(io.Discard, io.LimitReader(rc, n))
	err := rc.Close()
	if cerr != nil {
		err = cerr
	}
	return n, err
}
