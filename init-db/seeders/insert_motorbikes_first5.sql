-- insert_motorbikes_first5.sql
-- Seed the products table with the first 5 motorbikes from motorbikes.json
-- Assumptions:
--  - `products` table matches the GORM ProductModel in the project (columns: name, description, slug, model, price, status, priority, created_at, updated_at, ...)
--  - brand_id and category_id are intentionally left NULL here; if you have those ids available, set them explicitly.
--  - This file inserts 5 rows. Run inside a transaction or with psql to seed your DB.

BEGIN;

INSERT INTO products (name, description, slug, model, price, status, priority, created_at, updated_at)
VALUES
  ('Honda CD 70', NULL, 'honda-cd-70', 'CD 70', 85000.00, 1, 1, NOW(), NOW()),
  ('Honda Livo', NULL, 'honda-livo', 'Livo', 150000.00, 1, 1, NOW(), NOW()),
  ('Honda CB125F', NULL, 'honda-cb125f', 'CB125F', 155000.00, 1, 1, NOW(), NOW()),
  ('Honda CB Shine', NULL, 'honda-cb-shine', 'CB Shine', 165000.00, 1, 1, NOW(), NOW()),
  ('Honda Dio', NULL, 'honda-dio', 'Dio', 140000.00, 1, 1, NOW(), NOW());

COMMIT;

-- Usage:
-- psql -h <host> -U <user> -d <db> -f init-db/seeders/insert_motorbikes_first5.sql
