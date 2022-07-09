package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string, 1)

	// goroutine to mock some external call and return result
	// into channel after 2 seconds
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	select {
	// res will get the data from channel c1 after 2 seconds
	case res := <-c1:
		fmt.Println(res)
	// time.After will send a value after 1 second
	// so this case will get executed
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)

	// goroutine to mock some external call and return result
	// into channel after 2 seconds
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()

	select {
	// res will get the data from channel c2 after 2 seconds
	// so this case will get executed since timeout will happen after 3 seconds
	case res := <-c2:
		fmt.Println(res)
	// time.After will send a value after 3 seconds
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}
