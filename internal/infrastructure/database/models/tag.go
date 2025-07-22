// Package models contains database-specific implementations.
// This package is part of the infrastructure layer in Clean Architecture.
// These models are tightly coupled to GORM and handle database persistence.
package models

import (
	"kossti/internal/domain/entities"
	"time"
)

// TagModel represents the database model for tags (GORM-specific)
type TagModel struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	Name      string    `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

// ToEntity converts GORM model to domain entity
func (t *TagModel) ToEntity() *entities.Tag {
	return &entities.Tag{
		ID:        t.ID,
		Name:      t.Name,
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}
}

// FromEntity converts domain entity to GORM model
func (t *TagModel) FromEntity(entity *entities.Tag) {
	t.ID = entity.ID
	t.Name = entity.Name
	t.CreatedAt = entity.CreatedAt
	t.UpdatedAt = entity.UpdatedAt
}

// TableName returns the table name for GORM
func (TagModel) TableName() string {
	return "tags"
}
