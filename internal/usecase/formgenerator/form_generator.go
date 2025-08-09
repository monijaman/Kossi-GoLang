// Package formgenerator provides use cases for form generator operations.
// This package is part of the use case layer in Clean Architecture.
// It contains business logic for form generator management.
package formgenerator

import (
	"context"
	"encoding/json"
	"errors"
	"kossti/internal/domain/entities"
	"kossti/internal/domain/repository"
	"time"
)

// CreateFormGenerator creates a new form generator or updates existing one
func CreateFormGenerator(ctx context.Context, repo repository.FormGeneratorRepository, categoryID uint, specificationIDs []uint, status string) (*entities.FormGenerator, error) {
	if categoryID == 0 {
		return nil, errors.New("category ID is required")
	}

	if len(specificationIDs) == 0 {
		return nil, errors.New("at least one specification ID is required")
	}

	if status == "" {
		status = "inactive"
	}

	// Convert specification IDs to JSON string
	specIDsJSON, err := json.Marshal(specificationIDs)
	if err != nil {
		return nil, errors.New("failed to encode specification IDs")
	}

	newFormGenerator := &entities.FormGenerator{
		CategoryID:      categoryID,
		SpecificationID: string(specIDsJSON),
		Status:          status,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	// Use upsert to create or update based on category_id
	result, err := repo.UpsertByCategoryID(ctx, newFormGenerator)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetFormGeneratorByID retrieves a form generator by ID
func GetFormGeneratorByID(ctx context.Context, repo repository.FormGeneratorRepository, id uint) (*entities.FormGenerator, error) {
	if id == 0 {
		return nil, errors.New("form generator ID is required")
	}

	formGenerator, err := repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return formGenerator, nil
}

// GetFormGeneratorByCategoryID retrieves form generators by category ID
func GetFormGeneratorByCategoryID(ctx context.Context, repo repository.FormGeneratorRepository, categoryID uint) (*entities.FormGenerator, error) {
	if categoryID == 0 {
		return nil, errors.New("category ID is required")
	}

	formGenerator, err := repo.GetByCategoryID(ctx, categoryID)
	if err != nil {
		return nil, err
	}

	return formGenerator, nil
}

// UpdateFormGenerator updates an existing form generator
func UpdateFormGenerator(ctx context.Context, repo repository.FormGeneratorRepository, id uint, categoryID *uint, specificationIDs []uint, status *string) (*entities.FormGenerator, error) {
	if id == 0 {
		return nil, errors.New("form generator ID is required")
	}

	// Get existing form generator
	existingFormGenerator, err := repo.GetByID(ctx, id)
	if err != nil {
		return nil, errors.New("form generator not found")
	}

	// Update fields if provided
	if categoryID != nil {
		existingFormGenerator.CategoryID = *categoryID
	}

	if len(specificationIDs) > 0 {
		specIDsJSON, err := json.Marshal(specificationIDs)
		if err != nil {
			return nil, errors.New("failed to encode specification IDs")
		}
		existingFormGenerator.SpecificationID = string(specIDsJSON)
	}

	if status != nil {
		existingFormGenerator.Status = *status
	}

	existingFormGenerator.UpdatedAt = time.Now()

	updatedFormGenerator, err := repo.Update(ctx, id, existingFormGenerator)
	if err != nil {
		return nil, err
	}

	return updatedFormGenerator, nil
}

// GetCategorySpecifications retrieves specifications for a category
func GetCategorySpecifications(ctx context.Context, repo repository.FormGeneratorRepository, categoryID uint) ([]*entities.Specification, error) {
	if categoryID == 0 {
		return nil, errors.New("category ID is required")
	}

	specifications, err := repo.GetCategorySpecifications(ctx, categoryID)
	if err != nil {
		return nil, err
	}

	return specifications, nil
}
