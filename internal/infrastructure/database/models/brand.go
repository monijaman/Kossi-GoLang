// Package models contains database-specific implementations.
// This package is part of the infrastructure layer in Clean Architecture.
// These models are tightly coupled to GORM and handle database persistence.
package models

import (
	"kossti/internal/domain/entities"
	"time"
)

// BrandModel represents the database model for brands (GORM-specific)
type BrandModel struct {
	ID        uint       `gorm:"primaryKey;autoIncrement"`
	Name      string     `gorm:"type:varchar(255);not null"`
	Slug      string     `gorm:"type:varchar(255);unique;not null"`
	CreatedAt time.Time  `gorm:"autoCreateTime"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime"`
	DeletedAt *time.Time `gorm:"index"`
}

// ToEntity converts GORM model to domain entity
func (b *BrandModel) ToEntity() *entities.Brand {
	return &entities.Brand{
		ID:        b.ID,
		Name:      b.Name,
		Slug:      b.Slug,
		CreatedAt: b.CreatedAt,
		UpdatedAt: b.UpdatedAt,
		DeletedAt: b.DeletedAt,
	}
}

// FromEntity converts domain entity to GORM model
func (b *BrandModel) FromEntity(entity *entities.Brand) {
	b.ID = entity.ID
	b.Name = entity.Name
	b.Slug = entity.Slug
	b.CreatedAt = entity.CreatedAt
	b.UpdatedAt = entity.UpdatedAt
	b.DeletedAt = entity.DeletedAt
}

// TableName returns the table name for GORM
func (BrandModel) TableName() string {
	return "brands"
}
