// Package models contains database-specific implementations.
// This package is part of the infrastructure layer in Clean Architecture.
// These models are tightly coupled to GORM and handle database persistence.
package models

import (
	"kossti/internal/domain/entities"
)

// CacheModel represents the database model for cache entries (GORM-specific)
type CacheModel struct {
	Key        string `gorm:"primaryKey;type:varchar(255)"`
	Value      string `gorm:"type:text;not null"`
	Expiration int    `gorm:"not null"`
}

// ToEntity converts GORM model to domain entity
func (c *CacheModel) ToEntity() *entities.Cache {
	return &entities.Cache{
		Key:        c.Key,
		Value:      c.Value,
		Expiration: c.Expiration,
	}
}

// FromEntity converts domain entity to GORM model
func (c *CacheModel) FromEntity(entity *entities.Cache) {
	c.Key = entity.Key
	c.Value = entity.Value
	c.Expiration = entity.Expiration
}

// TableName returns the table name for GORM
func (CacheModel) TableName() string {
	return "cache"
}

// CacheLockModel represents the database model for cache locks (GORM-specific)
type CacheLockModel struct {
	Key        string `gorm:"primaryKey;type:varchar(255)"`
	Owner      string `gorm:"type:varchar(255);not null"`
	Expiration int    `gorm:"not null"`
}

// ToEntity converts GORM model to domain entity
func (cl *CacheLockModel) ToEntity() *entities.CacheLock {
	return &entities.CacheLock{
		Key:        cl.Key,
		Owner:      cl.Owner,
		Expiration: cl.Expiration,
	}
}

// FromEntity converts domain entity to GORM model
func (cl *CacheLockModel) FromEntity(entity *entities.CacheLock) {
	cl.Key = entity.Key
	cl.Owner = entity.Owner
	cl.Expiration = entity.Expiration
}

// TableName returns the table name for GORM
func (CacheLockModel) TableName() string {
	return "cache_locks"
}
