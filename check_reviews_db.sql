-- Check product_reviews table structure and data
-- Run this to diagnose the "user has already reviewed this product" error

-- 1. Check if product_id column exists
SELECT column_name, data_type, is_nullable 
FROM information_schema.columns 
WHERE table_name = 'product_reviews' 
ORDER BY ordinal_position;

-- 2. Check existing reviews
SELECT id, product_id, user_id, rating, reviews, created_at 
FROM product_reviews 
ORDER BY id DESC 
LIMIT 20;

-- 3. Check for duplicate reviews (same user, same product)
SELECT product_id, user_id, COUNT(*) as review_count
FROM product_reviews 
GROUP BY product_id, user_id
HAVING COUNT(*) > 1;

-- 4. Check reviews for product_id = 2
SELECT id, product_id, user_id, rating, reviews 
FROM product_reviews 
WHERE product_id = 2;

-- 5. Check if product_id column has NULL values
SELECT COUNT(*) as null_product_ids
FROM product_reviews 
WHERE product_id IS NULL OR product_id = 0;

-- SOLUTIONS:

-- If product_id column doesn't exist, add it:
-- ALTER TABLE product_reviews ADD COLUMN IF NOT EXISTS product_id BIGINT;

-- If you want to delete the existing review for user 1 and product 2:
-- DELETE FROM product_reviews WHERE user_id = 1 AND product_id = 2;

-- If you want to see which user is trying to create duplicate:
-- This is the current user_id (check your JWT token or it might be hardcoded to 1)
