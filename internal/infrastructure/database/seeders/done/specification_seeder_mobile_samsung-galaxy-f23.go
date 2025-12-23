package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederMobileSamsungGalaxyF23 seeds specifications/options for product 'samsung-galaxy-f23'
type SpecificationSeederMobileSamsungGalaxyF23 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyF23 creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyF23() *SpecificationSeederMobileSamsungGalaxyF23 {
	return &SpecificationSeederMobileSamsungGalaxyF23{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy F23"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyF23) getBanglaTranslations() map[string]string {
	common := s.GetCommonBanglaTranslations()
	specific := map[string]string{
		// Add specific translations here if needed
		"Snapdragon 750G":                 "স্ন্যাপড্রাগন ৭৫০জি",
		"Adreno 619":                      "অ্যাড্রেনো ৬১৯",
		"Octa-core":                       "অক্টা-কোর",
		"4 / 6 / 8 GB":                    "৪ / ৬ / ৮ GB",
		"128 GB + microSD":                "১২৮ GB + মাইক্রোSD",
		"Glass front, plastic frame/back": "সামনে গ্লাস, প্লাস্টিক ফ্রেম/পিছন",
		"Wi-Fi 802.11 a/b/g/n/ac":         "ওয়াই-ফাই ৮০২.১১ a/b/g/n/ac",
		"1080p @ 30/60fps":                "১০৮০p @ ৩০/৬০fps",
		"GSM / HSPA / LTE / 5G":           "GSM / HSPA / LTE / 5G",
		"Side fingerprint, Accelerometer, Gyro, Compass": "পাশের ফিঙ্গারপ্রিন্ট, অ্যাক্সেলেরোমিটার, গাইরো, কম্পাস",
		"Stereo":                     "স্টেরিও",
		"IP67":                       "IP67",
		"Yes":                        "হ্যাঁ",
		"2023":                       "২০২৩",
		"Available":                  "উপলব্ধ",
		"5000 mAh":                   "৫০০০ mAh",
		"1080 × 2408 px":             "১০৮০ × ২৪০৮ px",
		"198 g":                      "১৯৮ g",
		"50 MP + 8 MP + 2 MP":        "৫০ MP + ৮ MP + ২ MP",
		"Android 13, One UI 5.1":     "অ্যান্ড্রয়েড ১৩, ওয়ান UI ৫.১",
		"25 W wired":                 "২৫ W তারযুক্ত",
		"Black, Green, Copper":       "কালো, সবুজ, তামা",
		"120 Hz":                     "১২০ Hz",
		"None":                       "কোনটি না",
		"13 MP":                      "১৩ MP",
		"Yes, 3.5 mm":                "হ্যাঁ, ৩.৫ mm",
		"5.0":                        "৫.০",
		"OIS":                        "OIS",
		"Nano-SIM + Nano-SIM":        "ন্যানো-SIM + ন্যানো-SIM",
		"164 × 76 × 8.4 mm":          "১৬৪ × ৭৬ × ৮.৪ mm",
		"USB-C":                      "USB-C",
		"No":                         "না",
		"GPS, GLONASS, GALILEO, BDS": "GPS, GLONASS, GALILEO, BDS",
	}
	for k, v := range specific {
		common[k] = v
	}
	return common
}

// Seed inserts specification records for the product identified by slug 'samsung-galaxy-f23'
func (s *SpecificationSeederMobileSamsungGalaxyF23) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-f23"
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
		"camera_video_resolution":       "1080p @ 30/60fps",
		"audio_quality":                 "",
		"battery_type":                  "Li-Ion (non-removable)",
		"display_size":                  "6.6 inches",
		"processor":                     "Snapdragon 750G",
		"cpu_type":                      "Octa-core",
		"storage":                       "128 GB + microSD",
		"build_material":                "Glass front, plastic frame/back",
		"wifi_support":                  "Wi-Fi 802.11 a/b/g/n/ac",
		"front_camera_video_resolution": "1080p @ 30fps",
		"network_technology":            "GSM / HSPA / LTE / 5G",
		"sensors":                       "Side fingerprint, Accelerometer, Gyro, Compass",
		"loudspeaker_quality":           "Stereo",
		"gpu_type":                      "Adreno 619",
		"water_resistance":              "",
		"nfc_support":                   "Yes",
		"announcement_date":             "2023",
		"device_status":                 "Available",
		"battery":                       "5000 mAh",
		"resolution":                    "1080 × 2408 px",
		"weight":                        "198 g",
		"rear_camera":                   "50 MP + 8 MP + 2 MP",
		"operating_system":              "Android 13, One UI 5.1",
		"fast_charging":                 "25 W wired",
		"5g_support":                    "Yes",
		"special_features":              "",
		"ram":                           "4 / 6 / 8 GB",
		"display_type":                  "IPS LCD, 120 Hz",
		"dimensions":                    "164 × 76 × 8.4 mm",
		"usb_type":                      "USB-C",
		"wireless_charging":             "No",
		"positioning_system":            "GPS, GLONASS, GALILEO, BDS",
		"sim_card_type":                 "Nano-SIM + Nano-SIM",
		"available_colors":              "Black, Green, Copper",
		"chipset":                       "Snapdragon 750G",
		"refresh_rate":                  "120 Hz",
		"optical_zoom":                  "None",
		"front_camera":                  "13 MP",
		"audio_jack":                    "Yes, 3.5 mm",
		"bluetooth_version":             "5.0",
		"camera_features":               "OIS",
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
