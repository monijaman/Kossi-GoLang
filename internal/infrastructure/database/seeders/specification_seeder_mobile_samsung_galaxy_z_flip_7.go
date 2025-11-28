package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileSamsungGalaxyZFlip7 seeds specifications/options for product 'samsung-galaxy-z-flip-7'
type SpecificationSeederMobileSamsungGalaxyZFlip7 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyZFlip7 creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyZFlip7() *SpecificationSeederMobileSamsungGalaxyZFlip7 {
	return &SpecificationSeederMobileSamsungGalaxyZFlip7{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy Z Flip 7"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyZFlip7) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080 x 2640 pixels (Main)": "১০৮০ x ২৬৪০ পিক্সেল (Main)",
		"12 MP": "১২ মেগাপিক্সেল",
		"120Hz": "১২০Hz",
		"187 g": "১৮৭ g",
		"256 GB / 512 GB": "২৫৬ জিবি / ৫১২ GB",
		"4,200 mAh": "৪,২০০ এমএএইচ",
		"50 MP + 12 MP": "৫০ মেগাপিক্সেল + ১২ মেগাপিক্সেল",
		"5G": "৫G",
		"6.7 inches (Main) / 3.9 inches (Cover)": "৬.৭ ইঞ্চি (Main) / ৩.৯ ইঞ্চি (Cover)",
		"8 GB / 12 GB": "৮ জিবি / ১২ GB",
		"Adreno 830": "অ্যাড্রেনো ৮৩০",
		"Android 15, One UI 7.1.1": "অ্যান্ড্রয়েড ১৫, ওয়ান ইউআই ৭.১.১",
		"Corning Gorilla Armor 2": "Cবাning Gবাilla Armবা ২",
		"Foldable Dynamic LTPO AMOLED 2X, 120Hz, HDR10+": "Foldable Dynamic এলটিপিও অ্যামোলেড ২X, ১২০Hz, এইচডিআর১০+",
		"Glass front/back, aluminum frame": "গ্লাস সামনে/পেছনে, অ্যালুমিনিয়াম ফ্রেম",
		"IPX8": "IPX৮",
		"July 2025": "জুলাই ২০২৫",
		"Qualcomm SM8750-AC Snapdragon 8 Elite (3 nm)": "কোয়ালকম SM৮৭৫০-AC স্ন্যাপড্রাগন ৮ এলিট (৩ ন্যানোমিটার)",
		"Snapdragon 8 Elite": "স্ন্যাপড্রাগন ৮ এলিট",
		"Unfolded: 165.1 x 71.9 x 6.9 mm / Folded: 85.1 x 71.9 x 14.9 mm": "Unfolded: ১৬৫.১ x ৭১.৯ x ৬.৯ মিমি / Folded: ৮৫.১ x ৭১.৯ x ১৪.৯ মিমি",
	}
}

// Seed inserts specification records for the product identified by slug 'samsung-galaxy-z-flip-7'
func (s *SpecificationSeederMobileSamsungGalaxyZFlip7) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-z-flip-7"

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

	// Override model-specific values for Samsung Galaxy Z Flip 7
	specs["Display Size"] = "6.7 inches (Main) / 3.9 inches (Cover)"
	specs["Processor"] = "Snapdragon 8 Elite"
	specs["Chipset"] = "Qualcomm SM8750-AC Snapdragon 8 Elite (3 nm)"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "Adreno 830"
	specs["Ram"] = "8 GB / 12 GB"
	specs["Storage"] = "256 GB / 512 GB"
	specs["Display Type"] = "Foldable Dynamic LTPO AMOLED 2X, 120Hz, HDR10+"
	specs["Resolution"] = "1080 x 2640 pixels (Main)"
	specs["Screen Protection"] = "Corning Gorilla Armor 2"
	specs["Refresh Rate"] = "120Hz"
	specs["Build Material"] = "Glass front/back, aluminum frame"
	specs["Weight"] = "187 g"
	specs["Dimensions"] = "Unfolded: 165.1 x 71.9 x 6.9 mm / Folded: 85.1 x 71.9 x 14.9 mm"
	specs["Water Resistance"] = "IPX8"
	specs["Network Technology"] = "5G"
	specs["Rear Camera"] = "50 MP + 12 MP"
	specs["Front Camera"] = "12 MP"
	specs["Battery"] = "4,200 mAh"
	specs["Operating System"] = "Android 15, One UI 7.1.1"
	specs["Available Colors"] = "Mint, Graphite, Cream, Lavender"
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
