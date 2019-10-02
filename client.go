package main

import (
	"image"
	"image/draw"
	"image/jpeg"
	"log"
	"net/http"
	"time"
)

func sendGetRequest() {
	// Make HTTP Get request to retrieve the image from the server
	log.Print("Requesting image to server")

	start := time.Now()
	resp, err := http.Get("http://localhost:4242/image")
	elapsed := time.Since(start)
	log.Printf("Get image took %s", elapsed)

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	m, err := jpeg.Decode(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	b := m.Bounds()
	newImage := image.NewGray16(image.Rect(0, 0, b.Dx(), b.Dy()))
	draw.Draw(newImage, newImage.Bounds(), m, b.Min, draw.Src)

	log.Print("Pixels: ", newImage)
}

// ClientGetImage requests an image to the server
func ClientGetImage(w http.ResponseWriter, r *http.Request) {
	// Make HTTP Get request to retrieve the image from the server
	log.Print("Requesting image to server")
	resp, err := http.Get("http://localhost:4242/image")

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	m, _, err := image.Decode(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	b := m.Bounds()
	newImage := image.NewGray16(image.Rect(0, 0, b.Dx(), b.Dy()))
	draw.Draw(newImage, newImage.Bounds(), m, b.Min, draw.Src)

	log.Print("Pixels: ", newImage)
}
