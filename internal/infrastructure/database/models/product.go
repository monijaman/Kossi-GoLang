// Package models contains database-specific implementations.
// This package is part of the infrastructure layer in Clean Architecture.
// These models are tightly coupled to GORM and handle database persistence.
package models

import (
	"kossti/internal/domain/entities"
	"time"
)

// ProductModel represents the database model for products (GORM-specific)
type ProductModel struct {
	ID          uint       `gorm:"primaryKey;autoIncrement"`
	Name        string     `gorm:"type:varchar(255);not null"`
	Description *string    `gorm:"type:text"`
	Slug        string     `gorm:"type:varchar(255);unique;not null"`
	CategoryID  string     `gorm:"type:varchar(255);not null"`
	BrandID     *int       `gorm:""`
	Model       *string    `gorm:"type:varchar(255)"`
	Price       *float64   `gorm:"type:decimal(10,2)"`
	Status      bool       `gorm:"default:false"`
	Priority    int        `gorm:"default:1"`
	CreatedBy   *string    `gorm:"type:varchar(255)"`
	ViewsCount  int64      `gorm:"default:0"`
	CreatedAt   time.Time  `gorm:"autoCreateTime"`
	UpdatedAt   time.Time  `gorm:"autoUpdateTime"`
	UpdatedBy   *string    `gorm:"type:varchar(255)"`
	DeletedAt   *time.Time `gorm:"index"`
}

// ToEntity converts GORM model to domain entity
func (p *ProductModel) ToEntity() *entities.Product {
	var price float64
	if p.Price != nil {
		price = *p.Price
	}

	var categoryID *uint
	// Convert string CategoryID to uint if needed for business logic
	// This is a simplified conversion - you might need more sophisticated logic

	var brandID *uint
	if p.BrandID != nil {
		brandIDUint := uint(*p.BrandID)
		brandID = &brandIDUint
	}

	return &entities.Product{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		Price:       price,
		CategoryID:  categoryID,
		BrandID:     brandID,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
		DeletedAt:   p.DeletedAt,
	}
}

// FromEntity converts domain entity to GORM model
func (p *ProductModel) FromEntity(entity *entities.Product) {
	p.ID = entity.ID
	p.Name = entity.Name
	p.Description = entity.Description
	p.Price = &entity.Price

	if entity.BrandID != nil {
		brandID := int(*entity.BrandID)
		p.BrandID = &brandID
	}

	p.CreatedAt = entity.CreatedAt
	p.UpdatedAt = entity.UpdatedAt
	p.DeletedAt = entity.DeletedAt
}

// TableName returns the table name for GORM
func (ProductModel) TableName() string {
	return "products"
}
