package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/olahol/melody"
	"github.com/onodera-punpun/sako/monero"
)

func history(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(
		"static/templates/layout.html",
		"static/templates/sidebar.html",
		"static/templates/history.html",
	)
	if err != nil {
		log.Print(err)
	}

	sidebar, err := sidebar()
	if err != nil {
		log.Print(err)
		return
	}

	reverseTransfers, err := wallet.IncomingTransfers()
	if err != nil {
		log.Print(err)
		return
	}
	var transfers []monero.Transfer
	for i := len(reverseTransfers) - 1; i >= 0; i-- {
		transfers = append(transfers, reverseTransfers[i])
	}

	if err := t.Execute(w, struct {
		Template  string
		Sidebar   Sidebar
		Transfers []monero.Transfer
	}{
		"history", sidebar, transfers,
	}); err != nil {
		log.Print(err)
	}

	mel.HandleConnect(historyInfo)
}

func historyInfo(s *melody.Session) {
	go func() {
		t := time.NewTicker(15 * time.Second)
		defer t.Stop()

		for {
			price, err := cryptoComparePrice()
			if err != nil {
				log.Print(err)
				return
			}

			msg, err := json.Marshal(struct {
				Price Price
			}{
				price,
			})
			if err != nil {
				log.Print(err)
				continue
			}
			s.Write(msg)

			<-t.C
		}
	}()
}
