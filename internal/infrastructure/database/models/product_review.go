// Package models contains database-specific implementations.
// This package is part of the infrastructure layer in Clean Architecture.
// These models are tightly coupled to GORM and handle database persistence.
package models

import (
	"encoding/json"
	"kossti/internal/domain/entities"
	"time"
)

// ProductReviewModel represents the database model for product reviews (GORM-specific)
type ProductReviewModel struct {
	ID                uint       `gorm:"primaryKey;autoIncrement"`
	ProductID         uint       `gorm:"not null;index"`
	UserID            uint       `gorm:"not null;index"`
	Rating            string     `gorm:"type:varchar(50);default:''"`
	Reviews           *string    `gorm:"type:text"`
	SourceURL         *string    `gorm:"type:text"`
	AdditionalDetails []byte     `gorm:"type:json"`
	Priority          int        `gorm:"default:1"`
	Status            bool       `gorm:"default:false"`
	CreatedAt         time.Time  `gorm:"autoCreateTime"`
	UpdatedAt         time.Time  `gorm:"autoUpdateTime"`
	DeletedAt         *time.Time `gorm:"index"`
}

// TableName specifies the table name for GORM
func (ProductReviewModel) TableName() string {
	return "product_reviews"
}

// ToEntity converts GORM model to domain entity
func (pr *ProductReviewModel) ToEntity() *entities.ProductReview {
	return &entities.ProductReview{
		ID:                pr.ID,
		ProductID:         pr.ProductID,
		UserID:            pr.UserID,
		Rating:            pr.Rating,
		Review:            pr.Reviews,
		AdditionalDetails: json.RawMessage(pr.AdditionalDetails),
		SourceURL:         pr.SourceURL,
		CreatedAt:         pr.CreatedAt,
		UpdatedAt:         pr.UpdatedAt,
		DeletedAt:         pr.DeletedAt,
	}
}

// FromEntity converts domain entity to GORM model
func (pr *ProductReviewModel) FromEntity(entity *entities.ProductReview) {
	pr.ID = entity.ID
	pr.ProductID = entity.ProductID
	pr.UserID = entity.UserID
	pr.Rating = entity.Rating
	pr.Reviews = entity.Review
	pr.SourceURL = entity.SourceURL
	if len(entity.AdditionalDetails) > 0 {
		pr.AdditionalDetails = []byte(entity.AdditionalDetails)
	} else {
		pr.AdditionalDetails = nil
	}
	pr.CreatedAt = entity.CreatedAt
	pr.UpdatedAt = entity.UpdatedAt
	pr.DeletedAt = entity.DeletedAt
}
