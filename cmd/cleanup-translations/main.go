package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	// Load environment
	godotenv.Load(".env.production")

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL not set")
	}

	fmt.Println("🔧 Cleaning up Bengali translations...")
	fmt.Println("   Removing: (অনুবাদ প্রয়োজন)")

	// Connect to database
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	sqlDB, _ := gormDB.DB()

	// Step 1: Count before
	var countBefore int64
	sqlDB.QueryRow("SELECT COUNT(*) FROM product_translations WHERE locale='bn' AND translated_name LIKE '%(অনুবাদ প্রয়োজন)%'").Scan(&countBefore)
	fmt.Printf("   Found %d translations with phrase\n", countBefore)

	// Step 2: Fetch all Bengali translations
	rows, err := sqlDB.Query("SELECT id, translated_name FROM product_translations WHERE locale='bn'")
	if err != nil {
		log.Fatalf("Query error: %v", err)
	}
	defer rows.Close()

	// Step 3: Clean and update
	regexPattern := regexp.MustCompile(`\s*\(অনুবাদ প্রয়োজন\)`)
	updatedCount := 0

	for rows.Next() {
		var id int64
		var translatedName string

		if err := rows.Scan(&id, &translatedName); err != nil {
			continue
		}

		// Remove the phrase
		cleanedName := regexPattern.ReplaceAllString(translatedName, "")
		cleanedName = strings.TrimSpace(cleanedName)

		// Only update if it changed
		if cleanedName != translatedName {
			sqlDB.Exec("UPDATE product_translations SET translated_name = $1, updated_at = NOW() WHERE id = $2", cleanedName, id)
			updatedCount++

			if updatedCount%100 == 0 {
				fmt.Printf("   ✅ Updated %d translations...\n", updatedCount)
			}
		}
	}

	// Step 4: Verify
	var countAfter int64
	var totalBN int64
	sqlDB.QueryRow("SELECT COUNT(*) FROM product_translations WHERE locale='bn' AND translated_name LIKE '%(অনুবাদ প্রয়োজন)%'").Scan(&countAfter)
	sqlDB.QueryRow("SELECT COUNT(*) FROM product_translations WHERE locale='bn'").Scan(&totalBN)

	fmt.Println("\n" + strings.Repeat("=", 70))
	fmt.Printf("✅ Cleanup Complete!\n")
	fmt.Printf("   Removed phrase from: %d translations\n", updatedCount)
	fmt.Printf("   Remaining with phrase: %d\n", countAfter)
	fmt.Printf("   Total Bengali translations: %d\n", totalBN)
	fmt.Println(strings.Repeat("=", 70))

	// Show samples
	fmt.Println("\n📋 Sample cleaned translations:")
	sampleRows, _ := sqlDB.Query("SELECT product_id, translated_name FROM product_translations WHERE locale='bn' LIMIT 15")
	defer sampleRows.Close()

	count := 0
	for sampleRows.Next() && count < 15 {
		var productID int64
		var translatedName string
		sampleRows.Scan(&productID, &translatedName)
		fmt.Printf("   [%d] %s\n", productID, translatedName)
		count++
	}
}
