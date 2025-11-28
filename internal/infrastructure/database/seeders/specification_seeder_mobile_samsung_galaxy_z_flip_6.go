package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileSamsungGalaxyZFlip6 seeds specifications/options for product 'samsung-galaxy-z-flip-6'
type SpecificationSeederMobileSamsungGalaxyZFlip6 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyZFlip6 creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyZFlip6() *SpecificationSeederMobileSamsungGalaxyZFlip6 {
	return &SpecificationSeederMobileSamsungGalaxyZFlip6{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy Z Flip 6"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyZFlip6) getBanglaTranslations() map[string]string {
	return map[string]string{
		"10 MP": "১০ মেগাপিক্সেল",
		"1080 x 2640 pixels": "১০৮০ x ২৬৪০ পিক্সেল",
		"12 GB": "১২ GB",
		"120Hz": "১২০Hz",
		"187 g": "১৮৭ g",
		"256 GB / 512 GB": "২৫৬ জিবি / ৫১২ GB",
		"4,000 mAh": "৪,০০০ এমএএইচ",
		"50 MP + 12 MP": "৫০ মেগাপিক্সেল + ১২ মেগাপিক্সেল",
		"5G": "৫G",
		"6.7 inches (Main) / 3.4 inches (Cover)": "৬.৭ ইঞ্চি (Main) / ৩.৪ ইঞ্চি (Cover)",
		"Adreno 750": "অ্যাড্রেনো ৭৫০",
		"Android 14, One UI 6.1.1": "অ্যান্ড্রয়েড ১৪, ওয়ান ইউআই ৬.১.১",
		"Corning Gorilla Glass Victus 2": "কর্নিং গরিলা গ্লাস ভিক্টাস ২",
		"Foldable Dynamic LTPO AMOLED 2X, 120Hz, HDR10+": "Foldable Dynamic এলটিপিও অ্যামোলেড ২X, ১২০Hz, এইচডিআর১০+",
		"IP48": "IP৪৮",
		"July 2024": "জুলাই ২০২৪",
		"Plastic front (open), glass back (Gorilla Glass Victus 2), aluminum frame": "Plastic সামনে (open), গ্লাস পেছনে (গরিলা গ্লাস Victus ২), অ্যালুমিনিয়াম ফ্রেম",
		"Qualcomm SM8650-AC Snapdragon 8 Gen 3 (4 nm)": "কোয়ালকম SM৮৬৫০-AC স্ন্যাপড্রাগন ৮ জেন ৩ (৪ ন্যানোমিটার)",
		"Silver Shadow, Yellow, Blue, Mint, Crafted Black, White, Peach": "রূপালী শ্যাডো, হলুদ, নীল, মিন্ট, ক্রাফটেড কালো, সাদা, পীচ",
		"Snapdragon 8 Gen 3": "স্ন্যাপড্রাগন ৮ জেন ৩",
		"Unfolded: 165.1 x 71.9 x 6.9 mm / Folded: 85.1 x 71.9 x 14.9 mm": "Unfolded: ১৬৫.১ x ৭১.৯ x ৬.৯ মিমি / Folded: ৮৫.১ x ৭১.৯ x ১৪.৯ মিমি",
	}
}

// Seed inserts specification records for the product identified by slug 'samsung-galaxy-z-flip-6'
func (s *SpecificationSeederMobileSamsungGalaxyZFlip6) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-z-flip-6"

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

	// Override model-specific values for Samsung Galaxy Z Flip 6
	specs["Display Size"] = "6.7 inches (Main) / 3.4 inches (Cover)"
	specs["Processor"] = "Snapdragon 8 Gen 3"
	specs["Chipset"] = "Qualcomm SM8650-AC Snapdragon 8 Gen 3 (4 nm)"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "Adreno 750"
	specs["Ram"] = "12 GB"
	specs["Storage"] = "256 GB / 512 GB"
	specs["Display Type"] = "Foldable Dynamic LTPO AMOLED 2X, 120Hz, HDR10+"
	specs["Resolution"] = "1080 x 2640 pixels"
	specs["Screen Protection"] = "Corning Gorilla Glass Victus 2"
	specs["Refresh Rate"] = "120Hz"
	specs["Build Material"] = "Plastic front (open), glass back (Gorilla Glass Victus 2), aluminum frame"
	specs["Weight"] = "187 g"
	specs["Dimensions"] = "Unfolded: 165.1 x 71.9 x 6.9 mm / Folded: 85.1 x 71.9 x 14.9 mm"
	specs["Water Resistance"] = "IP48"
	specs["Network Technology"] = "5G"
	specs["Rear Camera"] = "50 MP + 12 MP"
	specs["Front Camera"] = "10 MP"
	specs["Battery"] = "4,000 mAh"
	specs["Operating System"] = "Android 14, One UI 6.1.1"
	specs["Available Colors"] = "Silver Shadow, Yellow, Blue, Mint, Crafted Black, White, Peach"
	specs["Announcement Date"] = "July 2024"
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
