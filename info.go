package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/olahol/melody"
)

// Info is a stuct with all the values needed in the info templates.
type Info struct {
	Type     string
	Price    Price
	GraphXMR Graph
	GraphBTC Graph
	GraphETH Graph
}

func info(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(
		"static/html/layout.html",
		"static/html/sidebar.html",
		"static/html/info.html",
	)
	if err != nil {
		log.Print(err)
	}

	if err := t.Execute(w, "info"); err != nil {
		log.Print(err)
	}

	mel.HandleConnect(handleConnectInfo)
}

func handleConnectInfo(s *melody.Session) {
	go func() {
		t := time.NewTicker(10 * time.Second)
		defer t.Stop()

		for {
			if s.IsClosed() {
				return
			}

			go updateSidebar(s)
			go updateInfo(s)

			<-t.C
		}
	}()
}

func updateInfo(s *melody.Session) {
	data := Info{Type: "info"}
	var err error

	data.Price, err = cryptoComparePrice("XMR")
	if err != nil {
		log.Print(err)
	}

	data.GraphXMR, err = cryptoCompareGraph("XMR")
	if err != nil {
		log.Print(err)
	}
	data.GraphBTC, err = cryptoCompareGraph("BTC")
	if err != nil {
		log.Print(err)
	}
	data.GraphETH, err = cryptoCompareGraph("ETH")
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
