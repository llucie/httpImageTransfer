package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.TextFormatter{})

	// Initialize router
	router := mux.NewRouter()

	args := os.Args

	if args[1] == "client" {
		log.Print("Run Client, listening on port 8080")
		router.Methods("GET").Path("/image").HandlerFunc(ClientGetImage)
		log.Fatal(http.ListenAndServe(":8080", router))
	} else if args[1] == "server" {
		log.Print("Run Server, listening on port 4242")
		router.Methods("GET").Path("/image").HandlerFunc(ServerGetImage)
		log.Fatal(http.ListenAndServe(":4242", router))
	} else {
		log.Fatal("Wrong command. Choose [client] or [server] option")
	}
}
