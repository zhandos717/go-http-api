package main

import (
	"github.com/zhandos717/go-http-api/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
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

func InitDb() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./data.db")
	db.LogMode(true)
	if err != nil {
		panic("failed to connect database")
	}
	if !db.HasTable(&models.Payments{}) && !db.HasTable(&models.Categories{}) {
		db.CreateTable(&models.Payments{}, &models.Categories{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&models.Payments{}, &models.Categories{})
	}

	return db
}

func PostCategory(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	var category models.Categories
	c.Bind(&category)
	if category.Name != "" && category.Type > 0 && category.Type < 3 {
		db.Create(&category)
		c.JSON(201, gin.H{"success": category})
	} else {
		c.JSON(422, gin.H{"error": category})
	}
}
func PostPayment(c *gin.Context) {

	db := InitDb()
	defer db.Close()
	var payment models.Payments
	c.Bind(&payment)

	if payment.Name != "" && payment.Comment != "" && payment.Price > 0 && payment.Category != "" {
		db.Create(&payment)
		c.JSON(201, gin.H{"success": payment})
	} else {
		c.JSON(422, gin.H{"error": payment})
	}
}

func GetCategory(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	id := c.Params.ByName("id")
	var category models.Categories

	db.First(&category, id)

	if category.Id != 0 {
		c.JSON(200, category)
	} else {
		c.JSON(404, gin.H{"error": "User not found"})
	}
}

func GetPayment(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	id := c.Params.ByName("id")
	var payment models.Payments

	db.First(&payment, id)

	if payment.Id != 0 {
		c.JSON(200, payment)
	} else {
		c.JSON(404, gin.H{"error": "User not found"})
	}
}

func GetCategories(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	var category []models.Categories
	db.Find(&category)
	c.JSON(200, category)
}

func GetPayments(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	var payment []models.Payments
	db.Find(&payment)
	c.JSON(200, payment)
}

func UpdatePayment(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	id := c.Params.ByName("id")
	var payment models.Payments
	db.First(&payment, id)
	if payment.Name != "" && payment.Comment != "" && payment.Price > 0 && payment.Category != "" {
		if payment.Id != 0 {
			var newPayment models.Payments
			c.Bind(&newPayment)
			result := models.Payments{
				Id:       payment.Id,
				Name:     newPayment.Name,
				Comment:  newPayment.Comment,
				Price:    newPayment.Price,
				Category: newPayment.Category,
			}
			db.Save(&result)
			c.JSON(200, gin.H{"success": result})
		} else {
			c.JSON(404, gin.H{"error": "User not found"})
		}
	} else {
		c.JSON(422, gin.H{"error": "Fields are empty"})
	}
}

func UpdateCategory(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	id := c.Params.ByName("id")
	var category models.Categories
	db.First(&category, id)
	if category.Name != "" && category.Type > 0 && category.Type < 3 {
		if category.Id != 0 {
			var newCategory models.Categories
			c.Bind(&newCategory)
			result := models.Payments{
				Id:   category.Id,
				Name: newCategory.Name,
				Type: newCategory.Type,
			}
			db.Save(&result)
			c.JSON(200, gin.H{"success": result})
		} else {
			c.JSON(404, gin.H{"error": "User not found"})
		}
	} else {
		c.JSON(422, gin.H{"error": "Fields are empty"})
	}
}

func DeletePayment(c *gin.Context) {
	db := InitDb()

	defer db.Close()

	id := c.Params.ByName("id")
	var payment models.Payments
	db.First(&payment, id)

	if payment.Id != 0 {
		db.Delete(&payment)
		c.JSON(200, gin.H{"success": "User #" + id + " deleted"})
	} else {
		c.JSON(404, gin.H{"error": "User not found"})
	}
}

func DeleteCategory(c *gin.Context) {
	db := InitDb()

	defer db.Close()

	id := c.Params.ByName("id")
	var category models.Categories
	db.First(&category, id)

	if category.Id != 0 {
		db.Delete(&category)
		c.JSON(200, gin.H{"success": "User #" + id + " deleted"})
	} else {
		c.JSON(404, gin.H{"error": "User not found"})
	}
}
