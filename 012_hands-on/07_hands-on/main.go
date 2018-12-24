package main

import (
	"os"
	"text/template"
)

var tpl *template.Template

type menu struct {
	Dinner, Breakfast, Lunch string
}

type restaurant struct {
	Name string
	Menu menu
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	m := []restaurant{
		restaurant{
			Name: "Rest 1",
			Menu: menu{
				Dinner:    "Miojo",
				Breakfast: "Pão",
				Lunch:     "Arroz",
			},
		},
		restaurant{
			Name: "Rest 2",
			Menu: menu{
				Dinner:    "Lasanha",
				Breakfast: "Pão de Queijo",
				Lunch:     "Feijoada",
			},
		},
	}

	err := tpl.Execute(os.Stdout, m)
	if err != nil {
		panic(err)
	}
}
