package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileOneplus13r seeds specifications/options for product 'oneplus-13r'
type SpecificationSeederMobileOneplus13r struct {
	BaseSeeder
}

// NewSpecificationSeederMobileOneplus13r creates a new seeder instance
func NewSpecificationSeederMobileOneplus13r() *SpecificationSeederMobileOneplus13r {
	return &SpecificationSeederMobileOneplus13r{BaseSeeder: BaseSeeder{name: "Specifications for OnePlus 13R"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileOneplus13r) getBanglaTranslations() map[string]string {
	return map[string]string{
		"120Hz": "১২০Hz",
		"1264 x 2780 pixels": "১২৬৪ x ২৭৮০ পিক্সেল",
		"128 GB / 256 GB": "১২৮ জিবি / ২৫৬ GB",
		"16 MP": "১৬ মেগাপিক্সেল",
		"163.3 x 75.3 x 8.8 mm": "১৬৩.৩ x ৭৫.৩ x ৮.৮ মিমি",
		"207 g": "২০৭ g",
		"5,500 mAh": "৫,৫০০ এমএএইচ",
		"50 MP + 8 MP + 2 MP": "৫০ মেগাপিক্সেল + ৮ মেগাপিক্সেল + ২ মেগাপিক্সেল",
		"5G": "৫G",
		"6.78 inches": "৬.৭৮ ইঞ্চি",
		"8 GB / 16 GB": "৮ জিবি / ১৬ GB",
		"Adreno 750": "অ্যাড্রেনো ৭৫০",
		"Android 15, OxygenOS 15": "অ্যান্ড্রয়েড ১৫, OxygenOS ১৫",
		"Corning Gorilla Glass Victus 2": "কর্নিং গরিলা গ্লাস ভিক্টাস ২",
		"Glass front/back, aluminum frame": "গ্লাস সামনে/পেছনে, অ্যালুমিনিয়াম ফ্রেম",
		"IP64": "IP৬৪",
		"Iron Gray, Cool Blue": "Iron ধূসর, Cool নীল",
		"January 2025": "জানুয়ারি ২০২৫",
		"LTPO AMOLED, 120Hz, HDR10+, 4500 nits": "এলটিপিও অ্যামোলেড, ১২০Hz, এইচডিআর১০+, ৪৫০০ nits",
		"Qualcomm SM8650-AB Snapdragon 8 Gen 3 (4 nm)": "কোয়ালকম SM৮৬৫০-AB স্ন্যাপড্রাগন ৮ জেন ৩ (৪ ন্যানোমিটার)",
		"Snapdragon 8 Gen 3": "স্ন্যাপড্রাগন ৮ জেন ৩",
	}
}

// Seed inserts specification records for the product identified by slug 'oneplus-13r'
func (s *SpecificationSeederMobileOneplus13r) Seed(db *gorm.DB) error {
	productSlug := "oneplus-13r"

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

	// Override model-specific values for OnePlus 13R
	specs["Display Size"] = "6.78 inches"
	specs["Processor"] = "Snapdragon 8 Gen 3"
	specs["Chipset"] = "Qualcomm SM8650-AB Snapdragon 8 Gen 3 (4 nm)"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "Adreno 750"
	specs["Ram"] = "8 GB / 16 GB"
	specs["Storage"] = "128 GB / 256 GB"
	specs["Display Type"] = "LTPO AMOLED, 120Hz, HDR10+, 4500 nits"
	specs["Resolution"] = "1264 x 2780 pixels"
	specs["Screen Protection"] = "Corning Gorilla Glass Victus 2"
	specs["Refresh Rate"] = "120Hz"
	specs["Build Material"] = "Glass front/back, aluminum frame"
	specs["Weight"] = "207 g"
	specs["Dimensions"] = "163.3 x 75.3 x 8.8 mm"
	specs["Water Resistance"] = "IP64"
	specs["Network Technology"] = "5G"
	specs["Rear Camera"] = "50 MP + 8 MP + 2 MP"
	specs["Front Camera"] = "16 MP"
	specs["Battery"] = "5,500 mAh"
	specs["Operating System"] = "Android 15, OxygenOS 15"
	specs["Available Colors"] = "Iron Gray, Cool Blue"
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
