package main

import (
	"encoding/json"
	"math"
	"net/http"
	"reflect"
	"time"

	"github.com/olihawkins/decimals"
)

// CryptoCompare is a stuct with all CryptoCompare values.
type CryptoCompare struct {
	GraphTime  []int     `json:"CryptoCompareGraphTime"`
	GraphPrice []float64 `json:"CryptoCompareGraphPrice"`

	Price         string `json:"CryptoComparePrice"`
	ChangePercent string `json:"CryptoCompareChangePercent"`
	ChangePrice   string `json:"CryptoCompareChangePrice"`
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

	// Fetch 48 hours of price history.
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

	// Write history to CryptoCompare struct.
	for _, p := range histohour.Data {
		c.GraphTime = append(c.GraphTime, p.Time)
		c.GraphPrice = append(c.GraphPrice, p.Close)
	}

	// Refelected value of price, for later usage.
	p := reflect.Indirect(reflect.ValueOf(price)).FieldByName(
		config.Currency).Float()

	// Symbol to use.
	var sym string
	switch config.Currency {
	case "EUR":
		sym = "€"
	case "USD":
		sym = "$"
	}

	// Did the price go up or down?
	dir := "+"
	if p < c.GraphPrice[0] {
		dir = "-"
	}

	// Write other values to CryptoCompare struct.
	c.Price = sym + decimals.FormatFloat(p, 2)
	c.ChangePercent = dir + " " + decimals.FormatFloat(math.Abs(((p/
		c.GraphPrice[0])-1)*100), 2) + "%"
	c.ChangePrice = sym + decimals.FormatFloat(math.Abs(p-c.GraphPrice[0]), 2)

	return c, nil
}
