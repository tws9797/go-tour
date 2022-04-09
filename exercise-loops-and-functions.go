package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {

	z := float64(1)

	var t float64

	for {
		//Check how close z² is to x
		z, t = z-(z*z-x)/(2*z), z
		if math.Abs(t-z) < 1e-6 {
			break
		}
	}

	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
