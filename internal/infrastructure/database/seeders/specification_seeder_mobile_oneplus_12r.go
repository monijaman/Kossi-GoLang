package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileOneplus12r seeds specifications/options for product 'oneplus-12r'
type SpecificationSeederMobileOneplus12r struct {
	BaseSeeder
}

// NewSpecificationSeederMobileOneplus12r creates a new seeder instance
func NewSpecificationSeederMobileOneplus12r() *SpecificationSeederMobileOneplus12r {
	return &SpecificationSeederMobileOneplus12r{BaseSeeder: BaseSeeder{name: "Specifications for OnePlus 12R"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileOneplus12r) getBanglaTranslations() map[string]string {
	return map[string]string{
		"120Hz": "১২০Hz",
		"1264 x 2780 pixels": "১২৬৪ x ২৭৮০ pixels",
		"128 GB / 256 GB": "১২৮ GB / ২৫৬ GB",
		"16 MP": "১৬ MP",
		"163.3 x 75.3 x 8.8 mm": "১৬৩.৩ x ৭৫.৩ x ৮.৮ মিমি",
		"207 g": "২০৭ g",
		"5,500 mAh": "৫,৫০০ এমএএইচ",
		"50 MP + 8 MP + 2 MP": "৫০ MP + ৮ MP + ২ MP",
		"5G": "৫G",
		"6.78 inches": "৬.৭৮ ইঞ্চি",
		"8 GB / 16 GB": "৮ GB / ১৬ GB",
		"Adreno 740": "Adreno ৭৪০",
		"Android 14, OxygenOS 14": "Android ১৪, OxygenOS ১৪",
		"Corning Gorilla Glass Victus 2": "Corning Gorilla Glass Victus ২",
		"Glass front/back, aluminum frame": "গ্লাস সামনে/back, অ্যালুমিনিয়াম ফ্রেম",
		"IP64": "IP৬৪",
		"Iron Gray, Cool Blue, Electric Violet": "Iron ধূসর, Cool নীল, Electric Violet",
		"January 2024": "January ২০২৪",
		"LTPO4 AMOLED, 120Hz, HDR10+, 4500 nits": "LTPO৪ AMOLED, ১২০Hz, HDR১০+, ৪৫০০ nits",
		"Qualcomm SM8550-AB Snapdragon 8 Gen 2 (4 nm)": "Qualcomm SM৮৫৫০-AB Snapdragon ৮ Gen ২ (৪ nm)",
		"Snapdragon 8 Gen 2": "Snapdragon ৮ Gen ২",
	}
}

// Seed inserts specification records for the product identified by slug 'oneplus-12r'
func (s *SpecificationSeederMobileOneplus12r) Seed(db *gorm.DB) error {
	productSlug := "oneplus-12r"

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

	// Override model-specific values for OnePlus 12R
	specs["Display Size"] = "6.78 inches"
	specs["Processor"] = "Snapdragon 8 Gen 2"
	specs["Chipset"] = "Qualcomm SM8550-AB Snapdragon 8 Gen 2 (4 nm)"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "Adreno 740"
	specs["Ram"] = "8 GB / 16 GB"
	specs["Storage"] = "128 GB / 256 GB"
	specs["Display Type"] = "LTPO4 AMOLED, 120Hz, HDR10+, 4500 nits"
	specs["Resolution"] = "1264 x 2780 pixels"
	specs["Screen Protection"] = "Corning Gorilla Glass Victus 2"
	specs["Refresh Rate"] = "120Hz"
	specs["Build Material"] = "Glass front/back, aluminum frame"
	specs["Weight"] = "207 g"
	specs["Dimensions"] = "163.3 x 75.3 x 8.8 mm"
	specs["Water Resistance"] = "IP64"
	specs["Network Technology"] = "5G"
	specs["Rear Camera"] = "50 MP + 8 MP + 2 MP"
	specs["Front Camera"] = "16 MP"
	specs["Battery"] = "5,500 mAh"
	specs["Operating System"] = "Android 14, OxygenOS 14"
	specs["Available Colors"] = "Iron Gray, Cool Blue, Electric Violet"
	specs["Announcement Date"] = "January 2024"
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
