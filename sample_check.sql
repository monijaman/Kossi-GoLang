SELECT st.specification_id, st.value, s.value as english_value
FROM specification_translations st
JOIN specifications s ON st.specification_id = s.id
WHERE st.specification_id IN (26772, 26773, 26774, 26775, 26776)
AND st.locale = 'bn'
LIMIT 10;
