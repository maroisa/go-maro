package main

import (
	"go-maro/internal/server"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	PORT := ":3000"

	server.RegisterRoutes(mux)

	log.Println("Listening on http://localhost" + PORT)
	err := http.ListenAndServe(PORT, server.Logger(mux))
	if err != nil {
		log.Fatalln(err)
	}
}
