# N+1 Query Problem - Diagnosis and Solution

## Problem Description

### What is the N+1 Query Problem?

The N+1 query problem occurs when an application executes one query to fetch a list of N records, and then executes N additional queries to fetch related data for each record. This results in 1 + N total database queries.

### Example Scenario

When fetching 100 products with their categories and brands:

**Bad (N+1 Queries):**
```
1. SELECT * FROM products LIMIT 100;           -- 1 query
2. SELECT * FROM categories WHERE id = 1;      -- 100 queries for categories
3. SELECT * FROM categories WHERE id = 2;
...
102. SELECT * FROM brands WHERE id = 5;        -- 100 queries for brands
103. SELECT * FROM brands WHERE id = 3;
...
Total: 201 queries!
```

**Good (Optimized):**
```
1. SELECT * FROM products LIMIT 100;
2. SELECT * FROM categories WHERE id IN (1, 2, 3, ...);
3. SELECT * FROM brands WHERE id IN (5, 3, 7, ...);
Total: 3 queries!
```

## How to Identify N+1 Query Problems

### Symptoms
- Slow API response times for list endpoints
- Response time increases linearly with result count
- Database connection pool exhaustion
- High database CPU usage

### Detection Methods

1. **Enable GORM Logging:**
```go
db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
    Logger: logger.Default.LogMode(logger.Info), // Shows all SQL queries
})
```

2. **Check Query Count:**
Look for patterns where queries repeat with different IDs.

3. **Performance Monitoring:**
Monitor database query counts vs API response times.

## Solution: GORM Preload

### Implementation Steps

#### Step 1: Add Relationship Fields to Database Models

**File:** `internal/infrastructure/database/models/product.go`

```go
type ProductModel struct {
    ID          uint       `gorm:"primaryKey;autoIncrement"`
    Name        string     `gorm:"type:varchar(255);not null"`
    CategoryID  *uint      `gorm:""`
    BrandID     *uint      `gorm:""`
    
    // Add relationship fields
    Category    *CategoryModel `gorm:"foreignKey:CategoryID"`
    Brand       *BrandModel    `gorm:"foreignKey:BrandID"`
    
    // ... other fields
}
```

**Important Notes:**
- These fields are for ORM relationships only
- No database schema changes required
- Existing foreign key columns remain unchanged
- Use pointer types (`*CategoryModel`) to handle NULL foreign keys

#### Step 2: Update ToEntity Conversion

```go
func (p *ProductModel) ToEntity() *entities.Product {
    // Handle preloaded category
    var category *entities.Category
    if p.Category != nil {
        category = p.Category.ToEntity()
    }
    
    // Handle preloaded brand
    var brand *entities.Brand
    if p.Brand != nil {
        brand = p.Brand.ToEntity()
    }

    return &entities.Product{
        ID:         p.ID,
        Name:       p.Name,
        CategoryID: p.CategoryID,
        BrandID:    p.BrandID,
        Category:   category,  // Include preloaded data
        Brand:      brand,     // Include preloaded data
        // ... other fields
    }
}
```

#### Step 3: Add Relationship Fields to Domain Entities

**File:** `internal/domain/entities/product.go`

```go
type Product struct {
    ID         uint       `json:"id"`
    Name       string     `json:"name"`
    CategoryID *uint      `json:"category_id,omitempty"`
    BrandID    *uint      `json:"brand_id,omitempty"`
    
    // Add relationship fields
    Category   *Category  `json:"category,omitempty"`
    Brand      *Brand     `json:"brand,omitempty"`
    
    // ... other fields
}
```

#### Step 4: Add Preload to Repository Queries

**File:** `internal/interface/repository/postgres/product.go`

```go
// List method
func (r *PostgresProductRepo) List(ctx context.Context, limit, offset int) ([]*entities.Product, error) {
    var productModels []models.ProductModel
    query := r.db.WithContext(ctx).
        Preload("Category").        // ← Add this
        Preload("Brand").           // ← Add this
        Where("deleted_at IS NULL")

    if limit > 0 {
        query = query.Limit(limit)
    }
    if offset > 0 {
        query = query.Offset(offset)
    }

    if err := query.Find(&productModels).Error; err != nil {
        return nil, err
    }
    
    // Convert to entities - preloaded data is included
    products := make([]*entities.Product, len(productModels))
    for i, model := range productModels {
        products[i] = model.ToEntity()
    }

    return products, nil
}
```

**Apply to all list-type queries:**
- `List()`
- `Search()`
- `GetByCategory()`
- `GetByBrand()`
- `GetPopular()`
- `GetWithFilters()`
- `GetSimilarProducts()`
- `GetByID()` (for single items)
- `GetBySlug()` (for single items)

#### Step 5: Update Handler Logic (Optional but Recommended)

**File:** `internal/interface/handler/product/handler.go`

