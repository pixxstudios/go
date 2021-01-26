package main

import "fmt"

func main() {
	// a := []int{1, 2, 3, 4}

	// fmt.Println(a)

	// s := a[:]
	// fmt.Println(s)

	// a[0] = 10
	// fmt.Println(a)
	// fmt.Println(s)

	// limit capacity
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(a)

	n1 := a[:2:2] // handles the overwriting of values in backing array
	fmt.Println(n1)

	// this overwrites the values in original backing array
	n1 = append(n1, 6, 7, 8)
	fmt.Println(a)
	fmt.Println(n1)
}
