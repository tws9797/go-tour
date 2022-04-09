package main

import "fmt"

//var statement declares a list of variables; as in function argument lists, the type is last
var c, python, java bool

//var statement can be at package or function level
func main() {
	var i int
	fmt.Println(i, c, python, java)
}
