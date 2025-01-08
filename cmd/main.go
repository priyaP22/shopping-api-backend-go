package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"shopping-api-backend-go/docs"
	"shopping-api-backend-go/internal/services"
	"shopping-api-backend-go/web"

	"github.com/gin-contrib/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // PostgreSQL driver
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var db *sql.DB

// @title Shopping API
// @version 1.0
// @description A simple API to manage shopping items with PostgreSQL
// @BasePath /
func main() {
	// Load environment variables
	if err := godotenv.Load("../.env"); err != nil { // Load from the root directory
		log.Fatalf("Error loading .env file")
	}

	// Initialize DB connection
	db = services.InitDB()

	// Create table if it doesn't exist
	if err := services.CreateTableIfNotExists(db); err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}

	// Initialize Gin router
	r := web.InitializeRouter(db)

	// CORS configuration
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://*.app.github.dev", "http://localhost:5000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "HEAD", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
	}))

	// Dynamically set Swagger host
	codespaceName := os.Getenv("CODESPACE_NAME")
	githubDomain := os.Getenv("GITHUB_COSPACE_DOMAIN")
	var swaggerHost string

	if codespaceName != "" && githubDomain != "" {
		// Construct the base URL for GitHub Codespaces
		swaggerHost = fmt.Sprintf("%s-8080.%s", codespaceName, githubDomain)
	} else {
		// Default to localhost for local development
		swaggerHost = "https://localhost:8080"
	}

	// Set host in Swagger documentation
	docs.SwaggerInfo.Host = swaggerHost

	// Swagger Endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Start server
	r.Run(":8080") // Default is localhost:8080
}
