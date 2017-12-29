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

	if err := t.Execute(w, "info"); err != nil {
		log.Print(err)
	}

	mel.HandleConnect(updateInfo)
}

func updateInfo(s *melody.Session) {
	go func() {
		t := time.NewTicker(15 * time.Second)
		defer t.Stop()

		// TODO: The continue here could result in an endless loop.
		for {
			sidebar, err := sidebar()
			if err != nil {
				log.Print(err)
				continue
			}

			graph, err := cryptoGraph()
			if err != nil {
				log.Print(err)
				continue
			}

			price, err := cryptoPrice()
			if err != nil {
				log.Print(err)
				continue
			}

			msg, err := json.Marshal(struct {
				Sidebar Sidebar
				Price   Price
				Graph   Graph
			}{
				sidebar, price, graph,
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
