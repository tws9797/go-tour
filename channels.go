package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}

	//Send sum to channel c
	c <- sum
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	//Channel must be created before use
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	//Data flows in the direction of the arrow
	//Receive from c, and assign value to x and y
	x, y := <-c, <-c

	/*
		By default, sends and receives block until the other side is ready
		This allows goroutines to synchronize without explicit locks or condition variables
	*/
	fmt.Println(x, y, x+y)
}
