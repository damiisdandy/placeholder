package main

import (
	"image/color"
	"testing"
)

func TestHexShortToLong(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		{"fff", "ffffff"},
		{"123", "112233"},
		{"000", "000000"},
		{"abc", "aabbcc"},
		{"141", "114411"},
	}

	for _, test := range tests {
		output := HexShortToLong(test.input)
		if output != test.output {
			t.Errorf("Expected %s, got %s", test.output, output)
		}
	}
}

func TestHexToRGB(t *testing.T) {
	tests := []struct {
		input  string
		output color.RGBA
	}{
		{"#fff", color.RGBA{255, 255, 255, 255}},
		{"#123", color.RGBA{17, 34, 51, 255}},
		{"#000", color.RGBA{0, 0, 0, 255}},
		{"#141414", color.RGBA{20, 20, 20, 255}},
		{"#d243e3", color.RGBA{210, 67, 227, 255}},
	}

	for _, test := range tests {
		output, err := HexToRGB(test.input, 1.0)
		if err != nil {
			t.Errorf("Error converting hex to RGBA: %s", err)
		}
		if output != test.output {
			t.Errorf("Expected %v, got %v", test.output, output)
		}
	}
}
