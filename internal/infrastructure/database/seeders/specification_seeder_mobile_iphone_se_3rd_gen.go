package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileIphoneSe3rdGen seeds specifications/options for product 'iphone-se-3rd-gen'
type SpecificationSeederMobileIphoneSe3rdGen struct {
	BaseSeeder
}

// NewSpecificationSeederMobileIphoneSe3rdGen creates a new seeder instance
func NewSpecificationSeederMobileIphoneSe3rdGen() *SpecificationSeederMobileIphoneSe3rdGen {
	return &SpecificationSeederMobileIphoneSe3rdGen{BaseSeeder: BaseSeeder{name: "Specifications for iPhone SE (3rd Gen)"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileIphoneSe3rdGen) getBanglaTranslations() map[string]string {
	return map[string]string{
		"12 MP": "১২ মেগাপিক্সেল",
		"138.4 x 67.3 x 7.3 mm": "১৩৮.৪ x ৬৭.৩ x ৭.৩ মিমি",
		"144 g": "১৪৪ g",
		"2,018 mAh": "২,০১৮ এমএএইচ",
		"4 GB": "৪ GB",
		"4-core Apple GPU": "৪-কোর অ্যাপল জিপিইউ",
		"4.7 inches": "৪.৭ ইঞ্চি",
		"5G": "৫G",
		"60Hz": "৬০Hz",
		"64 GB / 128 GB / 256 GB": "৬৪ জিবি / ১২৮ জিবি / ২৫৬ GB",
		"7 MP": "৭ মেগাপিক্সেল",
		"750 x 1334 pixels": "৭৫০ x ১৩৩৪ পিক্সেল",
		"Aluminum frame, glass front/back": "অ্যালুমিনিয়াম ফ্রেম, গ্লাস সামনে/পেছনে",
		"Apple A15 Bionic": "অ্যাপল A১৫ বায়োনিক",
		"Apple A15 Bionic (5 nm)": "অ্যাপল A১৫ বায়োনিক (৫ ন্যানোমিটার)",
		"IP67": "IP৬৭",
		"March 2022": "মার্চ ২০২২",
		"Midnight, Starlight, Red": "মিডনাইট, Starলাইট, লাল",
		"iOS 15.4, upgradable": "আইওএস ১৫.৪, আপগ্রেডযোগ্য",
	}
}

// Seed inserts specification records for the product identified by slug 'iphone-se-3rd-gen'
func (s *SpecificationSeederMobileIphoneSe3rdGen) Seed(db *gorm.DB) error {
	productSlug := "iphone-se-3rd-gen"

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

	// Override model-specific values for iPhone SE (3rd Gen)
	specs["Display Size"] = "4.7 inches"
	specs["Processor"] = "Apple A15 Bionic"
	specs["Chipset"] = "Apple A15 Bionic (5 nm)"
	specs["Cpu Type"] = "Hexa-core"
	specs["Gpu Type"] = "4-core Apple GPU"
	specs["Ram"] = "4 GB"
	specs["Storage"] = "64 GB / 128 GB / 256 GB"
	specs["Display Type"] = "Retina IPS LCD"
	specs["Resolution"] = "750 x 1334 pixels"
	specs["Screen Protection"] = "Ion-strengthened glass"
	specs["Refresh Rate"] = "60Hz"
	specs["Build Material"] = "Aluminum frame, glass front/back"
	specs["Weight"] = "144 g"
	specs["Dimensions"] = "138.4 x 67.3 x 7.3 mm"
	specs["Water Resistance"] = "IP67"
	specs["Network Technology"] = "5G"
	specs["Rear Camera"] = "12 MP"
	specs["Front Camera"] = "7 MP"
	specs["Battery"] = "2,018 mAh"
	specs["Operating System"] = "iOS 15.4, upgradable"
	specs["Available Colors"] = "Midnight, Starlight, Red"
	specs["Announcement Date"] = "March 2022"
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
