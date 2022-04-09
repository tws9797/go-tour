package main

import "fmt"

//Vertex is defined in structs.go

func main() {
	v := Vertex{1, 2}

	//Struct fields are accessed using a dot
	v.X = 4
	fmt.Println(v.X)
}
