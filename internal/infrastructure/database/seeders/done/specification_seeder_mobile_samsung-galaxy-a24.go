package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederMobileSamsungGalaxyA24 seeds specifications/options for product 'samsung-galaxy-a24'
type SpecificationSeederMobileSamsungGalaxyA24 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyA24 creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyA24() *SpecificationSeederMobileSamsungGalaxyA24 {
	return &SpecificationSeederMobileSamsungGalaxyA24{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy A24"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyA24) getBanglaTranslations() map[string]string {
	common := s.GetCommonBanglaTranslations()
	specific := map[string]string{
		"6.5 inches":                    "৬.৫ ইঞ্চি",
		"Super AMOLED":                  "সুপার অ্যামোলেড",
		"2408 × 1080 pixels (~406 ppi)": "২৪০৮ × ১০৮০ পিক্সেল (~৪০৬ পিপিআই)",
		"90 Hz":                         "৯০ হার্জ",
		"Helio G99 (6 nm)":              "হেলিও জি৯৯ (৬ ন্যানোমিটার)",
		"MediaTek Helio G99":            "মিডিয়াটেক হেলিও জি৯৯",
		"Octa-core (2×2.2 GHz Cortex-A76 + 6×2.0 GHz Cortex-A55)": "অক্টা-কোর (২×২.২ গিগাহার্জ কোর্টেক্স-এ৭৬ + ৬×২.০ গিগাহার্জ কোর্টেক্স-এ৫৫)",
		"Mali-G57 MC2":               "মালি জি৫৭ এমসি২",
		"4 / 6 / 8 GB":               "৪ / ৬ / ৮ জিবি",
		"128 GB + microSD":           "১২৮ জিবি + মাইক্রোএসডি",
		"128 GB":                     "১২৮ জিবি",
		"microSDXC (dedicated slot)": "মাইক্রোএসডি (নির্দিষ্ট স্লট)",
		"50 MP + 2 MP + 2 MP":        "৫০ মেগাপিক্সেল + ২ মেগাপিক্সেল + ২ মেগাপিক্সেল",
		"OIS (main)":                 "ওআইএস (মেইন)",
		"50 MP (wide) + 2 MP (macro) + 2 MP (depth)": "৫০ মেগাপিক্সেল (ওয়াইড) + ২ মেগাপিক্সেল (ম্যাক্রো) + ২ মেগাপিক্সেল (ডেপথ)",
		"1080p @ 30fps":                     "১০৮০পি @ ৩০ এফপিএস",
		"None":                              "কোনোটিই নেই",
		"13 MP":                             "১৩ মেগাপিক্সেল",
		"5000 mAh":                          "৫০০০ মিলিঅ্যাম্পিয়ার আওয়ার",
		"Li-Ion (non-removable)":            "লি-আয়ন (অপসারণযোগ্য নয়)",
		"25 W wired":                        "২৫ ওয়াট তারযুক্ত",
		"25W wired":                         "২৫ ওয়াট তারযুক্ত",
		"No":                                "না",
		"Android 13, One UI 5.1":            "অ্যান্ড্রয়েড ১৩, ওয়ান ইউআই ৫.১",
		"Stereo":                            "স্টেরিও",
		"Yes, 3.5 mm":                       "হ্যাঁ, ৩.৫ মিলিমিটার",
		"Glass front, plastic frame / back": "গ্লাস ফ্রন্ট, প্লাস্টিক ফ্রেম / ব্যাক",
		"195 g":                             "১৯৫ গ্রাম",
		"165.4 × 76.9 × 8.4 mm":             "১৬৫.৪ × ৭৬.৯ × ৮.৪ মিলিমিটার",
		"GSM / HSPA / LTE (no 5G)":          "জিএসএম / এইচএসপিএ / এলটিই (৫জি নেই)",
		"Wi-Fi 802.11 a/b/g/n/ac":           "ওয়াই-ফাই ৮০২.১১ এ/বি/জি/এন/এসি",
		"5.3":                               "৫.৩",
		"Yes (market dependent)":            "হ্যাঁ (বাজার-নির্ভর)",
		"USB-C 2.0":                         "ইউএসবি-সি ২.০",
		"GPS, GLONASS, GALILEO, BDS":        "জিপিএস, গ্লোনাস, গ্যালিলিও, বিডিএস",
		"Side-mounted fingerprint, Accelerometer, Gyro, Proximity, Compass": "সাইড-মাউন্টেড ফিঙ্গারপ্রিন্ট, অ্যাকসেলারোমিটার, জাইরো, প্রক্সিমিটি, কম্পাস",
		"Nano-SIM or Dual SIM":                    "ন্যানো-সিম বা ডুয়াল সিম",
		"Eye Care Certification (low blue light)": "আই কেয়ার সার্টিফিকেশন (লো ব্লু লাইট)",
		"Black, Dark Red, Light Green, Silver":    "কালো, ডার্ক রেড, লাইট গ্রিন, সিলভার",
		"April 2023":                              "এপ্রিল ২০২৩",
		"Available":                               "উপলব্ধ",
		"Not specified":                           "উল্লেখ করা হয়নি",
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

// Seed inserts specification records for the product identified by slug 'samsung-galaxy-a24'
func (s *SpecificationSeederMobileSamsungGalaxyA24) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-a24"
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
		"display_size":                  "6.5 inches",
		"display_type":                  "Super AMOLED",
		"resolution":                    "2408 × 1080 pixels (~406 ppi)",
		"refresh_rate":                  "90 Hz",
		"screen_protection":             "Corning Gorilla Glass 5",
		"processor":                     "MediaTek Helio G99",
		"chipset":                       "Helio G99 (6 nm)",
		"cpu_type":                      "Octa-core (2×2.2 GHz Cortex-A76 + 6×2.0 GHz Cortex-A55)",
		"gpu_type":                      "Mali-G57 MC2",
		"processor_speed":               "2.2 GHz",
		"ram":                           "4 / 6 / 8 GB",
		"storage":                       "128 GB + microSD",
		"internal_memory_capacity":      "128 GB",
		"card_slot_type":                "microSDXC (dedicated slot)",
		"rear_camera":                   "50 MP + 2 MP + 2 MP",
		"camera_features":               "OIS (main)",
		"quad_camera_setup":             "50 MP (wide) + 2 MP (macro) + 2 MP (depth)",
		"camera_video_resolution":       "1080p @ 30fps",
		"optical_zoom":                  "None",
		"front_camera":                  "13 MP",
		"front_camera_video_resolution": "1080p @ 30fps",
		"battery":                       "5000 mAh",
		"battery_type":                  "Li-Ion (non-removable)",
		"fast_charging":                 "25 W wired",
		"charging_speed":                "25W wired",
		"wireless_charging":             "No",
		"operating_system":              "Android 13, One UI 5.1",
		"audio_quality":                 "Not specified",
		"loudspeaker_quality":           "Stereo",
		"audio_jack":                    "Yes, 3.5 mm",
		"build_material":                "Glass front, plastic frame / back",
		"weight":                        "195 g",
		"dimensions":                    "165.4 × 76.9 × 8.4 mm",
		"water_resistance":              "None",
		"network_technology":            "GSM / HSPA / LTE (no 5G)",
		"2g_bands":                      "GSM 850 / 900 / 1800 / 1900",
		"3g_bands":                      "HSDPA 850 / 900 / 1700(AWS) / 1900 / 2100",
		"4g_bands":                      "LTE band 1(2100), 2(1900), 3(1800), 4(1700/2100), 5(850), 7(2600), 8(900), 12(700), 13(700), 17(700), 20(800), 26(850), 28(700), 32(1500), 38(2600), 40(2300), 41(2500), 66(1700/2100)",
		"5g_bands":                      "No",
		"5g_support":                    "No",
		"wifi_support":                  "Wi-Fi 802.11 a/b/g/n/ac",
		"bluetooth_version":             "5.3",
		"nfc_support":                   "Yes (market dependent)",
		"usb_type":                      "USB-C 2.0",
		"positioning_system":            "GPS, GLONASS, GALILEO, BDS",
		"sensors":                       "Side-mounted fingerprint, Accelerometer, Gyro, Proximity, Compass",
		"special_features":              "Eye Care Certification (low blue light)",
		"sim_card_type":                 "Nano-SIM or Dual SIM",
		"sar_rating":                    "0.41 W/kg (head) 1.35 W/kg (body)",
		"sar_rating_eu":                 "0.41 W/kg (head) 1.35 W/kg (body)",
		"available_colors":              "Black, Dark Red, Light Green, Silver",
		"model_variants":                "SM-A245F, SM-A245M, SM-A245N",
		"announcement_date":             "April 2023",
		"device_status":                 "Available",
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
