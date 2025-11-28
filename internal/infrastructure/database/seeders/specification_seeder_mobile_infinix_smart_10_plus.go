package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileInfinixSmart10Plus seeds specifications/options for product 'infinix-smart-10-plus'
type SpecificationSeederMobileInfinixSmart10Plus struct {
	BaseSeeder
}

// NewSpecificationSeederMobileInfinixSmart10Plus creates a new seeder instance
func NewSpecificationSeederMobileInfinixSmart10Plus() *SpecificationSeederMobileInfinixSmart10Plus {
	return &SpecificationSeederMobileInfinixSmart10Plus{BaseSeeder: BaseSeeder{name: "Specifications for Infinix Smart 10 Plus"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileInfinixSmart10Plus) getBanglaTranslations() map[string]string {
	return map[string]string{
		"128 GB": "১২৮ GB",
		"164 x 76 x 8.5 mm": "১৬৪ x ৭৬ x ৮.৫ মিমি",
		"190 g": "১৯০ g",
		"4 GB": "৪ GB",
		"4G": "৪G",
		"50 MP + AI": "৫০ মেগাপিক্সেল + এআই",
		"6,000 mAh": "৬,০০০ এমএএইচ",
		"6.6 inches": "৬.৬ ইঞ্চি",
		"720 x 1612 pixels": "৭২০ x ১৬১২ পিক্সেল",
		"8 MP": "৮ মেগাপিক্সেল",
		"90Hz": "৯০Hz",
		"Android 13 Go Edition": "অ্যান্ড্রয়েড ১৩ Go Edition",
		"Black, Green": "কালো, সবুজ",
		"IPS LCD, 90Hz": "আইপিএস এলসিডি, ৯০Hz",
		"January 2024": "জানুয়ারি ২০২৪",
		"Mali-G57 MP1": "মালি-G৫৭ মেগাপিক্সেল১",
		"Unisoc T606": "ইউনিসক T৬০৬",
		"Unisoc T606 (12 nm)": "ইউনিসক T৬০৬ (১২ ন্যানোমিটার)",
	}
}

// Seed inserts specification records for the product identified by slug 'infinix-smart-10-plus'
func (s *SpecificationSeederMobileInfinixSmart10Plus) Seed(db *gorm.DB) error {
	productSlug := "infinix-smart-10-plus"

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

	// Override model-specific values for Infinix Smart 10 Plus
	specs["Display Size"] = "6.6 inches"
	specs["Processor"] = "Unisoc T606"
	specs["Chipset"] = "Unisoc T606 (12 nm)"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "Mali-G57 MP1"
	specs["Ram"] = "4 GB"
	specs["Storage"] = "128 GB"
	specs["Display Type"] = "IPS LCD, 90Hz"
	specs["Resolution"] = "720 x 1612 pixels"
	specs["Screen Protection"] = "Glass front"
	specs["Refresh Rate"] = "90Hz"
	specs["Build Material"] = "Plastic body"
	specs["Weight"] = "190 g"
	specs["Dimensions"] = "164 x 76 x 8.5 mm"
	specs["Water Resistance"] = "No"
	specs["Network Technology"] = "4G"
	specs["Rear Camera"] = "50 MP + AI"
	specs["Front Camera"] = "8 MP"
	specs["Battery"] = "6,000 mAh"
	specs["Operating System"] = "Android 13 Go Edition"
	specs["Available Colors"] = "Black, Green"
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
