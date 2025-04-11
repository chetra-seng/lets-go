package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static"))

	// {$} special character to prevent subtree path pattern(anything that end with trailing slash) aka catch all
	mux.HandleFunc("GET /{$}", home)
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("GET /snippet/view/{id}", viewSnippet)
	mux.HandleFunc("GET /snippet/create", createSnippetForm)
	mux.HandleFunc("POST /snippet/create", createSnippet)

	log.Print("Started on port 4000")

	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		log.Fatal(err)
	}
}
