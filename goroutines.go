package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

/*
	Go is a concurrent language and not a parallel one
	The main function runs in its own Goroutine and it's called the main Goroutine
*/
func main() {

	//A goroutine is a lightweight thread managed by the Go runtime
	//starts a new goroutine running say("hello")
	go say("world")
	say("hello")
}
