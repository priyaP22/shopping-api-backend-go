package services

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"shopping-api-backend-go/internal/models"
)

// DB holds the global database connection.
var db *sql.DB

// InitDB initializes and returns the database connection
func InitDB() *sql.DB {
	dbHost := os.Getenv("POSTGRES_HOST")
	dbPort := os.Getenv("POSTGRES_PORT")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")

	dbURL := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)
	var err error
	db, err = sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	return db
}

// DB function returns the global database connection.
func DB() *sql.DB {
	if db == nil {
		db = InitDB()
	}
	return db
}

// CreateTableIfNotExists creates the shopping_items table if it doesn't exist
func CreateTableIfNotExists(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS shopping_items (
			name TEXT PRIMARY KEY,
			amount INTEGER NOT NULL
		)
	`)
	return err
}

// GetItemByName retrieves an item by its name from the database
func GetItemByName(db *sql.DB, name string) (models.ShoppingItem, error) {
	var item models.ShoppingItem
	err := db.QueryRow("SELECT name, amount FROM shopping_items WHERE name = $1", name).Scan(&item.Name, &item.Amount)
	return item, err
}

// UpdateItem updates an existing shopping item
func UpdateItem(db *sql.DB, name string, item models.ShoppingItem) error {
	_, err := db.Exec("UPDATE shopping_items SET name = $1, amount = $2 WHERE name = $3", item.Name, item.Amount, name)
	return err
}

// DeleteItem deletes a shopping item from the database
func DeleteItem(db *sql.DB, name string) error {
	_, err := db.Exec("DELETE FROM shopping_items WHERE name = $1", name)
	return err
}

// GetAllItems retrieves all shopping items from the database
func GetAllItems(db *sql.DB) ([]models.ShoppingItem, error) {
	rows, err := db.Query("SELECT name, amount FROM shopping_items")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []models.ShoppingItem
	for rows.Next() {
		var item models.ShoppingItem
		if err := rows.Scan(&item.Name, &item.Amount); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	// If no items were found, ensure we return an empty array
	if len(items) == 0 {
		return []models.ShoppingItem{}, nil // Return an empty array instead of nil
	}

	return items, nil // Otherwise, return the list of items
}



// AddItem adds a new shopping item to the list
func AddItem(db *sql.DB, item models.ShoppingItem) error {
	_, err := db.Exec("INSERT INTO shopping_items (name, amount) VALUES ($1, $2)", item.Name, item.Amount)
	return err
}
