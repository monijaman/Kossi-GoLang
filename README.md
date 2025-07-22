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

| Flag | Description | Safety Level |
|------|-------------|--------------|
| `-migrate` | Run safe migration (recommended) | ✅ **Safe** |
| `-fresh` | Drop all + recreate (development) | 🚨 **Dangerous** |
| `-drop` | Drop all tables | 🚨 **Dangerous** |
| `-fk` | Add foreign keys only | ✅ **Safe** |
| `-indexes` | Create indexes only | ✅ **Safe** |
| `-dsn` | Custom database connection string | ℹ️ **Config** |

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

| Feature | AutoMigrate | Migration Manager |
|---------|-------------|-------------------|
| Model Count | Manual (1-5 models) | Automatic (60+ models) |
| Progress Tracking | ❌ No | ✅ Yes |
| Foreign Keys | ❌ Manual | ✅ Automatic |
| Indexes | ❌ Manual | ✅ Automatic |
| Error Handling | ❌ Basic | ✅ Advanced |
| Production Ready | ⚠️ Limited | ✅ Yes |

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

# Show help and all options
go run ./cmd/migrate/main.go -h

# Build the migration tool
go build -o bin/kossti-migrate ./cmd/migrate
```

**Migration Options:**
- `-migrate` - Safe migration (production ready)
- `-fresh` - Drop all + recreate (development only)
- `-drop` - Drop all tables (dangerous)
- `-fk` - Add foreign keys only
- `-indexes` - Create indexes only
- `-dsn` - Custom database connection string

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
