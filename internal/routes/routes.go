package routes

import (
	"net/http"

	"github.com/zhandos717/go-http-api/internal/controllers"
	"github.com/zhandos717/go-http-api/internal/middleware"
)

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/payment", middleware.ErrorHandler(middleware.Method(http.MethodPost, controllers.PostPayment)))
	mux.HandleFunc("/api/category", middleware.ErrorHandler(middleware.Method(http.MethodPost, controllers.PostCategory)))
	mux.HandleFunc("/api/payments", middleware.ErrorHandler(middleware.Method(http.MethodGet, controllers.GetPayments)))
	mux.HandleFunc("/api/categories", middleware.ErrorHandler(middleware.Method(http.MethodGet, controllers.GetCategories)))
	mux.HandleFunc("/api/payment/", middleware.ErrorHandler(middleware.Method(http.MethodGet, controllers.GetPayment)))
	mux.HandleFunc("/api/category/", middleware.ErrorHandler(middleware.Method(http.MethodGet, controllers.GetCategory)))
	mux.HandleFunc("/api/update/payment/", middleware.ErrorHandler(middleware.Method(http.MethodPut, controllers.UpdatePayment)))
	mux.HandleFunc("/api/update/category/", middleware.ErrorHandler(middleware.Method(http.MethodPut, controllers.UpdateCategory)))
	mux.HandleFunc("/api/delete/payment/", middleware.ErrorHandler(middleware.Method(http.MethodDelete, controllers.DeletePayment)))
	mux.HandleFunc("/api/delete/category/", middleware.ErrorHandler(middleware.Method(http.MethodDelete, controllers.DeleteCategory)))

	mux.Handle("/", http.HandlerFunc(middleware.NotFoundHandler)) // Custom 404 handler

	return mux
}
