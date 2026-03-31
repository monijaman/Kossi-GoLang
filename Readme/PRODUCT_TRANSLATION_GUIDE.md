# Product Translation to Bengali - Guide

## Quick Start

### 1. **Analyze Missing Translations (Recommended First)**

```bash
# Go to server directory
cd d:/GO/gocrit/gocrit_server

# Run analysis only (no changes)
go run cmd/product-translator/main.go -analyze

# Or see all products with status
go run cmd/product-translator/main.go -analyze -show-all
```

**Output Example:**
```
📊 Total Products: 45

📈 Summary:
   ✅ Products with Bengali (bn): 12
   ❌ Products missing Bengali:   33
   📊 Coverage: 26.7%
```

---

### 2. **Generate Bengali Translations**

#### Option A: **Simple Translation (Uses Built-in Mapping)**
```bash
go run cmd/product-translator/main.go -generate-bn
```

This uses a predefined dictionary for common tech terms. **Pros:** Fast, no API keys. **Cons:** Limited accuracy.

---

#### Option B: **Google Translate API (Recommended)**

1. **Setup Google Cloud Credentials:**
   ```bash
   # Install Google Translate package (if not already installed)
   go get cloud.google.com/go/translate

   # Set up service account
   export GOOGLE_APPLICATION_CREDENTIALS="path/to/service-account-key.json"
   ```

2. **Update the tool** (I'll show you how below)

---

### 3. **Specific Flags**

| Flag | Purpose | Example |
|------|---------|---------|
| `-analyze` | Only analyze, don't create | `go run cmd/product-translator/main.go -analyze` |
| `-generate-bn` | Create Bengali translations | `go run cmd/product-translator/main.go -generate-bn` |
| `-show-all` | Show all products | `go run cmd/product-translator/main.go -show-all` |
| `-verbose` | Detailed output | `go run cmd/product-translator/main.go -generate-bn -verbose` |

---

## Usage Examples

### Example 1: Check Coverage
```bash
go run cmd/product-translator/main.go -analyze -show-all
```
Shows all products and their Bengali translation status.

### Example 2: Auto-Generate All Missing
```bash
go run cmd/product-translator/main.go -generate-bn
```
Creates Bengali translations for all missing products.

### Example 3: Step-by-Step Approach
```bash
# Step 1: Analyze
go run cmd/product-translator/main.go -analyze

# Step 2: Review output
# ...

# Step 3: Generate
go run cmd/product-translator/main.go -generate-bn
```

---

## Database Schema

**product_translations table:**
```
id (Primary Key)
product_id (Foreign Key → products.id)
locale (e.g., "bn" for Bengali)
translated_name (The Bengali product name)
price (Optional - for localized pricing)
created_at
updated_at
```

---

## Manual Translation Method

If you have translations manually prepared, you can insert them directly:

### Option 1: Via API
```bash
# Create a single translation
curl -X POST http://localhost:8080/api/products/{productID}/translations \
  -H "Content-Type: application/json" \
  -d '{
    "locale": "bn",
    "translated_name": "স্যামসাং গ্যালাক্সি S25",
    "price": null
  }'
```

### Option 2: Via SQL
```sql
INSERT INTO product_translations (product_id, locale, translated_name, created_at, updated_at)
VALUES 
  (1, 'bn', 'স্যামসাং গ্যালাক্সি S25', NOW(), NOW()),
  (2, 'bn', 'আইফোন 15 প্রো', NOW(), NOW()),
  (3, 'bn', 'শাওমি রিডমি নোট 15', NOW(), NOW());
```

---

## Advanced: Google Translate Integration

Create `cmd/product-translator/translate.go` for API integration:

```go
package main

import (
	"cloud.google.com/go/translate"
	"context"
)

func translateWithGoogle(ctx context.Context, text string) (string, error) {
	client, err := translate.NewClient(ctx)
	if err != nil {
		return "", err
	}
	defer client.Close()

	resp, err := client.Translate(ctx, []string{text}, &translate.Options{
		TargetLanguage: "bn",
		SourceLanguage: "en",
	})
	if err != nil {
		return "", err
	}

	if len(resp) > 0 {
		return resp[0].TranslatedText, nil
	}
	return text, nil
}
```

---

## Status Check

To see real-time progress of translations:

```bash
# Check Bengali translation count in database
psql -h localhost -U postgres -d gocrit -c \
  "SELECT COUNT(*) as bengali_translations FROM product_translations WHERE locale='bn';"
```

---

## Troubleshooting

| Issue | Solution |
|-------|----------|
| "Failed to connect to database" | Check `.env` file - `DATABASE_URL` must be set |
| "translation not found" (not an error) | Normal - means product hasn't been translated yet |
| "translated_name field cannot be empty" | Translation function returned empty string - check mapping |

---

## Performance Notes

- **45 products:** ~2-3 seconds for analysis, ~5-10 seconds to generate
- **100+ products:** Consider batching translations
- **Google API:** ~10-20ms per translation request

---

## Next Steps

1. Run `go run cmd/product-translator/main.go -analyze` to see what's missing
2. Decide on translation method (simple, Google Translate, or manual)
3. Run generation command
4. Verify in database: `SELECT COUNT(*) FROM product_translations WHERE locale='bn';`

---

## Need Help?

For better translations with special characters and Bangla grammar:
- **Option 1:** Use `google-cloud-translate` package
- **Option 2:** Integrate `reverso.net` or `yandex` translator API
- **Option 3:** Prepare manual CSV with translations and bulk import
