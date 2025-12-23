package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederMobileVivoX200ProMini seeds specifications/options for product 'vivo-x200-pro-mini'
type SpecificationSeederMobileVivoX200ProMini struct {
	BaseSeeder
}

// NewSpecificationSeederMobileVivoX200ProMini creates a new seeder instance
func NewSpecificationSeederMobileVivoX200ProMini() *SpecificationSeederMobileVivoX200ProMini {
	return &SpecificationSeederMobileVivoX200ProMini{BaseSeeder: BaseSeeder{name: "Specifications for Vivo X200 Pro Mini"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileVivoX200ProMini) getBanglaTranslations() map[string]string {
	translations := map[string]string{
		"4K @ 30fps; 1080p @ 30fps":  "4K ৩০fps",
		"2x":                         "২x",
		"Li-Ion (non-removable)":     "লি-আয়ন",
		"Yes":                        "হ্যাঁ",
		"AMOLED display":             "AMOLED ডিসপ্লে",
		"MediaTek Dimensity 1100":    "মিডিয়াটেক ১১০০",
		"Dimensity 1100 (6 nm)":      "ডাইমেন্সিটি ১১০০",
		"90 Hz":                      "৯০ হার্জ",
		"128 GB":                     "১২৮ জিবি",
		"Wi-Fi 802.11 a/b/g/n/ac":    "ওয়াই-ফাই",
		"OIS, LED flash":             "OIS LED",
		"No":                         "না",
		"6.44 inches":                "৬.৪৪ ইঞ্চি",
		"6 / 8 GB":                   "৬/৮ জিবি",
		"Dual SIM (Nano-SIM)":        "ডুয়াল সিম",
		"Mali-G77 MC9":               "Mali-G77",
		"AMOLED, 90 Hz":              "AMOLED ৯০Hz",
		"185 g":                      "১৮৫ গ্রাম",
		"GSM / HSPA / LTE / 5G":      "GSM/LTE/5G",
		"USB-C 2.0":                  "USB-C",
		"Android 11, Funtouch OS 11": "অ্যান্ড্রয়েড ১১",
		"2400 × 1080 px":             "২৪০০x১০৮০",
		"4000 mAh":                   "৪০০০ mAh",
		"Fingerprint (under display), Accelerometer, Gyro, Proximity, Compass": "ফিঙ্গারপ্রিন্ট",
		"Black, Blue":                     "কালো, নীল",
		"5.2":                             "৫.২",
		"32 MP":                           "৩২ মেগা",
		"1080p @ 30fps":                   "১০৮০p",
		"GPS, GLONASS, GALILEO, BDS":      "GPS",
		"Stereo":                          "স্টেরিও",
		"Octa-core":                       "অক্টা-কোর",
		"Glass front/back, plastic frame": "গ্লাস ফ্রন্ট",
		"48 MP + 8 MP + 2 MP":             "৪৮+৮+২ মেগা",
		"33 W wired":                      "৩৩W",
		"Hi-Res Audio":                    "হাই-রেজ",
		"November 2022":                   "নভেম্বর ২০২২",
		"Available":                       "উপলব্ধ",
	}
	for k, v := range s.GetCommonBanglaTranslations() {
		if _, exists := translations[k]; !exists {
			translations[k] = v
		}
	}
	return translations
}

// Seed inserts specification records for the product identified by slug 'vivo-x200-pro-mini'
func (s *SpecificationSeederMobileVivoX200ProMini) Seed(db *gorm.DB) error {
	productSlug := "vivo-x200-pro-mini"
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
		"camera_video_resolution":       "4K @ 30fps; 1080p @ 30fps",
		"optical_zoom":                  "2x",
		"battery_type":                  "Li-Ion (non-removable)",
		"5g_support":                    "Yes",
		"special_features":              "AMOLED display",
		"processor":                     "MediaTek Dimensity 1100",
		"chipset":                       "Dimensity 1100 (6 nm)",
		"refresh_rate":                  "90 Hz",
		"storage":                       "128 GB",
		"wifi_support":                  "Wi-Fi 802.11 a/b/g/n/ac",
		"nfc_support":                   "Yes",
		"camera_features":               "OIS, LED flash",
		"wireless_charging":             "No",
		"display_size":                  "6.44 inches",
		"ram":                           "6 / 8 GB",
		"sim_card_type":                 "Dual SIM (Nano-SIM)",
		"audio_jack":                    "No",
		"gpu_type":                      "Mali-G77 MC9",
		"display_type":                  "AMOLED, 90 Hz",
		"weight":                        "185 g",
		"network_technology":            "GSM / HSPA / LTE / 5G",
		"usb_type":                      "USB-C 2.0",
		"operating_system":              "Android 11, Funtouch OS 11",
		"resolution":                    "2400 × 1080 px",
		"battery":                       "4000 mAh",
		"sensors":                       "Fingerprint (under display), Accelerometer, Gyro, Proximity, Compass",
		"available_colors":              "Black, Blue",
		"bluetooth_version":             "5.2",
		"front_camera":                  "32 MP",
		"front_camera_video_resolution": "1080p @ 30fps",
		"positioning_system":            "GPS, GLONASS, GALILEO, BDS",
		"loudspeaker_quality":           "Stereo",
		"cpu_type":                      "Octa-core",
		"build_material":                "Glass front/back, plastic frame",
		"rear_camera":                   "48 MP + 8 MP + 2 MP",
		"fast_charging":                 "33 W wired",
		"audio_quality":                 "Hi-Res Audio",
		"announcement_date":             "November 2022",
		"device_status":                 "Available",
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
