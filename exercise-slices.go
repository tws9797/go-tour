package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {

	p := make([][]uint8, dy)

	for y := range p {
		p[y] = make([]uint8, dx)
		for x := 0; x < dx; x++ {
			//replace the formula here
			p[y][x] = uint8((x ^ y) * (x ^ y))
		}
	}

	return p
}

func main() {
	pic.Show(Pic)
}
