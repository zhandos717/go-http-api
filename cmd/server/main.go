package main

import (
	"log"
	"net/http"

	"github.com/zhandos717/go-http-api/internal/routes"
)

func main() {
	r := routes.SetupRoutes()
	log.Println("Server started at :8080")
	http.ListenAndServe(":8080", r)
}
