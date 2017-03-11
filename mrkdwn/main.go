package main

import (
	"net/http"

	"github.com/tmantock/Go-Web/mrkdwn/controllers"
)

func main() {
	hc := controllers.NewHomeController()
	pc := controllers.NewPostController()
	http.HandleFunc("/", hc.HomeRoute)
	http.HandleFunc("/posts/", pc.PostRoute)
	http.ListenAndServe(":8080", nil)
}
