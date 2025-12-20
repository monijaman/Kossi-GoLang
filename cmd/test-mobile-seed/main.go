package main

import (
	"fmt"
	"log"
	"os"

	"kossti/internal/infrastructure/database/seeders"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("🧪 Testing Mobile Products & Specifications Seeding...")

	// Load env
	godotenv.Load()

	// Initialize database
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "host=switchyard.proxy.rlwy.net port=58014 user=postgres password=wqxrKmVOvJeQORpcsGhWOGiRuQdJDLBU dbname=railway sslmode=require"
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Database connection failed: %v", err)
	}

	fmt.Println("\n📍 Step 1: Seeding Specification Key Translations (Bangla)...")
	if err := seeders.SeedMobileSpecificationKeyTranslations(db); err != nil {
		log.Printf("❌ Error: %v", err)
		os.Exit(1)
	}

	fmt.Println("\n📍 Step 2: Seeding Mobile Specifications...")
	if err := seeders.SeedMobileSpecifications(db); err != nil {
		log.Printf("❌ Error: %v", err)
		os.Exit(1)
	}

	fmt.Println("\n✅ All tests completed successfully!")

	// Verify
	fmt.Println("\n📊 Verification:")
	verifySeedingResults(db)
}

func verifySeedingResults(db *gorm.DB) {
	var specKeyCount int64
	db.Table("specification_keys").Count(&specKeyCount)
	fmt.Printf("   - Specification Keys: %d\n", specKeyCount)

	var specCount int64
	db.Table("specifications").Count(&specCount)
	fmt.Printf("   - Specifications: %d\n", specCount)

	var specTranslationCount int64
	db.Table("specification_key_translations").Count(&specTranslationCount)
	fmt.Printf("   - Specification Key Translations: %d\n", specTranslationCount)

	var specValueTranslationCount int64
	db.Table("specification_translations").Count(&specValueTranslationCount)
	fmt.Printf("   - Specification Value Translations: %d\n", specValueTranslationCount)
}
