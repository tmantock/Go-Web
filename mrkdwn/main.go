package main

import (
	"net/http"

	"github.com/tmantock/Go-Web/mrkdwn/controllers"
)

func main() {
	hc := controllers.NewHomeController()
	http.HandleFunc("/", hc.HomeRoute)
	http.ListenAndServe(":8080", nil)
}
