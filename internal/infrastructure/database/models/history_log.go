// Package models contains database-specific implementations.
// This package is part of the infrastructure layer in Clean Architecture.
// These models are tightly coupled to GORM and handle database persistence.
package models

import (
	"kossti/internal/domain/entities"
	"time"
)

// HistoryLogModel represents the database model for history logs (GORM-specific)
type HistoryLogModel struct {
	ID             uint      `gorm:"primaryKey;autoIncrement"`
	ReferenceTable string    `gorm:"type:varchar(255);not null"`
	ReferenceID    uint      `gorm:"not null"`
	ActorID        *uint     `gorm:""`
	Action         string    `gorm:"type:varchar(255);not null"`
	Body           string    `gorm:"type:varchar(255);not null"`
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime"`
}

// ToEntity converts GORM model to domain entity
func (h *HistoryLogModel) ToEntity() *entities.HistoryLog {
	return &entities.HistoryLog{
		ID:        h.ID,
		UserID:    h.ReferenceID, // Mapping ReferenceID to UserID for domain entity
		Action:    h.Action,
		Details:   &h.Body,
		CreatedAt: h.CreatedAt,
		UpdatedAt: h.UpdatedAt,
	}
}

// FromEntity converts domain entity to GORM model
func (h *HistoryLogModel) FromEntity(entity *entities.HistoryLog) {
	h.ID = entity.ID
	h.ReferenceID = entity.UserID
	h.Action = entity.Action
	if entity.Details != nil {
		h.Body = *entity.Details
	}
	h.CreatedAt = entity.CreatedAt
	h.UpdatedAt = entity.UpdatedAt
}

// TableName returns the table name for GORM
func (HistoryLogModel) TableName() string {
	return "history_logs"
}
