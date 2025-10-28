# ✅ Migration Best Practice - No Extra Files Needed

## The Right Way: Update Existing Migration File

When adding new fields to models in the Go server, **DO NOT create separate migration files**. Instead:

### ✅ Correct Approach

1. **Update the Model** - Add fields to the struct in `internal/infrastructure/database/models/`
2. **Ensure Model is in Migrator** - Check `internal/infrastructure/database/migrator.go`
3. **Restart Server** - GORM AutoMigrate will handle it

### ❌ Wrong Approach (What I Did Initially)

- ❌ Creating separate `.sql` migration files
- ❌ Manual ALTER TABLE scripts
- ❌ Multiple migration files for one change

## What Was Fixed

### 1. Updated the Model ✅

**File:** `internal/infrastructure/database/models/product_review.go`

Added `ProductID` field:

```go
type ProductReviewModel struct {
    ID                uint       `gorm:"primaryKey;autoIncrement"`
    ProductID         uint       `gorm:"not null;index"`  // ← Added this!
    UserID            uint       `gorm:"not null;index"`
    Rating            float32    `gorm:"not null"`
    // ... rest of fields
}
```

### 2. Added Model to AutoMigrate ✅

**File:** `internal/infrastructure/database/migrator.go`

Added `ProductReviewModel` and other missing models:

```go
err := m.db.AutoMigrate(
    &models.UserModel{},
    &models.ProductModel{},
    &models.ImageModel{},
    &models.ProductTranslationModel{},
    &models.ProductReviewModel{},            // ← Added!
    &models.ProductReviewTranslationModel{}, // ← Added!
    &models.CategoryModel{},                 // ← Added!
    &models.BrandModel{},                    // ← Added!
    // ... and more
)
```

## How GORM AutoMigrate Works

When you restart the Go server:

1. **GORM reads your structs** with `gorm:` tags
2. **Compares with database** schema
3. **Automatically adds missing columns**
4. **Creates indexes** based on tags
5. **NO MANUAL SQL NEEDED!**

### Example: Adding a New Field

**Step 1:** Update the model

```go
type ProductReviewModel struct {
    ID          uint
    ProductID   uint       `gorm:"not null;index"`
    NewField    string     `gorm:"type:varchar(255)"` // ← Add new field
}
```

**Step 2:** Make sure model is in migrator.go AutoMigrate list (already done!)

**Step 3:** Restart server

```bash
go run cmd/app/main.go
```

**Done!** GORM adds the column automatically.

## Models Now in AutoMigrate

All these models will be auto-migrated on server restart:

✅ UserModel  
✅ ProductModel  
✅ ImageModel  
✅ ProductTranslationModel  
✅ **ProductReviewModel** ← Fixed!  
✅ **ProductReviewTranslationModel** ← Fixed!  
✅ CategoryModel  
✅ BrandModel  
✅ BrandCategoryModel  
✅ SpecificationModel  
✅ SpecificationKeyModel  
✅ FormGeneratorModel  
✅ FeedbackModel  
✅ RefreshTokenModel  
✅ PasswordResetTokenModel

## For Future Development

### Adding New Models

1. Create model in `internal/infrastructure/database/models/your_model.go`
2. Add to `migrator.go` AutoMigrate list
3. Restart server - Done!

### Adding Fields to Existing Models

1. Add field to struct in the model file
2. Model already in migrator? Nothing else needed!
3. Restart server - GORM adds the column!

### Never Do This

❌ Create `.sql` migration files  
❌ Write manual ALTER TABLE statements  
❌ Create "add_field_migration.sql" files  
❌ Keep separate migration scripts

## Benefits of This Approach

✅ **Single Source of Truth** - Models define schema  
✅ **No SQL Files** - Everything in Go code  
✅ **Automatic Sync** - GORM handles it  
✅ **Type Safe** - Compiler checks your models  
✅ **Easy to Track** - Changes in version control  
✅ **Cross-Database** - GORM adapts SQL for your DB

## Summary

**For product_id fix:**

- ✅ Updated `ProductReviewModel` with `ProductID` field
- ✅ Added `ProductReviewModel` to migrator.go
- ✅ No separate migration files needed
- ✅ Just restart server - GORM does the rest!

**Going forward:**

- Always update the model struct
- Ensure model is in migrator.go
- Let GORM AutoMigrate handle database changes
- No manual SQL files!

---

**Date:** October 26, 2025  
**Principle:** Models are the single source of truth  
**Tool:** GORM AutoMigrate handles everything 🚀
