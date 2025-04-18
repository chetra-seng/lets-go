package main

import (
	"html/template"
	"path/filepath"

	"chetraseng.com/internal/models"
)

type templateData struct {
	Snippet  models.Snippet
	Snippets []models.Snippet
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
		ts, err := template.ParseFiles("./ui/html/base.tmpl.html")

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
