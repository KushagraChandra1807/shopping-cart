package routes

import (
	"shopping-cart-backend/controllers"
	"shopping-cart-backend/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Welcome to the Shopping Cart API"})
	})

	r.POST("/users/login", controllers.LoginUser)
	r.GET("/items", controllers.GetItems)

	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())

	protected.POST("/carts", controllers.AddToCart)
	protected.GET("/carts", controllers.GetCart)
	protected.PUT("/carts", controllers.UpdateCartQuantity)

	protected.POST("/orders", controllers.CreateOrder)
	protected.GET("/orders", controllers.GetOrders)
}
