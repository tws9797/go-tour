package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(b []byte) (int, error) {
	for i, c := range b {
		switch {
		case c >= 'A' && c <= 'M' || c >= 'a' && c <= 'm':
			b[i] += 13
		case c >= 'N' && c <= 'Z' || c >= 'n' && c <= 'z':
			b[i] -= 13
		}
	}

	return r.r.Read(b)
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
