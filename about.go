package main

import (
	"html/template"
	"log"
	"net/http"
)

type aboutModel struct {
	Template string
	Sidebar  Sidebar
	Coincap  Coincap
}

func about(w http.ResponseWriter, r *http.Request) {
	sb, err := parseSidebar()
	if err != nil {
		log.Print(err)
	}

	c, err := parseCoincap()
	if err != nil {
		log.Print(err)
	}

	model := aboutModel{"about", sb, c}

	t, err := template.ParseFiles(
		"static/templates/layout.html",
		"static/templates/head.html",
		"static/templates/sidebar.html",
		"static/templates/about.html",
		"static/templates/settings.js",
	)
	if err != nil {
		log.Print(err)
	}

	if err := t.Execute(w, model); err != nil {
		log.Print(err)
	}
}
