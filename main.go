package main

import (
	"log"
	"net/http"

	"github.com/sunrisedo/monero"
)

var daemon *monero.DaemonClient
var wallet *monero.WalletClient

func main() {
	if err := parseConfig(); err != nil {
		log.Fatal(err)
	}

	daemon = monero.NewDaemonClient("http://" + config.Daemon + "/json_rpc")
	wallet = monero.NewWalletClient("http://"+config.RPC+"/json_rpc",
		config.Username, config.Password)

	// Set root handler.
	http.HandleFunc("/", info)

	// Set various other handlers.
	http.HandleFunc("/info", info)
	http.HandleFunc("/settings", settings)
	http.HandleFunc("/about", about)

	// Handle WebSockets.
	http.HandleFunc("/socket", socket)

	// Set location of our assets.
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(
		http.Dir("static"))))

	if err := http.ListenAndServe(config.Host, nil); err != nil {
		log.Fatal(err)
	}
}
