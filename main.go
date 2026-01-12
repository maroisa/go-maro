package main

import (
	"go-maro/middleware"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	PORT := ":3000"

	registerRoutes(mux)

	log.Println("Listening on http://localhost" + PORT)
	err := http.ListenAndServe(PORT, middleware.Logger(mux))
	if err != nil {
		log.Fatalln(err)
	}
}
