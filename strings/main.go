package main

import "fmt"

func main() {
	b := []byte{104, 101, 121}
	fmt.Println(b)         // print bytes
	fmt.Println(string(b)) // prints string equivalent of bytes array
}
