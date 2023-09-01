package main

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Manager struct {
	Clients ClientList
	sync.RWMutex
}

func MakeManager() *Manager {
	return &Manager{
		Clients: make(ClientList),
	}
}

func (M *Manager) serveWS(w http.ResponseWriter, r *http.Request) {
	log.Println("New Connection")

	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("failed to upgrade connection to full duplex")
	}

	client := newClient(conn, M)
	M.AddClient(client)

	go client.ReadMessage()
	go client.WriteMessage()
}

func (M *Manager) AddClient(c *Client) {
	M.Lock()
	defer M.Unlock()

	M.Clients[c] = true
}

func (M *Manager) RemoveClient(c *Client) {
	M.Lock()
	defer M.Unlock()

	if _, ok := M.Clients[c]; ok {
		c.conn.Close()
		delete(M.Clients, c)
	}
}
