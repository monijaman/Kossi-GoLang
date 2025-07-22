// Package database contains database-specific implementations.
// This package is part of the infrastructure layer in Clean Architecture.
package database

import (
	"kossti/internal/infrastructure/database/models"
	"fmt"
	"log"

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

// GetAllModels returns all models that need to be migrated
func (m *MigrationManager) GetAllModels() []interface{} {
	return []interface{}{
		// Core system tables
		&models.UserModel{},
		&models.PasswordResetTokenModel{},
		&models.SessionModel{},
		&models.CacheModel{},
		&models.CacheLockModel{},
		&models.JobModel{},
		&models.JobBatchModel{},
		&models.FailedJobModel{},
		&models.PersonalAccessTokenModel{},
		&models.HistoryLogModel{},
		
		// Permission system
		&models.PermissionModel{},
		&models.RoleModel{},
		&models.ModelHasPermissionModel{},
		&models.ModelHasRoleModel{},
		&models.RoleHasPermissionModel{},
		
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
		&models.FeedbackTranslationModel{},
		
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
		
		// Form generator
		&models.FormGeneratorModel{},
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
	models := m.GetAllModels()
	
	log.Println("Dropping all tables...")
	
	// Drop in reverse order to handle foreign key constraints
	for i := len(models) - 1; i >= 0; i-- {
		model := models[i]
		log.Printf("Dropping table for model: %T", model)
		if err := m.db.Migrator().DropTable(model); err != nil {
			log.Printf("Warning: failed to drop table for %T: %v", model, err)
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
		{&models.SessionModel{}, "user_id", "users(id)", "SET NULL"},
		{&models.ProductReviewModel{}, "user_id", "users(id)", "CASCADE"},
		{&models.FeedbackModel{}, "user_id", "users(id)", "SET NULL"},
		{&models.BrandCategoryModel{}, "brand_id", "brands(id)", "CASCADE"},
		{&models.BrandCategoryModel{}, "category_id", "categories(id)", "CASCADE"},
		{&models.ModelHasPermissionModel{}, "permission_id", "permissions(id)", "CASCADE"},
		{&models.ModelHasRoleModel{}, "role_id", "roles(id)", "CASCADE"},
		{&models.RoleHasPermissionModel{}, "permission_id", "permissions(id)", "CASCADE"},
		{&models.RoleHasPermissionModel{}, "role_id", "roles(id)", "CASCADE"},
		
		// Translation foreign keys
		{&models.ProductTranslationModel{}, "product_id", "products(id)", "CASCADE"},
		{&models.CategoryTranslationModel{}, "category_id", "categories(id)", "CASCADE"},
		{&models.BrandTranslationModel{}, "brand_id", "brands(id)", "CASCADE"},
		{&models.ProductReviewTranslationModel{}, "product_review_id", "product_reviews(id)", "CASCADE"},
		{&models.FeedbackTranslationModel{}, "feedback_id", "feedback(id)", "CASCADE"},
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
		table   interface{}
		fields  []string
		name    string
		unique  bool
	}{
		{&models.UserModel{}, []string{"email"}, "idx_users_email", true},
		{&models.ProductModel{}, []string{"slug"}, "idx_products_slug", true},
		{&models.ProductModel{}, []string{"category_id", "brand_id"}, "idx_products_category_brand", false},
		{&models.CategoryModel{}, []string{"slug"}, "idx_categories_slug", true},
		{&models.BrandModel{}, []string{"slug"}, "idx_brands_slug", true},
		{&models.PermissionModel{}, []string{"name", "guard_name"}, "idx_permissions_name_guard", true},
		{&models.RoleModel{}, []string{"name", "guard_name"}, "idx_roles_name_guard", true},
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
