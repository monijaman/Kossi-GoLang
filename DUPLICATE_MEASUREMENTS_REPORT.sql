-- =============================================================================
-- DUPLICATE SPECIFICATION KEYS - DETAILED ANALYSIS
-- =============================================================================

-- FINDINGS: Clear duplicates with different naming conventions
-- =============================================================================

-- GROUP 1: DIMENSIONS AND MEASUREMENTS
-- GENERIC (Used): ID 113 'Length' (165 products)
-- SPECIFIC (Unused): ID 612 'Length (mm)' (0 products) ❌ DELETE
--
-- GENERIC (Used): ID 136 'Width' (165 products)
-- SPECIFIC (Unused): ID 586 'Width (mm)' (0 products) ❌ DELETE
--
-- GENERIC (Used): ID 25 'Dimensions' (26 products)
-- SPECIFIC (Unused): ID 506 'Weight Without Stand (kg)' (0 products) ❌ DELETE
-- SPECIFIC (Unused): ID 507 'Weight With Stand (kg)' (0 products) ❌ DELETE
--
-- GENERIC (Used): ID 80 'Weight' (26 products)
-- SPECIFIC (Unused): ID 320 'Product Weight' (0 products) ❌ DELETE

-- GROUP 2: POWER SPECIFICATIONS
-- GENERIC (Used): ID 114 'Max Power' (165 products)
-- SPECIFIC (Unused): ID 585 'Max Power (hp)' (0 products) ❌ DELETE

-- GROUP 3: WHEELBASE (Vehicle)
-- GENERIC (Used): ID 135 'Wheelbase' (165 products)
-- SPECIFIC (Unused): ID 599 'Wheelbase (mm)' (0 products) ❌ DELETE

-- GROUP 4: PRICE
-- GENERIC (Used): ID 56 'Price' (0 products) ← Interesting: neither used
-- SPECIFIC (Used): ID 526 'Price (BDT)' (25 products) ✅ KEEP (more specific)
-- ACTION: Keep ID 526, can delete ID 56 if needed

-- GROUP 5: SAR (Specific Absorption Rate)
-- BOTH UNUSED:
-- ID 63: 'SAR (Specific Absorption Rate)' (0 products) ❌ DELETE
-- ID 64: 'SAR (Specific Absorption Rate) EU' (0 products) ❌ DELETE

-- GROUP 6: NETWORK SERVICES (Different purpose)
-- ID 453: '2G Service Available' (5 products) ✅ KEEP
-- ID 455: '4G/LTE Service Available' (5 products) ✅ KEEP
-- These are telecom operator services, NOT duplicates

-- GROUP 7: ABS TYPE (Vehicle)
-- ID 361: 'ABS Type' (9 products) ✅ KEEP
-- This is used, not a duplicate

-- =============================================================================
-- CLEANUP PLAN - DELETE THESE UNUSED KEYS:
-- =============================================================================

DELETE FROM specification_keys WHERE id IN (
  56,    -- 'Price' (unused, superseded by ID 526 'Price (BDT)')
  63,    -- 'SAR (Specific Absorption Rate)' (unused)
  64,    -- 'SAR (Specific Absorption Rate) EU' (unused)
  320,   -- 'Product Weight' (duplicate of ID 80 'Weight')
  506,   -- 'Weight Without Stand (kg)' (duplicate of ID 80 'Weight')
  507,   -- 'Weight With Stand (kg)' (duplicate of ID 25 'Dimensions')
  585,   -- 'Max Power (hp)' (duplicate of ID 114 'Max Power')
  586,   -- 'Width (mm)' (duplicate of ID 136 'Width')
  599,   -- 'Wheelbase (mm)' (duplicate of ID 135 'Wheelbase')
  612    -- 'Length (mm)' (duplicate of ID 113 'Length')
);

-- =============================================================================
-- SUMMARY:
-- =============================================================================

-- DUPLICATES IDENTIFIED: 10 keys
-- All are UNUSED or REDUNDANT with better alternatives
-- Generic versions (without units) are preferred over specific ones
-- OR specific meaningful ones (like 'Price (BDT)') are preferred over generic

-- Keys to KEEP (in use or better names):
--   ID 25: 'Dimensions' (26 products)
--   ID 56: 'Price' (0 products - but could delete)
--   ID 80: 'Weight' (26 products)
--   ID 113: 'Length' (165 products)
--   ID 114: 'Max Power' (165 products)
--   ID 135: 'Wheelbase' (165 products)
--   ID 136: 'Width' (165 products)
--   ID 361: 'ABS Type' (9 products)
--   ID 453: '2G Service Available' (5 products)
--   ID 455: '4G/LTE Service Available' (5 products)
--   ID 526: 'Price (BDT)' (25 products) ← PREFERRED over Price

-- Total to delete: 10 keys (10 more than before)
-- Expected result: 549 - 10 = 539 total keys
