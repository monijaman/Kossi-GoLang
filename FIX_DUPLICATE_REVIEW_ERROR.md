# ✅ Fix: "user has already reviewed this product" Error

## Problem

Getting error when trying to create a review:

```json
{
  "error": "user has already reviewed this product"
}
```

## Root Causes & Solutions

### Cause 1: You Already Have a Review (Most Likely) ✅

**The duplicate prevention is working!** User ID 1 already has a review for product 2.

**Check:**

```sql
-- Connect to your database
psql -U your_username -d your_database_name

-- Check existing reviews
SELECT id, product_id, user_id, rating, reviews
FROM product_reviews
WHERE user_id = 1 AND product_id = 2;
```

**Solutions:**

**Option A: Delete the existing review**

```sql
DELETE FROM product_reviews WHERE user_id = 1 AND product_id = 2;
```

**Option B: Try a different product**

```bash
POST /reviews/3  # Try product 3 instead
POST /reviews/5  # Or product 5
```

**Option C: Test with different user** (once JWT is implemented)
Currently user_id is hardcoded to 1 in the handler.

---

### Cause 2: product_id Column Doesn't Exist Yet ⚠️

If you haven't restarted the Go server after the fix, the `product_id` column might not exist.

**Check:**

```sql
-- Check if column exists
SELECT column_name
FROM information_schema.columns
WHERE table_name = 'product_reviews' AND column_name = 'product_id';
```

**Solution:**

```bash
# Restart the Go server to trigger GORM AutoMigrate
cd gocrit_server
go run cmd/app/main.go
```

GORM will automatically add the `product_id` column.

---

### Cause 3: Old Reviews with NULL/0 product_id 🔧

If you have old reviews in the database from before the fix, they might have NULL or 0 for product_id.

**Check:**

```sql
-- Check for problematic reviews
SELECT id, product_id, user_id, rating
FROM product_reviews
WHERE product_id IS NULL OR product_id = 0;
```

**Solution:**

```sql
-- Option 1: Delete old reviews without product_id
DELETE FROM product_reviews WHERE product_id IS NULL OR product_id = 0;

-- Option 2: Update them with correct product_id (if you know what they should be)
-- UPDATE product_reviews SET product_id = ? WHERE id = ?;
```

---

## Complete Diagnostic Script

Run this SQL script to understand your current state:

```sql
-- 1. Check table structure
\d product_reviews

-- 2. Check all reviews
SELECT id, product_id, user_id, rating, LEFT(reviews, 30) as review_text, created_at
FROM product_reviews
ORDER BY id DESC;

-- 3. Check for the specific conflict
SELECT * FROM product_reviews WHERE user_id = 1 AND product_id = 2;

-- 4. Check for NULL product_ids
SELECT COUNT(*) FROM product_reviews WHERE product_id IS NULL OR product_id = 0;
```

---

## Testing Steps

### Step 1: Clean Database (Fresh Start)

```sql
-- Delete all reviews (for testing)
DELETE FROM product_reviews;

-- Or just delete reviews for user 1
DELETE FROM product_reviews WHERE user_id = 1;
```

### Step 2: Create First Review

```bash
POST http://localhost:8080/reviews/2
Authorization: Bearer YOUR_TOKEN

{
  "rating": 5,
  "reviews": "First review for product 2"
}

# Expected: ✅ Success
# Response: { "message": "Review added successfully", "review": { "product_id": 2 } }
```

### Step 3: Try Duplicate (Should Fail)

```bash
POST http://localhost:8080/reviews/2
Authorization: Bearer YOUR_TOKEN

{
  "rating": 4,
  "reviews": "Second review for same product"
}

# Expected: ❌ Error
# Response: { "error": "user has already reviewed this product" }
```

### Step 4: Different Product (Should Succeed)

```bash
POST http://localhost:8080/reviews/3
Authorization: Bearer YOUR_TOKEN

{
  "rating": 5,
  "reviews": "Review for product 3"
}

# Expected: ✅ Success
# Response: { "message": "Review added successfully", "review": { "product_id": 3 } }
```

---

## How the Duplicate Check Works

```go
// From: internal/usecase/productreview/product_review.go

// Check if user already has a review for this product
existingReview, err := repo.GetByProductAndUser(ctx, productID, userID)
if err == nil && existingReview != nil {
    return nil, errors.New("user has already reviewed this product")
}
```

This is **correct behavior** to prevent spam and ensure one review per user per product.

---

## Quick Fix Commands

### If you want to test creating multiple reviews:

```sql
-- Allow testing by deleting existing reviews
DELETE FROM product_reviews WHERE user_id = 1;
```

### If product_id column is missing:

```bash
# Restart Go server
cd gocrit_server
go run cmd/app/main.go
```

### If you want to verify the fix worked:

```sql
-- Create a review, then check it was stored with product_id
SELECT * FROM product_reviews ORDER BY id DESC LIMIT 1;
-- Should show: product_id = 2 (or whatever product you reviewed)
```

---

## Summary

| Scenario                         | Solution                                               |
| -------------------------------- | ------------------------------------------------------ |
| Already reviewed product 2       | ✅ Try product 3, 4, 5, etc. OR delete existing review |
| product_id column missing        | ⚠️ Restart Go server to run migrations                 |
| Old reviews with NULL product_id | 🔧 Delete or update old reviews                        |
| Want to test multiple reviews    | 🧪 Delete reviews between tests                        |

**The error message is actually CORRECT behavior** - it's preventing duplicate reviews!

To continue testing, either:

1. Use different product IDs (3, 4, 5...)
2. Delete existing reviews for user 1
3. Wait for JWT implementation to test with different users

---

**Date:** October 26, 2025  
**Status:** Working as intended - Duplicate prevention active! ✅
