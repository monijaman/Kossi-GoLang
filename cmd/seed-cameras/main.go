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
		// Try Railway production database first
		railwayHost := os.Getenv("RAILWAY_DB_HOST")
		railwayPort := os.Getenv("RAILWAY_DB_PORT")
		railwayUser := os.Getenv("RAILWAY_DB_USER")
		railwayPassword := os.Getenv("RAILWAY_DB_PASSWORD")
		railwayDatabase := os.Getenv("RAILWAY_DB_NAME")

		if railwayHost != "" {
			dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require TimeZone=UTC",
				railwayHost, railwayUser, railwayPassword, railwayDatabase, railwayPort)
			fmt.Println("📡 Using Railway database connection")
		} else {
			// Build DSN from individual environment variables (local)
			host := os.Getenv("DB_HOST")
			if host == "" {
				host = "localhost"
			}
			port := os.Getenv("DB_PORT")
			if port == "" {
				port = "5428"
			}
			user := os.Getenv("DB_USER")
			if user == "" {
				user = "root"
			}
			password := os.Getenv("DB_PASSWORD")
			if password == "" {
				password = "root"
			}
			dbname := os.Getenv("DB_NAME")
			if dbname == "" {
				dbname = "kossti"
			}
			dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
				host, user, password, dbname, port)
			fmt.Println("🏠 Using local database connection")
		}
	}

	// Connect to database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	fmt.Println("📷 RUNNING CAMERA PRODUCT SEEDER")
	fmt.Println("=================================\n")

	// Run the Camera Product Seeder
	seeder := seeders.NewCameraProductSeeder()
	if err := seeder.Seed(db); err != nil {
		log.Printf("❌ Error running camera seeder: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("\n✅ Camera Product Seeder completed!\n")

	// Verify products were created
	fmt.Println("📊 VERIFICATION")
	fmt.Println("================\n")

	// Count total products
	var productCount int64
	db.Model(&models.ProductModel{}).Count(&productCount)
	fmt.Printf("📝 Total products in database: %d\n", productCount)

	// Find Cameras category
	var category models.CategoryModel
	db.Where("slug = ?", "cameras").First(&category)
	fmt.Printf("📷 Cameras category ID: %d\n", category.ID)

	// Count Camera products specifically
	var cameraCount int64
	db.Model(&models.ProductModel{}).
		Where("category_id = ?", category.ID).
		Count(&cameraCount)
	fmt.Printf("📷 Camera products created: %d\n", cameraCount)

	// Count Brands
	var brandCount int64
	db.Model(&models.BrandModel{}).Count(&brandCount)
	fmt.Printf("🏷️  Total Brands: %d\n", brandCount)

	// Count Camera brands specifically
	var cameraBrandCount int64
	db.Table("brand_categories").
		Where("category_id = ?", category.ID).
		Count(&cameraBrandCount)
	fmt.Printf("📷 Camera Brands: %d\n", cameraBrandCount)

	// Show sample Camera products by brand
	fmt.Println("\n📋 Sample Camera Products by Brand:")

	type BrandProducts struct {
		BrandName    string
		ProductCount int64
	}

	var brandProducts []BrandProducts
	db.Raw(`
		SELECT b.name as brand_name, COUNT(p.id) as product_count
		FROM products p
		JOIN brands b ON p.brand_id = b.id
		WHERE p.category_id = ?
		GROUP BY b.name
		ORDER BY COUNT(p.id) DESC
		LIMIT 10
	`, category.ID).Scan(&brandProducts)

	for i, bp := range brandProducts {
		fmt.Printf("%2d. %-20s : %3d products\n", i+1, bp.BrandName, bp.ProductCount)
	}

	// Show sample products
	fmt.Println("\n📸 Sample Camera Products:")
	var cameras []struct {
		Name      string
		Slug      string
		BrandName string
	}
	db.Raw(`
		SELECT p.name, p.slug, b.name as brand_name
		FROM products p
		LEFT JOIN brands b ON p.brand_id = b.id
		WHERE p.category_id = ?
		LIMIT 10
	`, category.ID).Scan(&cameras)

	for i, cam := range cameras {
		fmt.Printf("\n%2d. %s\n", i+1, cam.Name)
		fmt.Printf("    Slug: %s\n", cam.Slug)
		fmt.Printf("    Brand: %s\n", cam.BrandName)
	}

	fmt.Println("\n✅ Camera seeding verification completed successfully!")
	fmt.Println("\n🎉 All camera products have been created in the database!")
}
