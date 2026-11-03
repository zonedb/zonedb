package build

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"strings"
	"sync/atomic"

	"github.com/PuerkitoBio/goquery"
)

var (
	ianaTablesURL = "https://www.iana.org/domains/idn-tables"
	ianaBaseURL   = "https://www.iana.org"
)

var ianaTypos = map[string]string{
	"shiksa": "shiksha",
}

// FetchIDNTablesFromIANA fetches IDN table references from the IANA website.
func FetchIDNTablesFromIANA(ctx context.Context, zones map[string]*Zone, cache *ETagCache) error {
	tlds := TLDs(zones)
	baseURL, err := url.Parse(ianaBaseURL)
	if err != nil {
		return err
	}
	res, err := FetchWithETag(ctx, ianaTablesURL, cache)
	if err != nil {
		return err
	}
	if res == nil {
		return nil // 304 Not Modified
	}
	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return err
	}

	// Clear IANA-sourced IDN table policies so that the IANA page is
	// treated as the authoritative snapshot. Stale entries that IANA no
	// longer lists will be removed. Languages from other sources (CLDR,
	// manual, etc.) are preserved.
	ianaPrefix := ianaBaseURL + "/"
	for _, z := range zones {
		z.PruneStalePolicies(func(p Policy) bool {
			return p.Type == TypeIDNTable &&
				(strings.HasPrefix(p.Source, ianaPrefix) || (p.Source == "" && strings.HasPrefix(p.Value, ianaPrefix)))
		})
	}

	var (
		matchCount     uint64
		extractedCount uint64
	)
	doc.Find("body table#idn-table > tbody > tr").Each(func(i int, tr *goquery.Selection) {
		// Group header rows carry no table link; skip them before counting.
		if tr.HasClass("idn-group-row") {
			return
		}
		s := tr.Find("td").First()
		if s.Length() == 0 {
			return
		}
		atomic.AddUint64(&matchCount, 1)
		domain := domainFromRow(tr, s)
		forLabel := s.Find("a").Text()
		if forLabel == "" {
			forLabel = domain
		}
		partURL, exists := s.Find("a[href]").Attr("href")
		if !exists {
			Trace("@{r}missing href for %q\n", forLabel)
			return
		}
		u, err := baseURL.Parse(partURL)
		if err != nil {
			Trace("@{r}failed to parse %q for %q\n", partURL, forLabel)
			return
		}

		// At this point, "domain" looks like ".<tld>" and u should have u.String() which is an absolute working URL
		// The partURL's last component, after directory-separator, looks like "<tld>_<language>_<version.info>.txt"
		if domain == "" || domain[0] != '.' {
			Trace("@{r}bad domain name %q\n", domain)
			return
		}
		domain = domain[1:] // trim dot

		// Check IANA typos (currently shiksa → shiksha)
		if fixed, ok := ianaTypos[domain]; ok {
			domain = fixed
		}

		z, ok := zones[domain]
		if !ok {
			if len(tlds) > 100 {
				Trace("@{r}unknown zone %q from %s\n", domain, ianaBaseURL+partURL)
			}
			return
		}

		lang, err := langFromURL(partURL)
		if err != nil {
			Trace("@{y}unable to extract language tag or script for zone %q from %s\n", domain, partURL)
			return
		}

		z.AddPolicy(TypeIDNTable, lang, u.String(), ianaTablesURL, "")
		z.Languages = append(z.Languages, lang)

		atomic.AddUint64(&extractedCount, 1)
	})
	Trace("@{.}saw %d matches, extracted %d entries\n", matchCount, extractedCount)
	if len(tlds) > 100 {
		if extractedCount == 0 {
			return errors.New("failed to extract any URLs from IANA index page, HTML change?")
		}
		// Stale IANA policies are pruned above, before parsing. A partial
		// HTML change that still yields some rows would silently delete the
		// rest, so treat a large drop as an error rather than committing the
		// loss. The 2026-08 regrouping broke every row at once; a narrower
		// one (e.g. a single column moving) would not.
		if matchCount > 0 && extractedCount*2 < matchCount {
			return fmt.Errorf("extracted only %d of %d IANA IDN table rows, HTML change?", extractedCount, matchCount)
		}
	}

	return nil
}

// domainFromRow extracts the ".<tld>" label for an IDN table row.
//
// IANA's current layout groups rows by zone and carries the domain on the
// row (data-idn-domain) or cell (data-sort-value), with the <span> holding
// the language name. The older layout put the domain in the <span>. Both
// are handled so a future revert or cached copy still parses.
func domainFromRow(tr, td *goquery.Selection) string {
	if d, ok := tr.Attr("data-idn-domain"); ok && d != "" {
		return strings.TrimSpace(d)
	}
	if d, ok := td.Attr("data-sort-value"); ok && d != "" {
		return strings.TrimSpace(d)
	}
	return strings.TrimSpace(td.Find("span").First().Text())
}

func langFromURL(u string) (string, error) {
	i := strings.LastIndexByte(u, '/')
	if i == -1 {
		return "", errMalformedURL
	}
	sections := strings.Split(u[i+1:], "_")
	if len(sections) < 3 {
		return "", errMalformedURL
	}
	// Checking sections[0] against zone from caller doesn’t work because:
	//  1. IDN TLDs using other strings
	//  2. table sharing between zones (eg, “academy” appears as baseline for many)
	// This does tell us that we want to have a cache of values on a per-URL basis, to avoid fetching the same
	// URL N times.
	lang := sections[1]
	return normalizeLang(lang)
}

var errMalformedURL = errors.New("malformed IDN table URL")
