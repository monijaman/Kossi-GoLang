package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileIphone15Plus seeds specifications/options for product 'iphone-15-plus'
type SpecificationSeederMobileIphone15Plus struct {
	BaseSeeder
}

// NewSpecificationSeederMobileIphone15Plus creates a new seeder instance
func NewSpecificationSeederMobileIphone15Plus() *SpecificationSeederMobileIphone15Plus {
	return &SpecificationSeederMobileIphone15Plus{BaseSeeder: BaseSeeder{name: "Specifications for iPhone 15 Plus"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileIphone15Plus) getBanglaTranslations() map[string]string {
	return map[string]string{
		"12 MP": "১২ মেগাপিক্সেল",
		"128 GB / 256 GB / 512 GB": "১২৮ জিবি / ২৫৬ জিবি / ৫১২ GB",
		"1290 x 2796 pixels": "১২৯০ x ২৭৯৬ পিক্সেল",
		"160.9 x 77.8 x 7.8 mm": "১৬০.৯ x ৭৭.৮ x ৭.৮ মিমি",
		"201 g": "২০১ g",
		"4,383 mAh": "৪,৩৮৩ এমএএইচ",
		"48 MP + 12 MP": "৪৮ মেগাপিক্সেল + ১২ মেগাপিক্সেল",
		"5-core Apple GPU": "৫-কোর অ্যাপল জিপিইউ",
		"5G": "৫G",
		"6 GB": "৬ GB",
		"6.7 inches": "৬.৭ ইঞ্চি",
		"60Hz": "৬০Hz",
		"Aluminum frame, glass front/back": "অ্যালুমিনিয়াম ফ্রেম, গ্লাস সামনে/পেছনে",
		"Apple A16 Bionic": "অ্যাপল A১৬ বায়োনিক",
		"Apple A16 Bionic (4 nm)": "অ্যাপল A১৬ বায়োনিক (৪ ন্যানোমিটার)",
		"Black, Blue, Green, Yellow, Pink": "কালো, নীল, সবুজ, হলুদ, গোলাপী",
		"IP68": "IP৬৮",
		"September 2023": "সেপ্টেম্বর ২০২৩",
		"iOS 17, upgradable": "আইওএস ১৭, আপগ্রেডযোগ্য",
	}
}

// Seed inserts specification records for the product identified by slug 'iphone-15-plus'
func (s *SpecificationSeederMobileIphone15Plus) Seed(db *gorm.DB) error {
	productSlug := "iphone-15-plus"

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

	// Override model-specific values for iPhone 15 Plus
	specs["Display Size"] = "6.7 inches"
	specs["Processor"] = "Apple A16 Bionic"
	specs["Chipset"] = "Apple A16 Bionic (4 nm)"
	specs["Cpu Type"] = "Hexa-core"
	specs["Gpu Type"] = "5-core Apple GPU"
	specs["Ram"] = "6 GB"
	specs["Storage"] = "128 GB / 256 GB / 512 GB"
	specs["Display Type"] = "Super Retina XDR OLED"
	specs["Resolution"] = "1290 x 2796 pixels"
	specs["Screen Protection"] = "Ceramic Shield"
	specs["Refresh Rate"] = "60Hz"
	specs["Build Material"] = "Aluminum frame, glass front/back"
	specs["Weight"] = "201 g"
	specs["Dimensions"] = "160.9 x 77.8 x 7.8 mm"
	specs["Water Resistance"] = "IP68"
	specs["Network Technology"] = "5G"
	specs["Rear Camera"] = "48 MP + 12 MP"
	specs["Front Camera"] = "12 MP"
	specs["Battery"] = "4,383 mAh"
	specs["Operating System"] = "iOS 17, upgradable"
	specs["Available Colors"] = "Black, Blue, Green, Yellow, Pink"
	specs["Announcement Date"] = "September 2023"
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
