package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileGooglePixel9Pro seeds specifications/options for product 'google-pixel-9-pro'
type SpecificationSeederMobileGooglePixel9Pro struct {
	BaseSeeder
}

// NewSpecificationSeederMobileGooglePixel9Pro creates a new seeder instance
func NewSpecificationSeederMobileGooglePixel9Pro() *SpecificationSeederMobileGooglePixel9Pro {
	return &SpecificationSeederMobileGooglePixel9Pro{BaseSeeder: BaseSeeder{name: "Specifications for Google Pixel 9 Pro"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileGooglePixel9Pro) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1-120Hz":                             "১-১২০Hz",
		"128 GB / 256 GB / 512 GB / 1 TB":     "১২৮ জিবি / ২৫৬ জিবি / ৫১২ জিবি / ১ টিবি",
		"1280 x 2856 pixels":                  "১২৮০ x ২৮৫৬ পিক্সেল",
		"152.8 x 72 x 8.5 mm":                 "১৫২.৮ x ৭২ x ৮.৫ মিমি",
		"16 GB":                               "১৬ GB",
		"199 g":                               "১৯৯ g",
		"4,700 mAh":                           "৪,৭০০ এমএএইচ",
		"42 MP":                               "৪২ মেগাপিক্সেল",
		"50 MP + 48 MP + 48 MP":               "৫০ মেগাপিক্সেল + ৪৮ মেগাপিক্সেল + ৪৮ মেগাপিক্সেল",
		"5G":                                  "৫G",
		"6.3 inches":                          "৬.৩ ইঞ্চি",
		"Android 15":                          "অ্যান্ড্রয়েড ১৫",
		"August 2024":                         "আগস্ট ২০২৪",
		"Corning Gorilla Glass Victus 2":      "কর্নিং গরিলা গ্লাস ভিক্টাস ২",
		"Glass front/back, aluminum frame":    "গ্লাস সামনে/পেছনে, অ্যালুমিনিয়াম ফ্রেম",
		"Google Tensor G4":                    "গুগল টেনসর G৪",
		"Google Tensor G4 (4 nm)":             "গুগল টেনসর G৪ (৪ ন্যানোমিটার)",
		"IP68":                                "IP৬৮",
		"LTPO OLED, 120Hz, HDR10+, 3000 nits": "এলটিপিও ওএলইডি, ১২০Hz, এইচডিআর১০+, ৩০০০ nits",
		"Mali-G715 MC7":                       "মালি-G৭১৫ MC৭",
	}
}

// Seed inserts specification records for the product identified by slug 'google-pixel-9-pro'
func (s *SpecificationSeederMobileGooglePixel9Pro) Seed(db *gorm.DB) error {
	productSlug := "google-pixel-9-pro"

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

	// Override model-specific values for Google Pixel 9 Pro
	specs["Display Size"] = "6.3 inches"
	specs["Processor"] = "Google Tensor G4"
	specs["Chipset"] = "Google Tensor G4 (4 nm)"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "Mali-G715 MC7"
	specs["Ram"] = "16 GB"
	specs["Storage"] = "128 GB / 256 GB / 512 GB / 1 TB"
	specs["Display Type"] = "LTPO OLED, 120Hz, HDR10+, 3000 nits"
	specs["Resolution"] = "1280 x 2856 pixels"
	specs["Screen Protection"] = "Corning Gorilla Glass Victus 2"
	specs["Refresh Rate"] = "1-120Hz"
	specs["Build Material"] = "Glass front/back, aluminum frame"
	specs["Weight"] = "199 g"
	specs["Dimensions"] = "152.8 x 72 x 8.5 mm"
	specs["Water Resistance"] = "IP68"
	specs["Network Technology"] = "5G"
	specs["Rear Camera"] = "50 MP + 48 MP + 48 MP"
	specs["Front Camera"] = "42 MP"
	specs["Battery"] = "4,700 mAh"
	specs["Operating System"] = "Android 15"
	specs["Available Colors"] = "Obsidian, Porcelain, Hazel, Rose Quartz"
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
