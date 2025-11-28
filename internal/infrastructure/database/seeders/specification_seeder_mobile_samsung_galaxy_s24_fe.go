package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileSamsungGalaxyS24Fe seeds specifications/options for product 'samsung-galaxy-s24-fe'
type SpecificationSeederMobileSamsungGalaxyS24Fe struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyS24Fe creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyS24Fe() *SpecificationSeederMobileSamsungGalaxyS24Fe {
	return &SpecificationSeederMobileSamsungGalaxyS24Fe{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy S24 FE"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyS24Fe) getBanglaTranslations() map[string]string {
	return map[string]string{
		"10 MP": "১০ মেগাপিক্সেল",
		"1080 x 2340 pixels": "১০৮০ x ২৩৪০ পিক্সেল",
		"120Hz": "১২০Hz",
		"128 GB / 256 GB / 512 GB": "১২৮ জিবি / ২৫৬ জিবি / ৫১২ GB",
		"162 x 77.3 x 8 mm": "১৬২ x ৭৭.৩ x ৮ মিমি",
		"213 g": "২১৩ g",
		"4,700 mAh": "৪,৭০০ এমএএইচ",
		"50 MP + 8 MP + 12 MP": "৫০ মেগাপিক্সেল + ৮ মেগাপিক্সেল + ১২ মেগাপিক্সেল",
		"5G": "৫G",
		"6.7 inches": "৬.৭ ইঞ্চি",
		"8 GB": "৮ GB",
		"Aluminum frame, glass front/back": "অ্যালুমিনিয়াম ফ্রেম, গ্লাস সামনে/পেছনে",
		"Android 14, One UI 6.1": "অ্যান্ড্রয়েড ১৪, ওয়ান ইউআই ৬.১",
		"Blue, Graphite, Gray, Mint, Yellow": "নীল, গ্রাফাইট, ধূসর, মিন্ট, হলুদ",
		"Dynamic AMOLED 2X, 120Hz, HDR10+": "Dynamic অ্যামোলেড ২X, ১২০Hz, এইচডিআর১০+",
		"Exynos 2400e": "এক্সিনস ২৪০০e",
		"Exynos 2400e (4 nm)": "এক্সিনস ২৪০০e (৪ ন্যানোমিটার)",
		"IP68": "IP৬৮",
		"September 2024": "সেপ্টেম্বর ২০২৪",
		"Xclipse 940": "Xclipse ৯৪০",
	}
}

// Seed inserts specification records for the product identified by slug 'samsung-galaxy-s24-fe'
func (s *SpecificationSeederMobileSamsungGalaxyS24Fe) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-s24-fe"

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

	// Override model-specific values for Samsung Galaxy S24 FE
	specs["Display Size"] = "6.7 inches"
	specs["Processor"] = "Exynos 2400e"
	specs["Chipset"] = "Exynos 2400e (4 nm)"
	specs["Cpu Type"] = "Deca-core"
	specs["Gpu Type"] = "Xclipse 940"
	specs["Ram"] = "8 GB"
	specs["Storage"] = "128 GB / 256 GB / 512 GB"
	specs["Display Type"] = "Dynamic AMOLED 2X, 120Hz, HDR10+"
	specs["Resolution"] = "1080 x 2340 pixels"
	specs["Screen Protection"] = "Corning Gorilla Glass Victus+"
	specs["Refresh Rate"] = "120Hz"
	specs["Build Material"] = "Aluminum frame, glass front/back"
	specs["Weight"] = "213 g"
	specs["Dimensions"] = "162 x 77.3 x 8 mm"
	specs["Water Resistance"] = "IP68"
	specs["Network Technology"] = "5G"
	specs["Rear Camera"] = "50 MP + 8 MP + 12 MP"
	specs["Front Camera"] = "10 MP"
	specs["Battery"] = "4,700 mAh"
	specs["Operating System"] = "Android 14, One UI 6.1"
	specs["Available Colors"] = "Blue, Graphite, Gray, Mint, Yellow"
	specs["Announcement Date"] = "September 2024"
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
