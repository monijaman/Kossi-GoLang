# Clean Architecture Authentication System in Go — Presentation Guide

Welcome! This guide is designed to help you present and explain your Go authentication project using Clean Architecture to students.

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

## 5. Setup & Run (Step-by-Step)

### 1. Install Go and PostgreSQL

- [Go download](https://go.dev/dl/)
- [PostgreSQL download](https://www.postgresql.org/download/)

### 2. Create a Database

```sh
createdb -U postgres authdb
```

### 3. Set Environment Variables

Create a `.env` file in the project root:

```env
DATABASE_URL=postgres://postgres:yourpassword@localhost:5432/authdb?sslmode=disable
JWT_SECRET=your_jwt_key
KAFKA_BROKER=localhost:9092
```

### 4. Install Dependencies

```sh
go mod init kossti
go get github.com/lib/pq
go get github.com/joho/godotenv
go get golang.org/x/crypto/bcrypt
```

### 5. Run the Applications

#### Main Application (API Server)

```sh
go run ./cmd/app/main.go
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

## 6. Troubleshooting

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

## How to Migrate the Database Schema (AutoMigrate)

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

### Migration Manager Features

- ✅ **Batch Migration**: Migrates all 60+ models at once
- ✅ **Progress Logging**: See which model is being migrated
- ✅ **Foreign Key Management**: Automatically handles relationships
- ✅ **Index Creation**: Creates database indexes for performance
- ✅ **Production Safe**: Won't break existing data
- ✅ **CLI Tool**: Easy command-line interface

### Using the Migration CLI Tool

#### 1. Safe Migration (Production Recommended)

```bash
# Migrates all tables, adds foreign keys, and creates indexes
go run cmd/migrate/main.go -migrate
```

#### 2. Fresh Setup (Development Only)

```bash
# ⚠️ DANGEROUS: Drops all tables and recreates them
go run cmd/migrate/main.go -fresh
```

#### 3. Individual Operations

```bash
# Drop all tables (testing/cleanup)
go run cmd/migrate/main.go -drop

# Add foreign keys only
go run cmd/migrate/main.go -fk

# Create indexes only
go run cmd/migrate/main.go -indexes
```

#### 4. Custom Database Connection

```bash
# Use custom database connection
go run cmd/migrate/main.go -migrate -dsn "host=prod-db user=prod password=secret dbname=production port=5432 sslmode=require"
```

### Migration CLI Options

| Flag          | Description                           | Safety Level     |
| ------------- | ------------------------------------- | ---------------- |
| `-migrate`    | Run safe migration (recommended)      | ✅ **Safe**      |
| `-fresh`      | Drop all + recreate (development)     | 🚨 **Dangerous** |
| `-fresh-seed` | Drop, migrate, and seed (development) | 🚨 **Dangerous** |
| `-seed`       | Run all database seeders              | ✅ **Safe**      |
| `-drop`       | Drop all tables                       | 🚨 **Dangerous** |
| `-fk`         | Add foreign keys only                 | ✅ **Safe**      |
| `-indexes`    | Create indexes only                   | ✅ **Safe**      |
| `-dsn`        | Custom database connection string     | ℹ️ **Config**    |

### Programmatic Usage (In Your App)

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

### What Gets Migrated

The migration system handles **60+ database models** organized into:

#### Core System Tables (15 models)

- `users`, `sessions`, `password_reset_tokens`
- `cache`, `cache_locks`, `jobs`, `job_batches`, `failed_jobs`
- `personal_access_tokens`, `history_logs`

#### Permission System (5 models)

- `permissions`, `roles`
- `model_has_permissions`, `model_has_roles`, `role_has_permissions`

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

### Migration Best Practices

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

| Feature           | AutoMigrate         | Migration Manager      |
| ----------------- | ------------------- | ---------------------- |
| Model Count       | Manual (1-5 models) | Automatic (60+ models) |
| Progress Tracking | ❌ No               | ✅ Yes                 |
| Foreign Keys      | ❌ Manual           | ✅ Automatic           |
| Indexes           | ❌ Manual           | ✅ Automatic           |
| Error Handling    | ❌ Basic            | ✅ Advanced            |
| Production Ready  | ⚠️ Limited          | ✅ Yes                 |

### Troubleshooting Migrations

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

### Migration Output Example

```
🛡️ Running safe migration...
2025/07/22 10:30:45 Starting database migration...
2025/07/22 10:30:45 Migrating model 1/60: *models.UserModel
2025/07/22 10:30:45 Migrating model 2/60: *models.PasswordResetTokenModel
2025/07/22 10:30:45 Migrating model 3/60: *models.SessionModel
...
2025/07/22 10:30:47 All migrations completed successfully!
2025/07/22 10:30:47 Adding foreign key constraints...
2025/07/22 10:30:47 Foreign key constraints added!
2025/07/22 10:30:47 Creating additional indexes...
2025/07/22 10:30:47 Additional indexes created!
✅ Migration completed!
```

---

## 10. Database Seeding System 🌱

Our Clean Architecture includes a comprehensive database seeding system that populates your database with realistic test data following Laravel-style seeding patterns.

### Quick Start

```bash
# Run fresh setup with seeding (development only)
go run cmd/migrate/main.go -fresh-seed

# Run seeders only (existing database)
go run cmd/migrate/main.go -seed

# Verify seeded data
go run cmd/seedtest/main.go
```

### What Gets Seeded

The seeding system includes **400+ records** organized across multiple tables:

| **Seeder**           | **Records**       | **Description**                                              |
| -------------------- | ----------------- | ------------------------------------------------------------ |
| **Categories**       | 128 records       | Product categories (Electronics, Footwear, Automotive, etc.) |
| **Brands**           | 152 records       | Global brands (Apple, Nike, BMW, Google, etc.)               |
| **Brand-Categories** | 415 relationships | Smart brand-category mappings                                |

### Seeded Data Examples

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

### Seeding CLI Options

| **Flag**      | **Description**         | **Use Case**                     |
| ------------- | ----------------------- | -------------------------------- |
| `-seed`       | Run all seeders         | Adding data to existing database |
| `-fresh-seed` | Drop, migrate, and seed | Complete fresh setup (dev only)  |

### Architecture Overview

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

### Programmatic Usage

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

### Creating Custom Seeders

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

### Verification & Testing

**Quick Verification:**

```bash
go run cmd/seedtest/main.go
```

**Output:**

```
🌱 SEEDING VERIFICATION DEMO
================================
📂 Categories seeded: 128
🏷️ Brands seeded: 152
🔗 Brand-Category relationships: 415

📋 Sample Categories:
   - Footwear (slug: footwear)
   - Electronics (slug: electronics)
   - Automotive (slug: automotive)

🏢 Sample Brands:
   - Apple (slug: apple)
   - Nike (slug: nike)
   - BMW (slug: bmw)

🎯 Apple's Categories:
   - Electronics
   - Mobile Phones
   - Laptops
   - Technology
✅ Seeding verification completed successfully!
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

## 11. CLI Tools Available

### Main Application Server

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

- `POST /register` - User registration
- `POST /login` - User login
- `GET /api/users` - List users (paginated)
- `GET /api/users/all` - Get all users
- `GET /api/users/search` - Search users
- `GET /api/users/count` - Get user count
- `GET /api/users/{id}` - Get user by ID
- `GET /health` - Health check

### Database Migration Tool

```bash
# Safe migration (recommended for production)
go run ./cmd/migrate/main.go -migrate

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
- `-fresh` - Drop all + recreate (development only)
- `-fresh-seed` - Drop, migrate, and seed (development only)
- `-seed` - Run all database seeders
- `-drop` - Drop all tables (dangerous)
- `-fk` - Add foreign keys only
- `-indexes` - Create indexes only
- `-dsn` - Custom database connection string

### Database Seeding Verification Tool

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

### Build All Tools

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

Good luck and have fun learning Clean Architecture in Go!
