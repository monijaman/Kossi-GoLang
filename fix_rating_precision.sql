-- Fix rating column precision to store only 1 decimal place
-- This prevents values like 3.9000000953674316 and ensures 3.9

ALTER TABLE product_reviews ALTER COLUMN rating TYPE NUMERIC(2,1);

-- Update existing records to round to 1 decimal place (optional, but good for cleanup)
UPDATE product_reviews SET rating = ROUND(rating::numeric, 1);
