package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederMobileSamsungGalaxyZFlip7FE seeds specifications/options for product 'samsung-galaxy-z-flip-7-fe'
type SpecificationSeederMobileSamsungGalaxyZFlip7FE struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyZFlip7FE creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyZFlip7FE() *SpecificationSeederMobileSamsungGalaxyZFlip7FE {
	return &SpecificationSeederMobileSamsungGalaxyZFlip7FE{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy Z Flip 7 FE"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyZFlip7FE) getBanglaTranslations() map[string]string {
	common := s.GetCommonBanglaTranslations()
	specific := map[string]string{
		"Exynos 2400e":                           "এক্সিনস ২৪০০e",
		"Xclipse 940":                            "এক্সক্লিপস ৯৪০",
		"8 GB":                                   "৮ জিবি",
		"128 GB":                                 "১২৮ জিবি",
		"256 GB":                                 "২৫৬ জিবি",
		"Foldable Dynamic AMOLED 2X":             "ফোল্ডেবল ডায়নামিক অ্যামোলেড ২এক্স",
		"1080 x 2640 pixels":                     "১০৮০ × ২৬৪০ পিক্সেল",
		"120Hz":                                  "১২০হার্টজ",
		"4000 mAh":                               "৪০০০ মিলিঅ্যাম্পিয়ার আওয়ার",
		"12 MP + 12 MP":                          "১২ এমপি + ১২ এমপি",
		"10 MP":                                  "১০ এমপি",
		"Android 15, One UI 7":                   "অ্যান্ড্রয়েড ১৫, ওয়ান ইউআই ৭",
		"IPX8":                                   "আইপি এক্স৮",
		"6.7 inches (Main) / 3.4 inches (Cover)": "৬.৭ ইঞ্চি (মেইন) / ৩.৪ ইঞ্চি (কভার)",
		"190 g":                                  "১৯০ গ্রাম",
		"Unfolded: 165.1 x 71.9 x 6.9 mm":        "আনফোল্ডেড: ১৬৫.১ × ৭১.৯ × ৬.৯ মিমি",
		"Graphite, Cream, Lavender":              "গ্রাফাইট, ক্রিম, ল্যাভেন্ডার",
		"July 2025":                              "জুলাই ২০২৫",
		"Exynos 2400e (4 nm)":                    "এক্সিনস ২৪০০e (৪ ন্যানোমিটার)",
		"Deca-core":                              "ডেকা-কোর",
		"Foldable Dynamic AMOLED 2X, 120Hz":      "ফোল্ডেবল ডায়নামিক অ্যামোলেড ২এক্স, ১২০হার্টজ",
		"Corning Gorilla Glass Victus+":          "কর্নিং গরিলা গ্লাস ভিক্টাস+",
		"5G":                                     "৫জি",
		"4,000 mAh":                              "৪,০০০ মিলিঅ্যাম্পিয়ার আওয়ার",
		"Rumored":                                "রামরড",
	}
	for k, v := range specific {
		common[k] = v
	}
	return common
}

// Seed inserts specification records for the product identified by slug 'samsung-galaxy-z-flip-7-fe'
func (s *SpecificationSeederMobileSamsungGalaxyZFlip7FE) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-z-flip-7-fe"
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
		"storage":            "128 GB / 256 GB",
		"build_material":     "Aluminum frame, glass front/back",
		"available_colors":   "Graphite, Cream, Lavender",
		"announcement_date":  "July 2025",
		"chipset":            "Exynos 2400e (4 nm)",
		"cpu_type":           "Deca-core",
		"gpu_type":           "Xclipse 940",
		"operating_system":   "Android 15, One UI 7",
		"weight":             "190 g",
		"dimensions":         "Unfolded: 165.1 x 71.9 x 6.9 mm",
		"display_size":       "6.7 inches (Main) / 3.4 inches (Cover)",
		"ram":                "8 GB",
		"display_type":       "Foldable Dynamic AMOLED 2X, 120Hz",
		"resolution":         "1080 x 2640 pixels",
		"screen_protection":  "Corning Gorilla Glass Victus+",
		"refresh_rate":       "120Hz",
		"processor":          "Exynos 2400e",
		"water_resistance":   "IPX8",
		"network_technology": "5G",
		"rear_camera":        "12 MP + 12 MP",
		"front_camera":       "10 MP",
		"battery":            "4,000 mAh",
		"device_status":      "Rumored",
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
