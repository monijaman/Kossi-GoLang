package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobilePocoX6Pro5g seeds specifications/options for product 'poco-x6-pro-5g'
type SpecificationSeederMobilePocoX6Pro5g struct {
	BaseSeeder
}

// NewSpecificationSeederMobilePocoX6Pro5g creates a new seeder instance
func NewSpecificationSeederMobilePocoX6Pro5g() *SpecificationSeederMobilePocoX6Pro5g {
	return &SpecificationSeederMobilePocoX6Pro5g{BaseSeeder: BaseSeeder{name: "Specifications for POCO X6 Pro 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobilePocoX6Pro5g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"120Hz": "১২০Hz",
		"1220 x 2712 pixels": "১২২০ x ২৭১২ pixels",
		"16 MP": "১৬ MP",
		"186 g": "১৮৬ g",
		"256 GB / 512 GB": "২৫৬ GB / ৫১২ GB",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6.67 inches": "৬.৬৭ ইঞ্চি",
		"64 MP + 8 MP + 2 MP": "৬৪ MP + ৮ MP + ২ MP",
		"67W wired": "৬৭W তারযুক্ত",
		"8 GB / 12 GB": "৮ GB / ১২ GB",
		"AMOLED, 120Hz, Dolby Vision, HDR10+": "AMOLED, ১২০Hz, Dolby Vision, HDR১০+",
		"Android 14, HyperOS": "Android ১৪, HyperOS",
		"Black, Yellow, Gray": "কালো, হলুদ, ধূসর",
		"Dimensity 8300 Ultra": "Dimensity ৮৩০০ Ultra",
		"GSM / HSPA / LTE / 5G": "GSM / HSPA / LTE / ৫G",
		"Glass front (Gorilla Glass 5), plastic back or silicone polymer back, plastic frame": "গ্লাস সামনে (Gorilla Glass ৫), প্লাস্টিক পেছনে or silicone polymer back, plastic frame",
		"IP54": "IP৫৪",
		"January 2024": "January ২০২৪",
		"Mali-G615-MC6": "Mali-G৬১৫-MC৬",
		"Mediatek Dimensity 8300 Ultra (4 nm)": "Mediatek Dimensity ৮৩০০ Ultra (৪ nm)",
	}
}

// Seed inserts specification records for the product identified by slug 'poco-x6-pro-5g'
func (s *SpecificationSeederMobilePocoX6Pro5g) Seed(db *gorm.DB) error {
	productSlug := "poco-x6-pro-5g"

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

	// Override model-specific values for POCO X6 Pro 5G
	specs["Display Size"] = "6.67 inches"
	specs["Processor"] = "Dimensity 8300 Ultra"
	specs["Chipset"] = "Mediatek Dimensity 8300 Ultra (4 nm)"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "Mali-G615-MC6"
	specs["Ram"] = "8 GB / 12 GB"
	specs["Storage"] = "256 GB / 512 GB"
	specs["Display Type"] = "AMOLED, 120Hz, Dolby Vision, HDR10+"
	specs["Resolution"] = "1220 x 2712 pixels"
	specs["Refresh Rate"] = "120Hz"
	specs["Build Material"] = "Glass front (Gorilla Glass 5), plastic back or silicone polymer back, plastic frame"
	specs["Weight"] = "186 g"
	specs["Water Resistance"] = "IP54"
	specs["Network Technology"] = "GSM / HSPA / LTE / 5G"
	specs["Rear Camera"] = "64 MP + 8 MP + 2 MP"
	specs["Front Camera"] = "16 MP"
	specs["Battery"] = "5000 mAh"
	specs["Fast Charging"] = "67W wired"
	specs["Operating System"] = "Android 14, HyperOS"
	specs["Available Colors"] = "Black, Yellow, Gray"
	specs["Announcement Date"] = "January 2024"
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
