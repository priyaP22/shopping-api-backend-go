package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "shopping-api-backend-go/docs" // Import Swagger docs

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq" // PostgreSQL driver
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// ShoppingItem represents a shopping item model
type ShoppingItem struct {
	Name   string `json:"name" example:"Milk"`
	Amount int    `json:"amount" example:"2"`
}

var db *sql.DB

// @title Shopping API
// @version 1.0
// @description A simple API to manage shopping items with PostgreSQL
// @contact.name Your Name
// @contact.email yourname@example.com
// @host localhost:8080
// @BasePath /api
func main() {
	var err error

	// Connect to the PostgreSQL database
	dbHost := os.Getenv("POSTGRES_HOST")
	dbPort := os.Getenv("POSTGRES_PORT")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")

	dbURL := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)
	db, err = sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	// Verify the connection
	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	// Create the shopping_items table if it doesn't exist
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS shopping_items (
			name TEXT PRIMARY KEY,
			amount INTEGER NOT NULL
		)
	`)
	if err != nil {
		log.Fatalf("Failed to create the shopping_items table: %v", err)
	}

	r := gin.Default()

	// CORS Configuration
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5000"}, // Replace with your frontend's URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
	}))

	// Swagger Endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Health Check Endpoint
	r.GET("/health", func(c *gin.Context) {
		if err := db.Ping(); err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{"error": "Database connection failed"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "OK"})
	})

	// Routes
	r.GET("/api/shoppingItems/:name", getItemByName)
	r.PUT("/api/shoppingItems/:name", updateItem)
	r.DELETE("/api/shoppingItems/:name", deleteItem)
	r.GET("/api/shoppingItems", getAllItems)
	r.POST("/api/shoppingItems", addItem)

	// Start server
	r.Run() // Default is localhost:8080
}

// @Summary Get a shopping item by name
// @Description Retrieve a specific shopping item by its name
// @Tags Shopping Items API
// @Param name path string true "Item name"
// @Success 200 {object} ShoppingItem
// @Failure 404 {object} gin.H{"error": "Item not found"}
// @Router /api/shoppingItems/{name} [get]
func getItemByName(c *gin.Context) {
	name := c.Param("name")

	var item ShoppingItem
	err := db.QueryRow("SELECT name, amount FROM shopping_items WHERE name = $1", name).Scan(&item.Name, &item.Amount)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve item"})
		return
	}

	c.JSON(http.StatusOK, item)
}

// @Summary Update a shopping item by name
// @Description Update a specific shopping item by its name
// @Tags Shopping Items API
// @Param name path string true "Item name"
// @Param shoppingItem body ShoppingItem true "Updated shopping item"
// @Success 200 {object} ShoppingItem
// @Failure 404 {object} gin.H{"error": "Item not found"}
// @Router /api/shoppingItems/{name} [put]
func updateItem(c *gin.Context) {
	name := c.Param("name")
	var updatedItem ShoppingItem
	if err := c.ShouldBindJSON(&updatedItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Input validation
	if updatedItem.Amount <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Amount must be greater than zero"})
		return
	}

	result, err := db.Exec("UPDATE shopping_items SET amount = $1 WHERE name = $2", updatedItem.Amount, name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update item"})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check affected rows"})
		return
	}
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	c.JSON(http.StatusOK, updatedItem)
}

// @Summary Delete a shopping item by name
// @Description Delete a specific shopping item by its name
// @Tags Shopping Items API
// @Param name path string true "Item name"
// @Success 200 {object} gin.H{"message": "Item deleted"}
// @Failure 404 {object} gin.H{"error": "Item not found"}
// @Router /api/shoppingItems/{name} [delete]
func deleteItem(c *gin.Context) {
	name := c.Param("name")

	result, err := db.Exec("DELETE FROM shopping_items WHERE name = $1", name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete item"})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check affected rows"})
		return
	}
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	// Return 204 No Content after successful deletion
    c.Status(http.StatusNoContent) // Status 204
}

// @Summary Get all shopping items
// @Description Retrieve all shopping items
// @Tags Shopping Items API
// @Success 200 {array} ShoppingItem
// @Router /api/shoppingItems [get]
func getAllItems(c *gin.Context) {
	rows, err := db.Query("SELECT name, amount FROM shopping_items")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve items"})
		return
	}
	defer rows.Close()

	var items []ShoppingItem
	for rows.Next() {
		var item ShoppingItem
		if err := rows.Scan(&item.Name, &item.Amount); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse items"})
			return
		}
		items = append(items, item)
	}

	// If no items were found, ensure we return an empty array
	if len(items) == 0 {
		c.JSON(http.StatusOK, []ShoppingItem{}) // Return an empty array instead of null
		return
	}

	// Otherwise, return the list of items
	c.JSON(http.StatusOK, items)
}


// @Summary Add a new shopping item
// @Description Add a new item to the shopping list
// @Tags Shopping Items API
// @Param shoppingItem body ShoppingItem true "New shopping item"
// @Success 201 {object} ShoppingItem
// @Failure 400 {object} gin.H{"error": "Invalid request payload"}
// @Router /api/shoppingItems [post]
func addItem(c *gin.Context) {
	var newItem ShoppingItem
	if err := c.ShouldBindJSON(&newItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Input validation
	if newItem.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Item name cannot be empty"})
		return
	}
	if newItem.Amount <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Amount must be greater than zero"})
		return
	}

	_, err := db.Exec("INSERT INTO shopping_items (name, amount) VALUES ($1, $2)", newItem.Name, newItem.Amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add item"})
		return
	}

	c.JSON(http.StatusCreated, newItem)
}
