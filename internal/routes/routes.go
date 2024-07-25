package routes

import (
	"net/http"

	"github.com/zhandos717/go-http-api/internal/controllers"
)

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/payment", controllers.PostPayment)
	mux.HandleFunc("/api/category", controllers.PostCategory)
	mux.HandleFunc("/api/payments", controllers.GetPayments)
	mux.HandleFunc("/api/categories", controllers.GetCategories)
	mux.HandleFunc("/api/payment/", controllers.GetPayment)
	mux.HandleFunc("/api/category/", controllers.GetCategory)
	mux.HandleFunc("/api/update/payment/", controllers.UpdatePayment)
	mux.HandleFunc("/api/update/category/", controllers.UpdateCategory)
	mux.HandleFunc("/api/delete/payment/", controllers.DeletePayment)
	mux.HandleFunc("/api/delete/category/", controllers.DeleteCategory)

	return mux
}
