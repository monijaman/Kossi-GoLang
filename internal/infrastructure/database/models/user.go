// Package models contains database-specific implementations.
// This package is part of the infrastructure layer in Clean Architecture.
// These models are tightly coupled to GORM and handle database persistence.
package models

import (
	"kossti/internal/domain/entities"
	"time"
)

// UserModel represents the database model with GORM-specific tags (GORM-specific)
type UserModel struct {
	ID              uint       `gorm:"primaryKey;autoIncrement"`
	Name            string     `gorm:"type:varchar(255);not null"`
	Email           string     `gorm:"type:varchar(255);unique;not null"`
	EmailVerifiedAt *time.Time `gorm:""`
	Password        string     `gorm:"type:varchar(255);not null"`
	RememberToken   *string    `gorm:"type:varchar(100)"`
	CreatedAt       time.Time  `gorm:"autoCreateTime"`
	UpdatedAt       time.Time  `gorm:"autoUpdateTime"`
}

// ToEntity converts GORM model to domain entity
func (u *UserModel) ToEntity() *entities.User {
	return &entities.User{
		ID:              u.ID,
		Name:            u.Name,
		Email:           u.Email,
		EmailVerifiedAt: u.EmailVerifiedAt,
		Password:        u.Password,
		RememberToken:   u.RememberToken,
		CreatedAt:       u.CreatedAt,
		UpdatedAt:       u.UpdatedAt,
	}
}

// FromEntity converts domain entity to GORM model
func (u *UserModel) FromEntity(entity *entities.User) {
	u.ID = entity.ID
	u.Name = entity.Name
	u.Email = entity.Email
	u.EmailVerifiedAt = entity.EmailVerifiedAt
	u.Password = entity.Password
	u.RememberToken = entity.RememberToken
	u.CreatedAt = entity.CreatedAt
	u.UpdatedAt = entity.UpdatedAt
}

// TableName returns the table name for GORM
func (UserModel) TableName() string {
	return "users"
}
