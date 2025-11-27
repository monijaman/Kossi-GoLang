package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileGooglePixel9ProFold seeds specifications/options for product 'google-pixel-9-pro-fold'
type SpecificationSeederMobileGooglePixel9ProFold struct {
	BaseSeeder
}

// NewSpecificationSeederMobileGooglePixel9ProFold creates a new seeder instance
func NewSpecificationSeederMobileGooglePixel9ProFold() *SpecificationSeederMobileGooglePixel9ProFold {
	return &SpecificationSeederMobileGooglePixel9ProFold{BaseSeeder: BaseSeeder{name: "Specifications for Google Pixel 9 Pro Fold"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileGooglePixel9ProFold) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1-120Hz": "১-১২০Hz",
		"10 MP (Cover) + 10 MP (Inner)": "১০ MP (Cover) + ১০ MP (Inner)",
		"16 GB": "১৬ GB",
		"2076 x 2152 pixels": "২০৭৬ x ২১৫২ pixels",
		"256 GB / 512 GB": "২৫৬ GB / ৫১২ GB",
		"257 g": "২৫৭ g",
		"4,650 mAh": "৪,৬৫০ এমএএইচ",
		"48 MP + 10.5 MP + 10.8 MP": "৪৮ MP + ১০.৫ MP + ১০.৮ MP",
		"5G": "৫G",
		"8.0 inches (Main) / 6.3 inches (Cover)": "৮.০ ইঞ্চি (Main) / ৬.৩ ইঞ্চি (Cover)",
		"Android 15": "Android ১৫",
		"August 2024": "August ২০২৪",
		"Corning Gorilla Glass Victus 2": "Corning Gorilla Glass Victus ২",
		"Foldable LTPO OLED, 120Hz, HDR10+, 2700 nits": "Foldable LTPO OLED, ১২০Hz, HDR১০+, ২৭০০ nits",
		"Glass front/back, aluminum frame": "গ্লাস সামনে/back, অ্যালুমিনিয়াম ফ্রেম",
		"Google Tensor G4": "Google Tensor G৪",
		"Google Tensor G4 (4 nm)": "Google Tensor G৪ (৪ nm)",
		"IPX8": "IPX৮",
		"Mali-G715 MC7": "Mali-G৭১৫ MC৭",
		"Unfolded: 155.2 x 150.2 x 5.1 mm / Folded: 155.2 x 77.1 x 10.5 mm": "Unfolded: ১৫৫.২ x ১৫০.২ x ৫.১ মিমি / Folded: ১৫৫.২ x ৭৭.১ x ১০.৫ মিমি",
	}
}

// Seed inserts specification records for the product identified by slug 'google-pixel-9-pro-fold'
func (s *SpecificationSeederMobileGooglePixel9ProFold) Seed(db *gorm.DB) error {
	productSlug := "google-pixel-9-pro-fold"

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

	// Override model-specific values for Google Pixel 9 Pro Fold
	specs["Display Size"] = "8.0 inches (Main) / 6.3 inches (Cover)"
	specs["Processor"] = "Google Tensor G4"
	specs["Chipset"] = "Google Tensor G4 (4 nm)"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "Mali-G715 MC7"
	specs["Ram"] = "16 GB"
	specs["Storage"] = "256 GB / 512 GB"
	specs["Display Type"] = "Foldable LTPO OLED, 120Hz, HDR10+, 2700 nits"
	specs["Resolution"] = "2076 x 2152 pixels"
	specs["Screen Protection"] = "Corning Gorilla Glass Victus 2"
	specs["Refresh Rate"] = "1-120Hz"
	specs["Build Material"] = "Glass front/back, aluminum frame"
	specs["Weight"] = "257 g"
	specs["Dimensions"] = "Unfolded: 155.2 x 150.2 x 5.1 mm / Folded: 155.2 x 77.1 x 10.5 mm"
	specs["Water Resistance"] = "IPX8"
	specs["Network Technology"] = "5G"
	specs["Rear Camera"] = "48 MP + 10.5 MP + 10.8 MP"
	specs["Front Camera"] = "10 MP (Cover) + 10 MP (Inner)"
	specs["Battery"] = "4,650 mAh"
	specs["Operating System"] = "Android 15"
	specs["Available Colors"] = "Obsidian, Porcelain"
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
