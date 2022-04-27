package main

import (
	"fmt"
	"github.com/OnkelPony/my-lenslocked/controllers"
	"github.com/OnkelPony/my-lenslocked/templates"
	"github.com/OnkelPony/my-lenslocked/views"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}

func galleryHandler(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	fmt.Fprintf(w, "<h1>Welcome, %v ;-)</h1>", userID)
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Heartbeat("/ping"))

	r.Get("/", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS, "home.gohtml"))))

	r.With(middleware.Logger).Get("/contact", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS, "contact.gohtml"))))

	r.Get("/faq", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS, "faq.gohtml"))))

	r.With(middleware.Logger).Get("/time", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS, "time.gohtml"))))

	r.Get("/gallery/{userID}", galleryHandler)
	r.With(middleware.Logger).NotFound(notFoundHandler)
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
