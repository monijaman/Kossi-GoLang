package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileRedmi15c seeds specifications/options for product 'redmi-15c'
type SpecificationSeederMobileRedmi15c struct {
	BaseSeeder
}

// NewSpecificationSeederMobileRedmi15c creates a new seeder instance
func NewSpecificationSeederMobileRedmi15c() *SpecificationSeederMobileRedmi15c {
	return &SpecificationSeederMobileRedmi15c{BaseSeeder: BaseSeeder{name: "Specifications for Redmi 15C"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileRedmi15c) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1080p @ 30fps": "১০৮০p @ ৩০fps",
		"120 Hz": "১২০ Hz",
		"128 / 256 GB + microSD": "১২৮ / ২৫৬ জিবি + মাইক্রোএসডি",
		"1600 × 720 px": "১৬০০ × ৭২০ পিক্সেল",
		"171.6 × 79.5 × ~8 mm": "১৭১.৬ × ৭৯.৫ × ~৮ মিমি",
		"2025": "২০২৫",
		"205 g": "২০৫ g",
		"33 W wired": "৩৩ W তারযুক্ত",
		"4 / 6 / 8 GB": "৪ / ৬ / ৮ GB",
		"5.4": "৫.৪",
		"50 MP + secondary lens": "৫০ মেগাপিক্সেল + secondary lens",
		"6.9 inches": "৬.৯ ইঞ্চি",
		"6000 mAh (typ)": "৬০০০ এমএএইচ (typ)",
		"8 MP": "৮ মেগাপিক্সেল",
		"802.11 a/b/g/n": "৮০২.১১ a/b/g/n",
		"GSM / HSPA / LTE (4G)": "জিএসএম / এইচএসপিএ / এলটিই (৪G)",
		"Glass front, plastic back": "গ্লাস সামনে, প্লাস্টিক পেছনে",
		"Helio G81 Ultra": "হেলিও G৮১ আল্ট্রা",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Mali‑G52 MC2": "মালি‑G৫২ MC২",
		"MediaTek Helio G81 Ultra": "মিডিয়াটেক হেলিও G৮১ আল্ট্রা",
		"Moonlight Blue, Mint Green, Midnight Gray, Twilight Orange": "মুনলাইট নীল, মিন্ট সবুজ, মিডনাইট ধূসর, টোয়াইলাইট Orange",
		"Nano-SIM + Nano-SIM": "ন্যানো-সিম + ন্যানো-সিম",
		"Side fingerprint, Accelerometer": "সাইড ফিঙ্গারপ্রিন্ট, অ্যাক্সিলেরোমিটার",
		"Yes, 3.5 mm": "হ্যাঁ, ৩.৫ মিমি",
	}
}

// Seed inserts specification records for the product identified by slug 'redmi-15c'
func (s *SpecificationSeederMobileRedmi15c) Seed(db *gorm.DB) error {
	productSlug := "redmi-15c"

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

	// Override model-specific values for Redmi 15C
	specs["Display Size"] = "6.9 inches"
	specs["Processor"] = "MediaTek Helio G81 Ultra"
	specs["Chipset"] = "Helio G81 Ultra"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "Mali‑G52 MC2"
	specs["Ram"] = "4 / 6 / 8 GB"
	specs["Storage"] = "128 / 256 GB + microSD"
	specs["Display Type"] = "IPS LCD"
	specs["Resolution"] = "1600 × 720 px"
	specs["Refresh Rate"] = "120 Hz"
	specs["Build Material"] = "Glass front, plastic back"
	specs["Weight"] = "205 g"
	specs["Dimensions"] = "171.6 × 79.5 × ~8 mm"
	specs["Network Technology"] = "GSM / HSPA / LTE (4G)"
	specs["Wifi Support"] = "802.11 a/b/g/n"
	specs["Bluetooth Version"] = "5.4"
	specs["Nfc Support"] = "Yes"
	specs["Usb Type"] = "USB-C"
	specs["Rear Camera"] = "50 MP + secondary lens"
	specs["Camera Features"] = "LED flash"
	specs["Camera Video Resolution"] = "1080p @ 30fps"
	specs["Optical Zoom"] = "None"
	specs["Front Camera"] = "8 MP"
	specs["Front Camera Video Resolution"] = "1080p @ 30fps"
	specs["Operating System"] = "HyperOS / MIUI"
	specs["Battery"] = "6000 mAh (typ)"
	specs["Battery Type"] = "Li-Ion (non-removable)"
	specs["Fast Charging"] = "33 W wired"
	specs["Wireless Charging"] = "No"
	specs["5G Support"] = "No"
	specs["Sensors"] = "Side fingerprint, Accelerometer"
	specs["Sim Card Type"] = "Nano-SIM + Nano-SIM"
	specs["Loudspeaker Quality"] = "Single speaker"
	specs["Audio Jack"] = "Yes, 3.5 mm"
	specs["Available Colors"] = "Moonlight Blue, Mint Green, Midnight Gray, Twilight Orange"
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
