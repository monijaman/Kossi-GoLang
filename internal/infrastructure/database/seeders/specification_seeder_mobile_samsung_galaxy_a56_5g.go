package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileSamsungGalaxyA565g seeds specifications/options for product 'samsung-galaxy-a56-5g'
type SpecificationSeederMobileSamsungGalaxyA565g struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyA565g creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyA565g() *SpecificationSeederMobileSamsungGalaxyA565g {
	return &SpecificationSeederMobileSamsungGalaxyA565g{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy A56 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyA565g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080 x 2400 pixels (~393 ppi density)": "১০৮০ x ২৪০০ pixels (~৩৯৩ ppi density)",
		"120Hz": "১২০Hz",
		"128 GB / 256 GB": "১২৮ GB / ২৫৬ GB",
		"165.1 x 76.4 x 8.4 mm": "১৬৫.১ x ৭৬.৪ x ৮.৪ মিমি",
		"202 g (7.13 oz)": "২০২ g (৭.১৩ oz)",
		"32 MP": "৩২ MP",
		"50 MP + 12 MP + 5 MP": "৫০ MP + ১২ MP + ৫ MP",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6.7 inches": "৬.৭ ইঞ্চি",
		"8 GB": "৮ GB",
		"Android 14, One UI 6": "Android ১৪, One UI ৬",
		"Black, White, Blue": "কালো, সাদা, নীল",
		"Corning Gorilla Glass 5": "Corning Gorilla Glass ৫",
		"Exynos 1380": "Exynos ১৩৮০",
		"GSM / HSPA / LTE / 5G": "GSM / HSPA / LTE / ৫G",
		"Glass front, plastic frame, plastic back": "গ্লাস সামনে, plastic frame, প্লাস্টিক পেছনে",
		"IP67 dust/water resistant (up to 1m for 30 min)": "IP৬৭ dust/water resistant (up to ১m for ৩০ min)",
		"Mali-G68 MP5": "Mali-G৬৮ MP৫",
		"Octa-core (4x2.4 GHz Cortex-A78 & 4x2.0 GHz Cortex-A55)": "Octa-core (৪x২.৪ GHz Cortex-A৭৮ & ৪x২.০ GHz Cortex-A৫৫)",
		"October 2025": "October ২০২৫",
		"Samsung Exynos 1380 (5 nm)": "Samsung Exynos ১৩৮০ (৫ nm)",
		"Super AMOLED, 120Hz": "Super AMOLED, ১২০Hz",
	}
}

// Seed inserts specification records for the product identified by slug 'samsung-galaxy-a56-5g'
func (s *SpecificationSeederMobileSamsungGalaxyA565g) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-a56-5g"

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

	// Override model-specific values for Samsung Galaxy A56 5G
	specs["Display Size"] = "6.7 inches"
	specs["Processor"] = "Exynos 1380"
	specs["Chipset"] = "Samsung Exynos 1380 (5 nm)"
	specs["Cpu Type"] = "Octa-core (4x2.4 GHz Cortex-A78 & 4x2.0 GHz Cortex-A55)"
	specs["Gpu Type"] = "Mali-G68 MP5"
	specs["Ram"] = "8 GB"
	specs["Storage"] = "128 GB / 256 GB"
	specs["Display Type"] = "Super AMOLED, 120Hz"
	specs["Resolution"] = "1080 x 2400 pixels (~393 ppi density)"
	specs["Screen Protection"] = "Corning Gorilla Glass 5"
	specs["Refresh Rate"] = "120Hz"
	specs["Build Material"] = "Glass front, plastic frame, plastic back"
	specs["Weight"] = "202 g (7.13 oz)"
	specs["Dimensions"] = "165.1 x 76.4 x 8.4 mm"
	specs["Water Resistance"] = "IP67 dust/water resistant (up to 1m for 30 min)"
	specs["Network Technology"] = "GSM / HSPA / LTE / 5G"
	specs["Rear Camera"] = "50 MP + 12 MP + 5 MP"
	specs["Front Camera"] = "32 MP"
	specs["Battery"] = "5000 mAh"
	specs["Operating System"] = "Android 14, One UI 6"
	specs["Available Colors"] = "Black, White, Blue"
	specs["Announcement Date"] = "October 2025"
	specs["Device Status"] = "Upcoming"

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
