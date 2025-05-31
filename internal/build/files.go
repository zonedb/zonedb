package build

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"slices"
	"sort"
	"strings"

	"github.com/wsxiaoys/terminal/color"
	"golang.org/x/net/idna"
)

// BaseDir specifies the base working directory that contains the ZoneDB flat-file database.
var BaseDir = "."

// ReadZones reads the zone database from the local filesystem.
func ReadZones() (zones map[string]*Zone, errs []error) {
	zones, errs = ReadZonesFile()
	errs = append(errs, ReadMetadata(zones)...)
	return
}

// ReadZonesFile reads and parses the zones.txt file.
func ReadZonesFile() (zones map[string]*Zone, errs []error) {
	path := filepath.Join(BaseDir, "zones.txt")
	f, err := os.Open(path)
	if err != nil {
		LogFatal(err)
	}
	defer f.Close()
	zones = make(map[string]*Zone)
	var line int
	s := bufio.NewScanner(f)
	for s.Scan() {
		line++
		d := s.Text()
		n := Normalize(d)
		if n != d {
			err := fmt.Errorf("invalid domain name %s at %s:%d", d, path, line)
			errs = append(errs, err)
			LogError(err)
			continue
		}
		if d == "" {
			err := fmt.Errorf("blank line at %s:%d", path, line)
			errs = append(errs, err)
			LogWarning(err)
			continue
		}
		if _, ok := zones[d]; ok {
			err := fmt.Errorf("duplicate zone %s at %s:%d", d, path, line)
			errs = append(errs, err)
			LogWarning(err)
			continue
		}
		z := &Zone{Domain: d}
		zones[z.Domain] = z
		pd := z.ParentDomain()
		if pd != "" {
			p, ok := zones[pd]
			if !ok {
				err := fmt.Errorf("missing parent %s for subdomain %s at %s:%d", pd, d, path, line)
				errs = append(errs, err)
				LogWarning(err)
				continue
			}
			p.subdomains = append(p.subdomains, d)
		}
	}
	if err := s.Err(); err != nil {
		errs = append(errs, err)
		LogError(err)
	}

	for _, z := range zones {
		sort.Strings(z.subdomains)
	}

	color.Fprintf(os.Stderr, "@{.}Read %d zones\n", len(zones))
	return
}

// ReadMetadata reads JSON files for each zone in the metadata directory.
func ReadMetadata(zones map[string]*Zone) (errs []error) {
	dir := filepath.Join(BaseDir, "metadata")
	paths, _ := filepath.Glob(filepath.Join(dir, "*.json"))
	var read int
	for _, path := range paths {
		// Ensure filename equals ASCII/punycode domain name
		base := filepath.Base(path)
		ext := filepath.Ext(base)
		di := strings.TrimSuffix(base, ext)
		d, err := idna.ToUnicode(di)
		if err != nil {
			errs = append(errs, err)
			LogError(err)
			continue
		}

		// Ensure the domain exists in zones.txt
		z, ok := zones[d]
		if !ok {
			err = fmt.Errorf("domain not found in zones.txt: %s", base)
			errs = append(errs, err)
			LogWarning(err)
			continue
		}

		// Parse the JSON metadata
		f, err := os.Open(path)
		if err != nil {
			err = fmt.Errorf("cannot load %s: %s", base, err)
			errs = append(errs, err)
			LogError(err)
			continue
		}
		dec := json.NewDecoder(f)
		err = dec.Decode(z)
		f.Close()
		if err != nil && err != io.EOF {
			err = fmt.Errorf("unable to parse %s: %s", base, err)
			errs = append(errs, err)
			LogError(err)
			continue
		}
		read++

		// Ensure domain name matches
		if z.Domain != d {
			err = fmt.Errorf("domain %s doesn’t match filename (expected %s): %s", z.Domain, z.ASCII(), base)
			errs = append(errs, err)
			LogError(err)
			continue
		}

		// Transition on load
		z.transition()
	}
	color.Fprintf(os.Stderr, "@{.}Read %d metadata files\n", read)
	return
}

// WriteZonesFile writes the zones.txt file.
func WriteZonesFile(zones map[string]*Zone) error {
	path := filepath.Join(BaseDir, "zones.txt")
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	domains := SortedDomains(zones)
	for i := range domains {
		_, err = f.WriteString(domains[i])
		if err != nil {
			return err
		}
		_, err = f.WriteString("\n")
		if err != nil {
			return err
		}
	}
	color.Fprintf(os.Stderr, "@{.}Wrote %d zones\n", len(domains))
	return nil
}

// WriteMetadata writes the metadata/*.json files.
func WriteMetadata(zones map[string]*Zone, delete bool) error {
	var created, deleted []string
	var written int
	for _, z := range zones {
		z.Normalize()
		path := filepath.Join(BaseDir, "metadata", z.ASCII()+".json")
		if delete || !z.HasMetadata() {
			err := os.Remove(path)
			if err == nil {
				deleted = append(deleted, z.ASCII())
			} else if !errors.Is(err, fs.ErrNotExist) {
				LogError(err)
			}
			continue
		}
		var newFile bool
		f, err := os.Open(path) // open read-only
		newFile = errors.Is(err, fs.ErrNotExist)
		if f != nil {
			_ = f.Close() // nothing to sync, cannot error
		}
		err = writeMetadataFile(z, path)
		if err != nil {
			return err
		}
		written++
		if newFile {
			created = append(created, z.ASCII())
		}
	}
	color.Fprintf(os.Stderr, "@{.}Wrote %d metadata files\n", written)
	slices.Sort(created)
	slices.Sort(deleted)
	color.Fprintf(os.Stderr, "@{.}Created %d: %v\n", len(created), created)
	color.Fprintf(os.Stderr, "@{.}Deleted %d: %v\n", len(deleted), deleted)
	return nil
}

func writeMetadataFile(z *Zone, path string) (err error) {
	var f *os.File
	f, err = os.Create(path) // open read-write and truncate
	if err != nil {
		return
	}
	defer func() {
		if err != nil {
			return
		}
		err = f.Close()
	}()
	enc := json.NewEncoder(f)
	enc.SetIndent("", "\t")
	err = enc.Encode(&z)
	return
}
