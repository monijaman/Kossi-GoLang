package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederMobileSamsungGalaxyZFold7 seeds specifications/options for product 'samsung-galaxy-z-fold-7'
type SpecificationSeederMobileSamsungGalaxyZFold7 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyZFold7 creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyZFold7() *SpecificationSeederMobileSamsungGalaxyZFold7 {
	return &SpecificationSeederMobileSamsungGalaxyZFold7{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy Z Fold 7"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyZFold7) getBanglaTranslations() map[string]string {
	common := s.GetCommonBanglaTranslations()
	specific := map[string]string{
		"Snapdragon 8 Elite":                     "স্ন্যাপড্রাগন ৮ এলিট",
		"Adreno 830":                             "অ্যাড্রেনো ৮৩০",
		"12 GB":                                  "১২ জিবি",
		"16 GB":                                  "১৬ জিবি",
		"256 GB":                                 "২৫৬ জিবি",
		"512 GB":                                 "৫১২ জিবি",
		"1 TB":                                   "১ টিবি",
		"Foldable Dynamic LTPO AMOLED 2X":        "ফোল্ডেবল ডায়নামিক এলটিপিও অ্যামোলেড ২এক্স",
		"1856 x 2160 pixels (Main)":              "১৮৫৬ × ২১৬০ পিক্সেল (মেইন)",
		"120Hz":                                  "১২০হার্টজ",
		"4600 mAh":                               "৪৬০০ মিলিঅ্যাম্পিয়ার আওয়ার",
		"50 MP + 12 MP + 10 MP":                  "৫০ এমপি + ১২ এমপি + ১০ এমপি",
		"4 MP (Under Display) + 10 MP (Cover)":   "৪ এমপি (আন্ডার ডিসপ্লে) + ১০ এমপি (কভার)",
		"Android 15, One UI 7.1.1":               "অ্যান্ড্রয়েড ১৫, ওয়ান ইউআই ৭.১.১",
		"IPX8":                                   "আইপি এক্স৮",
		"7.6 inches (Main) / 6.3 inches (Cover)": "৭.৬ ইঞ্চি (মেইন) / ৬.৩ ইঞ্চি (কভার)",
		"235 g":                                  "২৩৫ গ্রাম",
		"Unfolded: 153.5 x 132.6 x 5.6 mm / Folded: 153.5 x 68.1 x 11.2 mm": "আনফোল্ডেড: ১৫৩.৫ × ১৩২.৬ × ৫.৬ মিমি / ফোল্ডেড: ১৫৩.৫ × ৬৮.১ × ১১.২ মিমি",
		"Phantom Black, Icy Blue, Cream":                                    "ফ্যান্টম ব্ল্যাক, আইসি ব্লু, ক্রিম",
		"Foldable Dynamic LTPO AMOLED 2X, 120Hz, HDR10+":                    "ফোল্ডেবল ডায়নামিক এলটিপিও অ্যামোলেড ২এক্স, ১২০হার্টজ, এইচডিআর১০+",
		"Corning Gorilla Armor 2":                                           "কর্নিং গরিলা আর্মর ২",
		"4,600 mAh":                                                         "৪,৬০০ মিলিঅ্যাম্পিয়ার আওয়ার",
		"Rumored":                                                           "রামরড",
		"12 GB / 16 GB":                                                     "১২ জিবি / ১৬ জিবি",
		"256 GB / 512 GB / 1 TB":                                            "২৫৬ জিবি / ৫১২ জিবি / ১ টিবি",
		"Qualcomm SM8750-AC Snapdragon 8 Elite (3 nm)":                      "কোয়ালকম এসএম৮৭৫০-এসি স্ন্যাপড্রাগন ৮ এলিট (৩ ন্যানোমিটার)",
		"Octa-core":                                                         "অক্টা-কোর",
		"5G":                                                                "৫জি",
		"July 2025":                                                         "জুলাই ২০২৫",
		"Glass front (Gorilla Armor 2), glass back (Victus 2), titanium frame": "গ্লাস ফ্রন্ট (গরিলা আর্মর ২), গ্লাস ব্যাক (ভিক্টাস ২), টাইটানিয়াম ফ্রেম",
		"Unfolded: 165.1 x 71.9 x 6.9 mm":                                      "আনফোল্ডেড: ১৬৫.১ × ৭১.৯ × ৬.৯ মিমি",
	}
	for k, v := range specific {
		common[k] = v
	}
	return common
}

// Seed inserts specification records for the product identified by slug 'samsung-galaxy-z-fold-7'
func (s *SpecificationSeederMobileSamsungGalaxyZFold7) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-z-fold-7"
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
		"available_colors":   "Phantom Black, Icy Blue, Cream",
		"display_type":       "Foldable Dynamic LTPO AMOLED 2X, 120Hz, HDR10+",
		"screen_protection":  "Corning Gorilla Armor 2",
		"refresh_rate":       "120Hz",
		"battery":            "4,600 mAh",
		"device_status":      "Rumored",
		"display_size":       "7.6 inches (Main) / 6.3 inches (Cover)",
		"processor":          "Snapdragon 8 Elite",
		"gpu_type":           "Adreno 830",
		"ram":                "12 GB / 16 GB",
		"storage":            "256 GB / 512 GB / 1 TB",
		"resolution":         "1856 x 2160 pixels (Main)",
		"operating_system":   "Android 15, One UI 7.1.1",
		"cpu_type":           "Octa-core",
		"water_resistance":   "IPX8",
		"network_technology": "5G",
		"announcement_date":  "July 2025",
		"weight":             "235 g",
		"dimensions":         "Unfolded: 153.5 x 132.6 x 5.6 mm / Folded: 153.5 x 68.1 x 11.2 mm",
		"chipset":            "Qualcomm SM8750-AC Snapdragon 8 Elite (3 nm)",
		"build_material":     "Glass front (Gorilla Armor 2), glass back (Victus 2), titanium frame",
		"rear_camera":        "50 MP + 12 MP + 10 MP",
		"front_camera":       "4 MP (Under Display) + 10 MP (Cover)",
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
