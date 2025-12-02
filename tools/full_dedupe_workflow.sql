-- ============================================================================
-- FULL DEDUPLICATION WORKFLOW FOR SPECIFICATIONS DATA
-- ============================================================================
-- This script deduplicates and tracks changes across three related tables:
-- 1. specification_keys (with normalization: lowercase + alphanumerics only)
-- 2. specification_key_translations (updated to use deduplicated keys)
-- 3. specifications (with deduplication + tracking)
-- 4. specification_translations (updated to use deduplicated specs)
--
-- SAFETY: All deletions are logged in *_dedupe_log tables for audit trail.
-- All operations wrapped in a transaction that you can ROLLBACK if needed.
--
-- DRY-RUN: Comment out the DELETE statements and run to see what will happen.
-- ============================================================================

BEGIN;

-- ============================================================================
-- PHASE 1: DEDUPLICATE specification_keys
-- ============================================================================

-- Create log table for specification_keys deletions
CREATE TABLE IF NOT EXISTS specification_keys_dedupe_log (
  old_id bigint NOT NULL,
  kept_id bigint NOT NULL,
  old_specification_key text NOT NULL,
  kept_specification_key text NOT NULL,
  normalized_key text NOT NULL,
  deleted_at timestamptz NOT NULL DEFAULT now()
);

-- Find and log duplicate specification_keys
WITH normalized AS (
  SELECT id, specification_key,
    regexp_replace(lower(specification_key), '[^a-z0-9]+', '', 'g') AS normalized_key
  FROM specification_keys
),
groups AS (
  SELECT normalized_key, min(id) AS kept_id, array_agg(id ORDER BY id) AS ids, count(*) AS cnt
  FROM normalized
  GROUP BY normalized_key
  HAVING COUNT(*) > 1
),
to_delete AS (
  SELECT n.id AS old_id,
         g.kept_id,
         n.specification_key AS old_specification_key,
         k.specification_key AS kept_specification_key,
         g.normalized_key
  FROM normalized n
  JOIN groups g ON n.normalized_key = g.normalized_key
  JOIN specification_keys k ON k.id = g.kept_id
  WHERE n.id <> g.kept_id
)
INSERT INTO specification_keys_dedupe_log (old_id, kept_id, old_specification_key, kept_specification_key, normalized_key)
SELECT old_id, kept_id, old_specification_key, kept_specification_key, normalized_key
FROM to_delete
ON CONFLICT DO NOTHING;

-- DRY-RUN: See what will be deleted
-- SELECT * FROM specification_keys_dedupe_log ORDER BY normalized_key, kept_id;

-- DELETE duplicates from specification_keys
DELETE FROM specification_keys WHERE id IN (SELECT old_id FROM specification_keys_dedupe_log);

-- ============================================================================
-- PHASE 2: UPDATE specification_key_translations to use kept keys
-- ============================================================================

-- Log which specification_key_translations were updated
CREATE TABLE IF NOT EXISTS specification_key_translations_update_log (
  translation_id bigint NOT NULL,
  old_specification_key_id bigint NOT NULL,
  new_specification_key_id bigint NOT NULL,
  old_key_name text,
  new_key_name text,
  updated_at timestamptz NOT NULL DEFAULT now()
);

-- Find and update specification_key_translations that reference deleted keys
INSERT INTO specification_key_translations_update_log (translation_id, old_specification_key_id, new_specification_key_id, old_key_name, new_key_name)
SELECT
  skt.id,
  skt.specification_key_id,
  l.kept_id,
  sk_old.specification_key,
  sk_new.specification_key
FROM specification_key_translations skt
JOIN specification_keys_dedupe_log l ON skt.specification_key_id = l.old_id
JOIN specification_keys sk_old ON sk_old.id = l.old_id
JOIN specification_keys sk_new ON sk_new.id = l.kept_id
ON CONFLICT DO NOTHING;

-- DRY-RUN: See what will be updated
-- SELECT * FROM specification_key_translations_update_log;

-- Update specification_key_translations to use the kept specification_key_id
UPDATE specification_key_translations skt
SET specification_key_id = l.kept_id
FROM specification_keys_dedupe_log l
WHERE skt.specification_key_id = l.old_id;

-- ============================================================================
-- PHASE 3: DEDUPLICATE specifications
-- ============================================================================

-- Create log table for specifications deletions
CREATE TABLE IF NOT EXISTS specifications_dedupe_log (
  old_id bigint NOT NULL,
  kept_id bigint NOT NULL,
  product_id bigint NOT NULL,
  specification_key_id bigint NOT NULL,
  old_value text NOT NULL,
  kept_value text NOT NULL,
  deleted_at timestamptz NOT NULL DEFAULT now()
);

