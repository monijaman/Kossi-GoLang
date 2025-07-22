// Package models contains database-specific implementations.
// This package is part of the infrastructure layer in Clean Architecture.
// These models are tightly coupled to GORM and handle database persistence.
package models

import (
	"kossti/internal/domain/entities"
	"time"
)

// PersonalAccessTokenModel represents the database model for personal access tokens (GORM-specific)
type PersonalAccessTokenModel struct {
	ID            uint       `gorm:"primaryKey;autoIncrement"`
	TokenableType string     `gorm:"type:varchar(255);not null"`
	TokenableID   uint       `gorm:"not null"`
	Name          string     `gorm:"type:varchar(255);not null"`
	Token         string     `gorm:"type:varchar(64);unique;not null"`
	Abilities     *string    `gorm:"type:text"`
	LastUsedAt    *time.Time `gorm:""`
	ExpiresAt     *time.Time `gorm:""`
	CreatedAt     time.Time  `gorm:"autoCreateTime"`
	UpdatedAt     time.Time  `gorm:"autoUpdateTime"`
}

// ToEntity converts GORM model to domain entity
func (p *PersonalAccessTokenModel) ToEntity() *entities.PersonalAccessToken {
	return &entities.PersonalAccessToken{
		ID:            p.ID,
		TokenableType: p.TokenableType,
		TokenableID:   p.TokenableID,
		Name:          p.Name,
		Token:         p.Token,
		Abilities:     p.Abilities,
		LastUsedAt:    p.LastUsedAt,
		ExpiresAt:     p.ExpiresAt,
		CreatedAt:     p.CreatedAt,
		UpdatedAt:     p.UpdatedAt,
	}
}

// FromEntity converts domain entity to GORM model
func (p *PersonalAccessTokenModel) FromEntity(entity *entities.PersonalAccessToken) {
	p.ID = entity.ID
	p.TokenableType = entity.TokenableType
	p.TokenableID = entity.TokenableID
	p.Name = entity.Name
	p.Token = entity.Token
	p.Abilities = entity.Abilities
	p.LastUsedAt = entity.LastUsedAt
	p.ExpiresAt = entity.ExpiresAt
	p.CreatedAt = entity.CreatedAt
	p.UpdatedAt = entity.UpdatedAt
}

// TableName returns the table name for GORM
func (PersonalAccessTokenModel) TableName() string {
	return "personal_access_tokens"
}
