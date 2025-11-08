# Production Migration Guide

## Quick Start

### Using Railway Database for Production

The migration tool now automatically detects and uses the production database configuration.

#### Method 1: Using GO_ENV (Recommended)

```bash
# Run any migration command with GO_ENV=production
GO_ENV=production go run ./cmd/migrate/main.go -migrate
GO_ENV=production go run ./cmd/migrate/main.go -fresh-seed
GO_ENV=production go run ./cmd/migrate/main.go -seed
```

#### Method 2: Auto-Detection

If `.env.production` file exists, the tool will automatically load it:

```bash
# The tool will detect .env.production and use Railway DB
go run ./cmd/migrate/main.go -migrate
```

## Environment Configuration

### .env.production (Railway Database)

```env
GO_ENV=production
DATABASE_URL=postgresql://postgres:wqxrKmVOvJeQORpcsGhWOGiRuQdJDLBU@switchyard.proxy.rlwy.net:58014/railway
```

### .env (Local Development)

```env
GO_ENV=development
DB_HOST=localhost
DB_PORT=5428
DB_USER=root
DB_PASSWORD=root
DB_NAME=kossti
DB_SSLMODE=disable
```

## Migration Commands

### Safe Production Migration

```bash
# Migrate schema (safe for production)
GO_ENV=production go run ./cmd/migrate/main.go -migrate
```

### Fresh Setup (Development/Testing Only)

```bash
# ⚠️ DANGEROUS: Drops all tables and recreates with seed data
GO_ENV=production go run ./cmd/migrate/main.go -fresh-seed
```

### Seeding Only

```bash
# Add seed data to existing database
GO_ENV=production go run ./cmd/migrate/main.go -seed
```

### Create Database

```bash
# Create database if it doesn't exist, then migrate
GO_ENV=production go run ./cmd/migrate/main.go -create-db
```

## All Available Options

| Command | Description | Safe for Production |
|---------|-------------|---------------------|
| `-migrate` | Run safe migration | ✅ Yes |
| `-create-db` | Create DB if not exists, then migrate | ✅ Yes |
| `-seed` | Run all seeders | ⚠️ Check first |
| `-fresh-seed` | Drop all, migrate, and seed | ❌ No (DANGEROUS) |
| `-fresh` | Drop all and migrate | ❌ No (DANGEROUS) |
| `-drop` | Drop all tables | ❌ No (DANGEROUS) |
| `-fk` | Add foreign keys only | ✅ Yes |
| `-indexes` | Create indexes only | ✅ Yes |

## Environment Loading Priority

1. If `GO_ENV=production` → loads `.env.production`
2. If `.env.production` file exists → loads it and sets `GO_ENV=production`
3. Otherwise → loads `.env` (local development)

## Database Connection

The tool will display the database being used (with password masked):

```
✅ Loaded .env.production
🔗 Using database: postgresql://postgres:****@switchyard.proxy.rlwy.net:58014/railway
```

## Troubleshooting

### Connection Refused Error

If you see:
```
Failed to connect to database: dial tcp [::1]:5428: connectex: No connection could be made
```

**Solution:** Make sure you're using `GO_ENV=production`:
```bash
GO_ENV=production go run ./cmd/migrate/main.go -migrate
```

### Wrong Database

If unsure which database is being used, check the log output:
```
🔗 Using database: postgresql://postgres:****@switchyard.proxy.rlwy.net:58014/railway
```

### Railway Connection String

Make sure your `.env.production` has the correct Railway connection:
```env
DATABASE_URL=postgresql://postgres:YOUR_PASSWORD@switchyard.proxy.rlwy.net:58014/railway
```

## Examples

### First Time Setup on Railway

```bash
# Create database and run initial migration with seed data
GO_ENV=production go run ./cmd/migrate/main.go -create-db
GO_ENV=production go run ./cmd/migrate/main.go -seed
```

### Update Schema on Railway

```bash
# Safe schema update (doesn't drop existing data)
GO_ENV=production go run ./cmd/migrate/main.go -migrate
```

### Reset Railway Database (Use with caution!)

```bash
# Complete reset with fresh data
GO_ENV=production go run ./cmd/migrate/main.go -fresh-seed
# Type 'YES' when prompted
```

### Local Development

```bash
# Uses local database from .env
go run ./cmd/migrate/main.go -migrate
go run ./cmd/migrate/main.go -seed
```

## What Gets Seeded

When you run `-seed` or `-fresh-seed`, the following data is added:

- ✅ **Categories** (128 records) - Product categories in English and Bengali
- ✅ **Brands** (152 records) - Global brands
- ✅ **Users** (Test users with hashed passwords)
- ✅ **Specification Keys** (Common product attributes)
- ✅ **Translations** (Bengali translations for categories, brands, specs)
- ✅ **Brand-Category Relations** (415 smart mappings)
- ✅ **Products** (Motorcycles from `init-db/seeders/motorbikes.json`)

## Security Notes

- ✅ `.env.production` is in `.gitignore` (credentials are safe)
- ✅ Password masking in logs (sensitive data protected)
- ✅ SSL enabled for Railway connections
- ⚠️ Never commit `.env.production` to git
- ⚠️ Use `-fresh` or `-drop` only in development/testing environments
