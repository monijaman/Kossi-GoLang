package repository

import (
	"context"
	"kossti/internal/domain/entities"
)

type RefreshTokenRepository interface {
	Save(ctx context.Context, token *entities.RefreshToken) error
	GetByToken(ctx context.Context, token string) (*entities.RefreshToken, error)
	DeleteByToken(ctx context.Context, token string) error
	DeleteByUserID(ctx context.Context, userID uint) error
}
