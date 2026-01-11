package main

import (
	"html/template"
	"net/http"
)

var templates *template.Template = template.Must(template.ParseGlob("templates/_*.html"))

func Render(w http.ResponseWriter, name string, data any) {
	templates.ParseFiles("templates/" + name)
	templates.ExecuteTemplate(w, "_layout.html", data)
}
