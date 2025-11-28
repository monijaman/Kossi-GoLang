package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobilePocoX6Neo5g seeds specifications/options for product 'poco-x6-neo-5g'
type SpecificationSeederMobilePocoX6Neo5g struct {
	BaseSeeder
}

// NewSpecificationSeederMobilePocoX6Neo5g creates a new seeder instance
func NewSpecificationSeederMobilePocoX6Neo5g() *SpecificationSeederMobilePocoX6Neo5g {
	return &SpecificationSeederMobilePocoX6Neo5g{BaseSeeder: BaseSeeder{name: "Specifications for POCO X6 Neo 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobilePocoX6Neo5g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"108 MP + 2 MP": "১০৮ মেগাপিক্সেল + ২ মেগাপিক্সেল",
		"1080 x 2400 pixels": "১০৮০ x ২৪০০ পিক্সেল",
		"120Hz": "১২০Hz",
		"128 GB / 256 GB": "১২৮ জিবি / ২৫৬ GB",
		"16 MP": "১৬ মেগাপিক্সেল",
		"175 g": "১৭৫ g",
		"33W wired": "৩৩W তারযুক্ত",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6.67 inches": "৬.৬৭ ইঞ্চি",
		"8 GB / 12 GB": "৮ জিবি / ১২ GB",
		"AMOLED, 120Hz": "অ্যামোলেড, ১২০Hz",
		"Android 13, MIUI 14": "অ্যান্ড্রয়েড ১৩, এমআইইউআই ১৪",
		"Astral Black, Horizon Blue, Martian Orange": "অ্যাস্ট্রাল কালো, হরাইজন নীল, মার্শিয়ান Orange",
		"Dimensity 6080": "ডাইমেনসিটি ৬০৮০",
		"GSM / HSPA / LTE / 5G": "জিএসএম / এইচএসপিএ / এলটিই / ৫G",
		"Glass front (Gorilla Glass 5), plastic back, plastic frame": "গ্লাস সামনে (গরিলা গ্লাস ৫), প্লাস্টিক পেছনে, প্লাস্টিক ফ্রেম",
		"IP54": "IP৫৪",
		"Mali-G57 MC2": "মালি-G৫৭ MC২",
		"March 2024": "মার্চ ২০২৪",
		"Mediatek Dimensity 6080 (6 nm)": "মিডিয়াটেক ডাইমেনসিটি ৬০৮০ (৬ ন্যানোমিটার)",
	}
}

// Seed inserts specification records for the product identified by slug 'poco-x6-neo-5g'
func (s *SpecificationSeederMobilePocoX6Neo5g) Seed(db *gorm.DB) error {
	productSlug := "poco-x6-neo-5g"

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

	// Override model-specific values for POCO X6 Neo 5G
	specs["Display Size"] = "6.67 inches"
	specs["Processor"] = "Dimensity 6080"
	specs["Chipset"] = "Mediatek Dimensity 6080 (6 nm)"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "Mali-G57 MC2"
	specs["Ram"] = "8 GB / 12 GB"
	specs["Storage"] = "128 GB / 256 GB"
	specs["Display Type"] = "AMOLED, 120Hz"
	specs["Resolution"] = "1080 x 2400 pixels"
	specs["Refresh Rate"] = "120Hz"
	specs["Build Material"] = "Glass front (Gorilla Glass 5), plastic back, plastic frame"
	specs["Weight"] = "175 g"
	specs["Water Resistance"] = "IP54"
	specs["Network Technology"] = "GSM / HSPA / LTE / 5G"
	specs["Rear Camera"] = "108 MP + 2 MP"
	specs["Front Camera"] = "16 MP"
	specs["Battery"] = "5000 mAh"
	specs["Fast Charging"] = "33W wired"
	specs["Operating System"] = "Android 13, MIUI 14"
	specs["Available Colors"] = "Astral Black, Horizon Blue, Martian Orange"
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
