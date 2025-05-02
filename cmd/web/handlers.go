package main

import (
	"errors"
	"fmt"

	"chetraseng.com/internal/models"
	"chetraseng.com/internal/validator"

	"net/http"
	"strconv"
)

// Map each fields to form field name for decoder
// Form decoder needs to know these field names
// Ingore struct with validator
type snippetCreateForm struct {
	Title               string `form:"title"`
	Content             string `form:"content"`
	Expires             int    `form:"expires"`
	validator.Validator `form:"-"`
}

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
	data.Form = snippetCreateForm{
		Expires: 365,
	}
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

	var form snippetCreateForm

	// Decode the form automatically from the request without having to do it manually
	err = app.decodePostForm(r, &form)

	// Incase of error in decoding
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form.CheckField(validator.NotBlank(form.Title), "title", "This field can not be blank")
	form.CheckField(validator.MaxChars(form.Title, 100), "title", "This field can not be more than 100 characters long")
	form.CheckField(validator.NotBlank(form.Content), "content", "This field can not be blank")
	form.CheckField(validator.PermittedValue(form.Expires, 1, 7, 365), "expires", "This field must equal 1, 7 or 365")

	if !form.Valid() {
		data := app.newTemplateData(r)
		data.Form = form
		app.render(w, r, http.StatusUnprocessableEntity, "create.tmpl.html", data)
		return
	}

	id, err := app.snippets.Insert(form.Title, form.Content, form.Expires)

	if err != nil {
		app.serverError(w, r, err)
		return
	}

	app.sessionManager.Put(r.Context(), "flash", "Snippet successfully created!")

	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
}
