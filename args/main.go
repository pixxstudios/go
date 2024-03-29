package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	num1, err := strconv.Atoi(os.Args[1])

	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	num2, err := strconv.Atoi(os.Args[2])

	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	sum := num1 + num2

	sum1 := Sum(1, 2, 3, 4)
	fmt.Println(sum1)

	fmt.Printf("%d + %d = %d\n", num1, num2, sum)
}
