package monero

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"

	digest "github.com/delphinus/go-digest-request"
)

// Wallet describes a wallet.
type Wallet struct {
	// The monero-wallet-rpc endpoint.
	URL string

	// The username and password to use for authentication.
	Username string
	Password string
}

// NewWallet creates a new wallet.
func NewWallet(url, u, p string) *Wallet {
	return &Wallet{url, u, p}
}

// request represents a JSON-RPC request sent by a client.
type request struct {
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

// encodeClientRequest encodes struff for a JSON-RPC client request.
func encodeRequest(m string, p interface{}) *bytes.Reader {
	r := &request{
		Version: "2.0",
		Method:  m,
		Params:  p,
		ID:      uint64(rand.Int63()),
	}
	d, _ := json.Marshal(r)

	return bytes.NewReader(d)
}

// response represents a JSON-RPC response returned to a client.
type response struct {
	Version string           `json:"jsonrpc"`
	Result  *json.RawMessage `json:"result"`
	Error   *json.RawMessage `json:"error"`
}

// decodeResponse decodes stuff from a JSON-RPC client request.
func decodeResponse(r io.Reader, t interface{}) error {
	var res response
	if err := json.NewDecoder(r).Decode(&res); err != nil {
		return err
	}

	if res.Error != nil {
		// TODO: Read out error.
		return fmt.Errorf("decode: Result is an error")
	}
	if res.Result == nil {
		return fmt.Errorf("decode: Result is null")
	}

	return json.Unmarshal(*res.Result, t)
}

// request requests and parses JSON from the Monero wallet RPC client into a
// specified interface.
func (w *Wallet) request(m string, p, t interface{}) error {
	c := digest.New(context.Background(), w.Username, w.Password)

	req, err := http.NewRequest("POST", w.URL, encodeRequest(m, p))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

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

// Address returns the wallet's address, a 95-character hex address string of
// the monero-wallet-rpc in session.
func (w *Wallet) Address() (string, error) {
	var t = struct {
		Address string `json:"address"`
	}{}
	if err := w.request("getaddress", nil, &t); err != nil {
		return "", err
	}

	return t.Address, nil
}

// Balance represents the values returned by `Balance`.
type Balance struct {
	// The total balance of the current monero-wallet-rpc in session.
	Balance float64

	// Funds that are sufficiently deep enough in the Monero blockchain to be
	// considered safe to spend.
	UnBalance float64
}

// Balance returns the wallet's balance.
func (w *Wallet) Balance() (Balance, error) {
	var t = struct {
		Balance   uint64 `json:"balance"`
		UnBalance uint64 `json:"unlocked_balance"`
	}{}
	if err := w.request("getbalance", nil, &t); err != nil {
		return Balance{}, err
	}

	return Balance{
		float64(t.Balance) / 1.e+12,
		float64(t.UnBalance) / 1.e+12,
	}, nil
}

// Height returns the wallet's current block height.
func (w *Wallet) Height() (int64, error) {
	var t = struct {
		Height int64 `json:"height"`
	}{}
	if err := w.request("getheight", nil, &t); err != nil {
		return 0, err
	}

	return t.Height, nil
}

// Transfer represents the values returned by `incoming_transfers`.
type Transfer struct {
	// The amount of the transfer.
	Amount float64

	// If the transfer has been spent.
	Spent bool

	// The has of the transaction, several incoming transfers may share the same
	// hash if they were in the same transaction.
	Hash string

	// The size of the transaction in kB.
	Size uint64
}

// IncomingTransfers returns a list of incoming transfers to the wallet.
func (w *Wallet) IncomingTransfers() ([]Transfer, error) {
	var t = struct {
		Transfers []struct {
			Amount uint64 `json:"amount"`
			Spent  bool   `json:"spent"`
			TxHash string `json:"tx_hash"`
			TxSize uint64 `json:"tx_size"`
		} `json:"transfers"`
	}{}
	if err := w.request("incoming_transfers", struct {
		TransferType string `json:"transfer_type"`
	}{
		"all",
	}, &t); err != nil {
		return []Transfer{}, err
	}

	var tr []Transfer
	for _, p := range t.Transfers {
		tr = append(tr, Transfer{
			float64(p.Amount) / 1.e+12,
			p.Spent,
			p.TxHash,
			p.TxSize,
		})
	}

	return tr, nil
}
