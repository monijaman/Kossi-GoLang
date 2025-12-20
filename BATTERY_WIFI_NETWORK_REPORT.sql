-- =============================================================================
-- BATTERY, WiFi, NETWORK TECHNOLOGY - DUPLICATE ANALYSIS
-- =============================================================================

-- BATTERY SPECIFICATIONS FOUND:
-- =============================================================================

-- USED BATTERY KEYS:
--   ID  11: 'Battery Capacity'       - 166 products ✅
--   ID  12: 'Battery Type'           - 169 products ✅
--   ID 374: 'Battery Voltage'        -   1 product ✅
--   ID 553: 'fast_charging'          - 136 products ✅

-- UNUSED BATTERY KEYS (DELETE THESE):
--   ID  45: 'New Battery Capacity'   -   0 products ❌
--   ID  48: 'Old Battery Capacity'   -   0 products ❌
--   ID 303: 'Battery Required'       -   0 products ❌

-- NOT IN PRODUCTION DB (only in SQL dump):
--   ID 551: 'battery' (lowercase)    - Does not exist in production
--   ID 659: 'Battery'                - Does not exist in production

-- =============================================================================
-- WiFi SPECIFICATIONS:
-- =============================================================================

-- UNUSED WiFi KEY:
--   ID  81: 'WiFi'                   -   0 products ❌

-- NOT IN PRODUCTION DB (only in SQL dump):
--   ID 656: 'Wifi Support'           - Does not exist in production

-- =============================================================================
-- NETWORK TECHNOLOGY:
-- =============================================================================

-- USED:
--   ID 653: 'Network Technology'     -   1 product ✅

-- USED (RELATED):
--   ID 561: 'network_technology'     - Not found in DB query
--   ID 476: 'Network Speed (Average)' - Telecom specific

-- =============================================================================
-- CLEANUP PLAN - DELETE THESE UNUSED KEYS:
-- =============================================================================

DELETE FROM specification_keys WHERE id IN (
  45,    -- 'New Battery Capacity'
  48,    -- 'Old Battery Capacity'
  81,    -- 'WiFi'
  303    -- 'Battery Required'
);

-- Note: IDs 551, 656, 659 don't exist in production, only in SQL dump
-- These are in the SQL file but were never created in the database

-- =============================================================================
-- SUMMARY OF DUPLICATES:
-- =============================================================================

-- BATTERY GROUP:
--   Canonical keys: ID 11 'Battery Capacity' (166), ID 12 'Battery Type' (169)
--   Keep: ID 374 'Battery Voltage' (1 product, specific use)
--   Keep: ID 553 'fast_charging' (136, different concept)
--   Delete: ID 45, 48, 303 (unused)

-- WiFi GROUP:
--   Used: None (ID 81 has 0 usage)
--   Delete: ID 81 'WiFi'

-- NETWORK TECHNOLOGY:
--   Keep: ID 653 'Network Technology' (1 product)
--   Related but separate: Network Speed, Network Type, etc. for telecom

-- Total to delete: 4 keys (45, 48, 81, 303)
-- Expected result: 553 - 4 = 549 total keys
