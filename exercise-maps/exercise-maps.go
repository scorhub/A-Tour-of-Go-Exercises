package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

// A Tour of Go
// Exercise: Maps
// https://tour.golang.org/moretypes/23

func WordCount(s string) map[string]int {

	result := make(map[string]int)

	var sentence = strings.Fields(s)

	for _, word := range sentence {
		result[word] += 1
	}

	return result

}

func main() {
	wc.Test(WordCount)
}
