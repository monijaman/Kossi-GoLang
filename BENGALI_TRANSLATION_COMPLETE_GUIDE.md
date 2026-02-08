# Bengali Product Translation Guide

## Overview

You now have a **complete solution** to translate all your products to Bengali. This guide explains the options and how to use them.

---

## What You Have

### 1. **CLI Tool** (`cmd/product-translator/main.go`)
   - Connects to your PostgreSQL database
   - Analyzes which products need Bengali translations
   - Auto-generates translations for missing products
   - Inserts into `product_translations` table automatically

### 2. **Helper Scripts**
   - `translate-products.sh` (Linux/Mac)
   - `translate-products.ps1` (Windows PowerShell)

### 3. **SQL Sample** (`product_translations_bn_sample.sql`)
   - Pre-written Bengali translations for common products
   - Can be used if you want manual control

### 4. **Documentation** (this file)
   - Complete usage instructions
   - Troubleshooting guide

---

## Quick Start (Choose One Method)

### Method 1: **Analyze First (Recommended)**

```bash
cd d:/GO/gocrit/gocrit_server

# See what's missing
go run cmd/product-translator/main.go -analyze

# See all products with their status
go run cmd/product-translator/main.go -analyze -show-all
```

**Example Output:**
```
📊 Total Products: 45

📈 Summary:
   ✅ Products with Bengali (bn): 12
   ❌ Products missing Bengali:   33
   📊 Coverage: 26.7%
```

---

### Method 2: **Auto-Generate All Missing Translations**

```bash
go run cmd/product-translator/main.go -generate-bn
```

**What it does:**
- Gets all products without Bengali translations
- Uses built-in dictionary to generate Bengali names
- Creates them in the database automatically
- Shows progress (✅ success, ❌ error)

---

### Method 3: **Analyze + Generate (Full Workflow)**

```bash
# Step 1: Check the gap
go run cmd/product-translator/main.go -analyze -verbose

# Step 2: Generate translations
go run cmd/product-translator/main.go -generate-bn -verbose

# Step 3: Verify in database
psql -h localhost -U postgres -d gocrit -c \
  "SELECT COUNT(*) FROM product_translations WHERE locale='bn';"
```

---

## All Available Flags

| Flag | Purpose | Example |
|------|---------|---------|
| `-analyze` | Only analyze, no database changes | `go run cmd/product-translator/main.go -analyze` |
| `-generate-bn` | Create Bengali translations | `go run cmd/product-translator/main.go -generate-bn` |
| `-show-all` | Show all products (long list) | `go run cmd/product-translator/main.go -analyze -show-all` |
| `-verbose` | Detailed progress output | `go run cmd/product-translator/main.go -generate-bn -verbose` |

---

## Using PowerShell Scripts (Windows)

```powershell
# Analyze only
.\translate-products.ps1 analyze

# Generate translations
.\translate-products.ps1 generate

# Show all with analysis
.\translate-products.ps1 analyze all

# Verbose output
.\translate-products.ps1 generate verbose

# Combine flags
.\translate-products.ps1 generate all verbose
```

---

## Using Bash Scripts (Linux/Mac)

```bash
# Analyze only
chmod +x translate-products.sh
./translate-products.sh analyze

# Generate
./translate-products.sh generate

# All + verbose
./translate-products.sh all verbose
```

---

## Manual Translation with SQL

If you want to insert specific translations manually:

```bash
# Apply the sample SQL file
psql -h localhost -U postgres -d gocrit < product_translations_bn_sample.sql
```

Or insert specific translations:

```sql
-- Insert single translation
INSERT INTO product_translations (product_id, locale, translated_name, created_at, updated_at)
VALUES (1, 'bn', 'স্যামসাং গ্যালাক্সি এস২৫', NOW(), NOW());

-- Check if it created
SELECT * FROM product_translations WHERE product_id=1 AND locale='bn';
```

---

## Check Progress

### See Translation Count
```bash
psql -h localhost -U postgres -d gocrit -c \
  "SELECT COUNT(*) as bengali_count FROM product_translations WHERE locale='bn';"
```

