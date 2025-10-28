-- Step 1: Check if product_id column exists and its values
\echo '===== Checking product_reviews table structure ====='
\d product_reviews

\echo ''
\echo '===== Checking existing review data ====='
SELECT id, product_id, user_id, rating, LEFT(reviews, 40) as review_excerpt 
FROM product_reviews 
ORDER BY id;

\echo ''
\echo '===== Count reviews with product_id = 0 or NULL ====='
SELECT 
    COUNT(*) FILTER (WHERE product_id IS NULL) as null_count,
    COUNT(*) FILTER (WHERE product_id = 0) as zero_count,
    COUNT(*) FILTER (WHERE product_id > 0) as valid_count,
    COUNT(*) as total_count
FROM product_reviews;

\echo ''
\echo '===== If product_id is 0 or NULL for review ID 1, update it ====='
-- You need to know what product ID review 1 should belong to
-- Replace ? with the actual product ID
-- UPDATE product_reviews SET product_id = ? WHERE id = 1;

\echo ''
\echo '===== Check available products ====='
SELECT id, name FROM products ORDER BY id LIMIT 10;
