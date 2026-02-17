package build

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/PuerkitoBio/goquery"
)

var (
	centralNicIndexURL = "https://centralnicregistry.com/idn-tables/"
	centralNicBaseURL  = "https://centralnicregistry.com"
)

type centralNicEntry struct {
	Mnemonic string
	Type     string // "Language" or "Script"
	Name     string
	Version  string
	URL      string // absolute URL to detail page
}

// compareVersions compares two version strings like "1.2" and "3.0" numerically.
// Returns -1 if a < b, 0 if a == b, 1 if a > b.
func compareVersions(a, b string) int {
	aParts := strings.SplitN(a, ".", 2)
	bParts := strings.SplitN(b, ".", 2)

	aMajor, _ := strconv.Atoi(aParts[0])
	bMajor, _ := strconv.Atoi(bParts[0])
	if aMajor != bMajor {
		if aMajor < bMajor {
			return -1
		}
		return 1
	}

	var aMinor, bMinor int
	if len(aParts) > 1 {
		aMinor, _ = strconv.Atoi(aParts[1])
	}
	if len(bParts) > 1 {
		bMinor, _ = strconv.Atoi(bParts[1])
	}
	if aMinor < bMinor {
		return -1
	}
	if aMinor > bMinor {
		return 1
	}
	return 0
}

// parseCentralNicIndex parses the CentralNic IDN tables index page and returns
// a map of mnemonic → entry, keeping only the highest version for each mnemonic.
func parseCentralNicIndex(doc *goquery.Document, baseURL *url.URL) map[string]centralNicEntry {
	entries := make(map[string]centralNicEntry)
	doc.Find("table.table-data tbody tr").Each(func(i int, s *goquery.Selection) {
		mnemonic := strings.TrimSpace(s.Find("td.td-mnemonic").Text())
		typ := strings.TrimSpace(s.Find("td.td-type").Text())
		version := strings.TrimSpace(s.Find("td.td-version").Text())

		nameLink := s.Find("td.td-name a")
		name := strings.TrimSpace(nameLink.Text())
		href, exists := nameLink.Attr("href")
		if !exists || mnemonic == "" {
			return
		}

		u, err := baseURL.Parse(href)
		if err != nil {
			Trace("@{y}CentralNic: bad href %q for %q\n", href, mnemonic)
			return
		}
		// Ensure trailing slash (CentralNic redirects without it)
		absURL := u.String()
		if !strings.HasSuffix(absURL, "/") {
			absURL += "/"
		}

		entry := centralNicEntry{
			Mnemonic: mnemonic,
			Type:     typ,
			Name:     name,
			Version:  version,
			URL:      absURL,
		}

		if existing, ok := entries[mnemonic]; ok {
			if compareVersions(version, existing.Version) <= 0 {
				return // existing version is same or higher
			}
		}
		entries[mnemonic] = entry
	})
	return entries
}

// parseCentralNicDetailPage parses a CentralNic IDN table detail page and
// returns the list of zone names that support this table.
func parseCentralNicDetailPage(doc *goquery.Document) []string {
	var zones []string
	doc.Find("ul.tags li a").Each(func(i int, s *goquery.Selection) {
		zone := strings.TrimSpace(s.Text())
		if zone != "" {
			zones = append(zones, zone)
		}
	})
	return zones
}

