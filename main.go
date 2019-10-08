package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func main() {
	setFormatter()

	log.Info("Starting Http image transfer")

	// Initialize router
	router := mux.NewRouter()

	args := os.Args

	if args[1] == "client" {
		log.Info("Run Client, listening on port 8080")
		router.Methods("GET").Path("/image").HandlerFunc(ClientGetImage)
		log.Fatal(http.ListenAndServe(":8080", router))
	} else if args[1] == "server" {
		log.Info("Run Server, listening on port 4242")
		createImg()
		router.Methods("GET").Path("/image").HandlerFunc(ServerGetImage)
		log.Fatal(http.ListenAndServe(":4242", router))
	} else {
		log.Fatal("Wrong command. Choose [client] or [server] option")
	}
}

func setFormatter() {
	customFormatter := new(log.TextFormatter)
	customFormatter.TimestampFormat = time.StampMilli
	customFormatter.FullTimestamp = true
	log.SetFormatter(customFormatter)

	// Create the log file if doesn't exist. And append to it if it already exists.
	var filename string = "logfile.log"
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)

	if err != nil {
		// Cannot open log file. Logging to stderr
		fmt.Println(err)
	} else {
		log.SetOutput(f)
	}
}
