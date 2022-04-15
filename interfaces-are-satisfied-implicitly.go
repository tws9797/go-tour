package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	//A type implements an interface by implementing its methods
	//No explicit declaration of intent, no "implements" keyword
	var i I = T{"hello"}
	i.M()
}
