package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileSamsungGalaxyZFold6 seeds specifications/options for product 'samsung-galaxy-z-fold-6'
type SpecificationSeederMobileSamsungGalaxyZFold6 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyZFold6 creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyZFold6() *SpecificationSeederMobileSamsungGalaxyZFold6 {
	return &SpecificationSeederMobileSamsungGalaxyZFold6{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy Z Fold 6"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyZFold6) getBanglaTranslations() map[string]string {
	return map[string]string{
		"12 GB": "১২ GB",
		"120Hz": "১২০Hz",
		"1856 x 2160 pixels": "১৮৫৬ x ২১৬০ পিক্সেল",
		"239 g": "২৩৯ g",
		"256 GB / 512 GB / 1 TB": "২৫৬ জিবি / ৫১২ জিবি / ১ টিবি",
		"4 MP (Under Display) + 10 MP (Cover)": "৪ মেগাপিক্সেল (ডিসপ্লের নিচে) + ১০ মেগাপিক্সেল (Cover)",
		"4,400 mAh": "৪,৪০০ এমএএইচ",
		"50 MP + 10 MP + 12 MP": "৫০ মেগাপিক্সেল + ১০ মেগাপিক্সেল + ১২ মেগাপিক্সেল",
		"5G": "৫G",
		"7.6 inches (Main) / 6.3 inches (Cover)": "৭.৬ ইঞ্চি (Main) / ৬.৩ ইঞ্চি (Cover)",
		"Adreno 750": "অ্যাড্রেনো ৭৫০",
		"Android 14, One UI 6.1.1": "অ্যান্ড্রয়েড ১৪, ওয়ান ইউআই ৬.১.১",
		"Corning Gorilla Glass Victus 2": "কর্নিং গরিলা গ্লাস ভিক্টাস ২",
		"Foldable Dynamic LTPO AMOLED 2X, 120Hz, HDR10+": "Foldable Dynamic এলটিপিও অ্যামোলেড ২X, ১২০Hz, এইচডিআর১০+",
		"Glass front (Gorilla Glass Victus 2), glass back (Victus 2), aluminum frame": "গ্লাস সামনে (গরিলা গ্লাস Victus ২), গ্লাস পেছনে (Victus ২), অ্যালুমিনিয়াম ফ্রেম",
		"IP48": "IP৪৮",
		"July 2024": "জুলাই ২০২৪",
		"Qualcomm SM8650-AC Snapdragon 8 Gen 3 (4 nm)": "কোয়ালকম SM৮৬৫০-AC স্ন্যাপড্রাগন ৮ জেন ৩ (৪ ন্যানোমিটার)",
		"Silver Shadow, Pink, Navy, Crafted Black, White": "রূপালী শ্যাডো, গোলাপী, নেভি, ক্রাফটেড কালো, সাদা",
		"Snapdragon 8 Gen 3": "স্ন্যাপড্রাগন ৮ জেন ৩",
		"Unfolded: 153.5 x 132.6 x 5.6 mm / Folded: 153.5 x 68.1 x 12.1 mm": "Unfolded: ১৫৩.৫ x ১৩২.৬ x ৫.৬ মিমি / Folded: ১৫৩.৫ x ৬৮.১ x ১২.১ মিমি",
	}
}

// Seed inserts specification records for the product identified by slug 'samsung-galaxy-z-fold-6'
func (s *SpecificationSeederMobileSamsungGalaxyZFold6) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-z-fold-6"

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

	// Override model-specific values for Samsung Galaxy Z Fold 6
	specs["Display Size"] = "7.6 inches (Main) / 6.3 inches (Cover)"
	specs["Processor"] = "Snapdragon 8 Gen 3"
	specs["Chipset"] = "Qualcomm SM8650-AC Snapdragon 8 Gen 3 (4 nm)"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "Adreno 750"
	specs["Ram"] = "12 GB"
	specs["Storage"] = "256 GB / 512 GB / 1 TB"
	specs["Display Type"] = "Foldable Dynamic LTPO AMOLED 2X, 120Hz, HDR10+"
	specs["Resolution"] = "1856 x 2160 pixels"
	specs["Screen Protection"] = "Corning Gorilla Glass Victus 2"
	specs["Refresh Rate"] = "120Hz"
	specs["Build Material"] = "Glass front (Gorilla Glass Victus 2), glass back (Victus 2), aluminum frame"
	specs["Weight"] = "239 g"
	specs["Dimensions"] = "Unfolded: 153.5 x 132.6 x 5.6 mm / Folded: 153.5 x 68.1 x 12.1 mm"
	specs["Water Resistance"] = "IP48"
	specs["Network Technology"] = "5G"
	specs["Rear Camera"] = "50 MP + 10 MP + 12 MP"
	specs["Front Camera"] = "4 MP (Under Display) + 10 MP (Cover)"
	specs["Battery"] = "4,400 mAh"
	specs["Operating System"] = "Android 14, One UI 6.1.1"
	specs["Available Colors"] = "Silver Shadow, Pink, Navy, Crafted Black, White"
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
