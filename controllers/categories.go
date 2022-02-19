package controllers

import (
	"github.com/gin-gonic/gin"
)

func PostCategory(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	var category Categories
	c.Bind(&category)
	if category.Name != "" && category.Type > 0 && category.Type < 3 {
		db.Create(&category)
		c.JSON(201, gin.H{"success": category})
	} else {
		c.JSON(422, gin.H{"error": category})
	}
}

func GetCategory(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	id := c.Params.ByName("id")
	var category Categories

	db.First(&category, id)

	if category.Id != 0 {
		c.JSON(200, category)
	} else {
		c.JSON(404, gin.H{"error": "User not found"})
	}
}

func GetCategories(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	var category []Categories
	db.Find(&category)
	c.JSON(200, category)
}

func GetPayments(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	var payment []Payments
	db.Find(&payment)
	c.JSON(200, payment)
}

func UpdateCategory(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	id := c.Params.ByName("id")
	var category Categories
	db.First(&category, id)
	if category.Name != "" && category.Type > 0 && category.Type < 3 {
		if category.Id != 0 {
			var newCategory Categories
			c.Bind(&newCategory)
			result := Payments{
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

func DeleteCategory(c *gin.Context) {
	db := InitDb()

	defer db.Close()

	id := c.Params.ByName("id")
	var category Categories
	db.First(&category, id)

	if category.Id != 0 {
		db.Delete(&category)
		c.JSON(200, gin.H{"success": "User #" + id + " deleted"})
	} else {
		c.JSON(404, gin.H{"error": "User not found"})
	}
}
