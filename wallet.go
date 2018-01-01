package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"

	"github.com/onodera-punpun/sako/digest"
	"github.com/sunrisedo/monero"
)

// Request represents a JSON-RPC request sent by a client.
type Request struct {
	// JSON-RPC protocol.
	Version string `json:"jsonrpc"`

	// A String containing the name of the method to be invoked.
	Method string `json:"method"`

	// Object to pass as request parameter to the method.
	Params interface{} `json:"params"`

	// The request iD. This can be of any type. It is used to match the respons
	// with the request that it is replying to.
	ID uint64 `json:"id"`
}

// Response represents a JSON-RPC response returned to a client.
type Response struct {
	Version string           `json:"jsonrpc"`
	Result  *json.RawMessage `json:"result"`
	Error   *json.RawMessage `json:"error"`
}

// encodeClientRequest encodes struff for a JSON-RPC client request.
func encodeRequest(m string, p interface{}) *bytes.Reader {
	c := &Request{
		Version: "2.0",
		Method:  m,
		Params:  p,
		ID:      uint64(rand.Int63()),
	}
	d, _ := json.Marshal(c)

	return bytes.NewReader(d)
}

// decodeResponse decodes stuff from a JSON-RPC client request.
func decodeResponse(r io.Reader, t interface{}) error {
	var res Response
	if err := json.NewDecoder(r).Decode(&res); err != nil {
		return err
	}

	if res.Error != nil {
		return fmt.Errorf("decode: Result is an error")
	}
	if res.Result == nil {
		return fmt.Errorf("decode: Result is null")
	}

	return json.Unmarshal(*res.Result, t)
}

// walletRequest requests and parses JSON from the Monero wallet RPC client into
// a specified interface.
func walletRequest(m string, t interface{}) error {
	c := &http.Client{Timeout: time.Second * 5}

	req, err := http.NewRequest("POST", "http://"+config.RPC+"/json_rpc",
		encodeRequest(m, nil))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	if err := digest.ApplyAuth(c, config.Username, config.Password,
		req); err != nil {
		return err
	}

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

func walletAddress() (string, error) {
	v, err := wallet.GetAddress()
	if err != nil {
		return "", err
	}

	return v.Address, nil
}

func walletBalance() (float64, float64, error) {
	var t = struct {
		Balance   uint64 `json:"balance"`
		UnBalance uint64 `json:"unlocked_balance"`
	}{}
	if err := walletRequest("getbalance", &t); err != nil {
		return 0, 0, err
	}

	return float64(t.Balance) / 1.e+12, float64(t.UnBalance) / 1.e+12, nil
}

func walletHeight() (int64, error) {
	v, err := wallet.GetHeight()
	if err != nil {
		return 0, err
	}

	return v.Height, nil
}

func walletTransactions() (monero.Transfer, error) {
	v, err := wallet.IncomingTransfers("all")
	if err != nil {
		return monero.Transfer{}, err
	}
	fmt.Println(v)

	return v, nil
}
