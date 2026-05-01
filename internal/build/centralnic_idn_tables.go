package build

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/sync/errgroup"
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
// applies changes to zones in the working set; IANA-sourced policies take
// precedence. allZones is used to prune stale policies across the full
// database and to drive an HTML-schema health-check sentinel.
func FetchIDNTablesFromCentralNic(ctx context.Context, zones, allZones map[string]*Zone, cache *ETagCache) error {
	// Phase 0: Fetch index with ETag check
	res, err := FetchWithETag(ctx, centralNicIndexURL, cache)
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

	// Phase 1: Prune stale CentralNic entries across all zones. A
	// filtered run must not leave stale policies elsewhere. Languages
	// from other sources are preserved.
	centralNicPrefix := centralNicBaseURL + "/"
	for _, z := range allZones {
		z.PruneStalePolicies(func(p Policy) bool {
			return p.Type == TypeIDNTable && strings.HasPrefix(p.Source, centralNicPrefix)
		})
	}

	// Phase 2: Parse index, fetch detail pages
	entries := parseCentralNicIndex(doc, baseURL)
	Trace("@{.}CentralNic: parsed %d entries (version-deduped) from index\n", len(entries))

	// Normalize mnemonics and collect detail page fetches
	type fetchJob struct {
		lang  string
		entry centralNicEntry
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
		fetched uint64
		skipped uint64
	)
	g, gctx := errgroup.WithContext(ctx)
	g.SetLimit(Concurrency)

	for _, job := range jobs {
		g.Go(func() error {
			detailRes, err := Fetch(gctx, job.entry.URL)
			if err != nil {
				// Abort the batch when the parent ctx is done; otherwise this
				// is a per-URL fetch failure and the other jobs keep going.
				if gctx.Err() != nil {
					return gctx.Err()
				}
				Trace("@{r}CentralNic: error fetching %s: %v\n", job.entry.URL, err)
				atomic.AddUint64(&skipped, 1)
				return nil
			}
			defer detailRes.Body.Close()

			detailDoc, err := goquery.NewDocumentFromReader(detailRes.Body)
			if err != nil {
				Trace("@{r}CentralNic: error parsing %s: %v\n", job.entry.URL, err)
				atomic.AddUint64(&skipped, 1)
				return nil
			}

			zoneNames := parseCentralNicDetailPage(detailDoc)
			atomic.AddUint64(&fetched, 1)

			mu.Lock()
			results = append(results, detailResult{
				lang:  job.lang,
				entry: job.entry,
				zones: zoneNames,
			})
			mu.Unlock()
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		return err
	}

	Trace("@{.}CentralNic: fetched %d detail pages, %d skipped\n", fetched, skipped)

	// Phase 3: Apply results, IANA takes precedence. Writes are gated on
	// the working set so a filtered run does not mutate zones the user
	// did not select. matched counts detail-page zone names that resolve
	// to any known zone and drives the HTML-change sentinel below —
	// separate from added so the sentinel stays meaningful under a filter.
	ianaPrefix := ianaBaseURL + "/"
	var matched, added uint64
	for _, r := range results {
		for _, zoneName := range r.zones {
			if _, known := allZones[zoneName]; !known {
				continue
			}
			matched++
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
			// Only add languages for TLDs. SLDs (e.g., ru.com) get policies
			// to document IDN support, but not language tags which would
			// incorrectly surface them in language-based search results.
			if z.IsTLD() {
				z.Languages = append(z.Languages, r.lang)
			}
			added++
		}
	}

	Trace("@{.}CentralNic: added %d IDN table policies (%d matched across all zones)\n", added, matched)
	// HTML-change sentinel: keyed on matched (not added) so a user
	// filter can't mask a real parser regression.
	if matched == 0 && len(entries) > 10 {
		return fmt.Errorf("CentralNic: no detail-page zones matched any known zone from %d entries, HTML change?", len(entries))
	}

	return nil
}
