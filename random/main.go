package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	const max = 5
	var uniques [max]int

loop:
	for found := 0; found < max; {
		n := rand.Intn(max) + 1
		fmt.Print(n, " ")

		for _, u := range uniques {
			if n == u {
				continue loop
			}
		}

		uniques[found] = n
		found++
	}

	fmt.Println("\n\nUniques ", uniques)
}
