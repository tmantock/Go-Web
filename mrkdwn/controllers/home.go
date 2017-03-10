package controllers

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

type HomeController struct{}

func NewHomeController() *HomeController {
	return &HomeController{}
}

func (hc HomeController) HomeRoute(w http.ResponseWriter, r *http.Request) {
	cwd, _ := os.Getwd()
	tpl, err := template.ParseFiles(filepath.Join(cwd, "./templates/index.gohtml"))
	if err != nil {
		panic(err)
	}

	tpl.Execute(w, nil)
}
