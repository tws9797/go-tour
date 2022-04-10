package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	//Slices is a dynamically-sized, flexible view into the elements of an array
	//The type []T is a slice with elements of type T

	//A slice is formed by specifying two indices, a low and high bound, separated by a colon:
	//a[low : high]

	//The following expression creates a slice which includes elements 1 through 3 of primes
	var s []int = primes[1:4]
	fmt.Println(s)
	fmt.Println(&s)
	fmt.Printf("%p\n", &s)
}
