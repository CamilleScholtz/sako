package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"time"
)

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

func cryptoGraph() ([]int, []float64, error) {
	var histohour = struct {
		Data []struct {
			Close float64 `json:"close"`
			Time  int     `json:"time"`
		} `json:"Data"`
	}{}
	if err := getAndUnmarshal(
		"https://min-api.cryptocompare.com/data/histohour?fsym=XMR&limit=48&tsym="+
			config.Currency, &histohour); err != nil {
		return []int{}, []float64{}, err
	}

	var t []int
	var p []float64
	for _, i := range histohour.Data {
		t = append(t, i.Time)
		p = append(p, i.Close)
	}

	return t, p, nil
}

func cryptoPrice() (float64, error) {
	var price = struct {
		USD float64 `json:"USD"`
		EUR float64 `json:"EUR"`
	}{}
	if err := getAndUnmarshal(
		"https://min-api.cryptocompare.com/data/price?fsym=XMR&tsyms="+config.
			Currency, &price); err != nil {
		return 0, err
	}

	return reflect.Indirect(reflect.ValueOf(price)).FieldByName(config.
		Currency).Float(), nil
}

func cryptoSymbol() (string, error) {
	switch config.Currency {
	case "EUR":
		return "€", nil
	case "USD":
		return "$", nil
	}

	return "", fmt.Errorf("cryptoSymbol: No valid currency name")
}
