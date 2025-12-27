-- Insert Panasonic brand if not exists
INSERT INTO brands (name, slug, status, created_at, updated_at)
VALUES ('Panasonic', 'panasonic', 1, NOW(), NOW())
ON CONFLICT (slug) DO NOTHING;

-- Insert brand translations if not exists
INSERT INTO brand_translations (brand_id, locale, name, created_at, updated_at)
SELECT b.id, 'en', 'Panasonic', NOW(), NOW()
FROM brands b
WHERE b.slug = 'panasonic'
ON CONFLICT (brand_id, locale) DO NOTHING;

INSERT INTO brand_translations (brand_id, locale, name, created_at, updated_at)
SELECT b.id, 'bn', 'প্যানাসনিক', NOW(), NOW()
FROM brands b
WHERE b.slug = 'panasonic'
ON CONFLICT (brand_id, locale) DO NOTHING;

-- Insert brand category if not exists
INSERT INTO brand_category (brand_id, category_id)
SELECT b.id, c.id
FROM brands b
CROSS JOIN categories c
WHERE b.slug = 'panasonic' AND c.slug = 'refrigerator'
ON CONFLICT (brand_id, category_id) DO NOTHING;

-- Insert Panasonic refrigerator products if not exists
INSERT INTO products (name, slug, category_id, brand_id, model, status, priority, created_at, updated_at)
SELECT 'Panasonic ' || unnest(ARRAY[
    'CH201H7B',
    'SCR-CH200',
    'NR-AF172SNAE',
    'NR-CY550GKXZ',
    'NR-CY550QKXZ',
    'NR-BX468XGX3',
    'NR-DZ600GKXZ',
    'NR-BN221SNW',
    'NR-BK305SNWA',
    'NR-BK305SAWS',
    'NR-BK265SNWA',
    'NR-BK703',
    'NR-BK266SNWA',
    'NR-BK221',
    'NR-AH186RAWA',
    'NR-AE51SHSG',
    'NR-BU303L',
    'NR-BL267PSWA'
]), 'panasonic-' || lower(replace(unnest(ARRAY[
    'CH201H7B',
    'SCR-CH200',
    'NR-AF172SNAE',
    'NR-CY550GKXZ',
    'NR-CY550QKXZ',
    'NR-BX468XGX3',
    'NR-DZ600GKXZ',
    'NR-BN221SNW',
    'NR-BK305SNWA',
    'NR-BK305SAWS',
    'NR-BK265SNWA',
    'NR-BK703',
    'NR-BK266SNWA',
    'NR-BK221',
    'NR-AH186RAWA',
    'NR-AE51SHSG',
    'NR-BU303L',
    'NR-BL267PSWA'
]), '-', '')), c.id, b.id, unnest(ARRAY[
    'CH201H7B',
    'SCR-CH200',
    'NR-AF172SNAE',
    'NR-CY550GKXZ',
    'NR-CY550QKXZ',
    'NR-BX468XGX3',
    'NR-DZ600GKXZ',
    'NR-BN221SNW',
    'NR-BK305SNWA',
    'NR-BK305SAWS',
    'NR-BK265SNWA',
    'NR-BK703',
    'NR-BK266SNWA',
    'NR-BK221',
    'NR-AH186RAWA',
    'NR-AE51SHSG',
    'NR-BU303L',
    'NR-BL267PSWA'
]), 1, 0, NOW(), NOW()
FROM brands b
CROSS JOIN categories c
WHERE b.slug = 'panasonic' AND c.slug = 'refrigerator'
ON CONFLICT (slug) DO NOTHING;