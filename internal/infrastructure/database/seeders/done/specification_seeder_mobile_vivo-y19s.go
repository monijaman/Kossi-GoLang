package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederMobileVivoY19s seeds specifications/options for product 'vivo-y19s'
type SpecificationSeederMobileVivoY19s struct {
	BaseSeeder
}

// NewSpecificationSeederMobileVivoY19s creates a new seeder instance
func NewSpecificationSeederMobileVivoY19s() *SpecificationSeederMobileVivoY19s {
	return &SpecificationSeederMobileVivoY19s{BaseSeeder: BaseSeeder{name: "Specifications for Vivo Y19s"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileVivoY19s) getBanglaTranslations() map[string]string {
	translations := map[string]string{
		"6.68 inches":       "6.68 ইঞ্চি",
		"Unisoc T612":       "Unisoc T612",
		"720 x 1608 pixels": "720x1608",
		"Glass front, plastic back, plastic frame": "গ্লাস ফ্রন্ট",
		"5 MP":                    "5 মেগা",
		"15W wired":               "15W",
		"Available":               "উপলব্ধ",
		"Unisoc T612 (12 nm)":     "Unisoc T612",
		"198 g":                   "198 গ্রাম",
		"GSM / HSPA / LTE":        "GSM/LTE",
		"50 MP + 0.08 MP":         "50+0.08 মেগা",
		"5500 mAh":                "5500 mAh",
		"Mali-G57":                "Mali-G57",
		"6 GB":                    "6 জিবি",
		"IP64":                    "IP64",
		"October 2024":            "অক্টোবর 2024",
		"Octa-core":               "অক্টা-কোর",
		"128 GB":                  "128 জিবি",
		"IPS LCD, 90Hz":           "IPS LCD 90Hz",
		"90Hz":                    "90Hz",
		"Android 14, Funtouch 14": "অ্যান্ড্রয়েড 14",
		"Pearl Silver, Glacier Blue, Diamond Black": "পার্ল সিলভার, গ্লেসিয়ার ব্লু, ডায়মন্ড ব্ল্যাক",
	}
	for k, v := range s.GetCommonBanglaTranslations() {
		if _, exists := translations[k]; !exists {
			translations[k] = v
		}
	}
	return translations
}

// Seed inserts specification records for the product identified by slug 'vivo-y19s'
func (s *SpecificationSeederMobileVivoY19s) Seed(db *gorm.DB) error {
	productSlug := "vivo-y19s"
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
		"display_size":       "6.68 inches",
		"processor":          "Unisoc T612",
		"resolution":         "720 x 1608 pixels",
		"build_material":     "Glass front, plastic back, plastic frame",
		"front_camera":       "5 MP",
		"fast_charging":      "15W wired",
		"device_status":      "Available",
		"chipset":            "Unisoc T612 (12 nm)",
		"weight":             "198 g",
		"network_technology": "GSM / HSPA / LTE",
		"rear_camera":        "50 MP + 0.08 MP",
		"battery":            "5500 mAh",
		"gpu_type":           "Mali-G57",
		"ram":                "6 GB",
		"water_resistance":   "IP64",
		"announcement_date":  "October 2024",
		"cpu_type":           "Octa-core",
		"storage":            "128 GB",
		"display_type":       "IPS LCD, 90Hz",
		"refresh_rate":       "90Hz",
		"operating_system":   "Android 14, Funtouch 14",
		"available_colors":   "Pearl Silver, Glacier Blue, Diamond Black",
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
