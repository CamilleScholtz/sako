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

	if err := t.Execute(w, "history"); err != nil {
		log.Print(err)
	}

	mel.HandleConnect(handleConnectHistory)
}

func updateHistory(s *melody.Session) error {
	price, err := cryptoComparePrice()
	if err != nil {
		return err
	}

	rt, err := wallet.IncomingTransfers()
	if err != nil {
		return err
	}
	var transfers []monero.Transfer
	for i := len(rt) - 1; i >= 0; i-- {
		transfers = append(transfers, rt[i])
	}

	msg, err := json.Marshal(struct {
		Type      string
		Price     Price
		Transfers []monero.Transfer
	}{
		"history", price, transfers,
	})
	if err != nil {
		return err
	}
	s.Write(msg)

	return nil
}

func handleConnectHistory(s *melody.Session) {
	defer s.Close()

	if err := updateSidebar(s); err != nil {
		log.Println(err)
		return
	}
	if err := updateHistory(s); err != nil {
		log.Println(err)
		return
	}

	go func() {
		fastTicker := time.NewTicker(5 * time.Second)
		slowTicker := time.NewTicker(20 * time.Second)
		defer func() {
			fastTicker.Stop()
			slowTicker.Stop()
		}()

		for {
			select {
			case <-fastTicker.C:
				if err := updateSidebar(s); err != nil {
					log.Println(err)
					return
				}
			case <-slowTicker.C:
				if err := updateHistory(s); err != nil {
					log.Println(err)
					return
				}
			}
		}
	}()
}
