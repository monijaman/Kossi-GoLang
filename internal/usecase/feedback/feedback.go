// Package feedback provides use cases for feedback operations.
// This package is part of the use case layer in Clean Architecture.
// It contains business logic for feedback management.
package feedback

import (
	"context"
	"errors"
	"kossti/internal/domain/entities"
	"kossti/internal/domain/repository"
	"strconv"
	"strings"
	"time"
)

var ErrDailyFeedbackLimitReached = errors.New("you can submit a maximum of 10 feedback entries per day")

// CreateFeedback creates a new feedback for a product
func CreateFeedback(ctx context.Context, repo repository.FeedbackRepository, userID uint, productID uint, content string) (*entities.Feedback, error) {
	return CreateFeedbackWithDetails(ctx, repo, userID, productID, content, "", nil)
}

func CreateFeedbackWithDetails(ctx context.Context, repo repository.FeedbackRepository, userID uint, productID uint, content, rating string, sourceURL *string) (*entities.Feedback, error) {
	if userID == 0 {
		return nil, errors.New("user ID is required")
	}

	if productID == 0 {
		return nil, errors.New("product ID is required")
	}

	if content == "" {
		return nil, errors.New("content is required")
	}
	if err := validateRating(rating); err != nil {
		return nil, err
	}

	// Limit submissions across all products to ten per user per UTC calendar
	// day. Translations are stored separately and do not count as submissions.
	existing, err := repo.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	now := time.Now().UTC()
	dailyCount := 0
	for _, item := range existing {
		createdAt := item.CreatedAt.UTC()
		if createdAt.Year() == now.Year() && createdAt.YearDay() == now.YearDay() {
			dailyCount++
		}
	}
	if dailyCount >= 10 {
		return nil, ErrDailyFeedbackLimitReached
	}

	if len(content) > 2000 {
		return nil, errors.New("content must be 2000 characters or less")
	}

	feedback := &entities.Feedback{
		ProductID: productID,
		UserID:    userID,
		Rating:    rating,
		SourceURL: sourceURL,
		Content:   content,
		ContentEN: content,
		Status:    1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	createdFeedback, err := repo.Create(ctx, feedback)
	if err != nil {
		return nil, err
	}

	return createdFeedback, nil
}

// GetFeedbackByID retrieves a feedback by ID
func GetFeedbackByID(ctx context.Context, repo repository.FeedbackRepository, id uint) (*entities.Feedback, error) {
	if id == 0 {
		return nil, errors.New("feedback ID is required")
	}

	feedback, err := repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return feedback, nil
}

// GetAllFeedback retrieves all feedback with pagination
func GetAllFeedback(ctx context.Context, repo repository.FeedbackRepository, limit, offset int) ([]*entities.Feedback, error) {
	if limit <= 0 {
		limit = 10
	}

	if offset < 0 {
		offset = 0
	}

	feedbacks, err := repo.GetAll(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	return feedbacks, nil
}

// UpdateFeedback updates an existing feedback
func UpdateFeedback(ctx context.Context, repo repository.FeedbackRepository, id uint, content *string, status *int, rating *string, sourceURL *string) (*entities.Feedback, error) {
	if id == 0 {
		return nil, errors.New("feedback ID is required")
	}

	// Get existing feedback
	existingFeedback, err := repo.GetByID(ctx, id)
	if err != nil {
		return nil, errors.New("feedback not found")
	}

	// Update fields if provided
	if content != nil {
		if *content == "" {
			return nil, errors.New("content cannot be empty")
		}
		if len(*content) > 2000 {
			return nil, errors.New("content must be 2000 characters or less")
		}
		existingFeedback.Content = *content
	}

	if status != nil {
		existingFeedback.Status = *status
	}
	if rating != nil {
		if err := validateRating(*rating); err != nil {
			return nil, err
		}
		existingFeedback.Rating = *rating
	}
	if sourceURL != nil {
		existingFeedback.SourceURL = sourceURL
	}

	existingFeedback.UpdatedAt = time.Now()

	updatedFeedback, err := repo.Update(ctx, id, existingFeedback)
	if err != nil {
		return nil, err
	}

	return updatedFeedback, nil
}

func validateRating(rating string) error {
	rating = strings.TrimSpace(rating)
	value, err := strconv.Atoi(rating)
	if err != nil || value < 1 || value > 5 {
		return errors.New("rating must be between 1 and 5")
	}
	return nil
}

// DeleteFeedback soft deletes a feedback
func DeleteFeedback(ctx context.Context, repo repository.FeedbackRepository, id uint) error {
	if id == 0 {
		return errors.New("feedback ID is required")
	}

	// Check if feedback exists
	_, err := repo.GetByID(ctx, id)
	if err != nil {
		return errors.New("feedback not found")
	}

	return repo.Delete(ctx, id)
}

// GetProductFeedback retrieves all feedback for a specific product
func GetProductFeedback(ctx context.Context, repo repository.FeedbackRepository, productID uint) ([]*entities.Feedback, error) {
	if productID == 0 {
		return nil, errors.New("product ID is required")
	}

	feedbacks, err := repo.GetProductFeedbacks(ctx, productID)
	if err != nil {
		return nil, err
	}

	return feedbacks, nil
}

// CreateFeedbackTranslation creates a translation for feedback
func CreateFeedbackTranslation(ctx context.Context, repo repository.FeedbackRepository, feedbackID uint, locale, translatedContent string) (*entities.FeedbackTranslation, error) {
	if feedbackID == 0 {
		return nil, errors.New("feedback ID is required")
	}

	if locale == "" {
		return nil, errors.New("locale is required")
	}

	if translatedContent == "" {
		return nil, errors.New("translated content is required")
	}

	// Check if feedback exists
	_, err := repo.GetByID(ctx, feedbackID)
	if err != nil {
		return nil, errors.New("feedback not found")
	}

	translation := &entities.FeedbackTranslation{
		FeedbackID:        feedbackID,
		Locale:            locale,
		TranslatedContent: translatedContent,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}

	createdTranslation, err := repo.CreateTranslation(ctx, translation)
	if err != nil {
		return nil, err
	}

	return createdTranslation, nil
}
