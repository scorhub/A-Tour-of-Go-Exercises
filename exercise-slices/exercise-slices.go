package main

import (
	"golang.org/x/tour/pic"
)

// A Tour of Go
// Exercise: Slices
// https://tour.golang.org/moretypes/18

func Pic(dx, dy int) [][]uint8 {
	picture := make([][]uint8, dx)

	for x := range picture {
		picture[x] = make([]uint8, dy)
		for y := range picture[x] {
			picture[x][y] = uint8(x ^ y)
		}

	}
	return picture
}

func main() {
	pic.Show(Pic)
}
