package server

import (
	"context"
	"encoding/json"
	"go-maro/internal/db"
	"go-maro/web"
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

		InsertedLink, err := s.DB.CreateLink(context.Background(), db.CreateLinkParams{
			Source: source,
			Alias:  alias,
		})

		if err != nil {
			errMessage := map[string]string{
				"aliasError": "duplicate alias!",
			}
			json, _ := json.Marshal(errMessage)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(400)
			w.Write(json)
			return
		}

		w.Header().Set("content-type", "application/json")
		jsonMessage, _ := json.Marshal(map[string]string{
			"alias": InsertedLink.Alias,
		})
		w.Write(jsonMessage)

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
