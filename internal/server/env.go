package server

import (
	"log"
	"os"
)

func GetDatabaseUrl() string {
	url := os.Getenv("DATABASE_URL")
	if url == "" {
		url = "postgres://rin:blank@127.0.0.1:5432/go_maro"
		log.Println("DATABASE_URL is not set. Default to:", url)
	}

	return url
}
func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":3000"
		log.Println("PORT is not set. Default to :3000")
	}
	return port
}
