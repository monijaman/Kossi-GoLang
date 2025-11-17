// Package models contains database-specific implementations.
// This package is part of the infrastructure layer in Clean Architecture.
// These models are tightly coupled to GORM and handle database persistence.
package models

import (
	"encoding/json"
	"fmt"
	"kossti/internal/domain/entities"
	"time"
)

// ProductableModel represents the database model for productable polymorphic relationships (GORM-specific)
type ProductableModel struct {
	ID              uint      `gorm:"primaryKey;autoIncrement"`
	ProductID       string    `gorm:"type:varchar(255);not null"`
	ProductableType string    `gorm:"type:varchar(255);not null"`
	ProductableID   uint      `gorm:"not null"`
	CreatedAt       time.Time `gorm:"autoCreateTime"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime"`
}

// ToEntity converts GORM model to domain entity
func (p *ProductableModel) ToEntity() *entities.Productable {
	return &entities.Productable{
		ProductID:       uint(p.ProductID[0]), // Convert string to uint if needed
		ProductableType: p.ProductableType,
		ProductableID:   p.ProductableID,
	}
}

// FromEntity converts domain entity to GORM model
func (p *ProductableModel) FromEntity(entity *entities.Productable) {
	p.ProductID = string(rune(entity.ProductID)) // Convert uint to string if needed
	p.ProductableType = entity.ProductableType
	p.ProductableID = entity.ProductableID
}

// TableName returns the table name for GORM
func (ProductableModel) TableName() string {
	return "productables"
}

// ProductTranslationModel represents the database model for product translations (GORM-specific)
type ProductTranslationModel struct {
	ID             uint      `gorm:"primaryKey;autoIncrement"`
	ProductID      uint      `gorm:"not null"`
	Locale         string    `gorm:"type:varchar(255);not null"`
	TranslatedName string    `gorm:"column:translated_name;type:varchar(255);not null"` // Database column is 'translated_name'
	Price          *string   `gorm:"type:varchar(255)"`                                 // Laravel uses varchar for price
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime"`
}

// ToEntity converts GORM model to domain entity
func (pt *ProductTranslationModel) ToEntity() *entities.ProductTranslation {
	return &entities.ProductTranslation{
		ID:             pt.ID,
		ProductID:      pt.ProductID,
		Locale:         pt.Locale,
		TranslatedName: pt.TranslatedName, // Direct mapping
		Price:          pt.Price,
		CreatedAt:      pt.CreatedAt,
		UpdatedAt:      pt.UpdatedAt,
	}
}

// FromEntity converts domain entity to GORM model
func (pt *ProductTranslationModel) FromEntity(entity *entities.ProductTranslation) {
	pt.ID = entity.ID
	pt.ProductID = entity.ProductID
	pt.Locale = entity.Locale
	pt.TranslatedName = entity.TranslatedName // Direct mapping
	pt.Price = entity.Price
	pt.CreatedAt = entity.CreatedAt
	pt.UpdatedAt = entity.UpdatedAt

	// Debug: Log the conversion
	fmt.Printf("FromEntity conversion: entity.TranslatedName='%s' -> model.TranslatedName='%s'\n",
		entity.TranslatedName, pt.TranslatedName)
}

// TableName returns the table name for GORM
func (ProductTranslationModel) TableName() string {
	return "product_translations"
}

// CommentModel represents the database model for user comments (GORM-specific)
type CommentModel struct {
	ID        uint       `gorm:"primaryKey;autoIncrement"`
	ProductID int        `gorm:"not null"`
	Username  string     `gorm:"type:varchar(255);not null"`
	Location  string     `gorm:"type:varchar(255);not null"`
	Comment   string     `gorm:"type:text;not null"`
	Src       string     `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time  `gorm:"autoCreateTime"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime"`
	DeletedAt *time.Time `gorm:"index"`
}

