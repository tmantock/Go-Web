package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

type hero struct {
	Name     string
	Identity string
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

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", superheroes)
	if err != nil {
		log.Fatalln(err)
	}
}
