package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileSamsungGalaxyS22Ultra seeds specifications/options for product 'samsung-galaxy-s22-ultra'
type SpecificationSeederMobileSamsungGalaxyS22Ultra struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyS22Ultra creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyS22Ultra() *SpecificationSeederMobileSamsungGalaxyS22Ultra {
	return &SpecificationSeederMobileSamsungGalaxyS22Ultra{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy S22 Ultra"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyS22Ultra) getBanglaTranslations() map[string]string {
	return map[string]string{
		"108 MP + 10 MP + 10 MP + 12 MP": "১০৮ MP + ১০ MP + ১০ MP + ১২ MP",
		"120Hz": "১২০Hz",
		"128 GB / 256 GB / 512 GB / 1 TB": "১২৮ GB / ২৫৬ GB / ৫১২ GB / ১ TB",
		"1440 x 3088 pixels": "১৪৪০ x ৩০৮৮ pixels",
		"163.3 x 77.9 x 8.9 mm": "১৬৩.৩ x ৭৭.৯ x ৮.৯ মিমি",
		"228 g": "২২৮ g",
		"40 MP": "৪০ MP",
		"5,000 mAh": "৫,০০০ এমএএইচ",
		"5G": "৫G",
		"6.8 inches": "৬.৮ ইঞ্চি",
		"8 GB / 12 GB": "৮ GB / ১২ GB",
		"Aluminum frame, glass front/back": "অ্যালুমিনিয়াম ফ্রেম, গ্লাস সামনে/back",
		"Android 12, upgradable": "Android ১২, upgradable",
		"Dynamic AMOLED 2X, 120Hz, HDR10+": "Dynamic AMOLED ২X, ১২০Hz, HDR১০+",
		"Exynos 2200 (4 nm)": "Exynos ২২০০ (৪ nm)",
		"Exynos 2200 / Snapdragon 8 Gen 1": "Exynos ২২০০ / Snapdragon ৮ Gen ১",
		"February 2022": "February ২০২২",
		"IP68": "IP৬৮",
		"Phantom Black, White, Burgundy, Green": "Phantom কালো, সাদা, Burgundy, সবুজ",
		"Xclipse 920": "Xclipse ৯২০",
	}
}

// Seed inserts specification records for the product identified by slug 'samsung-galaxy-s22-ultra'
func (s *SpecificationSeederMobileSamsungGalaxyS22Ultra) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-s22-ultra"

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

	// Override model-specific values for Samsung Galaxy S22 Ultra
	specs["Display Size"] = "6.8 inches"
	specs["Processor"] = "Exynos 2200 / Snapdragon 8 Gen 1"
	specs["Chipset"] = "Exynos 2200 (4 nm)"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "Xclipse 920"
	specs["Ram"] = "8 GB / 12 GB"
	specs["Storage"] = "128 GB / 256 GB / 512 GB / 1 TB"
	specs["Display Type"] = "Dynamic AMOLED 2X, 120Hz, HDR10+"
	specs["Resolution"] = "1440 x 3088 pixels"
	specs["Screen Protection"] = "Corning Gorilla Glass Victus+"
	specs["Refresh Rate"] = "120Hz"
	specs["Build Material"] = "Aluminum frame, glass front/back"
	specs["Weight"] = "228 g"
	specs["Dimensions"] = "163.3 x 77.9 x 8.9 mm"
	specs["Water Resistance"] = "IP68"
	specs["Network Technology"] = "5G"
	specs["Rear Camera"] = "108 MP + 10 MP + 10 MP + 12 MP"
	specs["Front Camera"] = "40 MP"
	specs["Battery"] = "5,000 mAh"
	specs["Operating System"] = "Android 12, upgradable"
	specs["Available Colors"] = "Phantom Black, White, Burgundy, Green"
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
