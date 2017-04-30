package main

import (
	"net/http"

	"encoding/json"

	"github.com/gorilla/mux"
)

type name struct {
	FirstName string `json:"firstname"`
	Surname   string `json:"lastname"`
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/names", getNames)
	http.ListenAndServe(":8080", r)
}

func getNames(w http.ResponseWriter, r *http.Request) {
	james := name{"James", "Bourne"}
	jenny := name{"Jenny", "Ferry"}

	n := []name{james, jenny}

	json.NewEncoder(w).Encode(n)
}
