package main

import (
	"image"
	"image/png"
	"os"

	"github.com/damiisdandy/placeholder/pkg/colorparser"
)

// TODO: validators

func main() {
	width := 500
	height := 500
	opacity := 1

	// Create a new image with the specified width and height
	img := image.NewRGBA(image.Rectangle{
		Min: image.Point{0, 0},          // bottom left corner of quad
		Max: image.Point{width, height}, // top right corner of quad
	})
	pixelColor, err := colorparser.HexToRGB("#fff", float32(opacity))
	if err != nil {
		panic(err)
	}

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			// Set the color of each pixel to red
			img.Set(x, y, pixelColor)
		}
	}

	file, err := os.Create("output.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	png.Encode(file, img)
}
