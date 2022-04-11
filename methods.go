package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

//Go does not have classes
//However, you can define methods on types
//The receiver appears in its own argument list between the func keyword and the method name
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}
