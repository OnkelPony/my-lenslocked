package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to my awesome site, future passive income!</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact page</h1><p><a href=\"mailto:zwo@centrum.cz\">Napiš mi</a>")
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>Page not found, jako fakt ne...</h1>")
}
func pathHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, r.URL.Path)
	fmt.Fprintln(w, r.URL.RawPath)
	// switch r.URL.Path {
	// case "/":
	// 	homeHandler(w, r)
	// case "contacts":
	// 	contactHandler(w, r)
	// default:
	// 	notFoundHandler(w, r)
	// }
}

func main() {
	http.HandleFunc("/", pathHandler)
	//	http.HandleFunc("/", homeHandler)
	// http.HandleFunc("/contacts", contactHandler)
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", nil)
}
