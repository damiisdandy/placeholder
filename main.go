package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"strconv"
	"strings"
)

// TODO: validators

// Converts a short hexadecimal color string to a long hexadecimal color string
func HexShortToLong(hex string) string {
	hex = hex[0:1] + hex[0:1] + hex[1:2] + hex[1:2] + hex[2:3] + hex[2:3]
	return hex
}

// Converts a hexadecimal color string to a RGBA color (A being from 0 to 1)
func HexToRGB(hex string, alpha float32) (color.RGBA, error) {
	hex = strings.TrimPrefix(hex, "#")
	if len(hex) == 3 {
		hex = HexShortToLong(hex)
	}

	r, err := strconv.ParseUint(hex[0:2], 16, 8)
	if err != nil {
		return color.RGBA{}, err
	}
	g, err := strconv.ParseUint(hex[2:4], 16, 8)
	if err != nil {
		return color.RGBA{}, err
	}
	b, err := strconv.ParseUint(hex[4:6], 16, 8)
	if err != nil {
		return color.RGBA{}, err
	}
	alpha = alpha * 255
	return color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: uint8(alpha)}, nil
}

func main() {
	width := 100
	height := 100
	opacity := 1

	// Create a new image with the specified width and height
	img := image.NewRGBA(image.Rectangle{
		Min: image.Point{0, 0},          // bottom left corner of quad
		Max: image.Point{width, height}, // top right corner of quad
	})
	pixelColor, err := HexToRGB("#1234d1", float32(opacity))
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
