package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileSamsungGalaxyA07 seeds specifications/options for product 'samsung-galaxy-a07'
type SpecificationSeederMobileSamsungGalaxyA07 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyA07 creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyA07() *SpecificationSeederMobileSamsungGalaxyA07 {
	return &SpecificationSeederMobileSamsungGalaxyA07{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy A07"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyA07) getBanglaTranslations() map[string]string {
	return map[string]string{
		"13 MP + 2 MP": "১৩ মেগাপিক্সেল + ২ মেগাপিক্সেল",
		"164.2 x 75.9 x 9.1 mm": "১৬৪.২ x ৭৫.৯ x ৯.১ মিমি",
		"186 g (6.56 oz)": "১৮৬ গ্রাম (৬.৫৬ oz)",
		"4 GB / 6 GB": "৪ জিবি / ৬ GB",
		"5 MP": "৫ মেগাপিক্সেল",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6.5 inches": "৬.৫ ইঞ্চি",
		"60Hz": "৬০Hz",
		"64 GB / 128 GB": "৬৪ জিবি / ১২৮ GB",
		"720 x 1600 pixels (~270 ppi density)": "৭২০ x ১৬০০ পিক্সেল (~২৭০ পিপিআই density)",
		"Android 14, One UI Core 6": "অ্যান্ড্রয়েড ১৪, ওয়ান ইউআই Cবাe ৬",
		"Black, Green, White": "কালো, সবুজ, সাদা",
		"Exynos 850": "এক্সিনস ৮৫০",
		"June 2025": "জুন ২০২৫",
		"Mali-G52": "মালি-G৫২",
		"Octa-core (4x2.0 GHz Cortex-A55 & 4x2.0 GHz Cortex-A55)": "অক্টা-কোর (৪x২.০ গিগাহার্টজ Cবাtex-A৫৫ & ৪x২.০ গিগাহার্টজ Cবাtex-A৫৫)",
		"Plastic frame, plastic back": "Plastic ফ্রেম, প্লাস্টিক পেছনে",
		"Samsung Exynos 850 (8 nm)": "Samsung এক্সিনস ৮৫০ (৮ ন্যানোমিটার)",
	}
}

// Seed inserts specification records for the product identified by slug 'samsung-galaxy-a07'
func (s *SpecificationSeederMobileSamsungGalaxyA07) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-a07"

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

	// Override model-specific values for Samsung Galaxy A07
	specs["Display Size"] = "6.5 inches"
	specs["Processor"] = "Exynos 850"
	specs["Chipset"] = "Samsung Exynos 850 (8 nm)"
	specs["Cpu Type"] = "Octa-core (4x2.0 GHz Cortex-A55 & 4x2.0 GHz Cortex-A55)"
	specs["Gpu Type"] = "Mali-G52"
	specs["Ram"] = "4 GB / 6 GB"
	specs["Storage"] = "64 GB / 128 GB"
	specs["Display Type"] = "PLS LCD"
	specs["Resolution"] = "720 x 1600 pixels (~270 ppi density)"
	specs["Screen Protection"] = "No"
	specs["Refresh Rate"] = "60Hz"
	specs["Build Material"] = "Plastic frame, plastic back"
	specs["Weight"] = "186 g (6.56 oz)"
	specs["Dimensions"] = "164.2 x 75.9 x 9.1 mm"
	specs["Water Resistance"] = "No"
	specs["Network Technology"] = "GSM / HSPA / LTE"
	specs["Rear Camera"] = "13 MP + 2 MP"
	specs["Front Camera"] = "5 MP"
	specs["Battery"] = "5000 mAh"
	specs["Operating System"] = "Android 14, One UI Core 6"
	specs["Available Colors"] = "Black, Green, White"
	specs["Announcement Date"] = "June 2025"
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
