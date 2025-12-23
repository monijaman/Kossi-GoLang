package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederMobileVivoV50 seeds specifications/options for product 'vivo-v50'
type SpecificationSeederMobileVivoV50 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileVivoV50 creates a new seeder instance
func NewSpecificationSeederMobileVivoV50() *SpecificationSeederMobileVivoV50 {
	return &SpecificationSeederMobileVivoV50{BaseSeeder: BaseSeeder{name: "Specifications for Vivo V50"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileVivoV50) getBanglaTranslations() map[string]string {
	translations := map[string]string{
		"AMOLED, 60 Hz":                   "AMOLED ৬০Hz",
		"Android 11, Funtouch OS 11":      "অ্যান্ড্রয়েড ১১",
		"Hi-Res Audio":                    "হাই-রেজ",
		"SM7250 (7 nm)":                   "SM7250",
		"8 GB":                            "৮ জিবি",
		"32 MP":                           "৩২ মেগা",
		"128 / 256 GB":                    "১২৮/২৫৬ জিবি",
		"4315 mAh":                        "৪৩১৫ mAh",
		"Black, Blue":                     "কালো, নীল",
		"Yes":                             "হ্যাঁ",
		"No":                              "না",
		"USB-C 2.0":                       "USB-C",
		"Dual SIM (Nano-SIM)":             "ডুয়াল সিম",
		"Glass front/back, plastic frame": "গ্লাস ফ্রন্ট",
		"AMOLED display, 5G support":      "AMOLED 5G",
		"Available":                       "উপলব্ধ",
		"2376 × 1080 px":                  "২৩৭৬x১০৮০",
		"171 g":                           "১৭১ গ্রাম",
		"GSM / HSPA / LTE / 5G":           "GSM/LTE/5G",
		"5.1":                             "৫.১",
		"Fingerprint (under display), Accelerometer, Gyro, Proximity, Compass": "ফিঙ্গারপ্রিন্ট",
		"December 2020":                "ডিসেম্বর ২০২০",
		"Snapdragon 765G":              "স্ন্যাপড্রাগন ৭৬৫G",
		"Octa-core":                    "অক্টা-কোর",
		"33 W wired":                   "৩৩W",
		"Stereo":                       "স্টেরিও",
		"6.56 inches":                  "৬.৫৬ ইঞ্চি",
		"60 Hz":                        "৬০ হার্জ",
		"48 MP + 8 MP + 2 MP":          "৪৮+৮+২ মেগা",
		"None":                         "নেই",
		"Yes, 3.5 mm":                  "৩.৫মিমি",
		"Wi-Fi 802.11 a/b/g/n/ac/6":    "ওয়াই-ফাই 6",
		"LED flash, OIS (main)":        "LED OIS",
		"GPS, GLONASS, GALILEO, BDS":   "GPS",
		"Li-Ion (non-removable)":       "লি-আয়ন",
		"1080p @ 30fps":                "১০৮০p",
		"Adreno 620":                   "অ্যাড্রেনো ৬২০",
		"4K @ 30fps; 1080p @ 30/60fps": "4K ৩০fps",
	}
	for k, v := range s.GetCommonBanglaTranslations() {
		if _, exists := translations[k]; !exists {
			translations[k] = v
		}
	}
	return translations
}

// Seed inserts specification records for the product identified by slug 'vivo-v50'
func (s *SpecificationSeederMobileVivoV50) Seed(db *gorm.DB) error {
	productSlug := "vivo-v50"
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
		"display_type":                  "AMOLED, 60 Hz",
		"operating_system":              "Android 11, Funtouch OS 11",
		"wireless_charging":             "No",
		"sim_card_type":                 "Dual SIM (Nano-SIM)",
		"build_material":                "Glass front/back, plastic frame",
		"5g_support":                    "Yes",
		"special_features":              "AMOLED display, 5G support",
		"audio_quality":                 "Hi-Res Audio",
		"device_status":                 "Available",
		"chipset":                       "SM7250 (7 nm)",
		"resolution":                    "2376 × 1080 px",
		"weight":                        "171 g",
		"network_technology":            "GSM / HSPA / LTE / 5G",
		"bluetooth_version":             "5.1",
		"sensors":                       "Fingerprint (under display), Accelerometer, Gyro, Proximity, Compass",
		"ram":                           "8 GB",
		"front_camera":                  "32 MP",
		"front_camera_video_resolution": "1080p @ 30fps",
		"battery_type":                  "Li-Ion (non-removable)",
		"gpu_type":                      "Adreno 620",
		"camera_video_resolution":       "4K @ 30fps; 1080p @ 30/60fps",
		"battery":                       "4315 mAh",
		"announcement_date":             "December 2020",
		"processor":                     "Snapdragon 765G",
		"cpu_type":                      "Octa-core",
		"storage":                       "128 / 256 GB",
		"usb_type":                      "USB-C 2.0",
		"fast_charging":                 "33 W wired",
		"loudspeaker_quality":           "Stereo",
		"display_size":                  "6.56 inches",
		"refresh_rate":                  "60 Hz",
		"nfc_support":                   "Yes",
		"rear_camera":                   "48 MP + 8 MP + 2 MP",
		"optical_zoom":                  "None",
		"audio_jack":                    "Yes, 3.5 mm",
		"wifi_support":                  "Wi-Fi 802.11 a/b/g/n/ac/6",
		"camera_features":               "LED flash, OIS (main)",
		"positioning_system":            "GPS, GLONASS, GALILEO, BDS",
		"available_colors":              "Black, Blue",
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
