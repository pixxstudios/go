package main

import "fmt"

func main() {
	message := make(chan string, 3)

	// go func() {
	// 	message <- "This is coming from a goroutine"
	// 	message <- "This is coming from a goroutine 2"
	// 	message <- "This is coming from a goroutine 3"
	// }()

	message <- "This is coming from a goroutine"
	message <- "This is coming from a goroutine 2"
	message <- "This is coming from a goroutine 3"

	fmt.Println(<-message)
	fmt.Println(<-message)
	fmt.Println(<-message)
}
