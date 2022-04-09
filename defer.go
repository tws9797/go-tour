package main

import "fmt"

func main() {
	//defer statement defers the execution of a function until the surrounding function returns
	defer fmt.Println("!")

	fmt.Println("hello")

	//The deferred call's arguments are evaluated immediately
	//but the function call is not executed until the surrounding function returns
	defer fmt.Println("world")

	fmt.Println("good")
}
