package main

import "net/http"

func registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		Render(w, "index.html", nil)
	})
}
