package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileTecnoSpark405g seeds specifications/options for product 'tecno-spark-40-5g'
type SpecificationSeederMobileTecnoSpark405g struct {
	BaseSeeder
}

// NewSpecificationSeederMobileTecnoSpark405g creates a new seeder instance
func NewSpecificationSeederMobileTecnoSpark405g() *SpecificationSeederMobileTecnoSpark405g {
	return &SpecificationSeederMobileTecnoSpark405g{BaseSeeder: BaseSeeder{name: "Specifications for Tecno SPARK 40 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileTecnoSpark405g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"120Hz": "১২০Hz",
		"163.7 x 75.6 x 8.5 mm": "১৬৩.৭ x ৭৫.৬ x ৮.৫ মিমি",
		"190 g": "১৯০ g",
		"256 GB": "২৫৬ GB",
		"5,000 mAh": "৫,০০০ এমএএইচ",
		"50 MP + AI": "৫০ মেগাপিক্সেল + এআই",
		"5G": "৫G",
		"6.6 inches": "৬.৬ ইঞ্চি",
		"720 x 1612 pixels": "৭২০ x ১৬১২ পিক্সেল",
		"8 GB": "৮ GB",
		"8 MP": "৮ মেগাপিক্সেল",
		"Android 14, HIOS 14": "অ্যান্ড্রয়েড ১৪, HIOS ১৪",
		"Dimensity 6080": "ডাইমেনসিটি ৬০৮০",
		"Gravity Black, Cyber White": "গ্র্যাভিটি কালো, সাইবার সাদা",
		"IPS LCD, 120Hz": "আইপিএস এলসিডি, ১২০Hz",
		"January 2024": "জানুয়ারি ২০২৪",
		"Mali-G57 MC2": "মালি-G৫৭ MC২",
		"Mediatek Dimensity 6080 (6 nm)": "মিডিয়াটেক ডাইমেনসিটি ৬০৮০ (৬ ন্যানোমিটার)",
	}
}

// Seed inserts specification records for the product identified by slug 'tecno-spark-40-5g'
func (s *SpecificationSeederMobileTecnoSpark405g) Seed(db *gorm.DB) error {
	productSlug := "tecno-spark-40-5g"

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

	// Override model-specific values for Tecno SPARK 40 5G
	specs["Display Size"] = "6.6 inches"
	specs["Processor"] = "Dimensity 6080"
	specs["Chipset"] = "Mediatek Dimensity 6080 (6 nm)"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "Mali-G57 MC2"
	specs["Ram"] = "8 GB"
	specs["Storage"] = "256 GB"
	specs["Display Type"] = "IPS LCD, 120Hz"
	specs["Resolution"] = "720 x 1612 pixels"
	specs["Screen Protection"] = "Glass front"
	specs["Refresh Rate"] = "120Hz"
	specs["Build Material"] = "Plastic body"
	specs["Weight"] = "190 g"
	specs["Dimensions"] = "163.7 x 75.6 x 8.5 mm"
	specs["Water Resistance"] = "No"
	specs["Network Technology"] = "5G"
	specs["Rear Camera"] = "50 MP + AI"
	specs["Front Camera"] = "8 MP"
	specs["Battery"] = "5,000 mAh"
	specs["Operating System"] = "Android 14, HIOS 14"
	specs["Available Colors"] = "Gravity Black, Cyber White"
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
