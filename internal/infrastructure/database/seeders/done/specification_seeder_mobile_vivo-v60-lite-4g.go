package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederMobileVivoV60Lite4G seeds specifications/options for product 'vivo-v60-lite-4g'
type SpecificationSeederMobileVivoV60Lite4G struct {
	BaseSeeder
}

// NewSpecificationSeederMobileVivoV60Lite4G creates a new seeder instance
func NewSpecificationSeederMobileVivoV60Lite4G() *SpecificationSeederMobileVivoV60Lite4G {
	return &SpecificationSeederMobileVivoV60Lite4G{BaseSeeder: BaseSeeder{name: "Specifications for Vivo V60 Lite 4G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileVivoV60Lite4G) getBanglaTranslations() map[string]string {
	translations := map[string]string{
		"1080p @ 30fps":                   "১০৮০p",
		"Glass front/back, plastic frame": "গ্লাস ফ্রন্ট",
		"128 GB":                          "১২৮ জিবি",
		"5.0":                             "৫.০",
		"18 W wired":                      "১৮W",
		"Available":                       "উপলব্ধ",
		"6 / 8 GB":                        "৬/৮ জিবি",
		"Black, Blue":                     "কালো, নীল",
		"No":                              "না",
		"Yes":                             "হ্যাঁ",
		"GSM / HSPA / LTE (4G)":           "GSM/LTE",
		"6.44 inches":                     "৬.৪৪ ইঞ্চি",
		"GPS, GLONASS, GALILEO, BDS":      "GPS",
		"LED flash":                       "LED ফ্ল্যাশ",
		"None":                            "নেই",
		"16 MP":                           "১৬ মেগা",
		"AMOLED display":                  "AMOLED ডিসপ্লে",
		"Wi-Fi 802.11 a/b/g/n/ac":         "ওয়াই-ফাই",
		"Mali-G76 MC4":                    "Mali-G76",
		"AMOLED, 60 Hz":                   "AMOLED ৬০Hz",
		"175 g":                           "১৭৫ গ্রাম",
		"64 MP + 2 MP":                    "৬৪+২ মেগা",
		"Stereo":                          "স্টেরিও",
		"Hi-Res Audio":                    "হাই-রেজ",
		"MediaTek Helio G95":              "হেলিও G95",
		"4100 mAh":                        "৪১০০ mAh",
		"Dual SIM (Nano-SIM)":             "ডুয়াল সিম",
		"2400 × 1080 px":                  "২৪০০x১০৮০",
		"60 Hz":                           "৬০ হার্জ",
		"USB-C 2.0":                       "USB-C",
		"Android 11, Funtouch OS 11":      "অ্যান্ড্রয়েড ১১",
		"Yes, 3.5 mm":                     "৩.৫মিমি",
		"Helio G95 (12 nm)":               "হেলিও G95",
		"Octa-core":                       "অক্টা-কোর",
		"Fingerprint (under display), Accelerometer, Gyro, Proximity, Compass": "ফিঙ্গারপ্রিন্ট",
		"May 2021":               "মে ২০২১",
		"Li-Ion (non-removable)": "লি-আয়ন",
	}
	for k, v := range s.GetCommonBanglaTranslations() {
		if _, exists := translations[k]; !exists {
			translations[k] = v
		}
	}
	return translations
}

// Seed inserts specification records for the product identified by slug 'vivo-v60-lite-4g'
func (s *SpecificationSeederMobileVivoV60Lite4G) Seed(db *gorm.DB) error {
	productSlug := "vivo-v60-lite-4g"
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
		"front_camera_video_resolution": "1080p @ 30fps",
		"build_material":                "Glass front/back, plastic frame",
		"network_technology":            "GSM / HSPA / LTE (4G)",
		"storage":                       "128 GB",
		"bluetooth_version":             "5.0",
		"battery_type":                  "Li-Ion (non-removable)",
		"fast_charging":                 "18 W wired",
		"device_status":                 "Available",
		"display_size":                  "6.44 inches",
		"camera_video_resolution":       "1080p @ 30fps",
		"positioning_system":            "GPS, GLONASS, GALILEO, BDS",
		"camera_features":               "LED flash",
		"optical_zoom":                  "None",
		"front_camera":                  "16 MP",
		"special_features":              "AMOLED display",
		"wifi_support":                  "Wi-Fi 802.11 a/b/g/n/ac",
		"gpu_type":                      "Mali-G76 MC4",
		"display_type":                  "AMOLED, 60 Hz",
		"weight":                        "175 g",
		"nfc_support":                   "No",
		"rear_camera":                   "64 MP + 2 MP",
		"loudspeaker_quality":           "Stereo",
		"audio_quality":                 "Hi-Res Audio",
		"processor":                     "MediaTek Helio G95",
		"ram":                           "6 / 8 GB",
		"battery":                       "4100 mAh",
		"sim_card_type":                 "Dual SIM (Nano-SIM)",
		"available_colors":              "Black, Blue",
		"resolution":                    "2400 × 1080 px",
		"refresh_rate":                  "60 Hz",
		"usb_type":                      "USB-C 2.0",
		"operating_system":              "Android 11, Funtouch OS 11",
		"wireless_charging":             "No",
		"audio_jack":                    "Yes, 3.5 mm",
		"chipset":                       "Helio G95 (12 nm)",
		"cpu_type":                      "Octa-core",
		"5g_support":                    "No",
		"sensors":                       "Fingerprint (under display), Accelerometer, Gyro, Proximity, Compass",
		"announcement_date":             "May 2021",
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
