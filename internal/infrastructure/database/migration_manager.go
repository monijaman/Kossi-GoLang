// Package database contains database-specific implementations.
// This package is part of the infrastructure layer in Clean Architecture.
package database

import (
	"fmt"
	"kossti/internal/infrastructure/database/migrations"
	"kossti/internal/infrastructure/database/models"
	"log"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// MigrationManager handles database migrations
type MigrationManager struct {
	db *gorm.DB
}

// NewMigrationManager creates a new migration manager
func NewMigrationManager(db *gorm.DB) *MigrationManager {
	return &MigrationManager{db: db}
}

// CreateDatabaseIfNotExists creates the database if it doesn't exist
func (m *MigrationManager) CreateDatabaseIfNotExists(dsn, dbName string) error {
	log.Printf("Checking if database '%s' exists...", dbName)

	// Parse DSN to get connection without database name
	// Connect to 'postgres' default database to create target database
	baseDSN := strings.Replace(dsn, "/"+dbName+"?", "/postgres?", 1)
	// Also handle key=value format
	if !strings.Contains(baseDSN, "://") {
		baseDSN = strings.Replace(dsn, "dbname="+dbName, "dbname=postgres", 1)
	}

	// Open connection to default postgres database
	baseDB, err := gorm.Open(postgres.Open(baseDSN), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to postgres database: %w", err)
	}

	// Get underlying SQL DB for raw queries
	sqlDB, err := baseDB.DB()
	if err != nil {
		return fmt.Errorf("failed to get underlying database connection: %w", err)
	}
	defer sqlDB.Close()

	// Check if database exists
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname = $1)"
	err = baseDB.Raw(query, dbName).Scan(&exists).Error
	if err != nil {
		return fmt.Errorf("failed to check if database exists: %w", err)
	}

	if exists {
		log.Printf("✅ Database '%s' already exists", dbName)
		return nil
	}

	// Create database if it doesn't exist
	log.Printf("🔨 Creating database '%s'...", dbName)
	createQuery := fmt.Sprintf("CREATE DATABASE %s", dbName)
	err = baseDB.Exec(createQuery).Error
	if err != nil {
		return fmt.Errorf("failed to create database '%s': %w", dbName, err)
	}

	log.Printf("✅ Database '%s' created successfully!", dbName)
	return nil
} // GetAllModels returns all models that need to be migrated
func (m *MigrationManager) GetAllModels() []interface{} {
	return []interface{}{
		// Core system tables
		&models.UserModel{},
		&models.PasswordResetTokenModel{},

		// Product system
		&models.ProductModel{},
		&models.ProductReviewModel{},
		&models.CategoryModel{},
		&models.BrandModel{},
		&models.BrandCategoryModel{},
		&models.ProductableModel{},

		// Translation tables
		&models.ProductTranslationModel{},
		&models.CategoryTranslationModel{},
		&models.BrandTranslationModel{},
		&models.ProductReviewTranslationModel{},

		// Comments system
		&models.CommentModel{},
		&models.CommentTranslationModel{},

		// Specifications system
		&models.SpecificationKeyModel{},
		&models.SpecificationModel{},
		&models.SpecificationTranslationModel{},
		&models.SpecificationKeyTranslationModel{},

		// Media and feedback
		&models.ImageModel{},
		&models.TagModel{},
		&models.FeedbackModel{},
		&models.FeedbackTranslationModel{},

		// Form generator
		&models.FormGeneratorModel{},
		// Refresh tokens
		&models.RefreshTokenModel{},
	}
}

// MigrateAll runs auto-migration for all models
func (m *MigrationManager) MigrateAll() error {
	models := m.GetAllModels()

	log.Println("Starting database migration...")

	for i, model := range models {
		log.Printf("Migrating model %d/%d: %T", i+1, len(models), model)
		if err := m.db.AutoMigrate(model); err != nil {
			return fmt.Errorf("failed to migrate %T: %w", model, err)
		}
	}

	log.Println("All migrations completed successfully!")
	return nil
}

// DropAllTables drops all tables (useful for testing)
func (m *MigrationManager) DropAllTables() error {
	log.Println("Dropping all tables...")

	// First, disable foreign key checks temporarily to avoid constraint issues
	log.Println("Disabling foreign key checks temporarily...")

	// Get all table names from the database
	var tableNames []string
	err := m.db.Raw(`
		SELECT tablename 
		FROM pg_tables 
		WHERE schemaname = 'public'
	`).Scan(&tableNames).Error
	if err != nil {
		return fmt.Errorf("failed to get table names: %w", err)
	}

	log.Printf("Found %d tables to drop", len(tableNames))

	// Drop all tables in a single transaction with CASCADE
	for _, tableName := range tableNames {
		log.Printf("Dropping table: %s", tableName)
		dropSQL := fmt.Sprintf("DROP TABLE IF EXISTS %s CASCADE", tableName)
		if err := m.db.Exec(dropSQL).Error; err != nil {
			log.Printf("Warning: failed to drop table %s: %v", tableName, err)
			// Continue with other tables
		}
	}

	log.Println("All tables dropped!")
	return nil
}

