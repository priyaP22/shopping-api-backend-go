// db/migrations.go
package db

import (
    "database/sql"
    "fmt"
    "log"
    "github.com/pressly/goose/v3"
)

// RunMigrations applies all pending migrations from the "migrations" folder.
func RunMigrations(db *sql.DB) error {
    // Set the dialect to PostgreSQL
    if err := goose.SetDialect("postgres"); err != nil {
        return fmt.Errorf("failed to set dialect: %v", err)
    }

    // Run the migrations from the "migrations" directory
    if err := goose.Up(db, "migrations"); err != nil {
        return fmt.Errorf("failed to run migrations: %v", err)
    }

    return nil
}
