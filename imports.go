package main

//This code groups the imports into a parenthesized, "factored" import statement
import (
	"fmt"
	"math"
)

//Alternative way (multiple import statements)
/*
import "fmt"
import "math"

Factored import statement is a better style
*/

func main() {
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
}
