package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileOneplusNordCe4Lite5g seeds specifications/options for product 'oneplus-nord-ce4-lite-5g'
type SpecificationSeederMobileOneplusNordCe4Lite5g struct {
	BaseSeeder
}

// NewSpecificationSeederMobileOneplusNordCe4Lite5g creates a new seeder instance
func NewSpecificationSeederMobileOneplusNordCe4Lite5g() *SpecificationSeederMobileOneplusNordCe4Lite5g {
	return &SpecificationSeederMobileOneplusNordCe4Lite5g{BaseSeeder: BaseSeeder{name: "Specifications for OnePlus Nord CE4 Lite 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileOneplusNordCe4Lite5g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080 x 2400 pixels": "১০৮০ x ২৪০০ pixels",
		"120Hz": "১২০Hz",
		"128 GB / 256 GB": "১২৮ GB / ২৫৬ GB",
		"16 MP": "১৬ MP",
		"162.9 x 75.6 x 8.1 mm": "১৬২.৯ x ৭৫.৬ x ৮.১ মিমি",
		"191 g": "১৯১ g",
		"5,110 mAh": "৫,১১০ এমএএইচ",
		"50 MP + 2 MP": "৫০ MP + ২ MP",
		"5G": "৫G",
		"6.67 inches": "৬.৬৭ ইঞ্চি",
		"8 GB": "৮ GB",
		"AMOLED, 120Hz, 2100 nits": "AMOLED, ১২০Hz, ২১০০ nits",
		"Adreno 619": "Adreno ৬১৯",
		"Android 14, OxygenOS 14": "Android ১৪, OxygenOS ১৪",
		"Glass front, plastic frame, plastic back": "গ্লাস সামনে, plastic frame, প্লাস্টিক পেছনে",
		"IP54": "IP৫৪",
		"June 2024": "June ২০২৪",
		"Qualcomm SM6375 Snapdragon 695 5G (6 nm)": "Qualcomm SM৬৩৭৫ Snapdragon ৬৯৫ ৫G (৬ nm)",
		"Snapdragon 695 5G": "Snapdragon ৬৯৫ ৫G",
		"Super Silver, Mega Blue, Ultra Orange": "Super রূপালী, Mega নীল, Ultra Orange",
	}
}

// Seed inserts specification records for the product identified by slug 'oneplus-nord-ce4-lite-5g'
func (s *SpecificationSeederMobileOneplusNordCe4Lite5g) Seed(db *gorm.DB) error {
	productSlug := "oneplus-nord-ce4-lite-5g"

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

	// Override model-specific values for OnePlus Nord CE4 Lite 5G
	specs["Display Size"] = "6.67 inches"
	specs["Processor"] = "Snapdragon 695 5G"
	specs["Chipset"] = "Qualcomm SM6375 Snapdragon 695 5G (6 nm)"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "Adreno 619"
	specs["Ram"] = "8 GB"
	specs["Storage"] = "128 GB / 256 GB"
	specs["Display Type"] = "AMOLED, 120Hz, 2100 nits"
	specs["Resolution"] = "1080 x 2400 pixels"
	specs["Screen Protection"] = "Asahi Glass"
	specs["Refresh Rate"] = "120Hz"
	specs["Build Material"] = "Glass front, plastic frame, plastic back"
	specs["Weight"] = "191 g"
	specs["Dimensions"] = "162.9 x 75.6 x 8.1 mm"
	specs["Water Resistance"] = "IP54"
	specs["Network Technology"] = "5G"
	specs["Rear Camera"] = "50 MP + 2 MP"
	specs["Front Camera"] = "16 MP"
	specs["Battery"] = "5,110 mAh"
	specs["Operating System"] = "Android 14, OxygenOS 14"
	specs["Available Colors"] = "Super Silver, Mega Blue, Ultra Orange"
	specs["Announcement Date"] = "June 2024"
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
