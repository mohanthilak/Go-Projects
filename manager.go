package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type WelcomeStruct struct {
	TimeCreated int64 `json:"timeCreated"`
}

type JumboFrameStruct struct {
	JumboMessage string `json:"jumboMessage"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  512,
	WriteBufferSize: 1024,
}

type Manager struct {
	Clients ClientList
	sync.RWMutex
	handlers map[string]EventHandler
}

func MakeManager() *Manager {
	m := &Manager{
		Clients:  make(ClientList),
		handlers: make(map[string]EventHandler),
	}
	m.SetUpHandlers()
	return m
}

func (M *Manager) SetUpHandlers() {
	M.handlers["send-hello-broadcast"] = SendHelloBroadcast
	M.handlers["Successfully Connected"] = SuccessfulConnection
	M.handlers["Jumbo-Message-Rejection"] = JumboFrameMessageHandler
}

func JumboFrameMessageHandler(event Event, client *Client) error {
	var jumboMessage JumboFrameStruct
	if err := json.Unmarshal(event.Payload, &jumboMessage); err != nil {
		log.Println("error while unmarshalling json", err)
	}
	log.Println(jumboMessage.JumboMessage)
	return nil
}

func SuccessfulConnection(event Event, client *Client) error {
	var timeC WelcomeStruct
	if err := json.Unmarshal(event.Payload, &timeC); err != nil {
		log.Println("error while ", err)
	}
	e := Event{
		Type:    "New User Connected",
		Payload: event.Payload,
	}
	if err := SendHelloBroadcast(e, client); err != nil {
		return err
	}
	return nil
}

func SendHelloBroadcast(event Event, client *Client) error {
	for c := range client.manager.Clients {
		c.engres <- event
	}
	return nil
}

func (M *Manager) RouteEvents(event Event, client *Client) error {
	if handler, ok := M.handlers[event.Type]; ok {
		if err := handler(event, client); err != nil {
			return err
		}
		return nil
	} else {
		return errors.New("no such route handler")
	}
}

func (M *Manager) serveWS(w http.ResponseWriter, r *http.Request) {
	log.Println("New Connection")

	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

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
