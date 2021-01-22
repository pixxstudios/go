package main

import (
	s "github.com/inancgumus/prettyslice"
)

func main() {
	var todo []string

	todo = append(todo, "eat")
	todo = append(todo, "sleep", "code", "repeat")

	s.Show("todo", todo)
}
