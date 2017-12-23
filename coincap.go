package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/olihawkins/decimals"
)

type coincap struct {
	Date          []float64 `json:"coincapDate"`
	Price         []float64 `json:"coincapPrice"`
	Current       string    `json:"coincapCurrent"`
	ChangePrice   string    `json:"coincapChangePrice"`
	ChangePercent string    `json:"coincapChangePercent"`
}

type coincapJSON struct {
	Price [][]float64 `json:"price"`
}

func (j *coincapJSON) get(url string) error {
	c := &http.Client{Timeout: 10 * time.Second}

	r, err := c.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(&j)
}

func coincapValues() (coincap, error) {
	c := coincap{}

	j := coincapJSON{}
	if err := j.get("https://coincap.io/history/1day/XMR"); err != nil {
		return c, err
	}

	for i := 0; i < len(j.Price); i++ {
		// Smooth graph by removing data.
		// TODO: Check if the `-1` is needed.
		if i%8 == 0 || i == len(j.Price)-1 {
			c.Date = append(c.Date, j.Price[i][0])
			c.Price = append(c.Price, j.Price[i][1])
		}
	}

	d := "+"
	if c.Price[0] > c.Price[len(c.Price)-1] {
		d = "-"
	}
	c.Current = "$" + decimals.FormatFloat(c.Price[len(c.Price)-1], 2)
	c.ChangePrice = d + "$" + decimals.FormatFloat(c.Price[0]-c.Price[len(c.Price)-1],
		2)
	c.ChangePercent = d + decimals.FormatFloat(((c.Price[0]-c.Price[len(
		c.Price)-1])/c.Price[0])*100, 2) + "%"

	return c, nil
}
