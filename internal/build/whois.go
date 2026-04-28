package build

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
	"sync/atomic"
	"time"

	"github.com/miekg/dns"
	"github.com/wsxiaoys/terminal/color"
	"golang.org/x/net/idna"
	"golang.org/x/sync/errgroup"
)

// QueryWhoisServers queries whois-servers.net for CNAME’d whois servers.
func QueryWhoisServers(ctx context.Context, zones map[string]*Zone) error {
	color.Fprintf(os.Stderr, "@{.}Querying whois-servers.net for %d zones...\n", len(zones))

	// Get name servers for whois-servers.net
	rrs, err := resolver.LookupNS(ctx, "whois-servers.net")
	if err != nil {
		return fmt.Errorf("looking up NS for whois-servers.net: %w", err)
	}
	if len(rrs) == 0 {
		return errors.New("no name servers for whois-servers.net")
	}
	host := rrs[0].Host

	// Iterate zones
	var found int32
	err = mapZones(ctx, zones, func(gctx context.Context, z *Zone) error {
		qctx, cancel := context.WithTimeout(gctx, 10*time.Second)
		defer cancel()
		name := z.ASCII() + ".whois-servers.net."
		rmsg, err := exchange(qctx, host, name, dns.TypeCNAME) // Single-depth CNAME query
		if err != nil {
			// Abort siblings only on group cancellation; per-zone qctx timeouts are per-zone.
			if gctx.Err() != nil {
				return gctx.Err()
			}
			if terr, ok := err.(timeouter); ok && !terr.Timeout() {
				color.Fprintf(os.Stderr, "@{r}Error querying %s for %s: %s\n", name, z, err.Error())
			}
			return nil
		}
		for _, rr := range rmsg.Answer {
			if cname, ok := rr.(*dns.CNAME); ok {
				addr := cname.Target

				// whois-servers.net occasionally returns whois.ripe.net (unusable)
				if addr == "whois.ripe.net." {
					return nil
				}
				if verifyWhois(qctx, addr) != nil {
					return nil
				}
				z.WhoisServer = Normalize(addr)
				atomic.AddInt32(&found, 1)
				return nil
			}
		}
		return nil
	})
	color.Fprintf(os.Stderr, "@{.}Found %d whois servers\n", found)
	return err
}

const rubyWhoisURL = "https://github.com/weppos/whois/raw/HEAD/data/tld.json"

// FetchRubyWhoisServers fetches whois servers from the Ruby Whois project.
func FetchRubyWhoisServers(ctx context.Context, zones map[string]*Zone, addNew bool, cache *ETagCache) error {
	res, err := FetchWithETag(ctx, rubyWhoisURL, cache)
	if err != nil {
		return err
	}
	if res == nil {
		return nil // 304 Not Modified
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
		return fmt.Errorf("decoding Ruby Whois JSON: %w", err)
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
			z = &Zone{Domain: d}
			color.Fprintf(os.Stderr, "@{g}New zone @{g!}%s@{g}\n", z)
			zones[d] = z
		}
		if rec.Host != "" && rec.Host != z.WhoisServer {
			err := verifyWhois(ctx, rec.Host)
			if err == nil {
				z.WhoisServer = Normalize(rec.Host)
				servers++
			}
		}
		if z.WhoisServer != "" && rec.URL != "" && rec.URL != z.WhoisURL {
			z.WhoisURL = rec.URL
			urls++
		}
	}
	color.Fprintf(os.Stderr, "@{.}Set %d whois servers, %d URLs from Ruby Whois\n", servers, urls)
	return nil
}

// QueryIANA fetches whois data from whois.iana.org.
func QueryIANA(ctx context.Context, zones map[string]*Zone) error {
	tlds := TLDs(zones)
	color.Fprintf(os.Stderr, "@{.}Querying whois.iana.org for %d TLDs...\n", len(tlds))
	g, gctx := errgroup.WithContext(ctx)
	g.SetLimit(Concurrency)
	for _, z := range tlds {
		g.Go(func() error {
			if err := tldWhois(gctx, z); err != nil {
				// Abort the batch only on group cancellation; per-TLD errors
				// are logged and skipped — many TLDs have flaky whois servers.
				if gctx.Err() != nil {
					return gctx.Err()
				}
				LogWarning(fmt.Errorf("%s: whois.iana.org: %w", z, err))
			}
			return nil
		})
	}
	return g.Wait()
}

var (
	whoisLine    = regexp.MustCompile(`^([^:]+)\:\s+(.+)$`)
	nserverValue = regexp.MustCompile(`^(\S+)`)
)

func tldWhois(ctx context.Context, z *Zone) error {
	b, err := queryWhois(ctx, "whois.iana.org", z.ASCII())
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
			if verifyWhois(ctx, v) != nil {
				break
			}
			z.WhoisServer = Normalize(v)
		}
	}
	if err = s.Err(); err != nil {
		return fmt.Errorf("scanning whois response: %w", err)
	}
	return nil
}

func queryWhois(ctx context.Context, addr, query string) ([]byte, error) {
	if !strings.Contains(addr, ":") {
		addr = addr + ":43"
	}
	c, err := dialer.DialContext(ctx, "tcp", addr)
	if err != nil {
		return nil, fmt.Errorf("dialing %s: %w", addr, err)
	}
	defer c.Close()
	// Propagate ctx deadline to the connection so read/write are also bounded.
	if dl, ok := ctx.Deadline(); ok {
		_ = c.SetDeadline(dl)
	}
	if _, err = fmt.Fprint(c, query, "\r\n"); err != nil {
		// Re-wrap with ctx.Err() so callers can detect cancellation/deadline
		// uniformly — SetDeadline surfaces os.ErrDeadlineExceeded otherwise.
		if ctxErr := ctx.Err(); ctxErr != nil {
			return nil, fmt.Errorf("writing whois query: %w", ctxErr)
		}
		return nil, fmt.Errorf("writing whois query: %w", err)
	}
	res, err := io.ReadAll(c)
	if err != nil {
		if ctxErr := ctx.Err(); ctxErr != nil {
			return nil, fmt.Errorf("reading whois response: %w", ctxErr)
		}
		return nil, fmt.Errorf("reading whois response: %w", err)
	}
	return res, nil
}

// VerifyWhois verifies that the whois servers respond on TCP port 43.
// Per-zone verification failures clear the bad server and are logged; ctx
// cancellation surfaces so an interrupted pass isn't reported as success.
func VerifyWhois(ctx context.Context, zones map[string]*Zone) error {
	color.Fprintf(os.Stderr, "@{.}Verifying whois servers for %d zones...\n", len(zones))
	err := mapZones(ctx, zones, func(gctx context.Context, z *Zone) error {
		if z.WhoisServer != "" {
			if err := verifyWhois(gctx, z.WhoisServer); err != nil {
				// Abort only on group cancellation; otherwise clear the unreachable server and continue.
				if gctx.Err() != nil {
					return gctx.Err()
				}
				LogWarning(fmt.Errorf("can’t verify whois server %s: %s", z.WhoisServer, err))
				z.WhoisServer = ""
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	return ctx.Err()
}

func verifyWhois(ctx context.Context, host string) error {
	host, err := idna.ToASCII(Normalize(host))
	if err != nil {
		return fmt.Errorf("normalizing host %q: %w", host, err)
	}
	return CanDial(ctx, "tcp", host+":43")
}
