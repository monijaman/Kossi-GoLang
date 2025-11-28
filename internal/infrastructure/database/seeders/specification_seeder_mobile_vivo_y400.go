package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileVivoY400 seeds specifications/options for product 'vivo-y400'
type SpecificationSeederMobileVivoY400 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileVivoY400 creates a new seeder instance
func NewSpecificationSeederMobileVivoY400() *SpecificationSeederMobileVivoY400 {
	return &SpecificationSeederMobileVivoY400{BaseSeeder: BaseSeeder{name: "Specifications for Vivo Y400"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileVivoY400) getBanglaTranslations() map[string]string {
	return map[string]string{
		"120Hz": "১২০Hz",
		"128 / 256 GB": "১২৮ / ২৫৬ GB",
		"190 g": "১৯০ g",
		"44W wired": "৪৪W তারযুক্ত",
		"50 MP + 2 MP": "৫০ মেগাপিক্সেল + ২ মেগাপিক্সেল",
		"6.68 inches": "৬.৬৮ ইঞ্চি",
		"6000 mAh": "৬০০০ এমএএইচ",
		"720 x 1608 pixels": "৭২০ x ১৬০৮ পিক্সেল",
		"8 GB": "৮ GB",
		"8 MP": "৮ মেগাপিক্সেল",
		"Adreno 613": "অ্যাড্রেনো ৬১৩",
		"Android 14, Funtouch 14": "অ্যান্ড্রয়েড ১৪, Funথেকেuch ১৪",
		"GSM / HSPA / LTE / 5G": "জিএসএম / এইচএসপিএ / এলটিই / ৫G",
		"Glass front, plastic back, plastic frame": "গ্লাস সামনে, প্লাস্টিক পেছনে, প্লাস্টিক ফ্রেম",
		"IP54": "IP৫৪",
		"IPS LCD, 120Hz": "আইপিএস এলসিডি, ১২০Hz",
		"November 2024": "নভেম্বর ২০২৪",
		"Purple, Green": "বেগুনি, সবুজ",
		"Snapdragon 4 Gen 2": "স্ন্যাপড্রাগন ৪ জেন ২",
		"Snapdragon 4 Gen 2 (4 nm)": "স্ন্যাপড্রাগন ৪ জেন ২ (৪ ন্যানোমিটার)",
	}
}

// Seed inserts specification records for the product identified by slug 'vivo-y400'
func (s *SpecificationSeederMobileVivoY400) Seed(db *gorm.DB) error {
	productSlug := "vivo-y400"

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

	// Override model-specific values for Vivo Y400
	specs["Display Size"] = "6.68 inches"
	specs["Processor"] = "Snapdragon 4 Gen 2"
	specs["Chipset"] = "Snapdragon 4 Gen 2 (4 nm)"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "Adreno 613"
	specs["Ram"] = "8 GB"
	specs["Storage"] = "128 / 256 GB"
	specs["Display Type"] = "IPS LCD, 120Hz"
	specs["Resolution"] = "720 x 1608 pixels"
	specs["Refresh Rate"] = "120Hz"
	specs["Build Material"] = "Glass front, plastic back, plastic frame"
	specs["Weight"] = "190 g"
	specs["Water Resistance"] = "IP54"
	specs["Network Technology"] = "GSM / HSPA / LTE / 5G"
	specs["Rear Camera"] = "50 MP + 2 MP"
	specs["Front Camera"] = "8 MP"
	specs["Battery"] = "6000 mAh"
	specs["Fast Charging"] = "44W wired"
	specs["Operating System"] = "Android 14, Funtouch 14"
	specs["Available Colors"] = "Purple, Green"
	specs["Announcement Date"] = "November 2024"
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
