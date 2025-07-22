// Package models contains database-specific implementations.
// This package is part of the infrastructure layer in Clean Architecture.
// These models are tightly coupled to GORM and handle database persistence.
package models

import (
	"kossti/internal/domain/entities"
	"time"
)

// FormGeneratorModel represents the database model for form generator (GORM-specific)
type FormGeneratorModel struct {
	ID              uint      `gorm:"primaryKey;autoIncrement"`
	CategoryID      uint      `gorm:"not null"`
	SpecificationID string    `gorm:"type:json;not null"`
	Status          string    `gorm:"type:varchar(255);default:'active'"`
	CreatedAt       time.Time `gorm:"autoCreateTime"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime"`
}

// ToEntity converts GORM model to domain entity
func (fg *FormGeneratorModel) ToEntity() *entities.FormGenerator {
	return &entities.FormGenerator{
		ID:              fg.ID,
		CategoryID:      fg.CategoryID,
		SpecificationID: fg.SpecificationID,
		Status:          fg.Status,
		CreatedAt:       fg.CreatedAt,
		UpdatedAt:       fg.UpdatedAt,
	}
}

// FromEntity converts domain entity to GORM model
func (fg *FormGeneratorModel) FromEntity(entity *entities.FormGenerator) {
	fg.ID = entity.ID
	fg.CategoryID = entity.CategoryID
	fg.SpecificationID = entity.SpecificationID
	fg.Status = entity.Status
	fg.CreatedAt = entity.CreatedAt
	fg.UpdatedAt = entity.UpdatedAt
}

// TableName returns the table name for GORM
func (FormGeneratorModel) TableName() string {
	return "form_generator"
}
