// Migration Example - How to use the MigrationManager
// Run this with: go run examples/migration_example.go
package main

import (
	"kossti/internal/infrastructure/database"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func runMigrationExample() {
	// Connect to database
	db, err := connectDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Create migration manager
	migrationManager := database.NewMigrationManager(db)

	// Choose your migration approach:
	fmt.Println("Choose migration option:")
	fmt.Println("1. Fresh setup (drop all + migrate)")
	fmt.Println("2. Migrate only (safe for production)")
	fmt.Println("3. Drop all tables")
	fmt.Println("4. Add foreign keys only")
	fmt.Println("5. Create indexes only")

	var choice int
	fmt.Print("Enter choice (1-5): ")
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		// DEVELOPMENT: Fresh setup (drops everything and recreates)
		fmt.Println("🚨 WARNING: This will DROP ALL TABLES!")
		fmt.Print("Type 'YES' to continue: ")
		var confirm string
		fmt.Scanln(&confirm)
		if confirm == "YES" {
			if err := freshSetup(migrationManager); err != nil {
				log.Fatal("Fresh setup failed:", err)
			}
		}

	case 2:
		// PRODUCTION: Safe migration (only creates/updates)
		if err := safeMigration(migrationManager); err != nil {
			log.Fatal("Migration failed:", err)
		}

	case 3:
		// TESTING: Drop all tables
		fmt.Println("🚨 WARNING: This will DROP ALL TABLES!")
		fmt.Print("Type 'YES' to continue: ")
		var confirm string
		fmt.Scanln(&confirm)
		if confirm == "YES" {
			if err := migrationManager.DropAllTables(); err != nil {
				log.Fatal("Drop tables failed:", err)
			}
		}

	case 4:
		// Add foreign keys only
		if err := migrationManager.AddForeignKeys(); err != nil {
			log.Fatal("Adding foreign keys failed:", err)
		}

	case 5:
		// Create indexes only
		if err := migrationManager.CreateIndexes(); err != nil {
			log.Fatal("Creating indexes failed:", err)
		}

	default:
		fmt.Println("Invalid choice")
		return
	}

	fmt.Println("✅ Migration completed successfully!")
}

// connectDB establishes database connection
func connectDB() (*gorm.DB, error) {
	// Database configuration
	dsn := "host=localhost user=postgres password=postgres dbname=gocrit_db port=5432 sslmode=disable TimeZone=UTC"
	
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return db, nil
}

// freshSetup performs a complete fresh setup (DEVELOPMENT ONLY)
func freshSetup(manager *database.MigrationManager) error {
	log.Println("🔥 Starting fresh setup...")
	
	// 1. Drop all existing tables
	if err := manager.DropAllTables(); err != nil {
		return fmt.Errorf("failed to drop tables: %w", err)
	}
	
	// 2. Run complete setup
	if err := manager.Setup(); err != nil {
		return fmt.Errorf("failed to setup database: %w", err)
	}
	
	log.Println("✅ Fresh setup completed!")
	return nil
}

// safeMigration performs safe migration (PRODUCTION SAFE)
func safeMigration(manager *database.MigrationManager) error {
	log.Println("🛡️ Starting safe migration...")
	
	// Only migrate (won't drop anything)
	if err := manager.MigrateAll(); err != nil {
		return fmt.Errorf("failed to migrate: %w", err)
	}
	
	// Add any missing foreign keys
	if err := manager.AddForeignKeys(); err != nil {
		return fmt.Errorf("failed to add foreign keys: %w", err)
	}
	
	// Create any missing indexes
	if err := manager.CreateIndexes(); err != nil {
		return fmt.Errorf("failed to create indexes: %w", err)
	}
	
	log.Println("✅ Safe migration completed!")
	return nil
}
