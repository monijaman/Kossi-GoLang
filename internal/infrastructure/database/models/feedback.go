// Package models contains database-specific implementations.
// This package is part of the infrastructure layer in Clean Architecture.
// These models are tightly coupled to GORM and handle database persistence.
package models

import (
	"kossti/internal/domain/entities"
	"time"
)

// FeedbackModel represents the database model for feedback (GORM-specific)
type FeedbackModel struct {
	ID        uint       `gorm:"primaryKey;autoIncrement"`
	UserID    uint       `gorm:"not null"`
	Content   string     `gorm:"type:text;not null"`
	Status    bool       `gorm:"default:true"`
	CreatedBy *uint      `gorm:"nullable"`
	UpdatedBy *uint      `gorm:"nullable"`
	CreatedAt time.Time  `gorm:"autoCreateTime"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime"`
	DeletedAt *time.Time `gorm:"index"`
}

// ToEntity converts GORM model to domain entity
func (f *FeedbackModel) ToEntity() *entities.Feedback {
	return &entities.Feedback{
		ID:        f.ID,
		UserID:    f.UserID,
		Content:   f.Content,
		Status:    f.Status,
		CreatedBy: f.CreatedBy,
		UpdatedBy: f.UpdatedBy,
		CreatedAt: f.CreatedAt,
		UpdatedAt: f.UpdatedAt,
		DeletedAt: f.DeletedAt,
	}
}

// FromEntity converts domain entity to GORM model
func (f *FeedbackModel) FromEntity(entity *entities.Feedback) {
	f.ID = entity.ID
	f.UserID = entity.UserID
	f.Content = entity.Content
	f.Status = entity.Status
	f.CreatedBy = entity.CreatedBy
	f.UpdatedBy = entity.UpdatedBy
	f.CreatedAt = entity.CreatedAt
	f.UpdatedAt = entity.UpdatedAt
	f.DeletedAt = entity.DeletedAt
}

// TableName returns the table name for GORM
func (FeedbackModel) TableName() string {
	return "feedback"
}

// FeedbackTranslationModel represents the database model for feedback translations
type FeedbackTranslationModel struct {
	ID                uint      `gorm:"primaryKey;autoIncrement"`
	FeedbackID        uint      `gorm:"not null"`
	Locale            string    `gorm:"type:varchar(10);not null"`
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
