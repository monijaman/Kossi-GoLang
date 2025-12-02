-- Dedupe `specification_keys` by a normalized key
-- Normalization: lower(specification_key) with non-alphanumerics removed
-- This script will:
-- 1) Create a log table `specification_keys_dedupe_log` if it doesn't exist
-- 2) Find groups of keys that normalize to the same value
-- 3) For each group, keep the row with the smallest `id` and mark other rows for deletion
-- 4) Insert deleted-row info into the log table (old_id -> kept_id mapping)
-- 5) Delete the duplicate rows from `specification_keys`

-- IMPORTANT: Run the SELECT parts first (comment out the DELETE) to verify results before committing.

BEGIN;

-- 1) Create a log table to record what we delete (if it doesn't exist)
CREATE TABLE IF NOT EXISTS specification_keys_dedupe_log (
  old_id bigint NOT NULL,
  kept_id bigint NOT NULL,
  old_specification_key text NOT NULL,
  kept_specification_key text NOT NULL,
  normalized_key text NOT NULL,
  deleted_at timestamptz NOT NULL DEFAULT now()
);

-- 2) Build normalized view and find groups
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

-- 3) Insert what will be deleted into the log table
INSERT INTO specification_keys_dedupe_log (old_id, kept_id, old_specification_key, kept_specification_key, normalized_key)
SELECT old_id, kept_id, old_specification_key, kept_specification_key, normalized_key
FROM to_delete
ON CONFLICT DO NOTHING;

-- 4) -- VERIFY: list what will be deleted (run this manually or leave to see output)
-- SELECT * FROM to_delete ORDER BY normalized_key, kept_id, old_id;

-- 5) Delete duplicates (UNCOMMENT to perform deletion after verification)
DELETE FROM specification_keys WHERE id IN (SELECT old_id FROM specification_keys_dedupe_log);

COMMIT;

-- After running this, you can inspect `specification_keys_dedupe_log` to see mapping from deleted ids -> kept_id
-- Next steps (not included here):
--   - Update `specification_key_translations` to point any rows that referenced deleted specification_key ids to the `kept_id`.
--   - Update `specifications` and `specification_translations` if necessary (similar mapping logic may be required there).
