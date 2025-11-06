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
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

// PasswordResetToken represents a password reset token
type PasswordResetToken struct {
	Email     string
	Token     string
	CreatedAt time.Time
}

// Session represents a user session
type Session struct {
	ID           string
	UserID       *uint
	IPAddress    *string
	UserAgent    *string
	Payload      string
	LastActivity int
}

// Cache represents a cache entry
type Cache struct {
	Key        string
	Value      string
	Expiration int
}

// CacheLock represents a cache lock
type CacheLock struct {
	Key        string
	Owner      string
	Expiration int
}

// PersonalAccessToken represents a personal access token
type PersonalAccessToken struct {
	ID            uint
	TokenableType string
	TokenableID   uint
	Name          string
	Token         string
	Abilities     *string
	LastUsedAt    *time.Time
	ExpiresAt     *time.Time
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

// HistoryLog represents a history log entry
type HistoryLog struct {
	ID        uint
	UserID    uint
	Action    string
	Details   *string
	CreatedAt time.Time
	UpdatedAt time.Time
}
