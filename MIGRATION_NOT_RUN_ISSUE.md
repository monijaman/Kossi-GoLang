# CRITICAL: Did You Restart the Server After Migration Changes?

## The Issue

You're getting wrong data because the `product_id` column might not exist in your database yet!

## What Happened

1. We added `ProductID uint` field to `ProductReviewModel`
2. We added `ProductReviewModel` to `migrator.go`
3. **BUT** you need to restart the server for GORM AutoMigrate to run!

## The Evidence

The query `WHERE product_id = 3 AND user_id = 1` is returning the record with `product_id = 1`.

This can only happen if:

- The `product_id` column doesn't exist (GORM ignores it in WHERE clause)
- OR the column exists but is NULL/0 for all records

## The Fix

### Step 1: Stop the Server Completely

Press `Ctrl+C` and make sure it's fully stopped.

### Step 2: Check Current Schema

Run this in a new terminal:

```bash
psql -U postgres -d kossti -c "\d product_reviews"
```

Look for the `product_id` column. If it's missing, that's the problem!

### Step 3: Restart the Server

```bash
cd gocrit_server
go run cmd/app/main.go
```

When the server starts, GORM AutoMigrate will:

- Add the `product_id` column if it doesn't exist
- Add indexes
- BUT it won't populate existing rows!

### Step 4: Update Existing Records

If you have existing reviews with NULL or 0 for product_id, update them:

```sql
-- Check which reviews have product_id = 0 or NULL
SELECT id, product_id, user_id FROM product_reviews WHERE product_id = 0 OR product_id IS NULL;

-- Update them with actual product IDs (you'll need to know which product each review belongs to)
-- Example:
UPDATE product_reviews SET product_id = 1 WHERE id = 1;
```

### Step 5: Test Again

After restarting and ensuring product_id column exists with correct data, try creating a review for product 3.

## Quick Diagnostic

Run this to see the current state:

```bash
psql -U postgres -d kossti -c "SELECT id, COALESCE(product_id, -1) as product_id, user_id, rating FROM product_reviews LIMIT 5;"
```

If you see `-1` or `0` for product_id, the column exists but has no data!
If you get an error "column product_id does not exist", the migration hasn't run yet!
