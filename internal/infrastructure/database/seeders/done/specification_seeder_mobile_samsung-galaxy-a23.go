package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederMobileSamsungGalaxyA23 seeds specifications/options for product 'samsung-galaxy-a23'
type SpecificationSeederMobileSamsungGalaxyA23 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyA23 creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyA23() *SpecificationSeederMobileSamsungGalaxyA23 {
	return &SpecificationSeederMobileSamsungGalaxyA23{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy A23"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyA23) getBanglaTranslations() map[string]string {
	common := s.GetCommonBanglaTranslations()
	specific := map[string]string{
		"6.6 inches":                        "৬.৬ ইঞ্চি",
		"PLS LCD":                           "পিএলএস LCD",
		"2408 × 1080 px (~400 ppi)":         "২৪০৮ × ১০৮০ px (~৪০০ ppi)",
		"90 Hz":                             "৯০ Hz",
		"Corning Gorilla Glass 5":           "কর্নিং গরিলা গ্লাস ৫",
		"Snapdragon 680 (4G)":               "স্ন্যাপড্রাগন ৬৮০ (৪G)",
		"Snapdragon 680 (6 nm)":             "স্ন্যাপড্রাগন ৬৮০ (৬ ন্যানোমিটার)",
		"Octa-core (4×2.4 GHz + 4×1.9 GHz)": "অক্টা-কোর (৪×২.৪ GHz + ৪×১.৯ GHz)",
		"Adreno 610":                        "অ্যাড্রেনো ৬১০",
		"2.4 GHz":                           "২.৪ GHz",
		"4 / 6 / 8 GB":                      "৪ / ৬ / ৮ GB",
		"64 / 128 GB + microSD":             "৬৪ / ১২৮ GB + মাইক্রোSD",
		"64 / 128 GB":                       "৬৪ / ১২৮ GB",
		"microSDXC (dedicated slot)":        "মাইক্রোSDXC (ডেডিকেটেড স্লট)",
		"50 MP + 5 MP + 2 MP + 2 MP":        "৫০ MP + ৫ MP + ২ MP + ২ MP",
		"OIS (main), LED flash":             "OIS (মেইন), LED ফ্ল্যাশ",
		"50 MP (wide) + 5 MP (ultrawide) + 2 MP (macro) + 2 MP (depth)": "৫০ MP (ওয়াইড) + ৫ MP (আল্ট্রাওয়াইড) + ২ MP (ম্যাক্রো) + ২ MP (ডেপ্থ)",
		"1080p @ 30fps":                       "১০৮০p @ ৩০fps",
		"None":                                "কোনটি নয়",
		"8 MP":                                "৮ MP",
		"5000 mAh":                            "৫০০০ mAh",
		"Li-Ion (non-removable)":              "লি-আয়ন (অপসারণযোগ্য নয়)",
		"25 W wired":                          "২৫ W তারযুক্ত",
		"25W wired":                           "২৫W তারযুক্ত",
		"No":                                  "না",
		"Android 12, One UI 4.1 (upgradable)": "অ্যান্ড্রয়েড ১২, ওয়ান UI ৪.১ (আপগ্রেডযোগ্য)",
		"Dolby Atmos":                         "ডলবি অ্যাটমস",
		"Mono / Stereo depending on variant":  "মনো / স্টেরিও ভ্যারিয়েন্ট অনুসারে",
		"3.5 mm headphone jack":               "৩.৫ mm হেডফোন জ্যাক",
		"Glass front (Gorilla Glass 5), plastic frame/back": "গ্লাস ফ্রন্ট (গরিলা গ্লাস ৫), প্লাস্টিক ফ্রেম/ব্যাক",
		"195 g":                                     "১৯৫ গ্রাম",
		"165.4 × 76.9 × 8.4 mm":                     "১৬৫.৪ × ৭৬.৯ × ৮.৪ mm",
		"GSM / HSPA / LTE (4G)":                     "জিএসএম / এইচএসপিএ / এলটিই (৪G)",
		"GSM 850 / 900 / 1800 / 1900":               "জিএসএম ৮৫০ / ৯০০ / ১৮০০ / ১৯০০",
		"HSDPA 850 / 900 / 1700(AWS) / 1900 / 2100": "এইচএসডিপিএ ৮৫০ / ৯০০ / ১৭০০(AWS) / ১৯০০ / ২১০০",
		"LTE band 1(2100), 2(1900), 3(1800), 4(1700/2100), 5(850), 7(2600), 8(900), 12(700), 13(700), 17(700), 20(800), 26(850), 28(700), 32(1500), 38(2600), 40(2300), 41(2500), 66(1700/2100)": "এলটিই ব্যান্ড ১(২১০০), ২(১৯০০), ৩(১৮০০), ৪(১৭০০/২১০০), ৫(৮৫০), ৭(২৬০০), ৮(৯০০), ১২(৭০০), ১৩(৭০০), ১৭(৭০০), ২০(৮০০), ২৬(৮৫০), ২৮(৭০০), ৩২(১৫০০), ৩৮(২৬০০), ৪০(২৩০০), ৪১(২৫০০), ৬৬(১৭০০/২১০০)",
		"Some variants (A23 5G), base A23 is 4G": "কিছু ভ্যারিয়েন্ট (A২৩ ৫G), বেস A২৩ হল ৪G",
		"Wi-Fi 802.11 a/b/g/n/ac":                "ওয়াই-ফাই ৮০২.১১ a/b/g/n/ac",
		"5.0":                                    "৫.০",
		"Yes (market dependent)":                 "হ্যাঁ (মার্কেট নির্ভর)",
		"USB-C 2.0":                              "USB-C ২.০",
		"GPS, GLONASS, GALILEO, BDS":             "GPS, GLONASS, GALILEO, BDS",
		"Side fingerprint, Accelerometer, Gyro, Proximity, Compass": "সাইড ফিঙ্গারপ্রিন্ট, অ্যাকসেলেরোমিটার, জাইরো, প্রক্সিমিটি, কম্পাস",
		"Nano-SIM or Dual SIM":              "ন্যানো-SIM বা ডুয়াল SIM",
		"0.41 W/kg (head) 1.35 W/kg (body)": "০.৪১ W/kg (হেড) ১.৩৫ W/kg (বডি)",
		"Black, White, Peach, Blue":         "ব্ল্যাক, হোয়াইট, পিচ, ব্লু",
		"SM-A235F, SM-A235M, SM-A235N":      "SM-A235F, SM-A235M, SM-A235N",
		"March 2022":                        "মার্চ ২০২২",
		"Available":                         "উপলব্ধ",
	}
	result := make(map[string]string)
	for k, v := range common {
		result[k] = v
	}
	for k, v := range specific {
		result[k] = v
	}
	return result
}

// Seed inserts specification records for the product identified by slug 'samsung-galaxy-a23'
func (s *SpecificationSeederMobileSamsungGalaxyA23) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-a23"
	var prod models.ProductModel
	if err := db.Where("slug = ?", productSlug).First(&prod).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		return err
	}
	productID := prod.ID
	existingkeyMapping := map[string]uint{
		// Core display
		"display_size":      66, // Screen Size
		"display_type":      27, // Display Type
		"resolution":        62, // Resolution
		"refresh_rate":      61, // Refresh Rate
		"screen_protection": 65, // Screen Protection
		// Performance
		"processor":       535, // processor (lowercase)
		"chipset":         19,  // Chipset
		"cpu_type":        21,  // CPU Type
		"gpu_type":        33,  // GPU Type
		"processor_speed": 57,  // Processor Speed
		// Memory & storage
		"ram":                      60, // RAM
		"storage":                  71, // Storage Capacity
		"internal_memory_capacity": 36, // Internal Memory Capacity
		"card_slot_type":           17, // Card Slot Type
		// Cameras
		"rear_camera":                   122, // Rear Camera
		"camera_features":               16,  // Camera Features
		"quad_camera_setup":             58,  // Quad Camera Setup
		"camera_video_resolution":       40,  // Main Camera Video Resolution
		"optical_zoom":                  193, // Optical Zoom
		"front_camera":                  32,  // Front Camera
		"front_camera_video_resolution": 32,  // Front Camera Video Resolution
		// Battery & charging
		"battery":           11,  // Battery Capacity
		"battery_type":      12,  // Battery Type
		"fast_charging":     553, // fast_charging (lowercase)
		"charging_speed":    18,  // Charging Speed
		"wireless_charging": 82,  // Wireless Charging
		// Software
		"operating_system": 49, // Operating System
		// Audio
		"audio_quality":       9,  // Audio Quality
		"loudspeaker_quality": 38, // Loudspeaker Quality
		"audio_jack":          2,  // 3.5mm Audio Jack
		// Build & body
		"build_material":   14, // Build Material
		"weight":           80, // Weight
		"dimensions":       25, // Dimensions
		"water_resistance": 78, // Water Resistance
		// Network & connectivity
		"network_technology": 653, // Network Technology
		"2g_bands":           1,   // 2G Bands
		"3g_bands":           3,   // 3G Bands
		"4g_bands":           4,   // 4G Bands
		"5g_bands":           5,   // 5G Bands
		"5g_support":         5,   // 5G Support (maps to 5G Bands)
		"wifi_support":       81,  // WiFi
		"bluetooth_version":  13,  // Bluetooth Version
		"nfc_support":        46,  // NFC Support
		"usb_type":           76,  // USB Type
		// Positioning & sensors
		"positioning_system": 54, // Positioning System
		"sensors":            67, // Sensors
		"special_features":   69, // Special Features
		// SIM & compliance
		"sim_card_type": 68, // SIM Card Type
		"sar_rating":    63, // SAR (Specific Absorption Rate)
		"sar_rating_eu": 64, // SAR (Specific Absorption Rate) EU
		// Commercial info
		"available_colors":  10, // Available Colors
		"model_variants":    42, // Model Variants
		"announcement_date": 8,  // Announcement Date
		"device_status":     23, // Device Status
	}
	specs := map[string]string{
		// Core display
		"display_size":      "6.6 inches",
		"display_type":      "PLS LCD",
		"resolution":        "2408 × 1080 px (~400 ppi)",
		"refresh_rate":      "90 Hz",
		"screen_protection": "Corning Gorilla Glass 5",
		// Performance
		"processor":       "Snapdragon 680 (4G)",
		"chipset":         "Snapdragon 680 (6 nm)",
		"cpu_type":        "Octa-core (4×2.4 GHz + 4×1.9 GHz)",
		"gpu_type":        "Adreno 610",
		"processor_speed": "2.4 GHz",
		// Memory & storage
		"ram":                      "4 / 6 / 8 GB",
		"storage":                  "64 / 128 GB + microSD",
		"internal_memory_capacity": "64 / 128 GB",
		"card_slot_type":           "microSDXC (dedicated slot)",
		// Cameras
		"rear_camera":                   "50 MP + 5 MP + 2 MP + 2 MP",
		"camera_features":               "OIS (main), LED flash",
		"quad_camera_setup":             "50 MP (wide) + 5 MP (ultrawide) + 2 MP (macro) + 2 MP (depth)",
		"camera_video_resolution":       "1080p @ 30fps",
		"optical_zoom":                  "None",
		"front_camera":                  "8 MP",
		"front_camera_video_resolution": "1080p @ 30fps",
		// Battery & charging
		"battery":           "5000 mAh",
		"battery_type":      "Li-Ion (non-removable)",
		"fast_charging":     "25 W wired",
		"charging_speed":    "25W wired",
		"wireless_charging": "No",
		// Software
		"operating_system": "Android 12, One UI 4.1 (upgradable)",
		// Audio
		"audio_quality":       "Dolby Atmos",
		"loudspeaker_quality": "Mono / Stereo depending on variant",
		"audio_jack":          "3.5 mm headphone jack",
		// Build & body
		"build_material":   "Glass front (Gorilla Glass 5), plastic frame/back",
		"weight":           "195 g",
		"dimensions":       "165.4 × 76.9 × 8.4 mm",
		"water_resistance": "None",
		// Network & connectivity
		"network_technology": "GSM / HSPA / LTE (4G)",
		"2g_bands":           "GSM 850 / 900 / 1800 / 1900",
		"3g_bands":           "HSDPA 850 / 900 / 1700(AWS) / 1900 / 2100",
		"4g_bands":           "LTE band 1(2100), 2(1900), 3(1800), 4(1700/2100), 5(850), 7(2600), 8(900), 12(700), 13(700), 17(700), 20(800), 26(850), 28(700), 32(1500), 38(2600), 40(2300), 41(2500), 66(1700/2100)",
		"5g_bands":           "No",
		"5g_support":         "Some variants (A23 5G), base A23 is 4G",
		"wifi_support":       "Wi-Fi 802.11 a/b/g/n/ac",
		"bluetooth_version":  "5.0",
		"nfc_support":        "Yes (market dependent)",
		"usb_type":           "USB-C 2.0",
		// Positioning & sensors
		"positioning_system": "GPS, GLONASS, GALILEO, BDS",
		"sensors":            "Side fingerprint, Accelerometer, Gyro, Proximity, Compass",
		"special_features":   "Dolby Atmos",
		// SIM & compliance
		"sim_card_type": "Nano-SIM or Dual SIM",
		"sar_rating":    "0.41 W/kg (head) 1.35 W/kg (body)",
		"sar_rating_eu": "0.41 W/kg (head) 1.35 W/kg (body)",
		// Commercial info
		"available_colors":  "Black, White, Peach, Blue",
		"model_variants":    "SM-A235F, SM-A235M, SM-A235N",
		"announcement_date": "March 2022",
		"device_status":     "Available",
	}
	banglaTranslations := s.getBanglaTranslations()
	for key, value := range specs {
		specKeyID, exists := existingkeyMapping[key]
		if !exists {
			log.Printf("⚠️  Key not found in existingkeyMapping: '%!s(MISSING)' (used in product: %!s(MISSING))", key, productSlug)
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
