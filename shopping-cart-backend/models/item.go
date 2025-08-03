package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	Name  string  `gorm:"uniqueIndex:idx_name_price" json:"name"`
	Price float64 `gorm:"uniqueIndex:idx_name_price" json:"price"`
}
