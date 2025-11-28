package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileXiaomi15sPro seeds specifications/options for product 'xiaomi-15s-pro'
type SpecificationSeederMobileXiaomi15sPro struct {
	BaseSeeder
}

// NewSpecificationSeederMobileXiaomi15sPro creates a new seeder instance
func NewSpecificationSeederMobileXiaomi15sPro() *SpecificationSeederMobileXiaomi15sPro {
	return &SpecificationSeederMobileXiaomi15sPro{BaseSeeder: BaseSeeder{name: "Specifications for Xiaomi 15S Pro"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileXiaomi15sPro) getBanglaTranslations() map[string]string {
	return map[string]string{
		"12 GB / 16 GB": "১২ জিবি / ১৬ GB",
		"120Hz": "১২০Hz",
		"120W wired, 50W wireless": "১২০W তারযুক্ত, ৫০W বেতার",
		"1440 x 3200 pixels": "১৪৪০ x ৩২০০ পিক্সেল",
		"215 g": "২১৫ g",
		"256 GB / 512 GB / 1 TB": "২৫৬ জিবি / ৫১২ জিবি / ১ টিবি",
		"32 MP": "৩২ মেগাপিক্সেল",
		"50 MP + 50 MP + 50 MP": "৫০ মেগাপিক্সেল + ৫০ মেগাপিক্সেল + ৫০ মেগাপিক্সেল",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6.73 inches": "৬.৭৩ ইঞ্চি",
		"Adreno 830": "অ্যাড্রেনো ৮৩০",
		"Android 15, HyperOS": "অ্যান্ড্রয়েড ১৫, হাইপার ওএস",
		"Black, White, Green": "কালো, সাদা, সবুজ",
		"GSM / HSPA / LTE / 5G": "জিএসএম / এইচএসপিএ / এলটিই / ৫G",
		"Glass front, glass back or eco leather back, aluminum frame": "গ্লাস সামনে, গ্লাস পেছনে বা eco চামড়া পেছনে, অ্যালুমিনিয়াম ফ্রেম",
		"IP68": "IP৬৮",
		"LTPO AMOLED, 120Hz, Dolby Vision, HDR10+": "এলটিপিও অ্যামোলেড, ১২০Hz, ডলবি ভিশন, এইচডিআর১০+",
		"May 2025": "মে ২০২৫",
		"Qualcomm Snapdragon 8 Gen 4": "কোয়ালকম স্ন্যাপড্রাগন ৮ জেন ৪",
		"Snapdragon 8 Gen 4": "স্ন্যাপড্রাগন ৮ জেন ৪",
	}
}

// Seed inserts specification records for the product identified by slug 'xiaomi-15s-pro'
func (s *SpecificationSeederMobileXiaomi15sPro) Seed(db *gorm.DB) error {
	productSlug := "xiaomi-15s-pro"

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

	// Override model-specific values for Xiaomi 15S Pro
	specs["Display Size"] = "6.73 inches"
	specs["Processor"] = "Snapdragon 8 Gen 4"
	specs["Chipset"] = "Qualcomm Snapdragon 8 Gen 4"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "Adreno 830"
	specs["Ram"] = "12 GB / 16 GB"
	specs["Storage"] = "256 GB / 512 GB / 1 TB"
	specs["Display Type"] = "LTPO AMOLED, 120Hz, Dolby Vision, HDR10+"
	specs["Resolution"] = "1440 x 3200 pixels"
	specs["Refresh Rate"] = "120Hz"
	specs["Build Material"] = "Glass front, glass back or eco leather back, aluminum frame"
	specs["Weight"] = "215 g"
	specs["Water Resistance"] = "IP68"
	specs["Network Technology"] = "GSM / HSPA / LTE / 5G"
	specs["Rear Camera"] = "50 MP + 50 MP + 50 MP"
	specs["Front Camera"] = "32 MP"
	specs["Battery"] = "5000 mAh"
	specs["Fast Charging"] = "120W wired, 50W wireless"
	specs["Operating System"] = "Android 15, HyperOS"
	specs["Available Colors"] = "Black, White, Green"
	specs["Announcement Date"] = "May 2025"
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
