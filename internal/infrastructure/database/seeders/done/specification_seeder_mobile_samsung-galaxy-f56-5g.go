package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederMobileSamsungGalaxyF565G seeds specifications/options for product 'samsung-galaxy-f56-5g'
type SpecificationSeederMobileSamsungGalaxyF565G struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyF565G creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyF565G() *SpecificationSeederMobileSamsungGalaxyF565G {
	return &SpecificationSeederMobileSamsungGalaxyF565G{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy F56 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyF565G) getBanglaTranslations() map[string]string {
	common := s.GetCommonBanglaTranslations()
	specific := map[string]string{
		"Li-Ion (non-removable)":     "লি-আয়ন (অপসারণযোগ্য নয়)",
		"25 W wired":                 "२৫ ওয়াট তারযুক্ত",
		"Yes":                        "হ্যাঁ",
		"GPS, GLONASS, GALILEO, BDS": "জিপিএস, গ্লোনাস, গ্যালিলিও, বিডিএস",
		"Nano-SIM + Nano-SIM":        "ন্যানো-সিম + ন্যানো-সিম",
		"Exynos 1480":                "এক্সিনস ১৪৮০",
		"Side fingerprint, Accelerometer, Gyro, Compass": "সাইড ফিঙ্গারপ্রিন্ট, অ্যাক্সেলেরোমিটার, গাইরো, কম্পাস",
		"Super AMOLED+, 120 Hz":                          "সুপার AMOLED+, ১২০ Hz",
		"1080p @ 30fps":                                  "১০৮০p @ ৩০fps",
		"4K @ 30fps; 1080p @ 30/60fps":                   "৪কে @ ৩০fps; ১০৮০p @ ৩০/৬০fps",
		"1080 × 2340 px":                                 "১০৮০ × ২৩৪০ px",
		"120 Hz":                                         "১২০ Hz",
		"No":                                             "না",
		"50 MP + 8 MP + 2 MP":                            "৫০ মেগাপিক্সেল + ৪ মেগাপিক্সেল + ২ মেগাপিক্সেল",
		"None":                                           "কোনটি না",
		"5000 mAh":                                       "৫০০০ মিলিঅ্যাম্পিয়ার আওয়ার",
		"2025":                                           "২০২৫",
		"Glass front & back":                             "সামনে এবং পিছনে গ্লাস",
		"Not specified":                                  "উল্লেখ করা হয়নি",
		"GSM / HSPA / LTE / 5G":                          "জিএসএম / এইচএসপিএ / এলটিই / ৫জি",
		"Wi-Fi 802.11 a/b/g/n/ac":                        "ওয়াই-ফাই ৮০২.১১ a/b/g/n/ac",
		"16 MP":                                          "১৬ MP",
		"Android 15, One UI 7":                           "অ্যান্ড্রয়েড ১৫, ওয়ান UI ৭",
		"Graphite, Blue, Violet":                         "গ্রাফাইট, নীল, বেগুনি",
		"180 g":                                          "১৮০ g",
		"5.3":                                            "৫.৩",
		"Available":                                      "উপলব্ধ",
		"128 / 256 GB":                                   "১২৮ / ২৫৬ GB",
		"OIS, LED flash":                                 "ওআইএস, এলইডি ফ্ল্যাশ",
		"Corning Gorilla Glass":                          "কর্নিং গরিলা গ্লাস",
		"microSDXC (dedicated slot)":                     "মাইক্রো এসডিএক্সসি (ডেডিকেটেড স্লট)",
		"2.0 GHz":                                        "২.০ GHz",
		"GSM 850 / 900 / 1800 / 1900":                    "জিএসএম ৮৫০ / ৯০০ / ১৮০০ / ১৯০০",
		"HSDPA 850 / 900 / 1900 / 2100":                  "এইচএসডিপিএ ৮৫০ / ৯০০ / ১৯০০ / ২১০০",
		"LTE":                                            "এলটিই",
		"SA/NSA/Sub6":                                    "এসএ/এনএসএ/সাব৬",
	}
	for k, v := range specific {
		common[k] = v
	}
	return common
}

// Seed inserts specification records for the product identified by slug 'samsung-galaxy-f56-5g'
func (s *SpecificationSeederMobileSamsungGalaxyF565G) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-f56-5g"
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
		"usb_type":                      "USB-C",
		"special_features":              "",
		"cpu_type":                      "",
		"battery_type":                  "Li-Ion (non-removable)",
		"fast_charging":                 "25 W wired",
		"5g_support":                    "Yes",
		"positioning_system":            "GPS, GLONASS, GALILEO, BDS",
		"sim_card_type":                 "Nano-SIM + Nano-SIM",
		"processor":                     "Exynos 1480",
		"ram":                           "8 GB",
		"sensors":                       "Side fingerprint, Accelerometer, Gyro, Compass",
		"display_type":                  "Super AMOLED+, 120 Hz",
		"front_camera_video_resolution": "1080p @ 30fps",
		"loudspeaker_quality":           "Yes",
		"display_size":                  "6.7 inches",
		"gpu_type":                      "",
		"storage":                       "128 / 256 GB",
		"camera_features":               "OIS, LED flash",
		"camera_video_resolution":       "4K @ 30fps; 1080p @ 30/60fps",
		"resolution":                    "1080 × 2340 px",
		"refresh_rate":                  "120 Hz",
		"screen_protection":             "Corning Gorilla Glass",
		"water_resistance":              "",
		"wireless_charging":             "No",
		"audio_quality":                 "",
		"device_status":                 "Available",
		"rear_camera":                   "50 MP + 8 MP + 2 MP",
		"optical_zoom":                  "None",
		"battery":                       "5000 mAh",
		"audio_jack":                    "No",
		"announcement_date":             "2025",
		"chipset":                       "Exynos 1480",
		"build_material":                "Glass front & back",
		"dimensions":                    "Not specified",
		"network_technology":            "GSM / HSPA / LTE / 5G",
		"2g_bands":                      "GSM 850 / 900 / 1800 / 1900",
		"3g_bands":                      "HSDPA 850 / 900 / 1900 / 2100",
		"4g_bands":                      "LTE",
		"5g_bands":                      "SA/NSA/Sub6",
		"wifi_support":                  "Wi-Fi 802.11 a/b/g/n/ac",
		"front_camera":                  "16 MP",
		"operating_system":              "Android 15, One UI 7",
		"available_colors":              "Graphite, Blue, Violet",
		"weight":                        "180 g",
		"bluetooth_version":             "5.3",
		"nfc_support":                   "Yes",
		"internal_memory_capacity":      "128 / 256 GB",
		"card_slot_type":                "microSDXC (dedicated slot)",
		"processor_speed":               "2.0 GHz",
		"charging_speed":                "25 W",
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
