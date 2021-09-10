package main

import "fmt"

type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name  string
	Breed string
}

func main() {
	dog := Dog{
		Name:  "Lyka",
		Breed: "German Shephered",
	}
	PrintInfo(&dog)
}

func PrintInfo(a Animal) {
	fmt.Println("The animal says ", a.Says(), " and has ", a.NumberOfLegs(), " legs")
}

func (a *Dog) Says() string {
	return "Woof"
}

func (a *Dog) NumberOfLegs() int {
	return 4
}
