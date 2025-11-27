package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileRedmiNote14Pro5g seeds specifications/options for product 'redmi-note-14-pro-5g'
type SpecificationSeederMobileRedmiNote14Pro5g struct {
	BaseSeeder
}

// NewSpecificationSeederMobileRedmiNote14Pro5g creates a new seeder instance
func NewSpecificationSeederMobileRedmiNote14Pro5g() *SpecificationSeederMobileRedmiNote14Pro5g {
	return &SpecificationSeederMobileRedmiNote14Pro5g{BaseSeeder: BaseSeeder{name: "Specifications for Redmi Note 14 Pro+ 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileRedmiNote14Pro5g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"12 GB / 16 GB": "১২ GB / ১৬ GB",
		"120Hz": "১২০Hz",
		"1220 x 2712 pixels": "১২২০ x ২৭১২ pixels",
		"162.5 x 74.7 x 8.7 mm": "১৬২.৫ x ৭৪.৭ x ৮.৭ মিমি",
		"20 MP": "২০ MP",
		"210 g": "২১০ g",
		"256 GB / 512 GB": "২৫৬ GB / ৫১২ GB",
		"50 MP + 50 MP + 8 MP": "৫০ MP + ৫০ MP + ৮ MP",
		"5G": "৫G",
		"6,200 mAh": "৬,২০০ এমএএইচ",
		"6.67 inches": "৬.৬৭ ইঞ্চি",
		"AMOLED, 68B colors, 120Hz, Dolby Vision, 3000 nits": "AMOLED, ৬৮B colors, ১২০Hz, Dolby Vision, ৩০০০ nits",
		"Adreno 710": "Adreno ৭১০",
		"Android 14, HyperOS": "Android ১৪, HyperOS",
		"Corning Gorilla Glass Victus 2": "Corning Gorilla Glass Victus ২",
		"Glass front/back, plastic frame": "গ্লাস সামনে/back, plastic frame",
		"IP68": "IP৬৮",
		"Qualcomm Snapdragon 7s Gen 3 (4 nm)": "Qualcomm Snapdragon ৭s Gen ৩ (৪ nm)",
		"September 2024": "September ২০২৪",
		"Snapdragon 7s Gen 3": "Snapdragon ৭s Gen ৩",
		"Starry White, Midnight Black, Mirror Porcelain White": "Starry সাদা, Midnight কালো, Mirror Porcelain সাদা",
	}
}

// Seed inserts specification records for the product identified by slug 'redmi-note-14-pro-5g'
func (s *SpecificationSeederMobileRedmiNote14Pro5g) Seed(db *gorm.DB) error {
	productSlug := "redmi-note-14-pro-5g"

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

	// Override model-specific values for Redmi Note 14 Pro+ 5G
	specs["Display Size"] = "6.67 inches"
	specs["Processor"] = "Snapdragon 7s Gen 3"
	specs["Chipset"] = "Qualcomm Snapdragon 7s Gen 3 (4 nm)"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "Adreno 710"
	specs["Ram"] = "12 GB / 16 GB"
	specs["Storage"] = "256 GB / 512 GB"
	specs["Display Type"] = "AMOLED, 68B colors, 120Hz, Dolby Vision, 3000 nits"
	specs["Resolution"] = "1220 x 2712 pixels"
	specs["Screen Protection"] = "Corning Gorilla Glass Victus 2"
	specs["Refresh Rate"] = "120Hz"
	specs["Build Material"] = "Glass front/back, plastic frame"
	specs["Weight"] = "210 g"
	specs["Dimensions"] = "162.5 x 74.7 x 8.7 mm"
	specs["Water Resistance"] = "IP68"
	specs["Network Technology"] = "5G"
	specs["Rear Camera"] = "50 MP + 50 MP + 8 MP"
	specs["Front Camera"] = "20 MP"
	specs["Battery"] = "6,200 mAh"
	specs["Operating System"] = "Android 14, HyperOS"
	specs["Available Colors"] = "Starry White, Midnight Black, Mirror Porcelain White"
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
