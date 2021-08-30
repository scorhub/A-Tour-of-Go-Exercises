package main

// A Tour of Go
// Exercise: Fibonacci closure
// https://tour.golang.org/moretypes/26

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.

func fibonacci() func() int {
	var a = 0
	var b = 1
	return func() int {
		fib := a
		a = b
		b = fib + a
		return fib
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
