package models

import (
	"html/template"
	"time"
)

type Post struct {
	Title   string
	Date    time.Time
	Summary string
	Body    template.HTML
	File    string
}
