package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	fmt.Println("Starting server at ws://127.0.0.1:8080/\nQuit the server with CONTROL-C.")
	hub := newHub()
	go hub.run()

	// handle requests received from desktop app
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		serveClient(hub, w, r)
	})

	var err error = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("ListenAndServe:", err)
	}
}
