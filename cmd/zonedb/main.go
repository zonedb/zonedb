package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/wsxiaoys/terminal/color"
	"github.com/zonedb/zonedb/internal/build"
)

// outputJSON outputs data as JSON or falls back to text output with color formatting
func outputJSON(jsonOutput bool, data interface{}, key string, textOutput func()) {
	if jsonOutput {
		jsonData := map[string]interface{}{
			key: data,
		}
		if jsonBytes, err := json.MarshalIndent(jsonData, "", "  "); err == nil {
			fmt.Println(string(jsonBytes))
		} else {
			build.LogError(fmt.Errorf("failed to marshal JSON: %v", err))
		}
	} else {
		textOutput()
	}
}

// outputZonesJSON outputs zones with filter context when available
func outputZonesJSON(jsonOutput bool, domains []string, filterTags string, filterZones string, filterRegexp string, textOutput func()) {
	if jsonOutput {
		var jsonData map[string]interface{}

		// Create structured output based on filters applied
		if filterTags != "" {
			tags := strings.Split(filterTags, ",")
			if len(tags) == 1 {
				// Single tag: {"zones": {"tags": {"brand": [...]}}}
				jsonData = map[string]interface{}{
					"zones": map[string]interface{}{
						"tags": map[string]interface{}{
							tags[0]: domains,
						},
					},
				}
			} else {
				// Multiple tags: {"zones": {"tags": {"all_of": ["tag1", "tag2"], "filtered": [...]}}}
				jsonData = map[string]interface{}{
					"zones": map[string]interface{}{
						"tags": map[string]interface{}{
							"all_of":   tags,
							"filtered": domains,
						},
					},
				}
			}
		} else if filterZones != "" {
			// Specific zones filter
			jsonData = map[string]interface{}{
				"zones": map[string]interface{}{
					"filter": map[string]interface{}{
						"domains": strings.Split(filterZones, ","),
					},
					"filtered": domains,
				},
			}
		} else if filterRegexp != "" {
			// Regex filter
			jsonData = map[string]interface{}{
				"zones": map[string]interface{}{
					"filter": map[string]interface{}{
						"regexp": filterRegexp,
					},
					"filtered": domains,
				},
			}
		} else {
			// No specific filter, use simple format
			jsonData = map[string]interface{}{
				"zones": domains,
			}
		}

		if jsonBytes, err := json.MarshalIndent(jsonData, "", "  "); err == nil {
			fmt.Println(string(jsonBytes))
		} else {
			build.LogError(fmt.Errorf("failed to marshal JSON: %v", err))
		}
	} else {
		textOutput()
	}
}

