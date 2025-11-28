package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileSamsungGalaxyS24 seeds specifications/options for product 'samsung-galaxy-s24'
type SpecificationSeederMobileSamsungGalaxyS24 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyS24 creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyS24() *SpecificationSeederMobileSamsungGalaxyS24 {
	return &SpecificationSeederMobileSamsungGalaxyS24{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy S24"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyS24) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080 x 2340 pixels": "১০৮০ x ২৩৪০ পিক্সেল",
		"12 MP": "১২ মেগাপিক্সেল",
		"120Hz": "১২০Hz",
		"128 GB / 256 GB / 512 GB": "১২৮ জিবি / ২৫৬ জিবি / ৫১২ GB",
		"147 x 70.6 x 7.6 mm": "১৪৭ x ৭০.৬ x ৭.৬ মিমি",
		"167 g": "১৬৭ g",
		"4,000 mAh": "৪,০০০ এমএএইচ",
		"50 MP + 10 MP + 12 MP": "৫০ মেগাপিক্সেল + ১০ মেগাপিক্সেল + ১২ মেগাপিক্সেল",
		"5G": "৫G",
		"6.2 inches": "৬.২ ইঞ্চি",
		"8 GB": "৮ GB",
		"Adreno 750 / Xclipse 940": "অ্যাড্রেনো ৭৫০ / Xclipse ৯৪০",
		"Aluminum frame, glass front/back": "অ্যালুমিনিয়াম ফ্রেম, গ্লাস সামনে/পেছনে",
		"Android 14, One UI 6.1": "অ্যান্ড্রয়েড ১৪, ওয়ান ইউআই ৬.১",
		"Corning Gorilla Glass Victus 2": "কর্নিং গরিলা গ্লাস ভিক্টাস ২",
		"Dynamic LTPO AMOLED 2X, 120Hz, HDR10+": "Dynamic এলটিপিও অ্যামোলেড ২X, ১২০Hz, এইচডিআর১০+",
		"IP68": "IP৬৮",
		"January 2024": "জানুয়ারি ২০২৪",
		"Onyx Black, Marble Grey, Cobalt Violet, Amber Yellow": "অনিক্স কালো, Marble ধূসর, Cobalt ভায়োলেট, Amber হলুদ",
		"Qualcomm SM8650-AC Snapdragon 8 Gen 3 (4 nm)": "কোয়ালকম SM৮৬৫০-AC স্ন্যাপড্রাগন ৮ জেন ৩ (৪ ন্যানোমিটার)",
		"Snapdragon 8 Gen 3 / Exynos 2400": "স্ন্যাপড্রাগন ৮ জেন ৩ / এক্সিনস ২৪০০",
	}
}

// Seed inserts specification records for the product identified by slug 'samsung-galaxy-s24'
func (s *SpecificationSeederMobileSamsungGalaxyS24) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-s24"

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

	// Override model-specific values for Samsung Galaxy S24
	specs["Display Size"] = "6.2 inches"
	specs["Processor"] = "Snapdragon 8 Gen 3 / Exynos 2400"
	specs["Chipset"] = "Qualcomm SM8650-AC Snapdragon 8 Gen 3 (4 nm)"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "Adreno 750 / Xclipse 940"
	specs["Ram"] = "8 GB"
	specs["Storage"] = "128 GB / 256 GB / 512 GB"
	specs["Display Type"] = "Dynamic LTPO AMOLED 2X, 120Hz, HDR10+"
	specs["Resolution"] = "1080 x 2340 pixels"
	specs["Screen Protection"] = "Corning Gorilla Glass Victus 2"
	specs["Refresh Rate"] = "120Hz"
	specs["Build Material"] = "Aluminum frame, glass front/back"
	specs["Weight"] = "167 g"
	specs["Dimensions"] = "147 x 70.6 x 7.6 mm"
	specs["Water Resistance"] = "IP68"
	specs["Network Technology"] = "5G"
	specs["Rear Camera"] = "50 MP + 10 MP + 12 MP"
	specs["Front Camera"] = "12 MP"
	specs["Battery"] = "4,000 mAh"
	specs["Operating System"] = "Android 14, One UI 6.1"
	specs["Available Colors"] = "Onyx Black, Marble Grey, Cobalt Violet, Amber Yellow"
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
