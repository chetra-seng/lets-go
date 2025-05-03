package main

import (
	"net/http"

	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./ui/static"))

	// Dynamic middle for session
	dynamic := alice.New(app.sessionManager.LoadAndSave)

	// {$} special character to prevent subtree path pattern(anything that end with trailing slash) aka catch all
	mux.Handle("GET /{$}", dynamic.ThenFunc(app.home))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))
	mux.Handle("GET /snippet/view/{id}", dynamic.ThenFunc(app.viewSnippet))
	mux.Handle("GET /snippet/create", dynamic.ThenFunc(app.createSnippetForm))
	mux.Handle("POST /snippet/create", dynamic.ThenFunc(app.createSnippet))
	mux.Handle("POST /user/signup", dynamic.ThenFunc(app.signupUser))
	mux.Handle("GET /user/signup", dynamic.ThenFunc(app.signupForm))
	mux.Handle("POST /user/login", dynamic.ThenFunc(app.loginUser))
	mux.Handle("GET /user/login", dynamic.ThenFunc(app.loginUserForm))
	mux.Handle("POST /user/logout", dynamic.ThenFunc(app.logoutUser))

	standard := alice.New(app.recoverPanic, app.logRequest, commonHeaders)

	return standard.Then(mux)
}
