package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileGooglePixel9 seeds specifications/options for product 'google-pixel-9'
type SpecificationSeederMobileGooglePixel9 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileGooglePixel9 creates a new seeder instance
func NewSpecificationSeederMobileGooglePixel9() *SpecificationSeederMobileGooglePixel9 {
	return &SpecificationSeederMobileGooglePixel9{BaseSeeder: BaseSeeder{name: "Specifications for Google Pixel 9"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileGooglePixel9) getBanglaTranslations() map[string]string {
	return map[string]string{
		"10.5 MP": "১০.৫ মেগাপিক্সেল",
		"1080 x 2424 pixels": "১০৮০ x ২৪২৪ পিক্সেল",
		"12 GB": "১২ GB",
		"128 GB / 256 GB": "১২৮ জিবি / ২৫৬ GB",
		"152.8 x 72 x 8.5 mm": "১৫২.৮ x ৭২ x ৮.৫ মিমি",
		"198 g": "১৯৮ g",
		"4,700 mAh": "৪,৭০০ এমএএইচ",
		"50 MP + 48 MP": "৫০ মেগাপিক্সেল + ৪৮ মেগাপিক্সেল",
		"5G": "৫G",
		"6.3 inches": "৬.৩ ইঞ্চি",
		"60-120Hz": "৬০-১২০Hz",
		"Android 15": "অ্যান্ড্রয়েড ১৫",
		"August 2024": "আগস্ট ২০২৪",
		"Corning Gorilla Glass Victus 2": "কর্নিং গরিলা গ্লাস ভিক্টাস ২",
		"Glass front/back, aluminum frame": "গ্লাস সামনে/পেছনে, অ্যালুমিনিয়াম ফ্রেম",
		"Google Tensor G4": "গুগল টেনসর G৪",
		"Google Tensor G4 (4 nm)": "গুগল টেনসর G৪ (৪ ন্যানোমিটার)",
		"IP68": "IP৬৮",
		"Mali-G715 MC7": "মালি-G৭১৫ MC৭",
		"OLED, 120Hz, HDR10+, 2700 nits": "ওএলইডি, ১২০Hz, এইচডিআর১০+, ২৭০০ nits",
	}
}

// Seed inserts specification records for the product identified by slug 'google-pixel-9'
func (s *SpecificationSeederMobileGooglePixel9) Seed(db *gorm.DB) error {
	productSlug := "google-pixel-9"

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

	// Override model-specific values for Google Pixel 9
	specs["Display Size"] = "6.3 inches"
	specs["Processor"] = "Google Tensor G4"
	specs["Chipset"] = "Google Tensor G4 (4 nm)"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "Mali-G715 MC7"
	specs["Ram"] = "12 GB"
	specs["Storage"] = "128 GB / 256 GB"
	specs["Display Type"] = "OLED, 120Hz, HDR10+, 2700 nits"
	specs["Resolution"] = "1080 x 2424 pixels"
	specs["Screen Protection"] = "Corning Gorilla Glass Victus 2"
	specs["Refresh Rate"] = "60-120Hz"
	specs["Build Material"] = "Glass front/back, aluminum frame"
	specs["Weight"] = "198 g"
	specs["Dimensions"] = "152.8 x 72 x 8.5 mm"
	specs["Water Resistance"] = "IP68"
	specs["Network Technology"] = "5G"
	specs["Rear Camera"] = "50 MP + 48 MP"
	specs["Front Camera"] = "10.5 MP"
	specs["Battery"] = "4,700 mAh"
	specs["Operating System"] = "Android 15"
	specs["Available Colors"] = "Obsidian, Porcelain, Wintergreen, Peony"
	specs["Announcement Date"] = "August 2024"
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
