// Package models contains database-specific implementations.
// This package is part of the infrastructure layer in Clean Architecture.
// These models are tightly coupled to GORM and handle database persistence.
package models

import (
	"kossti/internal/domain/entities"
	"time"
)

// JobModel represents the database model for jobs (GORM-specific)
type JobModel struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	Queue       string `gorm:"type:varchar(255);index;not null"`
	Payload     string `gorm:"type:text;not null"`
	Attempts    uint8  `gorm:"not null"`
	ReservedAt  *uint  `gorm:""`
	AvailableAt uint   `gorm:"not null"`
	CreatedAt   uint   `gorm:"not null"`
}

// ToEntity converts GORM model to domain entity
func (j *JobModel) ToEntity() *entities.Job {
	return &entities.Job{
		ID:          j.ID,
		Queue:       j.Queue,
		Payload:     j.Payload,
		Attempts:    uint(j.Attempts),
		ReservedAt:  j.ReservedAt,
		AvailableAt: j.AvailableAt,
		CreatedAt:   j.CreatedAt,
	}
}

// FromEntity converts domain entity to GORM model
func (j *JobModel) FromEntity(entity *entities.Job) {
	j.ID = entity.ID
	j.Queue = entity.Queue
	j.Payload = entity.Payload
	j.Attempts = uint8(entity.Attempts)
	j.ReservedAt = entity.ReservedAt
	j.AvailableAt = entity.AvailableAt
	j.CreatedAt = entity.CreatedAt
}

// TableName returns the table name for GORM
func (JobModel) TableName() string {
	return "jobs"
}

// JobBatchModel represents the database model for job batches (GORM-specific)
type JobBatchModel struct {
	ID           string  `gorm:"primaryKey;type:varchar(255)"`
	Name         string  `gorm:"type:varchar(255);not null"`
	TotalJobs    int     `gorm:"not null"`
	PendingJobs  int     `gorm:"not null"`
	FailedJobs   int     `gorm:"not null"`
	FailedJobIds string  `gorm:"type:text;not null"`
	Options      *string `gorm:"type:text"`
	CancelledAt  *int    `gorm:""`
	CreatedAt    int     `gorm:"not null"`
	FinishedAt   *int    `gorm:""`
}

// ToEntity converts GORM model to domain entity
func (jb *JobBatchModel) ToEntity() *entities.JobBatch {
	return &entities.JobBatch{
		ID:             jb.ID,
		Name:           jb.Name,
		TotalJobs:      jb.TotalJobs,
		PendingJobs:    jb.PendingJobs,
		FailedJobs:     jb.FailedJobs,
		FailedJobIDs:   jb.FailedJobIds,
		Options:        jb.Options,
		CancelledAt:    jb.CancelledAt,
		CreatedAt:      jb.CreatedAt,
		FinishedAt:     jb.FinishedAt,
	}
}

// FromEntity converts domain entity to GORM model
func (jb *JobBatchModel) FromEntity(entity *entities.JobBatch) {
	jb.ID = entity.ID
	jb.Name = entity.Name
	jb.TotalJobs = entity.TotalJobs
	jb.PendingJobs = entity.PendingJobs
	jb.FailedJobs = entity.FailedJobs
	jb.FailedJobIds = entity.FailedJobIDs
	jb.Options = entity.Options
	jb.CancelledAt = entity.CancelledAt
	jb.CreatedAt = entity.CreatedAt
	jb.FinishedAt = entity.FinishedAt
}

// TableName returns the table name for GORM
func (JobBatchModel) TableName() string {
	return "job_batches"
}

// FailedJobModel represents the database model for failed jobs (GORM-specific)
type FailedJobModel struct {
	ID         uint      `gorm:"primaryKey;autoIncrement"`
	UUID       string    `gorm:"type:varchar(255);unique;not null"`
	Connection string    `gorm:"type:text;not null"`
	Queue      string    `gorm:"type:text;not null"`
	Payload    string    `gorm:"type:text;not null"`
	Exception  string    `gorm:"type:text;not null"`
	FailedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

// ToEntity converts GORM model to domain entity
func (fj *FailedJobModel) ToEntity() *entities.FailedJob {
	return &entities.FailedJob{
		ID:         fj.ID,
		UUID:       fj.UUID,
		Connection: fj.Connection,
		Queue:      fj.Queue,
		Payload:    fj.Payload,
		Exception:  fj.Exception,
		FailedAt:   fj.FailedAt,
	}
}

// FromEntity converts domain entity to GORM model
func (fj *FailedJobModel) FromEntity(entity *entities.FailedJob) {
	fj.ID = entity.ID
	fj.UUID = entity.UUID
	fj.Connection = entity.Connection
	fj.Queue = entity.Queue
	fj.Payload = entity.Payload
	fj.Exception = entity.Exception
	fj.FailedAt = entity.FailedAt
}

// TableName returns the table name for GORM
func (FailedJobModel) TableName() string {
	return "failed_jobs"
}
