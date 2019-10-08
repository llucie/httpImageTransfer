package main

import (
	"io/ioutil"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

// ClientGetImage requests an image to the server
func ClientGetImage(w http.ResponseWriter, r *http.Request) {
	// Make HTTP Get request to retrieve the image from the server
	log.Print("Requesting image to server")

	start := time.Now()
	resp, err := http.Get("http://localhost:4242/image")
	elapsed := time.Since(start)

	if err != nil {
		log.Fatal("Failed to retrieve image, err = ", err)
	}
	defer resp.Body.Close()

	ioutil.ReadAll(resp.Body)
	log.Info("Get image took ", elapsed)
}
