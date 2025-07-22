// Package models contains database-specific implementations.
// This package is part of the infrastructure layer in Clean Architecture.
// These models are tightly coupled to GORM and handle database persistence.
package models

import (
	"kossti/internal/domain/entities"
)

// SessionModel represents the database model for user sessions (GORM-specific)
type SessionModel struct {
	ID           string  `gorm:"primaryKey;type:varchar(255)"`
	UserID       *uint   `gorm:"index"`
	IPAddress    *string `gorm:"type:varchar(45)"`
	UserAgent    *string `gorm:"type:text"`
	Payload      string  `gorm:"type:text;not null"`
	LastActivity int     `gorm:"index;not null"`
}

// ToEntity converts GORM model to domain entity
func (s *SessionModel) ToEntity() *entities.Session {
	return &entities.Session{
		ID:           s.ID,
		UserID:       s.UserID,
		IPAddress:    s.IPAddress,
		UserAgent:    s.UserAgent,
		Payload:      s.Payload,
		LastActivity: s.LastActivity,
	}
}

// FromEntity converts domain entity to GORM model
func (s *SessionModel) FromEntity(entity *entities.Session) {
	s.ID = entity.ID
	s.UserID = entity.UserID
	s.IPAddress = entity.IPAddress
	s.UserAgent = entity.UserAgent
	s.Payload = entity.Payload
	s.LastActivity = entity.LastActivity
}

// TableName returns the table name for GORM
func (SessionModel) TableName() string {
	return "sessions"
}
