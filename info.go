package main

import (
	"html/template"
	"log"
	"net/http"
)

type infoModel struct {
	Title   string
	Sidebar sidebar
	Coincap coincap
}

func info(w http.ResponseWriter, r *http.Request) {
	sb, err := sidebarValues()
	if err != nil {
		log.Print(err)
	}

	c, err := parseCoincap()
	if err != nil {
		log.Print(err)
	}

	model := infoModel{"sako / info [" + c.Current + "]", sb, c}

	t, err := template.ParseFiles("static/info.html")
	if err != nil {
		log.Print(err)
	}

	if err := t.Execute(w, model); err != nil {
		log.Print(err)
	}
}
