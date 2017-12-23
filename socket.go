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
)

func reader(ws *websocket.Conn) {
	defer ws.Close()

	ws.SetReadLimit(512)
	ws.SetReadDeadline(time.Now().Add(60 * time.Second))
	ws.SetPongHandler(func(string) error {
		ws.SetReadDeadline(time.Now().Add(60 * time.Second))
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
	coincapTicker := time.NewTicker(15 * time.Second)
	pingTicker := time.NewTicker(54 * time.Second)

	defer func() {
		coincapTicker.Stop()
		pingTicker.Stop()

		ws.Close()
	}()

	for {
		select {
		case <-coincapTicker.C:
			c, err := coincapValues()
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
