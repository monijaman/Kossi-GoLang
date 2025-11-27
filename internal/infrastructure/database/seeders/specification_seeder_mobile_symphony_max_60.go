package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileSymphonyMax60 seeds specifications/options for product 'symphony-max-60'
type SpecificationSeederMobileSymphonyMax60 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSymphonyMax60 creates a new seeder instance
func NewSpecificationSeederMobileSymphonyMax60() *SpecificationSeederMobileSymphonyMax60 {
	return &SpecificationSeederMobileSymphonyMax60{BaseSeeder: BaseSeeder{name: "Specifications for Symphony Max 60"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSymphonyMax60) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080p@30fps": "১০৮০p@৩০fps",
		"13 MP + 2 MP": "১৩ MP + ২ MP",
		"168.75 x 78.1 x 8.79 mm": "১৬৮.৭৫ x ৭৮.১ x ৮.৭৯ মিমি",
		"210 g": "২১০ g",
		"4 GB": "৪ GB",
		"5 MP": "৫ MP",
		"6.75 inches": "৬.৭৫ ইঞ্চি",
		"6000 mAh": "৬০০০ এমএএইচ",
		"64 GB (expandable up to 256GB)": "৬৪ GB (expandable পর্যন্ত ২৫৬GB)",
		"720 × 1600 px": "৭২০ × ১৬০০ px",
		"90 Hz": "৯০ Hz",
		"Android 15": "Android ১৫",
		"IMG GE8322": "IMG GE৮৩২২",
		"IP64 dust & splash resistant": "IP৬৪ dust & splash resistant",
		"Li-Po (non-removable)": "লি-পো (অপসারণযোগ্য নয়)",
		"Octa-core 1.6 GHz": "Octa-core ১.৬ GHz",
		"Plastic body, glass front": "Plastic body, গ্লাস সামনে",
		"September 2025": "September ২০২৫",
		"Side fingerprint, accelerometer, proximity": "Side ফিঙ্গারপ্রিন্ট, অ্যাক্সিলেরোমিটার, প্রক্সিমিটি",
		"Space Black, Frost Blue, Classic Navy Blue, Titanium Gray, Divine Gold": "Space কালো, Frost নীল, Classic Navy নীল, Titanium ধূসর, Divine সোনালী",
		"Unisoc SC9863A1": "Unisoc SC৯৮৬৩A১",
		"Unisoc SC9863A1 (28 nm)": "Unisoc SC৯৮৬৩A১ (২৮ nm)",
		"Wi-Fi, Bluetooth 5.0, OTG": "Wi-Fi, Bluetooth ৫.০, OTG",
	}
}

// Seed inserts specification records for the product identified by slug 'symphony-max-60'
func (s *SpecificationSeederMobileSymphonyMax60) Seed(db *gorm.DB) error {
	productSlug := "symphony-max-60"

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

	// Override model-specific values for Symphony Max 60
	specs["Display Size"] = "6.75 inches"
	specs["Display Type"] = "IPS LCD"
	specs["Resolution"] = "720 × 1600 px"
	specs["Refresh Rate"] = "90 Hz"
	specs["Processor"] = "Unisoc SC9863A1"
	specs["Chipset"] = "Unisoc SC9863A1 (28 nm)"
	specs["Cpu Type"] = "Octa-core 1.6 GHz"
	specs["Gpu Type"] = "IMG GE8322"
	specs["Ram"] = "4 GB"
	specs["Storage"] = "64 GB (expandable up to 256GB)"
	specs["Operating System"] = "Android 15"
	specs["Rear Camera"] = "13 MP + 2 MP"
	specs["Camera Features"] = "AI modes, Portrait, Night mode"
	specs["Camera Video Resolution"] = "1080p@30fps"
	specs["Front Camera"] = "5 MP"
	specs["Front Camera Video Resolution"] = "1080p@30fps"
	specs["Battery"] = "6000 mAh"
	specs["Battery Type"] = "Li-Po (non-removable)"
	specs["Charging"] = "Standard wired"
	specs["Network Technology"] = "GSM / HSPA / LTE"
	specs["Sim Card Type"] = "Dual Nano-SIM"
	specs["Connectivity"] = "Wi-Fi, Bluetooth 5.0, OTG"
	specs["Sensors"] = "Side fingerprint, accelerometer, proximity"
	specs["Special Features"] = "IP64 dust & splash resistant"
	specs["Dimensions"] = "168.75 x 78.1 x 8.79 mm"
	specs["Weight"] = "210 g"
	specs["Build Material"] = "Plastic body, glass front"
	specs["Available Colors"] = "Space Black, Frost Blue, Classic Navy Blue, Titanium Gray, Divine Gold"
	specs["Announcement Date"] = "September 2025"
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
