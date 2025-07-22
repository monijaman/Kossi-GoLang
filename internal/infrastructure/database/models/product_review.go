// Package models contains database-specific implementations.
// This package is part of the infrastructure layer in Clean Architecture.
// These models are tightly coupled to GORM and handle database persistence.
package models

import (
	"kossti/internal/domain/entities"
	"time"
)

// ProductReviewModel represents the database model for product reviews (GORM-specific)
type ProductReviewModel struct {
	ID                uint       `gorm:"primaryKey;autoIncrement"`
	UserID            uint       `gorm:"not null"`
	Rating            float32    `gorm:"not null"`
	Reviews           *string    `gorm:"type:text"`
	AdditionalDetails *string    `gorm:"type:json"`
	Priority          int        `gorm:"default:1"`
	Status            bool       `gorm:"default:false"`
	CreatedAt         time.Time  `gorm:"autoCreateTime"`
	UpdatedAt         time.Time  `gorm:"autoUpdateTime"`
	DeletedAt         *time.Time `gorm:"index"`
}

// ToEntity converts GORM model to domain entity
func (pr *ProductReviewModel) ToEntity() *entities.ProductReview {
	return &entities.ProductReview{
		ID:        pr.ID,
		ProductID: 0, // Would need to be mapped from another field or relationship
		UserID:    pr.UserID,
		Rating:    int(pr.Rating),
		Review:    pr.Reviews,
		CreatedAt: pr.CreatedAt,
		UpdatedAt: pr.UpdatedAt,
		DeletedAt: pr.DeletedAt,
	}
}

// FromEntity converts domain entity to GORM model
func (pr *ProductReviewModel) FromEntity(entity *entities.ProductReview) {
	pr.ID = entity.ID
	pr.UserID = entity.UserID
	pr.Rating = float32(entity.Rating)
	pr.Reviews = entity.Review
	pr.CreatedAt = entity.CreatedAt
	pr.UpdatedAt = entity.UpdatedAt
	pr.DeletedAt = entity.DeletedAt
}

// TableName returns the table name for GORM
func (ProductReviewModel) TableName() string {
	return "product_reviews"
}
