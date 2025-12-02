-- Verify deleted specification_keys
SELECT '=== DELETED specification_keys (from audit log) ===' AS info;
SELECT old_id, kept_id, old_specification_key, kept_specification_key, normalized_key 
FROM specification_keys_dedupe_log 
ORDER BY normalized_key, kept_id;

-- Verify they're actually gone from the live table
SELECT '=== Verifying deleted IDs are gone from specification_keys ===' AS info;
SELECT old_id, old_specification_key
FROM specification_keys_dedupe_log d
WHERE EXISTS (SELECT 1 FROM specification_keys sk WHERE sk.id = d.old_id)
ORDER BY old_id;

-- Summary
SELECT '=== SUMMARY ===' AS info;
SELECT 
  (SELECT COUNT(*) FROM specification_keys_dedupe_log) AS deleted_count,
  (SELECT COUNT(*) FROM specification_keys) AS remaining_keys_count,
  (SELECT COUNT(DISTINCT specification_key) FROM specification_keys) AS unique_keys_count;
