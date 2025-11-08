// Package models contains database-specific implementations.
// This package is part of the infrastructure layer in Clean Architecture.
// These models are tightly coupled to GORM and handle database persistence.
package models

import (
	"kossti/internal/domain/entities"
	"time"
)

// CategoryModel represents the database model for categories (GORM-specific)
type CategoryModel struct {
	ID        uint       `gorm:"primaryKey;autoIncrement"`
	Name      string     `gorm:"type:varchar(255);not null"`
	Slug      string     `gorm:"type:varchar(255);unique;not null"`
	Status    int        `gorm:"type:integer;default:1"`
	CreatedAt time.Time  `gorm:"autoCreateTime"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime"`
	DeletedAt *time.Time `gorm:"index"`
}

// ToEntity converts GORM model to domain entity
func (c *CategoryModel) ToEntity() *entities.Category {
	return &entities.Category{
		ID:        c.ID,
		Name:      c.Name,
		Slug:      c.Slug,
		Status:    c.Status,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
		DeletedAt: c.DeletedAt,
	}
}

// FromEntity converts domain entity to GORM model
func (c *CategoryModel) FromEntity(entity *entities.Category) {
	c.ID = entity.ID
	c.Name = entity.Name
	c.Slug = entity.Slug
	c.Status = entity.Status
	c.CreatedAt = entity.CreatedAt
	c.UpdatedAt = entity.UpdatedAt
	c.DeletedAt = entity.DeletedAt
}

// TableName returns the table name for GORM
func (CategoryModel) TableName() string {
	return "categories"
}
