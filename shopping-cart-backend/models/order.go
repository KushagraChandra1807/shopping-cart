package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID uint `json:"user_id"`
	CartID uint `json:"cart_id"`

	Cart Cart `gorm:"foreignKey:CartID" json:"cart"`
}
