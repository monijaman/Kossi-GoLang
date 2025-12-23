package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederMobileSamsungGalaxyZFlip7 seeds specifications/options for product 'samsung-galaxy-z-flip-7'
type SpecificationSeederMobileSamsungGalaxyZFlip7 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyZFlip7 creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyZFlip7() *SpecificationSeederMobileSamsungGalaxyZFlip7 {
	return &SpecificationSeederMobileSamsungGalaxyZFlip7{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy Z Flip 7"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyZFlip7) getBanglaTranslations() map[string]string {
	common := s.GetCommonBanglaTranslations()
	specific := map[string]string{
		"Snapdragon 8 Elite":                     "স্ন্যাপড্রাগন ৮ এলিট",
		"Adreno 830":                             "অ্যাড্রেনো ৮৩০",
		"8 GB":                                   "৮ জিবি",
		"12 GB":                                  "১২ জিবি",
		"256 GB":                                 "২৫৬ জিবি",
		"512 GB":                                 "৫১২ জিবি",
		"Foldable Dynamic LTPO AMOLED 2X":        "ফোল্ডেবল ডায়নামিক এলটিপিও অ্যামোলেড ২এক্স",
		"1080 x 2640 pixels (Main)":              "১০৮০ × ২৬৪০ পিক্সেল (মেইন)",
		"120Hz":                                  "১২০হার্টজ",
		"4200 mAh":                               "৪২০০ মিলিঅ্যাম্পিয়ার আওয়ার",
		"50 MP + 12 MP":                          "৫০ এমপি + ১২ এমপি",
		"12 MP":                                  "১২ এমপি",
		"Android 15, One UI 7.1.1":               "অ্যান্ড্রয়েড ১৫, ওয়ান ইউআই ৭.১.১",
		"IPX8":                                   "আইপি এক্স৮",
		"6.7 inches (Main) / 3.9 inches (Cover)": "৬.৭ ইঞ্চি (মেইন) / ৩.৯ ইঞ্চি (কভার)",
		"187 g":                                  "১৮৭ গ্রাম",
		"Unfolded: 165.1 x 71.9 x 6.9 mm / Folded: 85.1 x 71.9 x 14.9 mm": "আনফোল্ডেড: ১৬৫.১ × ৭১.৯ × ৬.৯ মিমি / ফোল্ডেড: ৮৫.১ × ৭১.৯ × ১৪.৯ মিমি",
		"Mint, Graphite, Cream, Lavender":                                 "মিন্ট, গ্রাফাইট, ক্রিম, ল্যাভেন্ডার",
		"Corning Gorilla Armor 2":                                         "কর্নিং গরিলা আর্মর ২",
		"Glass front/back, aluminum frame":                                "গ্লাস ফ্রন্ট/ব্যাক, অ্যালুমিনিয়াম ফ্রেম",
		"Rumored":                                                         "রামরড",
		"8 GB / 12 GB":                                                    "৮ জিবি / ১২ জিবি",
		"5G":                                                              "৫জি",
		"July 2025":                                                       "জুলাই ২০২৫",
		"Qualcomm SM8750-AC Snapdragon 8 Elite (3 nm)": "কোয়ালকম এসএম৮৭৫০-এসি স্ন্যাপড্রাগন ৮ এলিট (৩ ন্যানোমিটার)",
		"Octa-core": "অক্টা-কোর",
		"Foldable Dynamic LTPO AMOLED 2X, 120Hz, HDR10+": "ফোল্ডেবল ডায়নামিক এলটিপিও অ্যামোলেড ২এক্স, ১২০হার্টজ, এইচডিআর১০+",
		"4,200 mAh": "৪,২০০ মিলিঅ্যাম্পিয়ার আওয়ার",
	}
	for k, v := range specific {
		common[k] = v
	}
	return common
}

// Seed inserts specification records for the product identified by slug 'samsung-galaxy-z-flip-7'
func (s *SpecificationSeederMobileSamsungGalaxyZFlip7) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-z-flip-7"
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
		"storage":            "256 GB / 512 GB",
		"resolution":         "1080 x 2640 pixels (Main)",
		"screen_protection":  "Corning Gorilla Armor 2",
		"build_material":     "Glass front/back, aluminum frame",
		"device_status":      "Rumored",
		"refresh_rate":       "120Hz",
		"rear_camera":        "50 MP + 12 MP",
		"front_camera":       "12 MP",
		"processor":          "Snapdragon 8 Elite",
		"gpu_type":           "Adreno 830",
		"ram":                "8 GB / 12 GB",
		"network_technology": "5G",
		"available_colors":   "Mint, Graphite, Cream, Lavender",
		"announcement_date":  "July 2025",
		"chipset":            "Qualcomm SM8750-AC Snapdragon 8 Elite (3 nm)",
		"cpu_type":           "Octa-core",
		"display_type":       "Foldable Dynamic LTPO AMOLED 2X, 120Hz, HDR10+",
		"weight":             "187 g",
		"dimensions":         "Unfolded: 165.1 x 71.9 x 6.9 mm / Folded: 85.1 x 71.9 x 14.9 mm",
		"water_resistance":   "IPX8",
		"battery":            "4,200 mAh",
		"operating_system":   "Android 15, One UI 7.1.1",
		"display_size":       "6.7 inches (Main) / 3.9 inches (Cover)",
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
