package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileRedmi13c seeds specifications/options for product 'redmi-13c'
type SpecificationSeederMobileRedmi13c struct {
	BaseSeeder
}

// NewSpecificationSeederMobileRedmi13c creates a new seeder instance
func NewSpecificationSeederMobileRedmi13c() *SpecificationSeederMobileRedmi13c {
	return &SpecificationSeederMobileRedmi13c{BaseSeeder: BaseSeeder{name: "Specifications for Redmi 13C"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileRedmi13c) getBanglaTranslations() map[string]string {
	return map[string]string{
		"10 W": "১০ W",
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"128 / 256 GB + microSD": "১২৮ / ২৫৬ GB + microSD",
		"1600 × 720 px": "১৬০০ × ৭২০ px",
		"168 × 78 × 8.09 mm": "১৬৮ × ৭৮ × ৮.০৯ মিমি",
		"192 g": "১৯২ g",
		"2023": "২০২৩",
		"3.5 mm": "৩.৫ মিমি",
		"4 / 6 / 8 GB": "৪ / ৬ / ৮ GB",
		"5.3": "৫.৩",
		"50 MP + 2 MP + tiny lens": "৫০ MP + ২ MP + tiny lens",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6.74 inches": "৬.৭৪ ইঞ্চি",
		"8 MP": "৮ MP",
		"90 Hz": "৯০ Hz",
		"Black, Blue, Green": "কালো, নীল, সবুজ",
		"Fingerprint (side), Accelerometer, Proximity, Compass": "ফিঙ্গারপ্রিন্ট (side), অ্যাক্সিলেরোমিটার, প্রক্সিমিটি, কম্পাস",
		"Helio G85": "Helio G৮৫",
		"Li‑Ion (non-removable)": "Li‑Ion (অপসারণযোগ্য নয়)",
		"MIUI 14": "MIUI ১৪",
		"Mali‑G52 MC2": "Mali‑G৫২ MC২",
		"MediaTek Helio G85": "MediaTek Helio G৮৫",
		"Nano-SIM + Nano-SIM": "ন্যানো-সিম + ন্যানো-সিম",
		"Octa-core (2×Cortex-A75 + 6×Cortex-A55)": "Octa-core (২×Cortex-A৭৫ + ৬×Cortex-A৫৫)",
		"Plastic frame / back, glass front": "Plastic frame / back, গ্লাস সামনে",
		"Wi‑Fi 802.11 a/b/g/n/ac": "Wi‑Fi ৮০২.১১ a/b/g/n/ac",
	}
}

// Seed inserts specification records for the product identified by slug 'redmi-13c'
func (s *SpecificationSeederMobileRedmi13c) Seed(db *gorm.DB) error {
	productSlug := "redmi-13c"

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

	// Override model-specific values for Redmi 13C
	specs["Display Size"] = "6.74 inches"
	specs["Processor"] = "MediaTek Helio G85"
	specs["Chipset"] = "Helio G85"
	specs["Cpu Type"] = "Octa-core (2×Cortex-A75 + 6×Cortex-A55)"
	specs["Gpu Type"] = "Mali‑G52 MC2"
	specs["Ram"] = "4 / 6 / 8 GB"
	specs["Storage"] = "128 / 256 GB + microSD"
	specs["Display Type"] = "IPS LCD"
	specs["Resolution"] = "1600 × 720 px"
	specs["Refresh Rate"] = "90 Hz"
	specs["Build Material"] = "Plastic frame / back, glass front"
	specs["Weight"] = "192 g"
	specs["Dimensions"] = "168 × 78 × 8.09 mm"
	specs["Network Technology"] = "GSM / HSPA / LTE"
	specs["Wifi Support"] = "Wi‑Fi 802.11 a/b/g/n/ac"
	specs["Bluetooth Version"] = "5.3"
	specs["Nfc Support"] = "Yes"
	specs["Usb Type"] = "USB‑C"
	specs["Rear Camera"] = "50 MP + 2 MP + tiny lens"
	specs["Camera Features"] = "PDAF, HDR, Night, Portrait"
	specs["Camera Video Resolution"] = "1080p @ 30fps"
	specs["Optical Zoom"] = "None"
	specs["Front Camera"] = "8 MP"
	specs["Front Camera Video Resolution"] = "1080p @ 30fps"
	specs["Operating System"] = "MIUI 14"
	specs["Battery"] = "5000 mAh"
	specs["Battery Type"] = "Li‑Ion (non-removable)"
	specs["Fast Charging"] = "10 W"
	specs["Wireless Charging"] = "No"
	specs["5G Support"] = "No"
	specs["Positioning System"] = "GPS, GLONASS, GALILEO"
	specs["Sensors"] = "Fingerprint (side), Accelerometer, Proximity, Compass"
	specs["Sim Card Type"] = "Nano-SIM + Nano-SIM"
	specs["Loudspeaker Quality"] = "Single speaker"
	specs["Audio Jack"] = "3.5 mm"
	specs["Available Colors"] = "Black, Blue, Green"
	specs["Announcement Date"] = "2023"
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
