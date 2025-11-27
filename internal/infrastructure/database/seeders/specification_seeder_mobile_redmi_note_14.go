package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileRedmiNote14 seeds specifications/options for product 'redmi-note-14'
type SpecificationSeederMobileRedmiNote14 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileRedmiNote14 creates a new seeder instance
func NewSpecificationSeederMobileRedmiNote14() *SpecificationSeederMobileRedmiNote14 {
	return &SpecificationSeederMobileRedmiNote14{BaseSeeder: BaseSeeder{name: "Specifications for Redmi Note 14"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileRedmiNote14) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080 x 2400 pixels": "১০৮০ x ২৪০০ pixels",
		"120Hz": "১২০Hz",
		"128 GB / 256 GB": "১২৮ GB / ২৫৬ GB",
		"16 MP": "১৬ MP",
		"162.4 x 75.7 x 8 mm": "১৬২.৪ x ৭৫.৭ x ৮ মিমি",
		"190 g": "১৯০ g",
		"5,110 mAh": "৫,১১০ এমএএইচ",
		"50 MP + 2 MP": "৫০ MP + ২ MP",
		"5G": "৫G",
		"6 GB / 8 GB / 12 GB": "৬ GB / ৮ GB / ১২ GB",
		"6.67 inches": "৬.৬৭ ইঞ্চি",
		"Android 14, HyperOS": "Android ১৪, HyperOS",
		"Corning Gorilla Glass 5": "Corning Gorilla Glass ৫",
		"Dimensity 7025 Ultra": "Dimensity ৭০২৫ Ultra",
		"Glass front, plastic frame, plastic back": "গ্লাস সামনে, plastic frame, প্লাস্টিক পেছনে",
		"IMG BXM-8-256": "IMG BXM-৮-২৫৬",
		"IP64": "IP৬৪",
		"Mediatek Dimensity 7025 Ultra (6 nm)": "Mediatek Dimensity ৭০২৫ Ultra (৬ nm)",
		"OLED, 120Hz, 2100 nits": "OLED, ১২০Hz, ২১০০ nits",
		"September 2024": "September ২০২৪",
		"Starry White, Midnight Black, Phantom Blue": "Starry সাদা, Midnight কালো, Phantom নীল",
	}
}

// Seed inserts specification records for the product identified by slug 'redmi-note-14'
func (s *SpecificationSeederMobileRedmiNote14) Seed(db *gorm.DB) error {
	productSlug := "redmi-note-14"

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

	// Override model-specific values for Redmi Note 14
	specs["Display Size"] = "6.67 inches"
	specs["Processor"] = "Dimensity 7025 Ultra"
	specs["Chipset"] = "Mediatek Dimensity 7025 Ultra (6 nm)"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "IMG BXM-8-256"
	specs["Ram"] = "6 GB / 8 GB / 12 GB"
	specs["Storage"] = "128 GB / 256 GB"
	specs["Display Type"] = "OLED, 120Hz, 2100 nits"
	specs["Resolution"] = "1080 x 2400 pixels"
	specs["Screen Protection"] = "Corning Gorilla Glass 5"
	specs["Refresh Rate"] = "120Hz"
	specs["Build Material"] = "Glass front, plastic frame, plastic back"
	specs["Weight"] = "190 g"
	specs["Dimensions"] = "162.4 x 75.7 x 8 mm"
	specs["Water Resistance"] = "IP64"
	specs["Network Technology"] = "5G"
	specs["Rear Camera"] = "50 MP + 2 MP"
	specs["Front Camera"] = "16 MP"
	specs["Battery"] = "5,110 mAh"
	specs["Operating System"] = "Android 14, HyperOS"
	specs["Available Colors"] = "Starry White, Midnight Black, Phantom Blue"
	specs["Announcement Date"] = "September 2024"
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
