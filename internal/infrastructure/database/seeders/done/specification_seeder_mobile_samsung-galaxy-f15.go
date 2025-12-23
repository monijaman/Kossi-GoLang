package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederMobileSamsungGalaxyF15 seeds specifications/options for product 'samsung-galaxy-f15'
type SpecificationSeederMobileSamsungGalaxyF15 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyF15 creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyF15() *SpecificationSeederMobileSamsungGalaxyF15 {
	return &SpecificationSeederMobileSamsungGalaxyF15{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy F15"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyF15) getBanglaTranslations() map[string]string {
	common := s.GetCommonBanglaTranslations()
	specific := map[string]string{
		"6.6 inches":            "৬.৬ ইঞ্চি",
		"IPS LCD, 90 Hz":        "আইপিএস এলসিডি, ৯০ হার্জ",
		"1080 × 2408 px":        "১০৮০ × ২৪০৮ পিক্সেল",
		"90 Hz":                 "৯০ হার্জ",
		"Not specified":         "উল্লেখ করা হয়নি",
		"Helio G99":             "হেলিও G99",
		"MediaTek Helio G99":    "মিডিয়াটেক হেলিও G99",
		"Octa-core":             "অক্টা-কোর",
		"Mali-G57 MC2":          "মালি-G57 MC2",
		"4 / 6 GB":              "৪ / ৬ জিবি",
		"64 / 128 GB + microSD": "৬৪ / ১২৮ জিবি + মাইক্রোএসডি",
		"64 / 128 GB":           "৬৪ / ১২৮ জিবি",
		"microSD":               "মাইক্রোএসডি",
		"50 MP + 2 MP + 2 MP":   "৫০ মেগাপিক্সেল + ২ মেগাপিক্সেল + ২ মেগাপিক্সেল",
		"50 MP (wide) + 2 MP (macro) + 2 MP (depth)": "৫০ মেগাপিক্সেল (ওয়াইড) + ২ মেগাপিক্সেল (ম্যাক্রো) + ২ মেগাপিক্সেল (ডেপথ)",
		"LED flash":                         "LED ফ্ল্যাশ",
		"1080p @ 30fps":                     "১০৮০পি @ ৩০ এফপিএস",
		"None":                              "কোনটি নয়",
		"13 MP":                             "১৩ MP",
		"5000 mAh":                          "৫০০০ মিলিঅ্যাম্পিয়ার আওয়ার",
		"Li-Ion (non-removable)":            "লি-আয়ন (অপসারণযোগ্য নয়)",
		"15 W wired":                        "১৫ ওয়াট তারযুক্ত",
		"No":                                "না",
		"Android 13, One UI Core 5":         "অ্যান্ড্রয়েড ১৩, ওয়ান UI কোর ৫",
		"Mono":                              "মনো",
		"Yes, 3.5 mm":                       "হ্যাঁ, ৩.৫ mm",
		"Plastic frame & back, glass front": "প্লাস্টিক ফ্রেম & ব্যাক, গ্লাস ফ্রন্ট",
		"195 g":                             "১৯৫ গ্রাম",
		"164 × 76 × 8.5 mm":                 "১৬৪ × ৭৬ × ৮.৫ মিলিমিটার",
		"GSM / HSPA / LTE / 5G":             "জিএসএম / এইচএসপিএ / এলটিই / ৫জি",
		"Black, Green, Blue":                "কালো, সবুজ, নীল",
		"2023":                              "২০২৩",
		"Available":                         "উপলব্ধ",
		"Wi-Fi 802.11 a/b/g/n/ac":           "ওয়াই-ফাই ৮০২.১১ এ/বি/জি/এন/এসি",
		"5.0":                               "৫.০",
		"Yes (market dependent)":            "হ্যাঁ (মার্কেট নির্ভর)",
		"USB-C":                             "ইউএসবি-সি",
		"GPS, GLONASS, GALILEO, BDS":        "জিপিএস, গ্লোনাস, গ্যালিলিও, বিডিএস",
		"Accelerometer, Gyro, Proximity, Compass": "অ্যাক্সেলেরোমিটার, জাইরো, প্রক্সিমিটি, কম্পাস",
		"Nano-SIM + Nano-SIM":                     "ন্যানো-SIM + ন্যানো-SIM",
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

// Seed inserts specification records for the product identified by slug 'samsung-galaxy-f15'
func (s *SpecificationSeederMobileSamsungGalaxyF15) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-f15"
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
		"display_size":                  "6.6 inches",
		"display_type":                  "IPS LCD, 90 Hz",
		"resolution":                    "1080 × 2408 px",
		"refresh_rate":                  "90 Hz",
		"screen_protection":             "Not specified",
		"chipset":                       "MediaTek Helio G99",
		"processor":                     "Helio G99",
		"cpu_type":                      "Octa-core",
		"gpu_type":                      "Mali-G57 MC2",
		"ram":                           "4 / 6 GB",
		"storage":                       "64 / 128 GB + microSD",
		"internal_memory_capacity":      "64 / 128 GB",
		"card_slot_type":                "microSD",
		"rear_camera":                   "50 MP + 2 MP + 2 MP",
		"quad_camera_setup":             "50 MP (wide) + 2 MP (macro) + 2 MP (depth)",
		"camera_features":               "LED flash",
		"camera_video_resolution":       "1080p @ 30fps",
		"optical_zoom":                  "None",
		"front_camera":                  "13 MP",
		"front_camera_video_resolution": "1080p @ 30fps",
		"battery":                       "5000 mAh",
		"battery_type":                  "Li-Ion (non-removable)",
		"fast_charging":                 "15 W wired",
		"charging_speed":                "15 W wired",
		"wireless_charging":             "No",
		"operating_system":              "Android 13, One UI Core 5",
		"audio_quality":                 "",
		"loudspeaker_quality":           "Mono",
		"audio_jack":                    "Yes, 3.5 mm",
		"build_material":                "Plastic frame & back, glass front",
		"weight":                        "195 g",
		"dimensions":                    "164 × 76 × 8.5 mm",
		"water_resistance":              "",
		"network_technology":            "GSM / HSPA / LTE / 5G",
		"2g_bands":                      "Not specified",
		"3g_bands":                      "Not specified",
		"4g_bands":                      "Not specified",
		"5g_bands":                      "Not specified",
		"5g_support":                    "Yes",
		"wifi_support":                  "Wi-Fi 802.11 a/b/g/n/ac",
		"bluetooth_version":             "5.0",
		"nfc_support":                   "Yes (market dependent)",
		"usb_type":                      "USB-C",
		"positioning_system":            "GPS, GLONASS, GALILEO, BDS",
		"sensors":                       "Accelerometer, Gyro, Proximity, Compass",
		"special_features":              "",
		"sim_card_type":                 "Nano-SIM + Nano-SIM",
		"sar_rating":                    "Not specified",
		"sar_rating_eu":                 "Not specified",
		"available_colors":              "Black, Green, Blue",
		"model_variants":                "Not specified",
		"announcement_date":             "2023",
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
