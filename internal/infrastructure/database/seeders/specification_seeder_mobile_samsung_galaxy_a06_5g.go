package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileSamsungGalaxyA065g seeds specifications/options for product 'samsung-galaxy-a06-5g'
type SpecificationSeederMobileSamsungGalaxyA065g struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyA065g creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyA065g() *SpecificationSeederMobileSamsungGalaxyA065g {
	return &SpecificationSeederMobileSamsungGalaxyA065g{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy A06 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyA065g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"167.3 x 77.3 x 8.0 mm": "১৬৭.৩ x ৭৭.৩ x ৮.০ মিমি",
		"195 g": "১৯৫ g",
		"25W wired": "২৫W তারযুক্ত",
		"4 GB / 6 GB": "৪ জিবি / ৬ GB",
		"50 MP + 2 MP": "৫০ মেগাপিক্সেল + ২ মেগাপিক্সেল",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6.7 inches": "৬.৭ ইঞ্চি",
		"64 GB / 128 GB": "৬৪ জিবি / ১২৮ GB",
		"720 x 1600 pixels": "৭২০ x ১৬০০ পিক্সেল",
		"8 MP": "৮ মেগাপিক্সেল",
		"90Hz": "৯০Hz",
		"Android 14": "অ্যান্ড্রয়েড ১৪",
		"Black, Light Blue, Gold": "কালো, Light নীল, সোনালী",
		"GSM / HSPA / LTE / 5G": "জিএসএম / এইচএসপিএ / এলটিই / ৫G",
		"Glass front, plastic back, plastic frame": "গ্লাস সামনে, প্লাস্টিক পেছনে, প্লাস্টিক ফ্রেম",
		"Mali-G57 MC2": "মালি-G৫৭ MC২",
		"Mediatek Dimensity 6300": "মিডিয়াটেক ডাইমেনসিটি ৬৩০০",
		"Mediatek Dimensity 6300 (6 nm)": "মিডিয়াটেক ডাইমেনসিটি ৬৩০০ (৬ ন্যানোমিটার)",
		"October 2024": "অক্টোবর ২০২৪",
		"PLS LCD, 90Hz": "PLS এলসিডি, ৯০Hz",
	}
}

// Seed inserts specification records for the product identified by slug 'samsung-galaxy-a06-5g'
func (s *SpecificationSeederMobileSamsungGalaxyA065g) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-a06-5g"

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

	// Override model-specific values for Samsung Galaxy A06 5G
	specs["Display Size"] = "6.7 inches"
	specs["Processor"] = "Mediatek Dimensity 6300"
	specs["Chipset"] = "Mediatek Dimensity 6300 (6 nm)"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "Mali-G57 MC2"
	specs["Ram"] = "4 GB / 6 GB"
	specs["Storage"] = "64 GB / 128 GB"
	specs["Display Type"] = "PLS LCD, 90Hz"
	specs["Resolution"] = "720 x 1600 pixels"
	specs["Refresh Rate"] = "90Hz"
	specs["Build Material"] = "Glass front, plastic back, plastic frame"
	specs["Weight"] = "195 g"
	specs["Dimensions"] = "167.3 x 77.3 x 8.0 mm"
	specs["Water Resistance"] = "No"
	specs["Network Technology"] = "GSM / HSPA / LTE / 5G"
	specs["Rear Camera"] = "50 MP + 2 MP"
	specs["Front Camera"] = "8 MP"
	specs["Battery"] = "5000 mAh"
	specs["Fast Charging"] = "25W wired"
	specs["Operating System"] = "Android 14"
	specs["Available Colors"] = "Black, Light Blue, Gold"
	specs["Announcement Date"] = "October 2024"
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
