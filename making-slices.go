package main

import "fmt"

func main() {
	//Slices can be created with the built-in make function
	//The make function allocates a zeroed array and returns a slice that refers to that array
	a := make([]int, 5)
	printSlice("a", a)

	//To specify a capacity, pass a third argument to make
	b := make([]int, 0, 5)
	printSlice("b", b)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
