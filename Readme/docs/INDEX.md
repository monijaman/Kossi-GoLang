# Documentation Index

This directory contains all technical documentation for the gocrit_server project.

## 📚 Documentation Files

### Core Documentation

- **[../README.md](../README.md)** - Main project README with architecture, setup, and usage guide (1662 lines)

### API Documentation

- **[PRODUCT_API.md](PRODUCT_API.md)** - Comprehensive Product API documentation with examples
- **[LARAVEL_PRODUCT_API.md](LARAVEL_PRODUCT_API.md)** - Legacy Laravel Product API reference
- **[API_COMPARISON.md](API_COMPARISON.md)** - Comparison between Go and Laravel implementations

### Architecture & Design

- **[CLEAN_ARCHITECTURE.md](CLEAN_ARCHITECTURE.md)** - Clean Architecture principles and implementation
- **[FIELD_UPDATES_GUIDE.md](FIELD_UPDATES_GUIDE.md)** - Guide for updating database fields

### Migration Documentation

- **[MIGRATION_GUIDE.md](MIGRATION_GUIDE.md)** - Step-by-step migration guide from Laravel to Go
- **[MIGRATION_COMPLETED.md](MIGRATION_COMPLETED.md)** - Migration completion status and notes

### AWS S3 Integration

- **[S3_UPLOAD_GUIDE.md](S3_UPLOAD_GUIDE.md)** - Guide for implementing S3 file uploads
- **[S3_ACCESS_FIX.md](S3_ACCESS_FIX.md)** - Troubleshooting S3 access issues
- **[IMAGE_DELETE_IMPLEMENTATION.md](IMAGE_DELETE_IMPLEMENTATION.md)** - Image deletion implementation guide

### Postman & Testing

- **[POSTMAN_UPLOAD_GUIDE.md](POSTMAN_UPLOAD_GUIDE.md)** - Guide for testing file uploads with Postman

## 📁 Other Resources

### Postman Collections

Located in `../postman/`:

- Product API collections
- Authentication endpoints
- File upload examples

### SQL Scripts

Located in project root (for debugging/diagnostics only):

- `check_products.sql`
- `check_product_reviews_schema.sql`
- `check_reviews_db.sql`

### Docker & Configuration

- `../docker-compose.yml` - Docker setup for PostgreSQL
- `../.env.example` - Environment variables template
- `../Makefile` - Build and run commands

## 🚀 Quick Links

### Getting Started

1. Read [README.md](../README.md) for setup instructions
2. Check [CLEAN_ARCHITECTURE.md](CLEAN_ARCHITECTURE.md) for architecture overview
3. Use [PRODUCT_API.md](PRODUCT_API.md) for API reference

### For Migrations

1. Follow [MIGRATION_GUIDE.md](MIGRATION_GUIDE.md) for Laravel → Go migration
2. Run migrations as described in main README

### For AWS S3

1. Start with [S3_UPLOAD_GUIDE.md](S3_UPLOAD_GUIDE.md)
2. Troubleshoot with [S3_ACCESS_FIX.md](S3_ACCESS_FIX.md)

## 📝 Notes

- All documentation is in Markdown format
- Legacy bug fix and issue tracking files have been removed
- Migration-related temporary files have been cleaned up
- All SQL scripts in root are for debugging only, not for running migrations
