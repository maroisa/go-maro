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
		message := IndexData{}
		source := r.PostFormValue("source")
		alias := r.PostFormValue("alias")

		message.validateSource(source)
		message.validateAlias(alias)

		if message.isValid() == false {
			Render(w, "index.html", message)
			return
		}

	})

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(r.URL.Path))
	})

	mux.HandleFunc("/output.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFileFS(w, r, web.AssetsFS, "assets/output.css")
	})
}
