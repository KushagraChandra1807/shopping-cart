package models

import "gorm.io/gorm"

type CartItem struct {
	gorm.Model
	CartID   uint
	ItemID   uint
	Quantity int

	Item Item `gorm:"foreignKey:ItemID"`
}
