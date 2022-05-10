package controllers

import (
	"github.com/OnkelPony/my-lenslocked/views"
	"net/http"
)

type Users struct {
	Templates struct {
		New views.Template
	}
}

func (u Users) New(w http.ResponseWriter, r *http.Request) {
	// TODO: Render the signup page.
}
