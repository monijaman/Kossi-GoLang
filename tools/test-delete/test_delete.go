package main

import (
	"fmt"
	"kossti/internal/infrastructure/database"
	"kossti/internal/infrastructure/database/models"
	"log"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func main() {
	// Load .env
	if err := godotenv.Load(".env"); err != nil {
		log.Println("No .env file found")
	}

	// Connect to database
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Test hard delete
	imageID := uint(2)

	fmt.Printf("=== Testing Hard Delete for Image ID: %d ===\n", imageID)

	// First, verify the image exists
	var before models.ImageModel
	result := db.First(&before, imageID)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			fmt.Printf("Image ID %d not found in database\n", imageID)
			return
		}
		log.Fatalf("Error checking image: %v", result.Error)
	}

	fmt.Printf("BEFORE DELETE: Image ID=%d, Path=%s exists\n", before.ID, before.ImagePath)

	// Perform hard delete
	result = db.Unscoped().Delete(&models.ImageModel{}, imageID)
	if result.Error != nil {
		log.Fatalf("Failed to delete: %v", result.Error)
	}

	fmt.Printf("Delete executed. Rows affected: %d\n", result.RowsAffected)

	// Verify deletion
	var after models.ImageModel
	result = db.Unscoped().First(&after, imageID)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			fmt.Printf("✓ SUCCESS: Image ID %d has been deleted\n", imageID)
		} else {
			log.Fatalf("Error checking after delete: %v", result.Error)
		}
	} else {
		fmt.Printf("✗ FAILED: Image ID %d still exists after delete\n", imageID)
	}
}
