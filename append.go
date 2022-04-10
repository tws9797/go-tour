package main

import "fmt"

func main() {
	var s []int
	printSlice(s)

	//The first parameter s of append is a slice of type T
	//The rest are T values to append to the slice
	// append works on nil slices
	s = append(s, 0)
	fmt.Printf("%p\n", &s)
	printSlice(s)

	// the slice grows as needed
	s = append(s, 1)
	fmt.Printf("%p\n", &s)
	printSlice(s)

	// add more than one element at a time
	s = append(s, 2, 3, 4, 5, 6)
	fmt.Printf("%p\n", &s)
	printSlice(s)

	fmt.Println(cap(s), len(s))
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
