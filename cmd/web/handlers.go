package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, _ *http.Request) {
	files := []string{"./ui/html/base.tmpl.html", "./ui/html/pages/home.tmpl.html", "./ui/html/partials/nav.tmpl.html"}
	w.Header().Add("server", "Go")

	ts, err := template.ParseFiles(files...)

	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}

	err = ts.ExecuteTemplate(w, "base", nil)

	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}

}

func createSnippetForm(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("Hello from create form"))
}

func viewSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	msg := fmt.Sprintf("Hello from view %d", id)
	w.Write([]byte(msg))
}

func createSnippet(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Hello from created")
}