// FetchIDNTablesFromCentralNic fetches IDN table data from CentralNic and
// applies it to zones. IANA-sourced policies take precedence and are not
// overwritten.
func FetchIDNTablesFromCentralNic(zones map[string]*Zone, cache *ETagCache) error {
	// Phase 0: Fetch index with ETag check
	res, err := FetchWithETag(centralNicIndexURL, cache)
	if err != nil {
		return err
	}
	if res == nil {
		Trace("@{.}CentralNic index unchanged, skipping\n")
		return nil // 304 Not Modified
	}
	defer res.Body.Close()

	baseURL, err := url.Parse(centralNicBaseURL)
	if err != nil {
		return err
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return err
	}

	// Phase 1: Clear stale CentralNic entries
	centralNicPrefix := centralNicBaseURL + "/"
	for _, z := range zones {
		var filtered []Policy
		var nonCNLangs []string
		for _, p := range z.Policies {
			if p.Type == TypeIDNTable && strings.HasPrefix(p.Source, centralNicPrefix) {
				continue
			}
			filtered = append(filtered, p)
			if p.Type == TypeIDNTable && p.Key != "" {
				nonCNLangs = append(nonCNLangs, p.Key)
			}
		}
		if len(filtered) != len(z.Policies) {
			z.Policies = filtered
			z.Languages = nonCNLangs
		}
	}

	// Phase 2: Parse index, fetch detail pages
	entries := parseCentralNicIndex(doc, baseURL)
	Trace("@{.}CentralNic: parsed %d entries (version-deduped) from index\n", len(entries))

	// Normalize mnemonics and collect detail page fetches
	type fetchJob struct {
		lang     string
		entry    centralNicEntry
		detailID string // URL path suffix for the detail page
	}
	var jobs []fetchJob
	for _, entry := range entries {
		lang, err := normalizeLang(entry.Mnemonic)
		if err != nil {
			Trace("@{y}CentralNic: skipping mnemonic %q (%s): %v\n", entry.Mnemonic, entry.Name, err)
			continue
		}
		jobs = append(jobs, fetchJob{lang: lang, entry: entry})
	}

	// Concurrent fetch of detail pages
	type detailResult struct {
		lang  string
		entry centralNicEntry
		zones []string
	}
	var (
		results []detailResult
		mu      sync.Mutex
		wg      sync.WaitGroup
		fetched uint64
		skipped uint64
	)
	limiter := make(chan struct{}, Concurrency)

	for _, job := range jobs {
		limiter <- struct{}{}
		wg.Add(1)
		go func(j fetchJob) {
			defer func() {
				<-limiter
				wg.Done()
			}()

			detailRes, err := Fetch(j.entry.URL)
			if err != nil {
				Trace("@{r}CentralNic: error fetching %s: %v\n", j.entry.URL, err)
				atomic.AddUint64(&skipped, 1)
				return
			}
			defer detailRes.Body.Close()

			detailDoc, err := goquery.NewDocumentFromReader(detailRes.Body)
			if err != nil {
				Trace("@{r}CentralNic: error parsing %s: %v\n", j.entry.URL, err)
				atomic.AddUint64(&skipped, 1)
				return
			}

			zoneNames := parseCentralNicDetailPage(detailDoc)
			atomic.AddUint64(&fetched, 1)

			mu.Lock()
			results = append(results, detailResult{
				lang:  j.lang,
				entry: j.entry,
				zones: zoneNames,
			})
			mu.Unlock()
		}(job)
	}
	wg.Wait()

	Trace("@{.}CentralNic: fetched %d detail pages, %d skipped\n", fetched, skipped)

	// Phase 3: Apply results, IANA takes precedence
	ianaPrefix := ianaBaseURL + "/"
	var added uint64
	for _, r := range results {
		for _, zoneName := range r.zones {
			z, ok := zones[zoneName]
			if !ok {
				continue
			}

			// Check if IANA already has a policy for this language
			var hasIANA bool
			for _, p := range z.Policies {
				if p.Type == TypeIDNTable && p.Key == r.lang && strings.HasPrefix(p.Source, ianaPrefix) {
					hasIANA = true
					break
				}
			}
			if hasIANA {
				continue
			}

			z.AddPolicy(TypeIDNTable, r.lang, r.entry.URL, centralNicIndexURL, "")
			z.Languages = append(z.Languages, r.lang)
			atomic.AddUint64(&added, 1)
		}
	}

	Trace("@{.}CentralNic: added %d IDN table policies\n", added)
	if added == 0 && len(entries) > 10 {
		return fmt.Errorf("CentralNic: no policies added from %d entries, HTML change?", len(entries))
	}

	return nil
}
