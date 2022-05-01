package main

import "fmt"

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}

	//A sender can close a channel to indicate that no more values will be sent
	//Only the sender should close a channel, never the receiver. Sending on a closed channel will cause a panic
	close(c)
}

func main() {
	c := make(chan int, 10)

	fmt.Println(cap(c))

	//Cap is to check the capacity of a given type
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}
