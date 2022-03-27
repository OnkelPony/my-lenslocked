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
	HairStyle Hair
}

type Hair struct {
	Color  string
	Length int
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	kelly := User{
		Name:   "Kelly Divine",
		Bio:    "I'm porn actress, really good in anal.",
		Age:    37,
		Lovers: []string{"Giorgio", "Schorschi", "OnkelPony", "Pachatel"},
		Positions: map[string]string{
			"deepthroat": "like",
			"missionary": "dislike",
			"anal":       "love",
		},
		HairStyle: Hair{
			Color:  "red",
			Length: 27,
		},
	}

	err = t.Execute(os.Stdout, kelly)
	if err != nil {
		panic(err)
	}
}
