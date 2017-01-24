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
		"Geek":             "Spider-man",
		"Cool Billionaire": "Batman",
		"Alien":            "Superman", "Amzonian": "Wonder Woman",
		"Less Cool Billionaire": "Iron Man",
	}

	err := tpl.Execute(os.Stdout, superheroes)
	if err != nil {
		log.Fatalln(err)
	}
}
