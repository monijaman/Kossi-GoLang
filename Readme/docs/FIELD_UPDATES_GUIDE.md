# Field Updates & Database Changes Guide

This guide covers how to handle different types of field changes in your Clean Architecture Go application.

## 📑 **Table of Contents**

1. [Adding New Fields (Safe)](#-1-adding-new-fields-safe)
2. [Modifying Existing Fields (Requires Caution)](#-2-modifying-existing-fields-requires-caution)
3. [Removing Fields (Dangerous)](#-3-removing-fields-dangerous)
4. [Adding Indexes for Performance](#-4-adding-indexes-for-performance)
5. [Adding Foreign Key Relationships](#-5-adding-foreign-key-relationships)
6. [Migration Commands](#-migration-commands)
7. [Testing Field Changes](#-testing-field-changes)
8. [Best Practices](#-best-practices)
9. [Rollback Strategy](#-rollback-strategy)
10. [Environment-Specific Considerations](#-environment-specific-considerations)

---

## 🔄 **1. Adding New Fields (Safe)**

### Steps:
1. Update Domain Entity
2. Update Infrastructure Model  
3. Update Conversion Methods
4. Run Safe Migration

### Example: Adding Phone Field to User

**Domain Entity** (`internal/domain/entity/user.go`):
```go
type User struct {
    ID              UserID
    Name            string
    Email           string
    Phone           *string    // ✅ NEW FIELD
    EmailVerifiedAt *time.Time
    Password        string
    RememberToken   *string
    CreatedAt       time.Time
    UpdatedAt       time.Time
}
```

**Infrastructure Model** (`internal/infrastructure/database/models/user.go`):
```go
type UserModel struct {
    ID              uint       `gorm:"primaryKey;autoIncrement"`
    Name            string     `gorm:"type:varchar(255);not null"`
    Email           string     `gorm:"type:varchar(255);unique;not null"`
    Phone           *string    `gorm:"type:varchar(20)"` // ✅ NEW FIELD
    EmailVerifiedAt *time.Time `gorm:""`
    Password        string     `gorm:"type:varchar(255);not null"`
    RememberToken   *string    `gorm:"type:varchar(100)"`
    CreatedAt       time.Time  `gorm:"autoCreateTime"`
    UpdatedAt       time.Time  `gorm:"autoUpdateTime"`
}

// Update conversion methods
func (u *UserModel) ToEntity() *entities.User {
    return &entities.User{
        ID:              entities.UserID(strconv.Itoa(int(u.ID))),
        Name:            u.Name,
        Email:           u.Email,
        Phone:           u.Phone, // ✅ ADD THIS
        EmailVerifiedAt: u.EmailVerifiedAt,
        Password:        u.Password,
        RememberToken:   u.RememberToken,
        CreatedAt:       u.CreatedAt,
        UpdatedAt:       u.UpdatedAt,
    }
}

func (u *UserModel) FromEntity(entity *entities.User) {
    if id, err := strconv.Atoi(string(entity.ID)); err == nil {
        u.ID = uint(id)
    }
    u.Name = entity.Name
    u.Email = entity.Email
    u.Phone = entity.Phone // ✅ ADD THIS
    u.EmailVerifiedAt = entity.EmailVerifiedAt
    u.Password = entity.Password
    u.RememberToken = entity.RememberToken
    u.CreatedAt = entity.CreatedAt
    u.UpdatedAt = entity.UpdatedAt
}
```

**Migration**:
```bash
go run cmd/migrate/main.go -migrate
```

---

## 🔄 **2. Modifying Existing Fields (Requires Caution)**

### Steps:
1. Update models with new field definition
2. Consider data compatibility
3. Run migration
4. Test thoroughly

### Example: Changing Email Field Size

**Before**:
```go
Email string `gorm:"type:varchar(255);unique;not null"`
```

**After**:
```go
Email string `gorm:"type:varchar(500);unique;not null"` // Increased size
```

**Migration**:
```bash
go run cmd/migrate/main.go -migrate
```

---

## 🗑️ **3. Removing Fields (Dangerous)**

### ⚠️ **Safe Approach - Two-Step Process**:

**Step 1**: Make field optional and remove from business logic
```go
// Mark as deprecated, make nullable
OldField *string `gorm:"type:varchar(255)"` // Made optional
```

**Step 2**: After confirming no usage, remove completely
```go
// Remove field entirely
```

**Migration**:
```bash
go run cmd/migrate/main.go -migrate
```

---

## 🏗️ **4. Adding Indexes for Performance**

### Example: Adding Index to Phone Field

**Infrastructure Model**:
```go
type UserModel struct {
    // ... other fields
    Phone *string `gorm:"type:varchar(20);index"` // ✅ ADD INDEX
}
```

**Migration**:
```bash
go run cmd/migrate/main.go -indexes
```

---

## 🔗 **5. Adding Foreign Key Relationships**

### Example: Adding UserID to Product

**Domain Entity** (`internal/domain/entity/product.go`):
```go
type Product struct {
    ID       ProductID
    Name     string
    UserID   *UserID   // ✅ NEW RELATIONSHIP
    // ... other fields
}
```

**Infrastructure Model** (`internal/infrastructure/database/models/product.go`):
```go
type ProductModel struct {
    ID       uint      `gorm:"primaryKey;autoIncrement"`
    Name     string    `gorm:"type:varchar(255);not null"`
    UserID   *uint     `gorm:"index"` // ✅ NEW FOREIGN KEY
    // ... other fields
    
    // ✅ ADD RELATIONSHIP
    User     *UserModel `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
```

**Migration**:
```bash
go run cmd/migrate/main.go -migrate
go run cmd/migrate/main.go -fk  # Add foreign keys
```

---

## 📋 **Migration Commands**

### Safe Operations:
```bash
# Add new fields, modify existing fields (safe changes)
go run cmd/migrate/main.go -migrate

# Add only foreign keys
go run cmd/migrate/main.go -fk

# Add only indexes
go run cmd/migrate/main.go -indexes
```

### Dangerous Operations:
```bash
# Drop all tables and recreate (DEVELOPMENT ONLY!)
go run cmd/migrate/main.go -fresh

# Drop all tables
go run cmd/migrate/main.go -drop
```

---

## 🔍 **Testing Field Changes**

### 1. Unit Tests for Conversion Methods
```go
func TestUserModelConversion(t *testing.T) {
    entity := &entities.User{
        ID:    "1",
        Name:  "John Doe",
        Email: "john@example.com",
        Phone: &[]string{"1234567890"}[0], // Test new field
    }
    
    model := &models.UserModel{}
    model.FromEntity(entity)
    
    converted := model.ToEntity()
    
    assert.Equal(t, entity.Phone, converted.Phone)
}
```

### 2. Integration Tests
```go
func TestCreateUserWithPhone(t *testing.T) {
    user := &entities.User{
        Name:  "Test User",
        Email: "test@example.com", 
        Phone: &[]string{"1234567890"}[0],
    }
    
    err := userRepo.Create(user)
    assert.NoError(t, err)
    
    retrieved, err := userRepo.FindByEmail("test@example.com")
    assert.NoError(t, err)
    assert.Equal(t, user.Phone, retrieved.Phone)
}
```

---

## 🚨 **Best Practices**

### ✅ **Do:**
- Always backup database before major changes
- Make new fields nullable initially
- Test conversion methods thoroughly
- Use safe migrations first
- Add fields before removing old ones

### ❌ **Don't:**
- Remove fields without checking dependencies
- Change field types without data migration plan
- Use `-fresh` in production
- Forget to update conversion methods
- Skip testing after field changes

---

## 🔄 **Rollback Strategy**

If you need to rollback changes:

1. **For added fields** (usually safe):
   ```bash
   # Remove from code and run migration
   go run cmd/migrate/main.go -migrate
   ```

2. **For removed fields** (dangerous):
   ```bash
   # You'll need to restore from backup
   # This is why we recommend the two-step approach
   ```

3. **For modified fields**:
   ```bash
   # Revert the field definition and migrate
   go run cmd/migrate/main.go -migrate
   ```

---

## 📝 **Environment-Specific Considerations**

### Development:
```bash
# You can use fresh migrations
go run cmd/migrate/main.go -fresh
```

### Staging:
```bash
# Use safe migrations only
go run cmd/migrate/main.go -migrate
```

### Production:
```bash
# Always backup first, then safe migrate
pg_dump kossti > backup_$(date +%Y%m%d_%H%M%S).sql
go run cmd/migrate/main.go -migrate
```
