-- =============================================================================
-- CAMERA SPECIFICATION KEYS - DUPLICATE ANALYSIS
-- =============================================================================

-- CAMERA RELATED KEYS SUMMARY:
-- =============================================================================

-- 1. USED CAMERA SPECIFICATIONS:
--    ID  16: 'Camera Features'                   - 25 products ✅
--    ID  32: 'Front Camera Video Resolution'     - 26 products ✅
--    ID  40: 'Main Camera Video Resolution'      - 1 product ✅
--    ID  58: 'Quad Camera Setup'                 - 25 products ✅
--    ID 122: 'Rear Camera'                       - 26 products ✅

-- 2. UNUSED CAMERA SPECIFICATIONS (DELETE THESE):
--    ID  28: 'Dual or Triple Camera Setup'       - 0 products ❌
--    ID  31: 'Front Camera Resolution'           - 0 products ❌
--    ID  39: 'Main Camera Resolution'            - 0 products ❌
--    ID  50: 'Penta Camera Setup'                - 0 products ❌
--    ID  75: 'Triple Camera Setup'               - 0 products ❌

-- 3. UNUSED NEWER KEYS (LOWERCASE FORMAT):
--    ID 547: 'camera_video_resolution'           - 0 products ❌
--    ID 549: 'front_camera'                      - 0 products ❌
--    ID 601: 'Camera'                            - 0 products ❌
--    ID 650: 'Camera Video Resolution'           - 0 products ❌
--    ID 651: 'Front Camera'                      - 0 products ❌

-- =============================================================================
-- DUPLICATE GROUPS IDENTIFIED:
-- =============================================================================

-- GROUP 1: CAMERA SETUP (MULTIPLE SETUPS)
--   Used: ID 58 'Quad Camera Setup' (25 products)
--   Unused: ID 28 'Dual or Triple Camera Setup' (0)
--   Unused: ID 50 'Penta Camera Setup' (0)
--   Unused: ID 75 'Triple Camera Setup' (0)
--   NOTE: Only Quad Camera is in use; others are completely unused

-- GROUP 2: FRONT CAMERA SPECIFICATION
--   Used: ID 32 'Front Camera Video Resolution' (26 products)
--   Unused: ID 31 'Front Camera Resolution' (0)
--   Unused: ID 549 'front_camera' (0)
--   Unused: ID 651 'Front Camera' (0)
--   NOTE: Only Video Resolution is used; plain Resolution and lowercase versions unused

-- GROUP 3: MAIN CAMERA SPECIFICATION
--   Used: ID 40 'Main Camera Video Resolution' (1 product)
--   Unused: ID 39 'Main Camera Resolution' (0)
--   NOTE: Minimal usage (1 product only)

-- GROUP 4: GENERIC CAMERA KEYS
--   Unused: ID 547 'camera_video_resolution' (0) - lowercase duplicate
--   Unused: ID 601 'Camera' (0) - too generic
--   Unused: ID 650 'Camera Video Resolution' (0) - different from Main/Front versions

-- =============================================================================
-- CLEANUP PLAN
-- =============================================================================

-- DELETE COMPLETELY UNUSED CAMERA KEYS (10 total):
DELETE FROM specification_keys WHERE id IN (
  28,    -- 'Dual or Triple Camera Setup'
  31,    -- 'Front Camera Resolution'
  39,    -- 'Main Camera Resolution'
  50,    -- 'Penta Camera Setup'
  75,    -- 'Triple Camera Setup'
  547,   -- 'camera_video_resolution'
  549,   -- 'front_camera'
  601,   -- 'Camera'
  650,   -- 'Camera Video Resolution'
  651    -- 'Front Camera'
);

-- KEEP THESE (IN USE):
--   ID  16: 'Camera Features' (25 products)
--   ID  32: 'Front Camera Video Resolution' (26 products)
--   ID  40: 'Main Camera Video Resolution' (1 product)
--   ID  58: 'Quad Camera Setup' (25 products)
--   ID 122: 'Rear Camera' (26 products)

-- =============================================================================
-- SUMMARY:
-- To Delete: 10 specification keys (all unused)
-- To Keep: 5 specification keys (all active)
-- Expected Result: 558 - 10 = 548 total keys
-- =============================================================================
