package main

import "fmt"

func main() {

	//A slice has both a length and a capacity
	//The length of a slice is the number of elements it contains
	//The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice

	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	//Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	//Extend its length.
	//panic: runtime error: slice bounds out of range [:8] with capacity 6
	//s = s[:8]
	//printSlice(s)

	//Remember slice is a reference to underlying array
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
