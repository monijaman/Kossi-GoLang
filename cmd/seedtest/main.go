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

	fmt.Println("🌱 RUNNING TV PRODUCT REVIEW SEEDER")
	fmt.Println("====================================\n")

	// Run the TV Product Review Seeder
	seeder := seeders.NewTVProductReviewSeeder()
	if err := seeder.Seed(db); err != nil {
		log.Printf("❌ Error running seeder: %v\n", err)
	} else {
		fmt.Println("✅ TV Product Review Seeder completed!\n")
	}

	// Verify reviews were created
	fmt.Println("📊 VERIFICATION")
	fmt.Println("================\n")

	// Count total reviews
	var reviewCount int64
	db.Model(&models.ProductReviewModel{}).Count(&reviewCount)
	fmt.Printf("📝 Total product reviews: %d\n", reviewCount)

	// Count TV reviews specifically
	var tvReviewCount int64
	db.Model(&models.ProductReviewModel{}).
		Joins("JOIN products ON products.id = product_reviews.product_id").
		Where("products.category_id = ?", 124).
		Count(&tvReviewCount)
	fmt.Printf("📺 TV product reviews: %d\n", tvReviewCount)

	// Count translations
	var translationCount int64
	db.Model(&models.ProductReviewTranslationModel{}).Count(&translationCount)
	fmt.Printf("🌐 Review translations: %d\n", translationCount)

	// Show sample TV reviews
	fmt.Println("\n📋 Sample TV Reviews:")
	var reviews []models.ProductReviewModel
	db.Joins("JOIN products ON products.id = product_reviews.product_id").
		Where("products.category_id = ?", 124).
		Limit(5).
		Find(&reviews)

	for i, review := range reviews {
		fmt.Printf("\n%d. Review ID: %d\n", i+1, review.ID)
		if review.Reviews != nil {
			// Show first 100 characters
			previewLen := 100
			if len(*review.Reviews) < previewLen {
				previewLen = len(*review.Reviews)
			}
			fmt.Printf("   Preview: %s...\n", (*review.Reviews)[:previewLen])
		}
		fmt.Printf("   Rating: %s\n", review.Rating)
	}

	fmt.Println("\n✅ Seeding verification completed successfully!")
}
