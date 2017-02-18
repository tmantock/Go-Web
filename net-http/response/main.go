package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(r http.ResponseWriter, w *http.Request) {
	r.Header().Set("Tevin-Key", "This is from Mantock")
	r.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(r, "<h1>What's up!?</h1>")
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
