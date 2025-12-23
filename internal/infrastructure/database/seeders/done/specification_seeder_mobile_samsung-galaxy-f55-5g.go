package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederMobileSamsungGalaxyF555G seeds specifications/options for product 'samsung-galaxy-f55-5g'
type SpecificationSeederMobileSamsungGalaxyF555G struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyF555G creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyF555G() *SpecificationSeederMobileSamsungGalaxyF555G {
	return &SpecificationSeederMobileSamsungGalaxyF555G{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy F55 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyF555G) getBanglaTranslations() map[string]string {
	common := s.GetCommonBanglaTranslations()
	specific := map[string]string{
		// Add specific translations here if needed
		"MediaTek Dimensity 6080":   "মিডিয়াটেক ডাইমেনসিটি ৬০८०",
		"Mali-G57 MC3":              "মালি-জি৫৭ এমসি৩",
		"Octa-core":                 "অক্টা-কোর",
		"6 / 8 GB":                  "৬ / ৮ জিবি",
		"128 / 256 GB + microSD":    "১২৮ / ২৫৬ জিবি + মাইক্রোএসডি",
		"Glass front, plastic back": "সামনে গ্লাস, পেছনে প্লাস্টিক",
		"Wi-Fi 802.11 a/b/g/n/ac":   "ওয়াই-ফাই ৮০২.১১ a/b/g/n/ac",
		"1080p @ 30fps":             "১০৮০পি @ ৩০ এফপিএস",
		"GSM / HSPA / LTE / 5G":     "জিএসএম / এইচএসপিএ / এলটিই / ৫জি",
		"Side fingerprint, Accelerometer, Gyro, Compass": "পাশের ফিঙ্গারপ্রিন্ট, অ্যাক্সেলেরোমিটার, জাইরো, কম্পাস",
		"Yes":                             "হ্যাঁ",
		"2024":                            "২০২৪",
		"Available":                       "উপলব্ধ",
		"5000 mAh":                        "৫০০০ মিলিঅ্যাম্পিয়ার আওয়ার",
		"1080 × 2408 px":                  "১০৮০ × ২৪০৮ px",
		"188 g":                           "১৮৮ g",
		"64 MP + 5 MP + 2 MP":             "৬৪ MP + ৫ MP + ২ MP",
		"Android 14, One UI 6":            "অ্যান্ড্রয়েড ১৪, ওয়ান UI ৬",
		"25 W wired":                      "२৫ ওয়াট তারযুক্ত",
		"Black, Silver, Blue":             "ব্ল্যাক, সিলভার, ব্লু",
		"Dimensity 6080":                  "ডাইমেনসিটি ৬০৮০",
		"120 Hz":                          "১২০ Hz",
		"None":                            "কোনটি না",
		"13 MP":                           "১৩ MP",
		"No":                              "না",
		"5.2":                             "৫.২",
		"OIS, LED flash":                  "অপটিক্যাল ইমেজ স্ট্যাবিলাইজেশন, এলইডি ফ্ল্যাশ",
		"1080p @ 30/60fps":                "১০৮০p @ ৩০/৬০fps",
		"Nano-SIM + Nano-SIM":             "ন্যানো-সিম + ন্যানো-সিম",
		"164 × 76 × 8.3 mm":               "১৬৪ × ৭৬ × ৮.৩ mm",
		"USB-C":                           "ইউএসবি-সি",
		"GPS, GLONASS, GALILEO, BDS":      "জিপিএস, গ্লোনাস, গ্যালিলিও, বিডিএস",
		"Unfolded: 165.1 x 71.9 x 6.9 mm": "আনফোল্ডেড: ১৬৫.১ × ৭১.৯ × ৬.৯ মিমি",
	}
	for k, v := range specific {
		common[k] = v
	}
	return common
}

// Seed inserts specification records for the product identified by slug 'samsung-galaxy-f55-5g'
func (s *SpecificationSeederMobileSamsungGalaxyF555G) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-f55-5g"
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
		"refresh_rate":                  "120 Hz",
		"usb_type":                      "USB-C",
		"front_camera_video_resolution": "1080p @ 30fps",
		"fast_charging":                 "25 W wired",
		"special_features":              "",
		"sim_card_type":                 "Nano-SIM + Nano-SIM",
		"processor":                     "MediaTek Dimensity 6080",
		"build_material":                "Glass front, plastic back",
		"bluetooth_version":             "5.2",
		"5g_support":                    "Yes",
		"chipset":                       "Dimensity 6080",
		"resolution":                    "1080 × 2408 px",
		"weight":                        "188 g",
		"wifi_support":                  "Wi-Fi 802.11 a/b/g/n/ac",
		"sensors":                       "Side fingerprint, Accelerometer, Gyro, Compass",
		"loudspeaker_quality":           "Yes",
		"audio_quality":                 "",
		"audio_jack":                    "No",
		"ram":                           "6 / 8 GB",
		"network_technology":            "GSM / HSPA / LTE / 5G",
		"camera_features":               "OIS, LED flash",
		"optical_zoom":                  "None",
		"wireless_charging":             "No",
		"positioning_system":            "GPS, GLONASS, GALILEO, BDS",
		"available_colors":              "Black, Silver, Blue",
		"gpu_type":                      "Mali-G57 MC3",
		"storage":                       "128 / 256 GB + microSD",
		"display_type":                  "Super AMOLED, 120 Hz",
		"water_resistance":              "",
		"operating_system":              "Android 14, One UI 6",
		"battery_type":                  "Li-Ion (non-removable)",
		"device_status":                 "Available",
		"cpu_type":                      "Octa-core",
		"rear_camera":                   "64 MP + 5 MP + 2 MP",
		"nfc_support":                   "Yes",
		"camera_video_resolution":       "1080p @ 30/60fps",
		"front_camera":                  "13 MP",
		"announcement_date":             "2024",
		"display_size":                  "6.6 inches",
		"dimensions":                    "164 × 76 × 8.3 mm",
		"battery":                       "5000 mAh",
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
