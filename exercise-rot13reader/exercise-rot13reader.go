package main

// A Tour of Go
// Exercise: rot13Reader
// https://tour.golang.org/methods/23

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot rot13Reader) Read(b []byte) (n int, e error) {
	n, e = rot.r.Read(b)
	for i := range b {
		if (b[i] >= 'a' && b[i] <= 'm') || (b[i] >= 'A' && b[i] <= 'M') {
			b[i] += 13
		} else if (b[i] >= 'n' && b[i] <= 'z') || (b[i] >= 'N' && b[i] <= 'Z') {
			b[i] -= 13
		}
	}
	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
