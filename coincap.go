package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/mbanzon/currency"
	"github.com/olihawkins/decimals"
)

// Coincap is a stuct with all Coincap values.
type Coincap struct {
	Date          []float64 `json:"CoincapDate"`
	Price         []float64 `json:"CoincapPrice"`
	Current       string    `json:"CoincapCurrent"`
	ChangePrice   string    `json:"CoincapChangePrice"`
	ChangePercent string    `json:"CoincapChangePercent"`
}

type coincapJSON struct {
	Price [][]float64 `json:"price"`
}

// get fetches and parses JSON from a specified URL.
func (j *coincapJSON) get(url string) error {
	c := &http.Client{Timeout: 10 * time.Second}

	r, err := c.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(&j)
}

// parseCoincap fetches `/history/1day/XMR` JSON from Coincap and converts this
// to usable data. In here we also decide what currency to use.
func parseCoincap() (Coincap, error) {
	c := Coincap{}

	j := coincapJSON{}
	if err := j.get("https://coincap.io/history/1day/XMR"); err != nil {
		return c, err
	}

	c.Date, c.Price = parseGraph(j.Price)

	current := c.Price[len(c.Price)-1]
	changePrice := c.Price[0] - current
	changePercent := (changePrice / c.Price[0]) * 100

	dir := "+"
	if c.Price[0] > c.Price[len(c.Price)-1] {
		dir = "-"
	}

	var sym string
	switch config.Currency {
	case "USD":
		sym = "$"
	case "EUR":
		sym = "€"

		ecb, err := currency.NewConverter()
		if err != nil {
			return c, err
		}
		scc, err := ecb.GetSingleCurrencyConverter("USD", "EUR")
		if err != nil {
			return c, err
		}
		current = scc.Convert(current)
		changePrice = scc.Convert(changePrice)
		changePercent = scc.Convert(changePercent)
	default:
		return c, fmt.Errorf("coincap: %s is not a valid currency",
			config.Currency)
	}

	c.Current = sym + decimals.FormatFloat(current, 2)
	c.ChangePrice = dir + sym + decimals.FormatFloat(changePrice, 2)
	c.ChangePercent = dir + decimals.FormatFloat(changePercent, 2) + "%"

	return c, nil
}

// parseGraph converts the Coincap history data into something we can use with
// Chart.js.
func parseGraph(j [][]float64) ([]float64, []float64) {
	var d, p []float64
	for i := 0; i < len(j); i++ {
		// Smooth graph by removing data.
		// TODO: Check if the `-1` is needed.
		if i%8 == 0 || i == len(j)-1 {
			d = append(d, j[i][0])
			p = append(p, j[i][1])
		}
	}

	return d, p
}
