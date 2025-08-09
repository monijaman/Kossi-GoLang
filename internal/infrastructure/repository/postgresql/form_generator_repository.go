// Package postgresql provides PostgreSQL-specific implementations of repository interfaces.
// This package is part of the infrastructure layer in Clean Architecture.
// It contains concrete implementations that interact with PostgreSQL database using GORM.
package postgresql

import (
	"context"
	"kossti/internal/domain/entities"
	"kossti/internal/domain/repository"
	models "kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// formGeneratorRepository implements the FormGeneratorRepository interface using PostgreSQL
type formGeneratorRepository struct {
	db *gorm.DB
}

// NewFormGeneratorRepository creates a new PostgreSQL form generator repository
func NewFormGeneratorRepository(db *gorm.DB) repository.FormGeneratorRepository {
	return &formGeneratorRepository{
		db: db,
	}
}

// Create creates a new form generator
func (r *formGeneratorRepository) Create(ctx context.Context, formGenerator *entities.FormGenerator) (*entities.FormGenerator, error) {
	model := &models.FormGeneratorModel{}
	model.FromEntity(formGenerator)

	if err := r.db.WithContext(ctx).Create(model).Error; err != nil {
		return nil, err
	}

	entity := model.ToEntity()
	return entity, nil
}

// GetByID retrieves a form generator by ID
func (r *formGeneratorRepository) GetByID(ctx context.Context, id uint) (*entities.FormGenerator, error) {
	var model models.FormGeneratorModel
	if err := r.db.WithContext(ctx).First(&model, id).Error; err != nil {
		return nil, err
	}
	return model.ToEntity(), nil
}

// GetByCategoryID retrieves a form generator by category ID
func (r *formGeneratorRepository) GetByCategoryID(ctx context.Context, categoryID uint) (*entities.FormGenerator, error) {
	var model models.FormGeneratorModel
	if err := r.db.WithContext(ctx).Where("category_id = ?", categoryID).First(&model).Error; err != nil {
		return nil, err
	}
	return model.ToEntity(), nil
}

// Update updates a form generator
func (r *formGeneratorRepository) Update(ctx context.Context, id uint, formGenerator *entities.FormGenerator) (*entities.FormGenerator, error) {
	model := &models.FormGeneratorModel{}
	model.FromEntity(formGenerator)
	model.ID = id

	if err := r.db.WithContext(ctx).Model(&model).Updates(model).Error; err != nil {
		return nil, err
	}

	return model.ToEntity(), nil
}

// Delete deletes a form generator
func (r *formGeneratorRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.FormGeneratorModel{}, id).Error
}

// UpsertByCategoryID creates or updates a form generator based on category_id
func (r *formGeneratorRepository) UpsertByCategoryID(ctx context.Context, formGenerator *entities.FormGenerator) (*entities.FormGenerator, error) {
	var existingModel models.FormGeneratorModel

	// Check if record exists for this category
	err := r.db.WithContext(ctx).Where("category_id = ?", formGenerator.CategoryID).First(&existingModel).Error

	if err == gorm.ErrRecordNotFound {
		// Record doesn't exist, create new
		model := &models.FormGeneratorModel{}
		model.FromEntity(formGenerator)

		if err := r.db.WithContext(ctx).Create(model).Error; err != nil {
			return nil, err
		}

		return model.ToEntity(), nil
	} else if err != nil {
		// Some other error occurred
		return nil, err
	}

	// Record exists, update it
	existingModel.SpecificationID = formGenerator.SpecificationID
	existingModel.Status = formGenerator.Status

	if err := r.db.WithContext(ctx).Save(&existingModel).Error; err != nil {
		return nil, err
	}

	return existingModel.ToEntity(), nil
}

// GetCategorySpecifications retrieves specifications for a category
func (r *formGeneratorRepository) GetCategorySpecifications(ctx context.Context, categoryID uint) ([]*entities.Specification, error) {
	var formGenerator models.FormGeneratorModel
	err := r.db.WithContext(ctx).Where("category_id = ?", categoryID).First(&formGenerator).Error

	if err == gorm.ErrRecordNotFound {
		return []*entities.Specification{}, nil
	} else if err != nil {
		return nil, err
	}

	// Parse specification IDs from JSON and get specifications
	// For now, return empty slice as we need to implement SpecificationRepository
	// TODO: Implement proper specification retrieval based on specification_id JSON array
	return []*entities.Specification{}, nil
}
