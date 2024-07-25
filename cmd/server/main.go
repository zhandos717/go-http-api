package main

import (
	"log"
	"net/http"

	"github.com/zhandos717/go-http-api/internal/middleware"
	"github.com/zhandos717/go-http-api/internal/routes"
)

func main() {
	r := routes.SetupRoutes()

	// Wrap the router with the logging middleware
	loggedRouter := middleware.LoggingMiddleware(r)

	log.Println("Server started at :8080")
	if err := http.ListenAndServe(":8080", loggedRouter); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
