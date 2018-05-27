package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

func about(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(
		"static/html/layout.html",
		"static/html/about.html",
	)
	if err != nil {
		log.Print(err)
	}

	if err := t.Execute(w, "about"); err != nil {
		log.Print(err)
	}

	go aboutEvent()
}

func aboutEvent() {
	tick := time.NewTicker(8 * time.Second)
	defer tick.Stop()

	// TODO: Can I somehow do an instant first tick?
	go updateSidebar()
	go updatePrice()

	for {
		select {
		case <-tick.C:
			go updateSidebar()
			go updatePrice()
		case <-close:
			return
		}
	}
}
