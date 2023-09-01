package main

import (
	"log"

	"github.com/gorilla/websocket"
)

type ClientList map[*Client]bool

type Client struct {
	conn    *websocket.Conn
	manager *Manager
	engres  chan []byte
}

func newClient(conn *websocket.Conn, manager *Manager) *Client {
	return &Client{
		conn:    conn,
		manager: manager,
		engres:  make(chan []byte),
	}
}

func (c *Client) ReadMessage() {
	defer func() {
		c.manager.RemoveClient(c)
	}()
	for {
		messageType, payload, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Println("Error while reading message", err)
			}
			break
		}
		log.Println(messageType, string(payload))

		log.Println("Broadcasting the message")
		// c.manager.Clients.
		for el := range c.manager.Clients {
			el.engres <- payload
		}
	}
}

func (c *Client) WriteMessage() {
	defer func() {
		c.manager.RemoveClient(c)
	}()

	for {
		select {
		case message, ok := <-c.engres:
			{
				if !ok {
					if err := c.conn.WriteMessage(websocket.CloseMessage, nil); err != nil {
						log.Println("Connection already closed", err)
					}
					return
				}
				if err := c.conn.WriteMessage(websocket.TextMessage, message); err != nil {
					log.Println("Failed to send the message", err)
				}
				log.Println("Message sent")
			}
		}
	}

}
