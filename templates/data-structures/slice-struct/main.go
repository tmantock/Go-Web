package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type hero struct {
	Name     string
	Identity string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	batman := hero{
		Name:     "Batman",
		Identity: "Bruce Wayne",
	}

	spiderman := hero{
		Name:     "Spider-man",
		Identity: "Peter Parker",
	}

	ironMan := hero{
		Name:     "Iron Man",
		Identity: "Tony Stark",
	}

	wonderWoman := hero{
		Name:     "Wonder Woman",
		Identity: "Diana Prince",
	}

	hulk := hero{
		Name:     "Hulk",
		Identity: "Bruce Banner",
	}

	superheroes := []hero{batman, spiderman, ironMan, wonderWoman, hulk}

	err := tpl.Execute(os.Stdout, superheroes)
	if err != nil {
		log.Fatalln(err)
	}
}
