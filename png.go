package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"png/helpers"
)

func main() {
	var colors map[string][]color.RGBA = map[string][]color.RGBA{}
	var currentColor string

	colorsFile, err := os.Open("TailwindColors.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(colorsFile)
	for scanner.Scan() {
		text := scanner.Text()

		if text[0] != '#' {
			colors[text] = []color.RGBA{}
			currentColor = text
		} else {
			colors[currentColor] = append(
				colors[currentColor],
				helpers.HexcodeToRGB(text),
			)
		}
	}

	colorsFile.Close()

	width := 11
	height := 1

	fmt.Println("Width:", width)

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	for color, shades := range colors {
		img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

		for x, shade := range shades {
			img.Set(x, 0, shade)
		}

		// Encode as PNG.
		f, _ := os.Create("palettes/tailwind-" + color + ".png")
		png.Encode(f, img)
	}
}
