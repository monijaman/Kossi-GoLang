// Package usecase contains application business logic.
// Use cases orchestrate the flow of data to and from entities,
// and direct entities to use their business rules to achieve the goals of the use case.
package usecase

import (
	"kossti/internal/domain/entities"
	"kossti/internal/domain/repositories"
	"errors"
)

// PermissionUseCase handles permission-related business logic
type PermissionUseCase struct {
	permissionRepo repositories.PermissionRepository
	roleRepo       repositories.RoleRepository
}

// NewPermissionUseCase creates a new permission use case
func NewPermissionUseCase(
	permissionRepo repositories.PermissionRepository,
	roleRepo repositories.RoleRepository,
) *PermissionUseCase {
	return &PermissionUseCase{
		permissionRepo: permissionRepo,
		roleRepo:       roleRepo,
	}
}

// CreatePermission creates a new permission with business logic validation
func (uc *PermissionUseCase) CreatePermission(name, guardName string) (*entities.Permission, error) {
	// Business logic: Check if permission already exists
	existing, _ := uc.permissionRepo.GetByName(name)
	if existing != nil {
		return nil, errors.New("permission already exists")
	}

	// Business logic: Validate permission name
	if name == "" {
		return nil, errors.New("permission name cannot be empty")
	}

	permission := &entities.Permission{
		Name:      name,
		GuardName: guardName,
	}

	if err := uc.permissionRepo.Create(permission); err != nil {
		return nil, err
	}

	return permission, nil
}

// GetAllPermissions retrieves all permissions
func (uc *PermissionUseCase) GetAllPermissions() ([]*entities.Permission, error) {
	return uc.permissionRepo.GetAll()
}

// CreateRole creates a new role with business logic validation
func (uc *PermissionUseCase) CreateRole(name, guardName string) (*entities.Role, error) {
	// Business logic: Check if role already exists
	existing, _ := uc.roleRepo.GetByName(name)
	if existing != nil {
		return nil, errors.New("role already exists")
	}

	// Business logic: Validate role name
	if name == "" {
		return nil, errors.New("role name cannot be empty")
	}

	role := &entities.Role{
		Name:      name,
		GuardName: guardName,
	}

	if err := uc.roleRepo.Create(role); err != nil {
		return nil, err
	}

	return role, nil
}

// GetAllRoles retrieves all roles
func (uc *PermissionUseCase) GetAllRoles() ([]*entities.Role, error) {
	return uc.roleRepo.GetAll()
}
