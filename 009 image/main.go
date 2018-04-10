// https://go-tour-ru-ru.appspot.com/methods/25

package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

// MyImage - my image type
type MyImage struct {
	height, width int
}

func main() {
	img := MyImage{height: 200, width: 250}
	pic.ShowImage(img)
}

// ColorModel returns the Image's color model.
func (img MyImage) ColorModel() color.Model {
	return color.RGBAModel
}

// Bounds returns the domain for which At can return non-zero color.
// The bounds do not necessarily contain the point (0, 0).
func (img MyImage) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.width, img.height)
}

// At returns the color of the pixel at (x, y).
// At(Bounds().Min.X, Bounds().Min.Y) returns the upper-left pixel of the grid.
// At(Bounds().Max.X-1, Bounds().Max.Y-1) returns the lower-right one.
func (img MyImage) At(x, y int) color.Color {
	return color.RGBA{
		uint8(x*x - x ^ y),
		uint8(x * y),
		uint8(y*y + x ^ y),
		uint8(x * y)}
}
