package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileIphone17ProMax seeds specifications/options for product 'iphone-17-pro-max'
type SpecificationSeederMobileIphone17ProMax struct {
	BaseSeeder
}

// NewSpecificationSeederMobileIphone17ProMax creates a new seeder instance
func NewSpecificationSeederMobileIphone17ProMax() *SpecificationSeederMobileIphone17ProMax {
	return &SpecificationSeederMobileIphone17ProMax{BaseSeeder: BaseSeeder{name: "Specifications for iPhone 17 Pro Max"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileIphone17ProMax) getBanglaTranslations() map[string]string {
	return map[string]string{
		"12 GB": "১২ GB",
		"12 MP, f/2.2": "১২ MP, f/২.২",
		"120Hz ProMotion": "১২০Hz ProMotion",
		"165.3 x 76.3 x 8.5 mm (6.51 x 3.00 x 0.33 in)": "১৬৫.৩ x ৭৬.৩ x ৮.৫ মিমি (৬.৫১ x ৩.০০ x ০.৩৩ in)",
		"235 g (8.28 oz)": "২৩৫ g (৮.২৮ oz)",
		"256 GB / 512 GB / 1 TB": "২৫৬ GB / ৫১২ GB / ১ TB",
		"2796 x 1290 pixels (~456 ppi density)": "২৭৯৬ x ১২৯০ pixels (~৪৫৬ ppi density)",
		"2x + 5x Optical Zoom": "২x + ৫x অপটিক্যাল জুম",
		"4,500 mAh (typical)": "৪,৫০০ এমএএইচ (সাধারণ)",
		"45W wired; 25W MagSafe wireless; 5W reverse wireless": "৪৫W wired; ২৫W MagSafe wireless; ৫W reverse wireless",
		"48 MP + 12 MP + 12 MP": "৪৮ MP + ১২ MP + ১২ MP",
		"4K@24/25/30/60fps, 1080p@30/60/120/240fps, ProRes video recording": "৪K@২৪/২৫/৩০/৬০fps, ১০৮০p@৩০/৬০/১২০/২৪০fps, ProRes video recording",
		"4K@24/25/30/60fps, 1080p@30/60fps": "৪K@২৪/২৫/৩০/৬০fps, ১০৮০p@৩০/৬০fps",
		"5.4, A2DP, LE, EDR": "৫.৪, A২DP, LE, EDR",
		"6-core Apple GPU": "৬-core Apple GPU",
		"6.9 inches": "৬.৯ ইঞ্চি",
		"Apple A19 Pro": "Apple A১৯ Pro",
		"Black, White, Gold, Rose Gold, Blue, Deep Purple": "কালো, সাদা, সোনালী, Rose সোনালী, নীল, Deep বেগুনি",
		"GSM / CDMA / HSPA / EVDO / LTE / 5G": "GSM / CDMA / HSPA / EVDO / LTE / ৫G",
		"Hexa-core (2x Performance + 4x Efficiency)": "Hexa-core (২x Performance + ৪x Efficiency)",
		"IP68 dust tight and water resistant (immersible up to 6m for 30 min)": "IP৬৮ dust tight and water resistant (immersible up to ৬m for ৩০ min)",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Nano-SIM + eSIM": "ন্যানো-সিম + ই-সিম",
		"September 2024": "September ২০২৪",
		"Super Retina XDR LTPO OLED, 120Hz, HDR10, Dolby Vision": "Super Retina XDR LTPO OLED, ১২০Hz, HDR১০, Dolby Vision",
		"USB Type-C with USB 3.1 speeds": "USB Type-C with USB ৩.১ speeds",
		"Wi-Fi 802.11 a/b/g/n/ac/6e (802.11ax)": "Wi-Fi ৮০২.১১ a/b/g/n/ac/৬e (৮০২.১১ax)",
		"Yes - 25W MagSafe wireless; 5W reverse wireless": "Yes - ২৫W MagSafe wireless; ৫W reverse wireless",
		"iOS 18": "iOS ১৮",
	}
}

// Seed inserts specification records for the product identified by slug 'iphone-17-pro-max'
func (s *SpecificationSeederMobileIphone17ProMax) Seed(db *gorm.DB) error {
	productSlug := "iphone-17-pro-max"

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

	// Override model-specific values for iPhone 17 Pro Max
	specs["Display Size"] = "6.9 inches"
	specs["Processor"] = "Apple A19 Pro"
	specs["Chipset"] = "Apple A19 Pro"
	specs["Cpu Type"] = "Hexa-core (2x Performance + 4x Efficiency)"
	specs["Gpu Type"] = "6-core Apple GPU"
	specs["Ram"] = "12 GB"
	specs["Storage"] = "256 GB / 512 GB / 1 TB"
	specs["Display Type"] = "Super Retina XDR LTPO OLED, 120Hz, HDR10, Dolby Vision"
	specs["Resolution"] = "2796 x 1290 pixels (~456 ppi density)"
	specs["Screen Protection"] = "Ceramic Shield, harder than any smartphone glass"
	specs["Refresh Rate"] = "120Hz ProMotion"
	specs["Build Material"] = "Titanium Design, Ceramic Shield front"
	specs["Weight"] = "235 g (8.28 oz)"
	specs["Dimensions"] = "165.3 x 76.3 x 8.5 mm (6.51 x 3.00 x 0.33 in)"
	specs["Water Resistance"] = "IP68 dust tight and water resistant (immersible up to 6m for 30 min)"
	specs["Network Technology"] = "GSM / CDMA / HSPA / EVDO / LTE / 5G"
	specs["Wifi Support"] = "Wi-Fi 802.11 a/b/g/n/ac/6e (802.11ax)"
	specs["Bluetooth Version"] = "5.4, A2DP, LE, EDR"
	specs["Nfc Support"] = "Yes"
	specs["Usb Type"] = "USB Type-C with USB 3.1 speeds"
	specs["Rear Camera"] = "48 MP + 12 MP + 12 MP"
	specs["Camera Features"] = "Optical image stabilization (OIS), Sensor-shift stabilization, Photonic Engine"
	specs["Camera Video Resolution"] = "4K@24/25/30/60fps, 1080p@30/60/120/240fps, ProRes video recording"
	specs["Optical Zoom"] = "2x + 5x Optical Zoom"
	specs["Front Camera"] = "12 MP, f/2.2"
	specs["Front Camera Video Resolution"] = "4K@24/25/30/60fps, 1080p@30/60fps"
	specs["Operating System"] = "iOS 18"
	specs["Battery"] = "4,500 mAh (typical)"
	specs["Battery Type"] = "Li-Ion (non-removable)"
	specs["Fast Charging"] = "45W wired; 25W MagSafe wireless; 5W reverse wireless"
	specs["Wireless Charging"] = "Yes - 25W MagSafe wireless; 5W reverse wireless"
	specs["5G Support"] = "Yes"
	specs["Positioning System"] = "GPS, GLONASS, BDS, GALILEO, QZSS"
	specs["Sensors"] = "Face ID, accelerometer, gyro, proximity, compass, barometer"
	specs["Special Features"] = "Smart HDR, Night mode, Action mode, Portrait mode, Spatial video"
	specs["Sim Card Type"] = "Nano-SIM + eSIM"
	specs["Loudspeaker Quality"] = "Spatial Audio with stereo speakers"
	specs["Audio Quality"] = "Lossless and Hi-Res Lossless audio"
	specs["Audio Jack"] = "No"
	specs["Available Colors"] = "Black, White, Gold, Rose Gold, Blue, Deep Purple"
	specs["Announcement Date"] = "September 2024"
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
