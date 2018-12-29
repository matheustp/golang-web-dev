package main

import (
	"html/template"
	"io"
	"net/http"
)

func foo(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Foo ran")
}

func dog(w http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("dog.gohtml")
	if err != nil {
		panic(err)
	}
	tpl.Execute(w, nil)
}

func dogPic(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "dog.jpg")
}

func me(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Matheus")
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", foo)
	mux.HandleFunc("/me/", me)
	mux.HandleFunc("/dog", dog)
	mux.HandleFunc("/dog.jpg", dogPic)
	http.ListenAndServe(":8085", mux)
}
