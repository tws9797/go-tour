package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}

	fmt.Println(t.S)
}

func main() {

	var i I

	var t *T

	//If the concrete value inside the interface itself is nil , the method will be called with a nil receiver
	i = t
	describe(i)
	i.M()

	//An interface value that holds a nil concrete value is itself non-nil
	if i == nil {
		fmt.Println("NULL")
	} else {
		fmt.Println("NOT NULL")
	}

	i = &T{"hello"}
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
