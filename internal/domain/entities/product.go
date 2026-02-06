// Package entities contains domain entities that are framework-independent.
package entities

import (
	"encoding/json"
	"time"
)

// Product represents a product in the domain
type Product struct {
	ID          uint       `json:"id"`
	Name        string     `json:"name"`
	Description *string    `json:"description,omitempty"`
	Slug        string     `json:"slug"`
	Price       float64    `json:"price"`
	CategoryID  *uint      `json:"category_id,omitempty"`
	BrandID     *uint      `json:"brand_id,omitempty"`
	Category    *Category  `json:"category,omitempty"`
	Brand       *Brand     `json:"brand,omitempty"`
	ViewsCount  int64      `json:"views_count"`
	Status      bool       `json:"status"`
	Priority    int        `json:"priority"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}

// ProductReview represents a product review
type ProductReview struct {
	ID                uint
	ProductID         uint
	UserID            uint
	Rating            string
	Review            *string
	AdditionalDetails json.RawMessage `json:"additional_details,omitempty"`
	SourceURL         *string         `json:"source_url,omitempty"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         *time.Time
}

// Category represents a category
type Category struct {
	ID        uint
	Name      string
	Slug      string
	Status    int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

// Brand represents a brand
type Brand struct {
	ID        uint
	Name      string
	Slug      string
	Status    int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

// BrandCategory represents the brand-category relationship
type BrandCategory struct {
	ID         uint
	BrandID    uint
	CategoryID uint
}

// Productable represents polymorphic product relationships
type Productable struct {
	ProductID       uint
	ProductableID   uint
	ProductableType string
}

// Image represents an image
type Image struct {
	ID            uint
	ImageableType string
	ImageableID   uint
	ImagePath     string
	Status        int
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time
}

// Tag represents a tag
type Tag struct {
	ID        uint
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// FormGenerator represents form generator configuration
type FormGenerator struct {
	ID              uint
	CategoryID      uint
	SpecificationID string
	Status          string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

// Feedback represents user feedback for products
type Feedback struct {
	ID        uint
	UserID    uint
	Content   string
	Status    int
	CreatedBy *uint
	UpdatedBy *uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

// FeedbackTranslation represents feedback translations
type FeedbackTranslation struct {
	ID                uint
	FeedbackID        uint
	Locale            string
	TranslatedContent string
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
