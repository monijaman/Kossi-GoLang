package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederMobileVivoV50Lite seeds specifications/options for product 'vivo-v50-lite'
type SpecificationSeederMobileVivoV50Lite struct {
	BaseSeeder
}

// NewSpecificationSeederMobileVivoV50Lite creates a new seeder instance
func NewSpecificationSeederMobileVivoV50Lite() *SpecificationSeederMobileVivoV50Lite {
	return &SpecificationSeederMobileVivoV50Lite{BaseSeeder: BaseSeeder{name: "Specifications for Vivo V50 Lite"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileVivoV50Lite) getBanglaTranslations() map[string]string {
	translations := map[string]string{
		"No":                         "না",
		"Yes":                        "হ্যাঁ",
		"Black, Blue":                "কালো, নীল",
		"128 GB":                     "১২৮ জিবি",
		"6 GB":                       "৬ জিবি",
		"16 MP":                      "১৬ মেগা",
		"GPS, GLONASS, GALILEO, BDS": "GPS",
		"AMOLED-like display":        "AMOLED ডিসপ্লে",
		"Android 11, Funtouch OS 11": "অ্যান্ড্রয়েড ১১",
		"4315 mAh":                   "৪৩১৫ mAh",
		"Wi-Fi 802.11 a/b/g/n/ac":    "ওয়াই-ফাই",
		"1080p @ 30fps":              "১০৮০p",
		"Dual SIM (Nano-SIM)":        "ডুয়াল সিম",
		"18 W wired":                 "১৮W",
		"Mono":                       "মোনো",
		"6.44 inches":                "৬.৪৪ ইঞ্চি",
		"Octa-core":                  "অক্টা-কোর",
		"180 g":                      "১৮০ গ্রাম",
		"Fingerprint (side-mounted), Accelerometer, Gyro, Proximity, Compass": "ফিঙ্গারপ্রিন্ট",
		"Yes, 3.5 mm":                     "৩.৫মিমি",
		"60 Hz":                           "৬০ হার্জ",
		"USB-C 2.0":                       "USB-C",
		"Standard":                        "স্ট্যান্ডার্ড",
		"48 MP + 2 MP":                    "৪৮+২ মেগা",
		"None":                            "নেই",
		"Li-Ion (non-removable)":          "লি-আয়ন",
		"Available":                       "উপলব্ধ",
		"Dimensity 720 (7 nm)":            "ডাইমেন্সিটি ৭২০",
		"Glass front, plastic back/frame": "গ্লাস ফ্রন্ট",
		"IPS LCD, 60 Hz":                  "IPS LCD",
		"2400 × 1080 px":                  "২৪০০x১০৮০",
		"January 2021":                    "জানুয়ারি ২০২১",
		"MediaTek Dimensity 720":          "মিডিয়াটেক ৭২০",
		"Mali-G57 MC3":                    "Mali-G57",
		"GSM / HSPA / LTE":                "GSM/LTE",
		"5.1":                             "৫.১",
		"LED flash":                       "LED ফ্ল্যাশ",
	}
	for k, v := range s.GetCommonBanglaTranslations() {
		if _, exists := translations[k]; !exists {
			translations[k] = v
		}
	}
	return translations
}

// Seed inserts specification records for the product identified by slug 'vivo-v50-lite'
func (s *SpecificationSeederMobileVivoV50Lite) Seed(db *gorm.DB) error {
	productSlug := "vivo-v50-lite"
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
		"wireless_charging":             "No",
		"loudspeaker_quality":           "Mono",
		"available_colors":              "Black, Blue",
		"storage":                       "128 GB",
		"wifi_support":                  "Wi-Fi 802.11 a/b/g/n/ac",
		"front_camera":                  "16 MP",
		"positioning_system":            "GPS, GLONASS, GALILEO, BDS",
		"camera_features":               "LED flash",
		"special_features":              "AMOLED-like display",
		"display_size":                  "6.44 inches",
		"cpu_type":                      "Octa-core",
		"weight":                        "180 g",
		"5g_support":                    "No",
		"sensors":                       "Fingerprint (side-mounted), Accelerometer, Gyro, Proximity, Compass",
		"sim_card_type":                 "Dual SIM (Nano-SIM)",
		"audio_jack":                    "Yes, 3.5 mm",
		"refresh_rate":                  "60 Hz",
		"usb_type":                      "USB-C 2.0",
		"fast_charging":                 "18 W wired",
		"audio_quality":                 "Standard",
		"ram":                           "6 GB",
		"rear_camera":                   "48 MP + 2 MP",
		"optical_zoom":                  "None",
		"battery_type":                  "Li-Ion (non-removable)",
		"device_status":                 "Available",
		"chipset":                       "Dimensity 720 (7 nm)",
		"build_material":                "Glass front, plastic back/frame",
		"nfc_support":                   "No",
		"operating_system":              "Android 11, Funtouch OS 11",
		"display_type":                  "IPS LCD, 60 Hz",
		"resolution":                    "2400 × 1080 px",
		"camera_video_resolution":       "1080p @ 30fps",
		"front_camera_video_resolution": "1080p @ 30fps",
		"battery":                       "4315 mAh",
		"announcement_date":             "January 2021",
		"processor":                     "MediaTek Dimensity 720",
		"gpu_type":                      "Mali-G57 MC3",
		"network_technology":            "GSM / HSPA / LTE",
		"bluetooth_version":             "5.1",
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
