package main

import (
	"html/template"
	"io/fs"
	"path/filepath"
	"time"

	"chetraseng.com/internal/models"
	"chetraseng.com/ui"
)

type templateData struct {
	CurrentYear     int
	Snippet         models.Snippet
	Snippets        []models.Snippet
	Form            any
	Flash           string
	IsAuthenticated bool
	CSRFToken       string
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
	pages, err := fs.Glob(ui.Files, "html/pages/*.tmpl.html")

	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		// get only the base path aka name
		name := filepath.Base(page)

		// Create a slice to hold all patterns
		patterns := []string{
			"html/base.tmpl.html",
			"html/partials/*.tmpl.html",
			page,
		}

		// Parse base first
		// Register function map before calling parseFiles
		ts, err := template.New(name).Funcs(function).ParseFS(ui.Files, patterns...)

		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}
	return cache, nil
}
