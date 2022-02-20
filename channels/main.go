package main

import "fmt"

func main() {
	// normal (unbuffered) channel
	message := make(chan string)

	go func() { message <- "ping" }()

	fmt.Println(<-message)

	// buffered channel

	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)

}
