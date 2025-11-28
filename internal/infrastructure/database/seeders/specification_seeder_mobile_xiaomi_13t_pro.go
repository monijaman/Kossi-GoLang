package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileXiaomi13tPro seeds specifications/options for product 'xiaomi-13t-pro'
type SpecificationSeederMobileXiaomi13tPro struct {
	BaseSeeder
}

// NewSpecificationSeederMobileXiaomi13tPro creates a new seeder instance
func NewSpecificationSeederMobileXiaomi13tPro() *SpecificationSeederMobileXiaomi13tPro {
	return &SpecificationSeederMobileXiaomi13tPro{BaseSeeder: BaseSeeder{name: "Specifications for Xiaomi 13T Pro"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileXiaomi13tPro) getBanglaTranslations() map[string]string {
	return map[string]string{
		"12 GB / 16 GB": "১২ জিবি / ১৬ GB",
		"120W wired": "১২০W তারযুক্ত",
		"1220 x 2712 pixels": "১২২০ x ২৭১২ পিক্সেল",
		"144Hz": "১৪৪Hz",
		"20 MP": "২০ মেগাপিক্সেল",
		"200 g": "২০০ g",
		"256 GB / 512 GB / 1 TB": "২৫৬ জিবি / ৫১২ জিবি / ১ টিবি",
		"50 MP + 50 MP + 12 MP": "৫০ মেগাপিক্সেল + ৫০ মেগাপিক্সেল + ১২ মেগাপিক্সেল",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6.67 inches": "৬.৬৭ ইঞ্চি",
		"AMOLED, 144Hz, Dolby Vision, HDR10+": "অ্যামোলেড, ১৪৪Hz, ডলবি ভিশন, এইচডিআর১০+",
		"Alpine Blue, Meadow Green, Black": "আলপাইন নীল, মিডো সবুজ, কালো",
		"Android 13, upgradable to Android 14, HyperOS": "অ্যান্ড্রয়েড ১৩, আপগ্রেডযোগ্য থেকে অ্যান্ড্রয়েড ১৪, হাইপার ওএস",
		"Dimensity 9200+": "ডাইমেনসিটি ৯২০০+",
		"GSM / HSPA / LTE / 5G": "জিএসএম / এইচএসপিএ / এলটিই / ৫G",
		"Glass front (Gorilla Glass 5), glass back or silicone polymer back, plastic frame": "গ্লাস সামনে (গরিলা গ্লাস ৫), গ্লাস পেছনে বা সিলিকন পলিমার পেছনে, প্লাস্টিক ফ্রেম",
		"IP68": "IP৬৮",
		"Immortalis-G715 MC11": "Immবাtalis-G৭১৫ MC১১",
		"Mediatek Dimensity 9200+ (4 nm)": "মিডিয়াটেক ডাইমেনসিটি ৯২০০+ (৪ ন্যানোমিটার)",
		"September 2023": "সেপ্টেম্বর ২০২৩",
	}
}

// Seed inserts specification records for the product identified by slug 'xiaomi-13t-pro'
func (s *SpecificationSeederMobileXiaomi13tPro) Seed(db *gorm.DB) error {
	productSlug := "xiaomi-13t-pro"

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

	// Override model-specific values for Xiaomi 13T Pro
	specs["Display Size"] = "6.67 inches"
	specs["Processor"] = "Dimensity 9200+"
	specs["Chipset"] = "Mediatek Dimensity 9200+ (4 nm)"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "Immortalis-G715 MC11"
	specs["Ram"] = "12 GB / 16 GB"
	specs["Storage"] = "256 GB / 512 GB / 1 TB"
	specs["Display Type"] = "AMOLED, 144Hz, Dolby Vision, HDR10+"
	specs["Resolution"] = "1220 x 2712 pixels"
	specs["Refresh Rate"] = "144Hz"
	specs["Build Material"] = "Glass front (Gorilla Glass 5), glass back or silicone polymer back, plastic frame"
	specs["Weight"] = "200 g"
	specs["Water Resistance"] = "IP68"
	specs["Network Technology"] = "GSM / HSPA / LTE / 5G"
	specs["Rear Camera"] = "50 MP + 50 MP + 12 MP"
	specs["Front Camera"] = "20 MP"
	specs["Battery"] = "5000 mAh"
	specs["Fast Charging"] = "120W wired"
	specs["Operating System"] = "Android 13, upgradable to Android 14, HyperOS"
	specs["Available Colors"] = "Alpine Blue, Meadow Green, Black"
	specs["Announcement Date"] = "September 2023"
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
