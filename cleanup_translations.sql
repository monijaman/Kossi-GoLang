-- Remove "অনুবাদ প্রয়োজন" phrase from all Bengali translations

-- Count translations with the phrase before update
SELECT COUNT(*) as "Translations with phrase" FROM product_translations 
WHERE locale='bn' AND translated_name LIKE '%(অনুবাদ প্রয়োজন)%';

-- Remove the phrase from all Bengali translations
UPDATE product_translations 
SET translated_name = TRIM(REGEXP_REPLACE(translated_name, '\s*\(অনুবাদ প্রয়োজন\)', ''))
WHERE locale='bn' AND translated_name LIKE '%(অনুবাদ প্রয়োজন)%';

-- Verify: Count after update
SELECT COUNT(*) as "Updated translations" FROM product_translations 
WHERE locale='bn';

-- Show sample of cleaned translations
SELECT id, product_id, translated_name FROM product_translations 
WHERE locale='bn' 
LIMIT 30;
