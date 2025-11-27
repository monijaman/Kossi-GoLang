-- Get product basic info and translations for product ID 499
SELECT 
    p.id AS product_id,
    p.name AS product_name,
    p.slug AS product_slug,
    pt.locale,
    pt.name AS translated_name,
    pt.description AS translated_description
FROM 
    products p
LEFT JOIN 
    product_translations pt ON p.id = pt.product_id
WHERE 
    p.id = 499
ORDER BY 
    pt.locale;

-- Get specifications with Bengali translations for product ID 499
SELECT 
    p.id AS product_id,
    p.name AS product_name,
    s.id AS spec_id,
    sk.specification_key,
    s.value AS spec_value_english,
    st.locale AS translation_locale,
    st.value AS spec_value_translated
FROM 
    products p
LEFT JOIN 
    specifications s ON p.id = s.product_id
LEFT JOIN 
    specification_keys sk ON s.specification_key_id = sk.id
LEFT JOIN 
    specification_translations st ON s.id = st.specification_id
WHERE 
    p.id = 499
ORDER BY 
    sk.specification_key, st.locale;

-- Count specifications and translations for product ID 499
SELECT 
    p.id,
    p.name,
    COUNT(DISTINCT s.id) AS total_specifications,
    COUNT(DISTINCT CASE WHEN st.locale = 'bn' THEN st.id END) AS bengali_translations,
    COUNT(DISTINCT CASE WHEN st.locale = 'hi' THEN st.id END) AS hindi_translations
FROM 
    products p
LEFT JOIN 
    specifications s ON p.id = s.product_id
LEFT JOIN 
    specification_translations st ON s.id = st.specification_id
WHERE 
    p.id = 499
GROUP BY 
    p.id, p.name;
