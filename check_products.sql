-- Check which products exist in your database

-- 1. List all products
SELECT id, name, slug, status 
FROM products 
ORDER BY id 
LIMIT 20;

-- 2. Count total products
SELECT COUNT(*) as total_products FROM products;

-- 3. Check specific product IDs
SELECT id, name FROM products WHERE id IN (1, 2, 3, 4, 5);

-- Use one of these existing product IDs in your review request!
