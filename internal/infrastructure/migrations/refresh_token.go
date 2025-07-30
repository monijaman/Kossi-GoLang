package migrations

import (
	"kossti/internal/domain/entities"

	"gorm.io/gorm"
)

func MigrateRefreshToken(db *gorm.DB) error {
	return db.AutoMigrate(&entities.RefreshToken{})
}
