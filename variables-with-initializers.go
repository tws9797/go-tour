package main

import "fmt"

//A var declaration can include initializers, one per variable
var i, j int = 1, 2

func main() {
	//If an initializer is present, the type can be omitted;
	//the variable will take the type of the initializer
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}
