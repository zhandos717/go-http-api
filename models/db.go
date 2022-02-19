package models

import "github.com/jinzhu/gorm"

func InitDb() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./data.db")
	db.LogMode(true)
	if err != nil {
		panic("failed to connect database")
	}
	if !db.HasTable(&Payments{}) && !db.HasTable(&Categories{}) {
		db.CreateTable(&Payments{}, &Categories{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Payments{}, &Categories{})
	}

	return db
}
