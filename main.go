package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	adminHandler := demoMiddleware(http.StripPrefix("/admin", adminMux()))
	mux.Handle("/admin/", adminHandler)
	mux.Handle("/", demoHandler("root page"))
	mux.Handle("/dashboard", demoHandler("Dashboard"))
	http.ListenAndServe(":8080", mux)
}

func demoMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Before handling the request for ", r.URL.Path)
		next.ServeHTTP(w, r)
		log.Println("After handling the request")
	})
}

func adminMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", demoHandler("root admin page"))
	mux.HandleFunc("/users", demoHandler("A list of users..."))
	mux.HandleFunc("/users/{id}/edit", editUserHandler())
	return mux
}

// Used to differentiate routes
func demoHandler(msg string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, msg)
	}
}

func editUserHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the ID from the URL
		id := r.PathValue("id")
		fmt.Fprintf(w, "Editing user with ID: %s", id)
	}
}
