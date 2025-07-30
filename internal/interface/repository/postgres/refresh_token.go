package postgres

import (
	"context"
	"kossti/internal/domain/entities"
	"kossti/internal/domain/repository"

	"gorm.io/gorm"
)

type PostgresRefreshTokenRepo struct {
	db *gorm.DB
}

func NewPostgresRefreshTokenRepo(db *gorm.DB) *PostgresRefreshTokenRepo {
	return &PostgresRefreshTokenRepo{db: db}
}

func (r *PostgresRefreshTokenRepo) Save(ctx context.Context, token *entities.RefreshToken) error {
	return r.db.WithContext(ctx).Create(token).Error
}

func (r *PostgresRefreshTokenRepo) GetByToken(ctx context.Context, token string) (*entities.RefreshToken, error) {
	var rt entities.RefreshToken
	err := r.db.WithContext(ctx).Where("token = ?", token).First(&rt).Error
	if err != nil {
		return nil, err
	}
	return &rt, nil
}

func (r *PostgresRefreshTokenRepo) DeleteByToken(ctx context.Context, token string) error {
	return r.db.WithContext(ctx).Where("token = ?", token).Delete(&entities.RefreshToken{}).Error
}

func (r *PostgresRefreshTokenRepo) DeleteByUserID(ctx context.Context, userID uint) error {
	return r.db.WithContext(ctx).Where("user_id = ?", userID).Delete(&entities.RefreshToken{}).Error
}

var _ repository.RefreshTokenRepository = (*PostgresRefreshTokenRepo)(nil)
