package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

const (
	// Time allowed to write the file to the client.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the client.
	pongWait = 60 * time.Second

	// Send pings to client with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10
)

func reader(ws *websocket.Conn) {
	defer ws.Close()

	ws.SetReadLimit(512)
	ws.SetReadDeadline(time.Now().Add(pongWait))
	ws.SetPongHandler(func(string) error {
		ws.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		_, _, err := ws.ReadMessage()
		if err != nil {
			break
		}
	}
}

func writer(ws *websocket.Conn) {
	cryptoCompareTicker := time.NewTicker(15 * time.Second)
	pingTicker := time.NewTicker(pingPeriod)

	defer func() {
		cryptoCompareTicker.Stop()
		pingTicker.Stop()

		ws.Close()
	}()

	for {
		select {
		case <-cryptoCompareTicker.C:
			c, err := cryptoCompare()
			if err != nil {
				log.Printf("closing socket: %s", err)
				return
			}

			// TODO: Is this deadline needed?
			ws.SetWriteDeadline(time.Now().Add(writeWait))
			if err := ws.WriteJSON(c); err != nil {
				log.Printf("closing socket: %s", err)
				return
			}
		case <-pingTicker.C:
			ws.SetWriteDeadline(time.Now().Add(writeWait))
			if err := ws.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}

func socket(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		if _, ok := err.(websocket.HandshakeError); !ok {
			log.Println(err)
		}
		return
	}

	go writer(ws)
	reader(ws)
}
