package main

import "fmt"

func main() {

	//The interface type that specifies zero methods is known as the empty interface
	var i interface{}
	describe(i)

	//An empty interfaces may hold values of any type. (Every type implements at least zero method)
	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

/*
Empty interfaces are used by code that handles values of unknown type.
For example, fmt.Print takes any number of arguments of type interface{}
*/
