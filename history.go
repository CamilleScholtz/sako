package main

import (
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/gabstv/go-monero/walletrpc"
)

func history(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(
		"static/html/layout.html",
		"static/html/history.html",
	)
	if err != nil {
		log.Print(err)
	}

	if err := t.Execute(w, "history"); err != nil {
		log.Print(err)
	}

	go historyEvent()
}

func historyEvent() {
	tick := time.NewTicker(8 * time.Second)
	defer tick.Stop()

	// TODO: Can I somehow do an instant first tick?
	go updateSidebar()
	go updatePrice()
	go updateHistory()

	for {
		select {
		case <-tick.C:
			go updateSidebar()
			go updatePrice()
			go updateHistory()
		case <-close:
			return
		}
	}
}

func updateHistory() {
	var err error
	msg := struct {
		Price     Price
		Transfers *walletrpc.GetTransfersResponse
	}{}

	msg.Price, err = cryptoComparePrice("XMR")
	if err != nil {
		log.Print(err)
	}

	msg.Transfers, err = wallet.GetTransfers(walletrpc.GetTransfersRequest{
		In:      true,
		Out:     true,
		Pending: true,
		Failed:  true,
	})
	if err != nil {
		log.Print(err)
	}

	event <- Event{"history", msg}
}
