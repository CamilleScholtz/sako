package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/olahol/melody"
	"github.com/sunrisedo/monero"
)

var (
	daemon *monero.DaemonClient
	wallet *monero.WalletClient

	mel = melody.New()
)

func main() {
	if err := parseConfig(); err != nil {
		log.Fatal(err)
	}

	// Listen to the Monero deamon and wallet.
	daemon = monero.NewDaemonClient("http://" + config.Daemon + "/json_rpc")
	wallet = monero.NewWalletClient("http://"+config.RPC+"/json_rpc",
		config.Username, config.Password)

	r := mux.NewRouter()

	r.HandleFunc("/", info)
	r.HandleFunc("/info", info)
	r.HandleFunc("/info-ws", func(w http.ResponseWriter, r *http.Request) {
		mel.HandleRequest(w, r)
	})

	r.HandleFunc("/history", history)
	r.HandleFunc("/history-ws", func(w http.ResponseWriter, r *http.Request) {
		mel.HandleRequest(w, r)
	})

	//r.HandleFunc("/settings", settings)
	//r.HandleFunc("/settings-ws", func(w http.ResponseWriter, r *http.Request) {
	//	m.HandleRequest(w, r)
	//})

	//r.HandleFunc("/about", about)
	//r.HandleFunc("/about-ws", func(w http.ResponseWriter, r *http.Request) {
	//	m.HandleRequest(w, r)
	//})

	// Set location of the static assets.
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
