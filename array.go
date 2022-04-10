package main

import "fmt"

func main() {

	//The type [n]T is an array of n values of type T
	//declares a variable a as an array of two strings
	var a [2]string

	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])

	//An array's length is part of its type, so arrays cannot be resized
	primes := [6]int{2, 3, 4, 7, 11, 13}
	fmt.Println(primes)
	fmt.Println(&primes)
	fmt.Printf("%p\n", &primes)
}
