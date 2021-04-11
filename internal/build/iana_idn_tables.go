package build

import (
	"errors"
	"net/url"
	"strings"
	"sync/atomic"

	"github.com/PuerkitoBio/goquery"
)

const (
	// I override the URL during dev.
	ianaTablesURL = "https://www.iana.org/domains/idn-tables"
	ianaBaseURL   = "https://www.iana.org"
)

// FetchIDNTablesFromIANA fetches IDN table references from the IANA website.
func FetchIDNTablesFromIANA(zones map[string]*Zone) error {
	tlds := TLDs(zones)
	baseURL, err := url.Parse(ianaBaseURL)
	if err != nil {
		return err
	}
	res, err := Fetch(ianaTablesURL)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return err
	}

	var (
		matchCount     uint64
		extractedCount uint64
	)
	doc.Find("body table#idn-table > tbody > tr > td:first-child").Each(func(i int, s *goquery.Selection) {
		atomic.AddUint64(&matchCount, 1)
		domain := s.Find("span").Text()
		forLabel := s.Find("a").Text()
		if forLabel == "" {
			forLabel = domain
		}
		partURL, exists := s.Find("a[href]").Attr("href")
		if !exists {
			forLabel := s.Find("a").Text()
			if forLabel == "" {
				forLabel = domain
			}
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
		if domain[0] != '.' {
			Trace("@{r}bad domain name %q\n", domain)
			return
		}
		domain = domain[1:] // trim dot
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

		z.AddPolicy(TypeIDNTable, lang, u.String(), "") // "Fetched from "+ianaTablesURL)

		atomic.AddUint64(&extractedCount, 1)
	})
	Trace("@{.}saw %d matches, extracted %d entries\n", matchCount, extractedCount)
	if extractedCount == 0 && len(tlds) > 100 {
		return errors.New("failed to extract any URLs from IANA index page, HTML change?")
	}

	// TODO: Do we want to _remove_ URLs if zone not found here?
	//       How do we handle multiple sources of URLs if so?

	return nil
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
