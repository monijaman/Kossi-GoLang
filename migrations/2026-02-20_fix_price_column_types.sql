-- Migration: Convert start_price and end_price columns to varchar for Bengali text support
-- These columns store text/Bengali numerals, not numeric decimals

-- If columns exist as decimal, convert them to varchar:
ALTER TABLE product_translations
ALTER COLUMN start_price TYPE varchar(255),
ALTER COLUMN end_price TYPE varchar(255);

-- If columns don't exist, this would be handled by the previous migration
