package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileSamsungGalaxyS23Fe seeds specifications/options for product 'samsung-galaxy-s23-fe'
type SpecificationSeederMobileSamsungGalaxyS23Fe struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyS23Fe creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyS23Fe() *SpecificationSeederMobileSamsungGalaxyS23Fe {
	return &SpecificationSeederMobileSamsungGalaxyS23Fe{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy S23 FE"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyS23Fe) getBanglaTranslations() map[string]string {
	return map[string]string{
		"10 MP": "১০ MP",
		"1080 x 2340 pixels": "১০৮০ x ২৩৪০ pixels",
		"120Hz": "১২০Hz",
		"128 GB / 256 GB": "১২৮ GB / ২৫৬ GB",
		"158 x 76.5 x 8.2 mm": "১৫৮ x ৭৬.৫ x ৮.২ মিমি",
		"209 g": "২০৯ g",
		"4,500 mAh": "৪,৫০০ এমএএইচ",
		"50 MP + 8 MP + 12 MP": "৫০ MP + ৮ MP + ১২ MP",
		"5G": "৫G",
		"6.4 inches": "৬.৪ ইঞ্চি",
		"8 GB": "৮ GB",
		"Aluminum frame, glass front/back": "অ্যালুমিনিয়াম ফ্রেম, গ্লাস সামনে/back",
		"Android 13, upgradable": "Android ১৩, upgradable",
		"Corning Gorilla Glass 5": "Corning Gorilla Glass ৫",
		"Dynamic AMOLED 2X, 120Hz, HDR10+": "Dynamic AMOLED ২X, ১২০Hz, HDR১০+",
		"Exynos 2200 (4 nm)": "Exynos ২২০০ (৪ nm)",
		"Exynos 2200 / Snapdragon 8 Gen 1": "Exynos ২২০০ / Snapdragon ৮ Gen ১",
		"IP68": "IP৬৮",
		"Mint, Cream, Graphite, Purple, Indigo, Tangerine": "Mint, Cream, Graphite, বেগুনি, Indigo, Tangerine",
		"October 2023": "October ২০২৩",
		"Xclipse 920": "Xclipse ৯২০",
	}
}

// Seed inserts specification records for the product identified by slug 'samsung-galaxy-s23-fe'
func (s *SpecificationSeederMobileSamsungGalaxyS23Fe) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-s23-fe"

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

	// Override model-specific values for Samsung Galaxy S23 FE
	specs["Display Size"] = "6.4 inches"
	specs["Processor"] = "Exynos 2200 / Snapdragon 8 Gen 1"
	specs["Chipset"] = "Exynos 2200 (4 nm)"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "Xclipse 920"
	specs["Ram"] = "8 GB"
	specs["Storage"] = "128 GB / 256 GB"
	specs["Display Type"] = "Dynamic AMOLED 2X, 120Hz, HDR10+"
	specs["Resolution"] = "1080 x 2340 pixels"
	specs["Screen Protection"] = "Corning Gorilla Glass 5"
	specs["Refresh Rate"] = "120Hz"
	specs["Build Material"] = "Aluminum frame, glass front/back"
	specs["Weight"] = "209 g"
	specs["Dimensions"] = "158 x 76.5 x 8.2 mm"
	specs["Water Resistance"] = "IP68"
	specs["Network Technology"] = "5G"
	specs["Rear Camera"] = "50 MP + 8 MP + 12 MP"
	specs["Front Camera"] = "10 MP"
	specs["Battery"] = "4,500 mAh"
	specs["Operating System"] = "Android 13, upgradable"
	specs["Available Colors"] = "Mint, Cream, Graphite, Purple, Indigo, Tangerine"
	specs["Announcement Date"] = "October 2023"
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
