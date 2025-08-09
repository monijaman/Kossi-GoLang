package postgres

import (
	"kossti/internal/domain/repository"
	postgresqlRepo "kossti/internal/infrastructure/repository/postgresql"

	"gorm.io/gorm"
)

// NewFeedbackRepository creates a new PostgreSQL feedback repository
func NewFeedbackRepository(db *gorm.DB) repository.FeedbackRepository {
	return postgresqlRepo.NewFeedbackRepository(db)
}
