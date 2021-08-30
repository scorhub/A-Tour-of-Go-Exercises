package main

// A Tour of Go
// Exercise: Loops and Functions
// https://tour.golang.org/flowcontrol/8

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	var result float64
	z := x / 2
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		if z == result {
			return result
		}
		result = z
	}
	return result
}

func main() {
	fmt.Println(Sqrt(2))
}
