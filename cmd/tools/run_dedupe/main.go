package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"kossti/internal/infrastructure/database"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables with production support (same as migrate tool)
	goEnv := os.Getenv("GO_ENV")

	// Load .env.production if GO_ENV=production or if the file exists
	if goEnv == "production" {
		if err := godotenv.Load(".env.production"); err != nil {
			log.Println("Warning: .env.production not found, trying .env")
			godotenv.Load() // Fallback to .env
		} else {
			log.Println("✅ Loaded .env.production")
		}
	} else if _, err := os.Stat(".env.production"); err == nil {
		// .env.production exists, load it and set GO_ENV
		if err := godotenv.Load(".env.production"); err == nil {
			os.Setenv("GO_ENV", "production")
			log.Println("✅ Loaded .env.production (file exists)")
		}
	} else {
		// Load regular .env
		err := godotenv.Load()
		if err != nil {
			log.Println("No .env file found, using default values")
		}
	}

	fmt.Printf("🔗 Connecting to database...\n")

	// Connect using the same method as other CLI tools
	cfg := database.NewDatabaseConfig()
	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	fmt.Println("✅ Connected to database")

	// Read SQL file
	sqlFile := "../../dedupe_specification_keys.sql"
	sqlContent, err := ioutil.ReadFile(sqlFile)
	if err != nil {
		log.Fatal("Failed to read SQL file:", err)
	}

	fmt.Printf("📄 Read SQL file: %s (%d bytes)\n", sqlFile, len(sqlContent))

	// Execute SQL
	fmt.Println("⚡ Executing deduplication script...")
	result := db.Exec(string(sqlContent))
	if result.Error != nil {
		log.Fatal("Failed to execute SQL:", result.Error)
	}

	fmt.Printf("✅ Script executed successfully! Rows affected: %d\n", result.RowsAffected)

	// Show audit results
	fmt.Println("\n📊 Deduplication Results:")
	var auditResults []struct {
		Summary                string `gorm:"column:summary"`
		SpecificationsAffected string `gorm:"column:specifications_affected"`
	}
	db.Raw(`
		SELECT
			'Removed ' || COUNT(*) || ' duplicate specification keys' as summary,
			SUM(affected_specifications_count) || ' specifications updated' as specifications_affected
		FROM specification_keys_dedupe_audit
	`).Scan(&auditResults)

	for _, result := range auditResults {
		fmt.Printf("   %s\n", result.Summary)
		fmt.Printf("   %s\n", result.SpecificationsAffected)
	}

	fmt.Println("\n📋 Detailed Audit Log:")
	var auditDetails []struct {
		DuplicateKeyID              int    `gorm:"column:duplicate_key_id"`
		DuplicateKeyName            string `gorm:"column:duplicate_key_name"`
		KeptKeyID                   int    `gorm:"column:kept_key_id"`
		KeptKeyName                 string `gorm:"column:kept_key_name"`
		AffectedSpecificationsCount int    `gorm:"column:affected_specifications_count"`
	}
	db.Table("specification_keys_dedupe_audit").Order("duplicate_key_id").Find(&auditDetails)

	for _, detail := range auditDetails {
		fmt.Printf("   Removed ID %d '%s' → Kept ID %d '%s' (%d specs updated)\n",
			detail.DuplicateKeyID, detail.DuplicateKeyName,
			detail.KeptKeyID, detail.KeptKeyName,
			detail.AffectedSpecificationsCount)
	}
}
