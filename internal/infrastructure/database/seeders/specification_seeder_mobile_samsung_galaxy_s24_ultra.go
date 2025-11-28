package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileSamsungGalaxyS24Ultra seeds specifications/options for product 'samsung-galaxy-s24-ultra'
type SpecificationSeederMobileSamsungGalaxyS24Ultra struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyS24Ultra creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyS24Ultra() *SpecificationSeederMobileSamsungGalaxyS24Ultra {
	return &SpecificationSeederMobileSamsungGalaxyS24Ultra{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy S24 Ultra"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyS24Ultra) getBanglaTranslations() map[string]string {
	return map[string]string{
		"12 GB": "১২ GB",
		"12 MP": "১২ মেগাপিক্সেল",
		"120Hz": "১২০Hz",
		"1440 x 3120 pixels": "১৪৪০ x ৩১২০ পিক্সেল",
		"162.3 x 79 x 8.6 mm": "১৬২.৩ x ৭৯ x ৮.৬ মিমি",
		"200 MP + 50 MP + 10 MP + 12 MP": "২০০ মেগাপিক্সেল + ৫০ মেগাপিক্সেল + ১০ মেগাপিক্সেল + ১২ মেগাপিক্সেল",
		"232 g": "২৩২ g",
		"256 GB / 512 GB / 1 TB": "২৫৬ জিবি / ৫১২ জিবি / ১ টিবি",
		"5,000 mAh": "৫,০০০ এমএএইচ",
		"5G": "৫G",
		"6.8 inches": "৬.৮ ইঞ্চি",
		"Adreno 750": "অ্যাড্রেনো ৭৫০",
		"Android 14, One UI 6.1": "অ্যান্ড্রয়েড ১৪, ওয়ান ইউআই ৬.১",
		"Dynamic LTPO AMOLED 2X, 120Hz, HDR10+": "Dynamic এলটিপিও অ্যামোলেড ২X, ১২০Hz, এইচডিআর১০+",
		"IP68": "IP৬৮",
		"January 2024": "জানুয়ারি ২০২৪",
		"Qualcomm SM8650-AC Snapdragon 8 Gen 3 (4 nm)": "কোয়ালকম SM৮৬৫০-AC স্ন্যাপড্রাগন ৮ জেন ৩ (৪ ন্যানোমিটার)",
		"Snapdragon 8 Gen 3": "স্ন্যাপড্রাগন ৮ জেন ৩",
		"Titanium Gray, Titanium Black, Titanium Violet, Titanium Yellow": "টাইটানিয়াম ধূসর, টাইটানিয়াম কালো, টাইটানিয়াম ভায়োলেট, টাইটানিয়াম হলুদ",
		"Titanium frame, glass front/back": "টাইটানিয়াম ফ্রেম, গ্লাস সামনে/পেছনে",
	}
}

// Seed inserts specification records for the product identified by slug 'samsung-galaxy-s24-ultra'
func (s *SpecificationSeederMobileSamsungGalaxyS24Ultra) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-s24-ultra"

	var prod models.ProductModel
	if err := db.Where("slug = ?", productSlug).First(&prod).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		return err
	}
	productID := prod.ID

	specs := DefaultMobileSpecs()
	banglaTranslations := s.getBanglaTranslations()

	// Override model-specific values for Samsung Galaxy S24 Ultra
	specs["Display Size"] = "6.8 inches"
	specs["Processor"] = "Snapdragon 8 Gen 3"
	specs["Chipset"] = "Qualcomm SM8650-AC Snapdragon 8 Gen 3 (4 nm)"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "Adreno 750"
	specs["Ram"] = "12 GB"
	specs["Storage"] = "256 GB / 512 GB / 1 TB"
	specs["Display Type"] = "Dynamic LTPO AMOLED 2X, 120Hz, HDR10+"
	specs["Resolution"] = "1440 x 3120 pixels"
	specs["Screen Protection"] = "Corning Gorilla Armor"
	specs["Refresh Rate"] = "120Hz"
	specs["Build Material"] = "Titanium frame, glass front/back"
	specs["Weight"] = "232 g"
	specs["Dimensions"] = "162.3 x 79 x 8.6 mm"
	specs["Water Resistance"] = "IP68"
	specs["Network Technology"] = "5G"
	specs["Rear Camera"] = "200 MP + 50 MP + 10 MP + 12 MP"
	specs["Front Camera"] = "12 MP"
	specs["Battery"] = "5,000 mAh"
	specs["Operating System"] = "Android 14, One UI 6.1"
	specs["Available Colors"] = "Titanium Gray, Titanium Black, Titanium Violet, Titanium Yellow"
	specs["Announcement Date"] = "January 2024"
	specs["Device Status"] = "Available"

	for key, value := range specs {
		sk, err := CreateOrFindSpecificationKey(db, key)
		if err != nil {
			return err
		}

		var existing models.SpecificationModel
		if err := db.Where("product_id = ? AND specification_key_id = ?", productID, sk.ID).First(&existing).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				sModel := &models.SpecificationModel{
					ProductID:          productID,
					SpecificationKeyID: sk.ID,
					Value:              value,
					Status:             1,
				}
				if err := db.Create(sModel).Error; err != nil {
					return err
				}

				// Create Bangla translation for the specification
				banglaValue, exists := banglaTranslations[value]
				if exists && banglaValue != "" {
					var existingTranslation models.SpecificationTranslationModel
					if err := db.Where("specification_id = ? AND locale = ?", sModel.ID, "bn").First(&existingTranslation).Error; err != nil {
						if err == gorm.ErrRecordNotFound {
							translation := &models.SpecificationTranslationModel{
								SpecificationID: sModel.ID,
								Locale:          "bn",
								Value:           banglaValue,
							}
							if err := db.Create(translation).Error; err != nil {
								return err
							}
						} else {
							return err
						}
					}
				}
			} else {
				return err
			}
		} else {
			// If specification already exists, check and create Bangla translation if missing
			banglaValue, exists := banglaTranslations[value]
			if exists && banglaValue != "" {
				var existingTranslation models.SpecificationTranslationModel
				if err := db.Where("specification_id = ? AND locale = ?", existing.ID, "bn").First(&existingTranslation).Error; err != nil {
					if err == gorm.ErrRecordNotFound {
						translation := &models.SpecificationTranslationModel{
							SpecificationID: existing.ID,
							Locale:          "bn",
							Value:           banglaValue,
						}
						if err := db.Create(translation).Error; err != nil {
							return err
						}
					} else {
						return err
					}
				}
			}
		}
	}

	return nil
}
