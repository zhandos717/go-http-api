package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	HandleFunc()
}
func HandleFunc() {
	r := gin.Default()
	api := r.Group("api")
	{

		api.POST("/payment", PostPayment)
		api.POST("/category", PostCategory)

		api.GET("/payments", GetPayments)
		api.GET("/categories", GetCategories)

		api.GET("/payment/:id", GetPayment)
		api.GET("/category/:id", GetCategory)

		api.PUT("/payment/:id", UpdatePayment)
		api.PUT("/category/:id", UpdateCategory)

		api.DELETE("/category/:id", DeletePayment)
		api.DELETE("/payment/:id", DeleteCategory)
	}
	r.Run(":8080")
}
