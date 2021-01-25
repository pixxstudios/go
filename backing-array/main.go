package main

import "fmt"

func main() {
	a := [4]int{1, 2, 3, 4}

	fmt.Println(a)

	s := a[:]
	fmt.Println(s)

	a[0] = 10
	fmt.Println(a)
	fmt.Println(s)
}
