package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*.tpl"))
}

func main() {

	err := tpl.Execute(os.Stdout, "Peace")

	if err != nil {
		log.Fatalln(err)
	}

}
