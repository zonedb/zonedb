package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/wsxiaoys/terminal/color"
	"github.com/zonedb/zonedb/build"
)

func main() {
	flag.BoolVar(&build.Verbose, "v", false, "enable verbose logging")
	flag.StringVar(&build.BaseDir, "dir", "./", "working directory (location of zones.txt and metadata dir)")
	flag.IntVar(&build.Concurrency, "c", build.Concurrency, "number of concurrent connections")
	tlds := flag.Bool("tlds", false, "operate on top-level domains only")
	fZones := flag.String("zones", "", "zones to operate on (default: all)")
	addTags := flag.String("add-tags", "", "add tags to zones (comma-delimited)")
	removeTags := flag.String("remove-tags", "", "remove tags from zones (comma-delimited)")
	updateRoot := flag.Bool("update-root", false, "retrieve updates to the root zone file")
	updateNS := flag.Bool("update-ns", false, "update name servers")
	verifyNS := flag.Bool("verify-ns", false, "verify name servers")
	updateRubyWhois := flag.Bool("update-ruby-whois", false, "query Ruby Whois for whois servers")
	updateWhois := flag.Bool("update-whois", false, "query whois-servers.net for whois servers")
	verifyWhois := flag.Bool("verify-whois", false, "verify whois servers")
	updateIANA := flag.Bool("update-iana", false, "query IANA for metadata")
	updateAll := flag.Bool("update", false, "update all (root zone, whois, IANA data)")
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

	zones, errs := build.ReadZones()
	if len(errs) > 0 {
		build.LogFatal(fmt.Errorf("read failed with %d issue(s)", len(errs)))
	}
	workZones := zones

	if *tlds {
		workZones = build.TLDs(zones)
		color.Fprintf(os.Stderr, "@{.}Working on top-level domains\n")
	}

	if *fZones != "" {
		domains := strings.Split(*fZones, ",")
		filtered := make(map[string]*build.Zone, len(domains))
		for _, d := range domains {
			d = build.Normalize(d)
			if z, ok := workZones[d]; ok {
				filtered[d] = z
			}
		}
		workZones = filtered
	}
	color.Fprintf(os.Stderr, "@{.}Working on %d zone(s) out of %d\n", len(workZones), len(zones))

	// Add newly found zones?
	addNew := len(workZones) == len(zones)

	if *updateRoot || *updateAll {
		err := build.FetchRootZone(workZones, addNew)
		if err != nil {
			build.LogError(err)
		}
	}

	if *updateRubyWhois || *updateAll {
		err := build.FetchRubyWhoisServers(workZones, addNew)
		if err != nil {
			build.LogError(err)
		}
	}

	// whois-servers.net overrides Ruby Whois
	if *updateWhois || *updateAll {
		err := build.QueryWhoisServers(workZones)
		if err != nil {
			build.LogError(err)
		}
	}

	// IANA overrides the above
	if *updateIANA || *updateAll {
		err := build.QueryIANA(workZones)
		if err != nil {
			build.LogError(err)
		}
	}

	if *updateNS || *updateAll {
		err := build.FetchNameServers(workZones)
		if err != nil {
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

	if *verifyNS {
		build.VerifyNameServers(workZones)
	}

	if *verifyWhois {
		build.VerifyWhois(workZones)
	}

	// Fold newly added zones back in
	for d, z := range workZones {
		zones[d] = z
	}

	if *write {
		err := build.WriteZones(zones)
		if err != nil {
			build.LogFatal(err)
		}
	}

	if *generateGo {
		err := build.GenerateGo(zones)
		if err != nil {
			build.LogFatal(err)
		}
	}
}
