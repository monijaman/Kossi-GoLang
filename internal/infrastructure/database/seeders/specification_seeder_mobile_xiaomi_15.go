package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileXiaomi15 seeds specifications/options for product 'xiaomi-15'
type SpecificationSeederMobileXiaomi15 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileXiaomi15 creates a new seeder instance
func NewSpecificationSeederMobileXiaomi15() *SpecificationSeederMobileXiaomi15 {
	return &SpecificationSeederMobileXiaomi15{BaseSeeder: BaseSeeder{name: "Specifications for Xiaomi 15"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileXiaomi15) getBanglaTranslations() map[string]string {
	return map[string]string{
		"12 GB / 16 GB": "১২ জিবি / ১৬ GB",
		"1200 x 2670 pixels": "১২০০ x ২৬৭০ পিক্সেল",
		"120Hz": "১২০Hz",
		"152.8 x 71.5 x 8.2 mm": "১৫২.৮ x ৭১.৫ x ৮.২ মিমি",
		"191 g": "১৯১ g",
		"256 GB / 512 GB / 1 TB": "২৫৬ জিবি / ৫১২ জিবি / ১ টিবি",
		"32 MP": "৩২ মেগাপিক্সেল",
		"5,400 mAh": "৫,৪০০ এমএএইচ",
		"50 MP + 50 MP + 50 MP": "৫০ মেগাপিক্সেল + ৫০ মেগাপিক্সেল + ৫০ মেগাপিক্সেল",
		"5G": "৫G",
		"6.36 inches": "৬.৩৬ ইঞ্চি",
		"Adreno 830": "অ্যাড্রেনো ৮৩০",
		"Android 15, HyperOS": "অ্যান্ড্রয়েড ১৫, হাইপার ওএস",
		"Black, White, Green, Purple, Silver": "কালো, সাদা, সবুজ, বেগুনি, রূপালী",
		"Glass front/back, aluminum frame": "গ্লাস সামনে/পেছনে, অ্যালুমিনিয়াম ফ্রেম",
		"IP68": "IP৬৮",
		"LTPO OLED, 120Hz, Dolby Vision, HDR10+, 3200 nits": "এলটিপিও ওএলইডি, ১২০Hz, ডলবি ভিশন, এইচডিআর১০+, ৩২০০ nits",
		"October 2024": "অক্টোবর ২০২৪",
		"Qualcomm SM8750 Snapdragon 8 Elite (3 nm)": "কোয়ালকম SM৮৭৫০ স্ন্যাপড্রাগন ৮ এলিট (৩ ন্যানোমিটার)",
		"Snapdragon 8 Elite": "স্ন্যাপড্রাগন ৮ এলিট",
	}
}

// Seed inserts specification records for the product identified by slug 'xiaomi-15'
func (s *SpecificationSeederMobileXiaomi15) Seed(db *gorm.DB) error {
	productSlug := "xiaomi-15"

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

	// Override model-specific values for Xiaomi 15
	specs["Display Size"] = "6.36 inches"
	specs["Processor"] = "Snapdragon 8 Elite"
	specs["Chipset"] = "Qualcomm SM8750 Snapdragon 8 Elite (3 nm)"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "Adreno 830"
	specs["Ram"] = "12 GB / 16 GB"
	specs["Storage"] = "256 GB / 512 GB / 1 TB"
	specs["Display Type"] = "LTPO OLED, 120Hz, Dolby Vision, HDR10+, 3200 nits"
	specs["Resolution"] = "1200 x 2670 pixels"
	specs["Screen Protection"] = "Xiaomi Dragon Crystal Glass"
	specs["Refresh Rate"] = "120Hz"
	specs["Build Material"] = "Glass front/back, aluminum frame"
	specs["Weight"] = "191 g"
	specs["Dimensions"] = "152.8 x 71.5 x 8.2 mm"
	specs["Water Resistance"] = "IP68"
	specs["Network Technology"] = "5G"
	specs["Rear Camera"] = "50 MP + 50 MP + 50 MP"
	specs["Front Camera"] = "32 MP"
	specs["Battery"] = "5,400 mAh"
	specs["Operating System"] = "Android 15, HyperOS"
	specs["Available Colors"] = "Black, White, Green, Purple, Silver"
	specs["Announcement Date"] = "October 2024"
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
