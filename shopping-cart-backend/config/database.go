package config

import (
	"log"
	"shopping-cart-backend/models"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// Connect to SQLite
	database, err := gorm.Open(sqlite.Open("shopping.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed to connect database:", err)
	}
	DB = database

	// Auto migrate all models
	err = DB.AutoMigrate(
		&models.User{},
		&models.Item{},
		&models.Cart{},
		&models.CartItem{},
		&models.Order{},
	)
	if err != nil {
		log.Fatal("❌ Auto migration failed:", err)
	}

	log.Println("✅ Database connected and migrated")
}
