package server

import (
	"go-maro/web"
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		Render(w, "index.html", nil)
	})

	mux.HandleFunc("POST /{$}", func(w http.ResponseWriter, r *http.Request) {
		data := r.PostFormValue("alias")
		if data == "" {
			Render(w, "index.html", map[string]string{
				"AliasMessage": "alias cannot be empty!",
			})
			return
		}
		w.Write([]byte("alias: " + data))
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(r.URL.Path))
	})

	mux.HandleFunc("/output.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFileFS(w, r, web.AssetsFS, "assets/output.css")
	})
}
