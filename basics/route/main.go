package main

import (
	"net/http"

	"encoding/json"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type name struct {
	FirstName string `json:"firstname"`
	Surname   string `json:"lastname"`
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/names", getNames)
	handler := cors.Default().Handler(r)
	http.ListenAndServe(":8080", handler)
}

func getNames(w http.ResponseWriter, r *http.Request) {
	james := name{"James", "Bourne"}
	jenny := name{"Jenny", "Ferry"}

	n := []name{james, jenny}

	json.NewEncoder(w).Encode(n)
}
