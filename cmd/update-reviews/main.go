package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"gocrit/internal/infrastructure/database"
	"gorm.io/gorm"
)

func main() {
	// Initialize database connection
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Read the SQL file
	sqlFile := filepath.Join(".", "SQL_UPDATE_PRODUCT_REVIEWS_HTML.sql")
	sqlContent, err := os.ReadFile(sqlFile)
	if err != nil {
		log.Fatalf("Failed to read SQL file: %v", err)
	}

	// Execute the SQL
	result := db.Exec(string(sqlContent))
	if result.Error != nil {
		log.Fatalf("Failed to execute SQL: %v", result.Error)
	}

	fmt.Println("✅ Database updated successfully!")
	fmt.Printf("Rows affected: %d\n", result.RowsAffected)
}
