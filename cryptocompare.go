package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
)

// getAndUnmarshal fetches and parses JSON from a specified URL.
func getAndUnmarshal(url string, t interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	if r.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected http GET status: %s", r.Status)
	}

	return json.NewDecoder(r.Body).Decode(t)
}

// Graph is a stuct with all the values needed for a graph.
type Graph struct {
	Time  []int
	Value []float64
}

func cryptoGraph() (Graph, error) {
	c := Graph{}

	var histohour = struct {
		Data []struct {
			Close float64 `json:"close"`
			Time  int     `json:"time"`
		} `json:"Data"`
	}{}
	if err := getAndUnmarshal(
		"https://min-api.cryptocompare.com/data/histohour?fsym=XMR&limit=48&tsym="+
			config.Currency, &histohour); err != nil {
		return c, err
	}

	for _, i := range histohour.Data {
		c.Time = append(c.Time, i.Time)
		c.Value = append(c.Value, i.Close)
	}

	return c, nil
}

// Price is a stuct with all the values needed for a price.
type Price struct {
	Symbol string
	Value  float64
}

func cryptoPrice() (Price, error) {
	p := Price{}

	var price = struct {
		USD float64 `json:"USD"`
		EUR float64 `json:"EUR"`
	}{}
	if err := getAndUnmarshal(
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
