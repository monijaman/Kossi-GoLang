package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobilePocoX7Pro5g seeds specifications/options for product 'poco-x7-pro-5g'
type SpecificationSeederMobilePocoX7Pro5g struct {
	BaseSeeder
}

// NewSpecificationSeederMobilePocoX7Pro5g creates a new seeder instance
func NewSpecificationSeederMobilePocoX7Pro5g() *SpecificationSeederMobilePocoX7Pro5g {
	return &SpecificationSeederMobilePocoX7Pro5g{BaseSeeder: BaseSeeder{name: "Specifications for Poco X7 Pro 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobilePocoX7Pro5g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"120 Hz": "১২০ Hz",
		"160.75 × 75.24 × ~8.3 mm": "১৬০.৭৫ × ৭৫.২৪ × ~৮.৩ মিমি",
		"195 / 198 g": "১৯৫ / ১৯৮ g",
		"2025": "২০২৫",
		"256 / 512 GB (UFS 4.0)": "২৫৬ / ৫১২ জিবি (UFS ৪.০)",
		"2712 × 1220 px (1.5K)": "২৭১২ × ১২২০ পিক্সেল (১.৫K)",
		"4K, 1080p": "৪K, ১০৮০p",
		"5.4": "৫.৪",
		"50 MP + 8 MP + additional": "৫০ মেগাপিক্সেল + ৮ মেগাপিক্সেল + additional",
		"6.67 inches": "৬.৬৭ ইঞ্চি",
		"8 / 12 GB": "৮ / ১২ GB",
		"802.11 a/b/g/n/ac/ax": "৮০২.১১ a/b/g/n/ac/ax",
		"Black, Green, Yellow": "কালো, সবুজ, হলুদ",
		"Dimensity 8400 Ultra (4 nm)": "ডাইমেনসিটি ৮৪০০ আল্ট্রা (৪ nm)",
		"Dolby Atmos, 360° ambient light": "ডলবি অ্যাটমস, ৩৬০° ambient লাইট",
		"GSM / HSPA / LTE / 5G": "জিএসএম / এইচএসপিএ / এলটিই / ৫G",
		"Glass front, plastic / eco-leather back": "গ্লাস সামনে, প্লাস্টিক / eco-চামড়া পেছনে",
		"HyperOS 2": "হাইপার ওএস ২",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Mali‑G720": "মালি‑G৭২০",
		"MediaTek Dimensity 8400 Ultra": "মিডিয়াটেক ডাইমেনসিটি ৮৪০০ আল্ট্রা",
		"Nano-SIM + Nano-SIM": "ন্যানো-সিম + ন্যানো-সিম",
		"Octa-core (up to 3.25 GHz)": "অক্টা-কোর (পর্যন্ত ৩.২৫ গিগাহার্টজ)",
		"Proximity, Accelerometer, Gyro, Compass, IR": "প্রক্সিমিটি, অ্যাক্সিলেরোমিটার, জাইরো, কম্পাস, IR",
	}
}

// Seed inserts specification records for the product identified by slug 'poco-x7-pro-5g'
func (s *SpecificationSeederMobilePocoX7Pro5g) Seed(db *gorm.DB) error {
	productSlug := "poco-x7-pro-5g"

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

	// Override model-specific values for Poco X7 Pro 5G
	specs["Display Size"] = "6.67 inches"
	specs["Processor"] = "MediaTek Dimensity 8400 Ultra"
	specs["Chipset"] = "Dimensity 8400 Ultra (4 nm)"
	specs["Cpu Type"] = "Octa-core (up to 3.25 GHz)"
	specs["Gpu Type"] = "Mali‑G720"
	specs["Ram"] = "8 / 12 GB"
	specs["Storage"] = "256 / 512 GB (UFS 4.0)"
	specs["Display Type"] = "AMOLED"
	specs["Resolution"] = "2712 × 1220 px (1.5K)"
	specs["Refresh Rate"] = "120 Hz"
	specs["Build Material"] = "Glass front, plastic / eco-leather back"
	specs["Weight"] = "195 / 198 g"
	specs["Dimensions"] = "160.75 × 75.24 × ~8.3 mm"
	specs["Network Technology"] = "GSM / HSPA / LTE / 5G"
	specs["Wifi Support"] = "802.11 a/b/g/n/ac/ax"
	specs["Bluetooth Version"] = "5.4"
	specs["Usb Type"] = "USB‑C"
	specs["Rear Camera"] = "50 MP + 8 MP + additional"
	specs["Camera Features"] = "OIS, HDR"
	specs["Camera Video Resolution"] = "4K, 1080p"
	specs["Optical Zoom"] = "None"
	specs["Operating System"] = "HyperOS 2"
	specs["Battery Type"] = "Li-Ion (non-removable)"
	specs["Wireless Charging"] = "No"
	specs["5G Support"] = "Yes"
	specs["Positioning System"] = "GPS, GLONASS, Galileo, BDS"
	specs["Sensors"] = "Proximity, Accelerometer, Gyro, Compass, IR"
	specs["Special Features"] = "Dolby Atmos, 360° ambient light"
	specs["Sim Card Type"] = "Nano-SIM + Nano-SIM"
	specs["Loudspeaker Quality"] = "Dual stereo"
	specs["Audio Quality"] = "Hi‑Res Audio"
	specs["Audio Jack"] = "No"
	specs["Available Colors"] = "Black, Green, Yellow"
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
