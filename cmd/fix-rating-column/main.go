package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Load environment
	godotenv.Load()

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		// Use Railway database connection
		dsn = "host=switchyard.proxy.rlwy.net user=postgres password=wqxrKmVOvJeQORpcsGhWOGiRuQdJDLBU dbname=railway port=58014 sslmode=require"
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	fmt.Println("🔄 Converting rating column from numeric to varchar...")

	// Run the conversion SQL
	result := db.Exec(`
		ALTER TABLE product_review_translations 
		ALTER COLUMN rating TYPE varchar(50) USING rating::text;
	`)

	if result.Error != nil {
		log.Fatalf("❌ Failed to convert column: %v", result.Error)
	}

	fmt.Println("✅ Successfully converted rating column to varchar(50)")
	fmt.Println("✅ Rating column can now store Bengali numerals like '৪.१५'")
}
