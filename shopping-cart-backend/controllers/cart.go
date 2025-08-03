package controllers

import (
	"net/http"
	"shopping-cart-backend/config"
	"shopping-cart-backend/models"

	"github.com/gin-gonic/gin"
)

// AddToCart adds or updates an item in the user's cart
func AddToCart(c *gin.Context) {
	userID := c.GetUint("user_id")

	var input struct {
		ItemID   uint `json:"item_id"`
		Quantity int  `json:"quantity"`
	}
	if err := c.ShouldBindJSON(&input); err != nil || input.Quantity <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	var cart models.Cart
	if err := config.DB.Where("user_id = ?", userID).First(&cart).Error; err != nil {
		cart = models.Cart{UserID: userID}
		config.DB.Create(&cart)
	}

	var cartItem models.CartItem
	if err := config.DB.Where("cart_id = ? AND item_id = ?", cart.ID, input.ItemID).First(&cartItem).Error; err == nil {
		cartItem.Quantity += input.Quantity
		config.DB.Save(&cartItem)
	} else {
		cartItem = models.CartItem{
			CartID:   cart.ID,
			ItemID:   input.ItemID,
			Quantity: input.Quantity,
		}
		config.DB.Create(&cartItem)
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Item added to cart"})
}

// GetCart returns the user's cart with items
func GetCart(c *gin.Context) {
	userID := c.GetUint("user_id")

	var cart models.Cart
	err := config.DB.Preload("CartItems.Item").Where("user_id = ?", userID).First(&cart).Error
	if err != nil {
		cart = models.Cart{UserID: userID}
		config.DB.Create(&cart)
		cart.CartItems = []models.CartItem{}
	}

	c.JSON(http.StatusOK, cart)
}

// UpdateCartQuantity updates the quantity of a specific item in the cart
func UpdateCartQuantity(c *gin.Context) {
	userID := c.GetUint("user_id")

	var input struct {
		ItemID   uint `json:"item_id"`
		Quantity int  `json:"quantity"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	var cart models.Cart
	if err := config.DB.Where("user_id = ?", userID).First(&cart).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cart not found"})
		return
	}

	var cartItem models.CartItem
	if err := config.DB.Where("cart_id = ? AND item_id = ?", cart.ID, input.ItemID).First(&cartItem).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found in cart"})
		return
	}

	if input.Quantity <= 0 {
		config.DB.Delete(&cartItem)
		c.JSON(http.StatusOK, gin.H{"message": "Item removed from cart"})
	} else {
		cartItem.Quantity = input.Quantity
		config.DB.Save(&cartItem)
		c.JSON(http.StatusOK, gin.H{"message": "Quantity updated"})
	}
}
