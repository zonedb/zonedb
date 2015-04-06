package build

import (
	"fmt"
	"net"
	"net/http"
	"sync"
)

// Fetch fetches a URL.
func Fetch(u string) (*http.Response, error) {
	res, err := http.Get(u)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP error for %s: %d", u, res.Status)
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
