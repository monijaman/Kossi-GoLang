-- ============================================================================
-- Product Translations - Bengali (bn) Sample Data
-- ============================================================================
-- 
-- This file contains sample SQL inserts for Bengali product translations.
-- Modify the translated_name values to match your actual products.
--
-- HOW TO USE:
-- 1. Connect to your PostgreSQL database
-- 2. Run this script: psql -U postgres -h localhost -d gocrit < product_translations_bn_sample.sql
-- 3. Or copy-paste the INSERT statements in your database tool
--
-- ============================================================================

-- Check current Bengali translation count
SELECT COUNT(*) as "Current Bengali Translations" 
FROM product_translations 
WHERE locale = 'bn';

-- Insert sample Bengali translations (modify these to match your products)
-- Note: These will replace any existing translations for the same product_id + locale combination
-- If you want to avoid duplicates, use ON CONFLICT clause:

INSERT INTO product_translations (product_id, locale, translated_name, price, created_at, updated_at)
VALUES
  -- Smart Phones
  (1, 'bn', 'স্যামসাং গ্যালাক্সি এস২৫', NULL, NOW(), NOW()),
  (2, 'bn', 'আইফোন ১৫ প্রো', NULL, NOW(), NOW()),
  (3, 'bn', 'শাওমি রেডমি নোট ১৫', NULL, NOW(), NOW()),
  (4, 'bn', 'রিয়েলমি ১২', NULL, NOW(), NOW()),
  (5, 'bn', 'ভিভো এক্স ১০০', NULL, NOW(), NOW()),
  (6, 'bn', 'ওপিও এ৩ প্রো', NULL, NOW(), NOW()),
  (7, 'bn', 'ওয়ানপ্লাস ১৩', NULL, NOW(), NOW()),
  (8, 'bn', 'মোটো জি ১০০', NULL, NOW(), NOW()),
  (9, 'bn', 'নোকিয়া এক্স ১০০', NULL, NOW(), NOW()),
  (10, 'bn', 'গুগল পিক্সেল ৯', NULL, NOW(), NOW()),
  
  -- Tablets
  (11, 'bn', 'আইপ্যাড এয়ার', NULL, NOW(), NOW()),
  (12, 'bn', 'স্যামসাং গ্যালাক্সি ট্যাব এস১০', NULL, NOW(), NOW()),
  (13, 'bn', 'লেনোভো ট্যাব এম১１', NULL, NOW(), NOW()),
  
  -- Laptops
  (14, 'bn', 'ম্যাকবুক প্রো ১৬"', NULL, NOW(), NOW()),
  (15, 'bn', 'ডেল এক্সপিএস ১৩', NULL, NOW(), NOW()),
  (16, 'bn', 'এইচপি প্যাভিলিয়ন ১৫', NULL, NOW(), NOW()),
  (17, 'bn', 'লেনোভো থিংকপ্যাড এক্স১', NULL, NOW(), NOW()),
  (18, 'bn', 'এসাস ভিভোবুক ১৫', NULL, NOW(), NOW()),
  
  -- Accessories
  (19, 'bn', 'অ্যাপল এয়ারপডস প্রো', NULL, NOW(), NOW()),
  (20, 'bn', 'সোনি ডাব্লিউএইচ-১০০০এক্সএম৫', NULL, NOW(), NOW()),
  (21, 'bn', 'স্যামসাং বাডস২', NULL, NOW(), NOW()),
  (22, 'bn', 'অ্যাপল ওয়াচ সিরিজ ৯', NULL, NOW(), NOW()),
  (23, 'bn', 'অ্যাপল পাওয়ার ব্যাংক ২০০০০এমএএইচ', NULL, NOW(), NOW()),
  (24, 'bn', 'অ্যানকার ইউএসবি-সি চার্জার', NULL, NOW(), NOW())
ON CONFLICT (product_id, locale) DO UPDATE SET
  translated_name = EXCLUDED.translated_name,
  price = EXCLUDED.price,
  updated_at = NOW();

-- Verify the insertions
SELECT COUNT(*) as "Updated Bengali Translations" 
FROM product_translations 
WHERE locale = 'bn';

-- View all Bengali translations
SELECT pt.id, p.name, pt.translated_name, pt.locale, pt.created_at
FROM product_translations pt
JOIN products p ON pt.product_id = p.id
WHERE pt.locale = 'bn'
ORDER BY pt.product_id;
