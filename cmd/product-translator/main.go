package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"kossti/internal/domain/entities"
	postgrepo "kossti/internal/interface/repository/postgres"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	fmt.Println("🚀 Starting Product Translator...")

	// Load environment variables - prioritize .env.production
	if _, err := os.Stat(".env.production"); err == nil {
		if err := godotenv.Load(".env.production"); err != nil {
			log.Println("Warning: .env.production exists but failed to load")
		} else {
			fmt.Println("✅ Loaded .env.production")
		}
	} else {
		// Fallback to .env
		if err := godotenv.Load(); err != nil {
			log.Println("Warning: No .env or .env.production found")
		}
	}

	// Flags
	analysisOnly := flag.Bool("analyze", false, "Only analyze missing translations (don't create)")
	generateBN := flag.Bool("generate-bn", false, "Generate Bengali translations for missing products")
	verbose := flag.Bool("verbose", false, "Show detailed output")

	flag.Parse()

	// Connect to database
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		// Build DSN from individual .env variables
		dbHost := os.Getenv("DB_HOST")
		if dbHost == "" {
			dbHost = "localhost"
		}
		dbPort := os.Getenv("DB_PORT")
		if dbPort == "" {
			dbPort = "5432"
		}
		dbUser := os.Getenv("DB_USER")
		if dbUser == "" {
			dbUser = "postgres"
		}
		dbPassword := os.Getenv("DB_PASSWORD")
		if dbPassword == "" {
			dbPassword = "postgres"
		}
		dbName := os.Getenv("DB_NAME")
		if dbName == "" {
			dbName = "gocrit"
		}

		dsn = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
			dbUser, dbPassword, dbHost, dbPort, dbName)
	}

	fmt.Println("🔗 Connecting to database...")
	maskDSN := strings.Split(dsn, "@")[0] + "@****"
	fmt.Printf("   DSN pattern: %s\n", maskDSN)

	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		fmt.Printf("❌ Failed to connect to database: %v\n", err)
		fmt.Println("\nTroubleshooting tips:")
		fmt.Println("1. Check if .env.production file exists and is readable")
		fmt.Println("2. Check if DATABASE_URL is set in .env.production")
		fmt.Println("3. Verify network/firewall can reach the database")
		os.Exit(1)
	}

	fmt.Println("✅ Connected to database")

	// Initialize repository
	productRepo := postgrepo.NewPostgresProductRepo(gormDB)
	ctx := context.Background()

	// Get total product count
	fmt.Println("📚 Fetching product count...")
	type CountResult struct {
		Total     int
		WithBN    int
		MissingBN int
	}

	var counts CountResult
	sqlDB, _ := gormDB.DB()

	// Get total products
	err = sqlDB.QueryRow("SELECT COUNT(*) FROM products WHERE deleted_at IS NULL").Scan(&counts.Total)
	if err != nil {
		fmt.Printf("❌ Error counting products: %v\n", err)
		os.Exit(1)
	}

	// Get products with Bengali translations
	err = sqlDB.QueryRow("SELECT COUNT(DISTINCT product_id) FROM product_translations WHERE locale='bn'").Scan(&counts.WithBN)
	if err != nil {
		fmt.Printf("❌ Error counting Bengali translations: %v\n", err)
		os.Exit(1)
	}

	counts.MissingBN = counts.Total - counts.WithBN

	fmt.Printf("\n📊 Total Products: %d\n", counts.Total)
	fmt.Println(strings.Repeat("=", 70))

	// Summary
	fmt.Println("\n" + strings.Repeat("=", 70))
	fmt.Printf("📈 Summary:\n")
	fmt.Printf("   ✅ Products with Bengali (bn): %d\n", counts.WithBN)
	fmt.Printf("   ❌ Products missing Bengali:   %d\n", counts.MissingBN)
	fmt.Printf("   📊 Coverage: %.1f%%\n", float64(counts.WithBN)*100/float64(counts.Total))

	// If only analysis, stop here
	if *analysisOnly && !*generateBN {
		fmt.Println("\n🔍 Analysis complete. Use --generate-bn to create translations.")
		return
	}

	// Generate Bengali translations
	if *generateBN && counts.MissingBN > 0 {
		fmt.Printf("\n🔄 Generating Bengali translations for %d products...\n", counts.MissingBN)
		fmt.Println(strings.Repeat("=", 70))

		// Fetch missing products in batches
		fmt.Println("📚 Fetching products without Bengali translations...")

		var missingBN []*entities.Product
		rows, err := sqlDB.Query(`
			SELECT p.id, p.name, p.description, p.slug, p.start_price, p.end_price, p.category_id, p.brand_id, p.views_count, p.status, p.priority, p.created_at, p.updated_at
			FROM products p
			WHERE p.deleted_at IS NULL
			AND p.id NOT IN (SELECT product_id FROM product_translations WHERE locale='bn')
			ORDER BY p.id
		`)
		if err != nil {
			fmt.Printf("❌ Error fetching missing products: %v\n", err)
			os.Exit(1)
		}
		defer rows.Close()

		for rows.Next() {
			product := &entities.Product{}
			var nullStartPrice, nullEndPrice sql.NullFloat64
			var nullCategoryID sql.NullInt64
			var nullBrandID sql.NullInt64

			err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Slug, &nullStartPrice, &nullEndPrice,
				&nullCategoryID, &nullBrandID, &product.ViewsCount, &product.Status, &product.Priority,
				&product.CreatedAt, &product.UpdatedAt)
			if err != nil {
				fmt.Printf("❌ Error scanning product: %v\n", err)
				continue
			}

			// Handle nullable fields
			if nullStartPrice.Valid {
				v := nullStartPrice.Float64
				product.StartPrice = &v
			}
			if nullEndPrice.Valid {
				v := nullEndPrice.Float64
				product.EndPrice = &v
			}
			if nullCategoryID.Valid {
				id := uint(nullCategoryID.Int64)
				product.CategoryID = &id
			}
			if nullBrandID.Valid {
				id := uint(nullBrandID.Int64)
				product.BrandID = &id
			}
			missingBN = append(missingBN, product)
		}

		if *verbose {
			fmt.Printf("   Found %d products to translate\n", len(missingBN))
		}

		successCount := 0
		errorCount := 0

		for i, product := range missingBN {
			bengaliName := enToBengali(product.Name)

			// Convert float prices to strings for Bengali text support
			var startPriceStr *string
			var endPriceStr *string
			if product.StartPrice != nil {
				sp := fmt.Sprintf("%.2f", *product.StartPrice)
				startPriceStr = &sp
			}
			if product.EndPrice != nil {
				ep := fmt.Sprintf("%.2f", *product.EndPrice)
				endPriceStr = &ep
			}

			translation := &entities.ProductTranslation{
				ProductID:      product.ID,
				Locale:         "bn",
				TranslatedName: bengaliName,
				StartPrice:     startPriceStr, // Optional start price
				EndPrice:       endPriceStr,   // Optional end price
			}

			_, err := productRepo.CreateTranslation(ctx, translation)
			if err != nil {
				fmt.Printf("❌ [%d/%d] Failed to create: %s\n", i+1, len(missingBN), product.Name)
				fmt.Printf("   Error: %v\n", err)
				errorCount++
			} else {
				fmt.Printf("✅ [%d/%d] Created: %s => %s\n", i+1, len(missingBN), product.Name, bengaliName)
				successCount++
			}
		}

		fmt.Println("\n" + strings.Repeat("=", 70))
		fmt.Printf("✅ Successfully created: %d\n", successCount)
		fmt.Printf("❌ Failed: %d\n", errorCount)
		fmt.Printf("📊 Total now in Bengali: %d\n", counts.WithBN+successCount)
	}
}

