package controllers

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

var FM = template.FuncMap{
	"fDate": unixTime,
}

var tpl *template.Template

type HomeController struct{}

func NewHomeController() *HomeController {
	return &HomeController{}
}

func unixTime(t time.Time) string {
	return t.Format(time.UnixDate)
}

func (hc HomeController) HomeRoute(w http.ResponseWriter, r *http.Request) {
	cwd, _ := os.Getwd()
	posts := GetPosts()
	tpl = template.Must(template.New("").Funcs(FM).ParseFiles(filepath.Join(cwd, "./templates/index.gohtml")))
	// if err != nil {
	// 	panic(err)
	// }

	tpl.ExecuteTemplate(w, "index.gohtml", posts)
}
