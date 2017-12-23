package main

import (
	"log"
	"net/http"

	"github.com/sunrisedo/monero"
)

var wallet = monero.NewWalletClient("http://127.0.0.1:18082/json_rpc",
	"onodera", "seekrit")

func main() {
	if err := parseConfig(); err != nil {
		log.Fatal(err)
	}

	// Create handle functions.
	http.HandleFunc("/", info)
	http.HandleFunc("/info", info)
	http.HandleFunc("/settings", settings)
	http.HandleFunc("/socket", socket)

	// Set location of our assets.
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(
		http.Dir("static"))))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
