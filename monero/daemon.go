package monero

import (
	"bytes"
	"fmt"
	"net/http"
	"time"
)

// Daemon describes a daamon.
type Daemon struct {
	// The monerod endpoint.
	URL string
}

// NewDaemon creates a new daemon.
func NewDaemon(url string) *Daemon {
	return &Daemon{url}
}

// request requests and parses JSON from the Monero daemon RPC client into a
// specified interface.
func (d *Daemon) request(m string, p, t interface{}) error {
	c := &http.Client{Timeout: time.Second * 5}

	dat := encodeRequest(m, p)
	req, err := http.NewRequest("POST", d.URL, bytes.NewBuffer(dat))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Length", (string)(len(dat)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	res, err := c.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("request: Returned invalid statuscode %d",
			res.StatusCode)
	}

	return decodeResponse(res.Body, t)
}

// Height returns the daemon's current block height.
func (d *Daemon) Height() (int64, error) {
	var t = struct {
		Count int64 `json:"count"`
	}{}
	if err := d.request("getblockcount", nil, &t); err != nil {
		return 0, err
	}

	return t.Count, nil
}
