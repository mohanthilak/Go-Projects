package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

type linkStruct struct {
	Link         string
	ShortendLink string
}

var linkList []linkStruct

func createHandler(w http.ResponseWriter, r *http.Request) {
	var link linkStruct
	err := json.NewDecoder(r.Body).Decode(&link)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Println("link:", link.Link)
	for _, el := range linkList {
		if el.Link == link.Link {
			fmt.Fprintf(w, "Link is %s", el.ShortendLink)
			return
		}
	}
	link.ShortendLink = fmt.Sprintf("http://localhost:8000/%d", len(linkList)+1)
	log.Println(link)
	linkList = append(linkList, link)
	fmt.Fprintf(w, "%+v", link)
	return
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/create", createHandler).Methods("POST")

	server := &http.Server{
		Addr:         "127.0.0.1:8000",
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		log.Println("Starting the server on Port: 8000")
		err := server.ListenAndServe()
		if err != nil {
			log.Panicf("Error while starting server: %s\n Gracefully Shutting Down \n", err)
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	sig := <-c

	log.Println("Got Signal", sig)

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(ctx)
}
