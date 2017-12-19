package main

import (
	coinmarketcap "github.com/miguelmota/go-coinmarketcap"
)

type price struct {
	Price  float64 `json:"price_Price"`
	Change float64 `json:"price_Change"`
}

func priceValues() (price, error) {
	p := price{}

	pd, err := coinmarketcap.GetCoinData("monero")
	if err != nil {
		return p, err
	}

	p.Price = pd.PriceUsd
	p.Change = pd.PercentChange24h

	return p, nil
}
