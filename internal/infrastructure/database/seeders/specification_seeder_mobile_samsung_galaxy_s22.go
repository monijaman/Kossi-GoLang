package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileSamsungGalaxyS22 seeds specifications/options for product 'samsung-galaxy-s22'
type SpecificationSeederMobileSamsungGalaxyS22 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyS22 creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyS22() *SpecificationSeederMobileSamsungGalaxyS22 {
	return &SpecificationSeederMobileSamsungGalaxyS22{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy S22"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyS22) getBanglaTranslations() map[string]string {
	return map[string]string{
		"10 MP": "১০ মেগাপিক্সেল",
		"1080 x 2340 pixels": "১০৮০ x ২৩৪০ পিক্সেল",
		"120Hz": "১২০Hz",
		"128 GB / 256 GB": "১২৮ জিবি / ২৫৬ GB",
		"146 x 70.6 x 7.6 mm": "১৪৬ x ৭০.৬ x ৭.৬ মিমি",
		"167 g": "১৬৭ g",
		"3,700 mAh": "৩,৭০০ এমএএইচ",
		"50 MP + 10 MP + 12 MP": "৫০ মেগাপিক্সেল + ১০ মেগাপিক্সেল + ১২ মেগাপিক্সেল",
		"5G": "৫G",
		"6.1 inches": "৬.১ ইঞ্চি",
		"8 GB": "৮ GB",
		"Aluminum frame, glass front/back": "অ্যালুমিনিয়াম ফ্রেম, গ্লাস সামনে/পেছনে",
		"Android 12, upgradable": "অ্যান্ড্রয়েড ১২, আপগ্রেডযোগ্য",
		"Dynamic AMOLED 2X, 120Hz, HDR10+": "Dynamic অ্যামোলেড ২X, ১২০Hz, এইচডিআর১০+",
		"Exynos 2200 (4 nm)": "এক্সিনস ২২০০ (৪ ন্যানোমিটার)",
		"Exynos 2200 / Snapdragon 8 Gen 1": "এক্সিনস ২২০০ / স্ন্যাপড্রাগন ৮ জেন ১",
		"February 2022": "ফেব্রুয়ারি ২০২২",
		"IP68": "IP৬৮",
		"Phantom Black, White, Pink Gold, Green": "ফ্যান্টম কালো, সাদা, গোলাপী সোনালী, সবুজ",
		"Xclipse 920": "Xclipse ৯২০",
	}
}

// Seed inserts specification records for the product identified by slug 'samsung-galaxy-s22'
func (s *SpecificationSeederMobileSamsungGalaxyS22) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-s22"

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

	// Override model-specific values for Samsung Galaxy S22
	specs["Display Size"] = "6.1 inches"
	specs["Processor"] = "Exynos 2200 / Snapdragon 8 Gen 1"
	specs["Chipset"] = "Exynos 2200 (4 nm)"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "Xclipse 920"
	specs["Ram"] = "8 GB"
	specs["Storage"] = "128 GB / 256 GB"
	specs["Display Type"] = "Dynamic AMOLED 2X, 120Hz, HDR10+"
	specs["Resolution"] = "1080 x 2340 pixels"
	specs["Screen Protection"] = "Corning Gorilla Glass Victus+"
	specs["Refresh Rate"] = "120Hz"
	specs["Build Material"] = "Aluminum frame, glass front/back"
	specs["Weight"] = "167 g"
	specs["Dimensions"] = "146 x 70.6 x 7.6 mm"
	specs["Water Resistance"] = "IP68"
	specs["Network Technology"] = "5G"
	specs["Rear Camera"] = "50 MP + 10 MP + 12 MP"
	specs["Front Camera"] = "10 MP"
	specs["Battery"] = "3,700 mAh"
	specs["Operating System"] = "Android 12, upgradable"
	specs["Available Colors"] = "Phantom Black, White, Pink Gold, Green"
	specs["Announcement Date"] = "February 2022"
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
