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

type car struct {
	Model        string
	Manufacturer string
	Doors        int
}

type items struct {
	Heroes    []hero
	Transport []car
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

	accord := car{
		Model:        "Accord",
		Manufacturer: "Honda",
		Doors:        4,
	}

	civic := car{
		Model:        "Civic",
		Manufacturer: "Honda",
		Doors:        4,
	}

	modelx := car{
		Model:        "Model X",
		Manufacturer: "Tesla",
		Doors:        5,
	}

	miata := car{
		Model:        "Miata",
		Manufacturer: "Mazda",
		Doors:        2,
	}

	superheroes := []hero{batman, spiderman, ironMan, wonderWoman, hulk}
	cars := []car{accord, civic, modelx, miata}

	data := items{
		Heroes:    superheroes,
		Transport: cars,
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}
