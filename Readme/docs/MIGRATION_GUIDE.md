# Database Migration Guide

This document explains how to migrate all tables from the Laravel `crit_api` project to the Go `gocrit_server` project using PostgreSQL.

## Overview

The migration system converts Laravel database migrations to Go GORM models and creates all tables in PostgreSQL. The system includes:

- **User Management**: Users, password reset tokens, sessions
- **Product System**: Products, reviews, categories, brands, translations
- **Permission System**: Roles, permissions, pivot tables
- **Media & Content**: Images, tags, feedback
- **System Tables**: Cache, jobs, history logs

## Quick Start

### 1. Set Up Environment

Make sure you have a `.env` file in the project root with your PostgreSQL connection:

```bash
DATABASE_URL=postgres://username:password@localhost:5432/gocrit_db?sslmode=disable
```

### 2. Run Migrations

#### Option A: Using the Migration Tool
```bash
# Navigate to project root
cd d:\GO\gocrit\gocrit_server

# Run migrations
go run tools/migrate/migrate.go -migrate

# Or reset database (drop all tables and recreate)
go run tools/migrate/migrate.go -reset

# Or just drop all tables
go run tools/migrate/migrate.go -drop
```

#### Option B: Using the Main Application
The main application will automatically run migrations on startup:
```bash
go run cmd/main.go
```

## Migrated Tables

### Core System Tables
- `users` - User accounts with name, email, password
- `password_reset_tokens` - Password reset functionality  
- `sessions` - User session management
- `cache` & `cache_locks` - Application caching
- `jobs`, `job_batches`, `failed_jobs` - Background job system
- `personal_access_tokens` - API token authentication
- `history_logs` - Activity tracking

### Permission System
- `permissions` - Available permissions
- `roles` - User roles
- `model_has_permissions` - User/model permissions
- `model_has_roles` - User/model roles  
- `role_has_permissions` - Role permissions

### Product System
- `products` - Product catalog
- `product_reviews` - Product reviews and ratings
- `categories` - Product categories
- `brands` - Product brands
- `brand_category` - Brand-category relationships

### Content & Media
- `images` - File attachments (polymorphic)
- `tags` - Content tagging
- `feedback` - User feedback

## Model Structure

### Key Features
- **Clean Architecture**: Models are separated from domain entities
- **GORM Integration**: Full PostgreSQL support with proper field types
- **Soft Deletes**: Non-destructive deletion for most entities
- **Timestamps**: Automatic created_at/updated_at tracking
- **Foreign Keys**: Proper relational constraints
- **Indexes**: Performance optimizations

### Example Model Structure
```go
type ProductModel struct {
    ID         uint       `gorm:"primaryKey;autoIncrement"`
    Name       string     `gorm:"type:varchar(255);not null"`
    Slug       string     `gorm:"type:varchar(255);unique;not null"`
    CategoryID string     `gorm:"type:varchar(255);not null"`
    Price      *float64   `gorm:"type:decimal(10,2)"`
    CreatedAt  time.Time  `gorm:"autoCreateTime"`
    UpdatedAt  time.Time  `gorm:"autoUpdateTime"`
    DeletedAt  *time.Time `gorm:"index"`
}
```

## Files Created

### Model Files
- `user.go` - Updated user model (matches Laravel structure)
- `password_reset_token.go` - Password reset tokens
- `session.go` - User sessions
- `cache.go` - Cache and cache locks
- `job.go` - Job system tables
- `history_log.go` - Activity logs
- `product.go` - Product catalog
- `product_review.go` - Product reviews
- `personal_access_token.go` - API tokens
- `permission.go` - Permission system
- `additional_models.go` - Feedback, images, tags, categories, brands

### Infrastructure Files
- `migration_manager.go` - Database migration orchestration
- `tools/migrate/migrate.go` - CLI migration tool

## Migration Process

1. **Parse Laravel Migrations**: Read PHP migration files
2. **Create Go Models**: Convert to GORM-compatible structs
3. **Auto-Migration**: Use GORM's AutoMigrate feature
4. **Add Constraints**: Foreign keys and indexes
5. **Verify Setup**: Test database connectivity

## Troubleshooting

### Common Issues

#### Connection Issues
```bash
# Check PostgreSQL is running
pg_isready -h localhost -p 5432

# Verify database exists
psql -h localhost -U username -d gocrit_db -c "SELECT version();"
```

#### Migration Errors
```bash
# Reset entire database
go run tools/migrate/migrate.go -reset

# Check GORM logs (enable in main.go)
db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{
    Logger: logger.Default.LogMode(logger.Info),
})
```

#### Field Type Mismatches
- Review model struct tags
- Check PostgreSQL data types
- Verify nullable fields use pointers

## Next Steps

After successful migration:

1. **Test Data Access**: Verify CRUD operations work
2. **Add Seeders**: Create initial data
3. **Set Up Relations**: Define GORM associations
4. **Add Validation**: Business rule enforcement
5. **Performance Tuning**: Query optimization

## Laravel to Go Mapping

| Laravel Type | Go Type | GORM Tag |
|-------------|---------|----------|
| `string()` | `string` | `type:varchar(255)` |
| `text()` | `string` | `type:text` |
| `longText()` | `string` | `type:longtext` |
| `integer()` | `int` | `type:integer` |
| `bigInteger()` | `int64` | `type:bigint` |
| `boolean()` | `bool` | `type:boolean` |
| `decimal(10,2)` | `float64` | `type:decimal(10,2)` |
| `timestamp()` | `time.Time` | `` |
| `nullable()` | `*Type` | `` |
| `unique()` | `` | `unique` |
| `index()` | `` | `index` |

The migration is now complete! All Laravel tables have been successfully converted to Go GORM models and can be created in PostgreSQL.