// ToEntity converts GORM model to domain entity
func (c *CommentModel) ToEntity() *entities.Comment {
	return &entities.Comment{
		ID:        c.ID,
		ProductID: c.ProductID,
		Username:  c.Username,
		Location:  c.Location,
		Comment:   c.Comment,
		Src:       c.Src,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
		DeletedAt: c.DeletedAt,
	}
}

// FromEntity converts domain entity to GORM model
func (c *CommentModel) FromEntity(entity *entities.Comment) {
	c.ID = entity.ID
	c.ProductID = entity.ProductID
	c.Username = entity.Username
	c.Location = entity.Location
	c.Comment = entity.Comment
	c.Src = entity.Src
	c.CreatedAt = entity.CreatedAt
	c.UpdatedAt = entity.UpdatedAt
	c.DeletedAt = entity.DeletedAt
}

// TableName returns the table name for GORM
func (CommentModel) TableName() string {
	return "comments"
}

// CommentTranslationModel represents the database model for comment translations (GORM-specific)
type CommentTranslationModel struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	CommentID uint      `gorm:"not null"`
	Locale    string    `gorm:"type:varchar(255);not null"`
	Comment   string    `gorm:"type:text;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

// ToEntity converts GORM model to domain entity
func (ct *CommentTranslationModel) ToEntity() *entities.CommentTranslation {
	return &entities.CommentTranslation{
		ID:                ct.ID,
		CommentID:         ct.CommentID,
		Locale:            ct.Locale,
		TranslatedComment: ct.Comment,
		CreatedAt:         ct.CreatedAt,
		UpdatedAt:         ct.UpdatedAt,
	}
}

// FromEntity converts domain entity to GORM model
func (ct *CommentTranslationModel) FromEntity(entity *entities.CommentTranslation) {
	ct.ID = entity.ID
	ct.CommentID = entity.CommentID
	ct.Locale = entity.Locale
	ct.Comment = entity.TranslatedComment
	ct.CreatedAt = entity.CreatedAt
	ct.UpdatedAt = entity.UpdatedAt
}

// TableName returns the table name for GORM
func (CommentTranslationModel) TableName() string {
	return "comment_translations"
}

// ProductReviewTranslationModel represents the database model for product review translations (GORM-specific)
type ProductReviewTranslationModel struct {
	ID              uint   `gorm:"primaryKey;autoIncrement"`
	ProductReviewID uint   `gorm:"not null"`
	Locale          string `gorm:"type:varchar(255);not null"`
	Rating          string `gorm:"type:varchar(50);default:''"`
	Reviews         string `gorm:"type:text;not null"`
	// Store any structured additional details as JSON bytes
	AdditionalDetails []byte    `gorm:"type:json"`
	CreatedAt         time.Time `gorm:"autoCreateTime"`
	UpdatedAt         time.Time `gorm:"autoUpdateTime"`
}

// ToEntity converts GORM model to domain entity
func (prt *ProductReviewTranslationModel) ToEntity() *entities.ProductReviewTranslation {
	return &entities.ProductReviewTranslation{
		ID:                prt.ID,
		ProductReviewID:   prt.ProductReviewID,
		Locale:            prt.Locale,
		Rating:            prt.Rating,
		TranslatedReview:  prt.Reviews,
		AdditionalDetails: json.RawMessage(prt.AdditionalDetails),
		CreatedAt:         prt.CreatedAt,
		UpdatedAt:         prt.UpdatedAt,
	}
}

// FromEntity converts domain entity to GORM model
func (prt *ProductReviewTranslationModel) FromEntity(entity *entities.ProductReviewTranslation) {
	prt.ID = entity.ID
	prt.ProductReviewID = entity.ProductReviewID
	prt.Locale = entity.Locale
	prt.Rating = entity.Rating
	prt.Reviews = entity.TranslatedReview
	if len(entity.AdditionalDetails) > 0 {
		prt.AdditionalDetails = []byte(entity.AdditionalDetails)
	}
	prt.CreatedAt = entity.CreatedAt
	prt.UpdatedAt = entity.UpdatedAt
}

// TableName returns the table name for GORM
func (ProductReviewTranslationModel) TableName() string {
	return "product_review_translations"
}
