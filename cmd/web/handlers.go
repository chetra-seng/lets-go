package main

import (
	"chetraseng.com/internal/models"
	"errors"
	"fmt"
	// "html/template"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("server", "Go")
	// files := []string{"./ui/html/base.tmpl.html", "./ui/html/pages/home.tmpl.html", "./ui/html/partials/nav.tmpl.html"}
	//
	// ts, err := template.ParseFiles(files...)
	//
	// if err != nil {
	// 	app.serverError(w, r, err)
	// 	http.Error(w, "Internal server error", http.StatusInternalServerError)
	// }

	// err = ts.ExecuteTemplate(w, "base", nil)

	// if err != nil {
	// 	app.serverError(w, r, err)
	// 	http.Error(w, "Internal server error", http.StatusInternalServerError)
	// }

	snippets, err := app.snippets.Latest()

	if err != nil {
		app.serverError(w, r, err)
		return
	}
	for _, s := range snippets {
		fmt.Fprintf(w, "%+v\n", s)
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

	s, err := app.snippets.Get(id)

	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			http.NotFound(w, r)
		} else {
			app.serverError(w, r, err)
		}

		return
	}

	fmt.Fprintf(w, "%+v", s)
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
	expires := 7

	id, err := app.snippets.Insert(title, content, expires)

	if err != nil {
		app.serverError(w, r, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
}
