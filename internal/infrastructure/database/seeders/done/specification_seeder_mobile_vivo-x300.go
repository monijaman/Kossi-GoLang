package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederMobileVivoX300 seeds specifications/options for product 'vivo-x300'
type SpecificationSeederMobileVivoX300 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileVivoX300 creates a new seeder instance
func NewSpecificationSeederMobileVivoX300() *SpecificationSeederMobileVivoX300 {
	return &SpecificationSeederMobileVivoX300{BaseSeeder: BaseSeeder{name: "Specifications for Vivo X300"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileVivoX300) getBanglaTranslations() map[string]string {
	common := s.GetCommonBanglaTranslations()
	specific := map[string]string{
		"AMOLED, 120 Hz":                   "AMOLED 120Hz",
		"USB-C 3.1":                        "USB-C",
		"4K @ 30/60fps; 1080p @ 30/120fps": "4K 30fps",
		"Stereo":                           "স্টেরিও",
		"Black, Blue, Silver":              "কালো, নীল, সিলভার",
		"6.78 inches":                      "6.78 ইঞ্চি",
		"8 / 12 GB":                        "8/12 জিবি",
		"January 2023":                     "জানুয়ারি 2023",
		"2x + 5x":                          "2x + 5x",
		"Snapdragon 870":                   "স্ন্যাপড্রাগন 870",
		"Snapdragon 870 (7 nm)":            "স্ন্যাপড্রাগন 870",
		"Adreno 650":                       "Adreno 650",
		"GSM / HSPA / LTE / 5G":            "GSM/LTE/5G",
		"Yes":                              "হ্যাঁ",
		"OIS, LED flash":                   "LED ফ্ল্যাশ",
		"Hi-Res Audio":                     "হাই-রেজ",
		"Android 12, Funtouch OS 12":       "অ্যান্ড্রয়েড 12",
		"4400 mAh":                         "4400 mAh",
		"No":                               "না",
		"HDR10+, AMOLED display":           "HDR10+ AMOLED",
		"3200 × 1440 px":                   "3200x1440",
		"32 MP":                            "32 মেগা",
		"1080p @ 30fps":                    "1080p",
		"GPS, GLONASS, GALILEO, BDS":       "GPS",
		"Available":                        "উপলব্ধ",
		"5.2":                              "5.2",
		"48 MP + 13 MP + 13 MP":            "48+13+13 মেগা",
		"Fingerprint (under display), Accelerometer, Gyro, Proximity, Compass": "ফিঙ্গারপ্রিন্ট",
		"Dual SIM (Nano-SIM)":              "ডুয়াল সিম",
		"128 / 256 GB":                     "128/256 জিবি",
		"Wi-Fi 802.11 a/b/g/n/ac/6":        "ওয়াই-ফাই",
		"66 W wired":                       "66W",
		"Octa-core":                        "অক্টা-কোর",
		"120 Hz":                           "120 হার্জ",
		"Glass front/back, aluminum frame": "গ্লাস ফ্রন্ট",
		"202 g":                            "202 গ্রাম",
		"Li-Ion (non-removable)":           "লি-আয়ন",
	}
	for k, v := range specific {
		common[k] = v
	}
	return common
}

// Seed inserts specification records for the product identified by slug 'vivo-x300'
func (s *SpecificationSeederMobileVivoX300) Seed(db *gorm.DB) error {
	productSlug := "vivo-x300"
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
		"chipset":                       "Snapdragon 870 (7 nm)",
		"network_technology":            "GSM / HSPA / LTE / 5G",
		"nfc_support":                   "Yes",
		"camera_features":               "OIS, LED flash",
		"audio_quality":                 "Hi-Res Audio",
		"display_size":                  "6.78 inches",
		"usb_type":                      "USB-C 3.1",
		"operating_system":              "Android 12, Funtouch OS 12",
		"battery":                       "4400 mAh",
		"wireless_charging":             "No",
		"special_features":              "HDR10+, AMOLED display",
		"resolution":                    "3200 × 1440 px",
		"front_camera":                  "32 MP",
		"available_colors":              "Black, Blue, Silver",
		"ram":                           "8 / 12 GB",
		"front_camera_video_resolution": "1080p @ 30fps",
		"5g_support":                    "Yes",
		"positioning_system":            "GPS, GLONASS, GALILEO, BDS",
		"loudspeaker_quality":           "Stereo",
		"processor":                     "Snapdragon 870",
		"display_type":                  "AMOLED, 120 Hz",
		"camera_video_resolution":       "4K @ 30/60fps; 1080p @ 30/120fps",
		"device_status":                 "Available",
		"bluetooth_version":             "5.2",
		"rear_camera":                   "48 MP + 13 MP + 13 MP",
		"sensors":                       "Fingerprint (under display), Accelerometer, Gyro, Proximity, Compass",
		"sim_card_type":                 "Dual SIM (Nano-SIM)",
		"audio_jack":                    "No",
		"announcement_date":             "January 2023",
		"storage":                       "128 / 256 GB",
		"wifi_support":                  "Wi-Fi 802.11 a/b/g/n/ac/6",
		"fast_charging":                 "66 W wired",
		"cpu_type":                      "Octa-core",
		"gpu_type":                      "Adreno 650",
		"refresh_rate":                  "120 Hz",
		"build_material":                "Glass front/back, aluminum frame",
		"weight":                        "202 g",
		"optical_zoom":                  "2x + 5x",
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
