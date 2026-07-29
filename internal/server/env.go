package server

import (
	"log"
	"os"
)

func GetDatabaseUrl() string {
	url := os.Getenv("DATABASE_URL")
	if url == "" {
		log.Fatal("failed to get DATABASE_URL")
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
