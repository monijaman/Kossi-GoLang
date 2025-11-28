package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileRedmiNote124g seeds specifications/options for product 'redmi-note-12-4g'
type SpecificationSeederMobileRedmiNote124g struct {
	BaseSeeder
}

// NewSpecificationSeederMobileRedmiNote124g creates a new seeder instance
func NewSpecificationSeederMobileRedmiNote124g() *SpecificationSeederMobileRedmiNote124g {
	return &SpecificationSeederMobileRedmiNote124g{BaseSeeder: BaseSeeder{name: "Specifications for Redmi Note 12 4G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileRedmiNote124g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080 x 2400 pixels": "১০৮০ x ২৪০০ পিক্সেল",
		"120Hz": "১২০Hz",
		"13 MP": "১৩ মেগাপিক্সেল",
		"183.5 g": "১৮৩.৫ g",
		"33W wired": "৩৩W তারযুক্ত",
		"4 GB / 6 GB / 8 GB": "৪ জিবি / ৬ জিবি / ৮ GB",
		"50 MP + 8 MP + 2 MP": "৫০ মেগাপিক্সেল + ৮ মেগাপিক্সেল + ২ মেগাপিক্সেল",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6.67 inches": "৬.৬৭ ইঞ্চি",
		"64 GB / 128 GB": "৬৪ জিবি / ১২৮ GB",
		"AMOLED, 120Hz": "অ্যামোলেড, ১২০Hz",
		"Adreno 610": "অ্যাড্রেনো ৬১০",
		"Android 13, MIUI 14": "অ্যান্ড্রয়েড ১৩, এমআইইউআই ১৪",
		"Glass front (Gorilla Glass 3), plastic back, plastic frame": "গ্লাস সামনে (গরিলা গ্লাস ৩), প্লাস্টিক পেছনে, প্লাস্টিক ফ্রেম",
		"IP53": "IP৫৩",
		"March 2023": "মার্চ ২০২৩",
		"Onyx Gray, Mint Green, Ice Blue": "অনিক্স ধূসর, মিন্ট সবুজ, আইস নীল",
		"Qualcomm Snapdragon 685 (6 nm)": "কোয়ালকম স্ন্যাপড্রাগন ৬৮৫ (৬ ন্যানোমিটার)",
		"Snapdragon 685": "স্ন্যাপড্রাগন ৬৮৫",
	}
}

// Seed inserts specification records for the product identified by slug 'redmi-note-12-4g'
func (s *SpecificationSeederMobileRedmiNote124g) Seed(db *gorm.DB) error {
	productSlug := "redmi-note-12-4g"

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

	// Override model-specific values for Redmi Note 12 4G
	specs["Display Size"] = "6.67 inches"
	specs["Processor"] = "Snapdragon 685"
	specs["Chipset"] = "Qualcomm Snapdragon 685 (6 nm)"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "Adreno 610"
	specs["Ram"] = "4 GB / 6 GB / 8 GB"
	specs["Storage"] = "64 GB / 128 GB"
	specs["Display Type"] = "AMOLED, 120Hz"
	specs["Resolution"] = "1080 x 2400 pixels"
	specs["Refresh Rate"] = "120Hz"
	specs["Build Material"] = "Glass front (Gorilla Glass 3), plastic back, plastic frame"
	specs["Weight"] = "183.5 g"
	specs["Water Resistance"] = "IP53"
	specs["Network Technology"] = "GSM / HSPA / LTE"
	specs["Rear Camera"] = "50 MP + 8 MP + 2 MP"
	specs["Front Camera"] = "13 MP"
	specs["Battery"] = "5000 mAh"
	specs["Fast Charging"] = "33W wired"
	specs["Operating System"] = "Android 13, MIUI 14"
	specs["Available Colors"] = "Onyx Gray, Mint Green, Ice Blue"
	specs["Announcement Date"] = "March 2023"
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