```go
func convertProductToResponse(product *entities.Product, categoryRepo repository.CategoryRepository, brandRepo repository.BrandRepository) ProductResponse {
    response := ProductResponse{
        ID:   product.ID,
        Name: product.Name,
        // ... other fields
    }

    // Use preloaded data if available
    if product.Category != nil {
        response.Category = &CategoryResponse{
            ID:   product.Category.ID,
            Name: product.Category.Name,
            Slug: product.Category.Slug,
        }
    } else if product.CategoryID != nil && categoryRepo != nil {
        // Fallback for cases without preload
        category, err := categoryRepo.GetByID(context.Background(), *product.CategoryID)
        if err == nil && category != nil {
            response.Category = &CategoryResponse{
                ID:   category.ID,
                Name: category.Name,
                Slug: category.Slug,
            }
        }
    }

    // Same for brand...
    
    return response
}
```

## Advanced Preloading Techniques

### Nested Preloading

```go
// Preload products with categories, brands, and images
db.Preload("Category").
   Preload("Brand").
   Preload("Images").
   Find(&products)
```

### Conditional Preloading

```go
// Preload only active categories
db.Preload("Category", "status = ?", 1).
   Find(&products)
```

### Preload with Custom Query

```go
// Preload categories ordered by name
db.Preload("Category", func(db *gorm.DB) *gorm.DB {
    return db.Order("name ASC")
}).Find(&products)
```

### Multiple Levels Deep

```go
// Preload product -> category -> parent category
db.Preload("Category.ParentCategory").
   Find(&products)
```

## Performance Comparison

### Before Optimization
```
Endpoint: GET /products?limit=100
Query Count: 201 queries
Response Time: ~500ms
Database Load: High
```

### After Optimization
```
Endpoint: GET /products?limit=100
Query Count: 3 queries
Response Time: ~50ms
Database Load: Low
```

**Performance Improvement:** ~90% faster

## Checklist for Future Optimizations

When adding new list endpoints or optimizing existing ones:

- [ ] Identify all relationships that need to be included in the response
- [ ] Add relationship fields to the database model (e.g., `Category *CategoryModel`)
- [ ] Add relationship fields to the domain entity (e.g., `Category *Category`)
- [ ] Update `ToEntity()` method to handle preloaded data safely
- [ ] Add `.Preload("RelationName")` to repository queries
- [ ] Update handler to use preloaded data when available
- [ ] Test with query logging enabled to verify optimization
- [ ] Measure performance improvement

## Common Pitfalls to Avoid

### 1. Forgetting Nil Checks
```go
// ❌ BAD - will panic if Category is nil
category := p.Category.ToEntity()

// ✅ GOOD - safe handling
var category *entities.Category
if p.Category != nil {
    category = p.Category.ToEntity()
}
```

### 2. Not Using Preload Consistently
```go
// ❌ BAD - mixing preloaded and non-preloaded queries
products1 := repo.List()        // With preload
products2 := repo.GetByID(123)  // Without preload

// ✅ GOOD - consistent preloading
products1 := repo.List()        // With preload
products2 := repo.GetByID(123)  // Also with preload
```

### 3. Over-Preloading
```go
// ❌ BAD - loading unnecessary data
db.Preload("Category").
   Preload("Brand").
   Preload("Images").
   Preload("Reviews").
   Preload("Specifications").  // Too much data!
   Find(&products)

// ✅ GOOD - only load what's needed for the response
db.Preload("Category").
   Preload("Brand").
   Find(&products)
```

### 4. Preload in Count Queries
```go
// ❌ BAD - wasting resources
db.Preload("Category").Preload("Brand").Count(&count)

// ✅ GOOD - no preload for count
db.Model(&ProductModel{}).Count(&count)
```

## Database Schema Requirements

**Important:** Preloading requires:
1. Foreign key columns exist (e.g., `category_id`, `brand_id`)
2. Related tables exist (e.g., `categories`, `brands`)
3. No additional indexes required (but recommended for performance)

**Recommended Indexes:**
```sql
CREATE INDEX idx_products_category_id ON products(category_id);
CREATE INDEX idx_products_brand_id ON products(brand_id);
```

## Testing Preloaded Queries

### Enable Query Logging in Development

**File:** `cmd/app/main.go`

```go
db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
    Logger: logger.Default.LogMode(logger.Info),
})
```

### Verify Query Count

Look for patterns like:
```sql
-- Good: Batch loading
SELECT * FROM categories WHERE id IN (1, 2, 3, 4, 5);

-- Bad: Individual loading
SELECT * FROM categories WHERE id = 1;
SELECT * FROM categories WHERE id = 2;
...
```

## References

- [GORM Preloading Documentation](https://gorm.io/docs/preload.html)
- [N+1 Query Problem Explained](https://stackoverflow.com/questions/97197/what-is-the-n1-selects-problem-in-orm-object-relational-mapping)
- Project implementation: `gocrit_server/internal/interface/repository/postgres/product.go`

## Related Files

- `internal/infrastructure/database/models/product.go` - Model with relationships
- `internal/domain/entities/product.go` - Entity with relationship fields
- `internal/interface/repository/postgres/product.go` - Repository with preload
- `internal/interface/handler/product/handler.go` - Handler using preloaded data

---

**Last Updated:** February 6, 2026  
**Applied to:** kossti.com product listing optimization
