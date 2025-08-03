package controllers

import (
	"fmt"
	"net/http"
	"shopping-cart-backend/config"
	"shopping-cart-backend/models"

	"github.com/gin-gonic/gin"
)

// CreateItem handles POST /items
func CreateItem(c *gin.Context) {
	var item models.Item

	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item data"})
		return
	}

	var existing models.Item
	if err := config.DB.Where("name = ? AND price = ?", item.Name, item.Price).First(&existing).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Item already exists"})
		return
	}

	config.DB.Create(&item)
	c.JSON(http.StatusCreated, item)
}

// GetItems handles GET /items
func GetItems(c *gin.Context) {
	var items []models.Item
	config.DB.Find(&items)

	uniqueMap := make(map[string]models.Item)
	for _, item := range items {
		key := item.Name + "_" + fmt.Sprintf("%.2f", item.Price)
		if _, exists := uniqueMap[key]; !exists {
			uniqueMap[key] = item
		}
	}

	var uniqueItems []models.Item
	for _, item := range uniqueMap {
		uniqueItems = append(uniqueItems, item)
	}

	c.JSON(http.StatusOK, uniqueItems)
}
