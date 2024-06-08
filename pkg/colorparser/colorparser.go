package colorparser

import (
	"image/color"
	"strconv"
	"strings"
)

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
