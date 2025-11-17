-- Fix rating column type from numeric to varchar
-- Run this SQL in your PostgreSQL database

-- Convert the rating column to varchar, preserving existing data
ALTER TABLE product_review_translations 
ALTER COLUMN rating TYPE varchar(50) USING rating::text;

-- Verify the change
SELECT column_name, data_type, character_maximum_length
FROM information_schema.columns
WHERE table_name = 'product_review_translations' AND column_name = 'rating';
