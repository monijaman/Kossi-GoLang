// Package models contains database-specific implementations.
// This package is part of the infrastructure layer in Clean Architecture.
// These models are tightly coupled to GORM and handle database persistence.
package models

import (
	"kossti/internal/domain/entities"
	"time"
)

// PermissionModel represents the database model for permissions (GORM-specific)
type PermissionModel struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	Name      string    `gorm:"type:varchar(255);not null"`
	GuardName string    `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

// ToEntity converts GORM model to domain entity
func (p *PermissionModel) ToEntity() *entities.Permission {
	return &entities.Permission{
		ID:        p.ID,
		Name:      p.Name,
		GuardName: p.GuardName,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}

// FromEntity converts domain entity to GORM model
func (p *PermissionModel) FromEntity(entity *entities.Permission) {
	p.ID = entity.ID
	p.Name = entity.Name
	p.GuardName = entity.GuardName
	p.CreatedAt = entity.CreatedAt
	p.UpdatedAt = entity.UpdatedAt
}

// TableName returns the table name for GORM
func (PermissionModel) TableName() string {
	return "permissions"
}

// RoleModel represents the database model for roles (GORM-specific)
type RoleModel struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	Name      string    `gorm:"type:varchar(255);not null"`
	GuardName string    `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

// ToEntity converts GORM model to domain entity
func (r *RoleModel) ToEntity() *entities.Role {
	return &entities.Role{
		ID:        r.ID,
		Name:      r.Name,
		GuardName: r.GuardName,
		CreatedAt: r.CreatedAt,
		UpdatedAt: r.UpdatedAt,
	}
}

// FromEntity converts domain entity to GORM model
func (r *RoleModel) FromEntity(entity *entities.Role) {
	r.ID = entity.ID
	r.Name = entity.Name
	r.GuardName = entity.GuardName
	r.CreatedAt = entity.CreatedAt
	r.UpdatedAt = entity.UpdatedAt
}

// TableName returns the table name for GORM
func (RoleModel) TableName() string {
	return "roles"
}

// ModelHasPermissionModel represents the pivot table for model permissions (GORM-specific)
type ModelHasPermissionModel struct {
	PermissionID uint   `gorm:"primaryKey;not null"`
	ModelType    string `gorm:"primaryKey;type:varchar(255);not null"`
	ModelID      uint   `gorm:"primaryKey;not null"`
}

// ToEntity converts GORM model to domain entity
func (m *ModelHasPermissionModel) ToEntity() *entities.ModelHasPermission {
	return &entities.ModelHasPermission{
		PermissionID: m.PermissionID,
		ModelType:    m.ModelType,
		ModelID:      m.ModelID,
	}
}

// FromEntity converts domain entity to GORM model
func (m *ModelHasPermissionModel) FromEntity(entity *entities.ModelHasPermission) {
	m.PermissionID = entity.PermissionID
	m.ModelType = entity.ModelType
	m.ModelID = entity.ModelID
}

// TableName returns the table name for GORM
func (ModelHasPermissionModel) TableName() string {
	return "model_has_permissions"
}

// ModelHasRoleModel represents the pivot table for model roles (GORM-specific)
type ModelHasRoleModel struct {
	RoleID    uint   `gorm:"primaryKey;not null"`
	ModelType string `gorm:"primaryKey;type:varchar(255);not null"`
	ModelID   uint   `gorm:"primaryKey;not null"`
}

// ToEntity converts GORM model to domain entity
func (m *ModelHasRoleModel) ToEntity() *entities.ModelHasRole {
	return &entities.ModelHasRole{
		RoleID:    m.RoleID,
		ModelType: m.ModelType,
		ModelID:   m.ModelID,
	}
}

// FromEntity converts domain entity to GORM model
func (m *ModelHasRoleModel) FromEntity(entity *entities.ModelHasRole) {
	m.RoleID = entity.RoleID
	m.ModelType = entity.ModelType
	m.ModelID = entity.ModelID
}

// TableName returns the table name for GORM
func (ModelHasRoleModel) TableName() string {
	return "model_has_roles"
}

// RoleHasPermissionModel represents the pivot table for role permissions (GORM-specific)
type RoleHasPermissionModel struct {
	PermissionID uint `gorm:"primaryKey;not null"`
	RoleID       uint `gorm:"primaryKey;not null"`
}

// ToEntity converts GORM model to domain entity
func (r *RoleHasPermissionModel) ToEntity() *entities.RoleHasPermission {
	return &entities.RoleHasPermission{
		PermissionID: r.PermissionID,
		RoleID:       r.RoleID,
	}
}

// FromEntity converts domain entity to GORM model
func (r *RoleHasPermissionModel) FromEntity(entity *entities.RoleHasPermission) {
	r.PermissionID = entity.PermissionID
	r.RoleID = entity.RoleID
}

// TableName returns the table name for GORM
func (RoleHasPermissionModel) TableName() string {
	return "role_has_permissions"
}
