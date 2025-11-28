package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileTecnoPovaCurve5g seeds specifications/options for product 'tecno-pova-curve-5g'
type SpecificationSeederMobileTecnoPovaCurve5g struct {
	BaseSeeder
}

// NewSpecificationSeederMobileTecnoPovaCurve5g creates a new seeder instance
func NewSpecificationSeederMobileTecnoPovaCurve5g() *SpecificationSeederMobileTecnoPovaCurve5g {
	return &SpecificationSeederMobileTecnoPovaCurve5g{BaseSeeder: BaseSeeder{name: "Specifications for Tecno POVA Curve 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileTecnoPovaCurve5g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080 x 2400 pixels": "১০৮০ x ২৪০০ পিক্সেল",
		"120Hz": "১২০Hz",
		"16 MP": "১৬ মেগাপিক্সেল",
		"162 x 75 x 7.9 mm": "১৬২ x ৭৫ x ৭.৯ মিমি",
		"185 g": "১৮৫ g",
		"2024": "২০২৪",
		"256 GB": "২৫৬ GB",
		"5,000 mAh": "৫,০০০ এমএএইচ",
		"50 MP + 2 MP": "৫০ মেগাপিক্সেল + ২ মেগাপিক্সেল",
		"5G": "৫G",
		"6.67 inches": "৬.৬৭ ইঞ্চি",
		"8 GB": "৮ GB",
		"AMOLED, 120Hz, Curved": "অ্যামোলেড, ১২০Hz, Curved",
		"Android 14": "অ্যান্ড্রয়েড ১৪",
		"Black, Blue": "কালো, নীল",
		"Dimensity 7050": "ডাইমেনসিটি ৭০৫০",
		"Mali-G68 MC4": "মালি-G৬৮ MC৪",
		"Mediatek Dimensity 7050 (6 nm)": "মিডিয়াটেক ডাইমেনসিটি ৭০৫০ (৬ ন্যানোমিটার)",
	}
}

// Seed inserts specification records for the product identified by slug 'tecno-pova-curve-5g'
func (s *SpecificationSeederMobileTecnoPovaCurve5g) Seed(db *gorm.DB) error {
	productSlug := "tecno-pova-curve-5g"

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

	// Override model-specific values for Tecno POVA Curve 5G
	specs["Display Size"] = "6.67 inches"
	specs["Processor"] = "Dimensity 7050"
	specs["Chipset"] = "Mediatek Dimensity 7050 (6 nm)"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "Mali-G68 MC4"
	specs["Ram"] = "8 GB"
	specs["Storage"] = "256 GB"
	specs["Display Type"] = "AMOLED, 120Hz, Curved"
	specs["Resolution"] = "1080 x 2400 pixels"
	specs["Screen Protection"] = "Glass front"
	specs["Refresh Rate"] = "120Hz"
	specs["Build Material"] = "Plastic body"
	specs["Weight"] = "185 g"
	specs["Dimensions"] = "162 x 75 x 7.9 mm"
	specs["Water Resistance"] = "No"
	specs["Network Technology"] = "5G"
	specs["Rear Camera"] = "50 MP + 2 MP"
	specs["Front Camera"] = "16 MP"
	specs["Battery"] = "5,000 mAh"
	specs["Operating System"] = "Android 14"
	specs["Available Colors"] = "Black, Blue"
	specs["Announcement Date"] = "2024"
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
