package main

import "fmt"

func main() {
	message := make(chan string)

	go func() {
		message <- "This is coming from a goroutine"
	}()

	fmt.Println(<-message)
}
