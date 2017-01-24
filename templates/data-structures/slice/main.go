package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("variable.gohtml", "tpl.gohtml"))
}

func main() {
	superheroes := []string{"Spider-man", "Batman", "Superman", "Wonder Woman", "Iron Man"}

	err := tpl.Execute(os.Stdout, superheroes)
	if err != nil {
		log.Fatalln(err)
	}
}
