package models

import (
	"time"
)

type Post struct {
	Title   string
	Date    time.Time
	Summary string
	Body    string
	File    string
}
