package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederMobileSamsungGalaxyM335G seeds specifications/options for product 'samsung-galaxy-m33-5g'
type SpecificationSeederMobileSamsungGalaxyM335G struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyM335G creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyM335G() *SpecificationSeederMobileSamsungGalaxyM335G {
	return &SpecificationSeederMobileSamsungGalaxyM335G{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy M33 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyM335G) getBanglaTranslations() map[string]string {
	common := s.GetCommonBanglaTranslations()
	specific := map[string]string{
		"PowerCool Tech":                   "পাওয়ারকুল টেক",
		"50 MP + 5 MP + 2 MP + 2 MP":       "৫০ MP + ৫ MP + ২ MP + ২ MP",
		"5000 mAh (some markets 6000 mAh)": "৫০০০ mAh (কিছু মার্কেটে ৬০০০ mAh)",
		"Nano-SIM + Nano-SIM":              "ন্যানো-সিম + ন্যানো-সিম",
		"Mali-G68 MP4":                     "মালি-G৬৮ MP৪",
		"120 Hz":                           "১২০ Hz",
		"Glass front (Gorilla Glass 5), plastic frame / back": "সামনে গ্লাস (গরিলা গ্লাস ৫), ফ্রেম / পিছনে প্লাস্টিক",
		"GSM / HSPA / LTE / 5G":                               "জিএসএম / এইচএসপিএ / এলটিই / ৫জি",
		"128 GB + microSD":                                    "১২৮ GB + microSD",
		"Android 12 (upgradable)":                             "অ্যান্ড্রয়েড ১২ (আপগ্রেডযোগ্য)",
		"3.5 mm":                                              "৩.৫ mm",
		"165.4 × 76.9 × 9.4 mm":                               "১৬৫.৪ × ৭৬.৯ × ৯.৪ mm",
		"2408 × 1080 px":                                      "২৪০৮ × ১০৮০ px",
		"Wi-Fi 802.11 a/b/g/n/ac, dual-band":                  "ওয়াই-ফাই ৮০২.১১ a/b/g/n/ac, ডুয়াল-ব্যান্ড",
		"April 2022":                                          "এপ্রিল ২০২২",
		"Available":                                           "উপলব্ধ",
		"Exynos 1280":                                         "এক্সিনস ১২৮০",
		"USB-C 2.0":                                           "ইউএসবি-সি ২.০",
		"4K @ 30fps; 1080p @ 30/60/720fps":                    "৪কে @ ৩০fps; ১০৮০p @ ৩০/৬০/৭২০fps",
		"25 W wired":                                          "২৫ W তারযুক্ত",
		"Side fingerprint, Accelerometer, Gyro, Proximity, Compass": "সাইড ফিঙ্গারপ্রিন্ট, অ্যাক্সেলেরোমিটার, গাইরো, প্রক্সিমিটি, কম্পাস",
		"198 g":                     "১৯৮ g",
		"6.6 inches":                "৬.৬ ইঞ্চি",
		"Exynos 1280 (5 nm)":        "এক্সিনস ১২৮০ (৫ nm)",
		"PDAF, LED flash, HDR":      "পিডিএএফ, এলইডি ফ্ল্যাশ, এইচডিআর",
		"8 MP":                      "৮ MP",
		"4K @ 30fps; 1080p @ 30fps": "৪কে @ ৩০fps; ১০৮০p @ ৩০fps",
		"Li-Ion (non-removable)":    "লি-আয়ন (অপসারণযোগ্য নয়)",
		"No":                        "না",
		"6 / 8 GB":                  "৬ / ৮ GB",
		"5.1":                       "৫.১",
		"None":                      "কোনটি না",
		"Yes":                       "হ্যাঁ",
		"Mono":                      "মনো",
		"Deep Ocean Blue, Mystique Green, Emerald Brown": "ডীপ ওশান ব্লু, মিস্টিক গ্রীন, এমেরাল্ড ব্রাউন",
		"Octa-core (2×2.4 GHz + 6×2.0 GHz)":              "অক্টা-কোর (২×২.৪ GHz + ৬×২.০ GHz)",
		"TFT LCD, 120 Hz":                                "টিএফটি LCD, ১২০ Hz",
		"GPS, GLONASS, GALILEO, BDS":                     "জিপিএস, গ্লোনাস, গ্যালিলিও, বিডিএস",
		"GSM 850 / 900 / 1800 / 1900":                    "জিএসএম ৮৫০ / ৯০০ / ১৮০০ / ১৯০০",
		"HSDPA 850 / 900 / 1900 / 2100":                  "এইচএসডিপিএ ৮৫০ / ৯০০ / ১৯০০ / ২১০০",
		"LTE":                                            "এলটিই",
		"SA/NSA/Sub6":                                    "এসএ/এনএসএ/সাব৬",
		"128 GB":                                         "১২৮ GB",
		"microSDXC (dedicated slot)":                     "মাইক্রো এসডিএক্সসি (ডেডিকেটেড স্লট)",
		"2.4 GHz":                                        "২.৪ GHz",
		"25 W":                                           "২৫ W",
		"Gorilla Glass 5":                                "গরিলা গ্লাস ৫",
	}
	for k, v := range specific {
		common[k] = v
	}
	return common
}

// Seed inserts specification records for the product identified by slug 'samsung-galaxy-m33-5g'
func (s *SpecificationSeederMobileSamsungGalaxyM335G) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-m33-5g"
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
		"special_features":              "PowerCool Tech",
		"audio_quality":                 "",
		"rear_camera":                   "50 MP + 5 MP + 2 MP + 2 MP",
		"battery":                       "5000 mAh (some markets 6000 mAh)",
		"sim_card_type":                 "Nano-SIM + Nano-SIM",
		"gpu_type":                      "Mali-G68 MP4",
		"refresh_rate":                  "120 Hz",
		"build_material":                "Glass front (Gorilla Glass 5), plastic frame / back",
		"network_technology":            "GSM / HSPA / LTE / 5G",
		"2g_bands":                      "GSM 850 / 900 / 1800 / 1900",
		"3g_bands":                      "HSDPA 850 / 900 / 1900 / 2100",
		"4g_bands":                      "LTE",
		"5g_bands":                      "SA/NSA/Sub6",
		"storage":                       "128 GB + microSD",
		"operating_system":              "Android 12 (upgradable)",
		"audio_jack":                    "3.5 mm",
		"dimensions":                    "165.4 × 76.9 × 9.4 mm",
		"water_resistance":              "",
		"resolution":                    "2408 × 1080 px",
		"wifi_support":                  "Wi-Fi 802.11 a/b/g/n/ac, dual-band",
		"announcement_date":             "April 2022",
		"device_status":                 "Available",
		"processor":                     "Exynos 1280",
		"nfc_support":                   "",
		"usb_type":                      "USB-C 2.0",
		"camera_video_resolution":       "4K @ 30fps; 1080p @ 30/60/720fps",
		"fast_charging":                 "25 W wired",
		"sensors":                       "Side fingerprint, Accelerometer, Gyro, Proximity, Compass",
		"weight":                        "198 g",
		"display_size":                  "6.6 inches",
		"chipset":                       "Exynos 1280 (5 nm)",
		"camera_features":               "PDAF, LED flash, HDR",
		"front_camera":                  "8 MP",
		"front_camera_video_resolution": "4K @ 30fps; 1080p @ 30fps",
		"battery_type":                  "Li-Ion (non-removable)",
		"wireless_charging":             "No",
		"ram":                           "6 / 8 GB",
		"bluetooth_version":             "5.1",
		"optical_zoom":                  "None",
		"5g_support":                    "Yes",
		"loudspeaker_quality":           "Mono",
		"available_colors":              "Deep Ocean Blue, Mystique Green, Emerald Brown",
		"cpu_type":                      "Octa-core (2×2.4 GHz + 6×2.0 GHz)",
		"display_type":                  "TFT LCD, 120 Hz",
		"positioning_system":            "GPS, GLONASS, GALILEO, BDS",
		"internal_memory_capacity":      "128 GB",
		"card_slot_type":                "microSDXC (dedicated slot)",
		"processor_speed":               "2.4 GHz",
		"charging_speed":                "25 W",
		"screen_protection":             "Gorilla Glass 5",
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
