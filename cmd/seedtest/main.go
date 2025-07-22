package main

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"
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

	fmt.Println("🌱 SEEDING VERIFICATION DEMO")
	fmt.Println("================================")

	// Test categories
	var categoryCount int64
	db.Model(&models.CategoryModel{}).Count(&categoryCount)
	fmt.Printf("📂 Categories seeded: %d\n", categoryCount)

	// Test brands
	var brandCount int64
	db.Model(&models.BrandModel{}).Count(&brandCount)
	fmt.Printf("🏷️  Brands seeded: %d\n", brandCount)

	// Test brand-category relationships
	var bcCount int64
	db.Table("brand_category").Count(&bcCount)
	fmt.Printf("🔗 Brand-Category relationships: %d\n", bcCount)

	// Show some sample data
	fmt.Println("\n📋 Sample Categories:")
	var categories []models.CategoryModel
	db.Limit(5).Find(&categories)
	for _, cat := range categories {
		fmt.Printf("   - %s (slug: %s)\n", cat.Name, cat.Slug)
	}

	fmt.Println("\n🏢 Sample Brands:")
	var brands []models.BrandModel
	db.Limit(5).Find(&brands)
	for _, brand := range brands {
		fmt.Printf("   - %s (slug: %s)\n", brand.Name, brand.Slug)
	}

	// Show brand-category relationships for one brand
	fmt.Println("\n🎯 Apple's Categories:")
	var apple models.BrandModel
	if err := db.Where("name = ?", "Apple").First(&apple).Error; err == nil {
		var appleCategories []models.CategoryModel
		db.Table("categories").
			Joins("JOIN brand_category ON categories.id = brand_category.category_id").
			Where("brand_category.brand_id = ?", apple.ID).
			Find(&appleCategories)
		
		for _, cat := range appleCategories {
			fmt.Printf("   - %s\n", cat.Name)
		}
	}

	fmt.Println("\n✅ Seeding verification completed successfully!")
}
