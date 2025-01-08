package web

import (
	"database/sql"
	"shopping-api-backend-go/internal/handlers"

	"github.com/gin-gonic/gin"
)

// InitializeRouter initializes the routes and returns a Gin engine with all routes set up
func InitializeRouter(db *sql.DB) *gin.Engine {
	r := gin.Default()

	// Health Check Endpoint
	r.GET("/health", handlers.HealthCheck)

	// CRUD routes for shopping items
	r.GET("/api/shoppingItems/:name", handlers.GetItemByName)
	r.PUT("/api/shoppingItems/:name", handlers.UpdateItem)
	r.DELETE("/api/shoppingItems/:name", handlers.DeleteItem)
	r.GET("/api/shoppingItems", handlers.GetAllItems)
	r.POST("/api/shoppingItems", handlers.AddItem)

	return r
}
