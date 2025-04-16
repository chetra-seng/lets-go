package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	files := []string{"./ui/html/base.tmpl.html", "./ui/html/pages/home.tmpl.html", "./ui/html/partials/nav.tmpl.html"}
	w.Header().Add("server", "Go")

	ts, err := template.ParseFiles(files...)

	if err != nil {
		app.serverError(w, r, err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}

	err = ts.ExecuteTemplate(w, "base", nil)

	if err != nil {
		app.serverError(w, r, err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}

}

func (app *application) createSnippetForm(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("Hello from create form"))
}

func (app *application) viewSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	msg := fmt.Sprintf("Hello from view %d", id)
	w.Write([]byte(msg))
}

func (app *application) createSnippet(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Hello from created")
}
