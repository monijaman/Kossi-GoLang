package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileSymphonyInnova40 seeds specifications/options for product 'symphony-innova-40'
type SpecificationSeederMobileSymphonyInnova40 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSymphonyInnova40 creates a new seeder instance
func NewSpecificationSeederMobileSymphonyInnova40() *SpecificationSeederMobileSymphonyInnova40 {
	return &SpecificationSeederMobileSymphonyInnova40{BaseSeeder: BaseSeeder{name: "Specifications for Symphony Innova 40"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSymphonyInnova40) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080p@30fps": "১০৮০p@৩০fps",
		"120 Hz": "১২০ Hz",
		"128 GB (uMCP)": "১২৮ GB (uMCP)",
		"169 x 78 x 9 mm": "১৬৯ x ৭৮ x ৯ মিমি",
		"18W": "১৮W",
		"210 g": "২১০ g",
		"50 MP + 0.08 MP + 0.08 MP": "৫০ MP + ০.০৮ MP + ০.০৮ MP",
		"6 GB / 12 GB / 16 GB (expandable)": "৬ GB / ১২ GB / ১৬ GB (expandable)",
		"6.75 inches": "৬.৭৫ ইঞ্চি",
		"6000 mAh": "৬০০০ এমএএইচ",
		"720 × 1600 px": "৭২০ × ১৬০০ px",
		"8 MP": "৮ MP",
		"Android 15": "Android ১৫",
		"August 2025": "August ২০২৫",
		"IP52 rating": "IP৫২ rating",
		"Mali-G57": "Mali-G৫৭",
		"Octa-core (A75 + A55)": "Octa-core (A৭৫ + A৫৫)",
		"Side fingerprint, face unlock, accelerometer, proximity, light, compass": "Side ফিঙ্গারপ্রিন্ট, face unlock, অ্যাক্সিলেরোমিটার, প্রক্সিমিটি, light, কম্পাস",
		"Space Black, Titanium Silver, Cosmic Gold, Ice Blue, Radium Green": "Space কালো, Titanium রূপালী, Cosmic সোনালী, Ice নীল, Radium সবুজ",
		"Unisoc T615": "Unisoc T৬১৫",
		"Unisoc T615 (12 nm)": "Unisoc T৬১৫ (১২ nm)",
	}
}

// Seed inserts specification records for the product identified by slug 'symphony-innova-40'
func (s *SpecificationSeederMobileSymphonyInnova40) Seed(db *gorm.DB) error {
	productSlug := "symphony-innova-40"

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

	// Override model-specific values for Symphony Innova 40
	specs["Display Size"] = "6.75 inches"
	specs["Display Type"] = "IPS LCD"
	specs["Resolution"] = "720 × 1600 px"
	specs["Refresh Rate"] = "120 Hz"
	specs["Processor"] = "Unisoc T615"
	specs["Chipset"] = "Unisoc T615 (12 nm)"
	specs["Cpu Type"] = "Octa-core (A75 + A55)"
	specs["Gpu Type"] = "Mali-G57"
	specs["Ram"] = "6 GB / 12 GB / 16 GB (expandable)"
	specs["Storage"] = "128 GB (uMCP)"
	specs["Operating System"] = "Android 15"
	specs["Rear Camera"] = "50 MP + 0.08 MP + 0.08 MP"
	specs["Camera Features"] = "Night mode, Portrait, Pro mode"
	specs["Camera Video Resolution"] = "1080p@30fps"
	specs["Front Camera"] = "8 MP"
	specs["Front Camera Video Resolution"] = "1080p@30fps"
	specs["Battery"] = "6000 mAh"
	specs["Battery Type"] = "Li-Polymer"
	specs["Fast Charging"] = "18W"
	specs["Network Technology"] = "GSM / HSPA / LTE"
	specs["Sim Card Type"] = "Dual Nano-SIM"
	specs["Connectivity"] = "Wi-Fi, Bluetooth, OTG"
	specs["Sensors"] = "Side fingerprint, face unlock, accelerometer, proximity, light, compass"
	specs["Special Features"] = "IP52 rating"
	specs["Dimensions"] = "169 x 78 x 9 mm"
	specs["Weight"] = "210 g"
	specs["Build Material"] = "Plastic frame and back"
	specs["Available Colors"] = "Space Black, Titanium Silver, Cosmic Gold, Ice Blue, Radium Green"
	specs["Announcement Date"] = "August 2025"
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
