package main

import (
	"github.com/zhandos717/go-http-api/controllers"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	HandleFunc()
}
func HandleFunc() {
	r := gin.Default()
	api := r.Group("api")
	{

		api.POST("/payment", controllers.PostPayment)
		api.POST("/category", controllers.PostCategory)

		api.GET("/payments", controllers.GetPayments)
		api.GET("/categories", controllers.GetCategories)

		api.GET("/payment/:id", controllers.GetPayment)
		api.GET("/category/:id", controllers.GetCategory)

		api.PUT("/payment/:id", controllers.UpdatePayment)
		api.PUT("/category/:id", controllers.UpdateCategory)

		api.DELETE("/category/:id", controllers.DeletePayment)
		api.DELETE("/payment/:id", controllers.DeleteCategory)
	}
	r.Run(":8080")
}
