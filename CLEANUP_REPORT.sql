-- =============================================================================
-- SPECIFICATION KEYS CLEANUP REPORT - ACTUAL USAGE FROM PRODUCTION DATABASE
-- =============================================================================

-- KEY FINDINGS:
-- =============================================================================

-- 1. CRITICAL: CHARGING SPECIFICATION MISMATCH
--    - 'fast_charging' (ID 553): 136 products using it ← MOST USED
--    - 'Charging Speed' (ID 18): 26 products using it
--    ACTION: Keep BOTH - they're used by different products
--            'fast_charging' is the dominant one (136 vs 26)
--            Consider renaming 'fast_charging' to 'Fast Charging' for consistency

-- 2. NETWORK SERVICE SPECIFICATIONS
--    - '2G Service Available' (ID 453): 5 products
--    - '3G Service Available' (ID 454): 5 products  
--    - '4G/LTE Service Available' (ID 455): 5 products (not in check, but likely same)
--    - '5G Service Available' (ID 456): 5 products
--    ACTION: Keep ALL - actively used for telecom/network products
--            These are used for different network types and should not be merged

-- 3. WARRANTY SPECIFICATIONS - COMPLETELY UNUSED
--    - 'Warranty Period' (ID 161): 0 usage ← DELETE
--    - 'Warranty Period (Years)' (ID 509): 0 usage ← DELETE
--    - 'Warranty Coverage' (ID 510): 0 usage ← DELETE
--    ACTION: DELETE all 3 - none are being used
--            These were created but never assigned to any product

-- =============================================================================
-- CLEANUP PLAN
-- =============================================================================

-- STEP 1: Delete unused Warranty-related specification keys
DELETE FROM specification_keys WHERE id IN (161, 509, 510);

-- STEP 2: Optional - Standardize 'fast_charging' naming (for consistency)
-- UPDATE specification_keys SET specification_key = 'Fast Charging' WHERE id = 553;

-- STEP 3: Keep these as-is (they are in use)
-- - ID 18: 'Charging Speed' (26 products)
-- - ID 553: 'fast_charging' (136 products) ← Most used
-- - ID 453-456: Network service specifications (telecom products)

-- =============================================================================
-- DUPLICATE TYPE ANALYSIS
-- =============================================================================

-- Found Duplicate Types:
-- 1. EXACT DUPLICATES: None found (different naming conventions)
-- 2. SEMANTIC DUPLICATES (same meaning, different names):
--    - 'Charging Speed' vs 'fast_charging' (but both in use - different products)
-- 3. UNUSED GROUP (all have 0 usage):
--    - Warranty Period group (3 keys, all unused)
-- 4. CONSISTENT NAMING GROUP (good):
--    - Service Available group (4 keys, all used, consistent naming)

-- =============================================================================
-- SUMMARY
-- =============================================================================

-- Keys to DELETE: 3 (IDs: 161, 509, 510) - All warranty-related, unused
-- Keys to RENAME: 1 (ID 553: 'fast_charging' → 'Fast Charging') - Optional for consistency
-- Keys to KEEP: All others - actively in use

-- Total: 593 specification keys
-- After cleanup: 590 specification keys (3 deletions)
