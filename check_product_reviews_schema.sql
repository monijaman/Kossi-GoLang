-- Check the structure of product_reviews table
SELECT column_name, data_type, is_nullable
FROM information_schema.columns
WHERE table_name = 'product_reviews'
ORDER BY ordinal_position;

-- Check actual data in the table
SELECT id, product_id, user_id, rating, LEFT(reviews, 50) as review_excerpt
FROM product_reviews
ORDER BY id
LIMIT 10;

-- Check if product_id column exists
SELECT EXISTS (
    SELECT 1 
    FROM information_schema.columns 
    WHERE table_name = 'product_reviews' 
    AND column_name = 'product_id'
) as product_id_column_exists;
