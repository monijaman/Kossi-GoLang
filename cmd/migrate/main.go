// Simple migration CLI tool
// Run with: go run cmd/migrate/main.go
package main

import (
	"flag"
	"fmt"
	"kossti/internal/infrastructure/database"
	"kossti/internal/infrastructure/database/seeders"
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
		fresh     = flag.Bool("fresh", false, "Drop all tables and recreate (DANGEROUS)")
		migrate   = flag.Bool("migrate", false, "Run migrations (safe)")
		createDB  = flag.Bool("create-db", false, "Create database if not exists, then migrate")
		drop      = flag.Bool("drop", false, "Drop all tables (DANGEROUS)")
		fk        = flag.Bool("fk", false, "Add foreign keys only")
		indexes   = flag.Bool("indexes", false, "Create indexes only")
		seed      = flag.Bool("seed", false, "Run all seeders")
		freshSeed = flag.Bool("fresh-seed", false, "Drop, migrate, and seed (DANGEROUS)")
		dsn       = flag.String("dsn", defaultDSN, "Database DSN")
	)
	flag.Parse()

	// Extract database name from DSN for create-db operation
	dbName := extractDatabaseName(*dsn)

	// For create-db flag, first try to create database, then connect
	if *createDB {
		fmt.Println("🔨 Creating database if not exists, then running migration...")

		// Create a temporary migration manager with a dummy connection to create database
		tempManager := database.NewMigrationManager(nil)
		if err := tempManager.CreateDatabaseIfNotExists(*dsn, dbName); err != nil {
			log.Fatal("Failed to create database:", err)
		}
	}

	// Connect to database
	logLevel := logger.Warn
	if *seed || *freshSeed {
		logLevel = logger.Silent // Quiet for seeding operations
	}

	db, err := gorm.Open(postgres.Open(*dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Create migration manager
	manager := database.NewMigrationManager(db)

	// Execute based on flags
	switch {
	case *createDB:
		// Database creation already handled above, now just run normal setup
		if err := manager.Setup(); err != nil {
			log.Fatal("Failed to run migration:", err)
		}
		fmt.Println("✅ Database created and migration completed!")

	case *freshSeed:
		fmt.Println("🚨 FRESH SETUP WITH SEEDING: This will DROP ALL TABLES!")
		if !confirmDangerous() {
			return
		}
		if err := manager.DropAllTables(); err != nil {
			log.Fatal("Failed to drop tables:", err)
		}
		if err := manager.Setup(); err != nil {
			log.Fatal("Failed to setup:", err)
		}

		// Run seeders
		fmt.Println("🌱 Running seeders...")
		seederManager := seeders.SetupAllSeeders(db)
		if err := seederManager.RunAll(); err != nil {
			log.Fatal("Failed to run seeders:", err)
		}
		fmt.Println("✅ Fresh setup with seeding completed!")

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

	case *seed:
		fmt.Println("🌱 Running seeders...")
		seederManager := seeders.SetupAllSeeders(db)
		if err := seederManager.RunAll(); err != nil {
			log.Fatal("Failed to run seeders:", err)
		}
		fmt.Println("✅ Seeding completed!")

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

func extractDatabaseName(dsn string) string {
	// Extract database name from PostgreSQL DSN
	// Example: "host=localhost user=root password=root dbname=kossti port=5428"
	re := regexp.MustCompile(`dbname=(\w+)`)
	matches := re.FindStringSubmatch(dsn)
	if len(matches) > 1 {
		return matches[1]
	}

	// If not found in key=value format, try URL format
	// Example: "postgres://user:pass@host:port/dbname?options"
	if strings.Contains(dsn, "://") {
		parts := strings.Split(dsn, "/")
		if len(parts) > 3 {
			dbPart := parts[3]
			// Remove query parameters if present
			if questionIndex := strings.Index(dbPart, "?"); questionIndex != -1 {
				dbPart = dbPart[:questionIndex]
			}
			return dbPart
		}
	}

	// Default fallback
	return "kossti"
}

func confirmDangerous() bool {
	fmt.Print("This is DANGEROUS! Type 'YES' to continue: ")
	var confirm string
	fmt.Scanln(&confirm)
	return confirm == "YES" || confirm == "yes" || confirm == "y"
}

func printUsage() {
	fmt.Println("Migration Tool Usage:")
	fmt.Println("  go run cmd/migrate/main.go [flags]")
	fmt.Println("")
	fmt.Println("Environment Variables:")
	fmt.Println("  DATABASE_URL    Database connection string (reads from .env file)")
	fmt.Println("")
	fmt.Println("Flags:")
	fmt.Println("  -migrate      Run safe migration (recommended for production)")
	fmt.Println("  -create-db    Create database if not exists, then migrate (safe)")
	fmt.Println("  -fresh        Drop all tables and recreate (DANGEROUS - dev only)")
	fmt.Println("  -fresh-seed   Drop, migrate, and seed database (DANGEROUS - dev only)")
	fmt.Println("  -seed         Run all database seeders")
	fmt.Println("  -drop         Drop all tables (DANGEROUS)")
	fmt.Println("  -fk           Add foreign keys only")
	fmt.Println("  -indexes      Create indexes only")
	fmt.Println("  -dsn          Database connection string (overrides DATABASE_URL)")
	fmt.Println("")
	fmt.Println("Examples:")
	fmt.Println("  go run cmd/migrate/main.go -migrate")
	fmt.Println("  go run cmd/migrate/main.go -create-db")
	fmt.Println("  go run cmd/migrate/main.go -fresh")
	fmt.Println("  go run cmd/migrate/main.go -fresh-seed")
	fmt.Println("  go run cmd/migrate/main.go -seed")
	fmt.Println("  go run cmd/migrate/main.go -dsn 'host=prod-db user=prod password=secret dbname=production'")
	fmt.Println("")
	fmt.Println("Note: The tool automatically reads DATABASE_URL from .env file if present")
}
