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

// 🔄 One-time cleanup function to remove duplicate items (same name + price)
func removeDuplicateItems() {
	var items []models.Item
	config.DB.Find(&items)

	seen := make(map[string]uint) // key = name_price, value = ID
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
	// ✅ Connect to DB
	config.ConnectDB()

	// 🧹 Run one-time duplicate cleanup (comment out after first successful run)
	removeDuplicateItems()

	// ✅ Insert default items ONLY if they don't already exist
	insertItemIfNotExists("T-Shirt", 299)
	insertItemIfNotExists("Sneakers", 1599)
	insertItemIfNotExists("Jeans", 799)
	insertItemIfNotExists("iPhone 15", 69999)
	insertItemIfNotExists("MacBook Air", 99999)
	insertItemIfNotExists("Apple Watch", 31999)
	insertItemIfNotExists("AirPods Pro", 24999)

	// ✅ Create Gin engine
	r := gin.Default()

	// ✅ Setup CORS middleware for 5173
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// ✅ Handle preflight OPTIONS requests
	r.OPTIONS("/*path", func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Authorization, Content-Type")
		c.AbortWithStatus(204)
	})

	// ✅ Register routes
	routes.SetupRoutes(r)

	// ✅ Start server
	if err := r.Run(":8080"); err != nil {
		panic("❌ Failed to start server: " + err.Error())
	}
}
