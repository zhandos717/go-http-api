package models

import "time"

type Categories struct {
	Id        int    `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
	Name      string `gorm:"not null" form:"name" json:"name"`
	Type      uint64 `gorm:"not null" form:"type" json:"type"`
	CreatedAt time.Time
}
