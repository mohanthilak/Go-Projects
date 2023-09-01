package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// var upgrader = websocket.Upgrader{
// 	ReadBufferSize:  1024,
// 	WriteBufferSize: 1024,
// }

// func wsReader(conn *websocket.Conn) {
// 	for {
// 		messageType, p, err := conn.ReadMessage()
// 		if err != nil {
// 			log.Println(err)
// 			return
// 		}
// 		log.Println(messageType, string(p))
// 		wsWriter(conn, p)
// 	}
// }

// func wsWriter(conn *websocket.Conn, p []byte) {
// 	err := conn.WriteMessage(1, p)
// 	if err != nil {
// 		log.Println("damn that error", err)
// 	}
// }

func main() {
	router := mux.NewRouter()
	// wsr := router.PathPrefix("/ws").Subrouter()

	manager := MakeManager()

	router.HandleFunc("/ws", manager.serveWS)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{"success": true, "data": "Message Received", "error": nil})
	})

	server := http.Server{
		Addr:         "127.0.0.1:8000",
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  1 * time.Minute,
	}

	log.Println("listening at port 8000")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
