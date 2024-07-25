package middleware

import (
	"log"
	"net/http"
	"time"
)

// Method is a middleware function to enforce HTTP method
func Method(method string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			NotFoundHandler(w, r)
			return
		}
		next(w, r)
	}
}

// NotFoundHandler is a custom 404 handler
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Header().Set("Content-Type", "application/json")
	response := `{"error": "Resource not found"}`
	w.Write([]byte(response))
	log.Printf("404 - %s %s %s", r.Method, r.RequestURI, r.RemoteAddr)
}

// LoggingMiddleware logs the details of each request
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("%s %s %s", r.Method, r.RequestURI, r.RemoteAddr)
		next.ServeHTTP(w, r)
		log.Printf("Completed in %v", time.Since(start))
	})
}

// ErrorHandler is a middleware that catches errors and logs them
func ErrorHandler(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Error: %v", err)
				w.WriteHeader(http.StatusInternalServerError)
				w.Header().Set("Content-Type", "application/json")
				response := `{"error": "Internal Server Error"}`
				w.Write([]byte(response))
			}
		}()
		next(w, r)
	}
}