func main() {
	// Default options
	flag.BoolVar(&build.Verbose, "v", false, "enable verbose logging")
	flag.StringVar(&build.BaseDir, "dir", "./", "working directory (location of zones.txt and metadata dir)")
	flag.IntVar(&build.Concurrency, "c", build.Concurrency, "number of concurrent connections")

	// Filters
	tlds := flag.Bool("tlds", false, "work on top-level domains only")
	filterIDN := flag.Bool("idn", false, "work on IDN (non-ASCII) zones only")
	filterZones := flag.String("zones", "", "select specific working zones (comma-delimited)")
	filterRegexp := flag.String("x", "", "filter working zones by regular expression")
	filterGrep := flag.String("grep", "", "filter working zones by regular expression across all metadata")
	excludeGrep := flag.String("exclude-grep", "", "exclude working zones by regular expression across all metadata")
	filterTags := flag.String("tags", "", "filter working zones that match ALL tags (comma-delimited)")
	excludeTags := flag.String("exclude-tags", "", "exclude working zones with ANY supplied tags (comma-delimited)")

	// Query operations
	listZones := flag.Bool("list", false, "list working zones")
	listTags := flag.Bool("list-tags", false, "list tags in working zones")
	listLocations := flag.Bool("list-locations", false, "list locations in working zones")
	listWildcards := flag.Bool("list-wildcards", false, "list zones with wildcarded DNS")
	jsonOutput := flag.Bool("json", false, "output results in JSON format")

	// Test operations
	verifyNS := flag.Bool("verify-ns", false, "verify name servers")
	verifyWhois := flag.Bool("verify-whois", false, "verify whois servers")
	checkPS := flag.Bool("ps", false, "check against Public Suffix List")

	// Mutate operations
	addLanguages := flag.String("add-languages", "", "add BCP 47 language tags to zones (comma-delimited)")
	guessLanguages := flag.Bool("guess-languages", false, "guess BCP 47 languages for zones")
	setInfoURL := flag.String("set-info-url", "\u0000", "set zone(s) info URLs")
	updateInfoURL := flag.Bool("update-info-url", false, "update zone(s) info URLs")
	addRDAPURL := flag.String("add-rdap-url", "", "add RDAP URL to zones")
	addTags := flag.String("add-tags", "", "add tags to zones (comma-delimited)")
	addLocations := flag.String("add-locations", "", "add locations to zones (comma-delimited)")
	removeTags := flag.String("remove-tags", "", "remove tags from zones (comma-delimited)")
	removeLocations := flag.String("remove-locations", "", "remove locations from zones (comma-delimited)")
	updateRoot := flag.Bool("update-root", false, "retrieve updates to the root zone file")
	updateNS := flag.Bool("update-ns", false, "update name servers")
	updateWildcards := flag.Bool("update-wildcards", false, "update wildcard IPs")
	updateRubyWhois := flag.Bool("update-ruby-whois", false, "query Ruby Whois for whois servers")
	updateWhois := flag.Bool("update-whois", false, "query whois-servers.net for whois servers")
	updateIANA := flag.Bool("update-iana", false, "query IANA for metadata")
	updateIANASpecial := flag.Bool("update-iana-special", false, "query IANA for special use domains")
	updateICANN := flag.Bool("update-icann", false, "query ICANN for metadata")
	updateIDN := flag.Bool("update-idn", false, "query IANA for IDN tables")
	updateRDAP := flag.Bool("update-rdap", false, "query IANA for RDAP URLs")
	updateAll := flag.Bool("update", false, "update all (root zone, whois, IANA data, IDN tables)")

	// Delete operations
	deleteZones := flag.Bool("delete", false, "delete working zones")

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

	// Enable quiet mode for JSON output to suppress diagnostic messages
	if *jsonOutput {
		build.Quiet = true
	}

	startTime := time.Now()
	defer func() {
		elapsed := time.Since(startTime)
		elapsed -= elapsed % 1000000
		build.Trace("@{.}Time elapsed: %s\n", elapsed)
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
		build.Trace("@{.}Working on top-level domains\n")
	}

	if *filterIDN {
		filtered := make(map[string]*build.Zone, len(workZones))
		for d, z := range workZones {
			if z.IsIDN() {
				filtered[d] = z
			}
		}
		workZones = filtered
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

	if *filterGrep != "" {
		re, err := regexp.Compile(*filterGrep)
		if err != nil {
			errs = append(errs, err)
			build.LogFatal(err)
		}
		filtered := make(map[string]*build.Zone, len(workZones))
		for d, z := range workZones {
			j, _ := json.MarshalIndent(z, "", "\t")
			if re.Match(j) {
				filtered[d] = z
			}
		}
		workZones = filtered
	}

	if *excludeGrep != "" {
		re, err := regexp.Compile(*excludeGrep)
		if err != nil {
			errs = append(errs, err)
			build.LogFatal(err)
		}
		filtered := make(map[string]*build.Zone, len(workZones))
		for d, z := range workZones {
			j, _ := json.MarshalIndent(z, "", "\t")
			if !re.Match(j) {
				filtered[d] = z
			}
		}
		workZones = filtered
	}

	if *filterTags != "" {
		tags := build.NewSet(strings.Split(*filterTags, ",")...)
		filtered := make(map[string]*build.Zone, len(workZones))
		for d, z := range workZones {
			n := 0
			for _, t := range z.Tags {
				if _, ok := tags[t]; ok {
					n++
				}
			}
			if n == len(tags) {
				filtered[d] = z
			}
		}
		workZones = filtered
	}

	if *excludeTags != "" {
		tags := build.NewSet(strings.Split(*excludeTags, ",")...)
		filtered := make(map[string]*build.Zone, len(workZones))
		for d, z := range workZones {
			filtered[d] = z
			for _, t := range z.Tags {
				if _, ok := tags[t]; ok {
					delete(filtered, d)
					break
				}
			}
		}
		workZones = filtered
	}

	build.Trace("@{.}Working on %d zone(s) out of %d\n", len(workZones), len(zones))

	// Add newly found zones?
	addNew := len(workZones) == len(zones)

	if *listZones || len(workZones) < len(zones) {
		domains := build.SortedDomains(workZones)
		outputZonesJSON(*jsonOutput, domains, *filterTags, *filterZones, *filterRegexp, func() {
			build.Trace("@{.}Zones: @{c}%s\n", strings.Join(domains, " "))
		})
	}

	if *updateRoot || *updateAll {
		err := build.FetchRootZone(workZones, addNew)
		if err != nil {
			errs = append(errs, err)
			build.LogError(err)
		}
	}

	if *updateICANN || *updateAll {
		err := build.FetchGTLDsFromICANN(workZones)
		if err != nil {
			errs = append(errs, err)
			build.LogError(err)
		}
	}

	if *updateRubyWhois {
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

	if *updateIANASpecial || *updateAll {
		err := build.FetchSpecialUseDomainsFromIANA(workZones, addNew)
		if err != nil {
			errs = append(errs, err)
			build.LogError(err)
		}
	}

	if *updateNS || *updateAll {
		err := build.FetchNameServers(workZones, zones)
		if err != nil {
			build.LogError(err)
		}
	}

	if *updateWildcards || *updateAll {
		err := build.FindWildcards(workZones)
		if err != nil {
			build.LogError(err)
		}
	}

	if *updateIDN || *updateAll {
		err := build.FetchIDNTablesFromIANA(workZones)
		if err != nil {
			errs = append(errs, err)
			build.LogError(err)
		}
	}

	if *updateRDAP || *updateAll {
		err := build.FetchRDAPFromIANA(workZones)
		if err != nil {
			errs = append(errs, err)
			build.LogError(err)
		}
	}

	if *addRDAPURL != "" {
		build.AddRDAPURLs(workZones, []string{*addRDAPURL})
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
		outputJSON(*jsonOutput, tags.Values(), "tags", func() {
			color.Fprintf(os.Stderr, "@{.}Tags: @{c}%s\n", strings.Join(tags.Values(), " "))
		})
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
		outputJSON(*jsonOutput, locations.Values(), "locations", func() {
			color.Fprintf(os.Stderr, "@{.}Locations: @{c}%s\n", strings.Join(locations.Values(), " "))
		})
	}

	if *listWildcards {
		zones := build.NewSet()
		for _, z := range workZones {
			if len(z.Wildcards) > 0 {
				zones.Add(z.Domain)
			}
		}
		outputJSON(*jsonOutput, zones.Values(), "zones", func() {
			color.Fprintf(os.Stderr, "@{.}Zones: @{c}%s\n", strings.Join(zones.Values(), " "))
		})
	}

	if *addLanguages != "" {
		languages := strings.Split(*addLanguages, ",")
		build.AddLanguages(workZones, languages)
	}

	if *guessLanguages {
		build.GuessLanguages(workZones)
	}

	if *setInfoURL != "\u0000" {
		for _, z := range workZones {
			z.InfoURL = *setInfoURL
		}
		build.Trace("@{.}Set info URL to: @{c}%s\n", *setInfoURL)
	}

	if *updateInfoURL {
		build.UpdateInfoURLs(workZones)
	}

	if *verifyNS {
		build.VerifyNameServers(workZones)
	}

	build.CountNameServers(workZones)

	if *verifyWhois {
		build.VerifyWhois(workZones)
	}

	if *checkPS {
		build.CheckPublicSuffix(workZones)
	}

	err := build.ValidateTags(workZones)
	if err != nil {
		build.LogFatal(err)
	}

	// Fold newly added zones back in
	for d, z := range workZones {
		z.Normalize()
		zones[d] = z
	}

	// Delete zones
	if *deleteZones {
		if len(workZones) == len(zones) {
			build.LogFatal(errors.New("cannot delete all zones, please select a subset"))
		}
		deleted := build.SortedDomains(workZones)
		build.Trace("@{r!}Deleting %d zones: @{r}%s\n", len(workZones), strings.Join(deleted, " "))
		for d := range workZones {
			delete(zones, d)
		}
	}

	if *write {
		err := build.WriteZonesFile(zones)
		if err != nil {
			errs = append(errs, err)
			build.LogFatal(err)
		}
		err = build.WriteMetadata(workZones, *deleteZones)
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
