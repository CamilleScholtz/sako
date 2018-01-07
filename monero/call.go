package monero

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
)

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
func encodeRequest(m string, p interface{}) []byte {
	r := &request{
		Version: "2.0",
		Method:  m,
		Params:  p,
		ID:      uint64(rand.Int63()),
	}
	d, _ := json.Marshal(r)

	return d
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
