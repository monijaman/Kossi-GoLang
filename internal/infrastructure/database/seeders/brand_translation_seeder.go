package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// BrandTranslationSeeder handles seeding brand translations
type BrandTranslationSeeder struct {
	BaseSeeder
}

// NewBrandTranslationSeeder creates a new brand translation seeder
func NewBrandTranslationSeeder() *BrandTranslationSeeder {
	return &BrandTranslationSeeder{
		BaseSeeder: BaseSeeder{name: "Brand Translations"},
	}
}

// Seed implements the Seeder interface
func (bts *BrandTranslationSeeder) Seed(db *gorm.DB) error {
	// Bangladeshi banks only (removed: Citibank, Commercial Bank of Ceylon, Habib Bank, HSBC Bank, National Bank of Pakistan, State Bank of India)
	translations := map[string]map[string]string{
		"ab-bank":                      {"bn": "এবি ব্যাংক"},
		"agrani-bank":                  {"bn": "অগ্রণী ব্যাংক"},
		"al-arafah-islami-bank":        {"bn": "আল-আরাফা ইসলামী ব্যাংক"},
		"bangladesh-commerce-bank":     {"bn": "বাংলাদেশ কমার্স ব্যাংক"},
		"bangladesh-development-bank":  {"bn": "বাংলাদেশ উন্নয়ন ব্যাংক"},
		"bangladesh-krishi-bank":       {"bn": "বাংলাদেশ কৃষি ব্যাংক"},
		"bank-alfalah":                 {"bn": "ব্যাংক আল ফালাহ"},
		"bank-asia":                    {"bn": "ব্যাংক এশিয়া"},
		"basic-bank":                   {"bn": "বেসিক ব্যাংক"},
		"bengal-bank":                  {"bn": "বেঙ্গল ব্যাংক"},
		"brac-bank":                    {"bn": "ব্র্যাক ব্যাংক"},
		"citizens-bank":                {"bn": "সিটিজেন্স ব্যাংক"},
		"city-bank":                    {"bn": "সিটি ব্যাংক"},
		"community-bank":               {"bn": "কমিউনিটি ব্যাংক"},
		"dhaka-bank":                   {"bn": "ঢাকা ব্যাংক"},
		"dutch-bangla-bank":            {"bn": "ডাচ-বাংলা ব্যাংক"},
		"eastern-bank":                 {"bn": "ইস্টার্ন ব্যাংক"},
		"exim-bank":                    {"bn": "এক্সিম ব্যাংক"},
		"first-security-islami-bank":   {"bn": "ফার্স্ট সিকিউরিটি ইসলামী ব্যাংক"},
		"global-islami-bank":           {"bn": "গ্লোবাল ইসলামী ব্যাংক"},
		"icb-islamic-bank":             {"bn": "আইসিবি ইসলামী ব্যাংক"},
		"ific-bank":                    {"bn": "আইএফআইসি ব্যাংক"},
		"islami-bank":                  {"bn": "ইসলামী ব্যাংক"},
		"jamuna-bank":                  {"bn": "যমুনা ব্যাংক"},
		"janata-bank":                  {"bn": "জনতা ব্যাংক"},
		"meghna-bank":                  {"bn": "মেঘনা ব্যাংক"},
		"mercantile-bank":              {"bn": "মার্চেন্টাইল ব্যাংক"},
		"midland-bank":                 {"bn": "মিডল্যান্ড ব্যাংক"},
		"modhumoti-bank":               {"bn": "মধুমতি ব্যাংক"},
		"mutual-trust-bank":            {"bn": "মিউচুয়াল ট্রাস্ট ব্যাংক"},
		"national-bank":                {"bn": "ন্যাশনাল ব্যাংক"},
		"ncc-bank":                     {"bn": "এনসিসি ব্যাংক"},
		"nrb-bank":                     {"bn": "এনআরবি ব্যাংক"},
		"nrbc-bank":                    {"bn": "এনআরবিসি ব্যাংক"},
		"one-bank":                     {"bn": "ওয়ান ব্যাংক"},
		"padma-bank":                   {"bn": "পদ্মা ব্যাংক"},
		"premier-bank":                 {"bn": "প্রিমিয়ার ব্যাংক"},
		"prime-bank":                   {"bn": "প্রাইম ব্যাংক"},
		"probashi-kallyan-bank":        {"bn": "প্রবাসী কল্যাণ ব্যাংক"},
		"pubali-bank":                  {"bn": "পাবলিক ব্যাংক"},
		"rajshahi-krishi-unnayan-bank": {"bn": "রাজশাহী কৃষি উন্নয়ন ব্যাংক"},
		"rupali-bank":                  {"bn": "রূপালী ব্যাংক"},
		"sbac-bank":                    {"bn": "এসবিএসি ব্যাংক"},
		"shahjalal-islami-bank":        {"bn": "শাহজালাল ইসলামী ব্যাংক"},
		"shimanto-bank":                {"bn": "শিমান্ত ব্যাংক"},
		"social-islami-bank":           {"bn": "সোশ্যাল ইসলামী ব্যাংক"},
		"sonali-bank":                  {"bn": "সোনালী ব্যাংক"},
		"southeast-bank":               {"bn": "সাউথইস্ট ব্যাংক"},
		"standard-bank":                {"bn": "স্ট্যান্ডার্ড ব্যাংক"},
		"standard-chartered-bank":      {"bn": "স্ট্যান্ডার্ড চার্টার্ড ব্যাংক"},
		"trust-bank":                   {"bn": "ট্রাস্ট ব্যাংক"},
		"union-bank":                   {"bn": "ইউনিয়ন ব্যাংক"},
		"united-commercial-bank":       {"bn": "ইউনাইটেড কমার্শিয়াল ব্যাংক"},
		"uttara-bank":                  {"bn": "উত্তরা ব্যাংক"},
		"woori-bank":                   {"bn": "উরি ব্যাংক"},
	}

	for slug, locales := range translations {
		// Find the brand by slug
		var brand models.BrandModel
		result := db.Where("slug = ?", slug).First(&brand)

		if result.Error == nil {
			// Add translations for each locale
			for locale, name := range locales {
				var existingTranslation models.BrandTranslationModel
				transResult := db.Where("brand_id = ? AND locale = ?", brand.ID, locale).First(&existingTranslation)

				if transResult.Error == gorm.ErrRecordNotFound {
					translation := &models.BrandTranslationModel{
						BrandID: brand.ID,
						Locale:  locale,
						Name:    name,
					}

					if err := db.Create(translation).Error; err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}
