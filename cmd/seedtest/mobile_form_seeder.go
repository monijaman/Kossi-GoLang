package main

import (
	"fmt"
	"kossti/internal/infrastructure/database/seeders"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using default values")
	}

	// Get database URL from environment
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "host=switchyard.proxy.rlwy.net user=postgres password=wqxrKmVOvJeQORpcsGhWOGiRuQdJDLBU dbname=railway port=58014 sslmode=require TimeZone=UTC"
	}

	// Connect to database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	fmt.Println("\n📱 UPDATING MOBILE FORM GENERATOR (ID 5)")
	fmt.Println("=========================================\n")

	// Run the Mobile Form Generator Seeder
	if err := seeders.SeedMobileFormGenerator(db); err != nil {
		log.Printf("❌ Error running mobile form generator seeder: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("\n✅ Mobile Form Generator seeder completed successfully!")
}
