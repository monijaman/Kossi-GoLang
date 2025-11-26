package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// BrandTranslationInclusionSeeder ensures all brands have translations in the database
type BrandTranslationInclusionSeeder struct {
	BaseSeeder
}

// NewBrandTranslationInclusionSeeder creates a new seeder for ensuring brand translations
func NewBrandTranslationInclusionSeeder() *BrandTranslationInclusionSeeder {
	return &BrandTranslationInclusionSeeder{
		BaseSeeder: BaseSeeder{name: "Brand Translation Inclusion Seeder"},
	}
}

// Seed implements the Seeder interface
func (btis *BrandTranslationInclusionSeeder) Seed(db *gorm.DB) error {
	// Fetch all brands from the database
	var brands []models.BrandModel
	if err := db.Find(&brands).Error; err != nil {
		return err
	}

	// Define default translations (can be extended for other locales)
	defaultLocale := "bn"
	translations := map[string]string{
		"Toyota":         "টয়োটা",
		"Honda":          "হোন্ডা",
		"Nissan":         "নিসান",
		"Mitsubishi":     "মিতসুবিশি",
		"Mazda":          "মাজদা",
		"Subaru":         "সুবারু",
		"Suzuki":         "সুজুকি",
		"Daihatsu":       "ডাইহাতসু",
		"Hyundai":        "হুন্দাই",
		"Kia":            "কিয়া",
		"Proton":         "প্রোটন",
		"BMW":            "বিএমডব্লিউ",
		"Mercedes-Benz":  "মার্সিডিজ-বেঞ্জ",
		"Audi":           "অডি",
		"Volkswagen":     "ভক্সওয়াগন",
		"Tata":           "টাটা",
		"Ford":           "ফোর্ড",
		"Chevrolet":      "শেভ্রোলেট",
		"Jaguar":         "জাগুয়ার",
		"Volvo":          "ভলভো",
		"KTM":            "কেটিএম",
		"Hero":           "হিরো",
		"Konka":          "কনকা",
		"Robi":           "রবি",
		"SYM":            "এসওয়াইএম",
		"Teletalk":       "টেলিটক",
		"Lifan":          "লিফান",
		"Benelli":        "বেনেলি",
		"Road Prince":    "রোড প্রিন্স",
		"MyOne":          "মাইওয়ান",
		"Haojue":         "হাওজুয়ে",
		"Victor-R":       "ভিক্টর-আর",
		"Atlas Zongshen": "অ্যাটলাস জংশেন",
		"Revoo":          "রেভু",
		"Zongshen":       "জংশেন",
		"PHP":            "পিএইচপি",
		"Roadmaster":     "রোডমাস্টার",
		"Airtel":         "এয়ারটেল",
		"Banglalink":     "বাংলালিংক",
		"ECO+":           "ইকো+",
		"Haiko":          "হাইকো",
		"Vespa":          "ভেসপা",
		"Grameenphone":   "গ্রামীণফোন",
	}

	// Added translations for the requested brands
	translations = map[string]string{
		"Toyota":         "টয়োটা",
		"Honda":          "হোন্ডা",
		"Nissan":         "নিসান",
		"Mitsubishi":     "মিতসুবিশি",
		"Mazda":          "মাজদা",
		"Subaru":         "সুবারু",
		"Suzuki":         "সুজুকি",
		"Daihatsu":       "ডাইহাতসু",
		"Hyundai":        "হুন্দাই",
		"Kia":            "কিয়া",
		"Proton":         "প্রোটন",
		"BMW":            "বিএমডব্লিউ",
		"Mercedes-Benz":  "মার্সিডিজ-বেঞ্জ",
		"Audi":           "অডি",
		"Volkswagen":     "ভক্সওয়াগন",
		"Tata":           "টাটা",
		"Ford":           "ফোর্ড",
		"Chevrolet":      "শেভ্রোলেট",
		"Jaguar":         "জাগুয়ার",
		"Volvo":          "ভলভো",
		"KTM":            "কেটিএম",
		"Hero":           "হিরো",
		"Konka":          "কনকা",
		"Robi":           "রবি",
		"SYM":            "এসওয়াইএম",
		"Teletalk":       "টেলিটক",
		"Lifan":          "লিফান",
		"Benelli":        "বেনেলি",
		"Road Prince":    "রোড প্রিন্স",
		"MyOne":          "মাইওয়ান",
		"Haojue":         "হাওজুয়ে",
		"Victor-R":       "ভিক্টর-আর",
		"Atlas Zongshen": "অ্যাটলাস জংশেন",
		"Revoo":          "রেভু",
		"Zongshen":       "জংশেন",
		"PHP":            "পিএইচপি",
		"Roadmaster":     "রোডমাস্টার",
		"Airtel":         "এয়ারটেল",
		"Banglalink":     "বাংলালিংক",
		"ECO+":           "ইকো+",
		"Haiko":          "হাইকো",
		"Vespa":          "ভেসপা",
		"Grameenphone":   "গ্রামীণফোন",
	}

	// Ensure translations exist for each brand
	for _, brand := range brands {
		if translatedName, exists := translations[brand.Name]; exists {
			var existingTranslation models.BrandTranslationModel
			result := db.Where("brand_id = ? AND locale = ?", brand.ID, defaultLocale).First(&existingTranslation)

			if result.Error == gorm.ErrRecordNotFound {
				// Create new translation if it doesn't exist
				newTranslation := models.BrandTranslationModel{
					BrandID: brand.ID,
					Locale:  defaultLocale,
					Name:    translatedName,
				}
				if err := db.Create(&newTranslation).Error; err != nil {
					return err
				}
			}
		}
	}

	return nil
}
