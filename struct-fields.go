package main

import "fmt"

func main() {
	v := Vertex{1, 2}

	//Struct fields are accessed using a dot
	v.X = 4
	fmt.Println(v.X)
}
