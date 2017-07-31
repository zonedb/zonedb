package build

import (
	"errors"
	"net/url"
	"os"
	"strings"
	"sync/atomic"

	"github.com/PuerkitoBio/goquery"
)

const (
	// I override the URL during dev.
	prodIANATablesIndexURL    = "https://www.iana.org/domains/idn-tables"
	tablesIndexURLOverrideEnv = "IANA_IDN_TABLES_URL"
	prodIANABaseURL           = "https://www.iana.org"
)

// FetchIDNURLs parses one IANA web-page to figure out which character sets
// each zone supports and the canonical URL for each.
func FetchIDNURLs(zones map[string]*Zone) error {
	baseURL, err := url.Parse(prodIANABaseURL)
	if err != nil {
		panic(err)
	}

	ianaFetchURL := os.Getenv(tablesIndexURLOverrideEnv)
	if ianaFetchURL == "" {
		ianaFetchURL = prodIANATablesIndexURL
	}
	res, err := Fetch(ianaFetchURL)
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
		zone := s.Find("span").Text()
		forLabel := s.Find("a").Text()
		if forLabel == "" {
			forLabel = zone
		}
		partURL, exists := s.Find("a[href]").Attr("href")
		if !exists {
			forLabel := s.Find("a").Text()
			if forLabel == "" {
				forLabel = zone
			}
			Trace("@{r}FetchIDNLists: missing href for %q\n", forLabel)
			return
		}
		u, err := baseURL.Parse(partURL)
		if err != nil {
			Trace("@{r}FetchIDNLists: failed to parse %q for %q\n", partURL, forLabel)
			return
		}

		if strings.HasSuffix(partURL, ".xml") {
			Trace("@{y}FetchIDNLists: unhandled XML URL for %q: %q\n", forLabel, u.String())
		}

		// At this point, "zone" looks like ".<tld>" and u should have u.String() which is an absolute working URL
		// The partURL's last component, after directory-separator, looks like "<tld>_<languagetag>_<version.info>.txt"
		if zone[0] != '.' {
			Trace("@{r}FetchIDNLists: bad zone %q\n", zone)
			return
		}
		zone = zone[1:]
		if _, have := zones[zone]; !have {
			Trace("@{r}FetchIDNLists: unrecognized zone %q\n", zone)
			return
		}
		if zones[zone].IDNTableURLs == nil {
			zones[zone].IDNTableURLs = make(map[string]string)
		}

		language := languageFromURL(partURL)
		if language == "" {
			Trace("@{r}FetchIDNLists: for zone %q unable to extract language from %q\n", zone, partURL)
			return
		}

		zones[zone].IDNTableURLs[language] = u.String()
		atomic.AddUint64(&extractedCount, 1)
	})
	Trace("@{.}FetchIDNLists: saw %d matches, extracted %d entries\n", matchCount, extractedCount)
	if extractedCount == 0 {
		return errors.New("failed to extract any URLs from IANA index page, HTML change?")
	}

	// TODO: Do we want to _remove_ URLs if zone not found here?
	//       How do we handle multiple sources of URLs if so?

	return nil
}

func languageFromURL(u string) string {
	i := strings.LastIndexByte(u, '/')
	if i == -1 {
		return ""
	}
	sections := strings.Split(u[i+1:], "_")
	if len(sections) < 3 {
		return ""
	}
	// Checking sections[0] against zone from caller doesn't work because:
	//  1. IDN TLDs using other strings
	//  2. table sharing between zones (eg, 'academy' appears as baseline for many)
	// This does tell us that we want to have a cache of values on a per-URL basis, to avoid fetching the same
	// URL N times.
	return sections[1]
}
