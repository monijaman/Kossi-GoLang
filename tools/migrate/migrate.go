package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"kossti/internal/infrastructure/database"
)

func main() {
	// Define command line flags
	var (
		migrate = flag.Bool("migrate", false, "Run database migrations")
		drop    = flag.Bool("drop", false, "Drop all tables")
		reset   = flag.Bool("reset", false, "Drop all tables and then migrate")
		help    = flag.Bool("help", false, "Show help")
	)
	flag.Parse()

	if *help {
		showHelp()
		return
	}

	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: Could not load .env file: %v", err)
	}

	// Get database URL
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL environment variable is required")
	}

	// Connect to database
	fmt.Println("Connecting to database...")
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Test connection
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get database instance: %v", err)
	}

	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("Database ping failed: %v", err)
	}
	fmt.Println("Database connection successful!")

	// Initialize migration manager
	migrationManager := database.NewMigrationManager(db)

	// Execute commands based on flags
	switch {
	case *reset:
		fmt.Println("Resetting database (drop + migrate)...")
		if err := migrationManager.DropAllTables(); err != nil {
			log.Fatalf("Failed to drop tables: %v", err)
		}
		if err := migrationManager.Setup(); err != nil {
			log.Fatalf("Failed to run migrations: %v", err)
		}
		fmt.Println("Database reset completed successfully!")

	case *drop:
		fmt.Println("Dropping all tables...")
		if err := migrationManager.DropAllTables(); err != nil {
			log.Fatalf("Failed to drop tables: %v", err)
		}
		fmt.Println("All tables dropped successfully!")

	case *migrate:
		fmt.Println("Running database migrations...")
		if err := migrationManager.Setup(); err != nil {
			log.Fatalf("Failed to run migrations: %v", err)
		}
		fmt.Println("Database migrations completed successfully!")

	default:
		fmt.Println("No action specified. Use -help to see available options.")
		showHelp()
	}
}

func showHelp() {
	fmt.Println("Database Migration Tool")
	fmt.Println("=======================")
	fmt.Println("")
	fmt.Println("Usage:")
	fmt.Println("  go run migrate.go [options]")
	fmt.Println("")
	fmt.Println("Options:")
	fmt.Println("  -migrate    Run database migrations (create all tables)")
	fmt.Println("  -drop       Drop all tables")
	fmt.Println("  -reset      Drop all tables and then run migrations")
	fmt.Println("  -help       Show this help message")
	fmt.Println("")
	fmt.Println("Examples:")
	fmt.Println("  go run migrate.go -migrate")
	fmt.Println("  go run migrate.go -reset")
	fmt.Println("  go run migrate.go -drop")
	fmt.Println("")
	fmt.Println("Environment Variables:")
	fmt.Println("  DATABASE_URL    PostgreSQL connection string (required)")
}
