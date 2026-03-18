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
	ID            uint           `gorm:"primaryKey;autoIncrement"`
	Price         *float64       `gorm:"-"` // Legacy price field (ignored for persistence)
	Name          string         `gorm:"type:varchar(255);not null"`
	Description   *string        `gorm:"type:text"`
	Slug          string         `gorm:"type:varchar(255);unique;not null"`
	CategoryID    *uint          `gorm:""`
	BrandID       *uint          `gorm:""`
	Category      *CategoryModel `gorm:"foreignKey:CategoryID"`
	Brand         *BrandModel    `gorm:"foreignKey:BrandID"`
	Model         *string        `gorm:"type:varchar(255)"`
	StartPrice    *float64       `gorm:"type:decimal(10,2)"`
	EndPrice      *float64       `gorm:"type:decimal(10,2)"`
	Status        int            `gorm:"type:integer;default:1"`
	Priority      int            `gorm:"default:1"`
	CreatedBy     *string        `gorm:"type:varchar(255)"`
	ViewsCount    int64          `gorm:"default:0"`
	CreatedAt     time.Time      `gorm:"autoCreateTime"`
	UpdatedAt     time.Time      `gorm:"autoUpdateTime"`
	UpdatedBy     *string        `gorm:"type:varchar(255)"`
	DeletedAt     *time.Time     `gorm:"index"`
	AverageRating *float64       `gorm:"-"` // Computed field, not persisted
}

// ToEntity converts GORM model to domain entity
func (p *ProductModel) ToEntity() *entities.Product {
	var price float64
	var startPrice *float64
	var endPrice *float64
	if p.StartPrice != nil {
		startPrice = p.StartPrice
	}
	if p.EndPrice != nil {
		endPrice = p.EndPrice
	}

	var category *entities.Category
	if p.Category != nil {
		category = p.Category.ToEntity()
	}
	var brand *entities.Brand
	if p.Brand != nil {
		brand = p.Brand.ToEntity()
	}

	return &entities.Product{
		ID:            p.ID,
		Name:          p.Name,
		Description:   p.Description,
		Slug:          p.Slug,
		Price:         price,
		StartPrice:    startPrice,
		EndPrice:      endPrice,
		CategoryID:    p.CategoryID,
		BrandID:       p.BrandID,
		Category:      category,
		Brand:         brand,
		ViewsCount:    p.ViewsCount,
		Status:        p.Status > 0, // Convert int to bool: 1+ = true, 0 = false
		Priority:      p.Priority,
		CreatedAt:     p.CreatedAt,
		UpdatedAt:     p.UpdatedAt,
		DeletedAt:     p.DeletedAt,
		AverageRating: p.AverageRating,
	}
}

// FromEntity converts domain entity to GORM model
func (p *ProductModel) FromEntity(entity *entities.Product) {
	p.ID = entity.ID
	p.Name = entity.Name
	p.Description = entity.Description
	p.Slug = entity.Slug
	// Map start/end price if provided
	p.StartPrice = entity.StartPrice
	p.EndPrice = entity.EndPrice
	p.CategoryID = entity.CategoryID
	p.BrandID = entity.BrandID
	p.ViewsCount = entity.ViewsCount
	// Convert bool to int: true = 1, false = 0
	if entity.Status {
		p.Status = 1
	} else {
		p.Status = 0
	}
	p.Priority = entity.Priority
	p.CreatedAt = entity.CreatedAt
	p.UpdatedAt = entity.UpdatedAt
	p.DeletedAt = entity.DeletedAt
}

// TableName returns the table name for GORM
func (ProductModel) TableName() string {
	return "products"
}
