package main

import "net/http"

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./ui/static"))

	// {$} special character to prevent subtree path pattern(anything that end with trailing slash) aka catch all
	mux.HandleFunc("GET /{$}", app.home)
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("GET /snippet/view/{id}", app.viewSnippet)
	mux.HandleFunc("GET /snippet/create", app.createSnippetForm)
	mux.HandleFunc("POST /snippet/create", app.createSnippet)

	return commonHeaders(mux)
}
