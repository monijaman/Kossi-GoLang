package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileSamsungGalaxyZFlip7Fe seeds specifications/options for product 'samsung-galaxy-z-flip-7-fe'
type SpecificationSeederMobileSamsungGalaxyZFlip7Fe struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyZFlip7Fe creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyZFlip7Fe() *SpecificationSeederMobileSamsungGalaxyZFlip7Fe {
	return &SpecificationSeederMobileSamsungGalaxyZFlip7Fe{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy Z Flip 7 FE"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyZFlip7Fe) getBanglaTranslations() map[string]string {
	return map[string]string{
		"10 MP": "১০ মেগাপিক্সেল",
		"1080 x 2640 pixels": "১০৮০ x ২৬৪০ পিক্সেল",
		"12 MP + 12 MP": "১২ মেগাপিক্সেল + ১২ মেগাপিক্সেল",
		"120Hz": "১২০Hz",
		"128 GB / 256 GB": "১২৮ জিবি / ২৫৬ GB",
		"190 g": "১৯০ g",
		"4,000 mAh": "৪,০০০ এমএএইচ",
		"5G": "৫G",
		"6.7 inches (Main) / 3.4 inches (Cover)": "৬.৭ ইঞ্চি (Main) / ৩.৪ ইঞ্চি (Cover)",
		"8 GB": "৮ GB",
		"Aluminum frame, glass front/back": "অ্যালুমিনিয়াম ফ্রেম, গ্লাস সামনে/পেছনে",
		"Android 15, One UI 7": "অ্যান্ড্রয়েড ১৫, ওয়ান ইউআই ৭",
		"Exynos 2400e": "এক্সিনস ২৪০০e",
		"Exynos 2400e (4 nm)": "এক্সিনস ২৪০০e (৪ ন্যানোমিটার)",
		"Foldable Dynamic AMOLED 2X, 120Hz": "Foldable Dynamic অ্যামোলেড ২X, ১২০Hz",
		"IPX8": "IPX৮",
		"July 2025": "জুলাই ২০২৫",
		"Unfolded: 165.1 x 71.9 x 6.9 mm": "Unfolded: ১৬৫.১ x ৭১.৯ x ৬.৯ মিমি",
		"Xclipse 940": "Xclipse ৯৪০",
	}
}

// Seed inserts specification records for the product identified by slug 'samsung-galaxy-z-flip-7-fe'
func (s *SpecificationSeederMobileSamsungGalaxyZFlip7Fe) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-z-flip-7-fe"

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

	// Override model-specific values for Samsung Galaxy Z Flip 7 FE
	specs["Display Size"] = "6.7 inches (Main) / 3.4 inches (Cover)"
	specs["Processor"] = "Exynos 2400e"
	specs["Chipset"] = "Exynos 2400e (4 nm)"
	specs["Cpu Type"] = "Deca-core"
	specs["Gpu Type"] = "Xclipse 940"
	specs["Ram"] = "8 GB"
	specs["Storage"] = "128 GB / 256 GB"
	specs["Display Type"] = "Foldable Dynamic AMOLED 2X, 120Hz"
	specs["Resolution"] = "1080 x 2640 pixels"
	specs["Screen Protection"] = "Corning Gorilla Glass Victus+"
	specs["Refresh Rate"] = "120Hz"
	specs["Build Material"] = "Aluminum frame, glass front/back"
	specs["Weight"] = "190 g"
	specs["Dimensions"] = "Unfolded: 165.1 x 71.9 x 6.9 mm"
	specs["Water Resistance"] = "IPX8"
	specs["Network Technology"] = "5G"
	specs["Rear Camera"] = "12 MP + 12 MP"
	specs["Front Camera"] = "10 MP"
	specs["Battery"] = "4,000 mAh"
	specs["Operating System"] = "Android 15, One UI 7"
	specs["Available Colors"] = "Graphite, Cream, Lavender"
	specs["Announcement Date"] = "July 2025"
	specs["Device Status"] = "Rumored"

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
