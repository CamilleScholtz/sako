package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/olahol/melody"
)

func info(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(
		"static/templates/layout.html",
		"static/templates/sidebar.html",
		"static/templates/info.html",
	)
	if err != nil {
		log.Print(err)
	}

	sidebar, err := sidebar()
	if err != nil {
		log.Print(err)
		return
	}

	if err := t.Execute(w, struct {
		Template string
		Sidebar  Sidebar
	}{
		"info", sidebar,
	}); err != nil {
		log.Print(err)
	}

	mel.HandleConnect(updateInfo)
}

func updateInfo(s *melody.Session) {
	go func() {
		t := time.NewTicker(15 * time.Second)
		defer t.Stop()

		for {
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

			msg, err := json.Marshal(struct {
				Price Price
				Graph Graph
			}{
				price, graph,
			})
			if err != nil {
				log.Print(err)
				return
			}
			s.Write(msg)

			<-t.C
		}
	}()
}
