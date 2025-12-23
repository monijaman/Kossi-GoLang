package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederMobileSamsungGalaxyA565G seeds specifications/options for product 'samsung-galaxy-a56-5g'
type SpecificationSeederMobileSamsungGalaxyA565G struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyA565G creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyA565G() *SpecificationSeederMobileSamsungGalaxyA565G {
	return &SpecificationSeederMobileSamsungGalaxyA565G{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy A56 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyA565G) getBanglaTranslations() map[string]string {
	common := s.GetCommonBanglaTranslations()
	specific := map[string]string{
		"6.7 inches":                            "৬.৭ ইঞ্চি",
		"Super AMOLED, 120Hz":                   "সুপার অ্যামোলেড, ১২০ হার্জ",
		"1080 x 2400 pixels (~393 ppi density)": "১০৮০ × ২৪০০ পিক্সেল (~৩৯৩ পিপিআই)",
		"120Hz":                                 "১২০ হার্জ",
		"Corning Gorilla Glass 5":               "কর্নিং গরিলা গ্লাস ৫",
		"Samsung Exynos 1380 (5 nm)":            "স্যামসাং এক্সিনস ১৩৮০ (৫ ন্যানোমিটার)",
		"Exynos 1380":                           "এক্সিনস ১৩৮০",
		"Octa-core (4x2.4 GHz Cortex-A78 & 4x2.0 GHz Cortex-A55)": "অক্টা-কোর (৪x২.৪ GHz কর্টেক্স-A78 & ৪x২.০ GHz কর্টেক্স-A55)",
		"Mali-G68 MP5":         "মালি-G68 MP5",
		"8 GB":                 "৮ জিবি",
		"128 GB / 256 GB":      "১২৮ জিবি / ২৫৬ জিবি",
		"None":                 "কোনটি নয়",
		"50 MP + 12 MP + 5 MP": "৫০ MP + ১২ MP + ৫ MP",
		"50 MP (wide) + 12 MP (ultrawide) + 5 MP (macro)": "৫০ MP (ওয়াইড) + ১২ MP (আল্ট্রাওয়াইড) + ৫ MP (ম্যাক্রো)",
		"32 MP":                "৩২ মেগাপিক্সেল",
		"5000 mAh":             "৫০০০ মিলিঅ্যাম্পিয়ার আওয়ার",
		"Android 14, One UI 6": "অ্যান্ড্রয়েড ১৪, ওয়ান UI ৬",
		"Glass front, plastic frame, plastic back":        "গ্লাস ফ্রন্ট, প্লাস্টিক ফ্রেম, প্লাস্টিক ব্যাক",
		"202 g (7.13 oz)":                                 "২০২ গ্রাম (৭.১৩ আউন্স)",
		"165.1 x 76.4 x 8.4 mm":                           "১৬৫.১ × ৭৬.৪ × ৮.৪ মিলিমিটার",
		"IP67 dust/water resistant (up to 1m for 30 min)": "আইপি৬৭ ধুলো/পানি প্রতিরোধী (১ মিটার পর্যন্ত ৩০ মিনিট)",
		"GSM / HSPA / LTE / 5G":                           "জিএসএম / এইচএসপিএ / এলটিই / ৫জি",
		"Black, White, Blue":                              "ব্ল্যাক, হোয়াইট, ব্লু",
		"October 2025":                                    "অক্টোবর ২০২৫",
		"Upcoming":                                        "আসন্ন",
		"Not specified":                                   "উল্লেখ করা হয়নি",
	}
	result := make(map[string]string)
	for k, v := range common {
		result[k] = v
	}
	for k, v := range specific {
		result[k] = v
	}
	return result
}

// Seed inserts specification records for the product identified by slug 'samsung-galaxy-a56-5g'
func (s *SpecificationSeederMobileSamsungGalaxyA565G) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-a56-5g"
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
		"display_size":                  "6.7 inches",
		"display_type":                  "Super AMOLED, 120Hz",
		"resolution":                    "1080 x 2400 pixels (~393 ppi density)",
		"refresh_rate":                  "120Hz",
		"screen_protection":             "Corning Gorilla Glass 5",
		"chipset":                       "Samsung Exynos 1380 (5 nm)",
		"processor":                     "Exynos 1380",
		"cpu_type":                      "Octa-core (4x2.4 GHz Cortex-A78 & 4x2.0 GHz Cortex-A55)",
		"gpu_type":                      "Mali-G68 MP5",
		"ram":                           "8 GB",
		"storage":                       "128 GB / 256 GB",
		"internal_memory_capacity":      "128 GB / 256 GB",
		"card_slot_type":                "None",
		"rear_camera":                   "50 MP + 12 MP + 5 MP",
		"quad_camera_setup":             "50 MP (wide) + 12 MP (ultrawide) + 5 MP (macro)",
		"camera_features":               "Not specified",
		"camera_video_resolution":       "Not specified",
		"optical_zoom":                  "None",
		"front_camera":                  "32 MP",
		"front_camera_video_resolution": "Not specified",
		"battery":                       "5000 mAh",
		"battery_type":                  "Not specified",
		"fast_charging":                 "Not specified",
		"charging_speed":                "Not specified",
		"wireless_charging":             "Not specified",
		"operating_system":              "Android 14, One UI 6",
		"audio_quality":                 "Not specified",
		"loudspeaker_quality":           "Not specified",
		"audio_jack":                    "Not specified",
		"build_material":                "Glass front, plastic frame, plastic back",
		"weight":                        "202 g (7.13 oz)",
		"dimensions":                    "165.1 x 76.4 x 8.4 mm",
		"water_resistance":              "IP67 dust/water resistant (up to 1m for 30 min)",
		"network_technology":            "GSM / HSPA / LTE / 5G",
		"2g_bands":                      "Not specified",
		"3g_bands":                      "Not specified",
		"4g_bands":                      "Not specified",
		"5g_bands":                      "Not specified",
		"5g_support":                    "Not specified",
		"wifi_support":                  "Not specified",
		"bluetooth_version":             "Not specified",
		"nfc_support":                   "Not specified",
		"usb_type":                      "Not specified",
		"positioning_system":            "Not specified",
		"sensors":                       "Not specified",
		"special_features":              "Not specified",
		"sim_card_type":                 "Not specified",
		"sar_rating":                    "Not specified",
		"sar_rating_eu":                 "Not specified",
		"available_colors":              "Black, White, Blue",
		"model_variants":                "Not specified",
		"announcement_date":             "October 2025",
		"device_status":                 "Upcoming",
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
