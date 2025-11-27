package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileRealme12 seeds specifications/options for product 'realme-12'
type SpecificationSeederMobileRealme12 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileRealme12 creates a new seeder instance
func NewSpecificationSeederMobileRealme12() *SpecificationSeederMobileRealme12 {
	return &SpecificationSeederMobileRealme12{BaseSeeder: BaseSeeder{name: "Specifications for Realme 12"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileRealme12) getBanglaTranslations() map[string]string {
	return map[string]string{
		"108 MP + 2 MP": "১০৮ MP + ২ MP",
		"1080 x 2400 pixels": "১০৮০ x ২৪০০ pixels",
		"120Hz": "১২০Hz",
		"128 GB": "১২৮ GB",
		"188 g": "১৮৮ g",
		"45W wired": "৪৫W তারযুক্ত",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6 GB / 8 GB": "৬ GB / ৮ GB",
		"6.72 inches": "৬.৭২ ইঞ্চি",
		"8 MP": "৮ MP",
		"Android 14, Realme UI 5.0": "Android ১৪, Realme UI ৫.০",
		"Dimensity 6100+": "Dimensity ৬১০০+",
		"GSM / HSPA / LTE / 5G": "GSM / HSPA / LTE / ৫G",
		"Glass front, plastic back, plastic frame": "গ্লাস সামনে, প্লাস্টিক পেছনে, plastic frame",
		"IP54": "IP৫৪",
		"IPS LCD, 120Hz, 950 nits": "IPS LCD, ১২০Hz, ৯৫০ nits",
		"Mali-G57 MC2": "Mali-G৫৭ MC২",
		"March 2024": "March ২০২৪",
		"Mediatek Dimensity 6100+ (6 nm)": "Mediatek Dimensity ৬১০০+ (৬ nm)",
		"Twilight Purple, Woodland Green": "Twilight বেগুনি, Woodland সবুজ",
	}
}

// Seed inserts specification records for the product identified by slug 'realme-12'
func (s *SpecificationSeederMobileRealme12) Seed(db *gorm.DB) error {
	productSlug := "realme-12"

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

	// Override model-specific values for Realme 12
	specs["Display Size"] = "6.72 inches"
	specs["Processor"] = "Dimensity 6100+"
	specs["Chipset"] = "Mediatek Dimensity 6100+ (6 nm)"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "Mali-G57 MC2"
	specs["Ram"] = "6 GB / 8 GB"
	specs["Storage"] = "128 GB"
	specs["Display Type"] = "IPS LCD, 120Hz, 950 nits"
	specs["Resolution"] = "1080 x 2400 pixels"
	specs["Refresh Rate"] = "120Hz"
	specs["Build Material"] = "Glass front, plastic back, plastic frame"
	specs["Weight"] = "188 g"
	specs["Water Resistance"] = "IP54"
	specs["Network Technology"] = "GSM / HSPA / LTE / 5G"
	specs["Rear Camera"] = "108 MP + 2 MP"
	specs["Front Camera"] = "8 MP"
	specs["Battery"] = "5000 mAh"
	specs["Fast Charging"] = "45W wired"
	specs["Operating System"] = "Android 14, Realme UI 5.0"
	specs["Available Colors"] = "Twilight Purple, Woodland Green"
	specs["Announcement Date"] = "March 2024"
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
