package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/SlyMarbo/rss"
)

type infoModel struct {
	Template      string
	Sidebar       Sidebar
	CryptoCompare CryptoCompare
	Feed          []string
}

func info(w http.ResponseWriter, r *http.Request) {
	sb, err := sidebar()
	if err != nil {
		log.Print(err)
	}

	c, err := cryptoCompare()
	if err != nil {
		log.Print(err)
	}

	f, err := rss.Fetch("http://monero-observer.com/feed.rss")
	if err != nil {
		log.Print(err)
	}

	var feed []string
	for _, i := range f.Items {
		feed = append(feed, i.Title)
	}

	model := infoModel{"info", sb, c, feed}

	t, err := template.ParseFiles(
		"static/templates/layout.html",
		"static/templates/head.html",
		"static/templates/sidebar.html",
		"static/templates/info.html",
		"static/templates/info.js",
	)
	if err != nil {
		log.Print(err)
	}

	if err := t.Execute(w, model); err != nil {
		log.Print(err)
	}
}
