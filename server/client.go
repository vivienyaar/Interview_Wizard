package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/satori/go.uuid"
	"net/http"
	"time"
)

type Client struct {
	hub  *Hub
	conn *websocket.Conn
	send chan *Message
	id   uuid.UUID
}

type Message struct {
	Type     string `json:"type"`
	ClientId string `json:"clientId"`
	// TODO
}

func newClient(hub *Hub, conn *websocket.Conn) *Client {
	u0, err := uuid.NewV4()
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		return nil
	}
	return &Client{
		hub:  hub,
		conn: conn,
		send: make(chan *Message),
		id:   u0,
	}
}

func closeClientConn(c *Client) {
	c.hub.unregisterClient <- c
	c.conn.Close()
}

// Transfers inbound messages from the websocket to the hub
func (c *Client) readPump() {
	defer closeClientConn(c)
	for {
		_, msg, err := c.conn.ReadMessage()

		if err != nil {
			fmt.Println("Can't receive!")
			break
		}

		fmt.Println("[" + time.Now().Format(time.RFC850) + "] Received from Client (" + c.id.String() + ")")
		var msgIn Message
		if err = json.Unmarshal(msg, &msgIn); err != nil {
			fmt.Println("unmarshal error:", err)
		}

		// TODO: handle different client requests here
		/* switch msgIn.Type {
		case :
		case :
		} */

	}
}

func (c *Client) writePump() {
	defer closeClientConn(c)
	for {
		select {
		case msg, _ := <-c.send:
			fmt.Println("[" + time.Now().Format(time.RFC850) + "] Sending to client(" + c.id.String() + ")")
			msgByte, err := json.Marshal(msg)
			if err != nil {
				fmt.Println("marshal error:", err)
			}
			c.conn.WriteMessage(websocket.TextMessage, msgByte)
		}
	}
}

func serveClient(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("upgrader error:", err)
		return
	}
	client := newClient(hub, conn)
	client.hub.registerClient <- client

	go client.writePump()
	go client.readPump()
}
