package main

import (
	"embed"
	"net/http"
)

//go:embed assets/output.css
var cssFS embed.FS

func registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/{$}", func(w http.ResponseWriter, r *http.Request) {
		Render(w, "index.html", nil)
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(r.URL.Path))
	})

	mux.HandleFunc("/output.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFileFS(w, r, cssFS, "assets/output.css")
	})
}
