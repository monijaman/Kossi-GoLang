package postgres

import (
	"kossti/internal/domain/repository"
	postgresqlRepo "kossti/internal/infrastructure/repository/postgresql"

	"gorm.io/gorm"
)

// NewFormGeneratorRepository creates a new PostgreSQL form generator repository
func NewFormGeneratorRepository(db *gorm.DB) repository.FormGeneratorRepository {
	return postgresqlRepo.NewFormGeneratorRepository(db)
}
