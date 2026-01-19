package server

import (
	"go-maro/web"
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseFS(web.TemplatesFS, "templates/*.html"))

func Render(w http.ResponseWriter, name string, data any) {
	err := templates.ExecuteTemplate(w, name, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
