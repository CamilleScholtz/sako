package main

import (
	"time"

	coinmarketcap "github.com/miguelmota/go-coinmarketcap"
)

type graph struct {
	Date  []string  `json:"graph_Date"`
	Price []float64 `json:"graph_Price"`
}

func graphValues(days int64) (graph, error) {
	g := graph{}

	t := time.Now().Unix()
	gd, err := coinmarketcap.GetCoinGraphData("monero", t-(60*60*24*days), t)
	if err != nil {
		return g, err
	}

	for i, p := range gd.PriceUsd {
		if i%8 == 0 || i == len(gd.PriceUsd)-1 {
			g.Date = append(g.Date, time.Unix(int64(p[0]/1000), 0).UTC().Format(
				time.RFC3339))
			g.Price = append(g.Price, p[1])
		}
	}

	return g, nil
}
