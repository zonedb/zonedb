package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/wsxiaoys/terminal/color"
	"github.com/zonedb/zonedb/build"
)

func main() {
	// Default options
	flag.BoolVar(&build.Verbose, "v", false, "enable verbose logging")
	flag.StringVar(&build.BaseDir, "dir", "./", "working directory (location of zones.txt and metadata dir)")
	flag.IntVar(&build.Concurrency, "c", build.Concurrency, "number of concurrent connections")

	// Filters
	tlds := flag.Bool("tlds", false, "work on top-level domains only")
	filterZones := flag.String("zones", "", "select specific working zones (comma-delimited)")
	filterRegexp := flag.String("x", "", "select working zones on by regular expression")
	filterTags := flag.String("tags", "", "select working zones by tags (comma-delimited)")

	// Query operations
	listZones := flag.Bool("list", false, "list working zones")
	listTags := flag.Bool("list-tags", false, "list tags in working zones")
	listLocations := flag.Bool("list-locations", false, "list locations in working zones")

	// Test operations
	verifyNS := flag.Bool("verify-ns", false, "verify name servers")
	verifyWhois := flag.Bool("verify-whois", false, "verify whois servers")
	checkPS := flag.Bool("ps", false, "check against Public Suffix List")

	// Mutate operations
	addTags := flag.String("add-tags", "", "add tags to zones (comma-delimited)")
	addLocations := flag.String("add-locations", "", "add locations to zones (comma-delimited)")
	removeTags := flag.String("remove-tags", "", "remove tags from zones (comma-delimited)")
	removeLocations := flag.String("remove-locations", "", "remove locations from zones (comma-delimited)")
	updateRoot := flag.Bool("update-root", false, "retrieve updates to the root zone file")
	updateNS := flag.Bool("update-ns", false, "update name servers")
	updateRubyWhois := flag.Bool("update-ruby-whois", false, "query Ruby Whois for whois servers")
	updateWhois := flag.Bool("update-whois", false, "query whois-servers.net for whois servers")
	updateIANA := flag.Bool("update-iana", false, "query IANA for metadata")
	updateIDNTables := flag.Bool("update-idn", false, "fetch IDN tables")
	updateIDNURLs := flag.Bool("update-idn-urls", false, "fetch list of URLs for IDN for each zone")
	updateAll := flag.Bool("update", false, "update all (root zone, whois, IANA data, IDN tables)")

	// Write operations
	write := flag.Bool("w", false, "write zones.txt and metadata")
	generateGo := flag.Bool("generate-go", false, "generate Go source code to specified directory")

	flag.Usage = func() {
		color.Fprintf(os.Stderr, "@{!}Usage:@{|} %s [arguments] <command>\n\n", os.Args[0])
		color.Fprintf(os.Stderr, "@{!}Available arguments: \n")
		flag.PrintDefaults()
		color.Fprintf(os.Stderr, "\n")
		os.Exit(1)
	}
	flag.Parse()

	startTime := time.Now()
	defer func() {
		elapsed := time.Since(startTime)
		elapsed -= elapsed % 1000000
		color.Fprintf(os.Stderr, "@{.}Time elapsed: %s\n", elapsed)
	}()

	zones, errs := build.ReadZones()
	defer func() {
		if len(errs) > 0 {
			build.LogFatal(fmt.Errorf("operation failed with %d issue(s)", len(errs)))
		}
	}()
	if len(errs) > 0 {
		return
	}
	workZones := zones

	if *tlds {
		workZones = build.TLDs(zones)
		color.Fprintf(os.Stderr, "@{.}Working on top-level domains\n")
	}

	if *filterZones != "" {
		domains := strings.Split(*filterZones, ",")
		filtered := make(map[string]*build.Zone, len(domains))
		for _, d := range domains {
			d = build.Normalize(d)
			if z, ok := workZones[d]; ok {
				filtered[d] = z
			}
		}
		workZones = filtered
	}

	if *filterRegexp != "" {
		re, err := regexp.Compile(*filterRegexp)
		if err != nil {
			errs = append(errs, err)
			build.LogFatal(err)
		}
		filtered := make(map[string]*build.Zone, len(workZones))
		for d, z := range workZones {
			if re.MatchString(d) {
				filtered[d] = z
			}
		}
		workZones = filtered
	}

	if *filterTags != "" {
		tags := strings.Split(*filterTags, ",")
		filtered := make(map[string]*build.Zone, len(workZones))
		for d, z := range workZones {
			s := build.NewSet(z.Tags...)
			for _, t := range tags {
				if _, ok := s[t]; ok {
					filtered[d] = z
					break
				}
			}
		}
		workZones = filtered
	}

	color.Fprintf(os.Stderr, "@{.}Working on %d zone(s) out of %d\n", len(workZones), len(zones))

	// Add newly found zones?
	addNew := len(workZones) == len(zones)

	if *listZones || len(workZones) < len(zones) {
		domains := build.SortedDomains(workZones)
		color.Fprintf(os.Stderr, "@{.}Zones: @{c}%s\n", strings.Join(domains, " "))
	}

	if *updateRoot || *updateAll {
		err := build.FetchRootZone(workZones, addNew)
		if err != nil {
			errs = append(errs, err)
			build.LogError(err)
		}
	}

	if *updateRubyWhois || *updateAll {
		err := build.FetchRubyWhoisServers(workZones, addNew)
		if err != nil {
			errs = append(errs, err)
			build.LogError(err)
		}
	}

	// whois-servers.net overrides Ruby Whois
	if *updateWhois || *updateAll {
		err := build.QueryWhoisServers(workZones)
		if err != nil {
			errs = append(errs, err)
			build.LogError(err)
		}
	}

	// IANA overrides the above
	if *updateIANA || *updateAll {
		err := build.QueryIANA(workZones)
		if err != nil {
			errs = append(errs, err)
			build.LogError(err)
		}
	}

	if *updateNS || *updateAll {
		err := build.FetchNameServers(workZones)
		if err != nil {
			build.LogError(err)
		}
	}

	// Explicitly _not_ under updateAll at this time, until we're confident this
	// won't get us blacklisted on IANA's web-servers
	if *updateIDNURLs {
		err := build.FetchIDNURLs(workZones)
		if err != nil {
			errs = append(errs, err)
			build.LogError(err)
		}
	}

	// Explicitly _not_ under updateAll at this time, until we're confident this
	// won't get us blacklisted on IANA's web-servers
	if *updateIDNTables {
		err := build.FetchIDNTables(workZones)
		if err != nil {
			errs = append(errs, err)
			build.LogError(err)
		}
	}

	if *removeTags != "" {
		tags := strings.Split(*removeTags, ",")
		build.RemoveTags(workZones, tags)
	}

	if *addTags != "" {
		tags := strings.Split(*addTags, ",")
		build.AddTags(workZones, tags)
	}

	if *listTags {
		tags := build.NewSet()
		for _, z := range workZones {
			tags.Add(z.Tags...)
		}
		color.Fprintf(os.Stderr, "@{.}Tags: @{c}%s\n", strings.Join(tags.Values(), " "))
	}

	if *removeLocations != "" {
		locations := strings.Split(*removeLocations, ",")
		build.RemoveLocations(workZones, locations)
	}

	if *addLocations != "" {
		locations := strings.Split(*addLocations, ",")
		build.AddLocations(workZones, locations)
	}

	if *listLocations {
		locations := build.NewSet()
		for _, z := range workZones {
			locations.Add(z.Locations...)
		}
		color.Fprintf(os.Stderr, "@{.}Locations: @{c}%s\n", strings.Join(locations.Values(), " "))
	}

	if *verifyNS {
		build.VerifyNameServers(workZones)
	}

	if *verifyWhois {
		build.VerifyWhois(workZones)
	}

	if *checkPS {
		build.CheckPublicSuffix(workZones)
	}

	// Fold newly added zones back in
	for d, z := range workZones {
		zones[d] = z
	}

	if *write {
		err := build.WriteZones(zones)
		if err != nil {
			errs = append(errs, err)
			build.LogFatal(err)
		}
	}

	if *generateGo {
		err := build.GenerateGo(zones)
		if err != nil {
			errs = append(errs, err)
			build.LogFatal(err)
		}
	}
}
