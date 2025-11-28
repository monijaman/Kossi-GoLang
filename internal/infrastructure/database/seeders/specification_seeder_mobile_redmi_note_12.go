package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileRedmiNote12 seeds specifications/options for product 'redmi-note-12'
type SpecificationSeederMobileRedmiNote12 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileRedmiNote12 creates a new seeder instance
func NewSpecificationSeederMobileRedmiNote12() *SpecificationSeederMobileRedmiNote12 {
	return &SpecificationSeederMobileRedmiNote12{BaseSeeder: BaseSeeder{name: "Specifications for Redmi Note 12"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileRedmiNote12) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080 x 2400 pixels": "১০৮০ x ২৪০০ পিক্সেল",
		"120Hz": "১২০Hz",
		"128 GB / 256 GB": "১২৮ জিবি / ২৫৬ GB",
		"13 MP": "১৩ মেগাপিক্সেল",
		"165.9 x 76.2 x 8 mm": "১৬৫.৯ x ৭৬.২ x ৮ মিমি",
		"188 g": "১৮৮ g",
		"4 GB / 6 GB / 8 GB": "৪ জিবি / ৬ জিবি / ৮ GB",
		"48 MP + 8 MP + 2 MP": "৪৮ মেগাপিক্সেল + ৮ মেগাপিক্সেল + ২ মেগাপিক্সেল",
		"5,000 mAh": "৫,০০০ এমএএইচ",
		"5G": "৫G",
		"6.67 inches": "৬.৬৭ ইঞ্চি",
		"AMOLED, 120Hz, 1200 nits": "অ্যামোলেড, ১২০Hz, ১২০০ nits",
		"Adreno 619": "অ্যাড্রেনো ৬১৯",
		"Android 12, upgradable": "অ্যান্ড্রয়েড ১২, আপগ্রেডযোগ্য",
		"Corning Gorilla Glass 3": "কর্নিং গরিলা গ্লাস ৩",
		"Frosted Green, Matte Black, Mystique Blue": "ফ্রস্টed সবুজ, Matte কালো, Mystique নীল",
		"Glass front, plastic frame, plastic back": "গ্লাস সামনে, প্লাস্টিক ফ্রেম, প্লাস্টিক পেছনে",
		"IP53": "IP৫৩",
		"October 2022": "অক্টোবর ২০২২",
		"Qualcomm SM4375 Snapdragon 4 Gen 1 (6 nm)": "কোয়ালকম SM৪৩৭৫ স্ন্যাপড্রাগন ৪ জেন ১ (৬ ন্যানোমিটার)",
		"Snapdragon 4 Gen 1": "স্ন্যাপড্রাগন ৪ জেন ১",
	}
}

// Seed inserts specification records for the product identified by slug 'redmi-note-12'
func (s *SpecificationSeederMobileRedmiNote12) Seed(db *gorm.DB) error {
	productSlug := "redmi-note-12"

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

	// Override model-specific values for Redmi Note 12
	specs["Display Size"] = "6.67 inches"
	specs["Processor"] = "Snapdragon 4 Gen 1"
	specs["Chipset"] = "Qualcomm SM4375 Snapdragon 4 Gen 1 (6 nm)"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "Adreno 619"
	specs["Ram"] = "4 GB / 6 GB / 8 GB"
	specs["Storage"] = "128 GB / 256 GB"
	specs["Display Type"] = "AMOLED, 120Hz, 1200 nits"
	specs["Resolution"] = "1080 x 2400 pixels"
	specs["Screen Protection"] = "Corning Gorilla Glass 3"
	specs["Refresh Rate"] = "120Hz"
	specs["Build Material"] = "Glass front, plastic frame, plastic back"
	specs["Weight"] = "188 g"
	specs["Dimensions"] = "165.9 x 76.2 x 8 mm"
	specs["Water Resistance"] = "IP53"
	specs["Network Technology"] = "5G"
	specs["Rear Camera"] = "48 MP + 8 MP + 2 MP"
	specs["Front Camera"] = "13 MP"
	specs["Battery"] = "5,000 mAh"
	specs["Operating System"] = "Android 12, upgradable"
	specs["Available Colors"] = "Frosted Green, Matte Black, Mystique Blue"
	specs["Announcement Date"] = "October 2022"
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
