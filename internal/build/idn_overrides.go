package build

// IDNOverrides maps zones to a reference zone whose IDN table policies
// and languages should be copied. This is for zones whose operator has
// confirmed IDN support matching another zone they operate, but where
// the data is not listed on the IANA IDN tables page.
var IDNOverrides = map[string]string{
	"cc": "com", // Verisign confirmed .cc supports the same IDN tables as .com
}

// ApplyIDNOverrides copies IDN table policies and languages from
// reference zones to the override zones defined in IDNOverrides.
// Both the source and target zones must be present in the zones map.
func ApplyIDNOverrides(zones map[string]*Zone) {
	for target, source := range IDNOverrides {
		tz, ok := zones[target]
		if !ok {
			continue
		}
		sz, ok := zones[source]
		if !ok {
			Trace("@{y}IDN override: source zone %q not found for %q\n", source, target)
			continue
		}

		// Collect source IDN table policies.
		var (
			policies  []Policy
			languages []string
		)
		for _, p := range sz.Policies {
			if p.Type == TypeIDNTable {
				policies = append(policies, p)
				if p.Key != "" {
					languages = append(languages, p.Key)
				}
			}
		}

		if len(policies) == 0 {
			Trace("@{y}IDN override: source zone %q has no IDN table policies\n", source)
			continue
		}

		// Add source policies to target, preserving any existing
		// non-IANA policies the target already has.
		for _, p := range policies {
			tz.AddPolicy(p.Type, p.Key, p.Value, p.Source, p.Comment)
		}
		tz.Languages = append(tz.Languages, languages...)

		Trace("@{.}IDN override: copied %d policies from %q to %q\n", len(policies), source, target)
	}
}
