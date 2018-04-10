// https://go-tour-ru-ru.appspot.com/moretypes/18

package main

import "golang.org/x/tour/pic"

func main() {
	pic.Show(Pic)
}

func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dx)

	for x := 0; x < dx; x++ {
		pic[x] = make([]uint8, dy)
		for y := 0; y < dy; y++ {
			pic[x][y] = uint8(x*y - x ^ y)
		}
	}

	return pic
}
