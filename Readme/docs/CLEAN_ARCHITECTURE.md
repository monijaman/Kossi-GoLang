# Clean Architecture Implementation

This project follows Clean Architecture principles to ensure loose coupling and high testability.

## Architecture Layers

```
┌─────────────────────────────────────────────────────────────┐
│                    Interface Layer                          │
│                  (HTTP, gRPC, CLI)                         │
└─────────────────────┬───────────────────────────────────────┘
                      │
┌─────────────────────▼───────────────────────────────────────┐
│                 Application Layer                           │
│                  (Use Cases)                               │
└─────────────────────┬───────────────────────────────────────┘
                      │
┌─────────────────────▼───────────────────────────────────────┐
│                  Domain Layer                               │
│              (Entities + Repositories)                     │
└─────────────────────┬───────────────────────────────────────┘
                      │
┌─────────────────────▼───────────────────────────────────────┐
│               Infrastructure Layer                          │
│           (Database, External APIs, etc.)                  │
└─────────────────────────────────────────────────────────────┘
```

## Directory Structure

```
internal/
├── domain/                     # Domain Layer (Framework Independent)
│   ├── entities/               # Business entities
│   │   └── permission.go
│   └── repositories/           # Repository interfaces
│       └── permission_repository.go
├── usecase/                    # Application Layer
│   └── permission_usecase.go   # Business logic orchestration
├── infrastructure/             # Infrastructure Layer
│   ├── database/
│   │   └── models/             # GORM-specific models
│   │       └── permission.go
│   └── repositories/           # Repository implementations
│       ├── permission_repository_impl.go      # GORM implementation
│       └── mongo_permission_repository_impl.go # MongoDB example
└── config/                     # Dependency injection
    └── dependency_injection.go
```

## Key Benefits

### 1. **ORM Independence**
Your domain entities are completely independent of GORM:

```go
// Domain Entity (Framework Independent)
type Permission struct {
    ID        uint
    Name      string
    GuardName string
    CreatedAt time.Time
    UpdatedAt time.Time
}

// Infrastructure Model (GORM Specific)
type PermissionModel struct {
    ID        uint      `gorm:"primaryKey;autoIncrement"`
    Name      string    `gorm:"type:varchar(255);not null"`
    GuardName string    `gorm:"type:varchar(255);not null"`
    CreatedAt time.Time `gorm:"autoCreateTime"`
    UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
```

### 2. **Easy ORM Switching**
To switch from GORM to MongoDB, you only need to:
1. Create new repository implementations
2. Update dependency injection
3. **Zero changes to business logic**

### 3. **Testability**
Mock the repository interfaces for unit testing:

```go
type MockPermissionRepository struct{}

func (m *MockPermissionRepository) Create(permission *entities.Permission) error {
    // Mock implementation
    return nil
}

func (m *MockPermissionRepository) GetByID(id uint) (*entities.Permission, error) {
    // Mock implementation
    return &entities.Permission{ID: id, Name: "test"}, nil
}
```

### 4. **Business Logic Isolation**
Your use cases contain pure business logic:

```go
func (uc *PermissionUseCase) CreatePermission(name, guardName string) (*entities.Permission, error) {
    // Business logic: Check if permission already exists
    existing, _ := uc.permissionRepo.GetByName(name)
    if existing != nil {
        return nil, errors.New("permission already exists")
    }

    // Business logic: Validate permission name
    if name == "" {
        return nil, errors.New("permission name cannot be empty")
    }

    // This logic works with ANY repository implementation!
    permission := &entities.Permission{Name: name, GuardName: guardName}
    return permission, uc.permissionRepo.Create(permission)
}
```

## How to Extend

### Adding a New Entity
1. Create domain entity in `domain/entities/`
2. Add repository interface in `domain/repositories/`
3. Implement repository in `infrastructure/repositories/`
4. Create use case in `usecase/`
5. Wire dependencies in `config/`

### Switching ORMs
1. Create new repository implementations
2. Update dependency injection
3. **No changes needed in domain or use case layers!**

## Example Usage

```go
// Setup dependencies (GORM)
container := config.NewServiceContainerWithGORM(gormDB)

// Use business logic (framework agnostic)
permission, err := container.PermissionUseCase.CreatePermission("read_users", "web")

// Later switch to MongoDB by changing one line:
// container := config.NewServiceContainerWithMongoDB(mongoDB)
// Same business logic works perfectly!
```

This architecture ensures your code is maintainable, testable, and truly follows Clean Architecture principles!
