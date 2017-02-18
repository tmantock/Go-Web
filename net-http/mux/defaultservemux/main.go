package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/dog", dog)
	http.HandleFunc("/cat", cat)

	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Dog Handler")
}

func cat(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Cat Handler")
}
