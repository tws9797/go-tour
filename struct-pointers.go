package main

import "fmt"

type Vertex struct {
	X, Y int
}

func main() {
	v := Vertex{1, 2}

	//Struct fields can be accessed through a struct pointer
	p := &v
	p.X = 1e9 //equivalent to (*p).X
	fmt.Println(v)
}
