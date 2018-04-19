package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/gabstv/go-monero/walletrpc"
	"github.com/olahol/melody"
)

// History is a stuct with all the values needed in the history template.
type History struct {
	Type      string
	Price     Price
	Transfers *walletrpc.GetTransfersResponse
}

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

	mel.HandleConnect(handleConnectHistory)
}

func handleConnectHistory(s *melody.Session) {
	go func() {
		t := time.NewTicker(8 * time.Second)
		defer t.Stop()

		for {
			if s.IsClosed() {
				return
			}

			go updateLayout(s)
			go updateHistory(s)

			<-t.C
		}
	}()
}

func updateHistory(s *melody.Session) {
	data := History{Type: "history"}
	var err error

	data.Price, err = cryptoComparePrice("XMR")
	if err != nil {
		log.Print(err)
	}

	data.Transfers, err = wallet.GetTransfers(walletrpc.GetTransfersRequest{
		In:      true,
		Out:     true,
		Pending: true,
		Failed:  true,
	})
	if err != nil {
		log.Print(err)
	}

	msg, err := json.Marshal(data)
	if err != nil {
		log.Print(err)
		return
	}

	s.Write(msg)
}
