// Package entities contains domain entities that are framework-independent.
package entities

import (
	"time"
)

// Product represents a product in the domain
type Product struct {
	ID          uint
	Name        string
	Description *string
	Slug        string
	Price       float64
	CategoryID  *uint
	BrandID     *uint
	ViewsCount  int64
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

// ProductReview represents a product review
type ProductReview struct {
	ID        uint
	ProductID uint
	UserID    uint
	Rating    int
	Review    *string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

// Category represents a category
type Category struct {
	ID        uint
	Name      string
	Slug      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

// Brand represents a brand
type Brand struct {
	ID        uint
	Name      string
	Slug      string
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
