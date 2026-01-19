package server

import (
	"context"
	"go-maro/internal/db"
	"go-maro/web"
	"log"
	"net/http"
	"strings"
)

func (s *Server) RegisterRoutes() {
	s.Mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		Render(w, "index.html", nil)
	})

	s.Mux.HandleFunc("POST /{$}", func(w http.ResponseWriter, r *http.Request) {
		message := IndexData{}
		source := r.PostFormValue("source")
		alias := r.PostFormValue("alias")

		message.validateSource(source)
		message.validateAlias(alias)

		if message.isValid() == false {
			message.ServerMessage = ""
			Render(w, "index.html", message)
			return
		}

		insertedLink, err := s.DB.CreateLink(context.Background(), db.CreateLinkParams{
			Source: source,
			Alias:  alias,
		})

		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(400)
			Render(w, "index.html", nil)
			return
		}

		Render(w, "index.html", IndexData{
			ServerMessage: "http://" + r.Host + "/" + insertedLink.Alias,
		})

	})

	s.Mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimPrefix(r.URL.Path, "/")
		fetchedLink, err := s.DB.GetLink(context.Background(), path)
		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte("Not Found"))
			return
		}

		http.Redirect(w, r, "https://"+fetchedLink.Source, http.StatusFound)
	})

	s.Mux.HandleFunc("/output.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFileFS(w, r, web.AssetsFS, "assets/output.css")
	})
}
