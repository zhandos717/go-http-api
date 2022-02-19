package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/zhandos717/go-http-api/models"
)

func PostPayment(c *gin.Context) {

	db := models.InitDb()
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
func GetPayment(c *gin.Context) {
	db := models.InitDb()
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
func GetPayments(c *gin.Context) {
	db := models.InitDb()
	defer db.Close()
	var payment []models.Payments
	db.Find(&payment)
	c.JSON(200, payment)
}

func UpdatePayment(c *gin.Context) {
	db := models.InitDb()
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
func DeletePayment(c *gin.Context) {
	db := models.InitDb()

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
