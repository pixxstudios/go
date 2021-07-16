package main

import (
	"fmt"
	"os"
	"text/template"
)

func main() {
	tpl := template.Must(template.ParseGlob("*.gohtml"))

	err := tpl.Execute(os.Stdout, nil)

	if err != nil {
		fmt.Println(err)
	}

}
