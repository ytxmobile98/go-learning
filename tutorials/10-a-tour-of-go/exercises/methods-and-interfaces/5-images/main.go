package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct{}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 255, 255)
}

func (i Image) At(x, y int) color.Color {
	var r, g, b uint8

	if x == y {
		r, g, b = 255, 255, 255
	} else {
		r, g, b = 0, 0, 0
	}

	return color.RGBA{r, g, b, 255}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
