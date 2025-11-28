package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileTecnoSparkGo1 seeds specifications/options for product 'tecno-spark-go-1'
type SpecificationSeederMobileTecnoSparkGo1 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileTecnoSparkGo1 creates a new seeder instance
func NewSpecificationSeederMobileTecnoSparkGo1() *SpecificationSeederMobileTecnoSparkGo1 {
	return &SpecificationSeederMobileTecnoSparkGo1{BaseSeeder: BaseSeeder{name: "Specifications for Tecno SPARK Go 1"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileTecnoSparkGo1) getBanglaTranslations() map[string]string {
	return map[string]string{
		"120Hz": "১২০Hz",
		"13 MP": "১৩ মেগাপিক্সেল",
		"165.6 x 76.3 x 8.4 mm": "১৬৫.৬ x ৭৬.৩ x ৮.৪ মিমি",
		"180 g": "১৮০ g",
		"3 GB / 4 GB": "৩ জিবি / ৪ GB",
		"4G": "৪G",
		"5,000 mAh": "৫,০০০ এমএএইচ",
		"6.67 inches": "৬.৬৭ ইঞ্চি",
		"64 GB / 128 GB": "৬৪ জিবি / ১২৮ GB",
		"720 x 1600 pixels": "৭২০ x ১৬০০ পিক্সেল",
		"8 MP": "৮ মেগাপিক্সেল",
		"Android 14 Go Edition": "অ্যান্ড্রয়েড ১৪ Go Edition",
		"August 2024": "আগস্ট ২০২৪",
		"IP54": "IP৫৪",
		"IPS LCD, 120Hz": "আইপিএস এলসিডি, ১২০Hz",
		"Mali-G57": "মালি-G৫৭",
		"Startrail Black, Glittery White": "Startrail কালো, Glittery সাদা",
		"Unisoc T615": "ইউনিসক T৬১৫",
	}
}

// Seed inserts specification records for the product identified by slug 'tecno-spark-go-1'
func (s *SpecificationSeederMobileTecnoSparkGo1) Seed(db *gorm.DB) error {
	productSlug := "tecno-spark-go-1"

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

	// Override model-specific values for Tecno SPARK Go 1
	specs["Display Size"] = "6.67 inches"
	specs["Processor"] = "Unisoc T615"
	specs["Chipset"] = "Unisoc T615"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "Mali-G57"
	specs["Ram"] = "3 GB / 4 GB"
	specs["Storage"] = "64 GB / 128 GB"
	specs["Display Type"] = "IPS LCD, 120Hz"
	specs["Resolution"] = "720 x 1600 pixels"
	specs["Screen Protection"] = "Glass front"
	specs["Refresh Rate"] = "120Hz"
	specs["Build Material"] = "Plastic body"
	specs["Weight"] = "180 g"
	specs["Dimensions"] = "165.6 x 76.3 x 8.4 mm"
	specs["Water Resistance"] = "IP54"
	specs["Network Technology"] = "4G"
	specs["Rear Camera"] = "13 MP"
	specs["Front Camera"] = "8 MP"
	specs["Battery"] = "5,000 mAh"
	specs["Operating System"] = "Android 14 Go Edition"
	specs["Available Colors"] = "Startrail Black, Glittery White"
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