// enToBengali converts English product names to Bengali using a comprehensive mapping
// For more accurate translations of uncommon products, consider integrating Google Translate API
func enToBengali(english string) string {
	// Comprehensive English to Bengali mapping for tech products
	translations := map[string]string{
		// === DEVICE TYPES ===
		"smartphone":        "স্মার্টফোন",
		"phone":             "ফোন",
		"mobile phone":      "মোবাইল ফোন",
		"mobile":            "মোবাইল",
		"tablet":            "ট্যাবলেট",
		"laptop":            "ল্যাপটপ",
		"computer":          "কম্পিউটার",
		"desktop":           "ডেস্কটপ",
		"monitor":           "মনিটর",
		"smartwatch":        "স্মার্ট ঘড়ি",
		"watch":             "ঘড়ি",
		"earbuds":           "ইয়ারবাডস",
		"headphones":        "হেডফোনস",
		"headphone":         "হেডফোন",
		"earphone":          "ইয়ারফোন",
		"speaker":           "স্পিকার",
		"bluetooth speaker": "ব্লুটুথ স্পিকার",
		"camera":            "ক্যামেরা",
		"cannon camera":     "ক্যানন ক্যামেরা",
		"power bank":        "পাওয়ার ব্যাংক",
		"charger":           "চার্জার",
		"cable":             "ক্যাবল",
		"adapter":           "এডাপ্টার",

		// === DISPLAY & SCREEN ===
		"screen":  "স্ক্রিন",
		"display": "ডিসপ্লে",
		"amoled":  "অ্যামোলেড",
		"lcd":     "এলসিডি",
		"ips":     "আইপিএস",
		"retina":  "রেটিনা",

		// === MEMORY & STORAGE ===
		"memory":  "মেমরি",
		"ram":     "র্যাম",
		"storage": "স্টোরেজ",
		"ssd":     "এসএসডি",
		"hdd":     "এইচডিডি",
		"gb":      "গিগাবাইট",
		"tb":      "টেরাবাইট",
		"mb":      "মেগাবাইট",

		// === PROCESSOR & PERFORMANCE ===
		"processor":  "প্রসেসর",
		"cpu":        "সিপিইউ",
		"chipset":    "চিপসেট",
		"gpu":        "জিপিইউ",
		"snapdragon": "স্ন্যাপড্রেগন",
		"apple m":    "অ্যাপল এম",
		"intel":      "ইন্টেল",
		"amd":        "এএমডি",
		"core":       "কোর",
		"octa-core":  "অক্টা-কোর",
		"quad-core":  "কোয়াড-কোর",
		"dual-core":  "ডুয়াল-কোর",

		// === BRANDS ===
		"samsung":   "স্যামসাং",
		"apple":     "অ্যাপল",
		"xiaomi":    "শাওমি",
		"realme":    "রিয়েলমি",
		"vivo":      "ভিভো",
		"oppo":      "ওপিও",
		"oneplus":   "ওয়ানপ্লাস",
		"motorola":  "মোটোরোলা",
		"nokia":     "নোকিয়া",
		"sony":      "সোনি",
		"lg":        "এলজি",
		"google":    "গুগল",
		"pixel":     "পিক্সেল",
		"nexus":     "নেক্সাস",
		"honor":     "অনার",
		"asus":      "এসাস",
		"moto":      "মোটো",
		"redmi":     "রেডমি",
		"poco":      "পোকো",
		"itel":      "আইটেল",
		"micromax":  "মাইক্রোম্যাক্স",
		"lava":      "লাভা",
		"karbonn":   "কার্বন",
		"dell":      "ডেল",
		"hp":        "এইচপি",
		"lenovo":    "লেনোভো",
		"toshiba":   "টোশিবা",
		"canon":     "ক্যানন",
		"nikon":     "নিকন",
		"panasonic": "প্যানাসনিক",

		// === MODEL NAMES & SERIES ===
		"galaxy":   "গ্যালাক্সি",
		"iphone":   "আইফোন",
		"ipad":     "আইপ্যাড",
		"macbook":  "ম্যাকবুক",
		"imac":     "আইম্যাক",
		"note":     "নোট",
		"a series": "এ সিরিজ",
		"s series": "এস সিরিজ",
		"ultra":    "আল্ট্রা",
		"max":      "ম্যাক্স",
		"pro":      "প্রো",
		"plus":     "প্লাস",
		"mini":     "মিনি",
		"lite":     "লাইট",
		"edge":     "এজ",
		"fold":     "ফোল্ড",
		"flip":     "ফ্লিপ",
		"z":        "জেড",
		"air":      "এয়ার",
		"new":      "নতুন",

		// === CONNECTIVITY ===
		"5g":        "৫জি",
		"4g":        "৪জি",
		"3g":        "৩জি",
		"wifi":      "ওয়াইফাই",
		"wi-fi":     "ওয়াইফাই",
		"bluetooth": "ব্লুটুথ",
		"usb":       "ইউএসবি",
		"type-c":    "টাইপ-সি",
		"lightning": "লাইটনিং",
		"nfc":       "এনএফসি",

		// === FEATURES ===
		"megapixel":          "মেগাপিক্সেল",
		"mp":                 "এমপি",
		"battery":            "ব্যাটারি",
		"mah":                "এমএএইচ",
		"charging":           "চার্জিং",
		"fast charging":      "দ্রুত চার্জিং",
		"wireless charging":  "ওয়্যারলেস চার্জিং",
		"fingerprint":        "ফিঙ্গারপ্রিন্ট",
		"face unlock":        "ফেস আনলক",
		"facial recognition": "ফেসিয়াল রিকগনিশন",
		"accelerometer":      "অ্যাক্সিলারোমিটার",
		"gyroscope":          "জাইরোস্কোপ",
		"touch":              "টাচ",
		"sensor":             "সেন্সর",
		"water resistant":    "জলরোধী",
		"ip rating":          "আইপি রেটিং",
		"dual speaker":       "ডুয়াল স্পিকার",
		"stereo":             "স্টেরিও",

		// === COLORS ===
		"black":  "কালো",
		"white":  "সাদা",
		"silver": "রূপালী",
		"gold":   "সোনালী",
		"blue":   "নীল",
		"red":    "লাল",
		"green":  "সবুজ",
		"purple": "বেগুনি",
		"pink":   "গোলাপী",
		"gray":   "ধূসর",

		// === MISC ===
		"inch":       "ইঞ্চ",
		"hz":         "হার্জ",
		"fps":        "এফপিএস",
		"version":    "সংস্করণ",
		"model":      "মডেল",
		"edition":    "সংস্করণ",
		"variant":    "ভেরিয়েন্ট",
		"series":     "সিরিজ",
		"gen":        "জেনারেশন",
		"generation": "জেনারেশন",
	}

	lower := strings.ToLower(english)
	bengaliResult := english
	matchCount := 0

	// Find the longest matching word/phrase
	var bestMatch string
	var bestReplacement string

	for en, bn := range translations {
		if strings.Contains(lower, en) && len(en) > len(bestMatch) {
			bestMatch = en
			bestReplacement = bn
		}
	}

	if bestMatch != "" {
		// Find the original casing version and replace
		index := strings.Index(lower, bestMatch)
		original := english[index : index+len(bestMatch)]
		bengaliResult = strings.Replace(english, original, bestReplacement, -1)
		matchCount++
	}

	// If no match found, return as is with a note for manual translation
	if matchCount == 0 {
		return english + " (অনুবাদ প্রয়োজন)"
	}

	return bengaliResult
}
