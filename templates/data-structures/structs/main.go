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

	err := tpl.Execute(os.Stdout, batman)
	if err != nil {
		log.Fatalln(err)
	}
}
