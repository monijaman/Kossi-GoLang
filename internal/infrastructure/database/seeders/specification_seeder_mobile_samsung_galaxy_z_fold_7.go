package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileSamsungGalaxyZFold7 seeds specifications/options for product 'samsung-galaxy-z-fold-7'
type SpecificationSeederMobileSamsungGalaxyZFold7 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyZFold7 creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyZFold7() *SpecificationSeederMobileSamsungGalaxyZFold7 {
	return &SpecificationSeederMobileSamsungGalaxyZFold7{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy Z Fold 7"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyZFold7) getBanglaTranslations() map[string]string {
	return map[string]string{
		"12 GB / 16 GB": "১২ GB / ১৬ GB",
		"120Hz": "১২০Hz",
		"1856 x 2160 pixels (Main)": "১৮৫৬ x ২১৬০ pixels (Main)",
		"235 g": "২৩৫ g",
		"256 GB / 512 GB / 1 TB": "২৫৬ GB / ৫১২ GB / ১ TB",
		"4 MP (Under Display) + 10 MP (Cover)": "৪ MP (ডিসপ্লের নিচে) + ১০ MP (Cover)",
		"4,600 mAh": "৪,৬০০ এমএএইচ",
		"50 MP + 12 MP + 10 MP": "৫০ MP + ১২ MP + ১০ MP",
		"5G": "৫G",
		"7.6 inches (Main) / 6.3 inches (Cover)": "৭.৬ ইঞ্চি (Main) / ৬.৩ ইঞ্চি (Cover)",
		"Adreno 830": "Adreno ৮৩০",
		"Android 15, One UI 7.1.1": "Android ১৫, One UI ৭.১.১",
		"Corning Gorilla Armor 2": "Corning Gorilla Armor ২",
		"Foldable Dynamic LTPO AMOLED 2X, 120Hz, HDR10+": "Foldable Dynamic LTPO AMOLED ২X, ১২০Hz, HDR১০+",
		"Glass front (Gorilla Armor 2), glass back (Victus 2), titanium frame": "গ্লাস সামনে (Gorilla Armor ২), গ্লাস পেছনে (Victus ২), টাইটানিয়াম ফ্রেম",
		"IPX8": "IPX৮",
		"July 2025": "July ২০২৫",
		"Phantom Black, Icy Blue, Cream": "Phantom কালো, Icy নীল, Cream",
		"Qualcomm SM8750-AC Snapdragon 8 Elite (3 nm)": "Qualcomm SM৮৭৫০-AC Snapdragon ৮ Elite (৩ nm)",
		"Snapdragon 8 Elite": "Snapdragon ৮ Elite",
		"Unfolded: 153.5 x 132.6 x 5.6 mm / Folded: 153.5 x 68.1 x 11.2 mm": "Unfolded: ১৫৩.৫ x ১৩২.৬ x ৫.৬ মিমি / Folded: ১৫৩.৫ x ৬৮.১ x ১১.২ মিমি",
	}
}

// Seed inserts specification records for the product identified by slug 'samsung-galaxy-z-fold-7'
func (s *SpecificationSeederMobileSamsungGalaxyZFold7) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-z-fold-7"

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

	// Override model-specific values for Samsung Galaxy Z Fold 7
	specs["Display Size"] = "7.6 inches (Main) / 6.3 inches (Cover)"
	specs["Processor"] = "Snapdragon 8 Elite"
	specs["Chipset"] = "Qualcomm SM8750-AC Snapdragon 8 Elite (3 nm)"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "Adreno 830"
	specs["Ram"] = "12 GB / 16 GB"
	specs["Storage"] = "256 GB / 512 GB / 1 TB"
	specs["Display Type"] = "Foldable Dynamic LTPO AMOLED 2X, 120Hz, HDR10+"
	specs["Resolution"] = "1856 x 2160 pixels (Main)"
	specs["Screen Protection"] = "Corning Gorilla Armor 2"
	specs["Refresh Rate"] = "120Hz"
	specs["Build Material"] = "Glass front (Gorilla Armor 2), glass back (Victus 2), titanium frame"
	specs["Weight"] = "235 g"
	specs["Dimensions"] = "Unfolded: 153.5 x 132.6 x 5.6 mm / Folded: 153.5 x 68.1 x 11.2 mm"
	specs["Water Resistance"] = "IPX8"
	specs["Network Technology"] = "5G"
	specs["Rear Camera"] = "50 MP + 12 MP + 10 MP"
	specs["Front Camera"] = "4 MP (Under Display) + 10 MP (Cover)"
	specs["Battery"] = "4,600 mAh"
	specs["Operating System"] = "Android 15, One UI 7.1.1"
	specs["Available Colors"] = "Phantom Black, Icy Blue, Cream"
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
