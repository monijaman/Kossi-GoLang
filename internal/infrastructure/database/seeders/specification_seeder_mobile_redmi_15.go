package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileRedmi15 seeds specifications/options for product 'redmi-15'
type SpecificationSeederMobileRedmi15 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileRedmi15 creates a new seeder instance
func NewSpecificationSeederMobileRedmi15() *SpecificationSeederMobileRedmi15 {
	return &SpecificationSeederMobileRedmi15{BaseSeeder: BaseSeeder{name: "Specifications for Redmi 15"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileRedmi15) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"128 / 256 GB + microSD": "১২৮ / ২৫৬ জিবি + মাইক্রোএসডি",
		"169.48 × 80.45 × 8.40 mm": "১৬৯.৪৮ × ৮০.৪৫ × ৮.৪০ মিমি",
		"2025": "২০২৫",
		"214 g": "২১৪ g",
		"2340 × 1080 px": "২৩৪০ × ১০৮০ পিক্সেল",
		"33 W wired, 18 W reverse": "৩৩ W তারযুক্ত, ১৮ W রিভার্স",
		"5.0": "৫.০",
		"50 MP + secondary lens": "৫০ মেগাপিক্সেল + secondary lens",
		"6 / 8 GB (up to 16 GB with memory extension)": "৬ / ৮ জিবি (পর্যন্ত ১৬ GB with memবাy extension)",
		"6.9 inches": "৬.৯ ইঞ্চি",
		"7000 mAh (typ)": "৭০০০ এমএএইচ (typ)",
		"8 MP": "৮ মেগাপিক্সেল",
		"802.11 a/b/g/n/ac": "৮০২.১১ a/b/g/n/ac",
		"GSM / HSPA / LTE (4G)": "জিএসএম / এইচএসপিএ / এলটিই (৪G)",
		"Glass front, plastic frame/back": "গ্লাস সামনে, প্লাস্টিক ফ্রেম/পেছনে",
		"HyperOS 2": "হাইপার ওএস ২",
		"IP64": "IP৬৪",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Nano-SIM + Nano-SIM": "ন্যানো-সিম + ন্যানো-সিম",
		"Sandy Purple, Midnight Black, Titan Gray": "স্যান্ডি বেগুনি, মিডনাইট কালো, Titan ধূসর",
		"Side fingerprint, Accelerometer, etc.": "সাইড ফিঙ্গারপ্রিন্ট, অ্যাক্সিলেরোমিটার, etc.",
		"Snapdragon 685": "স্ন্যাপড্রাগন ৬৮৫",
		"Snapdragon 685 (6 nm)": "স্ন্যাপড্রাগন ৬৮৫ (৬ nm)",
		"Up to 144 Hz": "Up থেকে ১৪৪ Hz",
		"Yes, 3.5 mm": "হ্যাঁ, ৩.৫ মিমি",
	}
}

// Seed inserts specification records for the product identified by slug 'redmi-15'
func (s *SpecificationSeederMobileRedmi15) Seed(db *gorm.DB) error {
	productSlug := "redmi-15"

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

	// Override model-specific values for Redmi 15
	specs["Display Size"] = "6.9 inches"
	specs["Processor"] = "Snapdragon 685"
	specs["Chipset"] = "Snapdragon 685 (6 nm)"
	specs["Cpu Type"] = "Octa‑core"
	specs["Gpu Type"] = "Adreno (Qualcomm)"
	specs["Ram"] = "6 / 8 GB (up to 16 GB with memory extension)"
	specs["Storage"] = "128 / 256 GB + microSD"
	specs["Display Type"] = "IPS LCD"
	specs["Resolution"] = "2340 × 1080 px"
	specs["Refresh Rate"] = "Up to 144 Hz"
	specs["Build Material"] = "Glass front, plastic frame/back"
	specs["Weight"] = "214 g"
	specs["Dimensions"] = "169.48 × 80.45 × 8.40 mm"
	specs["Water Resistance"] = "IP64"
	specs["Network Technology"] = "GSM / HSPA / LTE (4G)"
	specs["Wifi Support"] = "802.11 a/b/g/n/ac"
	specs["Bluetooth Version"] = "5.0"
	specs["Nfc Support"] = "Yes"
	specs["Usb Type"] = "USB-C"
	specs["Rear Camera"] = "50 MP + secondary lens"
	specs["Camera Features"] = "LED flash, HDR"
	specs["Camera Video Resolution"] = "1080p @ 30fps"
	specs["Optical Zoom"] = "None"
	specs["Front Camera"] = "8 MP"
	specs["Front Camera Video Resolution"] = "1080p @ 30fps"
	specs["Operating System"] = "HyperOS 2"
	specs["Battery"] = "7000 mAh (typ)"
	specs["Battery Type"] = "Li-Ion (non-removable)"
	specs["Fast Charging"] = "33 W wired, 18 W reverse"
	specs["Wireless Charging"] = "No"
	specs["5G Support"] = "No"
	specs["Positioning System"] = "GPS, GLONASS, Galileo, BDS"
	specs["Sensors"] = "Side fingerprint, Accelerometer, etc."
	specs["Special Features"] = "Memory extension"
	specs["Sim Card Type"] = "Nano-SIM + Nano-SIM"
	specs["Loudspeaker Quality"] = "Mono speaker"
	specs["Audio Jack"] = "Yes, 3.5 mm"
	specs["Available Colors"] = "Sandy Purple, Midnight Black, Titan Gray"
	specs["Announcement Date"] = "2025"
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
