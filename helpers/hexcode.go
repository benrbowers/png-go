package helpers

import (
	"encoding/hex"
	"image/color"
)

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
