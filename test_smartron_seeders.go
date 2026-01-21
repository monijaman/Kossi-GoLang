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
		dsn = "host=localhost user=root password=root dbname=kossti port=5428 sslmode=disable TimeZone=UTC"
	}

	// Connect to database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	fmt.Println("🧪 TESTING SMARTRON T.BOOK SEEDERS")
	fmt.Println("=================================")

	// Run the Smartron t.book flex Seeder
	fmt.Println("Testing Smartron t.book flex Seeder...")
	seeder1 := seeders.NewSmartronTBookFlexSeeder()
	if err := seeder1.Seed(db); err != nil {
		log.Printf("❌ Error running Smartron t.book flex seeder: %v\n", err)
	} else {
		fmt.Println("✅ Smartron t.book flex Seeder completed!")
	}

	// Run the Smartron t.book Seeder
	fmt.Println("Testing Smartron t.book Seeder...")
	seeder2 := seeders.NewSmartronTBookSeeder()
	if err := seeder2.Seed(db); err != nil {
		log.Printf("❌ Error running Smartron t.book seeder: %v\n", err)
	} else {
		fmt.Println("✅ Smartron t.book Seeder completed!")
	}

	fmt.Println("🎉 All Smartron seeders tested!")
}
