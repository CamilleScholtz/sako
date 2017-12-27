package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
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

	r := mux.NewRouter()

	// Set root handler.
	r.HandleFunc("/", info)

	// Set various other handlers.
	r.HandleFunc("/info", info)
	r.HandleFunc("/settings", settings)
	r.HandleFunc("/about", about)

	// Handle WebSockets.
	r.HandleFunc("/socket", socket)

	// Set location of the static assets.
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/",
		http.FileServer(http.Dir("static"))))

	srv := &http.Server{
		Handler:      r,
		Addr:         config.Host,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
