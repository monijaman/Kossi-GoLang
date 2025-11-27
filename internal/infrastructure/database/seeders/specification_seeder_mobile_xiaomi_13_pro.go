package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileXiaomi13Pro seeds specifications/options for product 'xiaomi-13-pro'
type SpecificationSeederMobileXiaomi13Pro struct {
	BaseSeeder
}

// NewSpecificationSeederMobileXiaomi13Pro creates a new seeder instance
func NewSpecificationSeederMobileXiaomi13Pro() *SpecificationSeederMobileXiaomi13Pro {
	return &SpecificationSeederMobileXiaomi13Pro{BaseSeeder: BaseSeeder{name: "Specifications for Xiaomi 13 Pro"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileXiaomi13Pro) getBanglaTranslations() map[string]string {
	return map[string]string{
		"120Hz": "১২০Hz",
		"128 GB / 256 GB / 512 GB": "১২৮ GB / ২৫৬ GB / ৫১২ GB",
		"1440 x 3200 pixels": "১৪৪০ x ৩২০০ pixels",
		"162.9 x 74.6 x 8.4 mm": "১৬২.৯ x ৭৪.৬ x ৮.৪ মিমি",
		"210 g": "২১০ g",
		"32 MP": "৩২ MP",
		"4,820 mAh": "৪,৮২০ এমএএইচ",
		"50 MP + 50 MP + 50 MP": "৫০ MP + ৫০ MP + ৫০ MP",
		"5G": "৫G",
		"6.73 inches": "৬.৭৩ ইঞ্চি",
		"8 GB / 12 GB": "৮ GB / ১২ GB",
		"Adreno 740": "Adreno ৭৪০",
		"Android 13, upgradable": "Android ১৩, upgradable",
		"Ceramic White, Ceramic Black, Ceramic Flora Green, Mountain Blue": "Ceramic সাদা, Ceramic কালো, Ceramic Flora সবুজ, Mountain নীল",
		"December 2022": "December ২০২২",
		"Glass front, ceramic/polymer back, aluminum frame": "গ্লাস সামনে, ceramic/polymer back, অ্যালুমিনিয়াম ফ্রেম",
		"IP68": "IP৬৮",
		"LTPO AMOLED, 120Hz, Dolby Vision, HDR10+, 1900 nits": "LTPO AMOLED, ১২০Hz, Dolby Vision, HDR১০+, ১৯০০ nits",
		"Qualcomm SM8550-AB Snapdragon 8 Gen 2 (4 nm)": "Qualcomm SM৮৫৫০-AB Snapdragon ৮ Gen ২ (৪ nm)",
		"Snapdragon 8 Gen 2": "Snapdragon ৮ Gen ২",
	}
}

// Seed inserts specification records for the product identified by slug 'xiaomi-13-pro'
func (s *SpecificationSeederMobileXiaomi13Pro) Seed(db *gorm.DB) error {
	productSlug := "xiaomi-13-pro"

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

	// Override model-specific values for Xiaomi 13 Pro
	specs["Display Size"] = "6.73 inches"
	specs["Processor"] = "Snapdragon 8 Gen 2"
	specs["Chipset"] = "Qualcomm SM8550-AB Snapdragon 8 Gen 2 (4 nm)"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "Adreno 740"
	specs["Ram"] = "8 GB / 12 GB"
	specs["Storage"] = "128 GB / 256 GB / 512 GB"
	specs["Display Type"] = "LTPO AMOLED, 120Hz, Dolby Vision, HDR10+, 1900 nits"
	specs["Resolution"] = "1440 x 3200 pixels"
	specs["Screen Protection"] = "Corning Gorilla Glass Victus"
	specs["Refresh Rate"] = "120Hz"
	specs["Build Material"] = "Glass front, ceramic/polymer back, aluminum frame"
	specs["Weight"] = "210 g"
	specs["Dimensions"] = "162.9 x 74.6 x 8.4 mm"
	specs["Water Resistance"] = "IP68"
	specs["Network Technology"] = "5G"
	specs["Rear Camera"] = "50 MP + 50 MP + 50 MP"
	specs["Front Camera"] = "32 MP"
	specs["Battery"] = "4,820 mAh"
	specs["Operating System"] = "Android 13, upgradable"
	specs["Available Colors"] = "Ceramic White, Ceramic Black, Ceramic Flora Green, Mountain Blue"
	specs["Announcement Date"] = "December 2022"
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
