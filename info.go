package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

func info(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(
		"static/html/layout.html",
		"static/html/info.html",
	)
	if err != nil {
		log.Print(err)
	}

	if err := t.Execute(w, "info"); err != nil {
		log.Print(err)
	}

	go infoEvent()
}

func infoEvent() {
	tick := time.NewTicker(8 * time.Second)
	slow := time.NewTicker(64 * time.Second)
	defer func() {
		tick.Stop()
		slow.Stop()
	}()

	// TODO: Can I somehow do an instant first tick?
	go updateSidebar()
	go updateGraph()
	go updateSubmissions()
	go updateFunding()

	for {
		select {
		case <-tick.C:
			go updateSidebar()
			go updateGraph()
		case <-slow.C:
			go updateSubmissions()
			go updateFunding()
		case <-close:
			return
		}
	}
}

func updateGraph() {
	var err error
	msg := struct {
		Price Price
		XMR   Graph
		BTC   Graph
		ETH   Graph
	}{}

	// TODO: Only use this for the graph and not for the title?
	msg.Price, err = cryptoComparePrice("XMR")
	if err != nil {
		log.Print(err)
	}
	msg.XMR, err = cryptoCompareGraph("XMR")
	if err != nil {
		log.Print(err)
	}
	msg.BTC, err = cryptoCompareGraph("BTC")
	if err != nil {
		log.Print(err)
	}
	msg.ETH, err = cryptoCompareGraph("ETH")
	if err != nil {
		log.Print(err)
	}

	event <- Event{"graph", msg}
}

func updateSubmissions() {
	msg, err := redditSubmissions("monero")
	if err != nil {
		log.Print(err)
	}

	event <- Event{"submissions", msg}
}

func updateFunding() {
	msg, err := getMoneroFunding()
	if err != nil {
		log.Print(err)
	}

	event <- Event{"funding", msg}
}
