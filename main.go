package main

import (
	"context"
	"example/todo-list/Controllers"
	"example/todo-list/Infra/DB"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	var controller Controllers.ControllerStruct
	controller.DB = DB.NewLinkWorker(client)

	r := mux.NewRouter()

	log.Println("Hi there")

	r.HandleFunc("/create", controller.CreateHandler).Methods("POST")

	//Graceful Shutdown
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
