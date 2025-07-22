// Package repositories contains repository interfaces for the domain layer.
// These interfaces define what the domain needs from persistence layer,
// without depending on any specific implementation.
package repositories

import (
	"kossti/internal/domain/entities"
)

// PermissionRepository defines the interface for permission persistence operations
type PermissionRepository interface {
	Create(permission *entities.Permission) error
	GetByID(id uint) (*entities.Permission, error)
	GetByName(name string) (*entities.Permission, error)
	GetAll() ([]*entities.Permission, error)
	Update(permission *entities.Permission) error
	Delete(id uint) error
}

// RoleRepository defines the interface for role persistence operations
type RoleRepository interface {
	Create(role *entities.Role) error
	GetByID(id uint) (*entities.Role, error)
	GetByName(name string) (*entities.Role, error)
	GetAll() ([]*entities.Role, error)
	Update(role *entities.Role) error
	Delete(id uint) error
}

// ModelPermissionRepository handles model-permission relationships
type ModelPermissionRepository interface {
	AssignPermissionToModel(modelType string, modelID uint, permissionID uint) error
	RevokePermissionFromModel(modelType string, modelID uint, permissionID uint) error
	GetModelPermissions(modelType string, modelID uint) ([]*entities.Permission, error)
}

// ModelRoleRepository handles model-role relationships
type ModelRoleRepository interface {
	AssignRoleToModel(modelType string, modelID uint, roleID uint) error
	RevokeRoleFromModel(modelType string, modelID uint, roleID uint) error
	GetModelRoles(modelType string, modelID uint) ([]*entities.Role, error)
}

// RolePermissionRepository handles role-permission relationships
type RolePermissionRepository interface {
	AssignPermissionToRole(roleID uint, permissionID uint) error
	RevokePermissionFromRole(roleID uint, permissionID uint) error
	GetRolePermissions(roleID uint) ([]*entities.Permission, error)
}
