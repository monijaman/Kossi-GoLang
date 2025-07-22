// Package models contains database-specific implementations.
// This package is part of the infrastructure layer in Clean Architecture.
// These models are tightly coupled to GORM and handle database persistence.
package models

import (
	"kossti/internal/domain/entities"
	"time"
)

// PasswordResetTokenModel represents the database model for password reset tokens (GORM-specific)
type PasswordResetTokenModel struct {
	Email     string     `gorm:"primaryKey;type:varchar(255)"`
	Token     string     `gorm:"type:varchar(255);not null"`
	CreatedAt *time.Time `gorm:""`
}

// ToEntity converts GORM model to domain entity
func (p *PasswordResetTokenModel) ToEntity() *entities.PasswordResetToken {
	var createdAt time.Time
	if p.CreatedAt != nil {
		createdAt = *p.CreatedAt
	}
	
	return &entities.PasswordResetToken{
		Email:     p.Email,
		Token:     p.Token,
		CreatedAt: createdAt,
	}
}

// FromEntity converts domain entity to GORM model
func (p *PasswordResetTokenModel) FromEntity(entity *entities.PasswordResetToken) {
	p.Email = entity.Email
	p.Token = entity.Token
	p.CreatedAt = &entity.CreatedAt
}

// TableName returns the table name for GORM
func (PasswordResetTokenModel) TableName() string {
	return "password_reset_tokens"
}
