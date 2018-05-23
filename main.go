package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gabstv/go-monero/walletrpc"
	"github.com/gorilla/mux"
	"github.com/onodera-punpun/sako/monero"
	"github.com/onodera-punpun/sako/sse"
)

// Global variables.
var (
	wallet walletrpc.Client
	// TODO: Replace this with daemonrpc
	daemon *monero.Daemon

	event chan Event
	close chan bool
)

// Event describes an SSE event.
type Event struct {
	Name    string
	Message interface{}
}

func main() {
	if err := parseConfig(); err != nil {
		log.Fatal(err)
	}

	wallet = walletrpc.New(walletrpc.Config{
		Address: "http://" + config.RPC + "/json_rpc",
	})
	daemon = monero.NewDaemon("http://" + config.Daemon + "/json_rpc")

	event = make(chan Event)
	// TODO: Use context?
	close = make(chan bool)

	mux := mux.NewRouter()

	// Handle SSE.
	mux.HandleFunc("/event", func(w http.ResponseWriter, r *http.Request) {
		c, err := sse.Upgrade(w, r)
		if err != nil {
			log.Print(err)
		}

		for {
			select {
			case e := <-event:
				if err := c.WriteJSONEvent(e.Name, e.Message); err != nil {
					log.Print(e.Name, ": ", err)
				}
			case <-r.Context().Done():
				close <- true
				return
			}
		}
	})

	// Handle pages.
	mux.HandleFunc("/", info)
	mux.HandleFunc("/info", info)
	mux.HandleFunc("/history", history)
	//mux.HandleFunc("/settings", settings)
	//mux.HandleFunc("/about", about)

	// Handle static assets.
	mux.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.
		FileServer(http.Dir("static"))))

	srv := &http.Server{
		Handler:     mux,
		Addr:        config.Host,
		ReadTimeout: 15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
