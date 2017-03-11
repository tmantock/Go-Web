package controllers

import (
	"html/template"
	"os"
	"path/filepath"

	"io/ioutil"
	"strings"

	"time"

	"net/http"

	"fmt"

	"github.com/russross/blackfriday"
	"github.com/tmantock/Go-Web/mrkdwn/models"
)

type PostController struct{}
type Posts []models.Post

func NewPostController() *PostController {
	return &PostController{}
}

func (pc PostController) PostRoute(w http.ResponseWriter, r *http.Request) {
	spl := strings.Split(r.URL.Path, "/")
	f := spl[2]
	cwd, _ := os.Getwd()
	tpl, err := template.New("").Funcs(FM).ParseFiles(filepath.Join(cwd, "templates/post.gohtml"))
	if err != nil {
		panic(err)
	}

	p := renderPost(f)

	tpl.ExecuteTemplate(w, "post.gohtml", p)
}

func renderPost(fn string) models.Post {
	var p models.Post
	cwd, _ := os.Getwd()

	files, _ := filepath.Glob(filepath.Join(cwd, "posts/"+fn+".md"))

	for _, f := range files {
		file := strings.Replace(f, "posts/", "", -1)
		file = strings.Replace(file, ".md", "", -1)
		fileread, _ := ioutil.ReadFile(f)
		lines := strings.Split(string(fileread), "\n")
		title := string(lines[0])
		sd := string(lines[1])
		date, _ := time.Parse(time.UnixDate, sd)
		summary := string(lines[2])
		body := strings.Join(lines[3:len(lines)], "\n")
		body = string(blackfriday.MarkdownCommon([]byte(body)))
		b := template.HTML(body)
		p = models.Post{title, date, summary, b, file}
	}

	return p
}

func GetPosts() Posts {
	var p Posts
	cwd, _ := os.Getwd()

	files, _ := filepath.Glob(filepath.Join(cwd, "posts/*"))

	for _, f := range files {
		file := strings.Replace(f, cwd+"/posts/", "", -1)
		file = strings.Replace(file, ".md", "", -1)
		fmt.Println(file)
		fileread, _ := ioutil.ReadFile(f)
		lines := strings.Split(string(fileread), "\n")
		title := string(lines[0])
		sd := string(lines[1])
		date, _ := time.Parse(time.UnixDate, sd)
		summary := string(lines[2])
		body := strings.Join(lines[3:len(lines)], "\n")
		body = string(blackfriday.MarkdownCommon([]byte(body)))
		b := template.HTML(body)
		p = append(p, models.Post{title, date, summary, b, file})
	}

	return p
}
