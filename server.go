package main

import (
	"image"
	"net/http"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
)

var img *image.Gray16

func createImg() {
	img = image.NewGray16(image.Rect(0, 0, 2850, 2394))

	log.Info("Image bytes size: ", len(img.Pix))

	pix := 0
	for x := 0; x < len(img.Pix); x++ {
		img.Pix[x] = uint8(pix)
		pix++
		if pix == 256 {
			pix = 0
		}
	}
}

// ServerGetImage TODO
func ServerGetImage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Length", strconv.Itoa(len(img.Pix)))
	log.Info("Request received, writing image")

	start := time.Now()

	if _, err := w.Write(img.Pix); err != nil {
		log.Error("unable to write image.")
	}
	elapsed := time.Since(start)
	log.Info("Write image took ", elapsed)
}
