package controllers

import (
	"net/http"
	"shopping-cart-backend/config"
	"shopping-cart-backend/models"

	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
	userID := c.GetUint("user_id")

	var cart models.Cart
	if err := config.DB.Preload("CartItems").Where("user_id = ?", userID).First(&cart).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cart not found"})
		return
	}

	orderCart := models.Cart{UserID: userID}
	if err := config.DB.Create(&orderCart).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order cart"})
		return
	}

	for _, item := range cart.CartItems {
		newItem := models.CartItem{
			CartID:   orderCart.ID,
			ItemID:   item.ItemID,
			Quantity: item.Quantity,
		}
		config.DB.Create(&newItem)
	}

	order := models.Order{UserID: userID, CartID: orderCart.ID}
	if err := config.DB.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
		return
	}

	config.DB.Where("cart_id = ?", cart.ID).Delete(&models.CartItem{})
	c.JSON(http.StatusOK, gin.H{"message": "Order placed successfully", "order_id": order.ID})
}

func GetOrders(c *gin.Context) {
	userID := c.GetUint("user_id")

	var orders []models.Order
	err := config.DB.
		Preload("Cart").
		Preload("Cart.CartItems").
		Preload("Cart.CartItems.Item").
		Where("user_id = ?", userID).
		Find(&orders).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch orders"})
		return
	}

	c.JSON(http.StatusOK, orders)
}
