package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederMobileVivoX300Pro seeds specifications/options for product 'vivo-x300-pro'
type SpecificationSeederMobileVivoX300Pro struct {
	BaseSeeder
}

// NewSpecificationSeederMobileVivoX300Pro creates a new seeder instance
func NewSpecificationSeederMobileVivoX300Pro() *SpecificationSeederMobileVivoX300Pro {
	return &SpecificationSeederMobileVivoX300Pro{BaseSeeder: BaseSeeder{name: "Specifications for Vivo X300 Pro"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileVivoX300Pro) getBanglaTranslations() map[string]string {
	translations := map[string]string{
		"6.78 inches":                      "6.78 ইঞ্চি",
		"Adreno 660":                       "Adreno 660",
		"AMOLED, 120 Hz":                   "AMOLED 120Hz",
		"Yes":                              "হ্যাঁ",
		"2x + 5x":                          "2x + 5x",
		"32 MP":                            "32 মেগা",
		"4500 mAh":                         "4500 mAh",
		"HDR10+, AMOLED display":           "HDR10+ AMOLED",
		"Glass front/back, aluminum frame": "গ্লাস ফ্রন্ট",
		"4K @ 30/60fps; 1080p @ 30/120fps": "4K 30fps",
		"Android 12, Funtouch OS 12":       "অ্যান্ড্রয়েড 12",
		"Fingerprint (under display), Accelerometer, Gyro, Proximity, Compass": "ফিঙ্গারপ্রিন্ট",
		"Hi-Res Audio":               "হাই-রেজ",
		"GSM / HSPA / LTE / 5G":      "GSM/LTE/5G",
		"Snapdragon 888+ (5 nm)":     "স্ন্যাপড্রাগন 888+",
		"128 / 256 GB":               "128/256 জিবি",
		"120 Hz":                     "120 হার্জ",
		"Dual SIM (Nano-SIM)":        "ডুয়াল সিম",
		"204 g":                      "204 গ্রাম",
		"OIS, LED flash":             "LED ফ্ল্যাশ",
		"No":                         "না",
		"Wi-Fi 802.11 a/b/g/n/ac/6":  "ওয়াই-ফাই",
		"8 / 12 GB":                  "8/12 জিবি",
		"50 MP + 13 MP + 13 MP":      "50+13+13 মেগা",
		"GPS, GLONASS, GALILEO, BDS": "GPS",
		"Black, Blue, Silver":        "কালো, নীল, সিলভার",
		"March 2023":                 "মার্চ 2023",
		"66 W wired":                 "66W",
		"Stereo":                     "স্টেরিও",
		"Available":                  "উপলব্ধ",
		"Octa-core":                  "অক্টা-কোর",
		"1080p @ 30fps":              "1080p",
		"Snapdragon 888+":            "স্ন্যাপড্রাগন 888+",
		"3200 × 1440 px":             "3200x1440",
		"5.2":                        "5.2",
		"USB-C 3.1":                  "USB-C",
		"Li-Ion (non-removable)":     "লি-আয়ন",
	}
	for k, v := range s.GetCommonBanglaTranslations() {
		if _, exists := translations[k]; !exists {
			translations[k] = v
		}
	}
	return translations
}

// Seed inserts specification records for the product identified by slug 'vivo-x300-pro'
func (s *SpecificationSeederMobileVivoX300Pro) Seed(db *gorm.DB) error {
	productSlug := "vivo-x300-pro"
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
		"display_size":                  "6.78 inches",
		"gpu_type":                      "Adreno 660",
		"display_type":                  "AMOLED, 120 Hz",
		"nfc_support":                   "Yes",
		"optical_zoom":                  "2x + 5x",
		"front_camera":                  "32 MP",
		"battery":                       "4500 mAh",
		"special_features":              "HDR10+, AMOLED display",
		"build_material":                "Glass front/back, aluminum frame",
		"camera_video_resolution":       "4K @ 30/60fps; 1080p @ 30/120fps",
		"operating_system":              "Android 12, Funtouch OS 12",
		"sensors":                       "Fingerprint (under display), Accelerometer, Gyro, Proximity, Compass",
		"audio_quality":                 "Hi-Res Audio",
		"network_technology":            "GSM / HSPA / LTE / 5G",
		"chipset":                       "Snapdragon 888+ (5 nm)",
		"storage":                       "128 / 256 GB",
		"refresh_rate":                  "120 Hz",
		"sim_card_type":                 "Dual SIM (Nano-SIM)",
		"weight":                        "204 g",
		"camera_features":               "OIS, LED flash",
		"wireless_charging":             "No",
		"wifi_support":                  "Wi-Fi 802.11 a/b/g/n/ac/6",
		"ram":                           "8 / 12 GB",
		"rear_camera":                   "50 MP + 13 MP + 13 MP",
		"positioning_system":            "GPS, GLONASS, GALILEO, BDS",
		"audio_jack":                    "No",
		"available_colors":              "Black, Blue, Silver",
		"announcement_date":             "March 2023",
		"fast_charging":                 "66 W wired",
		"5g_support":                    "Yes",
		"loudspeaker_quality":           "Stereo",
		"device_status":                 "Available",
		"cpu_type":                      "Octa-core",
		"front_camera_video_resolution": "1080p @ 30fps",
		"processor":                     "Snapdragon 888+",
		"resolution":                    "3200 × 1440 px",
		"bluetooth_version":             "5.2",
		"usb_type":                      "USB-C 3.1",
		"battery_type":                  "Li-Ion (non-removable)",
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
