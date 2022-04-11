package main

import "fmt"

func adder() func(int) int {

	//Sum is being reference by the function below, therefore the value
	sum := 0

	//A closure is a function value that references variables from outside its body
	//The function may access and assign to the referenced variables
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	//The adder function returns a closure. Each closure is bound to its own sum variable
	pos, neg := adder(), adder()

	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2*i))
	}
}
