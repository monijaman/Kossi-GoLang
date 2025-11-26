package main

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"
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

	fmt.Println("🚗 RUNNING CAR SEEDER")
	fmt.Println("=====================\n")

	// Run the Car Seeder
	seeder := seeders.NewCarSeeder()
	if err := seeder.Seed(db); err != nil {
		log.Printf("❌ Error running car seeder: %v\n", err)
	} else {
		fmt.Println("✅ Car Seeder completed!\n")
	}

	// Run the Car Brand Translation Seeder
	transSeeder := seeders.NewCarBrandTranslationSeeder()
	if err := transSeeder.Seed(db); err != nil {
		log.Printf("❌ Error running translation seeder: %v\n", err)
	} else {
		fmt.Println("✅ Car Brand Translation Seeder completed!\n")
	}

	// Verify reviews were created
	fmt.Println("📊 VERIFICATION")
	fmt.Println("================\n")

	// Count total products
	var productCount int64
	db.Model(&models.ProductModel{}).Count(&productCount)
	fmt.Printf("📝 Total products: %d\n", productCount)

	// Count Car products specifically
	var carCount int64
	db.Model(&models.ProductModel{}).
		Where("category_id = ?", 18).
		Count(&carCount)
	fmt.Printf("🚗 Car products: %d\n", carCount)

	// Count Brands
	var brandCount int64
	db.Model(&models.BrandModel{}).Count(&brandCount)
	fmt.Printf("🏷️ Total Brands: %d\n", brandCount)

	// Count Brand Translations
	var brandTransCount int64
	db.Model(&models.BrandTranslationModel{}).Where("locale = ?", "bn").Count(&brandTransCount)
	fmt.Printf("🗣️ Bangla Brand Translations: %d\n", brandTransCount)

	// Show sample Car products
	fmt.Println("\n📋 Sample Car Products:")
	var cars []models.ProductModel
	db.Where("category_id = ?", 18).
		Limit(5).
		Find(&cars)

	for i, car := range cars {
		fmt.Printf("\n%d. Car ID: %d\n", i+1, car.ID)
		fmt.Printf("   Name: %s\n", car.Name)
		fmt.Printf("   Slug: %s\n", car.Slug)
	}

	fmt.Println("\n✅ Seeding verification completed successfully!")
}
