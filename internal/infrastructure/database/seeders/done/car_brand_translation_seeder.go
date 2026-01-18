package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// CarBrandTranslationSeeder handles seeding car brand translations
type CarBrandTranslationSeeder struct {
	BaseSeeder
}

// NewCarBrandTranslationSeeder creates a new car brand translation seeder
func NewCarBrandTranslationSeeder() *CarBrandTranslationSeeder {
	return &CarBrandTranslationSeeder{
		BaseSeeder: BaseSeeder{name: "Car Brand Translations"},
	}
}

// Seed implements the Seeder interface
func (cbts *CarBrandTranslationSeeder) Seed(db *gorm.DB) error {
	// Car brand translations (English -> Bangla)
	// Keys are slugs (lowercase, spaces replaced by hyphens)
	translations := map[string]map[string]string{
		"toyota":             {"bn": "টয়োটা"},
		"honda":              {"bn": "হোন্ডা"},
		"nissan":             {"bn": "নিসান"},
		"mitsubishi":         {"bn": "মিতসুবিশি"},
		"mazda":              {"bn": "মাজদা"},
		"subaru":             {"bn": "সুবারু"},
		"suzuki":             {"bn": "সুজুকি"},
		"daihatsu":           {"bn": "ডাইহাতসু"},
		"hyundai":            {"bn": "হুন্দাই"},
		"kia":                {"bn": "কিয়া"},
		"proton":             {"bn": "প্রোটন"},
		"bmw":                {"bn": "বিএমডব্লিউ"},
		"mercedes-benz":      {"bn": "মার্সিডিজ-বেঞ্জ"},
		"audi":               {"bn": "অডি"},
		"volkswagen":         {"bn": "ভক্সওয়াগন"},
		"tata":               {"bn": "টাটা"},
		"ford":               {"bn": "ফোর্ড"},
		"changan":            {"bn": "চাঙ্গান"},
		"chery":              {"bn": "চেরি"},
		"chevrolet":          {"bn": "শেভ্রোলেট"},
		"geely":              {"bn": "জিলি"},
		"gmc":                {"bn": "জিএমসি"},
		"great-wall-/-haval": {"bn": "গ্রেট ওয়াল / হাভাল"}, // slug might need adjustment based on generation logic
		"isuzu":              {"bn": "ইসুজু"},
		"jaguar":             {"bn": "জাগুয়ার"},
		"jeep":               {"bn": "জিপ"},
		"lexus":              {"bn": "লেক্সাস"},
		"mahindra":           {"bn": "মাহিন্দ্রা"},
		"mg":                 {"bn": "এমজি"},
		"mini":               {"bn": "মিনি"},
		"peugeot":            {"bn": "পিউজো"},
		"porsche":            {"bn": "পোরশে"},
		"range-rover":        {"bn": "রেঞ্জ রোভার"},
		"ssangyong":          {"bn": "স্যাংইয়ং"},
		"volvo":              {"bn": "ভলভো"},
	}

	for slug, locales := range translations {
		// Find the brand by slug
		var brand models.BrandModel
		// Note: The slug generation in CarSeeder might produce different slugs for complex names.
		// We should try to match as best as possible.
		// "Great Wall / Haval" -> "great-wall-/-haval" (based on simple generator)
		
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
				} else if transResult.Error == nil {
					// Update existing translation if it exists (e.g. if CarSeeder created it with English name)
					existingTranslation.Name = name
					if err := db.Save(&existingTranslation).Error; err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}
