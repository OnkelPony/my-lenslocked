package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"html/template"
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tpl, err := template.ParseFiles("templates/home.gohtml")
	if err != nil {
		log.Printf("parsing template: %v", err)
		http.Error(w, "There was an error parsing template.", http.StatusInternalServerError)
		return
	}
	err = tpl.Execute(w, "prase")
	if err != nil {
		log.Printf("executing template: %v", err)
		http.Error(w, "There was an error executing template.", http.StatusInternalServerError)
		return
	}
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact page</h1><p><a href=\"mailto:zwo@centrum.cz\">Napiš mi</a>")
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `
<h1>FAQ</h1><br>
<ul>
	<li>Q: Na jaké palivo lítá větroň?</li>
	<li>A: Vózduch</li>
	<br>
	<li>Q: Kolik válců má čtvřválcový motor?</li>
	<li>A: Bílý</li>
</ul>
`)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}

func galleryHandler(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	fmt.Fprintf(w, "<h1>Welcome, %v ;-)</h1>", userID)
}

//TODO: Add middleware Logger to single route.
func main() {
	r := chi.NewRouter()
	r.Use(middleware.Heartbeat("/ping"))
	r.Get("/", homeHandler)
	r.With(middleware.Logger).Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.Get("/gallery/{userID}", galleryHandler)
	r.With(middleware.Logger).NotFound(notFoundHandler)
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
