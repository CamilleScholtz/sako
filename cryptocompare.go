package main

import (
	"encoding/json"
	"net/http"
	"reflect"
	"time"

	"github.com/olihawkins/decimals"
)

// CryptoCompare is a stuct with all CryptoCompare values.
type CryptoCompare struct {
	GraphTime  []int     `json:"CryptoCompareGraphTime"`
	GraphPrice []float64 `json:"CryptoCompareGraphPrice"`

	Symbol        string  `json:"CryptoCompareSymbol"`
	Price         float64 `json:"CryptoComparePrice"`
	ChangePercent float64 `json:"CryptoCompareChangePercent"`
	ChangePrice   float64 `json:"CryptoCompareChangePrice"`
}

// getAndUnmarshal fetches and parses JSON from a specified URL.
func getAndUnmarshal(url string, t interface{}) error {
	c := &http.Client{Timeout: 10 * time.Second}

	r, err := c.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(t)
}

func cryptoCompare() (CryptoCompare, error) {
	c := CryptoCompare{}

	// Fetch 24 hour history.
	var histoday = struct {
		Data []struct {
			Close float64 `json:"close"`
			Time  int     `json:"time"`
		} `json:"Data"`
	}{}
	if err := getAndUnmarshal(
		"https://min-api.cryptocompare.com/data/histoday?fsym=XMR&limit=45&tsym="+
			config.Currency, &histoday); err != nil {
		return c, err
	}

	// Fetch current price.
	var price = struct {
		USD float64 `json:"USD"`
		EUR float64 `json:"EUR"`
	}{}
	if err := getAndUnmarshal(
		"https://min-api.cryptocompare.com/data/price?fsym=XMR&tsyms="+
			config.Currency, &price); err != nil {
		return c, err
	}

	for _, p := range histoday.Data {
		c.GraphTime = append(c.GraphTime, p.Time)
		c.GraphPrice = append(c.GraphPrice, p.Close)
	}

	switch config.Currency {
	case "EUR":
		c.Symbol = "€"
	case "USD":
		c.Symbol = "$"
	}
	// TODO: These values sometimes only display one decimal, I don't want that.
	c.Price = reflect.Indirect(reflect.ValueOf(price)).FieldByName(
		config.Currency).Float()
	c.ChangePercent = decimals.RoundFloat(((c.GraphPrice[len(c.GraphPrice)-1]/
		c.Price)-1)*100, 2)
	c.ChangePrice = decimals.RoundFloat(c.Price-c.GraphPrice[len(c.GraphPrice)-
		1], 2)

	return c, nil
}
