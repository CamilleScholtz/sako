package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gabstv/go-monero/walletrpc"
	"github.com/gorilla/mux"
	"github.com/olahol/melody"
	"github.com/onodera-punpun/sako/monero"
)

var (
	wallet walletrpc.Client
	// TODO: Replace this with daemonrpc
	daemon *monero.Daemon
	mel    = melody.New()
)

func main() {
	if err := parseConfig(); err != nil {
		log.Fatal(err)
	}

	wallet = walletrpc.New(walletrpc.Config{
		Address: "http://" + config.RPC + "/json_rpc",
	})
	daemon = monero.NewDaemon("http://" + config.Daemon + "/json_rpc")

	r := mux.NewRouter()

	// Handle pages.
	r.HandleFunc("/", info)
	r.HandleFunc("/info", info)
	r.HandleFunc("/history", history)
	//r.HandleFunc("/settings", settings)
	//r.HandleFunc("/about", about)

	// Handle WebSockets.
	r.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		mel.HandleRequest(w, r)
	})

	// Handle static assets.
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/",
		http.FileServer(http.Dir("static"))))

	s := &http.Server{
		Handler:      r,
		Addr:         config.Host,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(s.ListenAndServe())
}
