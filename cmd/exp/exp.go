package main

import (
	"html/template"
	"os"
)

type User struct {
	Name string
	Bio  template.HTML
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	user := User{
		Name: "Kelly Divine",
		Bio:  `<script> alert("You've been pwned!") </script>`,
	}

	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}
