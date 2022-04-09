package main

import "fmt"

func main() {
	//When declaring a variable without specifying an explicit type (either by using the := syntax or var = expression syntax)
	//the variable's type is inferred from the value on the right-hand side

	v := 42

	//v = "a" will show error (try to change int type to string type)
	//v = 'a' is valid (show character in int value)

	//But when the right-hand side contains an untyped numeric constant
	//The new variable may be an int, float64, or complex128
	//depending on the precision of the constant:
	/*
		i := 42           // int
		f := 3.142        // float64
		g := 0.867 + 0.5i // complex128
	*/
	fmt.Printf("v is of type %T\n", v)
}
