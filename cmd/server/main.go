package main

import (
	"log"
	"net/http"
	"os"

	"github.com/zhandos717/go-http-api/internal/routes"
)

func main() {
	// Открытие или создание файла для логов
	logFile, err := os.OpenFile("server.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Could not open log file : %v", err)
	}
	defer logFile.Close()

	// Настройка логгера для записи в файл
	log.SetOutput(logFile)

	r := routes.SetupRoutes()

	log.Println("Server started at :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
