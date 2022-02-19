package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

func InitDb() *gorm.DB {

	db, err := gorm.Open("sqlite3", "./data.db")

	db.LogMode(true)

	if err != nil {
		panic("failed to connect database")
	}

	// if !db.HasTable(&models.Payments{}) && !db.HasTable(&models.Categories{}) {
	// 	db.CreateTable(&models.Payments{}, &models.Categories{})
	// 	db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable()
	// }

	db.AutoMigrate(&Payments{}, &Categories{})

	return db
}
