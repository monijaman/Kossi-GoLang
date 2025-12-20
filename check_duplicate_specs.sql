-- Query to find unused and duplicate specification keys
-- Check which specification keys have actual data in the specifications table

-- 1. Find all specification keys with no actual values
SELECT sk.id, sk.specification_key, COUNT(s.id) as usage_count
FROM specification_keys sk
LEFT JOIN specifications s ON sk.id = s.specification_key_id
WHERE sk.specification_key IN (
    'Warranty Period',
    'Warranty Period (Years)',
    'Warranty Coverage',
    '2G Service Available',
    '3G Service Available',
    '4G/LTE Service Available',
    '5G Service Available',
    'fast_charging',
    'Charging Speed'
)
GROUP BY sk.id, sk.specification_key
ORDER BY sk.id;

-- 2. Show duplicate meaning groups and their usage
SELECT 
    'Warranty Group' as category,
    'Warranty Period' as key1, 'Warranty Period (Years)' as key2, 'Warranty Coverage' as key3
UNION ALL
SELECT 
    'Service Available Group',
    '2G Service Available',
    '3G Service Available',
    '5G Service Available'
UNION ALL
SELECT 
    'Charging Group',
    'fast_charging',
    'Charging Speed',
    NULL;

-- 3. Check specifications table for each duplicate
SELECT 
    sk.specification_key,
    COUNT(s.id) as total_values,
    COUNT(DISTINCT s.product_id) as product_count
FROM specification_keys sk
LEFT JOIN specifications s ON sk.id = s.specification_key_id
WHERE sk.id IN (161, 509, 510, 453, 454, 456, 553, 18)
GROUP BY sk.id, sk.specification_key
ORDER BY total_values DESC;
