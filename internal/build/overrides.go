package build

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"

	"github.com/tailscale/hujson"
	"github.com/wsxiaoys/terminal/color"
)

// OverridesFile is the human-edited HuJSON file of build-time overrides
const OverridesFile = "content/overrides.hujson"

// Overrides holds human-authored corrections applied after the upstream
// IANA/ICANN feeds so they win over auto-derived data.
type Overrides struct {
	// BrandExempt maps a zone to the reason it must not be tagged "brand"
	// (Spec 13 TLDs that nonetheless allow open public registration).
	BrandExempt map[string]string `json:"brand_exempt"`
}

// ReadOverrides parses the overrides.hujson file, normalizing zone keys.
func ReadOverrides() (*Overrides, error) {
	path := filepath.Join(BaseDir, OverridesFile)
	raw, err := os.ReadFile(path)
	if errors.Is(err, fs.ErrNotExist) {
		LogWarning(fmt.Errorf("%s not found; overrides not applied", path))
		return &Overrides{}, nil
	}
	if err != nil {
		return nil, fmt.Errorf("reading %s: %w", path, err)
	}
	std, err := hujson.Standardize(raw)
	if err != nil {
		return nil, fmt.Errorf("parsing %s: %w", path, err)
	}
	var o Overrides
	dec := json.NewDecoder(bytes.NewReader(std))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&o); err != nil {
		return nil, fmt.Errorf("decoding %s: %w", path, err)
	}
	if o.BrandExempt != nil {
		normalized := make(map[string]string, len(o.BrandExempt))
		for zone, reason := range o.BrandExempt {
			normalized[Normalize(zone)] = reason
		}
		o.BrandExempt = normalized
	}
	return &o, nil
}

// ApplyBrandExempt removes the "brand" tag from each zone in BrandExempt,
// returning an error for any entry whose zone is not present in zones.
func (o *Overrides) ApplyBrandExempt(zones map[string]*Zone) []error {
	var errs []error
	var stripped int
	// Sort zones so log and error order is deterministic across builds.
	zoneNames := make([]string, 0, len(o.BrandExempt))
	for zone := range o.BrandExempt {
		zoneNames = append(zoneNames, zone)
	}
	sort.Strings(zoneNames)
	for _, zone := range zoneNames {
		z, ok := zones[zone]
		if !ok {
			err := fmt.Errorf("brand_exempt override for unknown zone %q in %s", zone, OverridesFile)
			errs = append(errs, err)
			LogError(err)
			continue
		}
		if z.RemoveTags(TagBrand) > 0 {
			stripped++
			Trace("@{.}brand_exempt: removed brand from %s (%s)\n", zone, o.BrandExempt[zone])
		}
	}
	if stripped > 0 {
		color.Fprintf(os.Stderr, "@{.}Applied brand_exempt to %d zone(s)\n", stripped)
	}
	return errs
}
