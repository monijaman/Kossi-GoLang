// Package entities contains domain entities that are framework-independent.
// These entities represent the business logic and should not depend on any external frameworks.
package entities

import (
	"time"
)

// Permission represents a permission in the domain
type Permission struct {
	ID        uint
	Name      string
	GuardName string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Role represents a role in the domain
type Role struct {
	ID        uint
	Name      string
	GuardName string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// ModelHasPermission represents the relationship between models and permissions
type ModelHasPermission struct {
	PermissionID uint
	ModelType    string
	ModelID      uint
}

// ModelHasRole represents the relationship between models and roles
type ModelHasRole struct {
	RoleID    uint
	ModelType string
	ModelID   uint
}

// RoleHasPermission represents the relationship between roles and permissions
type RoleHasPermission struct {
	PermissionID uint
	RoleID       uint
}
