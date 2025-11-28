package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileSamsungGalaxyA365g seeds specifications/options for product 'samsung-galaxy-a36-5g'
type SpecificationSeederMobileSamsungGalaxyA365g struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyA365g creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyA365g() *SpecificationSeederMobileSamsungGalaxyA365g {
	return &SpecificationSeederMobileSamsungGalaxyA365g{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy A36 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyA365g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080 x 2400 pixels (~399 ppi density)": "১০৮০ x ২৪০০ পিক্সেল (~৩৯৯ পিপিআই density)",
		"120Hz": "১২০Hz",
		"128 GB / 256 GB": "১২৮ জিবি / ২৫৬ GB",
		"162.5 x 75.8 x 8.1 mm": "১৬২.৫ x ৭৫.৮ x ৮.১ মিমি",
		"195 g (6.88 oz)": "১৯৫ গ্রাম (৬.৮৮ oz)",
		"20 MP": "২০ মেগাপিক্সেল",
		"4500 mAh": "৪৫০০ এমএএইচ",
		"48 MP + 8 MP + 5 MP": "৪৮ মেগাপিক্সেল + ৮ মেগাপিক্সেল + ৫ মেগাপিক্সেল",
		"6 GB / 8 GB": "৬ জিবি / ৮ GB",
		"6.6 inches": "৬.৬ ইঞ্চি",
		"Adreno 644": "অ্যাড্রেনো ৬৪৪",
		"Android 14, One UI 6": "অ্যান্ড্রয়েড ১৪, ওয়ান ইউআই ৬",
		"Black, White, Green": "কালো, সাদা, সবুজ",
		"Corning Gorilla Glass 5": "কর্নিং গরিলা গ্লাস ৫",
		"GSM / HSPA / LTE / 5G": "জিএসএম / এইচএসপিএ / এলটিই / ৫G",
		"Glass front, plastic frame, plastic back": "গ্লাস সামনে, প্লাস্টিক ফ্রেম, প্লাস্টিক পেছনে",
		"IP67 dust/water resistant (up to 1m for 30 min)": "IP৬৭ dust/জল প্রতিরোধী (পর্যন্ত ১m জন্য ৩০ min)",
		"Octa-core (1x2.4 GHz Cortex-A710 & 3x2.36 GHz Cortex-A710 & 4x1.8 GHz Cortex-A510)": "অক্টা-কোর (১x২.৪ গিগাহার্টজ Cবাtex-A৭১০ & ৩x২.৩৬ গিগাহার্টজ Cবাtex-A৭১০ & ৪x১.৮ গিগাহার্টজ Cবাtex-A৫১০)",
		"Qualcomm SM7450-AB Snapdragon 7 Gen 1 (4 nm)": "কোয়ালকম SM৭৪৫০-AB স্ন্যাপড্রাগন ৭ জেন ১ (৪ ন্যানোমিটার)",
		"September 2025": "সেপ্টেম্বর ২০২৫",
		"Snapdragon 7 Gen 1": "স্ন্যাপড্রাগন ৭ জেন ১",
		"Super AMOLED, 120Hz": "Super অ্যামোলেড, ১২০Hz",
	}
}

// Seed inserts specification records for the product identified by slug 'samsung-galaxy-a36-5g'
func (s *SpecificationSeederMobileSamsungGalaxyA365g) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-a36-5g"

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

	// Override model-specific values for Samsung Galaxy A36 5G
	specs["Display Size"] = "6.6 inches"
	specs["Processor"] = "Snapdragon 7 Gen 1"
	specs["Chipset"] = "Qualcomm SM7450-AB Snapdragon 7 Gen 1 (4 nm)"
	specs["Cpu Type"] = "Octa-core (1x2.4 GHz Cortex-A710 & 3x2.36 GHz Cortex-A710 & 4x1.8 GHz Cortex-A510)"
	specs["Gpu Type"] = "Adreno 644"
	specs["Ram"] = "6 GB / 8 GB"
	specs["Storage"] = "128 GB / 256 GB"
	specs["Display Type"] = "Super AMOLED, 120Hz"
	specs["Resolution"] = "1080 x 2400 pixels (~399 ppi density)"
	specs["Screen Protection"] = "Corning Gorilla Glass 5"
	specs["Refresh Rate"] = "120Hz"
	specs["Build Material"] = "Glass front, plastic frame, plastic back"
	specs["Weight"] = "195 g (6.88 oz)"
	specs["Dimensions"] = "162.5 x 75.8 x 8.1 mm"
	specs["Water Resistance"] = "IP67 dust/water resistant (up to 1m for 30 min)"
	specs["Network Technology"] = "GSM / HSPA / LTE / 5G"
	specs["Rear Camera"] = "48 MP + 8 MP + 5 MP"
	specs["Front Camera"] = "20 MP"
	specs["Battery"] = "4500 mAh"
	specs["Operating System"] = "Android 14, One UI 6"
	specs["Available Colors"] = "Black, White, Green"
	specs["Announcement Date"] = "September 2025"
	specs["Device Status"] = "Upcoming"

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
