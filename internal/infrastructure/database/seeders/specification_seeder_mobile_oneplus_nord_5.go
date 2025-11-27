package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileOneplusNord5 seeds specifications/options for product 'oneplus-nord-5'
type SpecificationSeederMobileOneplusNord5 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileOneplusNord5 creates a new seeder instance
func NewSpecificationSeederMobileOneplusNord5() *SpecificationSeederMobileOneplusNord5 {
	return &SpecificationSeederMobileOneplusNord5{BaseSeeder: BaseSeeder{name: "Specifications for OnePlus Nord 5"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileOneplusNord5) getBanglaTranslations() map[string]string {
	return map[string]string{
		"120Hz": "১২০Hz",
		"1240 x 2772 pixels": "১২৪০ x ২৭৭২ pixels",
		"128 GB / 256 GB / 512 GB": "১২৮ GB / ২৫৬ GB / ৫১২ GB",
		"16 MP": "১৬ MP",
		"162.6 x 75 x 8 mm": "১৬২.৬ x ৭৫ x ৮ মিমি",
		"199.5 g": "১৯৯.৫ g",
		"5,500 mAh": "৫,৫০০ এমএএইচ",
		"50 MP + 8 MP": "৫০ MP + ৮ MP",
		"5G": "৫G",
		"6.74 inches": "৬.৭৪ ইঞ্চি",
		"8 GB / 12 GB / 16 GB": "৮ GB / ১২ GB / ১৬ GB",
		"Adreno 732": "Adreno ৭৩২",
		"Android 14, OxygenOS 14.1": "Android ১৪, OxygenOS ১৪.১",
		"Fluid AMOLED, 120Hz, HDR10+, 2150 nits": "Fluid AMOLED, ১২০Hz, HDR১০+, ২১৫০ nits",
		"Glass front, plastic frame, glass back": "গ্লাস সামনে, plastic frame, গ্লাস পেছনে",
		"IP65": "IP৬৫",
		"July 2024": "July ২০২৪",
		"Mercurial Silver, Oasis Green, Obsidian Midnight": "Mercurial রূপালী, Oasis সবুজ, Obsidian Midnight",
		"Qualcomm Snapdragon 7+ Gen 3 (4 nm)": "Qualcomm Snapdragon ৭+ Gen ৩ (৪ nm)",
		"Snapdragon 7+ Gen 3": "Snapdragon ৭+ Gen ৩",
	}
}

// Seed inserts specification records for the product identified by slug 'oneplus-nord-5'
func (s *SpecificationSeederMobileOneplusNord5) Seed(db *gorm.DB) error {
	productSlug := "oneplus-nord-5"

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

	// Override model-specific values for OnePlus Nord 5
	specs["Display Size"] = "6.74 inches"
	specs["Processor"] = "Snapdragon 7+ Gen 3"
	specs["Chipset"] = "Qualcomm Snapdragon 7+ Gen 3 (4 nm)"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "Adreno 732"
	specs["Ram"] = "8 GB / 12 GB / 16 GB"
	specs["Storage"] = "128 GB / 256 GB / 512 GB"
	specs["Display Type"] = "Fluid AMOLED, 120Hz, HDR10+, 2150 nits"
	specs["Resolution"] = "1240 x 2772 pixels"
	specs["Screen Protection"] = "Asahi Glass"
	specs["Refresh Rate"] = "120Hz"
	specs["Build Material"] = "Glass front, plastic frame, glass back"
	specs["Weight"] = "199.5 g"
	specs["Dimensions"] = "162.6 x 75 x 8 mm"
	specs["Water Resistance"] = "IP65"
	specs["Network Technology"] = "5G"
	specs["Rear Camera"] = "50 MP + 8 MP"
	specs["Front Camera"] = "16 MP"
	specs["Battery"] = "5,500 mAh"
	specs["Operating System"] = "Android 14, OxygenOS 14.1"
	specs["Available Colors"] = "Mercurial Silver, Oasis Green, Obsidian Midnight"
	specs["Announcement Date"] = "July 2024"
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
