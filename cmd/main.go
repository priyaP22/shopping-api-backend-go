package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"shopping-api-backend-go/docs"
	"shopping-api-backend-go/internal/services"
	"shopping-api-backend-go/web"
	"time"

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
	log.Println("Application starting...")

	// Get the ENV variable
	env := os.Getenv("ENV")
	if env == "" {
		env = "development" // Default to "development" if ENV is not set
	}

	// Build the .env file path based on ENV
	envFile := "../.env." + env

	// Load the environment variables from the .env file
	if err := godotenv.Load(envFile); err != nil {
		log.Fatalf("Error loading .env file: %s", err)
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
		swaggerHost = "localhost:8080"
	}

	// Set host in Swagger documentation
	docs.SwaggerInfo.Host = swaggerHost

	// Swagger Endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Create an http.Server
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	// Go routine to start the server
	go func() {
		log.Println("Server starting on :8080")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown signal received, initiating graceful shutdown...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown: %v\n", err)
	}

	// Clean up other resources like DB connections
	if err := db.Close(); err != nil {
		log.Printf("Error closing database connection: %v", err)
	}

	log.Println("Server exited gracefully")
}
