// Package models contains database-specific implementations.
// This package is part of the infrastructure layer in Clean Architecture.
// These models are tightly coupled to GORM and handle database persistence.
package models

import (
	"kossti/internal/domain/entities"
	"time"
)

// CategoryTranslationModel represents the database model for category translations (GORM-specific)
type CategoryTranslationModel struct {
	ID         uint      `gorm:"primaryKey;autoIncrement"`
	CategoryID uint      `gorm:"not null"`
	Locale     string    `gorm:"type:varchar(255);not null"`
	Name       string    `gorm:"type:varchar(255);not null"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}

// ToEntity converts GORM model to domain entity
func (ct *CategoryTranslationModel) ToEntity() *entities.CategoryTranslation {
	return &entities.CategoryTranslation{
		ID:              ct.ID,
		CategoryID:      ct.CategoryID,
		Locale:          ct.Locale,
		TranslatedName:  ct.Name,
		CreatedAt:       ct.CreatedAt,
		UpdatedAt:       ct.UpdatedAt,
	}
}

// FromEntity converts domain entity to GORM model
func (ct *CategoryTranslationModel) FromEntity(entity *entities.CategoryTranslation) {
	ct.ID = entity.ID
	ct.CategoryID = entity.CategoryID
	ct.Locale = entity.Locale
	ct.Name = entity.TranslatedName
	ct.CreatedAt = entity.CreatedAt
	ct.UpdatedAt = entity.UpdatedAt
}

// TableName returns the table name for GORM
func (CategoryTranslationModel) TableName() string {
	return "category_translations"
}

// BrandTranslationModel represents the database model for brand translations (GORM-specific)
type BrandTranslationModel struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	BrandID   uint      `gorm:"not null"`
	Locale    string    `gorm:"type:varchar(255);not null"`
	Name      string    `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

// ToEntity converts GORM model to domain entity
func (bt *BrandTranslationModel) ToEntity() *entities.BrandTranslation {
	return &entities.BrandTranslation{
		ID:              bt.ID,
		BrandID:         bt.BrandID,
		Locale:          bt.Locale,
		TranslatedName:  bt.Name,
		CreatedAt:       bt.CreatedAt,
		UpdatedAt:       bt.UpdatedAt,
	}
}

// FromEntity converts domain entity to GORM model
func (bt *BrandTranslationModel) FromEntity(entity *entities.BrandTranslation) {
	bt.ID = entity.ID
	bt.BrandID = entity.BrandID
	bt.Locale = entity.Locale
	bt.Name = entity.TranslatedName
	bt.CreatedAt = entity.CreatedAt
	bt.UpdatedAt = entity.UpdatedAt
}

// TableName returns the table name for GORM
func (BrandTranslationModel) TableName() string {
	return "brand_translations"
}

// SpecificationKeyModel represents the database model for specification keys (GORM-specific)
type SpecificationKeyModel struct {
	ID               uint      `gorm:"primaryKey;autoIncrement"`
	SpecificationKey string    `gorm:"type:varchar(255);unique;not null"`
	CreatedAt        time.Time `gorm:"autoCreateTime"`
	UpdatedAt        time.Time `gorm:"autoUpdateTime"`
}

// ToEntity converts GORM model to domain entity
func (sk *SpecificationKeyModel) ToEntity() *entities.SpecificationKey {
	return &entities.SpecificationKey{
		ID:               sk.ID,
		SpecificationKey: sk.SpecificationKey,
		CreatedAt:        sk.CreatedAt,
		UpdatedAt:        sk.UpdatedAt,
	}
}

// FromEntity converts domain entity to GORM model
func (sk *SpecificationKeyModel) FromEntity(entity *entities.SpecificationKey) {
	sk.ID = entity.ID
	sk.SpecificationKey = entity.SpecificationKey
	sk.CreatedAt = entity.CreatedAt
	sk.UpdatedAt = entity.UpdatedAt
}

// TableName returns the table name for GORM
func (SpecificationKeyModel) TableName() string {
	return "specification_keys"
}

// SpecificationModel represents the database model for product specifications (GORM-specific)
type SpecificationModel struct {
	ID                 uint      `gorm:"primaryKey;autoIncrement"`
	ProductID          uint      `gorm:"not null"`
	SpecificationKeyID uint      `gorm:"not null"`
	Value              string    `gorm:"type:varchar(500);not null"`
	CreatedAt          time.Time `gorm:"autoCreateTime"`
	UpdatedAt          time.Time `gorm:"autoUpdateTime"`
}

// ToEntity converts GORM model to domain entity
func (s *SpecificationModel) ToEntity() *entities.Specification {
	return &entities.Specification{
		ID:               s.ID,
		ProductID:        s.ProductID,
		SpecificationKey: s.Value, // Map SpecificationKeyID to key for simplicity
		Value:            s.Value,
		CreatedAt:        s.CreatedAt,
		UpdatedAt:        s.UpdatedAt,
	}
}

// FromEntity converts domain entity to GORM model
func (s *SpecificationModel) FromEntity(entity *entities.Specification) {
	s.ID = entity.ID
	s.ProductID = entity.ProductID
	// s.SpecificationKeyID would need lookup logic
	s.Value = entity.Value
	s.CreatedAt = entity.CreatedAt
	s.UpdatedAt = entity.UpdatedAt
}

// TableName returns the table name for GORM
func (SpecificationModel) TableName() string {
	return "specifications"
}

// SpecificationTranslationModel represents the database model for specification translations (GORM-specific)
type SpecificationTranslationModel struct {
	ID              uint      `gorm:"primaryKey;autoIncrement"`
	SpecificationID uint      `gorm:"not null"`
	Locale          string    `gorm:"type:varchar(255);not null"`
	Value           string    `gorm:"type:varchar(500);not null"`
	CreatedAt       time.Time `gorm:"autoCreateTime"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime"`
}

// ToEntity converts GORM model to domain entity
func (st *SpecificationTranslationModel) ToEntity() *entities.SpecificationTranslation {
	return &entities.SpecificationTranslation{
		ID:              st.ID,
		SpecificationID: st.SpecificationID,
		Locale:          st.Locale,
		TranslatedValue: st.Value,
		CreatedAt:       st.CreatedAt,
		UpdatedAt:       st.UpdatedAt,
	}
}

// FromEntity converts domain entity to GORM model
func (st *SpecificationTranslationModel) FromEntity(entity *entities.SpecificationTranslation) {
	st.ID = entity.ID
	st.SpecificationID = entity.SpecificationID
	st.Locale = entity.Locale
	st.Value = entity.TranslatedValue
	st.CreatedAt = entity.CreatedAt
	st.UpdatedAt = entity.UpdatedAt
}

// TableName returns the table name for GORM
func (SpecificationTranslationModel) TableName() string {
	return "specification_translations"
}

// SpecificationKeyTranslationModel represents the database model for specification key translations (GORM-specific)
type SpecificationKeyTranslationModel struct {
	ID                 uint      `gorm:"primaryKey;autoIncrement"`
	SpecificationKeyID uint      `gorm:"not null"`
	Locale             string    `gorm:"type:varchar(255);not null"`
	SpecificationKey   string    `gorm:"type:varchar(255);not null"`
	CreatedAt          time.Time `gorm:"autoCreateTime"`
	UpdatedAt          time.Time `gorm:"autoUpdateTime"`
}

// ToEntity converts GORM model to domain entity
func (skt *SpecificationKeyTranslationModel) ToEntity() *entities.SpecificationKeyTranslation {
	return &entities.SpecificationKeyTranslation{
		ID:                  skt.ID,
		SpecificationKeyID:  skt.SpecificationKeyID,
		Locale:              skt.Locale,
		TranslatedKey:       skt.SpecificationKey,
		CreatedAt:           skt.CreatedAt,
		UpdatedAt:           skt.UpdatedAt,
	}
}

// FromEntity converts domain entity to GORM model
func (skt *SpecificationKeyTranslationModel) FromEntity(entity *entities.SpecificationKeyTranslation) {
	skt.ID = entity.ID
	skt.SpecificationKeyID = entity.SpecificationKeyID
	skt.Locale = entity.Locale
	skt.SpecificationKey = entity.TranslatedKey
	skt.CreatedAt = entity.CreatedAt
	skt.UpdatedAt = entity.UpdatedAt
}

// TableName returns the table name for GORM
func (SpecificationKeyTranslationModel) TableName() string {
	return "specification_key_translations"
}
