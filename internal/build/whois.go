package build

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"regexp"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/wsxiaoys/terminal/color"
	"golang.org/x/net/idna"
)

// QueryWhoisServers queries whois-servers.net for CNAME’d whois servers.
func QueryWhoisServers(zones map[string]*Zone) error {
	color.Fprintf(os.Stderr, "@{.}Querying whois-servers.net for %d zones...\n", len(zones))
	var found int32
	mapZones(zones, func(z *Zone) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		name := z.ASCII() + ".whois-servers.net."
		cname, _ := resolver.LookupCNAME(ctx, name)
		// whois-servers.net occasionally returns whois.ripe.net (unusable)
		if cname == "whois.ripe.net." {
			return
		}
		if verifyWhois(cname) != nil {
			return
		}
		z.WhoisServer = Normalize(cname)
		atomic.AddInt32(&found, 1)
		return
	})
	color.Fprintf(os.Stderr, "@{.}Found %d whois servers\n", found)
	return nil
}

const rubyWhoisURL = "https://github.com/weppos/whois/raw/master/data/tld.json"

// FetchRubyWhoisServers fetches whois servers from the Ruby Whois project.
func FetchRubyWhoisServers(zones map[string]*Zone, addNew bool) error {
	res, err := Fetch(rubyWhoisURL)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	records := make(map[string]struct {
		Host    string `json:"host"`
		Adapter string `json:"adapter"`
		URL     string `json:"url"`
	})
	d := json.NewDecoder(res.Body)
	err = d.Decode(&records)
	if err != nil {
		return err
	}
	var servers, urls int
	for d, rec := range records {
		// Skip broken records
		if d == "_" || d == "whois.nic.beauty" {
			continue
		}
		// Skip empty records
		if rec.Host == "" && rec.Adapter == "none" && rec.URL == "" {
			continue
		}
		d = Normalize(d)
		z := zones[d]
		if z == nil {
			if !addNew {
				continue
			}
			color.Fprintf(os.Stderr, "@{g}New zone @{g!}%s@{g}\n", d)
			z = &Zone{Domain: d}
			zones[d] = z
		}
		if rec.Host != "" && rec.Host != z.WhoisServer {
			err := verifyWhois(rec.Host)
			if err == nil {
				z.WhoisServer = Normalize(rec.Host)
				servers++
			}
		}
		if rec.URL != "" && rec.URL != z.WhoisURL {
			z.WhoisURL = rec.URL
			urls++
		}
	}
	color.Fprintf(os.Stderr, "@{.}Set %d whois servers, %d URLs from Ruby Whois\n", servers, urls)
	return nil
}

// QueryIANA fetches whois data from whois.iana.org.
func QueryIANA(zones map[string]*Zone) error {
	tlds := TLDs(zones)
	color.Fprintf(os.Stderr, "@{.}Querying whois.iana.org for %d TLDs...\n", len(tlds))
	limiter := make(chan struct{}, Concurrency)
	var wg sync.WaitGroup
	for _, z := range tlds {
		limiter <- struct{}{}
		wg.Add(1)
		go func(z *Zone) {
			defer func() {
				<-limiter
				wg.Done()
			}()
			err := tldWhois(z)
			if err != nil {
				LogWarning(err)
			}
		}(z)
	}
	wg.Wait()
	return nil
}

var (
	whoisLine    = regexp.MustCompile(`^([^:]+)\:\s+(.+)$`)
	nserverValue = regexp.MustCompile(`^(\S+)`)
)

func tldWhois(z *Zone) error {
	b, err := queryWhois("whois.iana.org", z.ASCII())
	if err != nil {
		return err
	}
	s := bufio.NewScanner(bytes.NewBuffer(b))
	for s.Scan() {
		m := whoisLine.FindStringSubmatch(s.Text())
		if m == nil {
			continue
		}
		k, v := strings.ToLower(m[1]), m[2]
		switch k {
		case "domain":
			d := Normalize(v)
			if d != z.Domain {
				return fmt.Errorf("IANA whois returned %s, expected %s", d, z.Domain)
			}

		case "nserver":
			vm := nserverValue.FindStringSubmatch(v)
			if vm == nil {
				break
			}
			z.NameServers = append(z.NameServers, Normalize(vm[1]))

		case "whois":
			if verifyWhois(v) != nil {
				break
			}
			z.WhoisServer = Normalize(v)
		}
	}
	if err = s.Err(); err != nil {
		return err
	}
	return nil
}

func queryWhois(addr, query string) ([]byte, error) {
	if !strings.Contains(addr, ":") {
		addr = addr + ":43"
	}
	c, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	if _, err = fmt.Fprint(c, query, "\r\n"); err != nil {
		return nil, err
	}
	res, err := ioutil.ReadAll(c)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// VerifyWhois verifies that the whois servers respond on TCP port 43.
func VerifyWhois(zones map[string]*Zone) {
	color.Fprintf(os.Stderr, "@{.}Verifying whois servers for %d zones...\n", len(zones))
	mapZones(zones, func(z *Zone) {
		if z.WhoisServer != "" {
			if err := verifyWhois(z.WhoisServer); err != nil {
				LogWarning(fmt.Errorf("can’t verify whois server %s: %s", z.WhoisServer, err))
				z.WhoisServer = ""
			}
		}
	})
}

func verifyWhois(host string) error {
	host, err := idna.ToASCII(Normalize(host))
	if err != nil {
		return err
	}
	err = CanDial("tcp", host+":43")
	return err
}
