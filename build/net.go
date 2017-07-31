package build

import (
	"fmt"
	"net"
	"net/http"
	"sync"
)

const httpUserAgent = "zonedb/1.0 (https://github.com/zonedb/zonedb)"

// Fetch fetches a URL.
func Fetch(u string) (*http.Response, error) {
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", httpUserAgent)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP error for %s: %s", u, res.Status)
	}
	return res, nil
}

// CanDial verifies if possible to connect to a given network and address.
// Returns nil for successful dial or an error.
func CanDial(network, addr string) error {
	var k = network + " " + addr
	mtx.RLock()
	err, ok := dialCache[k]
	mtx.RUnlock()
	if ok {
		return err
	}
	c, err := net.DialTimeout(network, addr, Timeout)
	mtx.Lock()
	dialCache[k] = err
	mtx.Unlock()
	if err != nil {
		return err
	}
	c.Close()
	return nil
}

var dialCache = make(map[string]error)
var mtx sync.RWMutex
