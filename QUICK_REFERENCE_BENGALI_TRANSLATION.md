# Bengali Product Translation - Quick Reference

## 🚀 Get Started in 30 Seconds

```bash
cd d:/GO/gocrit/gocrit_server

# See what's missing
go run cmd/product-translator/main.go -analyze

# Create Bengali translations
go run cmd/product-translator/main.go -generate-bn
```

---

## 📋 All Commands

### Analysis
```bash
# Basic analysis
go run cmd/product-translator/main.go -analyze

# Show all products
go run cmd/product-translator/main.go -analyze -show-all

# Detailed output
go run cmd/product-translator/main.go -analyze -verbose
```

### Generation
```bash
# Generate translations for missing products
go run cmd/product-translator/main.go -generate-bn

# With details
go run cmd/product-translator/main.go -generate-bn -verbose
```

### Combined
```bash
# Analyze + Show all
go run cmd/product-translator/main.go -analyze -show-all -verbose

# Full workflow
go run cmd/product-translator/main.go -analyze
go run cmd/product-translator/main.go -generate-bn
```

---

## ✅ Verify Results

```bash
# Count Bengali translations
psql -h localhost -U postgres -d gocrit -c \
  "SELECT COUNT(*) as total FROM product_translations WHERE locale='bn';"

# See all translations
psql -h localhost -U postgres -d gocrit -c \
  "SELECT p.name, pt.translated_name 
   FROM product_translations pt
   JOIN products p ON pt.product_id = p.id
   WHERE pt.locale='bn'
   LIMIT 10;"
```

---

## 📁 Files Created

| File | Purpose |
|------|---------|
| `cmd/product-translator/main.go` | CLI tool (main) |
| `translate-products.sh` | Bash helper script |
| `translate-products.ps1` | PowerShell helper script |
| `product_translations_bn_sample.sql` | Sample SQL inserts |
| `PRODUCT_TRANSLATION_GUIDE.md` | Detailed guide |
| `BENGALI_TRANSLATION_COMPLETE_GUIDE.md` | Full documentation |

---

## 🔧 Features

✅ Analyze missing translations  
✅ Auto-generate Bengali names  
✅ Insert into database automatically  
✅ Progress reporting  
✅ Built-in tech product dictionary  
✅ Support for 100+ terms  

---

## 🎯 Translation Quality

**Good for:**
- Standard tech products
- Known brands
- Common device types

**Example:**
- Input: "Samsung Galaxy S25"
- Output: "স্যামসাং গ্যালাক্সি এস২৫"

**If you need better:**
- Use Google Translate API
- Manual CSV import
- Custom SQL updates

---

## ⚡ Quick Fixes

| Issue | Fix |
|-------|-----|
| Can't run command | Make sure you're in `d:/GO/gocrit/gocrit_server` directory |
| Database connection error | Check `.env` has `DATABASE_URL` |
| No translations created | Run `-analyze` first to see status |
| Need to delete all | `DELETE FROM product_translations WHERE locale='bn';` |

---

## 📞 Database Queries

```sql
-- See how many need translation
SELECT COUNT(*) FROM products 
WHERE id NOT IN (
  SELECT product_id FROM product_translations WHERE locale='bn'
);

-- Update a translation
UPDATE product_translations 
SET translated_name = 'সঠিক নাম'
WHERE product_id = 1 AND locale='bn';

-- Delete specific translation
DELETE FROM product_translations 
WHERE product_id = 1 AND locale='bn';
```

---

## 🎓 Learning Path

1. **First run**: `go run cmd/product-translator/main.go -analyze`
2. **Check output**: See coverage percentage
3. **Review**: `-show-all` to see which are missing
4. **Generate**: `-generate-bn` to create them
5. **Verify**: SQL query to confirm

---

## 💡 Pro Tips

- Use `-verbose` flag to see exactly what's happening
- Run analysis before generation to understand scope
- Keep a backup of database before bulk generation
- Test with one product first if unsure
- Use SQL queries to verify results

---

## 📚 Full Documentation

For comprehensive guide, see: `BENGALI_TRANSLATION_COMPLETE_GUIDE.md`

For translation details, see: `PRODUCT_TRANSLATION_GUIDE.md`

---

## ✨ Next Steps

- [ ] Run `-analyze` to see coverage
- [ ] Review missing products
- [ ] Run `-generate-bn` to create translations
- [ ] Verify in database
- [ ] (Optional) Integrate Google Translate for better quality
- [ ] (Optional) Set up translations in frontend
