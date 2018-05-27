package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// TODO: Possibly use an universal model.
type settingsModel struct {
	string
	Config Config
}

func settings(w http.ResponseWriter, r *http.Request) {
	// Handle POST requests.
	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			log.Fatal(err)
		}

		l, err := locateConfig()
		if err != nil {
			log.Fatal(err)
		}

		f, err := ioutil.ReadFile(l)
		if err != nil {
			log.Fatal(err)
		}

		f = modifyConfig(f, "currency", r.FormValue("currency"))

		if err := ioutil.WriteFile(l, f, 0644); err != nil {
			log.Fatal(err)
		}

		// Re-parse the config file.
		if err := parseConfig(); err != nil {
			log.Fatal(err)
		}
	}

	t, err := template.ParseFiles(
		"static/html/layout.html",
		"static/html/settings.html",
	)
	if err != nil {
		log.Print(err)
	}

	if err := t.Execute(w, "settings"); err != nil {
		log.Print(err)
	}

	go settingsEvent()
}

func settingsEvent() {
	tick := time.NewTicker(8 * time.Second)
	defer tick.Stop()

	// TODO: Can I somehow do an instant first tick?
	go updateSidebar()
	go updatePrice()
	go updateSettings()

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

func updateSettings() {
	event <- Event{"settings", config}
}
