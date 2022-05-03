package controllers

import (
	"github.com/OnkelPony/my-lenslocked/views"
	"html/template"
	"net/http"
)

func StaticHandler(tpl views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}

func FAQ(tpl views.Template) http.HandlerFunc {
	qAndA := []struct {
		Question string
		Answer   template.HTML
		QType    string
	}{
		{
			Question: "Jak se máš?",
			Answer:   "De to!",
			QType:    "blbá",
		},
		{
			Question: "Za kolik to máš?",
			Answer:   "Za pade.",
			QType:    "přímá",
		},
		{
			Question: "Is ammunition included?",
			Answer:   "Of course!",
			QType:    "chytrá",
		},
		{
			Question: "How do I contact support?",
			Answer:   `Email us - <a href="mailto:support@lenslocked.com">support@lenslocked.com</a>`,
			QType:    "praktická",
		},
	}
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, qAndA)
	}
}
