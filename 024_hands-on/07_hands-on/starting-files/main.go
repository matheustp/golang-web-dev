package main

import (
	"net/http"
	"text/template"
)

func tpl(w http.ResponseWriter, req *http.Request) {
	temp, err := template.ParseGlob("templates/*")
	if err != nil {
		panic(err)
	}
	temp.Execute(w, "index.gohtml")
}
func main() {
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("public"))))
	http.HandleFunc("/", tpl)
	http.ListenAndServe(":8080", nil)
}
