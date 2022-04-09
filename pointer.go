package main

import "fmt"

func main() {

	i, j := 42, 2701

	//Pointer holds the memory address of a value
	//The type *T is a pointer to a T value. Its zero value is nil
	var v *int
	fmt.Println(v)

	//The & operator generates a pointer to its operand
	p := &i
	fmt.Println(*p)

	//The * operator denotes the pointer's underlying value
	//This is known as dereferencing
	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)

}
