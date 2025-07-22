// Simple migration CLI tool
// Run with: go run cmd/migrate/main.go
package main

import (
	"flag"
	"fmt"
	"kossti/internal/infrastructure/database"
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

	// Get database URL from environment or use default
	defaultDSN := os.Getenv("DATABASE_URL")
	if defaultDSN == "" {
		defaultDSN = "host=localhost user=root password=root dbname=kossti port=5428 sslmode=disable TimeZone=UTC"
	}

	// Define flags
	var (
		fresh    = flag.Bool("fresh", false, "Drop all tables and recreate (DANGEROUS)")
		migrate  = flag.Bool("migrate", false, "Run migrations (safe)")
		drop     = flag.Bool("drop", false, "Drop all tables (DANGEROUS)")
		fk       = flag.Bool("fk", false, "Add foreign keys only")
		indexes  = flag.Bool("indexes", false, "Create indexes only")
		dsn      = flag.String("dsn", defaultDSN, "Database DSN")
	)
	flag.Parse()

	// Connect to database
	db, err := gorm.Open(postgres.Open(*dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Create migration manager
	manager := database.NewMigrationManager(db)

	// Execute based on flags
	switch {
	case *fresh:
		fmt.Println("🚨 FRESH SETUP: This will DROP ALL TABLES!")
		if !confirmDangerous() {
			return
		}
		if err := manager.DropAllTables(); err != nil {
			log.Fatal("Failed to drop tables:", err)
		}
		if err := manager.Setup(); err != nil {
			log.Fatal("Failed to setup:", err)
		}
		fmt.Println("✅ Fresh setup completed!")

	case *migrate:
		fmt.Println("🛡️ Running safe migration...")
		if err := manager.MigrateAll(); err != nil {
			log.Fatal("Migration failed:", err)
		}
		if err := manager.AddForeignKeys(); err != nil {
			log.Fatal("Adding foreign keys failed:", err)
		}
		if err := manager.CreateIndexes(); err != nil {
			log.Fatal("Creating indexes failed:", err)
		}
		fmt.Println("✅ Migration completed!")

	case *drop:
		fmt.Println("🚨 DROP ALL TABLES!")
		if !confirmDangerous() {
			return
		}
		if err := manager.DropAllTables(); err != nil {
			log.Fatal("Failed to drop tables:", err)
		}
		fmt.Println("✅ All tables dropped!")

	case *fk:
		fmt.Println("Adding foreign keys...")
		if err := manager.AddForeignKeys(); err != nil {
			log.Fatal("Adding foreign keys failed:", err)
		}
		fmt.Println("✅ Foreign keys added!")

	case *indexes:
		fmt.Println("Creating indexes...")
		if err := manager.CreateIndexes(); err != nil {
			log.Fatal("Creating indexes failed:", err)
		}
		fmt.Println("✅ Indexes created!")

	default:
		printUsage()
	}
}

func confirmDangerous() bool {
	fmt.Print("This is DANGEROUS! Type 'YES' to continue: ")
	var confirm string
	fmt.Scanln(&confirm)
	return confirm == "YES"
}

func printUsage() {
	fmt.Println("Migration Tool Usage:")
	fmt.Println("  go run cmd/migrate/main.go [flags]")
	fmt.Println("")
	fmt.Println("Environment Variables:")
	fmt.Println("  DATABASE_URL    Database connection string (reads from .env file)")
	fmt.Println("")
	fmt.Println("Flags:")
	fmt.Println("  -migrate    Run safe migration (recommended for production)")
	fmt.Println("  -fresh      Drop all tables and recreate (DANGEROUS - dev only)")
	fmt.Println("  -drop       Drop all tables (DANGEROUS)")
	fmt.Println("  -fk         Add foreign keys only")
	fmt.Println("  -indexes    Create indexes only")
	fmt.Println("  -dsn        Database connection string (overrides DATABASE_URL)")
	fmt.Println("")
	fmt.Println("Examples:")
	fmt.Println("  go run cmd/migrate/main.go -migrate")
	fmt.Println("  go run cmd/migrate/main.go -fresh")
	fmt.Println("  go run cmd/migrate/main.go -dsn 'host=prod-db user=prod password=secret dbname=production'")
	fmt.Println("")
	fmt.Println("Note: The tool automatically reads DATABASE_URL from .env file if present")
}
