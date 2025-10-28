-- Quick Fix: Delete existing review for user 1, product 2
-- Run this in your PostgreSQL terminal

-- Option 1: Delete just the specific review causing conflict
DELETE FROM product_reviews WHERE user_id = 1 AND product_id = 2;

-- Option 2: Delete ALL reviews for user 1 (clean slate for testing)
-- DELETE FROM product_reviews WHERE user_id = 1;

-- Option 3: See what reviews exist first
-- SELECT id, product_id, user_id, rating, LEFT(reviews, 50) as review_text, created_at 
-- FROM product_reviews 
-- WHERE user_id = 1 
-- ORDER BY id DESC;

-- After deleting, try creating the review again!
