package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileSymphonyAtom5 seeds specifications/options for product 'symphony-atom-5'
type SpecificationSeederMobileSymphonyAtom5 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSymphonyAtom5 creates a new seeder instance
func NewSpecificationSeederMobileSymphonyAtom5() *SpecificationSeederMobileSymphonyAtom5 {
	return &SpecificationSeederMobileSymphonyAtom5{BaseSeeder: BaseSeeder{name: "Specifications for Symphony Atom 5"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSymphonyAtom5) getBanglaTranslations() map[string]string {
	return map[string]string{
		"164 x 75.8 x 8.9 mm": "১৬৪ x ৭৫.৮ x ৮.৯ মিমি",
		"190 g": "১৯০ g",
		"4 GB": "৪ GB",
		"4G": "৪G",
		"5,000 mAh": "৫,০০০ এমএএইচ",
		"50 MP + 2 MP": "৫০ মেগাপিক্সেল + ২ মেগাপিক্সেল",
		"6.5 inches": "৬.৫ ইঞ্চি",
		"60Hz": "৬০Hz",
		"64 GB": "৬৪ GB",
		"720 x 1600 pixels": "৭২০ x ১৬০০ পিক্সেল",
		"8 MP": "৮ মেগাপিক্সেল",
		"Android 13": "অ্যান্ড্রয়েড ১৩",
		"December 2023": "ডিসেম্বর ২০২৩",
		"Green, Black": "সবুজ, কালো",
		"Helio G37": "হেলিও G৩৭",
		"Mediatek Helio G37 (12 nm)": "মিডিয়াটেক হেলিও G৩৭ (১২ ন্যানোমিটার)",
		"PowerVR GE8320": "পাওয়ারভিআর GE৮৩২০",
	}
}

// Seed inserts specification records for the product identified by slug 'symphony-atom-5'
func (s *SpecificationSeederMobileSymphonyAtom5) Seed(db *gorm.DB) error {
	productSlug := "symphony-atom-5"

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

	// Override model-specific values for Symphony Atom 5
	specs["Display Size"] = "6.5 inches"
	specs["Processor"] = "Helio G37"
	specs["Chipset"] = "Mediatek Helio G37 (12 nm)"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "PowerVR GE8320"
	specs["Ram"] = "4 GB"
	specs["Storage"] = "64 GB"
	specs["Display Type"] = "IPS LCD"
	specs["Resolution"] = "720 x 1600 pixels"
	specs["Screen Protection"] = "Glass front"
	specs["Refresh Rate"] = "60Hz"
	specs["Build Material"] = "Plastic body"
	specs["Weight"] = "190 g"
	specs["Dimensions"] = "164 x 75.8 x 8.9 mm"
	specs["Water Resistance"] = "No"
	specs["Network Technology"] = "4G"
	specs["Rear Camera"] = "50 MP + 2 MP"
	specs["Front Camera"] = "8 MP"
	specs["Battery"] = "5,000 mAh"
	specs["Operating System"] = "Android 13"
	specs["Available Colors"] = "Green, Black"
	specs["Announcement Date"] = "December 2023"
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
