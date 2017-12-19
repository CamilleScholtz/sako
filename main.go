package main

import (
	"log"
	"net/http"

	"github.com/sunrisedo/monero"
)

var wallet = monero.NewWalletClient("http://127.0.0.1:18082/json_rpc",
	"onodera", "seekrit")

func main() {
	// Create handle functions.
	http.HandleFunc("/", info)
	http.HandleFunc("/socket", socket)

	// Set location of our assets.
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(
		http.Dir("assets"))))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
