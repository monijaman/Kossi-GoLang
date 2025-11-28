package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileXiaomi13Ultra seeds specifications/options for product 'xiaomi-13-ultra'
type SpecificationSeederMobileXiaomi13Ultra struct {
	BaseSeeder
}

// NewSpecificationSeederMobileXiaomi13Ultra creates a new seeder instance
func NewSpecificationSeederMobileXiaomi13Ultra() *SpecificationSeederMobileXiaomi13Ultra {
	return &SpecificationSeederMobileXiaomi13Ultra{BaseSeeder: BaseSeeder{name: "Specifications for Xiaomi 13 Ultra"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileXiaomi13Ultra) getBanglaTranslations() map[string]string {
	return map[string]string{
		"12 GB / 16 GB": "১২ জিবি / ১৬ GB",
		"120Hz": "১২০Hz",
		"1440 x 3200 pixels": "১৪৪০ x ৩২০০ পিক্সেল",
		"163.2 x 74.6 x 9.1 mm": "১৬৩.২ x ৭৪.৬ x ৯.১ মিমি",
		"227 g": "২২৭ g",
		"256 GB / 512 GB / 1 TB": "২৫৬ জিবি / ৫১২ জিবি / ১ টিবি",
		"32 MP": "৩২ মেগাপিক্সেল",
		"5,000 mAh": "৫,০০০ এমএএইচ",
		"50 MP + 50 MP + 50 MP + 50 MP": "৫০ মেগাপিক্সেল + ৫০ মেগাপিক্সেল + ৫০ মেগাপিক্সেল + ৫০ মেগাপিক্সেল",
		"5G": "৫G",
		"6.73 inches": "৬.৭৩ ইঞ্চি",
		"Adreno 740": "অ্যাড্রেনো ৭৪০",
		"Android 13, upgradable": "অ্যান্ড্রয়েড ১৩, আপগ্রেডযোগ্য",
		"April 2023": "এপ্রিল ২০২৩",
		"Black, Olive Green, White": "কালো, অলিভ সবুজ, সাদা",
		"Glass front, eco leather back, aluminum frame": "গ্লাস সামনে, eco চামড়া পেছনে, অ্যালুমিনিয়াম ফ্রেম",
		"IP68": "IP৬৮",
		"LTPO AMOLED, 120Hz, Dolby Vision, HDR10+, 2600 nits": "এলটিপিও অ্যামোলেড, ১২০Hz, ডলবি ভিশন, এইচডিআর১০+, ২৬০০ nits",
		"Qualcomm SM8550-AB Snapdragon 8 Gen 2 (4 nm)": "কোয়ালকম SM৮৫৫০-AB স্ন্যাপড্রাগন ৮ জেন ২ (৪ ন্যানোমিটার)",
		"Snapdragon 8 Gen 2": "স্ন্যাপড্রাগন ৮ জেন ২",
	}
}

// Seed inserts specification records for the product identified by slug 'xiaomi-13-ultra'
func (s *SpecificationSeederMobileXiaomi13Ultra) Seed(db *gorm.DB) error {
	productSlug := "xiaomi-13-ultra"

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

	// Override model-specific values for Xiaomi 13 Ultra
	specs["Display Size"] = "6.73 inches"
	specs["Processor"] = "Snapdragon 8 Gen 2"
	specs["Chipset"] = "Qualcomm SM8550-AB Snapdragon 8 Gen 2 (4 nm)"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "Adreno 740"
	specs["Ram"] = "12 GB / 16 GB"
	specs["Storage"] = "256 GB / 512 GB / 1 TB"
	specs["Display Type"] = "LTPO AMOLED, 120Hz, Dolby Vision, HDR10+, 2600 nits"
	specs["Resolution"] = "1440 x 3200 pixels"
	specs["Screen Protection"] = "Corning Gorilla Glass Victus"
	specs["Refresh Rate"] = "120Hz"
	specs["Build Material"] = "Glass front, eco leather back, aluminum frame"
	specs["Weight"] = "227 g"
	specs["Dimensions"] = "163.2 x 74.6 x 9.1 mm"
	specs["Water Resistance"] = "IP68"
	specs["Network Technology"] = "5G"
	specs["Rear Camera"] = "50 MP + 50 MP + 50 MP + 50 MP"
	specs["Front Camera"] = "32 MP"
	specs["Battery"] = "5,000 mAh"
	specs["Operating System"] = "Android 13, upgradable"
	specs["Available Colors"] = "Black, Olive Green, White"
	specs["Announcement Date"] = "April 2023"
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
