package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederMobileSamsungGalaxyM365G seeds specifications/options for product 'samsung-galaxy-m36-5g'
type SpecificationSeederMobileSamsungGalaxyM365G struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyM365G creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyM365G() *SpecificationSeederMobileSamsungGalaxyM365G {
	return &SpecificationSeederMobileSamsungGalaxyM365G{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy M36 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyM365G) getBanglaTranslations() map[string]string {
	common := s.GetCommonBanglaTranslations()
	specific := map[string]string{
		"6.7 inches":            "৬.৭ ইঞ্চি",
		"Super AMOLED, 120 Hz":  "সুপার অ্যামোলেড, ১২০ হার্জ",
		"50 MP + 8 MP + 2 MP":   "৫০ মেগাপিক্সেল + ৮ মেগাপিক্সেল + ২ মেগাপিক্সেল",
		"Android 15 (One UI 7)": "অ্যান্ড্রয়েড ১৫ (ওয়ান UI ৭)",
		"Side fingerprint, Accelerometer, Gyro, Proximity, Compass": "সাইড ফিঙ্গারপ্রিন্ট, অ্যাক্সেলেরোমিটার, গাইরো, প্রক্সিমিটি, কম্পাস",
		"Exynos 1380":                           "এক্সিনস ১৩৮০",
		"Mali-G68 MP5":                          "মালি-জি৬৮ এমপি৫",
		"6 / 8 GB":                              "৬ / ৮ জিবি",
		"5000 mAh":                              "৫০০০ মিলিঅ্যাম্পিয়ার আওয়ার",
		"Li-Po (non-removable)":                 "লি-পো (অপসারণযোগ্য নয়)",
		"128 / 256 GB + microSD":                "১২৮ / ২৫৬ জিবি + মাইক্রোএসডি",
		"222 g":                                 "২২২ গ্রাম",
		"162.3 × 78.6 × 9.1 mm":                 "১৬২.৩ × ৭৮.৬ × ৯.১ মিলিমিটার",
		"OIS, LED flash":                        "ওআইএস, এলইডি ফ্ল্যাশ",
		"Stereo":                                "স্টেরিও",
		"Available":                             "উপলব্ধ",
		"Exynos 1380 (5 nm)":                    "এক্সিনস ১৩৮০ (৫ ন্যানোমিটার)",
		"1080 × 2340 px":                        "১০৮০ × ২৩৪০ পিক্সেল",
		"120 Hz":                                "১২০ হার্জ",
		"Wi-Fi 802.11 a/b/g/n/ac":               "ওয়াই-ফাই ৮০২.১১ এ/বি/জি/এন/এসি",
		"4K @ 30fps; 1080p @ 30fps":             "৪কে @ ৩০ এফপিএস; ১০৮০পি @ ৩০ এফপিএস",
		"25 W wired":                            "২৫ W তারযুক্ত",
		"6 years of Android / Security updates": "৬ বছরের অ্যান্ড্রয়েড / সিকিউরিটি আপডেট",
		"Octa-core (4×2.4 GHz Cortex-A78 + 4×2.0 GHz Cortex-A55)": "অক্টা-কোর (৪×২.৪ গিগাহার্টজ Cortex-A78 + ৪×২.০ গিগাহার্টজ Cortex-A55)",
		"5.3":  "৫.৩",
		"None": "কোনটি না",
		"Thunder Grey, DayBreak Blue, Moonlight Blue":                    "থান্ডার গ্রে, ডেব্রেক ব্লু, মুনলাইট ব্লু",
		"Glass front (Gorilla Glass Victus+), plastic frame, glass back": "সামনে গ্লাস (গরিলা গ্লাস ভিক্টাস+), প্লাস্টিক ফ্রেম, গ্লাস ব্যাক",
		"Yes":                           "হ্যাঁ",
		"13 MP":                         "১৩ MP",
		"No":                            "না",
		"Nano-SIM + Nano-SIM":           "ন্যানো-সিম + ন্যানো-সিম",
		"July 2025":                     "জুলাই ২০২৫",
		"GSM / HSPA / LTE / 5G":         "জিএসএম / এইচএসপিএ / এলটিই / ৫জি",
		"USB-C 2.0":                     "ইউএসবি-সি ২.০",
		"4K @ 30fps; 1080p @ 30/60fps":  "৪কে @ ৩০ এফপিএস; ১০৮০পি @ ৩০/৬০ এফপিএস",
		"GSM 850 / 900 / 1800 / 1900":   "জিএসএম ৮৫০ / ৯০০ / ১৮০০ / ১৯০০",
		"HSDPA 850 / 900 / 1900 / 2100": "এইচএসডিপিএ ৮৫০ / ৯০০ / ১৯০০ / ২১০০",
		"LTE":                           "এলটিই",
		"SA/NSA/Sub6":                   "এসএ/এনএসএ/সাব৬",
		"128 / 256 GB":                  "১২৮ / ২৫৬ GB",
		"microSDXC (dedicated slot)":    "মাইক্রো এসডিএক্সসি (ডেডিকেটেড স্লট)",
		"2.4 GHz":                       "২.৪ GHz",
		"25 W":                          "২৫ W",
		"Gorilla Glass Victus+":         "গরিলা গ্লাস ভিক্টাস+",
		"25W":                           "২৫ওয়াট",
		"Up to 2.4 GHz":                 "২.৪ গিগাহার্টজ পর্যন্ত",
		"128GB 6GB RAM, 128GB 8GB RAM, 256GB 8GB RAM": "১২৮জিবি ৬জিবি র্যাম, ১২৮জিবি ৮জিবি র্যাম, ২৫৬জিবি ৮জিবি র্যাম",
		"50 MP, f/1.8, (wide), 1/2.76\", 0.64µm, PDAF\n8 MP, f/2.2, (ultrawide), 1/4.0\", 1.12µm\n2 MP, f/2.4, (macro)": "৫০ এমপি, এফ/১.৮, (ওয়াইড), ১/২.৭৬\", ০.৬৪মাইক্রোমিটার, পিডিএএফ\n৮ এমপি, এফ/২.২, (আল্ট্রাওয়াইড), ১/৪.০\", ১.১২মাইক্রোমিটার\n২ এমপি, এফ/২.৪, (ম্যাক্রো)",
	}
	for k, v := range specific {
		common[k] = v
	}
	return common
}

// Seed inserts specification records for the product identified by slug 'samsung-galaxy-m36-5g'
func (s *SpecificationSeederMobileSamsungGalaxyM365G) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-m36-5g"
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
		"display_size":                  "6.7 inches",
		"display_type":                  "Super AMOLED, 120 Hz",
		"rear_camera":                   "50 MP + 8 MP + 2 MP",
		"operating_system":              "Android 15 (One UI 7)",
		"sensors":                       "Side fingerprint, Accelerometer, Gyro, Proximity, Compass",
		"processor":                     "Exynos 1380",
		"gpu_type":                      "Mali-G68 MP5",
		"ram":                           "6 / 8 GB",
		"battery":                       "5000 mAh",
		"battery_type":                  "Li-Po (non-removable)",
		"storage":                       "128 / 256 GB + microSD",
		"weight":                        "222 g",
		"dimensions":                    "162.3 × 78.6 × 9.1 mm",
		"camera_features":               "OIS, LED flash",
		"loudspeaker_quality":           "Stereo",
		"device_status":                 "Available",
		"chipset":                       "Exynos 1380 (5 nm)",
		"resolution":                    "1080 × 2340 px",
		"refresh_rate":                  "120 Hz",
		"positioning_system":            "",
		"wifi_support":                  "Wi-Fi 802.11 a/b/g/n/ac",
		"front_camera_video_resolution": "4K @ 30fps; 1080p @ 30fps",
		"fast_charging":                 "25 W wired",
		"special_features":              "6 years of Android / Security updates",
		"audio_quality":                 "",
		"cpu_type":                      "Octa-core (4×2.4 GHz Cortex-A78 + 4×2.0 GHz Cortex-A55)",
		"bluetooth_version":             "5.3",
		"optical_zoom":                  "None",
		"available_colors":              "Thunder Grey, DayBreak Blue, Moonlight Blue",
		"build_material":                "Glass front (Gorilla Glass Victus+), plastic frame, glass back",
		"nfc_support":                   "Yes",
		"front_camera":                  "13 MP",
		"wireless_charging":             "No",
		"5g_support":                    "Yes",
		"sim_card_type":                 "Nano-SIM + Nano-SIM",
		"announcement_date":             "July 2025",
		"water_resistance":              "",
		"network_technology":            "GSM / HSPA / LTE / 5G",
		"2g_bands":                      "GSM 850 / 900 / 1800 / 1900",
		"3g_bands":                      "HSDPA 850 / 900 / 1900 / 2100",
		"4g_bands":                      "LTE",
		"5g_bands":                      "SA/NSA/Sub6",
		"usb_type":                      "USB-C 2.0",
		"camera_video_resolution":       "4K @ 30fps; 1080p @ 30/60fps",
		"audio_jack":                    "No",
		"internal_memory_capacity":      "128 / 256 GB",
		"card_slot_type":                "microSDXC (dedicated slot)",
		"processor_speed":               "2.4 GHz",
		"charging_speed":                "25 W",
		"screen_protection":             "Gorilla Glass Victus+",
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
