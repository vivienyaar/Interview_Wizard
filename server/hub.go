package main

import (
	"fmt"
	"time"
)

type Hub struct {
	registerClient   chan *Client
	unregisterClient chan *Client
	clients          map[*Client]bool
}

func newHub() *Hub {
	return &Hub{
		registerClient:   make(chan *Client),
		unregisterClient: make(chan *Client),
		clients:          make(map[*Client]bool),
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.registerClient:
			fmt.Println("["+time.Now().Format(time.RFC850)+"] Register Client:", client.id.String())
			h.clients[client] = true
			newMessage := Message{
				Type:     "registerSucc",
				ClientId: client.id.String(),
			}
			client.send <- &newMessage
		case client := <-h.unregisterClient:
			if _, ok := h.clients[client]; ok {
				fmt.Println("["+time.Now().Format(time.RFC850)+"] Unregister Client:", client.id.String())
				delete(h.clients, client)
				close(client.send)
			}
		}
	}
}
