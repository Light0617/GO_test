package main

import (
	"log"
	"os"
	"text/template"
)

type user struct {
	Name  string
	Admin bool
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index3.html"))
}

func main() {
	u1 := user{
		Name:  "Toddy",
		Admin: false,
	}

	u2 := user{
		Name:  "Mary",
		Admin: true,
	}

	u3 := user{
		Name:  "",
		Admin: false,
	}

	users := []user{u1, u2, u3}

	err := tpl.Execute(os.Stdout, users)
	if err != nil {
		log.Fatalln(err)
	}
}
