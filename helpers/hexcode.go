package helpers

import (
	"encoding/hex"
	"image/color"
)

// func ColorHexToRGB(hexstr string) (r, g, b uint8) {
// 	r = hexcodeToRGB(hexstr[1:3])
// 	g = hexcodeToRGB(hexstr[3:5])
// 	b = hexcodeToRGB(hexstr[5:7])

// 	return r, g, b
// }

func HexcodeToRGB(hexcode string) color.RGBA {
	if hexcode[0] == '#' {
		hexcode = hexcode[1:]
	}

	rgb, err := hex.DecodeString(hexcode)

	if err != nil {
		panic(err)
	}

	return color.RGBA{rgb[0], rgb[1], rgb[2], 0xff}
}
