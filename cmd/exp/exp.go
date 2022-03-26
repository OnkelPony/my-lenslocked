package main

import (
	"html/template"
	"os"
)

type User struct {
	Name   string
	Bio    string
	Age    int
	Father string
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	user := User{
		Name:   "Kelly Divine",
		Bio:    `I'm really good Go dev, pyčo!`,
		Age:    108,
		Father: "Ken Divine",
	}

	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}
