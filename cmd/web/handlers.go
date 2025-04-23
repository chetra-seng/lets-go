package main

import (
	"errors"
	"fmt"

	"chetraseng.com/internal/models"

	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	snippets, err := app.snippets.Latest()

	if err != nil {
		app.serverError(w, r, err)
		return
	}

	tmplData := app.newTemplateData(r)
	tmplData.Snippets = snippets
	app.render(w, r, http.StatusOK, "home.tmpl.html", tmplData)

}

func (app *application) createSnippetForm(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	app.render(w, r, 200, "create.tmpl.html", data)
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

	tmplData := app.newTemplateData(r)
	tmplData.Snippet = s

	app.render(w, r, http.StatusOK, "view.tmpl.html", tmplData)
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	title := r.PostForm.Get("title")
	content := r.PostForm.Get("content")
	expires, err := strconv.Atoi(r.PostForm.Get("expires"))

	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	id, err := app.snippets.Insert(title, content, expires)

	if err != nil {
		app.serverError(w, r, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
}
