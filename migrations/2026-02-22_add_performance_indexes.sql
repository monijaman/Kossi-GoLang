-- Performance optimization indexes for product filtering
-- This migration adds indexes to support fast category and brand filtering

-- Index on products.category_id for fast filtering by category
CREATE INDEX IF NOT EXISTS idx_products_category_id ON products(category_id) WHERE deleted_at IS NULL;

-- Index on products.brand_id for fast filtering by brand  
CREATE INDEX IF NOT EXISTS idx_products_brand_id ON products(brand_id) WHERE deleted_at IS NULL;

-- Index on products.status for active status filtering
CREATE INDEX IF NOT EXISTS idx_products_status ON products(status) WHERE deleted_at IS NULL;

-- Index on products.priority and views_count for sorting (composite index)
CREATE INDEX IF NOT EXISTS idx_products_priority_views ON products(priority ASC, views_count DESC) WHERE deleted_at IS NULL;

-- Index on products.name for ILIKE search optimization
CREATE INDEX IF NOT EXISTS idx_products_name_trgm ON products USING gin(name gin_trgm_ops);

-- Index on categories.slug for fast slug lookups
CREATE INDEX IF NOT EXISTS idx_categories_slug ON categories(slug) WHERE deleted_at IS NULL;

-- Index on brands.slug for fast slug lookups  
CREATE INDEX IF NOT EXISTS idx_brands_slug ON brands(slug) WHERE deleted_at IS NULL;

-- Composite index for products.start_price and end_price for price filtering
CREATE INDEX IF NOT EXISTS idx_products_price_range ON products(start_price, end_price) WHERE deleted_at IS NULL;

-- Index on images.imageable_id and imageable_type for batch image lookups
CREATE INDEX IF NOT EXISTS idx_images_imageable ON images(imageable_type, imageable_id);

-- Note: The gin_trgm_ops index for text search requires the pg_trgm extension
-- If not already installed, run: CREATE EXTENSION IF NOT EXISTS pg_trgm;
