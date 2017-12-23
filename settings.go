package main

import (
	"html/template"
	"log"
	"net/http"
)

type settingsModel struct {
	Title   string
	Sidebar sidebar
	Coincap coincap
}

func settings(w http.ResponseWriter, r *http.Request) {
	sb, err := sidebarValues()
	if err != nil {
		log.Print(err)
	}

	c, err := parseCoincap()
	if err != nil {
		log.Print(err)
	}

	model := infoModel{"sako / settings [" + c.Current + "]", sb, c}

	t, err := template.ParseFiles("static/settings.html")
	if err != nil {
		log.Print(err)
	}

	if err := t.Execute(w, model); err != nil {
		log.Print(err)
	}
}
