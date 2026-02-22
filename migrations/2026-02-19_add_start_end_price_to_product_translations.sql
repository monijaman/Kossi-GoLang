-- Migration: Add start_price and end_price to product_translations as VARCHAR for Bengali text support
-- These columns store text/Bengali numerals, not numeric decimals
ALTER TABLE product_translations
ADD COLUMN IF NOT EXISTS start_price varchar(255),
ADD COLUMN IF NOT EXISTS end_price varchar(255);

-- Optionally, migrate legacy price if needed (uncomment if you have a price column)
-- UPDATE product_translations SET start_price = price WHERE start_price IS NULL AND price IS NOT NULL;
-- ALTER TABLE product_translations DROP COLUMN IF EXISTS price;