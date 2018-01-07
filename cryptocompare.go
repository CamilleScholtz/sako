package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"time"
)

// cryptoCompareRequest requests and parses JSON from a specified URL into a
// specified interface.
func cryptoCompareRequest(url string, t interface{}) error {
	c := &http.Client{Timeout: time.Second * 5}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	// TODO: Use this?
	//req.Header.Set("Connection", "Keep-Alive")
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

	return json.NewDecoder(res.Body).Decode(t)
}

// Graph represents some of the values found in a CryptoCompare
// `/data/histohour` response.
type Graph struct {
	Time  []int
	Value []float64
}

// cryptoCompareGraph request and returns graphing information from the
// CryptoCompare API.
func cryptoCompareGraph() (Graph, error) {
	g := Graph{}

	var histohour = struct {
		Data []struct {
			Close float64 `json:"close"`
			Time  int     `json:"time"`
		} `json:"Data"`
	}{}
	if err := cryptoCompareRequest(
		"https://min-api.cryptocompare.com/data/histohour?fsym=XMR&limit=48&tsym="+
			config.Currency, &histohour); err != nil {
		return g, err
	}

	for _, i := range histohour.Data {
		g.Time = append(g.Time, i.Time)
		g.Value = append(g.Value, i.Close)
	}

	return g, nil
}

// Price represents some of the values found in a CryptoCompare `/data/price`
// response. It also includes a currency symbol.
type Price struct {
	Symbol string
	Value  float64
}

// cryptoComparePrice request and returns price information from the
// CryptoCompare API.
func cryptoComparePrice() (Price, error) {
	p := Price{}

	var price = struct {
		USD float64 `json:"USD"`
		EUR float64 `json:"EUR"`
	}{}
	if err := cryptoCompareRequest(
		"https://min-api.cryptocompare.com/data/price?fsym=XMR&tsyms="+config.
			Currency, &price); err != nil {
		return p, err
	}

	switch config.Currency {
	case "EUR":
		p.Symbol = "€"
	case "USD":
		p.Symbol = "$"
	}
	p.Value = reflect.Indirect(reflect.ValueOf(price)).FieldByName(config.
		Currency).Float()

	return p, nil
}
