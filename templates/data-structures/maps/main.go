package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	superheroes := map[string]string{
		"Peter Parker": "Spider-man",
		"Bruce Wayne":  "Batman",
		"Clark Kent":   "Superman",
		"Diana Prince": "Wonder Woman",
		"Tony Stark":   "Iron Man",
	}

	err := tpl.Execute(os.Stdout, superheroes)
	if err != nil {
		log.Fatalln(err)
	}
}