// AddForeignKeys adds foreign key constraints after migration
func (m *MigrationManager) AddForeignKeys() error {
	log.Println("Adding foreign key constraints...")

	// Add foreign key constraints
	foreignKeys := []struct {
		table      interface{}
		field      string
		references string
		onDelete   string
	}{
		{&models.ProductReviewModel{}, "user_id", "users(id)", "CASCADE"},
		{&models.BrandCategoryModel{}, "brand_id", "brands(id)", "CASCADE"},
		{&models.BrandCategoryModel{}, "category_id", "categories(id)", "CASCADE"},

		// Translation foreign keys
		{&models.ProductTranslationModel{}, "product_id", "products(id)", "CASCADE"},
		{&models.CategoryTranslationModel{}, "category_id", "categories(id)", "CASCADE"},
		{&models.BrandTranslationModel{}, "brand_id", "brands(id)", "CASCADE"},
		{&models.ProductReviewTranslationModel{}, "product_review_id", "product_reviews(id)", "CASCADE"},
		{&models.CommentTranslationModel{}, "comment_id", "comments(id)", "CASCADE"},

		// Specification foreign keys
		{&models.SpecificationModel{}, "product_id", "products(id)", "CASCADE"},
		{&models.SpecificationModel{}, "specification_key_id", "specification_keys(id)", "CASCADE"},
		{&models.SpecificationTranslationModel{}, "specification_id", "specifications(id)", "CASCADE"},
		{&models.SpecificationKeyTranslationModel{}, "specification_key_id", "specification_keys(id)", "CASCADE"},

		// Form generator foreign key
		{&models.FormGeneratorModel{}, "category_id", "categories(id)", "CASCADE"},
	}

	for _, fk := range foreignKeys {
		if m.db.Migrator().HasConstraint(fk.table, fk.field) {
			log.Printf("Foreign key constraint already exists for %T.%s", fk.table, fk.field)
			continue
		}

		log.Printf("Adding foreign key: %T.%s -> %s", fk.table, fk.field, fk.references)
		// Note: GORM AutoMigrate should handle most foreign keys automatically
		// This is here for any custom constraints needed
	}

	log.Println("Foreign key constraints added!")
	return nil
}

// CreateIndexes creates additional indexes for performance
func (m *MigrationManager) CreateIndexes() error {
	log.Println("Creating additional indexes...")

	indexes := []struct {
		table  interface{}
		fields []string
		name   string
		unique bool
	}{
		{&models.UserModel{}, []string{"email"}, "idx_users_email", true},
		{&models.ProductModel{}, []string{"slug"}, "idx_products_slug", true},
		{&models.ProductModel{}, []string{"category_id", "brand_id"}, "idx_products_category_brand", false},
		{&models.CategoryModel{}, []string{"slug"}, "idx_categories_slug", true},
		{&models.BrandModel{}, []string{"slug"}, "idx_brands_slug", true},
	}

	for _, idx := range indexes {
		if m.db.Migrator().HasIndex(idx.table, idx.name) {
			log.Printf("Index already exists: %s", idx.name)
			continue
		}

		log.Printf("Creating index: %s on %T(%v)", idx.name, idx.table, idx.fields)
		// Note: GORM AutoMigrate should handle most indexes from struct tags
		// This is here for any additional custom indexes needed
	}

	log.Println("Additional indexes created!")
	return nil
}

// Setup runs the complete database setup
func (m *MigrationManager) Setup() error {
	// First create all tables
	if err := m.MigrateAll(); err != nil {
		return err
	}

	if err := m.AddForeignKeys(); err != nil {
		return err
	}

	if err := m.CreateIndexes(); err != nil {
		return err
	}

	return nil
}

// UpdateBanglalinkReviews runs the Banglalink review update migration
func (m *MigrationManager) UpdateBanglalinkReviews() error {
	if m.db == nil {
		return fmt.Errorf("database connection not initialized")
	}

	// Call the migration function that updates Banglalink reviews with HTML content
	return migrations.UpdateBanglalinkReviewsWithComparison(m.db)
}

// SetupWithDatabaseCreation runs complete database setup including database creation
func (m *MigrationManager) SetupWithDatabaseCreation(dsn, dbName string) error {
	// First create database if it doesn't exist
	if err := m.CreateDatabaseIfNotExists(dsn, dbName); err != nil {
		return err
	}

	// Then run normal setup
	return m.Setup()
}
