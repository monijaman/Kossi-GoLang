package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobilePocoX75g seeds specifications/options for product 'poco-x7-5g'
type SpecificationSeederMobilePocoX75g struct {
	BaseSeeder
}

// NewSpecificationSeederMobilePocoX75g creates a new seeder instance
func NewSpecificationSeederMobilePocoX75g() *SpecificationSeederMobilePocoX75g {
	return &SpecificationSeederMobilePocoX75g{BaseSeeder: BaseSeeder{name: "Specifications for POCO X7 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobilePocoX75g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080 x 2400 pixels": "১০৮০ x ২৪০০ পিক্সেল",
		"120Hz": "১২০Hz",
		"16 MP": "১৬ মেগাপিক্সেল",
		"189 g": "১৮৯ g",
		"256 GB / 512 GB": "২৫৬ জিবি / ৫১২ GB",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6.67 inches": "৬.৬৭ ইঞ্চি",
		"64 MP + 8 MP + 2 MP": "৬৪ মেগাপিক্সেল + ৮ মেগাপিক্সেল + ২ মেগাপিক্সেল",
		"67W wired": "৬৭W তারযুক্ত",
		"8 GB / 12 GB": "৮ জিবি / ১২ GB",
		"AMOLED, 120Hz": "অ্যামোলেড, ১২০Hz",
		"Android 14, HyperOS": "অ্যান্ড্রয়েড ১৪, হাইপার ওএস",
		"Black, Blue, Yellow": "কালো, নীল, হলুদ",
		"Dimensity 7050": "ডাইমেনসিটি ৭০৫০",
		"GSM / HSPA / LTE / 5G": "জিএসএম / এইচএসপিএ / এলটিই / ৫G",
		"Glass front (Gorilla Glass 5), plastic back, plastic frame": "গ্লাস সামনে (গরিলা গ্লাস ৫), প্লাস্টিক পেছনে, প্লাস্টিক ফ্রেম",
		"IP54": "IP৫৪",
		"January 2025": "জানুয়ারি ২০২৫",
		"Mali-G68 MC4": "মালি-G৬৮ MC৪",
		"Mediatek Dimensity 7050 (6 nm)": "মিডিয়াটেক ডাইমেনসিটি ৭০৫০ (৬ ন্যানোমিটার)",
	}
}

// Seed inserts specification records for the product identified by slug 'poco-x7-5g'
func (s *SpecificationSeederMobilePocoX75g) Seed(db *gorm.DB) error {
	productSlug := "poco-x7-5g"

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

	// Override model-specific values for POCO X7 5G
	specs["Display Size"] = "6.67 inches"
	specs["Processor"] = "Dimensity 7050"
	specs["Chipset"] = "Mediatek Dimensity 7050 (6 nm)"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "Mali-G68 MC4"
	specs["Ram"] = "8 GB / 12 GB"
	specs["Storage"] = "256 GB / 512 GB"
	specs["Display Type"] = "AMOLED, 120Hz"
	specs["Resolution"] = "1080 x 2400 pixels"
	specs["Refresh Rate"] = "120Hz"
	specs["Build Material"] = "Glass front (Gorilla Glass 5), plastic back, plastic frame"
	specs["Weight"] = "189 g"
	specs["Water Resistance"] = "IP54"
	specs["Network Technology"] = "GSM / HSPA / LTE / 5G"
	specs["Rear Camera"] = "64 MP + 8 MP + 2 MP"
	specs["Front Camera"] = "16 MP"
	specs["Battery"] = "5000 mAh"
	specs["Fast Charging"] = "67W wired"
	specs["Operating System"] = "Android 14, HyperOS"
	specs["Available Colors"] = "Black, Blue, Yellow"
	specs["Announcement Date"] = "January 2025"
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
