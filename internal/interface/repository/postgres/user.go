package postgres

// UserRepo is an implementation of the UserRepository interface for Postgres using GORM.
// - This file is part of the Interface Adapters Layer in Clean Code Architecture.
// - Adapts domain repository interfaces to actual database operations using GORM.
// - Keeps business logic decoupled from database details.
// - Makes it easy to swap out or mock the DB implementation for testing or future changes.

import (
	"context"
	"errors"

	"kossti/internal/domain/entities"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type PostgresUserRepo struct {
	db *gorm.DB
}

func NewPostgresUserRepo(db *gorm.DB) *PostgresUserRepo {
	return &PostgresUserRepo{db: db}
}

func (r *PostgresUserRepo) Create(ctx context.Context, user *entities.User) (*entities.User, error) {
	// Convert domain entity to GORM model
	var model models.UserModel
	model.FromEntity(user)

	// Create record in database
	if err := r.db.WithContext(ctx).Create(&model).Error; err != nil {
		return nil, err
	}

	// Convert back to domain entity with generated ID
	return model.ToEntity(), nil
}

func (r *PostgresUserRepo) GetByID(ctx context.Context, id uint) (*entities.User, error) {
	var model models.UserModel
	
	if err := r.db.WithContext(ctx).First(&model, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return model.ToEntity(), nil
}

func (r *PostgresUserRepo) GetByEmail(ctx context.Context, email string) (*entities.User, error) {
	var model models.UserModel
	
	if err := r.db.WithContext(ctx).Where("email = ?", email).First(&model).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return model.ToEntity(), nil
}

func (r *PostgresUserRepo) GetByUsername(ctx context.Context, username string) (*entities.User, error) {
	var model models.UserModel
	
	if err := r.db.WithContext(ctx).Where("name = ?", username).First(&model).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return model.ToEntity(), nil
}

func (r *PostgresUserRepo) Update(ctx context.Context, user *entities.User) error {
	var model models.UserModel
	model.FromEntity(user)
	return r.db.WithContext(ctx).Save(&model).Error
}

func (r *PostgresUserRepo) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.UserModel{}, id).Error
}

func (r *PostgresUserRepo) Count(ctx context.Context) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&models.UserModel{}).Count(&count).Error
	return count, err
}

func (r *PostgresUserRepo) List(ctx context.Context, limit, offset int) ([]*entities.User, error) {
	var models []models.UserModel
	
	if err := r.db.WithContext(ctx).Limit(limit).Offset(offset).Find(&models).Error; err != nil {
		return nil, err
	}

	users := make([]*entities.User, len(models))
	for i, model := range models {
		users[i] = model.ToEntity()
	}

	return users, nil
}

func (r *PostgresUserRepo) Search(ctx context.Context, query string, limit, offset int) ([]*entities.User, error) {
	var models []models.UserModel
	
	searchPattern := "%" + query + "%"
	if err := r.db.WithContext(ctx).
		Where("name ILIKE ? OR email ILIKE ?", searchPattern, searchPattern).
		Limit(limit).
		Offset(offset).
		Find(&models).Error; err != nil {
		return nil, err
	}

	users := make([]*entities.User, len(models))
	for i, model := range models {
		users[i] = model.ToEntity()
	}

	return users, nil
}

func (r *PostgresUserRepo) GetByEmailOrUsername(ctx context.Context, emailOrUsername string) (*entities.User, error) {
	var model models.UserModel
	
	if err := r.db.WithContext(ctx).
		Where("email = ? OR name = ?", emailOrUsername, emailOrUsername).
		First(&model).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return model.ToEntity(), nil
}
