package main

import (
	"context"
	"go-maro/internal/db"
	"go-maro/internal/middleware"
	"go-maro/internal/server"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

func main() {
	pool, err := pgxpool.New(context.Background(), server.GetDatabaseUrl())
	if err != nil {
		log.Fatalln("Unable to create connection pool:", err.Error())
	}

	port := server.GetPort()

	goose.SetBaseFS(db.EmbedMigrations)
	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatalln(err)
	}

	if err = goose.Up(stdlib.OpenDBFromPool(pool), "migrations"); err != nil {
		log.Fatalln(err)
	}

	queries := db.New(pool)
	srv := server.NewServer(queries)

	log.Println("Listening on http://localhost" + port)
	http.ListenAndServe(port, middleware.Logger(srv.Mux))
}
