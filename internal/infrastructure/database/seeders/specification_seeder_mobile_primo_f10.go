package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobilePrimoF10 seeds specifications/options for product 'primo-f10'
type SpecificationSeederMobilePrimoF10 struct {
	BaseSeeder
}

// NewSpecificationSeederMobilePrimoF10 creates a new seeder instance
func NewSpecificationSeederMobilePrimoF10() *SpecificationSeederMobilePrimoF10 {
	return &SpecificationSeederMobilePrimoF10{BaseSeeder: BaseSeeder{name: "Specifications for Primo F10"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobilePrimoF10) getBanglaTranslations() map[string]string {
	return map[string]string{
		"160 x 77 x 9.5 mm": "১৬০ x ৭৭ x ৯.৫ মিমি",
		"170 g": "১৭০ g",
		"2 GB": "২ GB",
		"3,000 mAh": "৩,০০০ এমএএইচ",
		"32 GB": "৩২ GB",
		"480 x 960 pixels": "৪৮০ x ৯৬০ পিক্সেল",
		"4G": "৪G",
		"5 MP": "৫ মেগাপিক্সেল",
		"6.0 inches": "৬.০ ইঞ্চি",
		"60Hz": "৬০Hz",
		"8 MP": "৮ মেগাপিক্সেল",
		"Android 11 Go Edition": "অ্যান্ড্রয়েড ১১ Go Edition",
		"Blue, Black": "নীল, কালো",
		"IMG8322": "IMG৮৩২২",
		"July 2021": "জুলাই ২০২১",
		"Unisoc SC9863A": "ইউনিসক SC৯৮৬৩A",
		"Unisoc SC9863A (28 nm)": "ইউনিসক SC৯৮৬৩A (২৮ ন্যানোমিটার)",
	}
}

// Seed inserts specification records for the product identified by slug 'primo-f10'
func (s *SpecificationSeederMobilePrimoF10) Seed(db *gorm.DB) error {
	productSlug := "primo-f10"

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

	// Override model-specific values for Primo F10
	specs["Display Size"] = "6.0 inches"
	specs["Processor"] = "Unisoc SC9863A"
	specs["Chipset"] = "Unisoc SC9863A (28 nm)"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "IMG8322"
	specs["Ram"] = "2 GB"
	specs["Storage"] = "32 GB"
	specs["Display Type"] = "IPS LCD"
	specs["Resolution"] = "480 x 960 pixels"
	specs["Screen Protection"] = "Glass front"
	specs["Refresh Rate"] = "60Hz"
	specs["Build Material"] = "Plastic body"
	specs["Weight"] = "170 g"
	specs["Dimensions"] = "160 x 77 x 9.5 mm"
	specs["Water Resistance"] = "No"
	specs["Network Technology"] = "4G"
	specs["Rear Camera"] = "8 MP"
	specs["Front Camera"] = "5 MP"
	specs["Battery"] = "3,000 mAh"
	specs["Operating System"] = "Android 11 Go Edition"
	specs["Available Colors"] = "Blue, Black"
	specs["Announcement Date"] = "July 2021"
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
