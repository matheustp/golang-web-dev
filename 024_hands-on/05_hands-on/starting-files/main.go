package main

import (
	"html/template"
	"net/http"
)

func tpl(w http.ResponseWriter, req *http.Request) {
	temp, err := template.ParseGlob("templates/*")
	if err != nil {
		panic(err)
	}
	temp.Execute(w, "index.gohtml")
}

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.HandleFunc("/", tpl)
	http.Handle("/pics/", fs)
	http.ListenAndServe(":8080", nil)
}
