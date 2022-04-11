package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

//A method is just a function with a receiver argument
//Abs is a regular function
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(Abs(v))
}
