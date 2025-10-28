-- Comprehensive Database Check for Review Issue
-- Copy and paste these queries one by one into psql

-- 1. Check if product_id column exists and its properties
SELECT 
    column_name, 
    data_type, 
    is_nullable, 
    column_default
FROM information_schema.columns 
WHERE table_name = 'product_reviews' 
ORDER BY ordinal_position;

-- 2. Check ALL reviews in the database
SELECT 
    id, 
    product_id, 
    user_id, 
    rating, 
    LEFT(reviews, 50) as review_text,
    created_at 
FROM product_reviews 
ORDER BY id DESC 
LIMIT 20;

-- 3. Specifically check for reviews from user_id = 1
SELECT 
    id, 
    product_id, 
    user_id, 
    rating, 
    LEFT(reviews, 50) as review_text
FROM product_reviews 
WHERE user_id = 1;

-- 4. Check for the EXACT conflict (user 1, product 2)
SELECT 
    id, 
    product_id, 
    user_id, 
    rating, 
    reviews,
    created_at,
    updated_at
FROM product_reviews 
WHERE user_id = 1 AND product_id = 2;

-- 5. Check for reviews with NULL or 0 product_id
SELECT 
    id, 
    product_id, 
    user_id, 
    rating
FROM product_reviews 
WHERE product_id IS NULL OR product_id = 0;

-- 6. Count total reviews
SELECT COUNT(*) as total_reviews FROM product_reviews;

-- 7. Count reviews by product
SELECT 
    product_id, 
    COUNT(*) as review_count 
FROM product_reviews 
GROUP BY product_id 
ORDER BY product_id;

-- ==== SOLUTIONS ====

-- If you find a review with user_id=1 and product_id=2, delete it:
-- DELETE FROM product_reviews WHERE user_id = 1 AND product_id = 2;

-- If you find reviews with NULL product_id, delete them:
-- DELETE FROM product_reviews WHERE product_id IS NULL OR product_id = 0;

-- If you want to start completely fresh:
-- DELETE FROM product_reviews;

-- If product_id column doesn't exist, add it:
-- ALTER TABLE product_reviews ADD COLUMN IF NOT EXISTS product_id BIGINT;
-- CREATE INDEX IF NOT EXISTS idx_product_reviews_product_id ON product_reviews(product_id);
