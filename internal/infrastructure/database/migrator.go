package database

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// Migrator handles database migrations
type Migrator struct {
	db *gorm.DB
}

// NewMigrator creates a new migrator instance
func NewMigrator(db *gorm.DB) *Migrator {
	return &Migrator{db: db}
}

// RunMigrations executes all database migrations
func (m *Migrator) RunMigrations() error {
	log.Println("Running database migrations...")

	err := m.db.AutoMigrate(
		&models.UserModel{},
		&models.ProductModel{},
		&models.ImageModel{},
		&models.ProductTranslationModel{},
		&models.ProductReviewModel{},            // Added for product reviews
		&models.ProductReviewTranslationModel{}, // Added for review translations
		&models.CategoryModel{},                 // Added for categories
		&models.BrandModel{},                    // Added for brands
		&models.BrandCategoryModel{},            // Added for brand-category relations
		&models.SpecificationModel{},            // Added for specifications
		&models.SpecificationKeyModel{},         // Added for specification keys
		&models.FormGeneratorModel{},            // Added for form generator
		&models.FeedbackModel{},                 // Added for feedback
		&models.RefreshTokenModel{},             // Added for refresh tokens
		&models.PasswordResetTokenModel{},       // Added for password resets
		// Add other models here as you create them
	)

	if err != nil {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	log.Println("Database migrations completed successfully")
	return nil
}

// DropTables drops all tables (useful for development)
func (m *Migrator) DropTables() error {
	log.Println("Dropping database tables...")

	err := m.db.Migrator().DropTable(
		&models.UserModel{},
		&models.ProductModel{},
		&models.ImageModel{},
		&models.ProductTranslationModel{},
		&models.ProductReviewModel{},
		&models.ProductReviewTranslationModel{},
		&models.CategoryModel{},
		&models.BrandModel{},
		&models.BrandCategoryModel{},
		&models.SpecificationModel{},
		&models.SpecificationKeyModel{},
		&models.FormGeneratorModel{},
		&models.FeedbackModel{},
		&models.RefreshTokenModel{},
		&models.PasswordResetTokenModel{},
	)

	if err != nil {
		return fmt.Errorf("failed to drop tables: %w", err)
	}

	log.Println("Database tables dropped successfully")
	return nil
}
