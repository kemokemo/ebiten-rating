package main

import (
	"bytes"
	_ "embed"
	"image"
	_ "image/png"
	"log"
)

var (
	//go:embed assets/Star.png
	star_png []byte

	StarImage image.Image
)

func init() {
	var err error
	StarImage, err = loadSingleImage(star_png)
	if err != nil {
		log.Println("failed to load star image, ", err)
		return
	}
}

func loadSingleImage(b []byte) (image.Image, error) {
	img, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	return img, nil
}
