package seeders

import (
	"kossti/internal/domain/entities"
	infra_seeders "kossti/internal/infrastructure/database/seeders"

	"gorm.io/gorm"
)

// CreateOrFindSpecificationKey is a thin wrapper that forwards the call to the
// infrastructure seeders implementation. We provide this wrapper so the
// init-db/seeders package files can call CreateOrFindSpecificationKey without
// changing all existing usages; it keeps the invocation local to this package.
func CreateOrFindSpecificationKey(db *gorm.DB, key string) (*entities.SpecificationKey, error) {
	return infra_seeders.CreateOrFindSpecificationKey(db, key)
}
