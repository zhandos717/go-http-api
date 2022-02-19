package controllers

import (
	"github.com/gin-gonic/gin"
)

func PostPayment(c *gin.Context) {

	db := InitDb()
	defer db.Close()
	var payment Payments
	c.Bind(&payment)

	if payment.Name != "" && payment.Comment != "" && payment.Price > 0 && payment.Category != "" {
		db.Create(&payment)
		c.JSON(201, gin.H{"success": payment})
	} else {
		c.JSON(422, gin.H{"error": payment})
	}
}

func GetPayment(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	id := c.Params.ByName("id")
	var payment Payments

	db.First(&payment, id)

	if payment.Id != 0 {
		c.JSON(200, payment)
	} else {
		c.JSON(404, gin.H{"error": "User not found"})
	}
}

func UpdatePayment(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	id := c.Params.ByName("id")
	var payment Payments
	db.First(&payment, id)
	if payment.Name != "" && payment.Comment != "" && payment.Price > 0 && payment.Category != "" {
		if payment.Id != 0 {
			var newPayment Payments
			c.Bind(&newPayment)
			result := Payments{
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
	db := InitDb()

	defer db.Close()

	id := c.Params.ByName("id")
	var payment Payments
	db.First(&payment, id)

	if payment.Id != 0 {
		db.Delete(&payment)
		c.JSON(200, gin.H{"success": "User #" + id + " deleted"})
	} else {
		c.JSON(404, gin.H{"error": "User not found"})
	}
}
