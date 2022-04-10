package main

import "fmt"

type Vertex struct {
	X, Y int
}

//A struct literal denotes a newly allocated struct value by listing the values of its fields
//You can list just a subset of fields by using the Name: syntax. (And the order of named fields is irrelevant.)
var (
	v1 = Vertex{1, 2} // has type Vertex
	v2 = Vertex{X: 1} // Y:0 is implicit
	v3 = Vertex{}     // X:0 and Y:0
	v4 = Vertex{Y: 3, X: 2}

	//The special prefix & returns a pointer to the struct value
	p = &Vertex{1, 2} // has type *Vertex
)

func main() {
	fmt.Println(v1, p, v2, v3, v4)

	//Dereferencing struct pointer
	fmt.Println(*p)

	//Print values using fmt.Println(), the default format will be used
	//For compound objects, the elements are printed using these rules
	/*
		struct:             {field0 field1 ...}
		array, slice:       [elem0 elem1 ...]
		maps:               map[key1:value1 key2:value2 ...]
		pointer to above:   &{}, &[], &map[]
	*/
	//https://stackoverflow.com/questions/56112289/how-to-print-the-address-of-struct-variable-in-go

	//To print address of compound objects
	fmt.Printf("%p\n", p)
}
