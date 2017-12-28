package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/olahol/melody"
)

type infoData struct {
	Template string
	Sidebar  Sidebar
}

func info(w http.ResponseWriter, r *http.Request) {
	var err error

	d := infoData{Template: "info"}
	d.Sidebar, err = sidebar()
	if err != nil {
		log.Print(err)
	}

	t, err := template.ParseFiles(
		"static/templates/layout.html",
		"static/templates/sidebar.html",
		"static/templates/info.html",
		"static/templates/info.js",
	)
	if err != nil {
		log.Print(err)
	}

	if err := t.Execute(w, d); err != nil {
		log.Print(err)
	}

	mel.HandleConnect(updateInfo)
}

func updateInfo(s *melody.Session) {
	go func() {
		t := time.NewTicker(15 * time.Second)
		defer t.Stop()

		for {
			cgt, cgp, err := cryptoGraph()
			if err != nil {
				log.Print(err)
				return
			}

			cs, err := cryptoSymbol()
			if err != nil {
				log.Print(err)
				return
			}

			cp, err := cryptoPrice()
			if err != nil {
				log.Print(err)
				return
			}

			msg, err := json.Marshal(struct {
				Symbol     string
				Price      float64
				GraphTime  []int
				GraphPrice []float64
			}{
				cs, cp, cgt, cgp,
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
