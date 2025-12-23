package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederMobileSamsungGalaxyF365G seeds specifications/options for product 'samsung-galaxy-f36-5g'
type SpecificationSeederMobileSamsungGalaxyF365G struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyF365G creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyF365G() *SpecificationSeederMobileSamsungGalaxyF365G {
	return &SpecificationSeederMobileSamsungGalaxyF365G{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy F36 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyF365G) getBanglaTranslations() map[string]string {
	common := s.GetCommonBanglaTranslations()
	specific := map[string]string{
		// Add specific translations here if needed
		"Exynos 1380":                       "এক্সিনস ১৩৮০",
		"Mali-G68 MP5":                      "মালি-G68 MP5",
		"Octa-core (4×2.4 GHz + 4×2.0 GHz)": "অক্টা-কোর (4×2.4 GHz + 4×2.0 GHz)",
		"6 / 8 GB":                          "৬ / ৮ GB",
		"128 / 256 GB + microSD":            "১২৮ / ২৫৬ GB + মাইক্রোSD",
		"Glass front, plastic frame, eco-leather back":   "সামনে গ্লাস, প্লাস্টিক ফ্রেম, ইকো-লেদার পিছন",
		"Wi-Fi 802.11 a/b/g/n/ac/6":                      "ওয়াই-ফাই ৮০২.১১ a/b/g/n/ac/6",
		"4K @ 30fps; 1080p @ 30fps":                      "4K @ 30fps; 1080p @ 30fps",
		"GSM / HSPA / LTE / 5G":                          "GSM / HSPA / LTE / 5G",
		"Side fingerprint, Accelerometer, Gyro, Compass": "পাশের ফিঙ্গারপ্রিন্ট, অ্যাক্সেলেরোমিটার, গাইরো, কম্পাস",
		"Yes":                                "হ্যাঁ",
		"2025":                               "২০২৫",
		"Available":                          "উপলব্ধ",
		"5000 mAh":                           "৫০০০ মিলিঅ্যাম্পিয়ার আওয়ার",
		"1080 × 2340 px":                     "১০৮০ × ২৩৪০ px",
		"197 g":                              "১৯৭ g",
		"50 MP + 8 MP + 2 MP":                "৫০ MP + ৮ MP + ২ MP",
		"Android 15, One UI 7":               "অ্যান্ড্রয়েড ১৫, ওয়ান UI ৭",
		"25 W wired":                         "২৫ W তারযুক্ত",
		"Luxe Violet, Coral Red, Onyx Black": "লুক্স ভায়োলেট, করাল রেড, ওনিক্স ব্ল্যাক",
		"Exynos 1380 (5 nm)":                 "এক্সিনস ১৩৮০ (5 nm)",
		"120 Hz":                             "১২০ Hz",
		"None":                               "কোনটি না",
		"13 MP":                              "১৩ MP",
		"No":                                 "না",
		"5.3":                                "৫.৩",
		"OIS, LED flash":                     "OIS, LED ফ্ল্যাশ",
		"4K @ 30fps; 1080p @ 30/60fps":       "4K @ 30fps; 1080p @ 30/60fps",
		"Nano-SIM + Nano-SIM":                "ন্যানো-SIM + ন্যানো-SIM",
		"164.4 × 77.9 × 7.7 mm":              "১৬৪.৪ × ৭৭.৯ × ৭.৭ mm",
		"USB-C 2.0":                          "USB-C 2.0",
		"GPS, GLONASS, GALILEO, BDS, QZSS":   "জিপিএস, গ্লোনাস, গ্যালিলিও, বিডিএস, কিউজেএসএস",
	}
	for k, v := range specific {
		common[k] = v
	}
	return common
}

// Seed inserts specification records for the product identified by slug 'samsung-galaxy-f36-5g'
func (s *SpecificationSeederMobileSamsungGalaxyF365G) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-f36-5g"
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
		"sim_card_type":                 "Nano-SIM + Nano-SIM",
		"announcement_date":             "2025",
		"processor":                     "Exynos 1380",
		"cpu_type":                      "Octa-core (4×2.4 GHz + 4×2.0 GHz)",
		"display_type":                  "Super AMOLED, 120 Hz",
		"weight":                        "197 g",
		"battery_type":                  "Li-Ion (non-removable)",
		"loudspeaker_quality":           "Yes",
		"device_status":                 "Available",
		"water_resistance":              "",
		"wifi_support":                  "Wi-Fi 802.11 a/b/g/n/ac/6",
		"rear_camera":                   "50 MP + 8 MP + 2 MP",
		"camera_features":               "OIS, LED flash",
		"sensors":                       "Side fingerprint, Accelerometer, Gyro, Compass",
		"audio_quality":                 "",
		"storage":                       "128 / 256 GB + microSD",
		"resolution":                    "1080 × 2340 px",
		"front_camera_video_resolution": "4K @ 30fps; 1080p @ 30fps",
		"wireless_charging":             "No",
		"special_features":              "",
		"audio_jack":                    "No",
		"dimensions":                    "164.4 × 77.9 × 7.7 mm",
		"fast_charging":                 "25 W wired",
		"gpu_type":                      "Mali-G68 MP5",
		"build_material":                "Glass front, plastic frame, eco-leather back",
		"usb_type":                      "USB-C 2.0",
		"positioning_system":            "GPS, GLONASS, GALILEO, BDS, QZSS",
		"display_size":                  "6.7 inches",
		"nfc_support":                   "Yes",
		"operating_system":              "Android 15, One UI 7",
		"battery":                       "5000 mAh",
		"chipset":                       "Exynos 1380 (5 nm)",
		"ram":                           "6 / 8 GB",
		"optical_zoom":                  "None",
		"front_camera":                  "13 MP",
		"5g_support":                    "Yes",
		"available_colors":              "Luxe Violet, Coral Red, Onyx Black",
		"refresh_rate":                  "120 Hz",
		"network_technology":            "GSM / HSPA / LTE / 5G",
		"bluetooth_version":             "5.3",
		"camera_video_resolution":       "4K @ 30fps; 1080p @ 30/60fps",
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
