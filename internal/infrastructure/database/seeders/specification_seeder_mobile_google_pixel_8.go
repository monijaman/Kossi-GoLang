package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileGooglePixel8 seeds specifications/options for product 'google-pixel-8'
type SpecificationSeederMobileGooglePixel8 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileGooglePixel8 creates a new seeder instance
func NewSpecificationSeederMobileGooglePixel8() *SpecificationSeederMobileGooglePixel8 {
	return &SpecificationSeederMobileGooglePixel8{BaseSeeder: BaseSeeder{name: "Specifications for Google Pixel 8"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileGooglePixel8) getBanglaTranslations() map[string]string {
	return map[string]string{
		"10.5 MP": "১০.৫ মেগাপিক্সেল",
		"1080 x 2400 pixels": "১০৮০ x ২৪০০ পিক্সেল",
		"128 GB / 256 GB": "১২৮ জিবি / ২৫৬ GB",
		"150.5 x 70.8 x 8.9 mm": "১৫০.৫ x ৭০.৮ x ৮.৯ মিমি",
		"187 g": "১৮৭ g",
		"4,575 mAh": "৪,৫৭৫ এমএএইচ",
		"50 MP + 12 MP": "৫০ মেগাপিক্সেল + ১২ মেগাপিক্সেল",
		"5G": "৫G",
		"6.2 inches": "৬.২ ইঞ্চি",
		"60-120Hz": "৬০-১২০Hz",
		"8 GB": "৮ GB",
		"Android 14": "অ্যান্ড্রয়েড ১৪",
		"Glass front/back, aluminum frame": "গ্লাস সামনে/পেছনে, অ্যালুমিনিয়াম ফ্রেম",
		"Google Tensor G3": "গুগল টেনসর G৩",
		"Google Tensor G3 (4 nm)": "গুগল টেনসর G৩ (৪ ন্যানোমিটার)",
		"IP68": "IP৬৮",
		"Immortalis-G715s MC10": "Immবাtalis-G৭১৫s MC১০",
		"OLED, 120Hz, HDR10+, 2000 nits": "ওএলইডি, ১২০Hz, এইচডিআর১০+, ২০০০ nits",
		"October 2023": "অক্টোবর ২০২৩",
	}
}

// Seed inserts specification records for the product identified by slug 'google-pixel-8'
func (s *SpecificationSeederMobileGooglePixel8) Seed(db *gorm.DB) error {
	productSlug := "google-pixel-8"

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

	// Override model-specific values for Google Pixel 8
	specs["Display Size"] = "6.2 inches"
	specs["Processor"] = "Google Tensor G3"
	specs["Chipset"] = "Google Tensor G3 (4 nm)"
	specs["Cpu Type"] = "Nona-core"
	specs["Gpu Type"] = "Immortalis-G715s MC10"
	specs["Ram"] = "8 GB"
	specs["Storage"] = "128 GB / 256 GB"
	specs["Display Type"] = "OLED, 120Hz, HDR10+, 2000 nits"
	specs["Resolution"] = "1080 x 2400 pixels"
	specs["Screen Protection"] = "Corning Gorilla Glass Victus"
	specs["Refresh Rate"] = "60-120Hz"
	specs["Build Material"] = "Glass front/back, aluminum frame"
	specs["Weight"] = "187 g"
	specs["Dimensions"] = "150.5 x 70.8 x 8.9 mm"
	specs["Water Resistance"] = "IP68"
	specs["Network Technology"] = "5G"
	specs["Rear Camera"] = "50 MP + 12 MP"
	specs["Front Camera"] = "10.5 MP"
	specs["Battery"] = "4,575 mAh"
	specs["Operating System"] = "Android 14"
	specs["Available Colors"] = "Obsidian, Hazel, Rose, Mint"
	specs["Announcement Date"] = "October 2023"
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
