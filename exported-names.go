package main

import (
	"fmt"
	"math"
)

func main() {
	//In Go, a name is exported if it begins with a capital letter
	fmt.Println(math.Pi)

	//When importing a package, you can refer only to its exported names
	//Any "unexported" names are not accessible from outside the package
}
