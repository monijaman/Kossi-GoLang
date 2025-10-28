-- Find existing products to test reviews with
-- Run this to see which product IDs actually exist in your database

SELECT 
    id,
    name,
    slug,
    price,
    category_id,
    brand_id,
    created_at
FROM products
ORDER BY id ASC
LIMIT 10;

-- Count total products
SELECT COUNT(*) as total_products FROM products;

-- Check if specific products exist
SELECT 
    id,
    name,
    CASE 
        WHEN id = 1 THEN '✓ ID 1 exists'
        WHEN id = 2 THEN '✓ ID 2 exists'
        WHEN id = 3 THEN '✓ ID 3 exists'
        ELSE CONCAT('✓ ID ', id, ' exists')
    END as status
FROM products
WHERE id IN (1, 2, 3, 4, 5)
ORDER BY id;
