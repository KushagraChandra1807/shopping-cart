package routes

import (
	"shopping-cart-backend/controllers"
	"shopping-cart-backend/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// 🌐 Public route
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Welcome to the Shopping Cart API"})
	})

	// 🟢 Public API routes
	r.POST("/users/login", controllers.LoginUser) // Login endpoint
	r.GET("/items", controllers.GetItems)         // Fetch all items

	// 🔒 Protected routes with JWT middleware
	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())

	// 🛒 Cart routes
	protected.POST("/carts", controllers.AddToCart)         // Add item to cart
	protected.GET("/carts", controllers.GetCart)            // View cart
	protected.PUT("/carts", controllers.UpdateCartQuantity) // Update quantity or remove item

	// 📦 Order routes
	protected.POST("/orders", controllers.CreateOrder) // Create order from cart
	protected.GET("/orders", controllers.GetOrders)    // View user's orders
}
