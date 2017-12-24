package main

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/BurntSushi/toml"
)

// TODO: Possibly use an universal model.
type settingsModel struct {
	Template string
	Sidebar  Sidebar
	Coincap  Coincap
	Config   Config
}

func settings(w http.ResponseWriter, r *http.Request) {
	// Handle POST requests.
	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}

		// Set new values.
		config.Currency = r.FormValue("currency")

		d, err := os.Getwd()
		if err != nil {
			log.Print(err)
		}

		// Write to config file.
		// TODO: I can't seem to use the encode function without overwriting
		// formatting and comments.
		buf := new(bytes.Buffer)
		if err := toml.NewEncoder(buf).Encode(config); err != nil {
			log.Print(err)
		}
		ioutil.WriteFile(path.Join(d, "runtime", "config.toml"), buf.Bytes(),
			0664)

		// Re-parse config file.
		if err := parseConfig(); err != nil {
			log.Fatal(err)
		}
	}

	sb, err := parseSidebar()
	if err != nil {
		log.Print(err)
	}

	c, err := parseCoincap()
	if err != nil {
		log.Print(err)
	}

	model := settingsModel{"settings", sb, c, config}

	t, err := template.ParseFiles(
		"static/templates/layout.html",
		"static/templates/head.html",
		"static/templates/sidebar.html",
		"static/templates/settings.html",
		"static/templates/settings.js",
	)
	if err != nil {
		log.Print(err)
	}

	if err := t.Execute(w, model); err != nil {
		log.Print(err)
	}
}
