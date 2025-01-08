package handlers

import (
	"database/sql"
	"net/http"
	"shopping-api-backend-go/internal/models"
	"shopping-api-backend-go/internal/services"

	"github.com/gin-gonic/gin"
)

// ErrorResponse represents a standard error response
type ErrorResponse struct {
	Error string `json:"error"`
}

// ResponseMessage represents a standard response message
type ResponseMessage struct {
	Message string `json:"message"`
}

// HealthCheck checks the health of the service
// @Summary Health check
// @Description Check the health status of the API
// @Tags Health API
// @Success 200 {string} string "API is up and running"
// @Router /health [get]
func HealthCheck(c *gin.Context) {
	// Respond with a simple "API is up and running" message
	c.JSON(http.StatusOK, gin.H{
		"message": "API is up and running",
	})
}

// GetItemByName retrieves an item by its name
// @Summary Get a shopping item by name
// @Description Retrieve a specific shopping item by its name
// @Tags Shopping Items API
// @Param name path string true "Item name"
// @Success 200 {object} models.ShoppingItem
// @Failure 404 {object} ErrorResponse
// @Router /api/shoppingItems/{name} [get]
func GetItemByName(c *gin.Context) {
	name := c.Param("name")

	// Fetch item from the service layer
	item, err := services.GetItemByName(services.DB(), name)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, ErrorResponse{"Item not found"})
		} else {
			c.JSON(http.StatusInternalServerError, ErrorResponse{"Failed to retrieve item"})
		}
		return
	}

	// Return the shopping item
	c.JSON(http.StatusOK, item)
}

// UpdateItem updates a shopping item by its name
// @Summary Update a shopping item by name
// @Description Update a specific shopping item by its name
// @Tags Shopping Items API
// @Param name path string true "Item name"
// @Param shoppingItem body models.ShoppingItem true "Updated shopping item"
// @Success 200 {object} models.ShoppingItem
// @Failure 404 {object} ErrorResponse
// @Router /api/shoppingItems/{name} [put]
func UpdateItem(c *gin.Context) {
	name := c.Param("name")
	var updatedItem models.ShoppingItem
	if err := c.ShouldBindJSON(&updatedItem); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{"Invalid request payload"})
		return
	}

	// Input validation
	if updatedItem.Amount <= 0 {
		c.JSON(http.StatusBadRequest, ErrorResponse{"Amount must be greater than zero"})
		return
	}

	// Call the service layer to update the item
	err := services.UpdateItem(services.DB(), name, updatedItem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{"Failed to update item"})
		return
	}

	// Return the updated item
	c.JSON(http.StatusOK, updatedItem)
}

// DeleteItem deletes a shopping item by its name
// @Summary Delete a shopping item by name
// @Description Delete a specific shopping item by its name
// @Tags Shopping Items API
// @Param name path string true "Item name"
// @Success 200 {object} ResponseMessage
// @Failure 404 {object} ErrorResponse
// @Router /api/shoppingItems/{name} [delete]
func DeleteItem(c *gin.Context) {
	name := c.Param("name")

	// Call the service layer to delete the item
	err := services.DeleteItem(services.DB(), name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{"Failed to delete item"})
		return
	}

	// Return 204 No Content after successful deletion
	c.Status(http.StatusNoContent)
}

// GetAllItems retrieves all shopping items
// @Summary Get all shopping items
// @Description Retrieve all shopping items
// @Tags Shopping Items API
// @Success 200 {array} models.ShoppingItem
// @Router /api/shoppingItems [get]
func GetAllItems(c *gin.Context) {
	items, err := services.GetAllItems(services.DB())
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{"Failed to retrieve items"})
		return
	}

	// Return all items
	c.JSON(http.StatusOK, items)
}

// AddItem adds a new shopping item to the list
// @Summary Add a new shopping item
// @Description Add a new item to the shopping list
// @Tags Shopping Items API
// @Param shoppingItem body models.ShoppingItem true "New shopping item"
// @Success 201 {object} models.ShoppingItem
// @Failure 400 {object} ErrorResponse
// @Router /api/shoppingItems [post]
func AddItem(c *gin.Context) {
	var newItem models.ShoppingItem
	if err := c.ShouldBindJSON(&newItem); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{"Invalid request payload"})
		return
	}

	// Input validation
	if newItem.Name == "" {
		c.JSON(http.StatusBadRequest, ErrorResponse{"Item name cannot be empty"})
		return
	}
	if newItem.Amount <= 0 {
		c.JSON(http.StatusBadRequest, ErrorResponse{"Amount must be greater than zero"})
		return
	}

	// Call the service layer to add the item
	err := services.AddItem(services.DB(), newItem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{"Failed to add item"})
		return
	}

	// Return the newly added item
	c.JSON(http.StatusCreated, newItem)
}
