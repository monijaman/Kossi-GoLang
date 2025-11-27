package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileSamsungGalaxyZFold5 seeds specifications/options for product 'samsung-galaxy-z-fold-5'
type SpecificationSeederMobileSamsungGalaxyZFold5 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyZFold5 creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyZFold5() *SpecificationSeederMobileSamsungGalaxyZFold5 {
	return &SpecificationSeederMobileSamsungGalaxyZFold5{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy Z Fold 5"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyZFold5) getBanglaTranslations() map[string]string {
	return map[string]string{
		"12 GB": "১২ GB",
		"120Hz": "১২০Hz",
		"1812 x 2176 pixels": "১৮১২ x ২১৭৬ pixels",
		"253 g": "২৫৩ g",
		"256 GB / 512 GB / 1 TB": "২৫৬ GB / ৫১২ GB / ১ TB",
		"4 MP (Under Display) + 10 MP (Cover)": "৪ MP (ডিসপ্লের নিচে) + ১০ MP (Cover)",
		"4,400 mAh": "৪,৪০০ এমএএইচ",
		"50 MP + 10 MP + 12 MP": "৫০ MP + ১০ MP + ১২ MP",
		"5G": "৫G",
		"7.6 inches (Main) / 6.2 inches (Cover)": "৭.৬ ইঞ্চি (Main) / ৬.২ ইঞ্চি (Cover)",
		"Adreno 740": "Adreno ৭৪০",
		"Android 13, upgradable": "Android ১৩, upgradable",
		"Corning Gorilla Glass Victus 2": "Corning Gorilla Glass Victus ২",
		"Foldable Dynamic AMOLED 2X, 120Hz, HDR10+": "Foldable Dynamic AMOLED ২X, ১২০Hz, HDR১০+",
		"Glass front (Gorilla Glass Victus 2), glass back (Victus 2), aluminum frame": "গ্লাস সামনে (Gorilla Glass Victus ২), গ্লাস পেছনে (Victus ২), অ্যালুমিনিয়াম ফ্রেম",
		"IPX8": "IPX৮",
		"Icy Blue, Phantom Black, Cream, Gray, Blue": "Icy নীল, Phantom কালো, Cream, ধূসর, নীল",
		"July 2023": "July ২০২৩",
		"Qualcomm SM8550-AC Snapdragon 8 Gen 2 (4 nm)": "Qualcomm SM৮৫৫০-AC Snapdragon ৮ Gen ২ (৪ nm)",
		"Snapdragon 8 Gen 2": "Snapdragon ৮ Gen ২",
		"Unfolded: 154.9 x 129.9 x 6.1 mm / Folded: 154.9 x 67.1 x 13.4 mm": "Unfolded: ১৫৪.৯ x ১২৯.৯ x ৬.১ মিমি / Folded: ১৫৪.৯ x ৬৭.১ x ১৩.৪ মিমি",
	}
}

// Seed inserts specification records for the product identified by slug 'samsung-galaxy-z-fold-5'
func (s *SpecificationSeederMobileSamsungGalaxyZFold5) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-z-fold-5"

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

	// Override model-specific values for Samsung Galaxy Z Fold 5
	specs["Display Size"] = "7.6 inches (Main) / 6.2 inches (Cover)"
	specs["Processor"] = "Snapdragon 8 Gen 2"
	specs["Chipset"] = "Qualcomm SM8550-AC Snapdragon 8 Gen 2 (4 nm)"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "Adreno 740"
	specs["Ram"] = "12 GB"
	specs["Storage"] = "256 GB / 512 GB / 1 TB"
	specs["Display Type"] = "Foldable Dynamic AMOLED 2X, 120Hz, HDR10+"
	specs["Resolution"] = "1812 x 2176 pixels"
	specs["Screen Protection"] = "Corning Gorilla Glass Victus 2"
	specs["Refresh Rate"] = "120Hz"
	specs["Build Material"] = "Glass front (Gorilla Glass Victus 2), glass back (Victus 2), aluminum frame"
	specs["Weight"] = "253 g"
	specs["Dimensions"] = "Unfolded: 154.9 x 129.9 x 6.1 mm / Folded: 154.9 x 67.1 x 13.4 mm"
	specs["Water Resistance"] = "IPX8"
	specs["Network Technology"] = "5G"
	specs["Rear Camera"] = "50 MP + 10 MP + 12 MP"
	specs["Front Camera"] = "4 MP (Under Display) + 10 MP (Cover)"
	specs["Battery"] = "4,400 mAh"
	specs["Operating System"] = "Android 13, upgradable"
	specs["Available Colors"] = "Icy Blue, Phantom Black, Cream, Gray, Blue"
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
