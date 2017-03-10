package controllers

import (
	"os"
	"path/filepath"

	"io/ioutil"
	"strings"

	"time"

	"github.com/russross/blackfriday"
	"github.com/tmantock/Go-Web/mrkdwn/models"
)

type PostController struct{}
type Posts []models.Post

func NewPostController() *PostController {
	return &PostController{}
}

func (pc PostController) GetPosts() Posts {
	var p Posts
	cwd, _ := os.Getwd()
	files, _ := filepath.Glob(filepath.Join(cwd, "posts/*"))

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
		p = append(p, models.Post{title, date, summary, body, file})
	}

	return p
}
