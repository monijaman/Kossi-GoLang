// Package models contains database-specific implementations.
// This package is part of the infrastructure layer in Clean Architecture.
// These models are tightly coupled to GORM and handle database persistence.
package models

import (
	"kossti/internal/domain/entities"
)

// BrandCategoryModel represents the database model for brand_category pivot table (GORM-specific)
type BrandCategoryModel struct {
	ID         uint `gorm:"primaryKey;autoIncrement"`
	BrandID    uint `gorm:"not null"`
	CategoryID uint `gorm:"not null"`
}

// ToEntity converts GORM model to domain entity
func (bc *BrandCategoryModel) ToEntity() *entities.BrandCategory {
	return &entities.BrandCategory{
		ID:         bc.ID,
		BrandID:    bc.BrandID,
		CategoryID: bc.CategoryID,
	}
}

// FromEntity converts domain entity to GORM model
func (bc *BrandCategoryModel) FromEntity(entity *entities.BrandCategory) {
	bc.ID = entity.ID
	bc.BrandID = entity.BrandID
	bc.CategoryID = entity.CategoryID
}

// TableName returns the table name for GORM
func (BrandCategoryModel) TableName() string {
	return "brand_category"
}
