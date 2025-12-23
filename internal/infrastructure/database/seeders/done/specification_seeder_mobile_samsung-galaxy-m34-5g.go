package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederMobileSamsungGalaxyM345G seeds specifications/options for product 'samsung-galaxy-m34-5g'
type SpecificationSeederMobileSamsungGalaxyM345G struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyM345G creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyM345G() *SpecificationSeederMobileSamsungGalaxyM345G {
	return &SpecificationSeederMobileSamsungGalaxyM345G{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy M34 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyM345G) getBanglaTranslations() map[string]string {
	common := s.GetCommonBanglaTranslations()
	specific := map[string]string{
		"13 MP":                            "১৩ মেগাপিক্সেল",
		"July 2023":                        "জুলাই ২০২৩",
		"6.6 inches":                       "৬.৬ ইঞ্চি",
		"Exynos 1280":                      "এক্সিনস ১২৮০",
		"None":                             "কোনটি না",
		"208 g":                            "২০৮ গ্রাম",
		"25 W wired":                       "২৫ W তারযুক্ত",
		"Mono / Stereo (depends)":          "মনো / স্টেরিও (নির্ভর করে)",
		"Exynos 1280 (5 nm)":               "এক্সিনস ১২৮০ (৫ ন্যানোমিটার)",
		"161.7 × 77.2 × 8.8 mm":            "১৬১.৭ × ৭৭.২ × ৮.৮ মিলিমিটার",
		"Mali-G68 MP4":                     "মালি-জি৬৮ এমপি৪",
		"GSM / HSPA / LTE / 5G":            "জিএসএম / এইচএসপিএ / এলটিই / ৫জি",
		"Hybrid Dual SIM (nano + microSD)": "হাইব্রিড ডুয়াল সিম (ন্যানো + মাইক্রোএসডি)",
		"Midnight Blue, Prism Silver, Waterfall Blue": "মিডনাইট নীল, প্রিজম রৌপ্য, ওয়াটারফল নীল",
		"Available":                         "উপলব্ধ",
		"Super AMOLED, 120 Hz":              "সুপার অ্যামোলেড, ১২০ হার্জ",
		"2340 × 1080 px":                    "২৩৪০ × ১০৮০ পিক্সেল",
		"USB-C 2.0":                         "ইউএসবি-সি ২.০",
		"Android 13, One UI 5 (upgradable)": "অ্যান্ড্রয়েড ১৩, ওয়ান UI ৫ (আপগ্রেডযোগ্য)",
		"Octa-core (2×2.4 GHz + 6×2.0 GHz)": "অক্টা-কোর (২×২.৪ GHz + ৬×২.০ GHz)",
		"6 / 8 GB":                          "৬ / ৮ জিবি",
		"50 MP + 8 MP + 2 MP":               "৫০ মেগাপিক্সেল + ৮ মেগাপিক্সেল + ২ মেগাপিক্সেল",
		"6000 mAh":                          "৬০০০ মিলিঅ্যাম্পিয়ার আওয়ার",
		"No":                                "না",
		"Yes":                               "হ্যাঁ",
		"Accelerometer, Gyro, Proximity, Compass, fingerprint (side)": "অ্যাক্সেলেরোমিটার, গাইরো, প্রক্সিমিটি, কম্পাস, ফিঙ্গারপ্রিন্ট (সাইড)",
		"128 / 256 GB + microSD (up to 1TB)":                          "১২৮ / ২৫৬ জিবি + মাইক্রোএসডি (আপ টু ১টিবি)",
		"120 Hz":                                                      "১২০ হার্জ",
		"Li-Po (non-removable)":                                       "লি-পো (অপসারণযোগ্য নয়)",
		"Glass front (Gorilla Glass 5), plastic frame / back":         "সামনে গ্লাস (গরিলা গ্লাস ৫), ফ্রেম / পিছনে প্লাস্টিক",
		"Wi-Fi 802.11 a/b/g/n/ac":                                     "ওয়াই-ফাই ৮০২.১১ এ/বি/জি/এন/এসি",
		"OIS (main), LED flash":                                       "ওআইএস (মেইন), এলইডি ফ্ল্যাশ",
		"GSM 850 / 900 / 1800 / 1900":                                 "জিএসএম ৮৫০ / ৯০০ / ১৮০০ / ১৯০০",
		"HSDPA 850 / 900 / 1900 / 2100":                               "এইচএসডিপিএ ৮৫০ / ৯০০ / ১৯০০ / ২১০০",
		"LTE":                                                         "এলটিই",
		"SA/NSA/Sub6":                                                 "এসএ/এনএসএ/সাব৬",
		"128 / 256 GB":                                                "১২৮ / ২৫৬ GB",
		"microSDXC (hybrid)":                                          "মাইক্রোএসডিএক্সসি (হাইব্রিড)",
		"2.4 GHz":                                                     "২.৪ GHz",
		"25 W":                                                        "২৫ W",
		"Gorilla Glass 5":                                             "গরিলা গ্লাস ৫",
	}
	for k, v := range specific {
		common[k] = v
	}
	return common
}

// Seed inserts specification records for the product identified by slug 'samsung-galaxy-m34-5g'
func (s *SpecificationSeederMobileSamsungGalaxyM345G) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-m34-5g"
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
		"camera_video_resolution":       "",
		"front_camera":                  "13 MP",
		"announcement_date":             "July 2023",
		"display_size":                  "6.6 inches",
		"processor":                     "Exynos 1280",
		"optical_zoom":                  "None",
		"weight":                        "208 g",
		"fast_charging":                 "25 W wired",
		"loudspeaker_quality":           "Mono / Stereo (depends)",
		"chipset":                       "Exynos 1280 (5 nm)",
		"dimensions":                    "161.7 × 77.2 × 8.8 mm",
		"gpu_type":                      "Mali-G68 MP4",
		"water_resistance":              "",
		"network_technology":            "GSM / HSPA / LTE / 5G",
		"2g_bands":                      "GSM 850 / 900 / 1800 / 1900",
		"3g_bands":                      "HSDPA 850 / 900 / 1900 / 2100",
		"4g_bands":                      "LTE",
		"5g_bands":                      "SA/NSA/Sub6",
		"positioning_system":            "",
		"sim_card_type":                 "Hybrid Dual SIM (nano + microSD)",
		"available_colors":              "Midnight Blue, Prism Silver, Waterfall Blue",
		"device_status":                 "Available",
		"display_type":                  "Super AMOLED, 120 Hz",
		"resolution":                    "2340 × 1080 px",
		"nfc_support":                   "",
		"usb_type":                      "USB-C 2.0",
		"operating_system":              "Android 13, One UI 5 (upgradable)",
		"audio_jack":                    "",
		"cpu_type":                      "Octa-core (2×2.4 GHz + 6×2.0 GHz)",
		"ram":                           "6 / 8 GB",
		"rear_camera":                   "50 MP + 8 MP + 2 MP",
		"front_camera_video_resolution": "",
		"battery":                       "6000 mAh",
		"wireless_charging":             "No",
		"5g_support":                    "Yes",
		"sensors":                       "Accelerometer, Gyro, Proximity, Compass, fingerprint (side)",
		"storage":                       "128 / 256 GB + microSD (up to 1TB)",
		"refresh_rate":                  "120 Hz",
		"bluetooth_version":             "",
		"battery_type":                  "Li-Po (non-removable)",
		"special_features":              "",
		"audio_quality":                 "",
		"build_material":                "Glass front (Gorilla Glass 5), plastic frame / back",
		"wifi_support":                  "Wi-Fi 802.11 a/b/g/n/ac",
		"camera_features":               "OIS (main), LED flash",
		"internal_memory_capacity":      "128 / 256 GB",
		"card_slot_type":                "microSDXC (hybrid)",
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
