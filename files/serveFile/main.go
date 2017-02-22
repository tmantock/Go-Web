package main

import (
	"io"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/dog", dog)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="toby.jpg">`)
}

// func dogPic(w http.ResponseWriter, r *http.Request) {
// 	f, err := os.Open("toby.jpg")
// 	if err != nil {
// 		http.Error(w, "File Not Found", 404)
// 		return
// 	}
// defer f.Close()

// fi, err := f.Stat()
// if err != nil {
// 	http.Error(w, "File Not Found", 404)
// 	return
// }

//http.ServeContent(w, r, f.Name(), fi.ModTime(), f)

//io.Copy(w, f)

//http.ServeFile(w, r, "toby.jpg")
//}
