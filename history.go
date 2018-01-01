package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/olahol/melody"
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

	if err := t.Execute(w, "history"); err != nil {
		log.Print(err)
	}

	mel.HandleConnect(historyInfo)
}

func historyInfo(s *melody.Session) {
	go func() {
		t := time.NewTicker(15 * time.Second)
		defer t.Stop()

		for {
			sidebar, err := sidebar()
			if err != nil {
				log.Print(err)
				return
			}

			graph, err := cryptoCompareGraph()
			if err != nil {
				log.Print(err)
				return
			}

			price, err := cryptoComparePrice()
			if err != nil {
				log.Print(err)
				return
			}

			transfers, err := walletIncomingTransfers()
			if err != nil {
				log.Print(err)
				return
			}

			msg, err := json.Marshal(struct {
				Sidebar   Sidebar
				Price     Price
				Graph     Graph
				Transfers uint64
			}{
				sidebar, price, graph, transfers,
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
