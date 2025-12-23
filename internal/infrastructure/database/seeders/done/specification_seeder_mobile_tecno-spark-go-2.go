package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederMobileTecnoSPARKGo2 seeds specifications/options for product 'tecno-spark-go-2'
type SpecificationSeederMobileTecnoSPARKGo2 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileTecnoSPARKGo2 creates a new seeder instance
func NewSpecificationSeederMobileTecnoSPARKGo2() *SpecificationSeederMobileTecnoSPARKGo2 {
	return &SpecificationSeederMobileTecnoSPARKGo2{BaseSeeder: BaseSeeder{name: "Specifications for Tecno SPARK Go 2"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileTecnoSPARKGo2) getBanglaTranslations() map[string]string {
	commonTranslations := s.GetCommonBanglaTranslations()
	tecnoTranslations := map[string]string{
		"Helio A22":                        "হেলিও A22",
		"Mediatek Helio A22 (12 nm)":       "মিডিয়াটেক হেলিও A22 (১২ nm)",
		"PowerVR GE8320":                   "পাওয়ারভিআর GE8320",
		"Ice Jadeite, Aqua Blue":           "আইস জেডাইট, অ্যাকোয়া ব্লু",
		"Android 10 Go Edition":            "অ্যান্ড্রয়েড ১০ (গো এডিশন)",
		"2020":                             "২০২০",
		"Quad-core":                        "কোয়াড-কোর",
		"13 MP + AI":                       "১৩ মেগাপিক্সেল + এআই",
		"2 GB":                             "২ জিবি",
		"32 GB":                            "৩২ জিবি",
		"188 g":                            "১৮৮ গ্রাম",
		"6.52 inches":                      "৬.৫২ ইঞ্চি",
		"164.5 x 76 x 9 mm":                "১৬৪.৫ × ৭৬ × ৯ মিলিমিটার",
		"720 x 1600 pixels":                "৭২০ × ১৬০০ পিক্সেল",
		"60Hz":                             "৬০ হার্জ",
		"5,000 mAh":                        "৫,০০০ মিলিঅ্যাম্পিয়ার-আওয়ার",
		"8 MP":                             "৮ মেগাপিক্সেল",
		"Glass front":                      "গ্লাস ফ্রন্ট",
		"Plastic body":                     "প্লাস্টিক বডি",
		"No":                               "না",
		"4G":                               "৪জি",
		"Available":                        "উপলব্ধ",
		"Wi-Fi, Bluetooth 5.0, USB-C, OTG": "ওয়াই-ফাই, ব্লুটুথ ৫.০, USB-C, OTG",
		"Side fingerprint, accelerometer, proximity, compass": "সাইড ফিঙ্গারপ্রিন্ট, অ্যাক্সেলারোমিটার, প্রক্সিমিটি, কম্পাস",
	}
	for key, value := range tecnoTranslations {
		commonTranslations[key] = value
	}
	return commonTranslations
}

// Seed inserts specification records for the product identified by slug 'tecno-spark-go-2'
func (s *SpecificationSeederMobileTecnoSPARKGo2) Seed(db *gorm.DB) error {
	productSlug := "tecno-spark-go-2"
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
		"processor":          "Helio A22",
		"chipset":            "Mediatek Helio A22 (12 nm)",
		"cpu_type":           "Quad-core",
		"screen_protection":  "Glass front",
		"rear_camera":        "13 MP + AI",
		"ram":                "2 GB",
		"storage":            "32 GB",
		"refresh_rate":       "60Hz",
		"weight":             "188 g",
		"front_camera":       "8 MP",
		"battery":            "5,000 mAh",
		"device_status":      "Available",
		"display_size":       "6.52 inches",
		"gpu_type":           "PowerVR GE8320",
		"dimensions":         "164.5 x 76 x 9 mm",
		"water_resistance":   "No",
		"available_colors":   "Ice Jadeite, Aqua Blue",
		"announcement_date":  "2020",
		"display_type":       "IPS LCD",
		"resolution":         "720 x 1600 pixels",
		"build_material":     "Plastic body",
		"network_technology": "4G",
		"operating_system":   "Android 10 Go Edition",
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
