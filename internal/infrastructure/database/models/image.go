// Package models contains database-specific implementations.
// This package is part of the infrastructure layer in Clean Architecture.
// These models are tightly coupled to GORM and handle database persistence.
package models

import (
	"kossti/internal/domain/entities"
	"time"
)

// ImageModel represents the database model for images (GORM-specific)
type ImageModel struct {
	ID            uint       `gorm:"primaryKey;autoIncrement"`
	ImageableType string     `gorm:"type:varchar(255);not null"`
	ImageableID   uint       `gorm:"not null"`
	ImagePath     string     `gorm:"type:varchar(255);not null"`
	Status        int        `gorm:"type:integer;default:1"`
	DefaultPhoto  int        `gorm:"column:default_photo;type:integer;default:0"`
	CreatedAt     time.Time  `gorm:"autoCreateTime"`
	UpdatedAt     time.Time  `gorm:"autoUpdateTime"`
	DeletedAt     *time.Time `gorm:"index"`
}

// ToEntity converts GORM model to domain entity
func (i *ImageModel) ToEntity() *entities.Image {
	return &entities.Image{
		ID:            i.ID,
		ImageableType: i.ImageableType,
		ImageableID:   i.ImageableID,
		ImagePath:     i.ImagePath,
		Status:        i.Status,
		DefaultPhoto:  i.DefaultPhoto,
		CreatedAt:     i.CreatedAt,
		UpdatedAt:     i.UpdatedAt,
		DeletedAt:     i.DeletedAt,
	}
}

// FromEntity converts domain entity to GORM model
func (i *ImageModel) FromEntity(entity *entities.Image) {
	i.ID = entity.ID
	i.ImageableType = entity.ImageableType
	i.ImageableID = entity.ImageableID
	i.ImagePath = entity.ImagePath
	i.Status = entity.Status
	i.DefaultPhoto = entity.DefaultPhoto
	i.CreatedAt = entity.CreatedAt
	i.UpdatedAt = entity.UpdatedAt
	i.DeletedAt = entity.DeletedAt
}

// TableName returns the table name for GORM
func (ImageModel) TableName() string {
	return "images"
}
