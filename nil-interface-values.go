package main

import "fmt"

type I interface {
	M()
}

func main() {

	//A nil interface value holds neither value nor concrete type
	var i I
	describe(i)

	//Calling a method on a nil interface is a run-time error
	//There is no type inside the interface tuple to indicate which concrete method to call
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
