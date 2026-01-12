package main

import (
	"embed"
	"html/template"
	"net/http"
)

//go:embed templates/*.html
var templatesFS embed.FS

var templates = template.Must(template.ParseFS(templatesFS, "templates/*.html"))

func Render(w http.ResponseWriter, name string, data any) {
	err := templates.ExecuteTemplate(w, name ,data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
