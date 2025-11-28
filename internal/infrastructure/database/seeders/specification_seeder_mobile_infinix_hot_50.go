package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileInfinixHot50 seeds specifications/options for product 'infinix-hot-50'
type SpecificationSeederMobileInfinixHot50 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileInfinixHot50 creates a new seeder instance
func NewSpecificationSeederMobileInfinixHot50() *SpecificationSeederMobileInfinixHot50 {
	return &SpecificationSeederMobileInfinixHot50{BaseSeeder: BaseSeeder{name: "Specifications for Infinix Hot 50"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileInfinixHot50) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080 x 2460 pixels": "১০৮০ x ২৪৬০ পিক্সেল",
		"128 GB": "১২৮ GB",
		"168.7 x 76.6 x 8.4 mm": "১৬৮.৭ x ৭৬.৬ x ৮.৪ মিমি",
		"196 g": "১৯৬ g",
		"4G": "৪G",
		"5,000 mAh": "৫,০০০ এমএএইচ",
		"50 MP + 2 MP": "৫০ মেগাপিক্সেল + ২ মেগাপিক্সেল",
		"6.78 inches": "৬.৭৮ ইঞ্চি",
		"8 GB": "৮ GB",
		"8 MP": "৮ মেগাপিক্সেল",
		"90Hz": "৯০Hz",
		"Android 13, XOS 13.5": "অ্যান্ড্রয়েড ১৩, এক্সওএস ১৩.৫",
		"August 2024": "আগস্ট ২০২৪",
		"Helio G88": "হেলিও G৮৮",
		"IPS LCD, 90Hz": "আইপিএস এলসিডি, ৯০Hz",
		"Mali-G52 MC2": "মালি-G৫২ MC২",
		"Mediatek Helio G88 (12 nm)": "মিডিয়াটেক হেলিও G৮৮ (১২ ন্যানোমিটার)",
		"Racing Black, Surfing Green": "Racing কালো, Surfing সবুজ",
	}
}

// Seed inserts specification records for the product identified by slug 'infinix-hot-50'
func (s *SpecificationSeederMobileInfinixHot50) Seed(db *gorm.DB) error {
	productSlug := "infinix-hot-50"

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

	// Override model-specific values for Infinix Hot 50
	specs["Display Size"] = "6.78 inches"
	specs["Processor"] = "Helio G88"
	specs["Chipset"] = "Mediatek Helio G88 (12 nm)"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "Mali-G52 MC2"
	specs["Ram"] = "8 GB"
	specs["Storage"] = "128 GB"
	specs["Display Type"] = "IPS LCD, 90Hz"
	specs["Resolution"] = "1080 x 2460 pixels"
	specs["Screen Protection"] = "Glass front"
	specs["Refresh Rate"] = "90Hz"
	specs["Build Material"] = "Plastic body"
	specs["Weight"] = "196 g"
	specs["Dimensions"] = "168.7 x 76.6 x 8.4 mm"
	specs["Water Resistance"] = "No"
	specs["Network Technology"] = "4G"
	specs["Rear Camera"] = "50 MP + 2 MP"
	specs["Front Camera"] = "8 MP"
	specs["Battery"] = "5,000 mAh"
	specs["Operating System"] = "Android 13, XOS 13.5"
	specs["Available Colors"] = "Racing Black, Surfing Green"
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
