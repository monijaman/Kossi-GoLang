// Package entities contains domain entities that are framework-independent.
package entities

import (
	"encoding/json"
	"time"
)

// Comment represents a user comment
type Comment struct {
	ID        uint
	ProductID int
	Username  string
	Location  string
	Comment   string
	Src       string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

// Specification represents a product specification
type Specification struct {
	ID                 uint
	ProductID          uint
	SpecificationKeyID uint
	SpecificationKey   string // This will be populated through joins
	Value              string
	Status             int
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

// SpecificationKey represents a specification key
type SpecificationKey struct {
	ID               uint
	SpecificationKey string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

// ProductTranslation represents product translations
type ProductTranslation struct {
	ID             uint
	ProductID      uint
	Locale         string
	TranslatedName string
	Price          *string `json:"price,omitempty"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

// CategoryTranslation represents category translations
type CategoryTranslation struct {
	ID             uint
	CategoryID     uint
	Locale         string
	TranslatedName string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

// BrandTranslation represents brand translations
type BrandTranslation struct {
	ID             uint
	BrandID        uint
	Locale         string
	TranslatedName string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

// ProductReviewTranslation represents product review translations
type ProductReviewTranslation struct {
	ID               uint
	ProductReviewID  uint
	Locale           string
	TranslatedReview string
	// AdditionalDetails holds any extra structured metadata for the translation
	AdditionalDetails json.RawMessage `json:"additional_details,omitempty"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

// CommentTranslation represents comment translations
type CommentTranslation struct {
	ID                uint
	CommentID         uint
	Locale            string
	TranslatedComment string
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

// SpecificationTranslation represents specification translations
type SpecificationTranslation struct {
	ID              uint
	SpecificationID uint
	Locale          string
	TranslatedValue string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

// SpecificationKeyTranslation represents specification key translations
type SpecificationKeyTranslation struct {
	ID                 uint
	SpecificationKeyID uint
	Locale             string
	TranslatedKey      string
	CreatedAt          time.Time
	UpdatedAt          time.Time
}
