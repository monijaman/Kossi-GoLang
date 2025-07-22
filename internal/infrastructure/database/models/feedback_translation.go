// Package models contains database-specific implementations.
// This package is part of the infrastructure layer in Clean Architecture.
// These models are tightly coupled to GORM and handle database persistence.
package models

import (
	"kossti/internal/domain/entities"
	"time"
)

// FeedbackTranslationModel represents the database model for feedback translations (GORM-specific)
type FeedbackTranslationModel struct {
	ID                uint      `gorm:"primaryKey;autoIncrement"`
	FeedbackID        uint      `gorm:"not null"`
	Locale            string    `gorm:"type:varchar(255);not null"`
	TranslatedContent string    `gorm:"type:text;not null"`
	CreatedAt         time.Time `gorm:"autoCreateTime"`
	UpdatedAt         time.Time `gorm:"autoUpdateTime"`
}

// ToEntity converts GORM model to domain entity
func (ft *FeedbackTranslationModel) ToEntity() *entities.FeedbackTranslation {
	return &entities.FeedbackTranslation{
		ID:                ft.ID,
		FeedbackID:        ft.FeedbackID,
		Locale:            ft.Locale,
		TranslatedContent: ft.TranslatedContent,
		CreatedAt:         ft.CreatedAt,
		UpdatedAt:         ft.UpdatedAt,
	}
}

// FromEntity converts domain entity to GORM model
func (ft *FeedbackTranslationModel) FromEntity(entity *entities.FeedbackTranslation) {
	ft.ID = entity.ID
	ft.FeedbackID = entity.FeedbackID
	ft.Locale = entity.Locale
	ft.TranslatedContent = entity.TranslatedContent
	ft.CreatedAt = entity.CreatedAt
	ft.UpdatedAt = entity.UpdatedAt
}

// TableName returns the table name for GORM
func (FeedbackTranslationModel) TableName() string {
	return "feedback_translations"
}
