package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileOneplusOpen seeds specifications/options for product 'oneplus-open'
type SpecificationSeederMobileOneplusOpen struct {
	BaseSeeder
}

// NewSpecificationSeederMobileOneplusOpen creates a new seeder instance
func NewSpecificationSeederMobileOneplusOpen() *SpecificationSeederMobileOneplusOpen {
	return &SpecificationSeederMobileOneplusOpen{BaseSeeder: BaseSeeder{name: "Specifications for OnePlus Open"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileOneplusOpen) getBanglaTranslations() map[string]string {
	return map[string]string{
		"120Hz": "১২০Hz",
		"16 GB": "১৬ GB",
		"20 MP (Inner) + 32 MP (Cover)": "২০ MP (Inner) + ৩২ MP (Cover)",
		"2268 x 2440 pixels": "২২৬৮ x ২৪৪০ pixels",
		"239 g": "২৩৯ g",
		"4,805 mAh": "৪,৮০৫ এমএএইচ",
		"48 MP + 64 MP + 48 MP": "৪৮ MP + ৬৪ MP + ৪৮ MP",
		"512 GB": "৫১২ GB",
		"5G": "৫G",
		"7.82 inches (Main) / 6.31 inches (Cover)": "৭.৮২ ইঞ্চি (Main) / ৬.৩১ ইঞ্চি (Cover)",
		"Adreno 740": "Adreno ৭৪০",
		"Android 13, OxygenOS 13.2": "Android ১৩, OxygenOS ১৩.২",
		"Emerald Dusk, Voyager Black": "Emerald Dusk, Voyager কালো",
		"Foldable LTPO3 Flexi-fluid AMOLED, 120Hz, Dolby Vision, 2800 nits": "Foldable LTPO৩ Flexi-fluid AMOLED, ১২০Hz, Dolby Vision, ২৮০০ nits",
		"Glass front (Ceramic Guard), glass back, aluminum frame": "গ্লাস সামনে (Ceramic Guard), গ্লাস পেছনে, অ্যালুমিনিয়াম ফ্রেম",
		"IPX4": "IPX৪",
		"October 2023": "October ২০২৩",
		"Qualcomm SM8550-AB Snapdragon 8 Gen 2 (4 nm)": "Qualcomm SM৮৫৫০-AB Snapdragon ৮ Gen ২ (৪ nm)",
		"Snapdragon 8 Gen 2": "Snapdragon ৮ Gen ২",
		"Unfolded: 153.4 x 143.1 x 5.8 mm / Folded: 153.4 x 73.3 x 11.7 mm": "Unfolded: ১৫৩.৪ x ১৪৩.১ x ৫.৮ মিমি / Folded: ১৫৩.৪ x ৭৩.৩ x ১১.৭ মিমি",
	}
}

// Seed inserts specification records for the product identified by slug 'oneplus-open'
func (s *SpecificationSeederMobileOneplusOpen) Seed(db *gorm.DB) error {
	productSlug := "oneplus-open"

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

	// Override model-specific values for OnePlus Open
	specs["Display Size"] = "7.82 inches (Main) / 6.31 inches (Cover)"
	specs["Processor"] = "Snapdragon 8 Gen 2"
	specs["Chipset"] = "Qualcomm SM8550-AB Snapdragon 8 Gen 2 (4 nm)"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "Adreno 740"
	specs["Ram"] = "16 GB"
	specs["Storage"] = "512 GB"
	specs["Display Type"] = "Foldable LTPO3 Flexi-fluid AMOLED, 120Hz, Dolby Vision, 2800 nits"
	specs["Resolution"] = "2268 x 2440 pixels"
	specs["Screen Protection"] = "Ceramic Guard (Cover)"
	specs["Refresh Rate"] = "120Hz"
	specs["Build Material"] = "Glass front (Ceramic Guard), glass back, aluminum frame"
	specs["Weight"] = "239 g"
	specs["Dimensions"] = "Unfolded: 153.4 x 143.1 x 5.8 mm / Folded: 153.4 x 73.3 x 11.7 mm"
	specs["Water Resistance"] = "IPX4"
	specs["Network Technology"] = "5G"
	specs["Rear Camera"] = "48 MP + 64 MP + 48 MP"
	specs["Front Camera"] = "20 MP (Inner) + 32 MP (Cover)"
	specs["Battery"] = "4,805 mAh"
	specs["Operating System"] = "Android 13, OxygenOS 13.2"
	specs["Available Colors"] = "Emerald Dusk, Voyager Black"
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
