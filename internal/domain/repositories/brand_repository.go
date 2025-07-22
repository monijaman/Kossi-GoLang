// Package repositories defines domain repository interfaces.
// These interfaces are framework-independent and define business requirements.
package repositories

import "kossti/internal/domain/entities"

// BrandRepository defines the interface for brand persistence operations.
// This interface is implemented by infrastructure layer repositories.
type BrandRepository interface {
	Create(brand *entities.Brand) error
	GetByID(id uint) (*entities.Brand, error)
	GetBySlug(slug string) (*entities.Brand, error)
	GetAll() ([]*entities.Brand, error)
	Update(brand *entities.Brand) error
	Delete(id uint) error
}
