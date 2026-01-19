package main

import (
	"context"
	"go-maro/internal/db"
	"go-maro/internal/server"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	pool, err := pgxpool.New(context.Background(), server.GetDatabaseUrl())
	if err != nil {
		log.Fatalln("Unable to create connection pool:", err.Error())
	}

	port := server.GetPort()

	queries := db.New(pool)
	srv := server.NewServer(queries)

	log.Println("Listening on port", port)
	http.ListenAndServe(port, server.Logger(srv.Mux))
}
