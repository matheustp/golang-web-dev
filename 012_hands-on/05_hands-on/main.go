package main

import (
	"os"
	"text/template"
)

var tpl *template.Template

type menu struct {
	Dinner, Breakfast, Lunch string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	m := menu{
		Dinner:    "Miojo",
		Breakfast: "Pão",
		Lunch:     "Arroz",
	}

	err := tpl.Execute(os.Stdout, m)
	if err != nil {
		panic(err)
	}
}
