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
	UserID    *uint      `gorm:""`
	Content   string     `gorm:"type:text;not null"`
	Status    bool       `gorm:"default:true"`
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
	f.CreatedAt = entity.CreatedAt
	f.UpdatedAt = entity.UpdatedAt
	f.DeletedAt = entity.DeletedAt
}

// TableName returns the table name for GORM
func (FeedbackModel) TableName() string {
	return "feedback"
}
