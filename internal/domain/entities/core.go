// Package entities contains domain entities that are framework-independent.
// These entities represent the business logic and should not depend on any external frameworks.
package entities

import (
	"time"
)

// User represents a user in the domain
type User struct {
	ID              uint
	Name            string
	Email           string
	EmailVerifiedAt *time.Time
	Password        string
	RememberToken   *string
	Type            string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

// PasswordResetToken represents a password reset token
type PasswordResetToken struct {
	Email     string
	Token     string
	CreatedAt time.Time
}
