package models

import (
	"kossti/internal/domain/entities"
)

type RefreshTokenModel struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    uint   `gorm:"index;not null"`
	Token     string `gorm:"not null;unique"`
	ExpiresAt int64  `gorm:"not null"`
}

func (m *RefreshTokenModel) ToEntity() *entities.RefreshToken {
	return &entities.RefreshToken{
		ID:        m.ID,
		UserID:    m.UserID,
		Token:     m.Token,
		ExpiresAt: m.ExpiresAt,
	}
}

func (m *RefreshTokenModel) FromEntity(e *entities.RefreshToken) {
	m.ID = e.ID
	m.UserID = e.UserID
	m.Token = e.Token
	m.ExpiresAt = e.ExpiresAt
}

func (RefreshTokenModel) TableName() string {
	return "refresh_tokens"
}
