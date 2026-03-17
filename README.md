Go Backend of https://kossti.com/

# Clean Architecture Authentication System in Go — Presentation Guide

Welcome! This guide is designed to help you present and explain your Go authentication project using Clean Architecture to students.

Short description: This repository contains the Go backend for the Kossti product/review system. It follows Clean Architecture principles, exposes HTTP handlers for product and review management, uses GORM with PostgreSQL for persistence, and includes migration and seeding tools used during development and presentation.

---

## 📋 Table of Contents

### 🏗️ **Architecture & Concepts**

- [1. What is Clean Architecture?](#1-what-is-clean-architecture)
- [2. Project Structure (Visual)](#2-project-structure-visual)
- [3. Layers Explained (Simple)](#3-layers-explained-simple)
- [4. How the Parts Work Together](#4-how-the-parts-work-together)


### � **API Reference**

- [6.5. API Endpoints Reference](#65-api-endpoints-reference)
  - [Authentication & User Management](#-authentication--user-management)
  - [Product Management](#-product-management)
  - [Reviews & Feedback](#-reviews--feedback)
  - [Categories](#-categories)
  - [Brands](#-brands)
  - [Specifications](#-specifications)
  - [Form Generator](#-form-generator)
  - [Example API Requests](#-example-api-requests)


---

## 1. What is Clean Architecture?

- **Clean Architecture** is a way to organize code so that each part has a clear job.
- Business logic is at the center and does not depend on frameworks or databases.
- Makes code easier to test, change, and maintain.

---

## 2. Project Structure (Visual)

```
kossti/
├── cmd/                  # Application entry points
│   ├── app/              # Main application
│   │   └── main.go
│   └── migrate/          # Database migration CLI tool
│       └── main.go
├── internal/
│   ├── domain/           # Core business logic (entities, interfaces)
│   │   ├── entities/     # Domain entities (User, Product, etc.)
│   │   └── repositories/ # Repository interfaces
│   ├── usecase/          # Application logic (register, login)
│   ├── interface/        # Adapters (HTTP handlers, DB impl)
│   │   ├── handler/      # HTTP request handlers
│   │   └── repository/   # Database implementations
│   └── infrastructure/   # Frameworks (DB, JWT, Kafka)
│       └── database/     # Database models and migration manager
└── pkg/                  # Shared code (e.g., password hashing)
```

---

## 3. Layers Explained (Simple)

- **Domain Layer:** User struct, repository interface (pure business rules)
- **Use Case Layer:** Register and login logic (no frameworks)
- **Interface Layer:** HTTP handlers, DB implementations (adapts input/output)
- **Infrastructure Layer:** Database connection, JWT, Kafka (details that can change)

---

## 4. How the Parts Work Together

1. User sends a request (e.g., register) to the HTTP handler.
2. Handler calls the use case (register logic).
3. Use case uses the repository interface to save the user.
4. Repository implementation talks to the database.
5. Infrastructure handles the actual DB connection.

---

## 4.5. Quick Start (Working Tools) 🚀

If you want to get started immediately with the working components:

````bash
# 1. Install dependencies
go mod download

# 2. Install Air for hot reloading (optional)
go install github.com/cosmtrek/air@latest

# 3. Create database and run migration
go run ./cmd/migrate/main.go -create-db

# 4. Seed with sample data
go run ./cmd/migrate/main.go -seed

# 5. Verify setup
go run ./cmd/seedtest/main.go

### ✅ Seeding custom motorbikes dataset

This project supports custom seeders. A motorbikes dataset is provided in JSON and the product seeder reads it automatically when you run the seeder tool.

- Dataset path (relative to the `gocrit_server` root): `init-db/seeders/motorbikes.json`
- What it contains: name_en, name_bn, brand, model, price, images (array of URLs).

How to seed the motorbikes dataset:

1. Make sure your `DATABASE_URL` env var is set (or a `.env` is present in `gocrit_server`).
2. Run the seed command (this will run all seeders including the motorbikes seeder):

```bash
cd /d/GO/gocrit/gocrit_server
go run ./cmd/migrate/main.go -seed
````

3. To perform a fresh setup (drop tables, migrate, then seed — DANGEROUS on production):

```bash
go run ./cmd/migrate/main.go -fresh-seed
```

4. Verify motorbikes were inserted (quick check using the verification tool or direct DB queries):

```bash
go run ./cmd/seedtest/main.go
# SQL quick checks (psql / any DB client):
# SELECT count(*) FROM products WHERE slug LIKE '%';
# SELECT * FROM product_translations WHERE locale = 'bn' LIMIT 20;
```

Notes:

- The seeder will skip creating a product if a product with the same slug already exists.
- Images in the dataset use placeholder URLs by default; replace them with real S3 or hosted URLs if you have them.
- To expand or edit the motorbikes list, modify `init-db/seeders/motorbikes.json` and re-run `-seed` (or `-fresh-seed` for a full refresh).

````

**What's working:**
✅ Database migration system (24 models)
✅ Comprehensive seeding system (400+ records)
✅ Data verification tools
✅ Clean Architecture structure
✅ Air hot reloading (with working migration tools)

**What needs attention:**
⚠️ Main API server has entity import conflicts
⚠️ Use case layer needs entity reconciliation

### 🔥 Quick Fix for Air Issues

If you get `'tmp\main.exe' is not recognized` error with Air:

```bash
# Option 1: Use Air with migration tools (works perfectly)
air -c .air-migrate.toml

# Option 2: Use direct commands (reliable)
go run ./cmd/migrate/main.go -create-db
go run ./cmd/migrate/main.go -seed

# Option 3: Manual build test to see specific errors
mkdir tmp
go build -o ./tmp/main.exe ./cmd/app/main.go
````

### 🚨 Entity Import Conflict Issue

The main app currently has compilation errors due to conflicting entity packages:

**Error:** `cannot use user (variable of type *entity.User) as *entities.User value`

**Solution:** Use the working tools while this is resolved:

```bash
# ✅ WORKING: Migration and seeding tools
go run ./cmd/migrate/main.go -create-db  # Create database
go run ./cmd/migrate/main.go -seed       # Add sample data
go run ./cmd/seedtest/main.go           # Verify data

# ✅ WORKING: Air with migration tools
air -c .air-migrate.toml                # Hot reload for DB work

# ❌ NOT WORKING: Main API server (entity conflicts)
# go run ./cmd/app/main.go              # Has compilation errors
```

---

## 5. Setup & Run (Step-by-Step)

### 1. Install Go and PostgreSQL

- [Go download](https://go.dev/dl/) (Go 1.19 or higher required)
- [PostgreSQL download](https://www.postgresql.org/download/)

### 2. Install Project Dependencies

#### Install Go Dependencies

```bash
# Navigate to project directory
cd gocrit_server

# Install all Go dependencies
go mod download

# Verify dependencies
go mod verify
```

#### Install Development Tools (Optional but Recommended)

```bash
# Install Air for hot reloading during development
go install github.com/cosmtrek/air@latest

# Install golangci-lint for code linting
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Install migrate tool for database migrations (alternative CLI)
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

### 3. Create a Database

```bash
createdb -U postgres kossti
```

### 4. Set Environment Variables

Create a `.env` file in the project root:

```env
# Database Configuration
DATABASE_URL=postgres://root:root@localhost:5428/kossti?sslmode=disable&TimeZone=UTC
JWT_SECRET=your_jwt_secret_key_here
KAFKA_BROKER=localhost:9092

# Server Configuration
PORT=8080
GO_ENV=development
```

### 5. Run the Applications

#### Main Application (API Server)

**Standard Go Run:**

```bash
go run ./cmd/app/main.go
```

**With Air for Hot Reloading (For Working Tools):**

_Note: The main app has compilation issues. Use Air for the working migration tools:_

```bash
# Option 1: Use Air with the fixed .air.toml (will show build errors)
air

# Option 2: Use the working migration tools directly
go run ./cmd/migrate/main.go -migrate
go run ./cmd/migrate/main.go -seed
go run ./cmd/seedtest/main.go

# Option 3: Manual build test to see errors
go build -o ./tmp/main.exe ./cmd/app/main.go
```

**Air Configuration:**
Your project includes multiple `.air.toml` configuration files:

- `.air.toml` - For main application (has compilation issues)
- `.air-migrate.toml` - For migration tools (working alternative)

If the main Air fails, use the migration-focused config:

```bash
# Use the working migration Air config
air -c .air-migrate.toml

# This will run the migration tool with hot reloading
# Default command: ./tmp/migrate.exe -h (shows help)
```

To customize the migration Air config for specific tasks:

```toml
# Edit .air-migrate.toml and change args_bin:
args_bin = ["-migrate"]     # Auto-run migration on changes
# or
args_bin = ["-seed"]        # Auto-run seeding on changes
# or
args_bin = ["-h"]          # Show help (default, safe)
```

#### Database Migration Tool

```sh
# Run safe migration (recommended)
go run ./cmd/migrate/main.go -migrate

# See all migration options
go run ./cmd/migrate/main.go -h
```

### 6. Test the API

> **Note:**
> Make sure your Go server registers the `/register` and `/login` endpoints using `http.HandleFunc` or your web framework. If you get a 404 error, check your route registration in main.go.

- **Register:**
  - Method: POST
  - URL: `http://localhost:8080/register`
  - Body (JSON):
    ```json
    {
      "username": "alice",
      "email": "alice@example.com",
      "password": "password123"
    }
    ```
- **Login:**
  - Method: POST
  - URL: `http://localhost:8080/login`
  - Body (JSON):
    ```json
    {
      "email": "alice@example.com",
      "password": "password123"
    }
    ```

---

## 6.5. API Endpoints Reference

This section provides a comprehensive list of all available API endpoints in the system.

### 📝 Authentication & User Management

#### Authentication Endpoints

| Method | Endpoint              | Description                       | Auth Required |
| ------ | --------------------- | --------------------------------- | ------------- |
| POST   | `/register`           | User registration                 | No            |
| POST   | `/login`              | User login                        | No            |
| POST   | `/v1/registration`    | Alternative registration endpoint | No            |
| POST   | `/v1/login`           | V1 login endpoint                 | No            |
| POST   | `/v1/refresh-token`   | Refresh access token              | No            |
| POST   | `/v1/forgot-password` | Password reset request            | No            |
| POST   | `/v1/reset-password`  | Password reset                    | No            |
| POST   | `/v1/logout`          | User logout                       | Yes           |
| GET    | `/v1/check-token`     | Check token validity              | Yes           |
| GET    | `/v1/profile`         | Get user profile                  | Yes           |
| POST   | `/v1/profile`         | Update user profile               | Yes           |
| GET    | `/v1/checkrole`       | Check user role                   | Yes           |
| GET    | `/v1/search_users`    | Search users                      | Yes           |

#### User Endpoints

| Method | Endpoint        | Description                | Auth Required |
| ------ | --------------- | -------------------------- | ------------- |
| GET    | `/users/all`    | Get all users (public)     | No            |
| GET    | `/users`        | List users with pagination | Yes           |
| GET    | `/users/{id}`   | Get user by ID             | Yes           |
| GET    | `/users/search` | Search users               | Yes           |
| GET    | `/users/count`  | Get user count             | Yes           |

### 🛍️ Product Management

#### Product Endpoints

| Method | Endpoint                          | Description                  | Auth Required |
| ------ | --------------------------------- | ---------------------------- | ------------- |
| GET    | `/products`                       | List products (basic)        | No            |
| GET    | `/products?locale=&page=&sortby=` | List products with filtering | No            |
| POST   | `/products`                       | Create new product           | Yes           |
| GET    | `/products/{id}`                  | Get product by ID or slug    | No            |
| PATCH  | `/products/{id}`                  | Update product               | Yes           |
| GET    | `/products-by-slug/{slug}`        | Get product by slug          | No            |
| GET    | `/popular-products`               | Get popular products         | No            |
| POST   | `/products/{id}/increment-views`  | Increment product views      | No            |

#### Product Translation Endpoints

| Method | Endpoint                      | Description                      | Auth Required |
| ------ | ----------------------------- | -------------------------------- | ------------- |
| GET    | `/products/{id}/translations` | Get product translations         | No            |
| POST   | `/products/{id}/translations` | Create product translation       | Yes           |
| DELETE | `/products/{id}/translations` | Delete product translation       | Yes           |
| POST   | `/product-trans/{id}`         | Alternative translation endpoint | Yes           |

#### Product Image Endpoints

| Method | Endpoint                         | Description             | Auth Required |
| ------ | -------------------------------- | ----------------------- | ------------- |
| POST   | `/products/{id}/image`           | Add product image       | Yes           |
| GET    | `/products/{id}/images`          | Get product images      | No            |
| POST   | `/addproductimage/{productId}`   | Add product image (alt) | Yes           |
| GET    | `/get-product-image/{productId}` | Get product image       | No            |

### ⭐ Reviews & Feedback

#### Product Review Endpoints

| Method | Endpoint                          | Description               | Auth Required |
| ------ | --------------------------------- | ------------------------- | ------------- |
| GET    | `/products/{id}/reviews`          | Get product reviews       | No            |
| GET    | `/public-reviews/{id}`            | Get public reviews        | No            |
| GET    | `/reviews`                        | List all reviews          | No            |
| POST   | `/reviews/{id}`                   | Create review             | Yes           |
| GET    | `/reviews/{id}`                   | Get review by ID          | No            |
| POST   | `/product/{id}/review/{reviewid}` | Update review             | Yes           |
| POST   | `/review/translation`             | Create review translation | Yes           |

#### Review Image Endpoints

| Method | Endpoint              | Description                 | Auth Required |
| ------ | --------------------- | --------------------------- | ------------- |
| POST   | `/productimages`      | Upload review images        | Yes           |
| POST   | `/productimages/s3`   | Register S3 uploaded images | Yes           |
| GET    | `/productimages/{id}` | Get product images          | No            |
| POST   | `/imageremove/{id}`   | Remove image                | Yes           |
| POST   | `/default-image/{id}` | Make default image          | Yes           |
| POST   | `/s3/presign`         | Presign S3 PUT URL          | Yes           |
| POST   | `/s3/presign-get`     | Presign S3 GET URL          | Yes           |

#### Feedback Endpoints

| Method | Endpoint            | Description             | Auth Required |
| ------ | ------------------- | ----------------------- | ------------- |
| GET    | `/feedback`         | Get all feedback        | No            |
| POST   | `/feedback/`        | Create feedback         | Yes           |
| GET    | `/feedback/{id}`    | Get feedback by ID      | Yes           |
| PUT    | `/feedback/{id}`    | Update feedback         | Yes           |
| DELETE | `/feedback/{id}`    | Delete feedback         | Yes           |
| GET    | `/v1/feedback`      | Get all feedback (V1)   | No            |
| POST   | `/v1/feedback/`     | Create feedback (V1)    | Yes           |
| GET    | `/v1/feedback/{id}` | Get feedback by ID (V1) | Yes           |
| PUT    | `/v1/feedback/{id}` | Update feedback (V1)    | Yes           |
| DELETE | `/v1/feedback/{id}` | Delete feedback (V1)    | Yes           |

### 📂 Categories

#### Category Endpoints

| Method | Endpoint                | Description            | Auth Required |
| ------ | ----------------------- | ---------------------- | ------------- |
| GET    | `/categories`           | List all categories    | No            |
| POST   | `/categories`           | Create new category    | Yes           |
| GET    | `/categories/{id}`      | Get category by ID     | No            |
| PUT    | `/categories/{id}`      | Update category        | Yes           |
| DELETE | `/categories/{id}`      | Delete category        | Yes           |
| GET    | `/wide-categories`      | Get wide categories    | No            |
| PUT    | `/category-status/{id}` | Update category status | Yes           |

#### Category Translation Endpoints

| Method | Endpoint                     | Description                 | Auth Required |
| ------ | ---------------------------- | --------------------------- | ------------- |
| POST   | `/category-translation`      | Create category translation | Yes           |
| GET    | `/category-translation/{id}` | Get category translations   | No            |
| PUT    | `/category-translation/{id}` | Update category translation | Yes           |

#### Category-Brand Relations

| Method | Endpoint           | Description                    | Auth Required |
| ------ | ------------------ | ------------------------------ | ------------- |
| GET    | `/category-brands` | Get category-brand relations   | No            |
| POST   | `/category-brands` | Create category-brand relation | Yes           |

### 🏷️ Brands

#### Brand Endpoints

| Method | Endpoint             | Description         | Auth Required |
| ------ | -------------------- | ------------------- | ------------- |
| GET    | `/brands`            | List all brands     | No            |
| POST   | `/brands`            | Create new brand    | Yes           |
| GET    | `/brands/{id}`       | Get brand by ID     | No            |
| PUT    | `/brands/{id}`       | Update brand        | Yes           |
| DELETE | `/brands/{id}`       | Delete brand        | Yes           |
| GET    | `/wide-brands`       | Get wide brands     | No            |
| GET    | `/public-brands`     | Get public brands   | No            |
| POST   | `/brand-status/{id}` | Update brand status | Yes           |

#### Brand Translation Endpoints

| Method | Endpoint                  | Description              | Auth Required |
| ------ | ------------------------- | ------------------------ | ------------- |
| POST   | `/brand-translation`      | Create brand translation | Yes           |
| GET    | `/brand-translation/{id}` | Get brand translations   | No            |
| PUT    | `/brand-translation/{id}` | Update brand translation | Yes           |

### 📋 Specifications

#### Specification Endpoints

| Method | Endpoint                   | Description                      | Auth Required |
| ------ | -------------------------- | -------------------------------- | ------------- |
| POST   | `/specifications`          | Create specification             | Yes           |
| POST   | `/specifications/bulk`     | Bulk upsert specifications       | Yes           |
| GET    | `/specifications/{id}`     | Get specification by ID          | No            |
| PUT    | `/specifications/{id}`     | Update specification             | Yes           |
| DELETE | `/specifications/{id}`     | Delete specification             | Yes           |
| GET    | `/get-specifications/{id}` | Get specifications by product ID | No            |
| GET    | `/get-public-spec/{id}`    | Get public specification         | No            |
| GET    | `/specificationsearch`     | Search specifications            | No            |

#### Specification Translation Endpoints

| Method | Endpoint                   | Description                      | Auth Required |
| ------ | -------------------------- | -------------------------------- | ------------- |
| POST   | `/spec_translation`        | Create specification translation | Yes           |
| GET    | `/spec_translation/{id}`   | Get specification translations   | No            |
| PUT    | `/spec_translation/values` | Update translation values        | Yes           |
| POST   | `/spec_translation/values` | Update translation values (alt)  | Yes           |

#### Specification Key Endpoints

| Method | Endpoint           | Description                        | Auth Required |
| ------ | ------------------ | ---------------------------------- | ------------- |
| GET    | `/speckey`         | Get all specification keys         | No            |
| POST   | `/speckey`         | Create or update specification key | Yes           |
| GET    | `/speckey/{id}`    | Get specification key by ID        | No            |
| POST   | `/specremove/{id}` | Delete specification key           | Yes           |

#### Specification Key Translation Endpoints

| Method | Endpoint               | Description                   | Auth Required |
| ------ | ---------------------- | ----------------------------- | ------------- |
| GET    | `/speckey-translation` | Get all spec key translations | No            |
| POST   | `/speckey-translation` | Create spec key translation   | Yes           |

### 🔧 Form Generator

#### Form Generator Endpoints

| Method | Endpoint               | Description                 | Auth Required |
| ------ | ---------------------- | --------------------------- | ------------- |
| POST   | `/formgenerator/`      | Create form generator       | Yes           |
| GET    | `/formgenerator/{id}`  | Get form generator by ID    | No            |
| PUT    | `/formgenerator/{id}`  | Update form generator       | Yes           |
| GET    | `/category-specs/{id}` | Get category specifications | No            |

### 📊 Example API Requests

#### Register a New User

```bash
curl -X POST http://localhost:8080/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "alice",
    "email": "alice@example.com",
    "password": "password123"
  }'
```

#### Login

```bash
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "alice@example.com",
    "password": "password123"
  }'
```

#### Get All Products

```bash
curl -X GET http://localhost:8080/products
```

#### Get Product by ID

```bash
curl -X GET http://localhost:8080/products/1
```

#### Create a Product (Authenticated)

```bash
curl -X POST http://localhost:8080/products \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{
    "name": "iPhone 15",
    "description": "Latest iPhone model",
    "category_id": 1,
    "brand_id": 1
  }'
```

#### Get Product Reviews

```bash
curl -X GET http://localhost:8080/products/1/reviews
```

#### Get Categories

```bash
curl -X GET http://localhost:8080/categories
```

#### Get Brands

```bash
curl -X GET http://localhost:8080/brands
```

### 📌 Notes

- **Base URL**: All endpoints are relative to `http://localhost:8080` (default)
- **Authentication**: Endpoints marked "Yes" require a valid JWT token in the `Authorization` header
- **Content-Type**: Most POST/PUT requests require `Content-Type: application/json`
- **Response Format**: All responses are in JSON format
- **Pagination**: Many list endpoints support `limit`, `offset`, or `page` query parameters
- **Search**: Search endpoints typically accept a `q` or `search` query parameter
- **Localization**: Many endpoints support a `locale` query parameter (e.g., `en`, `de`, `fr`)

---

## 6. Development Tools & Troubleshooting

### 6.1. Air Hot Reloading Issues

**Problem: `cmd\main.go: directory not found`**

_Solution:_ The `.air.toml` file has been configured for this project structure. Make sure you're running Air from the project root:

```bash
# Correct project structure uses cmd/app/main.go, not cmd/main.go
air

# If still failing, try rebuilding the air config:
air init
```

**Problem: `'air' is not recognized as an internal or external command`**

_Solution:_ Air is not installed or not in PATH:

```bash
# Install Air
go install github.com/cosmtrek/air@latest

# Add Go bin to PATH (Windows)
set PATH=%PATH%;%USERPROFILE%\go\bin

# Add Go bin to PATH (Linux/Mac)
export PATH=$PATH:$(go env GOPATH)/bin

# Or run with full path
$(go env GOPATH)/bin/air
```

**Problem: Build errors with Air**

_Solution:_ Check the `.air.toml` configuration and build logs:

```bash
# Check build errors
cat build-errors.log

# Manual build test
go build -o ./tmp/main.exe ./cmd/app/main.go

# Clean tmp directory
rm -rf tmp
mkdir tmp
```

**Problem: Compilation errors when using Air**

_Solution:_ The project may have entity import conflicts. Use the working migration and seeding tools:

```bash
# Working tools that compile successfully:
go run ./cmd/migrate/main.go -migrate
go run ./cmd/migrate/main.go -seed
go run ./cmd/seedtest/main.go

# For development, focus on these working components:
go run ./cmd/migrate/main.go -create-db  # Create DB and migrate
go run ./cmd/migrate/main.go -fresh-seed # Fresh setup with data
go run ./cmd/seedtest/main.go           # Verify seeded data
```

**Problem: `'tmp\main.exe' is not recognized as an internal or external command`**

_Solution:_ Air build is failing or path issues on Windows:

```bash
# Step 1: Check if build succeeds manually
go build -o ./tmp/main.exe ./cmd/app/main.go

# Step 2: If build fails, check the errors
mkdir tmp
go build -o ./tmp/main.exe ./cmd/app/main.go

# Step 3: Fix .air.toml for Windows (ensure .exe extension)
# Update your .air.toml file:
# bin = "./tmp/main.exe"
# cmd = "go build -o ./tmp/main.exe ./cmd/app/main.go"

# Step 4: Alternative - Use the working migration tools with Air
# Change .air.toml to build migration tool instead:
# cmd = "go build -o ./tmp/migrate.exe ./cmd/migrate/main.go"
# bin = "./tmp/migrate.exe"
# args_bin = ["-migrate"]
```

**Problem: Air shows compilation errors**

_Solution:_ The main app has known entity conflicts. Use Air with the working migration tool:

```bash
# Option 1: Use migration tool with Air (modify .air.toml)
# Change these lines in .air.toml:
# cmd = "go build -o ./tmp/migrate.exe ./cmd/migrate/main.go"
# bin = "./tmp/migrate.exe"
# args_bin = ["-h"]  # This will show help

# Option 2: Skip Air and use direct commands
go run ./cmd/migrate/main.go -create-db
go run ./cmd/migrate/main.go -seed
go run ./cmd/seedtest/main.go

# Option 3: Check specific build errors
go build -o ./tmp/main.exe ./cmd/app/main.go 2>&1 | tee build-errors.log
```

### 6.2. Go Dependencies Issues

**Problem: `go: module not found`**

_Solution:_ Initialize and download dependencies:

```bash
# Initialize Go module (if not done)
go mod init kossti

# Download dependencies
go mod download

# Clean module cache if corrupted
go clean -modcache
go mod download
```

**Problem: Version conflicts**

_Solution:_ Update dependencies:

```bash
# Update all dependencies
go get -u ./...

# Tidy up go.mod and go.sum
go mod tidy

# Verify dependencies
go mod verify
```

**Problem: Entity import conflicts**

_Solution:_ The project has conflicting entity packages causing type mismatches:

```bash
# Error symptoms:
# cannot use user (variable of type *entity.User) as *entities.User value
# cannot use createdUser (variable of type *entities.User) as *entity.User value
# user.UpdateProfile undefined (type *entities.User has no field or method UpdateProfile)

# Quick diagnosis - check import conflicts:
grep -r "entity\." internal/usecase/
grep -r "entities\." internal/usecase/

# Temporary workaround - use the working migration tools:
go run ./cmd/migrate/main.go -create-db
go run ./cmd/migrate/main.go -seed
go run ./cmd/seedtest/main.go

# Long-term fix - standardize entity imports:
# Option 1: Use only 'entities' package everywhere
# Option 2: Use only 'entity' package everywhere
# Option 3: Create type aliases for compatibility
```

**Entity Conflict Troubleshooting Steps:**

```bash
# Step 1: Identify conflicting imports
find internal/ -name "*.go" -exec grep -l "entity\." {} \;
find internal/ -name "*.go" -exec grep -l "entities\." {} \;

# Step 2: Check which entity package is primary
ls internal/domain/entities/
ls internal/domain/entity/

# Step 3: Temporarily focus on working tools
# These compile successfully:
go run ./cmd/migrate/main.go -h          # Shows migration options
go run ./cmd/migrate/main.go -create-db  # Database setup
go run ./cmd/migrate/main.go -seed       # Data seeding
go run ./cmd/seedtest/main.go           # Verification

# Step 4: Use Air with working migration tools
air -c .air-migrate.toml  # Uses migration tools instead of main app
```

**Specific Compilation Errors Fix:**

If you see these exact errors:

```
cannot use user (variable of type *entity.User) as *entities.User value
cannot use createdUser (variable of type *entities.User) as *entity.User value
user.UpdateProfile undefined (type *entities.User has no field or method UpdateProfile)
```

**Root Cause:** The project has two entity packages:

- `internal/domain/entity/` (singular)
- `internal/domain/entities/` (plural)

**Immediate Solutions:**

```bash
# Option A: Focus on working components (recommended)
air -c .air-migrate.toml                # Use migration tools with hot reload
go run ./cmd/migrate/main.go -create-db  # Database operations work
go run ./cmd/migrate/main.go -seed       # Seeding works perfectly

# Option B: Check which entity package is actually used
ls internal/domain/entity*/             # See which directories exist
grep -r "type.*User struct" internal/   # Find where User is defined

# Option C: Quick fix for development
# Edit internal/usecase/user/user.go and change imports:
# FROM: "kossti/internal/domain/entity"
# TO:   "kossti/internal/domain/entities"
```

**Long-term Resolution Strategy:**

```bash
# 1. Identify the primary entity package
find internal/ -name "*.go" -exec grep -l "package entities" {} \;
find internal/ -name "*.go" -exec grep -l "package entity" {} \;

# 2. Standardize all imports to use the same package
# Either use 'entities' everywhere or 'entity' everywhere

# 3. Update repository interfaces to match
# Make sure all interfaces use the same entity types

# 4. Test with migration tools (these work regardless)
go run ./cmd/migrate/main.go -migrate
```

### 6.3. PostgreSQL Connection Issues

**Problem: `connection refused` or `database does not exist`**

_Solution:_ Use the automatic database creation:

```bash
# Create database if not exists, then migrate
go run ./cmd/migrate/main.go -create-db

# Or manually create database
createdb -U root kossti
```

**Problem: Docker PostgreSQL issues**

_Solution:_ Reset PostgreSQL data:

```bash
# Stop containers
docker-compose down

# Remove corrupted data
rm -rf postgres-data

# Restart with fresh data
docker-compose up postgres
```

---

## 7. Troubleshooting

- **Database connection failed:**
  - Make sure Postgres is running and your password is correct.
  - See the password setup section below.
- **psql not found (Windows):**
  - Add the Postgres `bin` directory to your PATH (see instructions below).
- **.env not loaded:**
  - Make sure you have `github.com/joho/godotenv` and your `.env` file is in the project root.

---

## 7. How to Set or Change the Password for the "postgres" User in PostgreSQL

1. Open a terminal or command prompt.
2. Connect to PostgreSQL as the "postgres" user:
   ```sh
   psql -U postgres
   ```
   (If prompted for a password and you don't know it, try running as the system user that installed Postgres, or use `sudo -u postgres psql` on Linux.)
3. Set a new password:
   ```sql
   ALTER USER postgres WITH PASSWORD 'yournewpassword';
   ```
4. Update your `.env` file:
   ```env
   DATABASE_URL=postgres://postgres:yournewpassword@localhost:5432/authdb?sslmode=disable
   ```
5. Restart your Go app.

**Note for Windows users:**
If you see an error like `'psql' is not recognized as an internal or external command`, add the PostgreSQL `bin` directory to your system PATH:

1. Find where PostgreSQL is installed (e.g., `C:\Program Files\PostgreSQL\17\bin`).
2. Open the Start menu and search for "Environment Variables".
3. Click "Edit the system environment variables" > "Environment Variables...".
4. Under "System variables", find and select the `Path` variable, then click "Edit...".
5. Click "New" and add the path to your PostgreSQL `bin` directory.
6. Click OK to save and restart your terminal.

---

## 8. Clean Architecture Benefits (Recap)

- Easy to test and maintain
- Swap out frameworks or databases with minimal changes
- Business logic is protected from outside changes

---

## 9. For You

- Ask questions about any layer or file!
- Try changing the database or adding a new feature (like password reset)
- Practice writing tests for the use cases

---

## 9. How to Migrate the Database Schema (AutoMigrate)

You can automatically create or update your database tables to match your Go structs using GORM's AutoMigrate feature.

### 1. Install GORM and the Postgres driver

```sh
go get gorm.io/gorm
go get gorm.io/driver/postgres
```

### 2. Update your main.go

- Import GORM and your User entity.
- Use the following code to connect and migrate:

```go
import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "kossti/internal/domain/entity"
    // ...
)

func main() {
    // ...
    db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
    if err != nil {
        log.Fatalf("GORM database connection failed: %v", err)
    }
    if err := db.AutoMigrate(&entity.User{}); err != nil {
        log.Fatalf("AutoMigrate failed: %v", err)
    }
    log.Println("Database migration complete!")
}
```

### 3. Run your app

```sh
go run ./cmd/main.go
```

- This will create or update the `users` table in your database to match your User struct.

### 3. Run your app

```sh
go run ./cmd/app/main.go
```

---

## 10. Advanced Database Migration System

For larger projects with multiple models, we have a comprehensive migration system that handles all database operations.

### 10.1. Migration Manager Features

- ✅ **Batch Migration**: Migrates all 24 models at once
- ✅ **Progress Logging**: See which model is being migrated
- ✅ **Foreign Key Management**: Automatically handles relationships
- ✅ **Index Creation**: Creates database indexes for performance
- ✅ **Production Safe**: Won't break existing data
- ✅ **CLI Tool**: Easy command-line interface

### 10.2. Using the Migration CLI Tool

#### 1. Safe Migration (Production Recommended)

```bash
# Migrates all tables, adds foreign keys, and creates indexes
go run cmd/migrate/main.go -migrate
```

#### 2. Database Creation + Migration (Recommended for New Projects)

```bash
# Creates database if not exists, then runs complete migration
go run cmd/migrate/main.go -create-db
```

#### 3. Fresh Setup (Development Only)

```bash
# ⚠️ DANGEROUS: Drops ALL tables (including orphaned tables) and recreates them
go run cmd/migrate/main.go -fresh

# ⚠️ DANGEROUS: Fresh setup with sample data
go run cmd/migrate/main.go -fresh-seed
```

**Note:** The `-fresh` command now uses PostgreSQL's system catalog to find and drop ALL tables in the public schema, including any orphaned tables not in the model list (e.g., legacy Laravel tables like `jobs`, `permissions`, `roles`, etc.). This ensures a completely clean slate.

#### 4. Individual Operations

```bash
# Drop all tables (testing/cleanup) - IMPROVED: Drops ALL tables including orphaned ones
go run cmd/migrate/main.go -drop

# Add foreign keys only
go run cmd/migrate/main.go -fk

# Create indexes only
go run cmd/migrate/main.go -indexes

# Run seeders only (add sample data)
go run cmd/migrate/main.go -seed
```

**Important Note on -drop and -fresh flags:**  
These commands now query PostgreSQL's system catalog (`pg_tables`) to find and drop ALL tables in the public schema, not just the ones defined in your Go models. This ensures complete cleanup of:

- All current model tables (24 tables)
- Legacy/orphaned tables (e.g., `jobs`, `permissions`, `roles`, `model_has_permissions`, etc. from other frameworks)
- Any other tables you may have created manually

This is particularly useful when migrating from other frameworks or cleaning up after testing different schemas.

#### 5. Custom Database Connection

```bash
# Use custom database connection
go run cmd/migrate/main.go -migrate -dsn "host=prod-db user=prod password=secret dbname=production port=5432 sslmode=require"

# Create new database with custom connection
go run cmd/migrate/main.go -create-db -dsn "host=localhost user=root password=root dbname=new_project port=5428 sslmode=disable TimeZone=UTC"
```

### 10.3. Migration CLI Options

| Flag          | Description                             | Safety Level     |
| ------------- | --------------------------------------- | ---------------- |
| `-migrate`    | Run safe migration (recommended)        | ✅ **Safe**      |
| `-create-db`  | Create database if not exists + migrate | ✅ **Safe**      |
| `-fresh`      | Drop all + recreate (development)       | 🚨 **Dangerous** |
| `-fresh-seed` | Drop, migrate, and seed (development)   | 🚨 **Dangerous** |
| `-seed`       | Run all database seeders                | ✅ **Safe**      |
| `-drop`       | Drop all tables                         | 🚨 **Dangerous** |
| `-fk`         | Add foreign keys only                   | ✅ **Safe**      |
| `-indexes`    | Create indexes only                     | ✅ **Safe**      |
| `-dsn`        | Custom database connection string       | ℹ️ **Config**    |

### 10.4. Programmatic Usage (In Your App)

```go
package main

import (
    "kossti/internal/infrastructure/database"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "log"
)

func setupDatabase() {
    // Connect to database
    db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
    if err != nil {
        log.Fatal("Database connection failed:", err)
    }

    // Create migration manager
    manager := database.NewMigrationManager(db)

    // Option 1: Complete setup (migrate + foreign keys + indexes)
    if err := manager.Setup(); err != nil {
        log.Fatal("Migration failed:", err)
    }

    // Option 2: Individual operations
    // if err := manager.MigrateAll(); err != nil {
    //     log.Fatal("Migration failed:", err)
    // }
    // if err := manager.AddForeignKeys(); err != nil {
    //     log.Fatal("Foreign keys failed:", err)
    // }
    // if err := manager.CreateIndexes(); err != nil {
    //     log.Fatal("Indexes failed:", err)
    // }

    log.Println("✅ Database setup completed!")
}
```

### 10.5. What Gets Migrated

The migration system handles **24 database models** organized into:

#### Core System Tables (2 models)

- `users`, `password_reset_tokens`

#### Product System (10 models)

- `products`, `product_reviews`, `categories`, `brands`
- `brand_categories`, `productables`, `images`, `tags`
- `feedback`, `form_generators`

#### Translation Tables (8 models)

- `product_translations`, `category_translations`, `brand_translations`
- `product_review_translations`, `feedback_translations`
- `comment_translations`, `specification_translations`, `specification_key_translations`

#### Content System (8 models)

- `comments`, `specifications`, `specification_keys`
- And their translation tables

### 10.6. Migration Best Practices

#### For Development

```bash
# First time setup
go run cmd/migrate/main.go -fresh

# Regular development
go run cmd/migrate/main.go -migrate
```

#### For Production

```bash
# Always use safe migration
go run cmd/migrate/main.go -migrate

# Never use -fresh or -drop in production!
```

#### For Testing

```bash
# Clean slate for tests
go run cmd/migrate/main.go -drop
go run cmd/migrate/main.go -migrate
```

### Migration vs AutoMigrate Comparison

| Feature           | AutoMigrate         | Migration Manager     |
| ----------------- | ------------------- | --------------------- |
| Model Count       | Manual (1-5 models) | Automatic (24 models) |
| Progress Tracking | ❌ No               | ✅ Yes                |
| Foreign Keys      | ❌ Manual           | ✅ Automatic          |
| Indexes           | ❌ Manual           | ✅ Automatic          |
| Error Handling    | ❌ Basic            | ✅ Advanced           |
| Production Ready  | ⚠️ Limited          | ✅ Yes                |

### 10.7. Troubleshooting Migrations

#### Common Issues

**"Table already exists" error:**

```bash
# This is normal - migration will skip existing tables
go run cmd/migrate/main.go -migrate
```

**Foreign key constraint errors:**

```bash
# Run foreign keys separately
go run cmd/migrate/main.go -fk
```

**Permission denied:**

```bash
# Make sure your database user has CREATE privileges
GRANT CREATE ON DATABASE gocrit_db TO your_user;
```

**Connection refused:**

```bash
# Check your database is running and connection string is correct
go run cmd/migrate/main.go -migrate -dsn "your_connection_string"
```

**Orphaned or legacy tables not being dropped:**

```bash
# The -fresh and -drop commands now handle ALL tables automatically
# They query pg_tables to find every table in the public schema

# If you still have issues, you can manually check what tables exist:
PGPASSWORD=root psql -h localhost -p 5428 -U root -d kossti -c "\dt"

# Then use -fresh to drop everything:
go run cmd/migrate/main.go -fresh
```

**Tables from previous Laravel migration:**

If you have legacy tables from a Laravel application (like `jobs`, `failed_jobs`, `permissions`, `roles`, `model_has_permissions`, `model_has_roles`, `role_has_permissions`), the improved `-fresh` and `-drop` commands will now automatically detect and remove them using PostgreSQL's system catalog.

### Migration Output Example

```
🛡️ Running safe migration...
2025/07/22 10:30:45 Starting database migration...
2025/07/22 10:30:45 Migrating model 1/36: *models.UserModel
2025/07/22 10:30:45 Migrating model 2/36: *models.PasswordResetTokenModel
2025/07/22 10:30:45 Migrating model 3/36: *models.SessionModel
...
2025/07/22 10:30:47 All migrations completed successfully!
2025/07/22 10:30:47 Adding foreign key constraints...
2025/07/22 10:30:47 Foreign key constraints added!
2025/07/22 10:30:47 Creating additional indexes...
2025/07/22 10:30:47 Additional indexes created!
✅ Migration completed!
```

### Quick Reference Commands

**Most Common Commands:**

```bash
# 🎯 NEW PROJECT: Create database and run migration
go run cmd/migrate/main.go -create-db

# 🔄 EXISTING PROJECT: Run safe migration only
go run cmd/migrate/main.go -migrate

# 🌱 DEVELOPMENT: Fresh setup with sample data
go run cmd/migrate/main.go -fresh-seed

# 🧪 TESTING: Clean database for tests
go run cmd/migrate/main.go -drop && go run cmd/migrate/main.go -migrate
```

**Custom Database Examples:**

```bash
# Create new database with custom name
go run cmd/migrate/main.go -create-db -dsn "host=localhost user=root password=root dbname=new_project port=5428 sslmode=disable TimeZone=UTC"

# Production migration with SSL
go run cmd/migrate/main.go -migrate -dsn "host=prod-server user=prod_user password=secure_pass dbname=prod_db port=5432 sslmode=require"

# Test different database locally
go run cmd/migrate/main.go -create-db -dsn "host=localhost user=root password=root dbname=test_db port=5428 sslmode=disable TimeZone=UTC"
```

---

## 11. Database Seeding System 🌱

Our Clean Architecture includes a comprehensive database seeding system that populates your database with realistic test data following Laravel-style seeding patterns.

### 11.1. Quick Start

```bash
# Run fresh setup with seeding (development only)
go run ./cmd/migrate/main.go -fresh-seed

# Run seeders only (existing database)
go run ./cmd/migrate/main.go -seed

# Verify seeded data
go run ./cmd/seedtest/main.go
```

### 11.2. What Gets Seeded

The seeding system includes **400+ records** organized across multiple tables:

| **Seeder**           | **Records**       | **Description**                                              |
| -------------------- | ----------------- | ------------------------------------------------------------ |
| **Categories**       | 128 records       | Product categories (Electronics, Footwear, Automotive, etc.) |
| **Brands**           | 152 records       | Global brands (Apple, Nike, BMW, Google, etc.)               |
| **Brand-Categories** | 415 relationships | Smart brand-category mappings                                |

### 11.3. Seeded Data Examples

**📂 Categories Include:**

- Electronics, Mobile Phones, Laptops
- Footwear, Clothing, Sports Equipment
- Automotive, Luxury Cars, Electric Vehicles
- Beauty Products, Personal Care, Skincare
- Food & Beverages, Restaurants, Coffee

**🏷️ Brands Include:**

- **Tech**: Apple, Google, Microsoft, Samsung, Intel
- **Automotive**: BMW, Mercedes-Benz, Tesla, Toyota, Audi
- **Fashion**: Nike, Adidas, Gucci, Louis Vuitton, Zara
- **F&B**: Coca-Cola, Starbucks, McDonald's, Nestlé

**🔗 Smart Relationships:**

- Apple → Electronics, Mobile Phones, Laptops, Technology
- Nike → Footwear, Clothing, Sports Equipment
- BMW → Automotive, Luxury Cars, Motorcycles
- Starbucks → Food & Beverages, Coffee, Restaurants

### 11.4. Seeding CLI Options

| **Flag**      | **Description**         | **Use Case**                     |
| ------------- | ----------------------- | -------------------------------- |
| `-seed`       | Run all seeders         | Adding data to existing database |
| `-fresh-seed` | Drop, migrate, and seed | Complete fresh setup (dev only)  |

### 11.5. Architecture Overview

The seeding system follows Clean Architecture principles:

```
internal/infrastructure/database/seeders/
├── seeder.go              # Core seeding framework & interfaces
├── category_seeder.go     # Category data seeding
├── brand_seeder.go        # Brand data seeding
└── brand_category_seeder.go # Relationship seeding
```

**Key Components:**

- **SeederManager**: Orchestrates all seeders with progress tracking
- **Seeder Interface**: Common contract for all seeders
- **BaseSeeder**: Shared functionality and naming
- **Helper Functions**: Duplicate detection, slug generation, relationship creation

### 11.6. Programmatic Usage

```go
package main

import (
    "kossti/internal/infrastructure/database/seeders"
    "gorm.io/gorm"
)

func setupSeeders(db *gorm.DB) {
    // Setup all seeders
    manager := seeders.SetupAllSeeders(db)

    // Run all seeders
    if err := manager.RunAll(); err != nil {
        log.Fatal("Seeding failed:", err)
    }

    // Or run specific seeders
    manager.AddSeeder(seeders.NewCategorySeeder())
    manager.AddSeeder(seeders.NewBrandSeeder())

    if err := manager.RunAll(); err != nil {
        log.Fatal("Seeding failed:", err)
    }
}
```

### 11.7. Creating Custom Seeders

```go
package seeders

import (
    "kossti/internal/domain/entities"
    "kossti/internal/infrastructure/database/models"
    "gorm.io/gorm"
)

// CustomSeeder example
type ProductSeeder struct {
    BaseSeeder
}

func NewProductSeeder() *ProductSeeder {
    return &ProductSeeder{
        BaseSeeder: BaseSeeder{name: "Products"},
    }
}

func (ps *ProductSeeder) Seed(db *gorm.DB) error {
    products := []struct {
        Name     string
        Category string
        Brand    string
    }{
        {"iPhone 15", "Mobile Phones", "Apple"},
        {"Air Jordan 1", "Footwear", "Nike"},
    }

    for _, product := range products {
        // Get related entities
        category, err := CreateOrFindCategory(db, product.Category, GenerateSlug(product.Category))
        if err != nil {
            return err
        }

        brand, err := CreateOrFindBrand(db, product.Brand, GenerateSlug(product.Brand))
        if err != nil {
            return err
        }

        // Create product entity
        productEntity := &entities.Product{
            Name:       product.Name,
            CategoryID: &category.ID,
            BrandID:    &brand.ID,
        }

        // Convert and save
        var productModel models.ProductModel
        productModel.FromEntity(productEntity)

        if err := db.Create(&productModel).Error; err != nil {
            return err
        }
    }

    return nil
}
```

### Seeder Output Example

```
🌱 Starting database seeding...
   🔄 Running Categories seeder...
   ✅ Categories seeder completed successfully
   🔄 Running Brands seeder...
   ✅ Brands seeder completed successfully
   🔄 Running Brand Categories seeder...
   ✅ Brand Categories seeder completed successfully
🎉 All seeders completed successfully!
```

### Best Practices

**✅ Do:**

- Use seeders for development and testing environments
- Run verification after seeding
- Keep seeder data realistic and diverse
- Follow Clean Architecture patterns in custom seeders

**❌ Don't:**

- Run seeders in production without careful consideration
- Modify core seeder files directly
- Skip duplicate detection in custom seeders
- Hardcode IDs in seeder relationships

---

## 12. CLI Tools Available

### 12.1. Main Application Server

```bash
# Start the HTTP API server (default port: 8080)
go run ./cmd/app/main.go

# Build the application
go build -o bin/kossti-server ./cmd/app
```

**Features:**

- ✅ User registration and login endpoints
- ✅ RESTful API for user management
- ✅ Automatic database migration on startup
- ✅ Health check endpoint
- ✅ Graceful shutdown
- ✅ Port conflict detection

**API Endpoints:**

**Authentication & Users:**

- `POST /register` - User registration
- `POST /login` - User login
- `GET /api/users` - List users (paginated)
- `GET /api/users/all` - Get all users
- `GET /api/users/search` - Search users
- `GET /api/users/count` - Get user count
- `GET /api/users/{id}` - Get user by ID

**Products:**

- `GET /api/products` - List products with pagination, search & filtering
- `POST /api/products` - Create a new product
- `GET /api/products/{id}` - Get product by ID
- `PATCH /api/products/{id}` - Update product
- `GET /api/products-by-slug/{slug}` - Get product by slug
- `GET /api/popular-products` - Get popular products (sorted by views)
- `POST /api/products/{id}/increment-views` - Increment product view count

**Product Images:**

- `POST /api/addproductimage/{productId}` - Upload image for product
- `GET /api/get-product-image/{productId}` - Get all images for product
- `POST /api/products/{product}/image` - Upload image (alternative endpoint)
- `GET /api/products/{product}/images` - Get images (alternative endpoint)

**Product Translations:**

- `POST /api/product-trans/{id}` - Create product translation

**System:**

- `GET /health` - Health check

**Query Parameters for Product Endpoints:**

For `GET /api/products`:

- `limit` - Number of products to return (default: 10)
- `offset` - Number of products to skip (default: 0)
- `search` - Search term for product name or description
- `category_id` - Filter by category ID
- `brand_id` - Filter by brand ID

For `GET /api/popular-products`:

- `limit` - Number of products to return (default: 10)

**Example Requests:**

```bash
# List all products
curl -X GET "http://localhost:8080/api/products"

# Search products
curl -X GET "http://localhost:8080/api/products?search=iphone&limit=5"

# Get products by category
curl -X GET "http://localhost:8080/api/products?category_id=1&limit=10"

# Get popular products
curl -X GET "http://localhost:8080/api/popular-products?limit=5"

# Create product
curl -X POST "http://localhost:8080/api/products" \
  -H "Content-Type: application/json" \
  -d '{"name":"iPhone 15","description":"Latest iPhone","price":999.99,"category_id":1}'

# Increment views
curl -X POST "http://localhost:8080/api/products/1/increment-views"

# Upload product image
curl -X POST "http://localhost:8080/api/addproductimage/1" \
  -F "image=@/path/to/image.jpg"

# Create translation
curl -X POST "http://localhost:8080/api/product-trans/1" \
  -H "Content-Type: application/json" \
  -d '{"locale":"bn","translated_name":"আইফোন ১৫","translated_description":"নতুন আইফোন"}'
```

### 12.2. Database Migration Tool

#### Database Creation Features

The migration system can automatically create the PostgreSQL database if it doesn't exist:

```bash
# Create database if not exists, then run migration
go run ./cmd/migrate/main.go -create-db
```

**How it works:**

1. **Connects to `postgres` default database** to check if target database exists
2. **Queries `pg_database`** system catalog to verify database existence
3. **Creates database** using `CREATE DATABASE` if it doesn't exist
4. **Runs normal migration** after database creation

**Database Name Detection:**

The system automatically extracts database name from your connection string:

```bash
# Key-value format
"host=localhost user=root password=root dbname=kossti port=5428"
# Extracts: kossti

# URL format
"postgres://root:root@localhost:5428/kossti?sslmode=disable"
# Extracts: kossti
```

**Example output:**

```
🔨 Creating database if not exists, then running migration...
🔄 Checking if database 'kossti' exists...
🔨 Creating database 'kossti'...
✅ Database 'kossti' created successfully!
🛡️ Starting database migration...
📋 Migrating model 1/60: *models.UserModel
📋 Migrating model 2/60: *models.PasswordResetTokenModel
...
✅ Database created and migration completed!
```

#### Migration Commands

```bash
# Safe migration (recommended for production)
go run ./cmd/migrate/main.go -migrate

# Create database if not exists, then migrate (safe)
go run ./cmd/migrate/main.go -create-db

# Fresh setup (development only - DANGEROUS)
go run ./cmd/migrate/main.go -fresh

# Fresh setup with seeding (development only - DANGEROUS)
go run ./cmd/migrate/main.go -fresh-seed

# Run seeders only
go run ./cmd/migrate/main.go -seed

# Show help and all options
go run ./cmd/migrate/main.go -h

# Build the migration tool
go build -o bin/kossti-migrate ./cmd/migrate
```

**Migration Options:**

- `-migrate` - Safe migration (production ready)
- `-create-db` - Create database if not exists, then migrate (safe)
- `-fresh` - Drop all + recreate (development only)
- `-fresh-seed` - Drop, migrate, and seed (development only)
- `-seed` - Run all database seeders
- `-drop` - Drop all tables (dangerous)
- `-fk` - Add foreign keys only
- `-indexes` - Create indexes only
- `-dsn` - Custom database connection string

### 12.3. Database Seeding Verification Tool

```bash
# Verify seeded data
go run ./cmd/seedtest/main.go

# Build the verification tool
go build -o bin/kossti-seedtest ./cmd/seedtest
```

**Features:**

- ✅ Shows count of seeded records
- ✅ Displays sample data from each table
- ✅ Verifies relationships (e.g., Apple's categories)
- ✅ Clean Architecture compliance testing

### 12.4. Build All Tools

```bash
# Using Makefile (Linux/Mac/WSL)
make build

# Or manually (Windows/any platform):
mkdir -p bin
go build -o bin/kossti-server ./cmd/app
go build -o bin/kossti-migrate ./cmd/migrate

# Then run the built binaries:
./bin/kossti-server
./bin/kossti-migrate -migrate
```

### Docker Usage (if configured)

```bash
# Run the application in Docker
docker-compose up -d

# Run migrations in Docker
docker-compose exec app ./bin/kossti-migrate -migrate
```

---

## 13. PostgreSQL Troubleshooting

### 13.1. Common PostgreSQL Docker Issues

#### Problem: "PostgreSQL Database directory appears to contain a database; Skipping initialization"

**Error Message:**

```
postgres-1  | PostgreSQL Database directory appears to contain a database; Skipping initialization
postgres-1  | FATAL: could not open directory "pg_notify": No such file or directory
postgres-1  | LOG: database system is shut down
```

**Cause:** Corrupted or incomplete PostgreSQL data in the `postgres-data` directory.

**Solution:**

1. **Stop the PostgreSQL container:**

   ```bash
   docker-compose down
   ```

2. **Remove the corrupted data directory:**

   ```bash
   rm -rf postgres-data
   ```

3. **Restart PostgreSQL with fresh initialization:**
   ```bash
   docker-compose up postgres
   ```

#### Problem: Database Does Not Exist

**Error:** `database "kossti" does not exist`

**Solution:** Use the automatic database creation feature:

```bash
# Create database if not exists, then migrate
go run ./cmd/migrate/main.go -create-db
```

**Manual Solution:**

```bash
# Connect to PostgreSQL as superuser
docker-compose exec postgres psql -U root -d postgres

# Create database manually
CREATE DATABASE kossti;

# Exit and run normal migration
go run ./cmd/migrate/main.go -migrate
```

#### Problem: Healthcheck Database Name Mismatch

**Error:** Healthcheck fails because it's checking the wrong database name.

**Solution:** Ensure the healthcheck database name matches your `POSTGRES_DB` environment variable in `docker-compose.yml`:

```yaml
environment:
  - POSTGRES_DB=kossti
healthcheck:
  test: ["CMD-SHELL", "pg_isready -U root -d kossti"] # Match this with POSTGRES_DB
```

#### Problem: Port Conflicts

**Error:** `bind: address already in use`

**Solution:** Change the host port in `docker-compose.yml`:

```yaml
ports:
  - "5429:5432" # Change 5428 to another available port
```

### 13.2. Connection Problems

**Troubleshooting Steps:**

1. **Check if PostgreSQL is running:**

   ```bash
   docker-compose ps
   ```

2. **Check PostgreSQL logs:**

   ```bash
   docker-compose logs postgres
   ```

3. **Test connection:**

   ```bash
   docker-compose exec postgres psql -U root -d kossti
   ```

4. **Verify port mapping:**
   ```bash
   netstat -an | grep 5428
   ```

### 13.3. Best Practices

- Always use `docker-compose down` before removing data directories
- Backup your database regularly if it contains important data
- Use specific PostgreSQL versions instead of `latest` for production
- Keep your `init-db` scripts in version control

---

## 🟢 Troubleshooting: Docker PostgreSQL Container Keeps Restarting

If your `postgres` container starts and then keeps restarting, try these steps:

1. **Stop and remove containers:**

   ```sh
   docker compose down
   ```

2. **Remove the data directory (this will delete all database data!):**

   ```sh
   rm -rf ./postgres-data
   ```

3. **(Optional) Check your `init-db` scripts** for errors or invalid SQL.

4. **Start fresh:**

   ```sh
   docker compose up -d
   ```

5. **Check logs for errors:**
   ```sh
   docker compose logs postgres
   ```

**Common causes:**

- Corrupted or permission-denied `postgres-data` directory
- Invalid SQL in `init-db` scripts
- Port already in use
- Environment variables missing or mismatched

**If you see:**

```
Error: Database is uninitialized and superuser password is not specified.
You must specify POSTGRES_PASSWORD to a non-empty value for the superuser.
```

**Solution:**

- Make sure your `docker-compose.yml` includes:
  ```yaml
  environment:
    - POSTGRES_USER=root
    - POSTGRES_PASSWORD=root
    - POSTGRES_DB=kossti
  ```
- Never use `POSTGRES_HOST_AUTH_METHOD=trust` in production.

**If you see:**

```
FATAL: could not open directory "pg_notify": No such file or directory
```

**Solution:**

- Remove the `postgres-data` directory and restart the container.

**If you see:**

```
database "kossti" does not exist
```

**Solution:**

- Use the automatic database creation feature:
  ```sh
  go run ./cmd/migrate/main.go -create-db
  ```

---

## 8. Clean Architecture Benefits (Recap)

✅ **Easy Testing**: Each layer can be tested independently  
✅ **Framework Independence**: Swap databases, web frameworks without breaking business logic  
✅ **Maintainable Code**: Clear separation of concerns makes code easier to understand  
✅ **Scalable Architecture**: Add new features without affecting existing functionality  
✅ **Database Agnostic**: Switch from PostgreSQL to MySQL with minimal changes

---

## 9. For You

- Ask questions about any layer or file!
- Try changing the database or adding a new feature (like password reset)
- Practice writing tests for the use cases
- Experiment with the seeding system to add your own data
- Build additional CLI tools using the existing patterns

---

Good luck and have fun learning Clean Architecture in Go!