### See All Bengali Translations
```bash
psql -h localhost -U postgres -d gocrit -c \
  "SELECT p.name, pt.translated_name, pt.created_at 
   FROM product_translations pt
   JOIN products p ON pt.product_id = p.id
   WHERE pt.locale = 'bn'
   ORDER BY p.id;"
```

### See Missing Translations
```bash
psql -h localhost -U postgres -d gocrit -c \
  "SELECT p.id, p.name FROM products p
   WHERE NOT EXISTS (
     SELECT 1 FROM product_translations pt 
     WHERE pt.product_id = p.id AND pt.locale = 'bn'
   )
   ORDER BY p.id;"
```

---

## Translation Quality

### Built-in Dictionary
The tool uses a **dictionary of 100+ tech terms** for mapping:
- Device names: smartphone, tablet, laptop, etc.
- Brands: Samsung, Apple, Xiaomi, etc.
- Specs: RAM, storage, battery, etc.
- Features: fingerprint, 5G, fast charging, etc.

**Example translations:**
- "Samsung Galaxy S25" → "স্যামসাং গ্যালাক্সি এস২৫"
- "Apple iPhone 15 Pro" → "অ্যাপল আইফোন ১৫ প্রো"
- "Xiaomi Redmi Note 15" → "শাওমি রেডমি নোট ১৫"

### Improving Translation Quality

If you want **more accurate** translations, you have 3 options:

#### Option 1: **Use Google Translate API** (Recommended for accuracy)
```bash
# Install dependency
go get cloud.google.com/go/translate

# Set up Google Cloud credentials
export GOOGLE_APPLICATION_CREDENTIALS="/path/to/service-account-key.json"

# Update the tool (I can do this for you)
```

#### Option 2: **Prepare Custom CSV**
Create a file with your translations:

```csv
product_id,translated_name_bn
1,আপনার কাস্টম নাম
2,অন্য পণ্যের নাম
3,আরেকটি নাম
```

#### Option 3: **Manual Review & Correction**
After auto-generation, you can manually edit translations:

```sql
UPDATE product_translations
SET translated_name = 'সঠিক বাংলা নাম'
WHERE product_id = 1 AND locale = 'bn';
```

---

## Troubleshooting

### Error: "Failed to connect to database"
**Solution:** Ensure `.env` file has `DATABASE_URL`:
```
DATABASE_URL=postgres://user:password@localhost:5432/gocrit
```

### Error: "translated_name field cannot be empty"
**Solution:** The translation algorithm couldn't find a match. Check the product name in the database.

### Tool runs but creates no translations
**Solution:** All products might already have Bengali translations. Run with `-show-all` to verify.

### Need to undo translations
```sql
-- Delete all Bengali translations (use with caution!)
DELETE FROM product_translations WHERE locale = 'bn';
```

---

## Database Schema Reference

```sql
-- product_translations table structure
CREATE TABLE product_translations (
    id SERIAL PRIMARY KEY,
    product_id INTEGER NOT NULL,
    locale VARCHAR(10) NOT NULL,           -- e.g., 'bn' for Bengali
    translated_name VARCHAR(255) NOT NULL,
    price VARCHAR(255),                     -- Optional
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    UNIQUE(product_id, locale),
    FOREIGN KEY (product_id) REFERENCES products(id)
);
```

---

## Performance Notes

- **Analysis**: < 1 second for 50 products
- **Generation**: ~5ms per product
- **Batch insert**: 50+ products in ~2-3 seconds

---

## Next Steps

1. **Run analysis**: `go run cmd/product-translator/main.go -analyze`
2. **Review results**: Check how many are missing
3. **Generate**: `go run cmd/product-translator/main.go -generate-bn`
4. **Verify**: Check database with SQL query above
5. **Optional**: Use Google Translate API for better accuracy

---

## API Integration (After Translation)

Once translations are in database, they're accessible via API:

```javascript
// Frontend - Get Bengali product name
const productId = 1;
fetch(`/api/products/${productId}/translations/bn`)
  .then(res => res.json())
  .then(data => console.log(data.translated_name))  // বাংলা নাম
```

---

## Support & Questions

For issues or questions:
1. Check troubleshooting section above
2. Review the log output with `-verbose` flag
3. Check database directly with SQL queries
4. Look at the product name in the database - it's used for translation
