package main

import (
	"os"
	"text/template"
)

type hotel struct {
	Name, Address, City, Zip, Region string
}

type region struct {
	Region string
	Hotels []hotel
}

type Regions []region

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	regions := Regions{
		region{
			Region: "Southern",
			Hotels: []hotel{
				hotel{
					Name:    "Teste1",
					Address: "Adr",
					City:    "Santos",
					Zip:     "11000000",
					Region:  "southern",
				},
				hotel{
					Name:    "Teste2",
					Address: "Adr2",
					City:    "Santo2s",
					Zip:     "110000020",
					Region:  "southern",
				},
			},
		},
		region{
			Region: "Central",
			Hotels: []hotel{
				hotel{
					Name:    "Teste4",
					Address: "Adr",
					City:    "Santos2",
					Zip:     "11000000",
					Region:  "central",
				},
				hotel{
					Name:    "Teste3",
					Address: "Adr2",
					City:    "Santo31s",
					Zip:     "110000020",
					Region:  "central",
				},
			},
		},
	}
	err := tpl.Execute(os.Stdout, regions)
	if err != nil {
		panic(err)
	}
}
