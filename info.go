package main

import (
	"html/template"
	"log"
	"net/http"

	coinmarketcap "github.com/miguelmota/go-coinmarketcap"
)

type infoModel struct {
	Title   string
	Sidebar sidebar
	Graph   graph
	Price   float64
	Change  float64
}

func info(w http.ResponseWriter, r *http.Request) {
	sb, err := sidebarValues()
	if err != nil {
		log.Print(err)
	}

	g, err := graphValues(1)
	if err != nil {
		log.Print(err)
	}

	c, err := coinmarketcap.GetCoinData("monero")
	if err != nil {
		log.Print(err)
	}

	model := infoModel{
		"sako",
		sb,
		g,
		c.PriceUsd,
		c.PercentChange24h,
	}

	t, err := template.ParseFiles("assets/info.html")
	if err != nil {
		log.Print(err)
	}

	if err := t.Execute(w, model); err != nil {
		log.Print(err)
	}
}
