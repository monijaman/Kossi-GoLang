package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SeedMobileProductTranslations seeds Bengali translations for mobile products
func SeedMobileProductTranslations(db *gorm.DB) error {
	fmt.Println("\n🔄 Seeding Mobile Product Translations (Bangla)...")

	// Get all mobile products
	var products []models.ProductModel
	db.Where("category_id = ?", 79).Find(&products)

	if len(products) == 0 {
		log.Println("⚠️  No products found for translation")
		return nil
	}

	translations := []models.ProductTranslationModel{}

	// Create Bangla translations for each product
	for _, product := range products {
		translation := models.ProductTranslationModel{
			ProductID:      product.ID,
			Locale:         "bn",
			TranslatedName: getBanglaTitleForProduct(product.Name),
		}
		translations = append(translations, translation)
	}

	// Batch create translations
	result := db.CreateInBatches(translations, 100)
	if result.Error != nil {
		log.Printf("❌ Error seeding product translations: %v", result.Error)
		return result.Error
	}

	fmt.Printf("✅ Successfully seeded %d product translations (Bangla)\n", result.RowsAffected)
	return nil
}

// getBanglaTitleForProduct returns Bangla translation for product names
func getBanglaTitleForProduct(englishName string) string {
	translations := map[string]string{
		// Apple
		"iPhone 17 Pro Max":   "আইফোন ১৭ প্রো ম্যাক্স",
		"iPhone 17 Pro":       "আইফোন ১৭ প্রো",
		"iPhone 17 Plus":      "আইফোন ১৭ প্লাস",
		"iPhone 17":           "আইফোন ১৭",
		"iPhone 16 Pro Max":   "আইফোন ১৬ প্রো ম্যাক্স",
		"iPhone 16 Pro":       "আইফোন ১৬ প্রো",
		"iPhone 16":           "আইফোন ১৬",
		"iPhone 15":           "আইফোন ১৫",
		"iPhone SE (3rd Gen)": "আইফোন এসই (তৃতীয় প্রজন্ম)",
		"iPhone SE":           "আইফোন এসই",

		// Samsung Galaxy S Series
		"Samsung Galaxy S25 Ultra": "স্যামসাং গ্যালাক্সি এস২৫ আলট্রা",
		"Samsung Galaxy S25+":      "স্যামসাং গ্যালাক্সি এস২৫+",
		"Samsung Galaxy S25":       "স্যামসাং গ্যালাক্সি এস২৫",
		"Samsung Galaxy S24 Ultra": "স্যামসাং গ্যালাক্সি এস২৪ আলট্রা",
		"Samsung Galaxy S24+":      "স্যামসাং গ্যালাক্সি এস২৪+",
		"Samsung Galaxy S24":       "স্যামসাং গ্যালাক্সি এস২৪",
		"Samsung Galaxy S24 FE":    "স্যামসাং গ্যালাক্সি এস২৪ এফই",
		"Samsung Galaxy S23 Ultra": "স্যামসাং গ্যালাক্সি এস২৩ আলট্রা",
		"Samsung Galaxy S23+":      "স্যামসাং গ্যালাক্সি এস২৩+",
		"Samsung Galaxy S23":       "স্যামসাং গ্যালাক্সি এস২৩",

		// Samsung Galaxy Z (Foldables)
		"Samsung Galaxy Z Fold 7":    "স্যামসাং গ্যালাক্সি জেড ফোল্ড ৭",
		"Samsung Galaxy Z Flip 7":    "স্যামসাং গ্যালাক্সি জেড ফ্লিপ ৭",
		"Samsung Galaxy Z Flip 7 FE": "স্যামসাং গ্যালাক্সি জেড ফ্লিপ ৭ এফই",
		"Samsung Galaxy Z Fold 6":    "স্যামসাং গ্যালাক্সি জেড ফোল্ড ৬",
		"Samsung Galaxy Z Flip 6":    "স্যামসাং গ্যালাক্সি জেড ফ্লিপ ৬",

		// Samsung Galaxy A Series
		"Samsung Galaxy A56 5G": "স্যামসাং গ্যালাক্সি এ৫৬ ৫জি",
		"Samsung Galaxy A55 5G": "স্যামসাং গ্যালাক্সি এ৫৫ ৫জি",
		"Samsung Galaxy A54 5G": "স্যামসাং গ্যালাক্সি এ৫৪ ৫জি",
		"Samsung Galaxy A25 5G": "স্যামসাং গ্যালাক্সি এ২৫ ৫জি",
		"Samsung Galaxy A15 5G": "স্যামসাং গ্যালাক্সি এ১৫ ৫জি",

		// Samsung Galaxy M Series
		"Samsung Galaxy M56 5G": "স্যামসাং গ্যালাক্সি এম৫৬ ৫জি",
		"Samsung Galaxy M36 5G": "স্যামসাং গ্যালাক্সি এম৩৬ ৫জি",

		// Samsung Galaxy F Series
		"Samsung Galaxy F56 5G": "স্যামসাং গ্যালাক্সি এফ৫৬ ৫জি",
		"Samsung Galaxy F36 5G": "স্যামসাং গ্যালাক্সি এফ৩৬ ৫জি",

		// Xiaomi Flagship Series
		"Xiaomi 15 Ultra": "শিয়াওমি ১৫ আলট্রা",
		"Xiaomi 15 Pro":   "শিয়াওমি ১৫ প্রো",
		"Xiaomi 15":       "শিয়াওমি ১৫",
		"Xiaomi 14 Ultra": "শিয়াওমি ১৪ আলট্রা",
		"Xiaomi 14 Pro":   "শিয়াওমি ১৪ প্রো",

		// Redmi Note Series
		"Redmi Note 15":         "রেডমি নোট ১৫",
		"Redmi Note 15 Pro 5G":  "রেডমি নোট ১৫ প্রো ৫জি",
		"Redmi Note 14":         "রেডমি নোট ১৪",
		"Redmi Note 14 Pro+ 5G": "রেডমি নোট ১৪ প্রো+ ৫জি",
		"Redmi Note 13":         "রেডমি নোট ১৩",
		"Redmi Note 12":         "রেডমি নোট ১২",

		// Redmi Numbered and POCO
		"Redmi 15":       "রেডমি ১৫",
		"Redmi 14":       "রেডমি ১৪",
		"Redmi 13C":      "রেডমি ১৩সি",
		"POCO F7":        "পোকো এফ৭",
		"POCO F7 Ultra":  "পোকো এফ৭ আলট্রা",
		"POCO F6 5G":     "পোকো এফ৬ ৫জি",
		"POCO X7 Pro 5G": "পোকো এক্স৭ প্রো ৫জি",
		"POCO X7 5G":     "পোকো এক্স৭ ৫জি",
		"POCO X6 Pro 5G": "পোকো এক্স৬ প্রো ৫জি",
		"POCO M7 Pro 5G": "পোকো এম৭ প্রো ৫জি",
		"POCO M6 Pro 5G": "পোকো এম৬ প্রো ৫জি",

		// Vivo
		"Vivo X300 Pro":    "ভিভো এক্স৩০০ প্রো",
		"Vivo X300":        "ভিভো এক্স৩০০",
		"Vivo X200 Pro":    "ভিভো এক্স২০০ প্রো",
		"Vivo V60 5G":      "ভিভো ভি৬০ ৫জি",
		"Vivo V60 Lite 5G": "ভিভো ভি৬০ লাইট ৫জি",
		"Vivo V50":         "ভিভো ভি৫০",
		"Vivo V40":         "ভিভো ভি৪০",
		"Vivo Y29":         "ভিভো ওয়াই২৯",
		"Vivo Y28":         "ভিভো ওয়াই২৮",

		// Oppo
		"Oppo Find N4 Fold":   "ওপিও ফাইন্ড এন৪ ফোল্ড",
		"Oppo Find N4 Flip":   "ওপিও ফাইন্ড এন৪ ফ্লিপ",
		"Oppo Reno 14 5G":     "ওপিও রেনো ১৪ ৫জি",
		"Oppo Reno 14 F 5G":   "ওপিও রেনো ১৪ এফ ৫জি",
		"Oppo Reno 13 5G":     "ওপিও রেনো ১৩ ৫জি",
		"Oppo Reno 13 Pro 5G": "ওপিও রেনো ১৩ প্রো ৫জি",
		"Oppo Reno 15":        "ওপিও রেনো ১৫",
		"Oppo Reno 15 Pro":    "ওপিও রেনো ১৫ প্রো",
		"Oppo A5 Pro":         "ওপিও এ৫ প্রো",
		"Oppo A5":             "ওপিও এ৫",
		"Oppo A60":            "ওপিও এ৬০",

		// Realme
		"Realme 15 Pro 5G": "রিয়েলমি ১৫ প্রো ৫জি",
		"Realme 15 5G":     "রিয়েলমি ১৫ ৫জি",
		"Realme 14 5G":     "রিয়েলমি ১৪ ৫জি",
		"Realme 13":        "রিয়েলমি ১৩",
		"Realme C75x":      "রিয়েলমি সি৭৫এক্স",
		"Realme C71":       "রিয়েলমি সি৭১",
		"Realme C65":       "রিয়েলমি সি৬৫",
		"Realme GT Neo 3":  "রিয়েলমি জিটিএনিও ৩",

		// Google
		"Google Pixel 9 Pro XL":   "গুগল পিক্সেল ৯ প্রো এক্সএল",
		"Google Pixel 9 Pro":      "গুগল পিক্সেল ৯ প্রো",
		"Google Pixel 9":          "গুগল পিক্সেল ৯",
		"Google Pixel 9 Pro Fold": "গুগল পিক্সেল ৯ প্রো ফোল্ড",
		"Google Pixel 8 Pro":      "গুগল পিক্সেল ৮ প্রো",

		// OnePlus
		"OnePlus 13":       "ওয়ানপ্লাস ১৩",
		"OnePlus 13R":      "ওয়ানপ্লাস ১৩আর",
		"OnePlus 12":       "ওয়ানপ্লাস ১২",
		"OnePlus 12R":      "ওয়ানপ্লাস ১২আর",
		"OnePlus Nord 5":   "ওয়ানপ্লাস নর্ড ৫",
		"OnePlus Nord 4":   "ওয়ানপ্লাস নর্ড ৪",
		"OnePlus Nord CE5": "ওয়ানপ্লাস নর্ড সিই৫",
		"OnePlus Open":     "ওয়ানপ্লাস ওপেন",

		// Walton
		"Walton XANON X1 Ultra": "ওয়ালটন জ্যানন এক্স১ আলট্রা",
		"Walton XANON X91":      "ওয়ালটন জ্যানন এক্স৯১",
		"Walton NEXG N75":       "ওয়ালটন নেক্সজি এন৭৫",
		"Walton NEXG N27":       "ওয়ালটন নেক্সজি এন২৭",
		"Walton Orbit Y13":      "ওয়ালটন অরবিট ওয়াই১৩",
		"Walton ZENX 2":         "ওয়ালটন জেনএক্স ২",
		// Walton (additional models & uppercase variants)
		"Walton ZENX 1T":        "ওয়ালটন জেনএক্স ১টি",
		"Walton Orbit Y71":      "ওয়ালটন অরবিট ওয়াই৭১",
		"Walton Primo S8":       "ওয়ালটন প্রিমো এস৮",
		"Walton Primo GH11":     "ওয়ালটন প্রিমো জিএইচ১১",
		"Walton Primo H10":      "ওয়ালটন প্রিমো এইচ১০",
		"Walton Primo ZX4":      "ওয়ালটন প্রিমো জেডএক্স৪",
		"Walton Primo R8":       "ওয়ালটন প্রিমো আর৮",
		"Walton Primo F10":      "ওয়ালটন প্রিমো এফ১০",
		"Walton ZENX 1":         "ওয়ালটন জেনএক্স ১",
		"Walton NEXG N10 ULTRA": "ওয়ালটন নেক্সজি এন১০ আলট্রা",
		"Walton NEXG N26":       "ওয়ালটন নেক্সজি এন২৬",
		"Walton NEXG N74":       "ওয়ালটন নেক্সজি এন৭৪",
		"Walton NEXG N71 PLUS":  "ওয়ালটন নেক্সজি এন৭১ প্লাস",
		"Walton NEXG N9":        "ওয়ালটন নেক্সজি এন৯",
		"Walton ORBIT Y12":      "ওয়ালটন অরবিট ওয়াই১২",
		"Walton ORBIT Y11":      "ওয়ালটন অরবিট ওয়াই১১",
		"Walton ORBIT Y13":      "ওয়ালটন অরবিট ওয়াই১৩",
		"Walton ORBIT Y70C":     "ওয়ালটন অরবিট ওয়াই৭০সি",
		"Walton ORBIT Y71":      "ওয়ালটন অরবিট ওয়াই৭১",
		"Walton XANON X1 ULTRA": "ওয়ালটন জ্যানন এক্স১ আলট্রা",
		"Walton XANON X21":      "ওয়ালটন জ্যানন এক্স২১",

		// Infinix
		"Infinix Note 50 Pro+ 5G": "ইনফিনিক্স নোট ৫০ প্রো+ ৫জি",
		"Infinix Note 50 Pro":     "ইনফিনিক্স নোট ৫০ প্রো",
		"Infinix Note 40 Pro":     "ইনফিনিক্স নোট ৪০ প্রো",
		"Infinix Hot 50":          "ইনফিনিক্স হট ৫০",
		"Infinix Smart 10":        "ইনফিনিক্স স্মার্ট ১০",

		// Tecno
		"Tecno CAMON 40 Pro 5G": "টেকনো ক্যামন ৪০ প্রো ৫জি",
		"Tecno CAMON 40 Pro":    "টেকনো ক্যামন ৪০ প্রো",
		"Tecno CAMON 40":        "টেকনো ক্যামন ৪০",
		"Tecno SPARK 40 5G":     "টেকনো স্পার্ক ৪০ ৫জি",
		"Tecno SPARK 40":        "টেকনো স্পার্ক ৪০",
		"Tecno POVA Slim 5G":    "টেকনো পোভা স্লিম ৫জি",

		// Symphony
		"Symphony G27":    "সিম্ফনি জি২৭",
		"Symphony Atom 5": "সিম্ফনি অ্যাটম ৫",
	}

	if bengaliName, exists := translations[englishName]; exists {
		return bengaliName
	}

	// Fallback: return the English name if no translation exists
	return englishName
}
