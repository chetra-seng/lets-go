package main

import (
	"net/http"

	"chetraseng.com/ui"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	// Dynamic middle for session
	dynamic := alice.New(app.sessionManager.LoadAndSave, noSurf, app.authenticate)

	protected := dynamic.Append(app.requiredAuthentication)

	// {$} special character to prevent subtree path pattern(anything that end with trailing slash) aka catch all
	mux.Handle("GET /{$}", dynamic.ThenFunc(app.home))
	mux.Handle("GET /static/", http.FileServerFS(ui.Files))
	mux.Handle("GET /snippet/view/{id}", dynamic.ThenFunc(app.viewSnippet))
	mux.Handle("GET /snippet/create", protected.ThenFunc(app.createSnippetForm))
	mux.Handle("POST /snippet/create", protected.ThenFunc(app.createSnippet))
	mux.Handle("POST /user/signup", dynamic.ThenFunc(app.signupUser))
	mux.Handle("GET /user/signup", dynamic.ThenFunc(app.signupForm))
	mux.Handle("POST /user/login", dynamic.ThenFunc(app.loginUser))
	mux.Handle("GET /user/login", dynamic.ThenFunc(app.loginUserForm))
	mux.Handle("POST /user/logout", dynamic.ThenFunc(app.logoutUser))

	standard := alice.New(app.recoverPanic, app.logRequest, commonHeaders)

	return standard.Then(mux)
}
