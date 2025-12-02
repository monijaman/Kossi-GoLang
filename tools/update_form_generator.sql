-- Update form_generator to use deduplicated specification_ids
-- The specification_id column is JSON (array of IDs), so we need to handle it as JSON
-- Map each old specification_id to its kept_id using the dedupe log

SELECT '=== Updating form_generator JSON specification_ids ===' AS info;

-- Simply update all form_generator rows by mapping old spec IDs to kept IDs
UPDATE form_generator fg
SET specification_id = (
  SELECT json_agg(COALESCE(m.kept_id, spec_id::bigint))
  FROM (SELECT json_array_elements(fg.specification_id)::text::bigint AS spec_id) specs
  LEFT JOIN specifications_dedupe_log m ON m.old_id = specs.spec_id
);

SELECT '=== Form generator update complete ===' AS info;

-- Show updated form_generator rows
SELECT 
  id,
  category_id,
  specification_id,
  updated_at
FROM form_generator
ORDER BY updated_at DESC
LIMIT 10;
