package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileSamsungGalaxyS23Ultra seeds specifications/options for product 'samsung-galaxy-s23-ultra'
type SpecificationSeederMobileSamsungGalaxyS23Ultra struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyS23Ultra creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyS23Ultra() *SpecificationSeederMobileSamsungGalaxyS23Ultra {
	return &SpecificationSeederMobileSamsungGalaxyS23Ultra{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy S23 Ultra"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyS23Ultra) getBanglaTranslations() map[string]string {
	return map[string]string{
		"12 MP": "১২ মেগাপিক্সেল",
		"120Hz": "১২০Hz",
		"1440 x 3088 pixels": "১৪৪০ x ৩০৮৮ পিক্সেল",
		"163.4 x 78.1 x 8.9 mm": "১৬৩.৪ x ৭৮.১ x ৮.৯ মিমি",
		"200 MP + 10 MP + 10 MP + 12 MP": "২০০ মেগাপিক্সেল + ১০ মেগাপিক্সেল + ১০ মেগাপিক্সেল + ১২ মেগাপিক্সেল",
		"234 g": "২৩৪ g",
		"256 GB / 512 GB / 1 TB": "২৫৬ জিবি / ৫১২ জিবি / ১ টিবি",
		"5,000 mAh": "৫,০০০ এমএএইচ",
		"5G": "৫G",
		"6.8 inches": "৬.৮ ইঞ্চি",
		"8 GB / 12 GB": "৮ জিবি / ১২ GB",
		"Adreno 740": "অ্যাড্রেনো ৭৪০",
		"Aluminum frame, glass front/back": "অ্যালুমিনিয়াম ফ্রেম, গ্লাস সামনে/পেছনে",
		"Android 13, upgradable": "অ্যান্ড্রয়েড ১৩, আপগ্রেডযোগ্য",
		"Corning Gorilla Glass Victus 2": "কর্নিং গরিলা গ্লাস ভিক্টাস ২",
		"Dynamic AMOLED 2X, 120Hz, HDR10+": "Dynamic অ্যামোলেড ২X, ১২০Hz, এইচডিআর১০+",
		"February 2023": "ফেব্রুয়ারি ২০২৩",
		"IP68": "IP৬৮",
		"Phantom Black, Green, Cream, Lavender": "ফ্যান্টম কালো, সবুজ, ক্রিম, Lavender",
		"Qualcomm SM8550-AC Snapdragon 8 Gen 2 (4 nm)": "কোয়ালকম SM৮৫৫০-AC স্ন্যাপড্রাগন ৮ জেন ২ (৪ ন্যানোমিটার)",
		"Snapdragon 8 Gen 2": "স্ন্যাপড্রাগন ৮ জেন ২",
	}
}

// Seed inserts specification records for the product identified by slug 'samsung-galaxy-s23-ultra'
func (s *SpecificationSeederMobileSamsungGalaxyS23Ultra) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-s23-ultra"

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

	// Override model-specific values for Samsung Galaxy S23 Ultra
	specs["Display Size"] = "6.8 inches"
	specs["Processor"] = "Snapdragon 8 Gen 2"
	specs["Chipset"] = "Qualcomm SM8550-AC Snapdragon 8 Gen 2 (4 nm)"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "Adreno 740"
	specs["Ram"] = "8 GB / 12 GB"
	specs["Storage"] = "256 GB / 512 GB / 1 TB"
	specs["Display Type"] = "Dynamic AMOLED 2X, 120Hz, HDR10+"
	specs["Resolution"] = "1440 x 3088 pixels"
	specs["Screen Protection"] = "Corning Gorilla Glass Victus 2"
	specs["Refresh Rate"] = "120Hz"
	specs["Build Material"] = "Aluminum frame, glass front/back"
	specs["Weight"] = "234 g"
	specs["Dimensions"] = "163.4 x 78.1 x 8.9 mm"
	specs["Water Resistance"] = "IP68"
	specs["Network Technology"] = "5G"
	specs["Rear Camera"] = "200 MP + 10 MP + 10 MP + 12 MP"
	specs["Front Camera"] = "12 MP"
	specs["Battery"] = "5,000 mAh"
	specs["Operating System"] = "Android 13, upgradable"
	specs["Available Colors"] = "Phantom Black, Green, Cream, Lavender"
	specs["Announcement Date"] = "February 2023"
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
