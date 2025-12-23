package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederMobileIPhone15 seeds specifications/options for product 'iphone-15'
type SpecificationSeederMobileIPhone15 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileIPhone15 creates a new seeder instance
func NewSpecificationSeederMobileIPhone15() *SpecificationSeederMobileIPhone15 {
	return &SpecificationSeederMobileIPhone15{BaseSeeder: BaseSeeder{name: "Specifications for iPhone 15"}}
}

func (s *SpecificationSeederMobileIPhone15) getBanglaTranslations() map[string]string {
	commonTranslations := s.GetCommonBanglaTranslations()

	// Add iPhone 15 specific translations
	iPhoneTranslations := map[string]string{
		"6.1 inches":                         "৬.১ ইঞ্চি",
		"Super Retina XDR OLED":              "সুপার রেটিনা XDR OLED",
		"2556 x 1179 pixels":                 "২৫৫৬ × ১১৭৯ পিক্সেল",
		"460 ppi":                            "৪৬০ পিপিআই",
		"60Hz":                               "৬০ হার্জ",
		"Ceramic Shield":                     "সিরামিক শিল্ড",
		"Apple A16 Bionic":                   "অ্যাপল A১৬ বায়োনিক",
		"Hexa-core":                          "হেক্সা-কোর",
		"5-core Apple GPU":                   "৫-কোর অ্যাপল GPU",
		"6 GB":                               "৬ জিবি",
		"128 GB / 256 GB / 512 GB":           "১২৮ জিবি / ২৫৬ জিবি / ৫১২ জিবি",
		"48 MP + 12 MP":                      "৪৮ মেগাপিক্সেল + ১২ মেগাপিক্সেল",
		"12 MP":                              "১২ মেগাপিক্সেল",
		"4K@24/30/60fps, 1080p@30/60/120fps": "৪কে @ ২৪/৩০/৬০ এফপিএস, ১০৮০পি @ ৩০/৬০/১২০ এফপিএস",
		"iOS 17, upgradable":                 "iOS ১৭, আপগ্রেডযোগ্য",
		"Li-Ion 3349 mAh":                    "লি-আয়ন ৩৩৪৯ মিলিঅ্যাম্পিয়ার আওয়ার",
		"Aluminum frame, glass front/back":   "অ্যালুমিনিয়াম ফ্রেম, গ্লাস ফ্রন্ট/ব্যাক",
		"147.6 x 71.6 x 7.8 mm":              "১৪৭.৬ × ৭১.৬ × ৭.৮ মিলিমিটার",
		"171 g":                              "১৭১ গ্রাম",
		"IP68":                               "IP৬৮",
		"Pink, Yellow, Green, Blue, Black":   "পিঙ্ক, ইয়েলো, গ্রিন, ব্লু, ব্ল্যাক",
		"September 2023":                     "সেপ্টেম্বর ২০২৩",
		"Available":                          "উপলব্ধ",
		"5G":                                 "৫জি",
		"Apple A16 Bionic (4 nm)":            "অ্যাপল A১৬ বায়োনিক (৪ nm)",
		"3349 mAh":                           "৩৩৪৯ মিলিঅ্যাম্পিয়ার আওয়ার",
		"3,349 mAh":                          "৩,৩৪৯ মিলিঅ্যাম্পিয়ার আওয়ার",
		"Li-Ion":                             "লি-আয়ন",
		"HDR, panorama":                      "HDR, প্যানোরামা",
		"2x":                                 "২এক্স",
		"Yes, 50% in 30 min":                 "হ্যাঁ, ৩০ মিনিটে ৫০%",
		"27W wired, PD2.0":                   "২৭ ওয়াট তারযুক্ত, PD২.০",
		"20W wired, MagSafe wireless, Qi wireless":  "২০ ওয়াট তারযুক্ত, MagSafe ওয়্যারলেস, Qi ওয়্যারলেস",
		"MagSafe wireless charging up to 15W":       "MagSafe ওয়্যারলেস চার্জিং পর্যন্ত ১৫ ওয়াট",
		"Stereo speakers":                           "স্টেরিও স্পিকার",
		"Yes, with stereo speakers":                 "হ্যাঁ, স্টেরিও স্পিকার সহ",
		"No":                                        "না",
		"GSM 850 / 900 / 1800 / 1900":               "জিএসএম ৮৫০ / ৯০০ / ১৮০০ / ১৯০০",
		"HSDPA 850 / 900 / 1700(AWS) / 1900 / 2100": "এইচএসডিপিএ ৮৫০ / ৯০০ / ১৭০০(এডব্লিউএস) / ১৯০০ / ২১০০",
		"LTE":                                   "এলটিই",
		"SA/NSA/Sub6/mmWave":                    "এসএ/এনএসএ/সাব৬/এমএমওয়েভ",
		"GSM / CDMA / HSPA / EVDO / LTE / 5G":   "জিএসএম / সিডিএমএ / এইচএসপিএ / ইভিডিও / এলটিই / ৫জি",
		"Wi-Fi 802.11 a/b/g/n/ac/6e, dual-band": "ওয়াই-ফাই ৮০২.১১ এ/বি/জি/এন/এসি/৬ই, ডুয়াল-ব্যান্ড",
		"5.3, A2DP, LE":                         "৫.৩, A2DP, LE",
		"Yes":                                   "হ্যাঁ",
		"Lightning, USB 2.0":                    "লাইটনিং, ইউএসবি ২.০",
		"USB 2.0":                               "ইউএসবি ২.০",
		"GPS, GLONASS, GALILEO, BDS, QZSS":      "জিপিএস, গ্লোনাস, গ্যালিলিও, বিডিএস, কিউজেএসএস",
		"Face ID, accelerometer, gyro, proximity, compass, barometer": "ফেস ID, অ্যাক্সেলেরোমিটার, জাইরো, প্রক্সিমিটি, কম্পাস, ব্যারোমিটার",
		"Face ID, Emergency SOS, Crash Detection":                     "ফেস ID, ইমার্জেন্সি SOS, ক্র্যাশ ডিটেকশন",
		"Face ID, Ceramic Shield, Dynamic Island":                     "ফেস ID, সিরামিক শিল্ড, ডায়নামিক আইল্যান্ড",
		"Nano-SIM and eSIM":                                           "ন্যানো-সিম এবং ই-সিম",
		"Nano SIM and eSIM":                                           "ন্যানো সিম এবং ই-সিম",
		"1.15 W/kg (head) 1.15 W/kg (body)":                           "১.১৫ W/kg (হেড) ১.১৫ W/kg (বডি)",
		"1.19 W/kg (head) 1.20 W/kg (body)":                           "১.১৯ W/kg (হেড) ১.২০ W/kg (বডি)",
		"0.98 W/kg (head) 0.98 W/kg (body)":                           "০.৯৮ W/kg (হেড) ০.৯৮ W/kg (বডি)",
		"0.99 W/kg (head) 0.99 W/kg (body)":                           "০.৯৯ W/kg (হেড) ০.৯৯ W/kg (বডি)",
		"A3090, A3092, A3094, A3096":                                  "A৩০৯০, A৩০৯২, A৩০৯৪, A৩০৯৬",
		"A3090, A2846, A2847, A2848, A2849":                           "A৩০৯০, A২৮৪৬, A২৮৪৭, A২৮৪৮, A২৮৪৯",
	}

	// Merge common translations with iPhone-specific ones
	for key, value := range iPhoneTranslations {
		commonTranslations[key] = value
	}

	return commonTranslations
}

