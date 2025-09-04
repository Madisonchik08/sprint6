package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "HTTP Server: ", log.Ldate|log.Ltime|log.Lshortfile)

	srv := server.NewServer(logger)
	logger.Println("Starting HTTP server")
	if err := srv.Start(); err != nil {
		logger.Fatalf("Error start server: %v", err)
	}
}
