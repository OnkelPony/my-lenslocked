package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to my awesome site, passive income!</h1>")
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
</ul>`)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}

type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contacts":
		contactHandler(w, r)
	case "/faq":
		faqHandler(w, r)
	default:
		notFoundHandler(w, r)
	}
}

// func pathHandler(w http.ResponseWriter, r *http.Request) {
// 	switch r.URL.Path {
// 	case "/":
// 		homeHandler(w, r)
// 	case "contacts":
// 		contactHandler(w, r)
// 	default:
// 		notFoundHandler(w, r)
// 	}
// }

func main() {
	var router Router
	// http.HandleFunc("/", pathHandler)
	//	http.HandleFunc("/", homeHandler)
	// http.HandleFunc("/contacts", contactHandler)
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", router)
}
