package main

import (
	"html/template"
	"path/filepath"
	"time"

	"chetraseng.com/internal/models"
)

type templateData struct {
	CurrentYear int
	Snippet     models.Snippet
	Snippets    []models.Snippet
	Form        any
	Flash       string
}

func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}

// Add function to template function map before using
var function = template.FuncMap{
	"humanDate": humanDate,
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	// all pages with full path
	pages, err := filepath.Glob("./ui/html/pages/*.tmpl.html")

	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		// get only the base path aka name
		name := filepath.Base(page)

		// Parse base first
		// Register function map before calling parseFiles
		ts, err := template.New(name).Funcs(function).ParseFiles("./ui/html/base.tmpl.html")

		if err != nil {
			return nil, err
		}

		// Parse all partials
		ts, err = ts.ParseGlob("./ui/html/partials/*.tmpl.html")

		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseFiles(page)

		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}
	return cache, nil
}
