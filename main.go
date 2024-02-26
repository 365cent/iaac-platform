package main

import (
	"html/template"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var templates *template.Template

func homeHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"Message": "Hello, world!",
	}
	templates.ExecuteTemplate(w, "index.html", data)
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", homeHandler)

	http.ListenAndServe(":8080", r)
}
