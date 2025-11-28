package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileRealme13Pro seeds specifications/options for product 'realme-13-pro'
type SpecificationSeederMobileRealme13Pro struct {
	BaseSeeder
}

// NewSpecificationSeederMobileRealme13Pro creates a new seeder instance
func NewSpecificationSeederMobileRealme13Pro() *SpecificationSeederMobileRealme13Pro {
	return &SpecificationSeederMobileRealme13Pro{BaseSeeder: BaseSeeder{name: "Specifications for Realme 13 Pro"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileRealme13Pro) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080 x 2412 pixels": "১০৮০ x ২৪১২ পিক্সেল",
		"120Hz": "১২০Hz",
		"128 GB / 256 GB / 512 GB": "১২৮ জিবি / ২৫৬ জিবি / ৫১২ GB",
		"183.5 g": "১৮৩.৫ g",
		"32 MP": "৩২ মেগাপিক্সেল",
		"45W wired": "৪৫W তারযুক্ত",
		"50 MP + 8 MP + 2 MP": "৫০ মেগাপিক্সেল + ৮ মেগাপিক্সেল + ২ মেগাপিক্সেল",
		"5200 mAh": "৫২০০ এমএএইচ",
		"6.7 inches": "৬.৭ ইঞ্চি",
		"8 GB / 12 GB": "৮ জিবি / ১২ GB",
		"AMOLED, 120Hz, 1B colors": "অ্যামোলেড, ১২০Hz, ১B রঙ",
		"Adreno 710": "অ্যাড্রেনো ৭১০",
		"Android 14, Realme UI 5.0": "অ্যান্ড্রয়েড ১৪, Realme UI ৫.০",
		"GSM / HSPA / LTE / 5G": "জিএসএম / এইচএসপিএ / এলটিই / ৫G",
		"Glass front, glass back, plastic frame": "গ্লাস সামনে, গ্লাস পেছনে, প্লাস্টিক ফ্রেম",
		"IP65": "IP৬৫",
		"July 2024": "জুলাই ২০২৪",
		"Monet Gold, Monet Purple, Emerald Green": "Monet সোনালী, Monet বেগুনি, Emerald সবুজ",
		"Qualcomm Snapdragon 7s Gen 2 (4 nm)": "কোয়ালকম স্ন্যাপড্রাগন ৭s জেন ২ (৪ ন্যানোমিটার)",
		"Snapdragon 7s Gen 2": "স্ন্যাপড্রাগন ৭s জেন ২",
	}
}

// Seed inserts specification records for the product identified by slug 'realme-13-pro'
func (s *SpecificationSeederMobileRealme13Pro) Seed(db *gorm.DB) error {
	productSlug := "realme-13-pro"

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

	// Override model-specific values for Realme 13 Pro
	specs["Display Size"] = "6.7 inches"
	specs["Processor"] = "Snapdragon 7s Gen 2"
	specs["Chipset"] = "Qualcomm Snapdragon 7s Gen 2 (4 nm)"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "Adreno 710"
	specs["Ram"] = "8 GB / 12 GB"
	specs["Storage"] = "128 GB / 256 GB / 512 GB"
	specs["Display Type"] = "AMOLED, 120Hz, 1B colors"
	specs["Resolution"] = "1080 x 2412 pixels"
	specs["Refresh Rate"] = "120Hz"
	specs["Build Material"] = "Glass front, glass back, plastic frame"
	specs["Weight"] = "183.5 g"
	specs["Water Resistance"] = "IP65"
	specs["Network Technology"] = "GSM / HSPA / LTE / 5G"
	specs["Rear Camera"] = "50 MP + 8 MP + 2 MP"
	specs["Front Camera"] = "32 MP"
	specs["Battery"] = "5200 mAh"
	specs["Fast Charging"] = "45W wired"
	specs["Operating System"] = "Android 14, Realme UI 5.0"
	specs["Available Colors"] = "Monet Gold, Monet Purple, Emerald Green"
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