// Seed inserts specification records for the product identified by slug 'iphone-15'
func (s *SpecificationSeederMobileIPhone15) Seed(db *gorm.DB) error {
	productSlug := "iphone-15"
	var prod models.ProductModel
	if err := db.Where("slug = ?", productSlug).First(&prod).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		return err
	}
	productID := prod.ID
	existingkeyMapping := map[string]uint{
		"Brand":                       310,
		"Model Name":                  316,
		"Door Type":                   142,
		"Capacity":                    138,
		"Refrigerator Capacity":       156,
		"Freezer Capacity":            146,
		"Energy Efficiency Rating":    143,
		"Energy Star Rating":          144,
		"Annual Energy Consumption":   137,
		"Product Dimensions":          319,
		"Product Weight":              697,
		"Color":                       311,
		"Compressor Type":             139,
		"Cooling Technology":          698,
		"Defrost Type":                141,
		"Temperature Control":         158,
		"Shelf Material":              699,
		"Number of Shelves":           154,
		"Door Bins":                   700,
		"Crisper Drawers":             701,
		"Ice Maker":                   702,
		"Water Dispenser":             703,
		"Noise Level":                 150,
		"Voltage":                     160,
		"Frequency (Hz)":              704,
		"App Control":                 705,
		"Voice Assistant Support":     385,
		"Warranty Period":             323,
		"Compressor Warranty (Years)": 707,
		"Special Features":            69,
	}
	specs := map[string]string{
		"processor":          "Apple A16 Bionic",
		"resolution":         "1179 x 2556 pixels",
		"screen_protection":  "Ceramic Shield",
		"build_material":     "Aluminum frame, glass front/back",
		"rear_camera":        "48 MP + 12 MP",
		"battery":            "3,349 mAh",
		"battery_type":       "Li-Ion",
		"charging":           "20W wired, MagSafe wireless, Qi wireless",
		"device_status":      "Available",
		"ram":                "6 GB",
		"dimensions":         "147.6 x 71.6 x 7.8 mm",
		"water_resistance":   "IP68",
		"network_technology": "5G",
		"network_bands":      "GSM / CDMA / HSPA / EVDO / LTE / 5G",
		"announcement_date":  "September 2023",
		"chipset":            "Apple A16 Bionic (4 nm)",
		"cpu_type":           "Hexa-core",
		"gpu_type":           "5-core Apple GPU",
		"front_camera":       "12 MP",
		"available_colors":   "Black, Blue, Green, Yellow, Pink",
		"display_size":       "6.1 inches",
		"storage":            "128 GB / 256 GB / 512 GB",
		"display_type":       "Super Retina XDR OLED",
		"refresh_rate":       "60Hz",
		"weight":             "171 g",
		"operating_system":   "iOS 17, upgradable",
		"usb_type":           "USB 2.0",
		"wireless_charging":  "Yes",
		"sensors":            "Face ID, accelerometer, gyro, proximity, compass, barometer",
		"positioning_system": "GPS, GLONASS, GALILEO, BDS, QZSS",
		"special_features":   "Face ID, Ceramic Shield, Dynamic Island",
		"sim_card_type":      "Nano SIM and eSIM",
		"sar_rating":         "1.19 W/kg (head) 1.20 W/kg (body)",
		"sar_rating_eu":      "0.99 W/kg (head) 0.99 W/kg (body)",
		"model_variants":     "A3090, A2846, A2847, A2848, A2849",
	}
	banglaTranslations := s.getBanglaTranslations()
	for key, value := range specs {
		specKeyID, exists := existingkeyMapping[key]
		if !exists {
			log.Printf("⚠️  Key not found in existingkeyMapping: '%s' (used in product: %s)", key, productSlug)
			continue // Skip keys that don't exist in mapping
		}
		var existing models.SpecificationModel
		if err := db.Where("product_id = ? AND specification_key_id = ?", productID, specKeyID).First(&existing).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				sModel := &models.SpecificationModel{
					ProductID:          productID,
					SpecificationKeyID: specKeyID,
					Value:              value,
					Status:             1,
				}
				if err := db.Create(sModel).Error; err != nil {
					return err
				}
				// Create Bangla translation for the specification
				banglaValue, exists := banglaTranslations[value]
				if exists && banglaValue != "" {
					translation := &models.SpecificationTranslationModel{
						SpecificationID: sModel.ID,
						Locale:          "bn",
						Value:           banglaValue,
					}
					if err := db.Create(translation).Error; err != nil {
						return err
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
