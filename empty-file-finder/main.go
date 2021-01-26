package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Provide a directory")
		return
	}

	files, err := ioutil.ReadDir(args[0])

	if err != nil {
		fmt.Println(err)
		return
	}

	var names []byte

	for _, file := range files {
		if file.Size() == 0 {
			// fmt.Println(file.Name())
			names = append(names, file.Name()...)
			names = append(names, '\n')
		}
	}

	err = ioutil.WriteFile("filenames.txt", names, 0644)

	if err != nil {
		fmt.Println(err)
		return
	}
}
