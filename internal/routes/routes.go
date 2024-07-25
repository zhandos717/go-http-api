package routes

import (
	"net/http"

	"github.com/zhandos717/go-http-api/internal/controllers"
	"github.com/zhandos717/go-http-api/internal/middleware"
)

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/payment", middleware.Method(http.MethodPost, controllers.PostPayment))
	mux.HandleFunc("/api/category", middleware.Method(http.MethodPost, controllers.PostCategory))
	mux.HandleFunc("/api/payments", middleware.Method(http.MethodGet, controllers.GetPayments))
	mux.HandleFunc("/api/categories", middleware.Method(http.MethodGet, controllers.GetCategories))
	mux.HandleFunc("/api/payment/", middleware.Method(http.MethodGet, controllers.GetPayment))
	mux.HandleFunc("/api/category/", middleware.Method(http.MethodGet, controllers.GetCategory))
	mux.HandleFunc("/api/update/payment/", middleware.Method(http.MethodPut, controllers.UpdatePayment))
	mux.HandleFunc("/api/update/category/", middleware.Method(http.MethodPut, controllers.UpdateCategory))
	mux.HandleFunc("/api/delete/payment/", middleware.Method(http.MethodDelete, controllers.DeletePayment))
	mux.HandleFunc("/api/delete/category/", middleware.Method(http.MethodDelete, controllers.DeleteCategory))

	return mux
}
