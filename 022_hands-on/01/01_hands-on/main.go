package main

import (
	"io"
	"net/http"
)

func inicio(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "inicio")
}

func dog(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Dog")
}

func me(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Matheus")
}

func main() {
	http.HandleFunc("/", inicio)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)
	http.ListenAndServe(":8080", nil)
}
