package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileSamsungGalaxyZFlip5 seeds specifications/options for product 'samsung-galaxy-z-flip-5'
type SpecificationSeederMobileSamsungGalaxyZFlip5 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyZFlip5 creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyZFlip5() *SpecificationSeederMobileSamsungGalaxyZFlip5 {
	return &SpecificationSeederMobileSamsungGalaxyZFlip5{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy Z Flip 5"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyZFlip5) getBanglaTranslations() map[string]string {
	return map[string]string{
		"10 MP": "১০ MP",
		"1080 x 2640 pixels": "১০৮০ x ২৬৪০ pixels",
		"12 MP + 12 MP": "১২ MP + ১২ MP",
		"120Hz": "১২০Hz",
		"187 g": "১৮৭ g",
		"256 GB / 512 GB": "২৫৬ GB / ৫১২ GB",
		"3,700 mAh": "৩,৭০০ এমএএইচ",
		"5G": "৫G",
		"6.7 inches (Main) / 3.4 inches (Cover)": "৬.৭ ইঞ্চি (Main) / ৩.৪ ইঞ্চি (Cover)",
		"8 GB": "৮ GB",
		"Adreno 740": "Adreno ৭৪০",
		"Android 13, upgradable": "Android ১৩, upgradable",
		"Corning Gorilla Glass Victus 2": "Corning Gorilla Glass Victus ২",
		"Foldable Dynamic AMOLED 2X, 120Hz, HDR10+": "Foldable Dynamic AMOLED ২X, ১২০Hz, HDR১০+",
		"IPX8": "IPX৮",
		"July 2023": "July ২০২৩",
		"Mint, Graphite, Cream, Lavender, Gray, Blue, Green, Yellow": "Mint, Graphite, Cream, Lavender, ধূসর, নীল, সবুজ, হলুদ",
		"Plastic front (open), glass back (Gorilla Glass Victus 2), aluminum frame": "Plastic front (open), গ্লাস পেছনে (Gorilla Glass Victus ২), অ্যালুমিনিয়াম ফ্রেম",
		"Qualcomm SM8550-AC Snapdragon 8 Gen 2 (4 nm)": "Qualcomm SM৮৫৫০-AC Snapdragon ৮ Gen ২ (৪ nm)",
		"Snapdragon 8 Gen 2": "Snapdragon ৮ Gen ২",
		"Unfolded: 165.1 x 71.9 x 6.9 mm / Folded: 85.1 x 71.9 x 15.1 mm": "Unfolded: ১৬৫.১ x ৭১.৯ x ৬.৯ মিমি / Folded: ৮৫.১ x ৭১.৯ x ১৫.১ মিমি",
	}
}

// Seed inserts specification records for the product identified by slug 'samsung-galaxy-z-flip-5'
func (s *SpecificationSeederMobileSamsungGalaxyZFlip5) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-z-flip-5"

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

	// Override model-specific values for Samsung Galaxy Z Flip 5
	specs["Display Size"] = "6.7 inches (Main) / 3.4 inches (Cover)"
	specs["Processor"] = "Snapdragon 8 Gen 2"
	specs["Chipset"] = "Qualcomm SM8550-AC Snapdragon 8 Gen 2 (4 nm)"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "Adreno 740"
	specs["Ram"] = "8 GB"
	specs["Storage"] = "256 GB / 512 GB"
	specs["Display Type"] = "Foldable Dynamic AMOLED 2X, 120Hz, HDR10+"
	specs["Resolution"] = "1080 x 2640 pixels"
	specs["Screen Protection"] = "Corning Gorilla Glass Victus 2"
	specs["Refresh Rate"] = "120Hz"
	specs["Build Material"] = "Plastic front (open), glass back (Gorilla Glass Victus 2), aluminum frame"
	specs["Weight"] = "187 g"
	specs["Dimensions"] = "Unfolded: 165.1 x 71.9 x 6.9 mm / Folded: 85.1 x 71.9 x 15.1 mm"
	specs["Water Resistance"] = "IPX8"
	specs["Network Technology"] = "5G"
	specs["Rear Camera"] = "12 MP + 12 MP"
	specs["Front Camera"] = "10 MP"
	specs["Battery"] = "3,700 mAh"
	specs["Operating System"] = "Android 13, upgradable"
	specs["Available Colors"] = "Mint, Graphite, Cream, Lavender, Gray, Blue, Green, Yellow"
	specs["Announcement Date"] = "July 2023"
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
