package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileGooglePixel8a seeds specifications/options for product 'google-pixel-8a'
type SpecificationSeederMobileGooglePixel8a struct {
	BaseSeeder
}

// NewSpecificationSeederMobileGooglePixel8a creates a new seeder instance
func NewSpecificationSeederMobileGooglePixel8a() *SpecificationSeederMobileGooglePixel8a {
	return &SpecificationSeederMobileGooglePixel8a{BaseSeeder: BaseSeeder{name: "Specifications for Google Pixel 8a"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileGooglePixel8a) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080 x 2400 pixels": "১০৮০ x ২৪০০ pixels",
		"120Hz": "১২০Hz",
		"128 GB / 256 GB": "১২৮ GB / ২৫৬ GB",
		"13 MP": "১৩ MP",
		"152.1 x 72.7 x 8.9 mm": "১৫২.১ x ৭২.৭ x ৮.৯ মিমি",
		"188 g": "১৮৮ g",
		"4,492 mAh": "৪,৪৯২ এমএএইচ",
		"5G": "৫G",
		"6.1 inches": "৬.১ ইঞ্চি",
		"64 MP + 13 MP": "৬৪ MP + ১৩ MP",
		"8 GB": "৮ GB",
		"Android 14": "Android ১৪",
		"Corning Gorilla Glass 3": "Corning Gorilla Glass ৩",
		"Glass front, plastic back, aluminum frame": "গ্লাস সামনে, প্লাস্টিক পেছনে, অ্যালুমিনিয়াম ফ্রেম",
		"Google Tensor G3": "Google Tensor G৩",
		"Google Tensor G3 (4 nm)": "Google Tensor G৩ (৪ nm)",
		"IP67": "IP৬৭",
		"Immortalis-G715s MC10": "Immortalis-G৭১৫s MC১০",
		"May 2024": "May ২০২৪",
		"OLED, 120Hz, HDR": "OLED, ১২০Hz, HDR",
	}
}

// Seed inserts specification records for the product identified by slug 'google-pixel-8a'
func (s *SpecificationSeederMobileGooglePixel8a) Seed(db *gorm.DB) error {
	productSlug := "google-pixel-8a"

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

	// Override model-specific values for Google Pixel 8a
	specs["Display Size"] = "6.1 inches"
	specs["Processor"] = "Google Tensor G3"
	specs["Chipset"] = "Google Tensor G3 (4 nm)"
	specs["Cpu Type"] = "Nona-core"
	specs["Gpu Type"] = "Immortalis-G715s MC10"
	specs["Ram"] = "8 GB"
	specs["Storage"] = "128 GB / 256 GB"
	specs["Display Type"] = "OLED, 120Hz, HDR"
	specs["Resolution"] = "1080 x 2400 pixels"
	specs["Screen Protection"] = "Corning Gorilla Glass 3"
	specs["Refresh Rate"] = "120Hz"
	specs["Build Material"] = "Glass front, plastic back, aluminum frame"
	specs["Weight"] = "188 g"
	specs["Dimensions"] = "152.1 x 72.7 x 8.9 mm"
	specs["Water Resistance"] = "IP67"
	specs["Network Technology"] = "5G"
	specs["Rear Camera"] = "64 MP + 13 MP"
	specs["Front Camera"] = "13 MP"
	specs["Battery"] = "4,492 mAh"
	specs["Operating System"] = "Android 14"
	specs["Available Colors"] = "Obsidian, Porcelain, Bay, Aloe"
	specs["Announcement Date"] = "May 2024"
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