-- Find and log duplicate specifications (same product_id + specification_key_id + value)
WITH groups AS (
  SELECT product_id, specification_key_id, value, min(id) AS kept_id, array_agg(id ORDER BY id) AS ids, count(*) AS cnt
  FROM specifications
  GROUP BY product_id, specification_key_id, value
  HAVING COUNT(*) > 1
),
to_delete AS (
  SELECT s.id AS old_id,
         g.kept_id,
         s.product_id,
         s.specification_key_id,
         s.value AS old_value,
         (SELECT value FROM specifications WHERE id = g.kept_id LIMIT 1) AS kept_value
  FROM specifications s
  JOIN groups g ON s.product_id = g.product_id
                AND s.specification_key_id = g.specification_key_id
                AND s.value = g.value
  WHERE s.id <> g.kept_id
)
INSERT INTO specifications_dedupe_log (old_id, kept_id, product_id, specification_key_id, old_value, kept_value)
SELECT old_id, kept_id, product_id, specification_key_id, old_value, kept_value
FROM to_delete
ON CONFLICT DO NOTHING;

-- DRY-RUN: See what will be deleted
-- SELECT * FROM specifications_dedupe_log ORDER BY product_id, specification_key_id;

-- DELETE duplicates from specifications
DELETE FROM specifications WHERE id IN (SELECT old_id FROM specifications_dedupe_log);

-- ============================================================================
-- PHASE 4: UPDATE specification_translations to use kept specifications
-- ============================================================================

-- Log which specification_translations were updated
CREATE TABLE IF NOT EXISTS specification_translations_update_log (
  translation_id bigint NOT NULL,
  old_specification_id bigint NOT NULL,
  new_specification_id bigint NOT NULL,
  updated_at timestamptz NOT NULL DEFAULT now()
);

-- Find and update specification_translations that reference deleted specifications
INSERT INTO specification_translations_update_log (translation_id, old_specification_id, new_specification_id)
SELECT
  st.id,
  st.specification_id,
  l.kept_id
FROM specification_translations st
JOIN specifications_dedupe_log l ON st.specification_id = l.old_id
ON CONFLICT DO NOTHING;

-- DRY-RUN: See what will be updated
-- SELECT * FROM specification_translations_update_log;

-- Update specification_translations to use the kept specification_id
UPDATE specification_translations st
SET specification_id = l.kept_id
FROM specifications_dedupe_log l
WHERE st.specification_id = l.old_id;


-- ============================================================================
-- SUMMARY & VERIFICATION
-- ============================================================================

SELECT '========== DEDUPE COMPLETE ==========' AS status;

SELECT 'specification_keys_dedupe_log' AS table_name, COUNT(*) AS count FROM specification_keys_dedupe_log
UNION ALL
SELECT 'specification_key_translations_update_log', COUNT(*) FROM specification_key_translations_update_log
UNION ALL
SELECT 'specifications_dedupe_log', COUNT(*) FROM specifications_dedupe_log
UNION ALL
SELECT 'specification_translations_update_log', COUNT(*) FROM specification_translations_update_log;

COMMIT;

-- ============================================================================
-- POST-RUN VERIFICATION QUERIES (run manually after script completes)
-- ============================================================================
-- Check for remaining duplicates in specification_keys:
-- SELECT specification_key, COUNT(*) AS cnt
-- FROM specification_keys
-- GROUP BY specification_key
-- HAVING COUNT(*) > 1;
--
-- Check for remaining duplicates in specifications:
-- SELECT product_id, specification_key_id, value, COUNT(*) AS cnt
-- FROM specifications
-- GROUP BY product_id, specification_key_id, value
-- HAVING COUNT(*) > 1;
--
-- Review the audit logs:
-- SELECT * FROM specification_keys_dedupe_log ORDER BY kept_id;
-- SELECT * FROM specification_key_translations_update_log ORDER BY translation_id;
-- SELECT * FROM specifications_dedupe_log ORDER BY kept_id;
-- SELECT * FROM specification_translations_update_log ORDER BY translation_id;
--
-- If you need to ROLLBACK everything, manually run in a new transaction:
-- BEGIN;
-- -- Restore specification_keys from log (insert deleted rows back)
-- INSERT INTO specification_keys (id, specification_key, created_at, updated_at)
-- SELECT old_id, old_specification_key, now(), now()
-- FROM specification_keys_dedupe_log;
-- -- Revert specification_key_translations updates
-- UPDATE specification_key_translations skt
-- SET specification_key_id = l.old_id
-- FROM specification_key_translations_update_log l
-- WHERE skt.id = l.translation_id;
-- -- Restore specifications from log
-- INSERT INTO specifications (id, product_id, specification_key_id, value, status, created_at, updated_at)
-- SELECT old_id, product_id, specification_key_id, old_value, 1, now(), now()
-- FROM specifications_dedupe_log;
-- -- Revert specification_translations updates
-- UPDATE specification_translations st
-- SET specification_id = l.old_specification_id
-- FROM specification_translations_update_log l
-- WHERE st.id = l.translation_id;
-- COMMIT;
--
-- ============================================================================
