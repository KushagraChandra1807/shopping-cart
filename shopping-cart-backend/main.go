package main

import (
	"fmt"
	"log"
	"time"

	"shopping-cart-backend/config"
	"shopping-cart-backend/models"
	"shopping-cart-backend/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func removeDuplicateItems() {
	var items []models.Item
	config.DB.Find(&items)

	seen := make(map[string]uint)
	var duplicates []uint

	for _, item := range items {
		key := fmt.Sprintf("%s_%.2f", item.Name, item.Price)
		if _, exists := seen[key]; exists {
			duplicates = append(duplicates, item.ID)
		} else {
			seen[key] = item.ID
		}
	}

	if len(duplicates) > 0 {
		if err := config.DB.Delete(&models.Item{}, duplicates).Error; err != nil {
			log.Fatalf("❌ Failed to delete duplicates: %v", err)
		}
		fmt.Printf("✅ Removed %d duplicate items.\n", len(duplicates))
	} else {
		fmt.Println("✅ No duplicate items found.")
	}
}

func insertItemIfNotExists(name string, price int) {
	var item models.Item
	if err := config.DB.Where("name = ?", name).First(&item).Error; err != nil {
		config.DB.Create(&models.Item{Name: name, Price: float64(price)})
	}
}

func main() {

	config.ConnectDB()

	removeDuplicateItems()

	insertItemIfNotExists("T-Shirt", 299)
	insertItemIfNotExists("Sneakers", 1599)
	insertItemIfNotExists("Jeans", 799)
	insertItemIfNotExists("iPhone 15", 69999)
	insertItemIfNotExists("MacBook Air", 99999)
	insertItemIfNotExists("Apple Watch", 31999)
	insertItemIfNotExists("AirPods Pro", 24999)

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.OPTIONS("/*path", func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Authorization, Content-Type")
		c.AbortWithStatus(204)
	})

	routes.SetupRoutes(r)

	if err := r.Run(":8080"); err != nil {
		panic("❌ Failed to start server: " + err.Error())
	}
}
