package models

import "time"

type Payments struct {
	Id        int    `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
	Name      string `gorm:"not null" form:"name" json:"name"`
	Price     uint64 `gorm:"not null" form:"price" json:"price"`
	Date      string `gorm:"not null" form:"date" json:"date"`
	Type      uint64 `gorm:"not null" form:"type" json:"type"`
	Comment   string `gorm:"not null" form:"comment" json:"comment"`
	Category  string `gorm:"not null" form:"category" json:"category"`
	CreatedAt time.Time
}
