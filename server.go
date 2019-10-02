package main

import (
	"bytes"
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"net/http"
	"strconv"
	"time"
)

// ServerGetImage TODO
func ServerGetImage(w http.ResponseWriter, r *http.Request) {
	//img := image.NewGray16(image.Rect(0, 0, 2394, 2850))
	img := image.NewGray16(image.Rect(0, 0, 2, 1))
	img.SetGray16(0, 0, color.Gray16{Y: 42})

	// Print created
	log.Print("Pixels: ", img)

	buffer := new(bytes.Buffer)
	if err := jpeg.Encode(buffer, img, nil); err != nil {
		log.Println("unable to encode image.")
	}

	w.Header().Set("Content-Type", "image/jpeg")
	w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))

	start := time.Now()
	if _, err := w.Write(buffer.Bytes()); err != nil {
		log.Println("unable to write image.")
	}
	elapsed := time.Since(start)
	log.Printf("Write image took %s", elapsed)
}
