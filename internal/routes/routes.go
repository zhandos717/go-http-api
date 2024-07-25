package routes

import (
	"net/http"

	"github.com/zhandos717/go-http-api/internal/controllers"
	"github.com/zhandos717/go-http-api/internal/middleware"
)

// Group is a struct that allows adding middlewares to groups of routes
type Group struct {
	prefix     string
	middleware []func(http.Handler) http.Handler
}

// NewGroup creates a new group with the specified prefix and middleware
func NewGroup(prefix string, middleware ...func(http.Handler) http.Handler) *Group {
	return &Group{
		prefix:     prefix,
		middleware: middleware,
	}
}

// HandleFunc registers a new route with a matcher for the URL path and wraps it with the group's middleware
func (g *Group) HandleFunc(mux *http.ServeMux, path string, handler http.HandlerFunc) {
	finalHandler := http.Handler(handler)
	for _, m := range g.middleware {
		finalHandler = m(finalHandler)
	}
	mux.Handle(g.prefix+path, finalHandler)
}

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	// Create a new group with logging and error handling middleware
	apiGroup := NewGroup("/api", middleware.LoggingMiddleware, func(next http.Handler) http.Handler {
		return http.HandlerFunc(middleware.ErrorHandler(next.ServeHTTP))
	})

	apiGroup.HandleFunc(mux, "/payment", middleware.Method(http.MethodPost, controllers.PostPayment))
	apiGroup.HandleFunc(mux, "/category", middleware.Method(http.MethodPost, controllers.PostCategory))
	apiGroup.HandleFunc(mux, "/payments", middleware.Method(http.MethodGet, controllers.GetPayments))
	apiGroup.HandleFunc(mux, "/categories", middleware.Method(http.MethodGet, controllers.GetCategories))
	apiGroup.HandleFunc(mux, "/payment/", middleware.Method(http.MethodGet, controllers.GetPayment))
	apiGroup.HandleFunc(mux, "/category/", middleware.Method(http.MethodGet, controllers.GetCategory))
	apiGroup.HandleFunc(mux, "/update/payment/", middleware.Method(http.MethodPut, controllers.UpdatePayment))
	apiGroup.HandleFunc(mux, "/update/category/", middleware.Method(http.MethodPut, controllers.UpdateCategory))
	apiGroup.HandleFunc(mux, "/delete/payment/", middleware.Method(http.MethodDelete, controllers.DeletePayment))
	apiGroup.HandleFunc(mux, "/delete/category/", middleware.Method(http.MethodDelete, controllers.DeleteCategory))

	mux.Handle("/", http.HandlerFunc(middleware.NotFoundHandler)) // Custom 404 handler

	return mux
}
