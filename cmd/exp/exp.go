package main

import (
	"html/template"
	"os"
)

type User struct {
	Name      string
	Bio       string
	Age       int
	Lovers    []string
	Positions map[string]string
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	user := User{
		Name:   "Kelly Divine",
		Bio:    "I'm porn actress, really good in anal.",
		Age:    37,
		Lovers: []string{"Giorgio", "Schorschi", "OnkelPony", "Pachatel"},
		Positions: map[string]string{
			"deepthroat": "like",
			"missionary": "dislike",
			"anal":       "love",
		},
	}

	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}
