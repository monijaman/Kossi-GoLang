package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederMobileSamsungGalaxyF345G seeds specifications/options for product 'samsung-galaxy-f34-5g'
type SpecificationSeederMobileSamsungGalaxyF345G struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyF345G creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyF345G() *SpecificationSeederMobileSamsungGalaxyF345G {
	return &SpecificationSeederMobileSamsungGalaxyF345G{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy F34 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyF345G) getBanglaTranslations() map[string]string {
	common := s.GetCommonBanglaTranslations()
	specific := map[string]string{
		// Add specific translations here if needed
		"Dimensity 1080":            "ডাইমেনসিটি ১০৮০",
		"Mali-G68 MC4":              "মালি-জিম এমকৃ",
		"Octa-core":                 "অক্টা-কোর",
		"6 / 8 GB":                  "৬ / ৮ জিবি",
		"128 / 256 GB + microSD":    "১২৮ / ২৫৬ GB + মাইক্রোSD",
		"Glass front, plastic back": "সামনে গ্লাস, পিছনে প্লাস্টিক",
		"Wi-Fi 802.11 a/b/g/n/ac":   "ওয়াই-ফাই ৮০২.১১ a/b/g/n/ac",
		"1080p @ 30fps":             "১০৮০p @ ৩০fps",
		"GSM / HSPA / LTE / 5G":     "জিএসএম / এইচএসপিএ / এলটিই / ৫জি",
		"Side fingerprint, Accelerometer, Gyro, Compass": "পাশের ফিঙ্গারপ্রিন্ট, অ্যাক্সেলেরোমিটার, গাইরো, কম্পাস",
		"Stereo":                 "স্টেরিও",
		"IP67":                   "আইপিখ৭",
		"Yes":                    "হ্যাঁ",
		"2024":                   "২০২৪",
		"Available":              "উপলব্ধ",
		"5000 mAh":               "৫০০০ মিলিঅ্যাম্পিয়ার আওয়ার",
		"1080 × 2340 px":         "১০৮০ × ২৩৪০ px",
		"199 g":                  "১৯৯ g",
		"48 MP + 8 MP + 5 MP":    "৪৮ মেগাপিক্সেল + ৪ মেগাপিক্সেল + ৫ মেগাপিক্সেল",
		"Android 13, One UI 5.1": "অ্যান্ড্রয়েড ১৩, ওয়ান UI ৫.১",
		"25 W wired":             "२৫ ওয়াট তারযুক্ত",
		"Awesome Lime, Awesome Violet, Awesome Graphite": "অসাম লাইম, অসাম ভায়োলেট, অসাম গ্রাফাইট",
		"MediaTek Dimensity 1080":                        "মিডিয়াটেক ডাইমেনসিটি ১০৮০",
		"120 Hz":                                         "१२० হার্জ",
		"None":                                           "কোনটি না",
		"13 MP":                                          "१३ মেগাপিক্সেল",
		"No":                                             "না",
		"5.3":                                            "৫.৩",
		"OIS":                                            "অপটিক্যাল ইমেজ স্ট্যাবিলাইজেশন",
		"4K @ 30fps; 1080p @ 30/60fps":                   "४কে @ ३०fps; १०८०p @ ३०/६०fps",
		"Nano-SIM + Nano-SIM":                            "ন্যানো-সিম + ন্যানো-সিম",
		"161 × 78 × 8.2 mm":                              "১৬১ × ৭৮ × ৮.২ mm",
		"USB-C":                                          "ইউএসবি-সি",
		"GPS, GLONASS, GALILEO, BDS":                     "জিপিএস, গ্লোনাস, গ্যালিলিও, বিডিএস",
	}
	for k, v := range specific {
		common[k] = v
	}
	return common
}

// Seed inserts specification records for the product identified by slug 'samsung-galaxy-f34-5g'
func (s *SpecificationSeederMobileSamsungGalaxyF345G) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-f34-5g"
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
		"resolution":                    "1080 × 2340 px",
		"water_resistance":              "IP67",
		"nfc_support":                   "Yes",
		"wireless_charging":             "No",
		"cpu_type":                      "Octa-core",
		"bluetooth_version":             "5.3",
		"positioning_system":            "GPS, GLONASS, GALILEO, BDS",
		"audio_quality":                 "",
		"audio_jack":                    "No",
		"announcement_date":             "2024",
		"device_status":                 "Available",
		"chipset":                       "Dimensity 1080",
		"weight":                        "199 g",
		"usb_type":                      "USB-C",
		"optical_zoom":                  "None",
		"battery_type":                  "Li-Ion (non-removable)",
		"network_technology":            "GSM / HSPA / LTE / 5G",
		"front_camera_video_resolution": "1080p @ 30fps",
		"rear_camera":                   "48 MP + 8 MP + 5 MP",
		"camera_features":               "OIS",
		"camera_video_resolution":       "4K @ 30fps; 1080p @ 30/60fps",
		"5g_support":                    "Yes",
		"display_size":                  "6.6 inches",
		"gpu_type":                      "Mali-G68 MC4",
		"wifi_support":                  "Wi-Fi 802.11 a/b/g/n/ac",
		"front_camera":                  "13 MP",
		"operating_system":              "Android 13, One UI 5.1",
		"battery":                       "5000 mAh",
		"available_colors":              "Awesome Lime, Awesome Violet, Awesome Graphite",
		"processor":                     "MediaTek Dimensity 1080",
		"refresh_rate":                  "120 Hz",
		"fast_charging":                 "25 W wired",
		"sensors":                       "Side fingerprint, Accelerometer, Gyro, Compass",
		"special_features":              "",
		"loudspeaker_quality":           "Stereo",
		"ram":                           "6 / 8 GB",
		"storage":                       "128 / 256 GB + microSD",
		"display_type":                  "Super AMOLED, 120 Hz",
		"build_material":                "Glass front, plastic back",
		"dimensions":                    "161 × 78 × 8.2 mm",
		"sim_card_type":                 "Nano-SIM + Nano-SIM",
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
