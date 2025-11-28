package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileOneplusNordCe5 seeds specifications/options for product 'oneplus-nord-ce5'
type SpecificationSeederMobileOneplusNordCe5 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileOneplusNordCe5 creates a new seeder instance
func NewSpecificationSeederMobileOneplusNordCe5() *SpecificationSeederMobileOneplusNordCe5 {
	return &SpecificationSeederMobileOneplusNordCe5{BaseSeeder: BaseSeeder{name: "Specifications for OnePlus Nord CE5"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileOneplusNordCe5) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080 x 2412 pixels": "১০৮০ x ২৪১২ পিক্সেল",
		"120Hz": "১২০Hz",
		"128 GB / 256 GB": "১২৮ জিবি / ২৫৬ GB",
		"16 MP": "১৬ মেগাপিক্সেল",
		"162.7 x 75.5 x 8.2 mm": "১৬২.৭ x ৭৫.৫ x ৮.২ মিমি",
		"186 g": "১৮৬ g",
		"5,500 mAh": "৫,৫০০ এমএএইচ",
		"50 MP + 8 MP": "৫০ মেগাপিক্সেল + ৮ মেগাপিক্সেল",
		"5G": "৫G",
		"6.7 inches": "৬.৭ ইঞ্চি",
		"8 GB / 12 GB": "৮ জিবি / ১২ GB",
		"AMOLED, 120Hz, HDR10+": "অ্যামোলেড, ১২০Hz, এইচডিআর১০+",
		"Adreno 720": "অ্যাড্রেনো ৭২০",
		"Android 15, OxygenOS 15": "অ্যান্ড্রয়েড ১৫, OxygenOS ১৫",
		"April 2025": "এপ্রিল ২০২৫",
		"Glass front, plastic frame, plastic back": "গ্লাস সামনে, প্লাস্টিক ফ্রেম, প্লাস্টিক পেছনে",
		"IP54": "IP৫৪",
		"Qualcomm Snapdragon 7 Gen 3 (4 nm)": "কোয়ালকম স্ন্যাপড্রাগন ৭ জেন ৩ (৪ ন্যানোমিটার)",
		"Snapdragon 7 Gen 3": "স্ন্যাপড্রাগন ৭ জেন ৩",
	}
}

// Seed inserts specification records for the product identified by slug 'oneplus-nord-ce5'
func (s *SpecificationSeederMobileOneplusNordCe5) Seed(db *gorm.DB) error {
	productSlug := "oneplus-nord-ce5"

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

	// Override model-specific values for OnePlus Nord CE5
	specs["Display Size"] = "6.7 inches"
	specs["Processor"] = "Snapdragon 7 Gen 3"
	specs["Chipset"] = "Qualcomm Snapdragon 7 Gen 3 (4 nm)"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "Adreno 720"
	specs["Ram"] = "8 GB / 12 GB"
	specs["Storage"] = "128 GB / 256 GB"
	specs["Display Type"] = "AMOLED, 120Hz, HDR10+"
	specs["Resolution"] = "1080 x 2412 pixels"
	specs["Screen Protection"] = "Corning Gorilla Glass"
	specs["Refresh Rate"] = "120Hz"
	specs["Build Material"] = "Glass front, plastic frame, plastic back"
	specs["Weight"] = "186 g"
	specs["Dimensions"] = "162.7 x 75.5 x 8.2 mm"
	specs["Water Resistance"] = "IP54"
	specs["Network Technology"] = "5G"
	specs["Rear Camera"] = "50 MP + 8 MP"
	specs["Front Camera"] = "16 MP"
	specs["Battery"] = "5,500 mAh"
	specs["Operating System"] = "Android 15, OxygenOS 15"
	specs["Available Colors"] = "Celadon Marble, Dark Chrome"
	specs["Announcement Date"] = "April 2025"
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
